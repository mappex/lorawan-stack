// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: lorawan-stack/api/packetbrokeragent.proto

package ttnpb

import (
	context "context"
	fmt "fmt"
	math "math"

	proto "github.com/gogo/protobuf/proto"
	types "github.com/gogo/protobuf/types"
	golang_proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = golang_proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

func init() {
	proto.RegisterFile("lorawan-stack/api/packetbrokeragent.proto", fileDescriptor_1a44242dc5cd678e)
}
func init() {
	golang_proto.RegisterFile("lorawan-stack/api/packetbrokeragent.proto", fileDescriptor_1a44242dc5cd678e)
}

var fileDescriptor_1a44242dc5cd678e = []byte{
	// 311 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x90, 0xa1, 0x6f, 0xc2, 0x40,
	0x14, 0x87, 0xdf, 0x13, 0x9b, 0x20, 0xd9, 0x04, 0x62, 0x82, 0x25, 0xbf, 0x4c, 0xcc, 0x4c, 0xec,
	0x2e, 0x19, 0xff, 0xc1, 0x92, 0x05, 0x45, 0x82, 0xd9, 0xc4, 0xdc, 0x95, 0xdc, 0x4a, 0xd3, 0xd2,
	0x6b, 0xda, 0x03, 0x82, 0x43, 0x22, 0x27, 0x27, 0x67, 0x96, 0x20, 0x91, 0x48, 0x24, 0x12, 0x89,
	0xa4, 0x77, 0x06, 0x89, 0x44, 0x2e, 0xa3, 0x9b, 0x20, 0xd3, 0xbf, 0x2f, 0xdf, 0xf7, 0xf2, 0x6a,
	0x77, 0x89, 0xc9, 0xd5, 0x48, 0xa5, 0xf7, 0x85, 0x55, 0xdd, 0x58, 0xaa, 0x2c, 0x92, 0x99, 0xea,
	0xc6, 0xda, 0x06, 0xb9, 0x89, 0x75, 0xae, 0x42, 0x9d, 0x5a, 0x91, 0xe5, 0xc6, 0x9a, 0xfa, 0xa5,
	0xb5, 0xa9, 0xf8, 0xc5, 0xc5, 0xb0, 0xd9, 0xb8, 0x0e, 0x8d, 0x09, 0x13, 0x2d, 0x8f, 0x6b, 0x30,
	0x78, 0x93, 0xba, 0x9f, 0xd9, 0x71, 0x05, 0x37, 0x6e, 0xfe, 0x7b, 0xfb, 0xba, 0x28, 0x54, 0xa8,
	0x8b, 0x8a, 0x78, 0x78, 0xa9, 0x9d, 0xb5, 0x8a, 0x4e, 0xa0, 0xea, 0xed, 0xda, 0x45, 0x67, 0x10,
	0x24, 0x51, 0xd1, 0x7b, 0xce, 0x92, 0x28, 0x8d, 0xeb, 0xb7, 0xe2, 0xb4, 0x24, 0x5a, 0xca, 0xea,
	0x91, 0x1a, 0x57, 0x73, 0xbb, 0xd2, 0x34, 0xae, 0x44, 0xd5, 0x17, 0x7f, 0x7d, 0xf1, 0xf4, 0xd3,
	0x7f, 0xfc, 0xe2, 0x55, 0x09, 0x5e, 0x97, 0xe0, 0x4d, 0x09, 0xda, 0x96, 0xa0, 0x5d, 0x09, 0xda,
	0x97, 0xa0, 0x43, 0x09, 0x9e, 0x38, 0xf0, 0xd4, 0x81, 0x66, 0x0e, 0x3c, 0x77, 0xa0, 0x85, 0x03,
	0x2d, 0x1d, 0x68, 0xe5, 0xc0, 0x6b, 0x07, 0xde, 0x38, 0xd0, 0xd6, 0x81, 0x77, 0x0e, 0xb4, 0x77,
	0xe0, 0x83, 0x03, 0x4d, 0x3c, 0x68, 0xea, 0xc1, 0xef, 0x1e, 0xf4, 0xe1, 0xc1, 0x9f, 0x1e, 0x34,
	0xf3, 0xa0, 0xb9, 0x07, 0x2f, 0x3c, 0x78, 0xe9, 0xc1, 0xaf, 0x32, 0x34, 0xc2, 0xf6, 0xb4, 0xed,
	0x45, 0x69, 0x58, 0x88, 0x54, 0xdb, 0x91, 0xc9, 0x63, 0x79, 0xfa, 0x85, 0x61, 0x53, 0x66, 0x71,
	0x28, 0xad, 0x4d, 0xb3, 0x20, 0x38, 0x3f, 0xde, 0xdd, 0xfc, 0x0e, 0x00, 0x00, 0xff, 0xff, 0x46,
	0xd8, 0x87, 0x11, 0x82, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// GsPbaClient is the client API for GsPba service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type GsPbaClient interface {
	PublishUplink(ctx context.Context, in *GatewayUplinkMessage, opts ...grpc.CallOption) (*types.Empty, error)
}

type gsPbaClient struct {
	cc *grpc.ClientConn
}

func NewGsPbaClient(cc *grpc.ClientConn) GsPbaClient {
	return &gsPbaClient{cc}
}

func (c *gsPbaClient) PublishUplink(ctx context.Context, in *GatewayUplinkMessage, opts ...grpc.CallOption) (*types.Empty, error) {
	out := new(types.Empty)
	err := c.cc.Invoke(ctx, "/ttn.lorawan.v3.GsPba/PublishUplink", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GsPbaServer is the server API for GsPba service.
type GsPbaServer interface {
	PublishUplink(context.Context, *GatewayUplinkMessage) (*types.Empty, error)
}

// UnimplementedGsPbaServer can be embedded to have forward compatible implementations.
type UnimplementedGsPbaServer struct {
}

func (*UnimplementedGsPbaServer) PublishUplink(ctx context.Context, req *GatewayUplinkMessage) (*types.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PublishUplink not implemented")
}

func RegisterGsPbaServer(s *grpc.Server, srv GsPbaServer) {
	s.RegisterService(&_GsPba_serviceDesc, srv)
}

func _GsPba_PublishUplink_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GatewayUplinkMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GsPbaServer).PublishUplink(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ttn.lorawan.v3.GsPba/PublishUplink",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GsPbaServer).PublishUplink(ctx, req.(*GatewayUplinkMessage))
	}
	return interceptor(ctx, in, info, handler)
}

var _GsPba_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ttn.lorawan.v3.GsPba",
	HandlerType: (*GsPbaServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PublishUplink",
			Handler:    _GsPba_PublishUplink_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "lorawan-stack/api/packetbrokeragent.proto",
}
