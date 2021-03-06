// Code generated by protoc-gen-go.
// source: randomnumber.proto
// DO NOT EDIT!

/*
Package randomnumber is a generated protocol buffer package.

It is generated from these files:
	randomnumber.proto

It has these top-level messages:
	GetIntRequest
	GetIntResponse
*/
package randomnumber

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type GetIntRequest struct {
	Min int32 `protobuf:"varint,1,opt,name=min" json:"min,omitempty"`
	Max int32 `protobuf:"varint,2,opt,name=max" json:"max,omitempty"`
}

func (m *GetIntRequest) Reset()                    { *m = GetIntRequest{} }
func (m *GetIntRequest) String() string            { return proto.CompactTextString(m) }
func (*GetIntRequest) ProtoMessage()               {}
func (*GetIntRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type GetIntResponse struct {
	RandomInt int32 `protobuf:"varint,1,opt,name=randomInt" json:"randomInt,omitempty"`
}

func (m *GetIntResponse) Reset()                    { *m = GetIntResponse{} }
func (m *GetIntResponse) String() string            { return proto.CompactTextString(m) }
func (*GetIntResponse) ProtoMessage()               {}
func (*GetIntResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func init() {
	proto.RegisterType((*GetIntRequest)(nil), "GetIntRequest")
	proto.RegisterType((*GetIntResponse)(nil), "GetIntResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion3

// Client API for RandomNumber service

type RandomNumberClient interface {
	GetInt(ctx context.Context, in *GetIntRequest, opts ...grpc.CallOption) (*GetIntResponse, error)
}

type randomNumberClient struct {
	cc *grpc.ClientConn
}

func NewRandomNumberClient(cc *grpc.ClientConn) RandomNumberClient {
	return &randomNumberClient{cc}
}

func (c *randomNumberClient) GetInt(ctx context.Context, in *GetIntRequest, opts ...grpc.CallOption) (*GetIntResponse, error) {
	out := new(GetIntResponse)
	err := grpc.Invoke(ctx, "/RandomNumber/GetInt", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for RandomNumber service

type RandomNumberServer interface {
	GetInt(context.Context, *GetIntRequest) (*GetIntResponse, error)
}

func RegisterRandomNumberServer(s *grpc.Server, srv RandomNumberServer) {
	s.RegisterService(&_RandomNumber_serviceDesc, srv)
}

func _RandomNumber_GetInt_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetIntRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RandomNumberServer).GetInt(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/RandomNumber/GetInt",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RandomNumberServer).GetInt(ctx, req.(*GetIntRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _RandomNumber_serviceDesc = grpc.ServiceDesc{
	ServiceName: "RandomNumber",
	HandlerType: (*RandomNumberServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetInt",
			Handler:    _RandomNumber_GetInt_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: fileDescriptor0,
}

func init() { proto.RegisterFile("randomnumber.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 145 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0x12, 0x2a, 0x4a, 0xcc, 0x4b,
	0xc9, 0xcf, 0xcd, 0x2b, 0xcd, 0x4d, 0x4a, 0x2d, 0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x57, 0x32,
	0xe6, 0xe2, 0x75, 0x4f, 0x2d, 0xf1, 0xcc, 0x2b, 0x09, 0x4a, 0x2d, 0x2c, 0x4d, 0x2d, 0x2e, 0x11,
	0x12, 0xe0, 0x62, 0xce, 0xcd, 0xcc, 0x93, 0x60, 0x54, 0x60, 0xd4, 0x60, 0x0d, 0x02, 0x31, 0xc1,
	0x22, 0x89, 0x15, 0x12, 0x4c, 0x50, 0x91, 0xc4, 0x0a, 0x25, 0x3d, 0x2e, 0x3e, 0x98, 0xa6, 0xe2,
	0x82, 0xfc, 0xbc, 0xe2, 0x54, 0x21, 0x19, 0x2e, 0x4e, 0x88, 0xe1, 0x9e, 0x79, 0x25, 0x50, 0xbd,
	0x08, 0x01, 0x23, 0x6b, 0x2e, 0x9e, 0x20, 0x30, 0xc7, 0x0f, 0x6c, 0xb5, 0x90, 0x36, 0x17, 0x1b,
	0x44, 0xbf, 0x10, 0x9f, 0x1e, 0x8a, 0xed, 0x52, 0xfc, 0x7a, 0xa8, 0x06, 0x2b, 0x31, 0x24, 0xb1,
	0x81, 0x1d, 0x6a, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0xd5, 0xe0, 0x26, 0xe5, 0xbe, 0x00, 0x00,
	0x00,
}
