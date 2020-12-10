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

package devicerepository_test

import (
	"context"
	"fmt"
	"net"
	"testing"
	"time"

	"github.com/smartystreets/assertions"
	"go.thethings.network/lorawan-stack/v3/pkg/cluster"
	"go.thethings.network/lorawan-stack/v3/pkg/component"
	componenttest "go.thethings.network/lorawan-stack/v3/pkg/component/test"
	"go.thethings.network/lorawan-stack/v3/pkg/config"
	"go.thethings.network/lorawan-stack/v3/pkg/errors"
	"go.thethings.network/lorawan-stack/v3/pkg/messageprocessors"
	dr_processor "go.thethings.network/lorawan-stack/v3/pkg/messageprocessors/devicerepository"
	"go.thethings.network/lorawan-stack/v3/pkg/rpcserver"
	"go.thethings.network/lorawan-stack/v3/pkg/ttnpb"
	"go.thethings.network/lorawan-stack/v3/pkg/util/test"
	"go.thethings.network/lorawan-stack/v3/pkg/util/test/assertions/should"
	"google.golang.org/grpc"
)

// mockProcessor is a mock messageprocessors.PayloadEncodeDecoder
type mockProcessor struct {
	ch  chan string
	err error
}

func (p *mockProcessor) EncodeDownlink(ctx context.Context, ids ttnpb.EndDeviceIdentifiers, version *ttnpb.EndDeviceVersionIdentifiers, message *ttnpb.ApplicationDownlink, parameter string) error {
	p.ch <- parameter
	return p.err
}

func (p *mockProcessor) DecodeUplink(ctx context.Context, ids ttnpb.EndDeviceIdentifiers, version *ttnpb.EndDeviceVersionIdentifiers, message *ttnpb.ApplicationUplink, parameter string) error {
	p.ch <- parameter
	return p.err
}

func (p *mockProcessor) DecodeDownlink(ctx context.Context, ids ttnpb.EndDeviceIdentifiers, version *ttnpb.EndDeviceVersionIdentifiers, message *ttnpb.ApplicationDownlink, parameter string) error {
	p.ch <- parameter
	return p.err
}

type mockDR struct {
	uplinkDecoders,
	downlinkDecoders,
	downlinkEncoders map[string]*ttnpb.MessagePayloadFormatter
}

func (dr *mockDR) ListBrands(_ context.Context, _ *ttnpb.ListEndDeviceBrandsRequest) (*ttnpb.ListEndDeviceBrandsResponse, error) {
	panic("not implemented")
}

func (dr *mockDR) GetBrand(_ context.Context, _ *ttnpb.GetEndDeviceBrandRequest) (*ttnpb.EndDeviceBrand, error) {
	panic("not implemented") // TODO: Implement
}

func (dr *mockDR) ListModels(_ context.Context, _ *ttnpb.ListEndDeviceModelsRequest) (*ttnpb.ListEndDeviceModelsResponse, error) {
	panic("not implemented") // TODO: Implement
}

func (dr *mockDR) GetModel(_ context.Context, _ *ttnpb.GetEndDeviceModelRequest) (*ttnpb.EndDeviceModel, error) {
	panic("not implemented") // TODO: Implement
}

func (dr *mockDR) GetTemplate(_ context.Context, _ *ttnpb.EndDeviceVersionIdentifiers) (*ttnpb.EndDeviceTemplate, error) {
	panic("not implemented") // TODO: Implement
}

func (dr *mockDR) key(ids *ttnpb.EndDeviceVersionIdentifiers) string {
	return fmt.Sprintf("%s:%s:%s:%s", ids.BrandID, ids.ModelID, ids.FirmwareVersion, ids.BandID)
}

var (
	mockError = fmt.Errorf("mock_error")
)

func (dr *mockDR) GetUplinkDecoder(_ context.Context, ids *ttnpb.EndDeviceVersionIdentifiers) (*ttnpb.MessagePayloadFormatter, error) {
	f, ok := dr.uplinkDecoders[dr.key(ids)]
	if !ok {
		return nil, mockError
	}
	return f, nil
}

func (dr *mockDR) GetDownlinkDecoder(_ context.Context, ids *ttnpb.EndDeviceVersionIdentifiers) (*ttnpb.MessagePayloadFormatter, error) {
	f, ok := dr.downlinkDecoders[dr.key(ids)]
	if !ok {
		return nil, mockError
	}
	return f, nil
}

func (dr *mockDR) GetDownlinkEncoder(_ context.Context, ids *ttnpb.EndDeviceVersionIdentifiers) (*ttnpb.MessagePayloadFormatter, error) {
	f, ok := dr.downlinkEncoders[dr.key(ids)]
	if !ok {
		return nil, mockError
	}
	return f, nil
}

// start mock device repository and return listen address.
func (dr *mockDR) start(ctx context.Context) string {
	srv := rpcserver.New(ctx)
	ttnpb.RegisterDeviceRepositoryServer(srv.Server, dr)
	lis, err := net.Listen("tcp", ":0")
	if err != nil {
		panic(err)
	}
	go srv.Serve(lis)
	return lis.Addr().String()
}

func mustHavePeer(ctx context.Context, c *component.Component, role ttnpb.ClusterRole) {
	for i := 0; i < 20; i++ {
		time.Sleep(20 * time.Millisecond)
		if _, err := c.GetPeer(ctx, role, nil); err == nil {
			return
		}
	}
	panic("could not connect to peer")
}

func TestDeviceRepository(t *testing.T) {
	ids := &ttnpb.EndDeviceVersionIdentifiers{
		BrandID:         "brand",
		ModelID:         "model",
		FirmwareVersion: "1.0",
		HardwareVersion: "1.1",
		BandID:          "band",
	}
	idsNotFound := &ttnpb.EndDeviceVersionIdentifiers{
		BrandID:         "brand2",
		ModelID:         "model1",
		FirmwareVersion: "1.0",
		HardwareVersion: "1.1",
		BandID:          "band",
	}
	idsUnknownFormatter := &ttnpb.EndDeviceVersionIdentifiers{
		BrandID:         "brand3",
		ModelID:         "model1",
		FirmwareVersion: "1.0",
		HardwareVersion: "1.1",
		BandID:          "band",
	}

	dr := &mockDR{
		uplinkDecoders:   make(map[string]*ttnpb.MessagePayloadFormatter),
		downlinkDecoders: make(map[string]*ttnpb.MessagePayloadFormatter),
		downlinkEncoders: make(map[string]*ttnpb.MessagePayloadFormatter),
	}
	dr.uplinkDecoders[dr.key(ids)] = &ttnpb.MessagePayloadFormatter{
		Formatter:          ttnpb.PayloadFormatter_FORMATTER_JAVASCRIPT,
		FormatterParameter: "uplink decoder",
	}
	dr.downlinkDecoders[dr.key(ids)] = &ttnpb.MessagePayloadFormatter{
		Formatter:          ttnpb.PayloadFormatter_FORMATTER_JAVASCRIPT,
		FormatterParameter: "downlink decoder",
	}
	dr.downlinkEncoders[dr.key(ids)] = &ttnpb.MessagePayloadFormatter{
		Formatter:          ttnpb.PayloadFormatter_FORMATTER_JAVASCRIPT,
		FormatterParameter: "downlink encoder",
	}
	dr.uplinkDecoders[dr.key(idsUnknownFormatter)] = &ttnpb.MessagePayloadFormatter{
		Formatter:          ttnpb.PayloadFormatter_FORMATTER_GRPC_SERVICE,
		FormatterParameter: "parameter",
	}
	drAddr := dr.start(test.Context())

	ctx := test.Context()
	mockProcessorJS := &mockProcessor{
		ch:  make(chan string, 1),
		err: nil,
	}

	// start mock device repository
	c := componenttest.NewComponent(t, &component.Config{
		ServiceBase: config.ServiceBase{
			GRPC: config.GRPC{
				Listen:                      ":0",
				AllowInsecureForCredentials: true,
			},
			Cluster: cluster.Config{
				DeviceRepository: drAddr,
			},
		},
	})
	componenttest.StartComponent(t, c)
	defer c.Close()

	mustHavePeer(ctx, c, ttnpb.ClusterRole_DEVICE_REPOSITORY)

	p := dr_processor.New(
		map[ttnpb.PayloadFormatter]messageprocessors.PayloadEncodeDecoder{
			ttnpb.PayloadFormatter_FORMATTER_JAVASCRIPT: mockProcessorJS,
		},
		func() (*grpc.ClientConn, error) {
			return c.GetPeerConn(ctx, ttnpb.ClusterRole_DEVICE_REPOSITORY, nil)
		},
		func() []grpc.CallOption {
			return []grpc.CallOption{}
		},
	)

	t.Run("DeviceNotFound", func(t *testing.T) {
		err := p.DecodeDownlink(test.Context(), ttnpb.EndDeviceIdentifiers{}, idsNotFound, nil, "")
		a := assertions.New(t)
		a.So(err.Error(), should.ContainSubstring, mockError.Error())

		select {
		case <-mockProcessorJS.ch:
			t.Error("Expected timeout but processor was called instead")
			t.FailNow()
		case <-time.After(30 * time.Millisecond):
		}
	})

	t.Run("NoProcessor", func(t *testing.T) {
		err := p.DecodeUplink(test.Context(), ttnpb.EndDeviceIdentifiers{}, idsUnknownFormatter, nil, "")
		a := assertions.New(t)
		a.So(errors.IsFailedPrecondition(err), should.BeTrue)

		select {
		case <-mockProcessorJS.ch:
			t.Error("Expected timeout but processor was called instead")
			t.FailNow()
		case <-time.After(30 * time.Millisecond):
		}
	})

	t.Run("UplinkDecoder", func(t *testing.T) {
		err := p.DecodeUplink(test.Context(), ttnpb.EndDeviceIdentifiers{}, ids, nil, "")
		a := assertions.New(t)
		a.So(err, should.BeNil)

		select {
		case s := <-mockProcessorJS.ch:
			a.So(s, should.Equal, "uplink decoder")
		case <-time.After(time.Second):
			t.Error("Timed out waiting for message processor")
			t.FailNow()
		}
	})
	t.Run("DownlinkDecoder", func(t *testing.T) {
		err := p.DecodeDownlink(test.Context(), ttnpb.EndDeviceIdentifiers{}, ids, nil, "")
		a := assertions.New(t)
		a.So(err, should.BeNil)

		select {
		case s := <-mockProcessorJS.ch:
			a.So(s, should.Equal, "downlink decoder")
		case <-time.After(time.Second):
			t.Error("Timed out waiting for message processor")
			t.FailNow()
		}
	})
	t.Run("DownlinkEncoder", func(t *testing.T) {
		err := p.EncodeDownlink(test.Context(), ttnpb.EndDeviceIdentifiers{}, ids, nil, "")
		a := assertions.New(t)
		a.So(err, should.BeNil)

		select {
		case s := <-mockProcessorJS.ch:
			a.So(s, should.Equal, "downlink encoder")
		case <-time.After(time.Second):
			t.Error("Timed out waiting for message processor")
			t.FailNow()
		}
	})
}
