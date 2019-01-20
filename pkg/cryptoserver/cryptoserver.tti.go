// Copyright © 2019 The Things Industries B.V.

package cryptoserver

import (
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"go.thethings.network/lorawan-stack/pkg/component"
	"go.thethings.network/lorawan-stack/pkg/ttnpb"
	"google.golang.org/grpc"
)

// CryptoServer implements the Crypto Server component.
//
// The Crypto Server exposes the NetworkCryptoService and ApplicationCryptoService services.
type CryptoServer struct {
	*component.Component

	grpc struct {
		networkCryptoService     *networkCryptoServiceServer
		applicationCryptoService *applicationCryptoServiceServer
	}
}

// New returns new *CryptoServer.
func New(c *component.Component, conf *Config) (*CryptoServer, error) {
	cs := &CryptoServer{
		Component: c,
	}

	network, err := conf.NewNetwork(cs.Context())
	if err != nil {
		return nil, err
	}
	if network != nil {
		cs.grpc.networkCryptoService = &networkCryptoServiceServer{
			Network: network,
		}
	}

	application, err := conf.NewApplication(cs.Context())
	if err != nil {
		return nil, err
	}
	if application != nil {
		cs.grpc.applicationCryptoService = &applicationCryptoServiceServer{
			Application: application,
		}
	}

	c.RegisterGRPC(cs)
	return cs, nil
}

// Roles of the gRPC service.
func (cs *CryptoServer) Roles() []ttnpb.PeerInfo_Role {
	return []ttnpb.PeerInfo_Role{ttnpb.PeerInfo_CRYPTO_SERVER}
}

// RegisterServices registers services provided by cs at s.
func (cs *CryptoServer) RegisterServices(s *grpc.Server) {
	if cs.grpc.networkCryptoService != nil {
		ttnpb.RegisterNetworkCryptoServiceServer(s, cs.grpc.networkCryptoService)
	}
	if cs.grpc.applicationCryptoService != nil {
		ttnpb.RegisterApplicationCryptoServiceServer(s, cs.grpc.applicationCryptoService)
	}
}

// RegisterHandlers registers gRPC handlers.
func (cs *CryptoServer) RegisterHandlers(s *runtime.ServeMux, conn *grpc.ClientConn) {}
