// Copyright © 2020 The Things Network Foundation, The Things Industries B.V.
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

package devicerepository

import (
	"context"
	"fmt"

	"go.thethings.network/lorawan-stack/v3/pkg/errors"
	"go.thethings.network/lorawan-stack/v3/pkg/messageprocessors"
	"go.thethings.network/lorawan-stack/v3/pkg/ttnpb"
	"golang.org/x/sync/singleflight"
	"google.golang.org/grpc"
)

const (
	downlinkEncoder = 1
	downlinkDecoder = 2
	uplinkDecoder   = 3
)

// GetClientConnFunc is a function for retrieving a *grpc.ClientConn to the device repository peer.
type GetClientConnFunc func() (*grpc.ClientConn, error)

// GetCallOptionsFunc is a function for retrieving grpc.CallOptions for the device repository grpc client.
type GetCallOptionsFunc func() []grpc.CallOption

type host struct {
	ctx context.Context

	getClientConn  GetClientConnFunc
	getCallOptions GetCallOptionsFunc

	singleflight singleflight.Group

	processors map[ttnpb.PayloadFormatter]messageprocessors.PayloadEncodeDecoder
}

var errFormatterNotSupported = errors.DefineFailedPrecondition("formatter_not_supported", "formatter `{formatter}` not supported")

// New creates a new PayloadEncodeDecoder that retrieves codecs from the Device Repository
// and uses an underlying PayloadEncodeDecoder to execute them.
func New(processors map[ttnpb.PayloadFormatter]messageprocessors.PayloadEncodeDecoder, getClientConn GetClientConnFunc, getCallOptions GetCallOptionsFunc) messageprocessors.PayloadEncodeDecoder {
	return &host{
		getClientConn:  getClientConn,
		getCallOptions: getCallOptions,

		processors: processors,
	}
}

func key(version *ttnpb.EndDeviceVersionIdentifiers, which int) string {
	return fmt.Sprintf("%s:%s:%s:%s:%v", version.BrandID, version.ModelID, version.FirmwareVersion, version.BandID, which)
}

// EncodeDownlink encodes a downlink message.
func (h *host) EncodeDownlink(ctx context.Context, ids ttnpb.EndDeviceIdentifiers, version *ttnpb.EndDeviceVersionIdentifiers, message *ttnpb.ApplicationDownlink, parameter string) error {
	cc, err := h.getClientConn()
	if err != nil {
		return err
	}
	resInterface, err, _ := h.singleflight.Do(key(version, downlinkEncoder), func() (interface{}, error) {
		return ttnpb.NewDeviceRepositoryClient(cc).GetDownlinkEncoder(ctx, version, h.getCallOptions()...)
	})
	if err != nil {
		return err
	}
	res := resInterface.(*ttnpb.MessagePayloadFormatter)
	p, ok := h.processors[res.Formatter]
	if !ok {
		return errFormatterNotSupported.WithAttributes("formatter", res.Formatter.String())
	}
	return p.EncodeDownlink(ctx, ids, version, message, res.FormatterParameter)
}

// DecodeUplink decodes an uplink message.
func (h *host) DecodeUplink(ctx context.Context, ids ttnpb.EndDeviceIdentifiers, version *ttnpb.EndDeviceVersionIdentifiers, message *ttnpb.ApplicationUplink, parameter string) error {
	cc, err := h.getClientConn()
	if err != nil {
		return err
	}
	resInterface, err, _ := h.singleflight.Do(key(version, uplinkDecoder), func() (interface{}, error) {
		return ttnpb.NewDeviceRepositoryClient(cc).GetUplinkDecoder(ctx, version, h.getCallOptions()...)
	})
	if err != nil {
		return err
	}
	res := resInterface.(*ttnpb.MessagePayloadFormatter)
	p, ok := h.processors[res.Formatter]
	if !ok {
		return errFormatterNotSupported.WithAttributes("formatter", res.Formatter.String())
	}
	return p.DecodeUplink(ctx, ids, version, message, res.FormatterParameter)
}

// DecodeDownlink decodes a downlink message.
func (h *host) DecodeDownlink(ctx context.Context, ids ttnpb.EndDeviceIdentifiers, version *ttnpb.EndDeviceVersionIdentifiers, message *ttnpb.ApplicationDownlink, parameter string) error {
	cc, err := h.getClientConn()
	if err != nil {
		return err
	}
	resInterface, err, _ := h.singleflight.Do(key(version, downlinkDecoder), func() (interface{}, error) {
		return ttnpb.NewDeviceRepositoryClient(cc).GetDownlinkDecoder(ctx, version, h.getCallOptions()...)
	})
	if err != nil {
		return err
	}
	res := resInterface.(*ttnpb.MessagePayloadFormatter)

	p, ok := h.processors[res.Formatter]
	if !ok {
		return errFormatterNotSupported.WithAttributes("formatter", res.Formatter.String())
	}
	return p.DecodeDownlink(ctx, ids, version, message, res.FormatterParameter)
}
