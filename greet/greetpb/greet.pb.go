// Code generated by protoc-gen-go. DO NOT EDIT.
// source: greet/greetpb/greet.proto

/*
Package greetpb is a generated protocol buffer package.

It is generated from these files:
	greet/greetpb/greet.proto

It has these top-level messages:
	Greeting
	GreetWithDeadlineRequest
	GreetWithDeadlineResponse
*/
package greetpb

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

type Greeting struct {
	FirstName string `protobuf:"bytes,1,opt,name=first_name,json=firstName" json:"first_name,omitempty"`
	LastName  string `protobuf:"bytes,2,opt,name=last_name,json=lastName" json:"last_name,omitempty"`
}

func (m *Greeting) Reset()                    { *m = Greeting{} }
func (m *Greeting) String() string            { return proto.CompactTextString(m) }
func (*Greeting) ProtoMessage()               {}
func (*Greeting) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Greeting) GetFirstName() string {
	if m != nil {
		return m.FirstName
	}
	return ""
}

func (m *Greeting) GetLastName() string {
	if m != nil {
		return m.LastName
	}
	return ""
}

type GreetWithDeadlineRequest struct {
	Greeting *Greeting `protobuf:"bytes,1,opt,name=greeting" json:"greeting,omitempty"`
}

func (m *GreetWithDeadlineRequest) Reset()                    { *m = GreetWithDeadlineRequest{} }
func (m *GreetWithDeadlineRequest) String() string            { return proto.CompactTextString(m) }
func (*GreetWithDeadlineRequest) ProtoMessage()               {}
func (*GreetWithDeadlineRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *GreetWithDeadlineRequest) GetGreeting() *Greeting {
	if m != nil {
		return m.Greeting
	}
	return nil
}

type GreetWithDeadlineResponse struct {
	Response string `protobuf:"bytes,1,opt,name=response" json:"response,omitempty"`
}

func (m *GreetWithDeadlineResponse) Reset()                    { *m = GreetWithDeadlineResponse{} }
func (m *GreetWithDeadlineResponse) String() string            { return proto.CompactTextString(m) }
func (*GreetWithDeadlineResponse) ProtoMessage()               {}
func (*GreetWithDeadlineResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *GreetWithDeadlineResponse) GetResponse() string {
	if m != nil {
		return m.Response
	}
	return ""
}

func init() {
	proto.RegisterType((*Greeting)(nil), "greet.Greeting")
	proto.RegisterType((*GreetWithDeadlineRequest)(nil), "greet.GreetWithDeadlineRequest")
	proto.RegisterType((*GreetWithDeadlineResponse)(nil), "greet.GreetWithDeadlineResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for GreetService service

type GreetServiceClient interface {
	GreetWithDeadline(ctx context.Context, in *GreetWithDeadlineRequest, opts ...grpc.CallOption) (*GreetWithDeadlineResponse, error)
}

type greetServiceClient struct {
	cc *grpc.ClientConn
}

func NewGreetServiceClient(cc *grpc.ClientConn) GreetServiceClient {
	return &greetServiceClient{cc}
}

func (c *greetServiceClient) GreetWithDeadline(ctx context.Context, in *GreetWithDeadlineRequest, opts ...grpc.CallOption) (*GreetWithDeadlineResponse, error) {
	out := new(GreetWithDeadlineResponse)
	err := grpc.Invoke(ctx, "/greet.GreetService/GreetWithDeadline", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for GreetService service

type GreetServiceServer interface {
	GreetWithDeadline(context.Context, *GreetWithDeadlineRequest) (*GreetWithDeadlineResponse, error)
}

func RegisterGreetServiceServer(s *grpc.Server, srv GreetServiceServer) {
	s.RegisterService(&_GreetService_serviceDesc, srv)
}

func _GreetService_GreetWithDeadline_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GreetWithDeadlineRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GreetServiceServer).GreetWithDeadline(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/greet.GreetService/GreetWithDeadline",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GreetServiceServer).GreetWithDeadline(ctx, req.(*GreetWithDeadlineRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _GreetService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "greet.GreetService",
	HandlerType: (*GreetServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GreetWithDeadline",
			Handler:    _GreetService_GreetWithDeadline_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "greet/greetpb/greet.proto",
}

func init() { proto.RegisterFile("greet/greetpb/greet.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 215 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x4c, 0x2f, 0x4a, 0x4d,
	0x2d, 0xd1, 0x07, 0x93, 0x05, 0x49, 0x10, 0x5a, 0xaf, 0xa0, 0x28, 0xbf, 0x24, 0x5f, 0x88, 0x15,
	0xcc, 0x51, 0x72, 0xe3, 0xe2, 0x70, 0x07, 0x31, 0x32, 0xf3, 0xd2, 0x85, 0x64, 0xb9, 0xb8, 0xd2,
	0x32, 0x8b, 0x8a, 0x4b, 0xe2, 0xf3, 0x12, 0x73, 0x53, 0x25, 0x18, 0x15, 0x18, 0x35, 0x38, 0x83,
	0x38, 0xc1, 0x22, 0x7e, 0x89, 0xb9, 0xa9, 0x42, 0xd2, 0x5c, 0x9c, 0x39, 0x89, 0x30, 0x59, 0x26,
	0xb0, 0x2c, 0x07, 0x48, 0x00, 0x24, 0xa9, 0xe4, 0xce, 0x25, 0x01, 0x36, 0x27, 0x3c, 0xb3, 0x24,
	0xc3, 0x25, 0x35, 0x31, 0x25, 0x27, 0x33, 0x2f, 0x35, 0x28, 0xb5, 0xb0, 0x34, 0xb5, 0xb8, 0x44,
	0x48, 0x9b, 0x8b, 0x23, 0x1d, 0x6a, 0x07, 0xd8, 0x54, 0x6e, 0x23, 0x7e, 0x3d, 0x88, 0x53, 0x60,
	0x56, 0x07, 0xc1, 0x15, 0x28, 0x99, 0x73, 0x49, 0x62, 0x31, 0xa8, 0xb8, 0x20, 0x3f, 0xaf, 0x38,
	0x55, 0x48, 0x8a, 0x8b, 0xa3, 0x08, 0xca, 0x86, 0xba, 0x0f, 0xce, 0x37, 0xca, 0xe0, 0xe2, 0x01,
	0x6b, 0x0c, 0x4e, 0x2d, 0x2a, 0xcb, 0x4c, 0x4e, 0x15, 0x8a, 0xe0, 0x12, 0xc4, 0x30, 0x48, 0x48,
	0x1e, 0xd9, 0x62, 0x2c, 0x6e, 0x95, 0x52, 0xc0, 0xad, 0x00, 0x62, 0x8f, 0x12, 0x83, 0x13, 0x67,
	0x14, 0x3b, 0x34, 0x44, 0x93, 0xd8, 0xc0, 0x81, 0x69, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0x97,
	0x67, 0xed, 0x5d, 0x69, 0x01, 0x00, 0x00,
}
