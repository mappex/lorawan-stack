// Copyright © 2018 The Things Network Foundation, The Things Industries B.V.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package udp

import (
	"encoding/base64"
	"fmt"
	"math"
	"strconv"
	"strings"
	"time"

	"go.thethings.network/lorawan-stack/pkg/ttnpb"
	"go.thethings.network/lorawan-stack/pkg/types"
	"go.thethings.network/lorawan-stack/pkg/version"
)

const (
	delta = 0.001 // For GPS comparisons
	lora  = "LORA"
	fsk   = "FSK"

	eirpDelta int32 = 2
)

var (
	ttnVersions = map[string]string{
		"ttn-lw-gateway-server": version.TTN,
	}
	invalidLocations = []ttnpb.Location{
		{Latitude: 0.0, Longitude: 0.0},
		{Latitude: 10.0, Longitude: 20.0},
	}
)

func validLocation(loc ttnpb.Location) bool {
	for _, invalidLoc := range invalidLocations {
		if (loc.Latitude > invalidLoc.Latitude-delta && loc.Latitude < invalidLoc.Latitude+delta) &&
			(loc.Longitude > invalidLoc.Longitude-delta && loc.Longitude < invalidLoc.Longitude+delta) {
			return false
		}
	}

	return true
}

// UpstreamMetadata related to an uplink.
type UpstreamMetadata struct {
	ID ttnpb.GatewayIdentifiers
	IP string
}

// ToGatewayUp converts the UDP message to a gateway upstream message.
func ToGatewayUp(data Data, md UpstreamMetadata) (*ttnpb.GatewayUp, error) {
	up := &ttnpb.GatewayUp{}
	up.UplinkMessages = make([]*ttnpb.UplinkMessage, 0)
	for rxIndex, rx := range data.RxPacket {
		if rx == nil {
			continue
		}
		convertedRx, err := convertUplink(data, rxIndex, md)
		if err != nil {
			return nil, err
		}
		up.UplinkMessages = append(up.UplinkMessages, &convertedRx)
	}
	if data.Stat != nil {
		up.GatewayStatus = convertStatus(*data.Stat, md)
	}
	if data.TxPacketAck != nil {
		result, ok := ttnAckError[data.TxPacketAck.Error]
		if !ok {
			result = ttnpb.TxAcknowledgment_UNKNOWN_ERROR
		}
		up.TxAcknowledgment = &ttnpb.TxAcknowledgment{
			Result: result,
		}
	}
	return up, nil
}

var (
	ttnAckError = map[TxError]ttnpb.TxAcknowledgment_Result{
		TxErrNone:            ttnpb.TxAcknowledgment_SUCCESS,
		TxErrTooLate:         ttnpb.TxAcknowledgment_TOO_LATE,
		TxErrTooEarly:        ttnpb.TxAcknowledgment_TOO_EARLY,
		TxErrCollisionBeacon: ttnpb.TxAcknowledgment_COLLISION_BEACON,
		TxErrCollisionPacket: ttnpb.TxAcknowledgment_COLLISION_PACKET,
		TxErrTxFreq:          ttnpb.TxAcknowledgment_TX_FREQ,
		TxErrTxPower:         ttnpb.TxAcknowledgment_TX_POWER,
		TxErrGPSUnlocked:     ttnpb.TxAcknowledgment_GPS_UNLOCKED,
	}
	semtechAckError = map[ttnpb.TxAcknowledgment_Result]TxError{
		ttnpb.TxAcknowledgment_SUCCESS:          TxErrNone,
		ttnpb.TxAcknowledgment_TOO_LATE:         TxErrTooLate,
		ttnpb.TxAcknowledgment_TOO_EARLY:        TxErrTooEarly,
		ttnpb.TxAcknowledgment_COLLISION_BEACON: TxErrCollisionBeacon,
		ttnpb.TxAcknowledgment_COLLISION_PACKET: TxErrCollisionPacket,
		ttnpb.TxAcknowledgment_TX_FREQ:          TxErrTxFreq,
		ttnpb.TxAcknowledgment_TX_POWER:         TxErrTxPower,
		ttnpb.TxAcknowledgment_GPS_UNLOCKED:     TxErrGPSUnlocked,
	}
)

func metadata(rx RxPacket, gatewayID ttnpb.GatewayIdentifiers) []*ttnpb.RxMetadata {
	return []*ttnpb.RxMetadata{
		{
			GatewayIdentifiers: gatewayID,
			AntennaIndex:       0,
			Timestamp:          uint64(rx.Tmst) * 1000,
			RSSI:               float32(rx.RSSI),
			SNR:                float32(rx.LSNR),
		},
	}
}

func fineTimestampMetadata(rx RxPacket, gatewayID ttnpb.GatewayIdentifiers) []*ttnpb.RxMetadata {
	md := make([]*ttnpb.RxMetadata, 0)
	for _, signal := range rx.RSig {
		signalMetadata := &ttnpb.RxMetadata{
			GatewayIdentifiers:    gatewayID,
			AntennaIndex:          uint32(signal.Ant),
			Timestamp:             uint64(rx.Tmst) * 1000,
			RSSI:                  float32(signal.RSSIS),
			ChannelRSSI:           float32(signal.RSSIC),
			RSSIStandardDeviation: float32(signal.RSSISD),
			SNR:                   float32(signal.LSNR),
			FrequencyOffset:       int64(signal.FOff),
		}
		if signal.ETime != "" {
			if etime, err := base64.RawStdEncoding.DecodeString(strings.TrimRight(signal.ETime, "=")); err == nil {
				signalMetadata.EncryptedFineTimestampKeyID = strconv.Itoa(int(rx.Aesk))
				signalMetadata.EncryptedFineTimestamp = etime
			}
		}
		md = append(md, signalMetadata)
	}
	return md
}

func convertUplink(data Data, rxIndex int, md UpstreamMetadata) (ttnpb.UplinkMessage, error) {
	up := ttnpb.UplinkMessage{}
	rx := *data.RxPacket[rxIndex]
	up.Settings = ttnpb.TxSettings{
		CodingRate:   rx.CodR,
		Frequency:    uint64(rx.Freq * 1000000),
		ChannelIndex: uint32(rx.Chan),
	}

	rawPayload, err := base64.RawStdEncoding.DecodeString(strings.TrimRight(rx.Data, "="))
	if err != nil {
		return up, errPayload.WithCause(err)
	}
	up.RawPayload = rawPayload

	if rx.RSig != nil && len(rx.RSig) > 0 {
		up.RxMetadata = fineTimestampMetadata(rx, md.ID)
	} else {
		up.RxMetadata = metadata(rx, md.ID)
	}

	if rx.Time != nil {
		goTime := time.Time(*rx.Time)
		for mdIndex := range up.RxMetadata {
			up.RxMetadata[mdIndex].Time = &goTime
		}
		up.RxMetadata[0].Time = &goTime
	}

	switch rx.Modu {
	case lora:
		up.Settings.Modulation = ttnpb.Modulation_LORA
		sf, err := rx.DatR.SpreadingFactor()
		if err != nil {
			return up, errSpreadingFactor.WithCause(err)
		}
		up.Settings.SpreadingFactor = uint32(sf)
		if up.Settings.Bandwidth, err = rx.DatR.Bandwidth(); err != nil {
			return up, errBandwidth.WithCause(err)
		}
	case fsk:
		up.Settings.Modulation = ttnpb.Modulation_FSK
		up.Settings.BitRate = rx.DatR.FSK
	default:
		return up, errModulation.WithAttributes("modulation", rx.Modu)
	}

	return up, nil
}

func addVersions(status *ttnpb.GatewayStatus, stat Stat) {
	if stat.FPGA != nil {
		status.Versions["fpga"] = strconv.Itoa(int(*stat.FPGA))
	}
	if stat.DSP != nil {
		status.Versions["dsp"] = strconv.Itoa(int(*stat.DSP))
	}
	if stat.HAL != nil {
		status.Versions["hal"] = *stat.HAL
	}
}

func addMetrics(status *ttnpb.GatewayStatus, stat Stat) {
	status.Metrics["rxnb"] = float32(stat.RxNb)
	status.Metrics["rxok"] = float32(stat.RxOk)
	status.Metrics["rxfw"] = float32(stat.RxFW)
	status.Metrics["ackr"] = float32(stat.ACKR)
	status.Metrics["dwnb"] = float32(stat.DWNb)
	status.Metrics["txnb"] = float32(stat.TxNb)
	if stat.Temp != nil {
		status.Metrics["temp"] = float32(*stat.Temp)
	}
	if stat.LPPS != nil {
		status.Metrics["lpps"] = float32(*stat.LPPS)
	}
	if stat.LMNW != nil {
		status.Metrics["lmnw"] = float32(*stat.LMNW)
	}
	if stat.LMST != nil {
		status.Metrics["lmst"] = float32(*stat.LMST)
	}
	if stat.LMOK != nil {
		status.Metrics["lmok"] = float32(*stat.LMOK)
	}
}

func convertStatus(stat Stat, md UpstreamMetadata) *ttnpb.GatewayStatus {
	status := &ttnpb.GatewayStatus{
		Metrics:  map[string]float32{},
		Versions: map[string]string{},
		IP:       []string{md.IP},
	}

	if stat.Lati != nil && stat.Long != nil {
		loc := &ttnpb.Location{Latitude: float32(*stat.Lati), Longitude: float32(*stat.Long)}
		if stat.Alti != nil {
			loc.Altitude = *stat.Alti
		}
		if validLocation(*loc) {
			status.AntennaLocations = []*ttnpb.Location{loc}
		}
	}

	currentTime := time.Time(stat.Time)
	status.Time = currentTime
	if stat.Boot != nil {
		bootTime := time.Time(*stat.Boot)
		status.BootTime = bootTime
	}

	addVersions(status, stat)
	for versionName, version := range ttnVersions {
		status.Versions[versionName] = version
	}
	addMetrics(status, stat)
	return status
}

// FromGatewayUp converts the upstream message to the UDP format.
func FromGatewayUp(up *ttnpb.GatewayUp) (rxs []*RxPacket, stat *Stat, ack *TxPacketAck) {
	rxs = make([]*RxPacket, 0, len(up.UplinkMessages))

	for _, msg := range up.UplinkMessages {
		var modulation string
		var dataRate types.DataRate
		switch msg.Settings.Modulation {
		case ttnpb.Modulation_LORA:
			modulation = lora
			dataRate.LoRa = fmt.Sprintf("SF%dBW%d", msg.Settings.SpreadingFactor, msg.Settings.Bandwidth/1000)
		case ttnpb.Modulation_FSK:
			modulation = fsk
			dataRate.FSK = msg.Settings.BitRate
		}
		rxs = append(rxs, &RxPacket{
			Freq: float64(msg.Settings.Frequency) / 1000000,
			Chan: uint8(msg.Settings.ChannelIndex),
			Modu: modulation,
			DatR: DataRate{DataRate: dataRate},
			CodR: msg.Settings.CodingRate,
			Size: uint16(len(msg.RawPayload)),
			Data: base64.StdEncoding.EncodeToString(msg.RawPayload),
			Tmst: uint32(msg.RxMetadata[0].Timestamp / 1000),
			RSSI: int16(msg.RxMetadata[0].RSSI),
			LSNR: float64(msg.RxMetadata[0].SNR),
		})
	}
	if up.GatewayStatus != nil {
		stat = &Stat{
			Time: ExpandedTime(up.GatewayStatus.Time),
		}
	}
	if up.TxAcknowledgment != nil {
		ack = &TxPacketAck{
			Error: semtechAckError[up.TxAcknowledgment.Result],
		}
	}
	return
}

// ToDownlinkMessage converts the UDP format to a downlink message.
func ToDownlinkMessage(tx *TxPacket) (*ttnpb.DownlinkMessage, error) {
	msg := &ttnpb.DownlinkMessage{
		Settings: ttnpb.TxSettings{
			CodingRate:         tx.CodR,
			Frequency:          uint64(tx.Freq * 1000000),
			InvertPolarization: tx.IPol,
			TxPower:            int32(tx.Powe) + eirpDelta,
		},
		TxMetadata: ttnpb.TxMetadata{
			Timestamp: uint64(tx.Tmst * 1000),
		},
	}
	var err error
	msg.RawPayload, err = base64.RawStdEncoding.DecodeString(strings.TrimRight(tx.Data, "="))
	if err != nil {
		return nil, err
	}
	switch tx.Modu {
	case lora:
		msg.Settings.Modulation = ttnpb.Modulation_LORA
		msg.Settings.EnableCRC = !tx.NCRC
		sf, err := tx.DatR.SpreadingFactor()
		if err != nil {
			return nil, err
		}
		msg.Settings.SpreadingFactor = uint32(sf)
		bw, err := tx.DatR.Bandwidth()
		if err != nil {
			return nil, err
		}
		msg.Settings.Bandwidth = bw

	case fsk:
		msg.Settings.Modulation = ttnpb.Modulation_FSK
		msg.Settings.BitRate = tx.DatR.FSK
	}
	return msg, nil
}

// FromDownlinkMessage converts to the downlink message to the UDP format.
func FromDownlinkMessage(msg *ttnpb.DownlinkMessage) (*TxPacket, error) {
	payload := msg.GetRawPayload()
	tmst := msg.TxMetadata.Timestamp / 1000
	tx := &TxPacket{
		CodR: msg.Settings.CodingRate,
		Freq: float64(msg.Settings.Frequency) / 1000000,
		Imme: msg.TxMetadata.Timestamp == 0,
		IPol: msg.Settings.InvertPolarization,
		Powe: uint8(msg.Settings.TxPower) - uint8(eirpDelta),
		Size: uint16(len(payload)),
		Tmst: uint32(tmst % math.MaxUint32),
		Data: base64.StdEncoding.EncodeToString(payload),
	}
	gpsTime := CompactTime(msg.TxMetadata.Time)
	tx.Time = &gpsTime

	switch msg.Settings.Modulation {
	case ttnpb.Modulation_LORA:
		tx.Modu = "LORA"
		tx.NCRC = !msg.Settings.EnableCRC
		tx.DatR.LoRa = fmt.Sprintf("SF%dBW%d", msg.Settings.SpreadingFactor, msg.Settings.Bandwidth/1000)
	case ttnpb.Modulation_FSK:
		tx.Modu = "FSK"
		tx.DatR.FSK = msg.Settings.BitRate
	default:
		return tx, errModulation.WithAttributes("modulation", msg.Settings.Modulation)
	}

	return tx, nil
}
