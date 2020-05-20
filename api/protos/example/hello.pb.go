// Code generated by protoc-gen-go. DO NOT EDIT.
// source: hello.proto

package com_salty_protos

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	any "github.com/golang/protobuf/ptypes/any"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// defined Req message struct
type HelloReq struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HelloReq) Reset()         { *m = HelloReq{} }
func (m *HelloReq) String() string { return proto.CompactTextString(m) }
func (*HelloReq) ProtoMessage()    {}
func (*HelloReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_61ef911816e0a8ce, []int{0}
}

func (m *HelloReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HelloReq.Unmarshal(m, b)
}
func (m *HelloReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HelloReq.Marshal(b, m, deterministic)
}
func (m *HelloReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HelloReq.Merge(m, src)
}
func (m *HelloReq) XXX_Size() int {
	return xxx_messageInfo_HelloReq.Size(m)
}
func (m *HelloReq) XXX_DiscardUnknown() {
	xxx_messageInfo_HelloReq.DiscardUnknown(m)
}

var xxx_messageInfo_HelloReq proto.InternalMessageInfo

func (m *HelloReq) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

//defined Res message struct
type HelloResp struct {
	Message              string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HelloResp) Reset()         { *m = HelloResp{} }
func (m *HelloResp) String() string { return proto.CompactTextString(m) }
func (*HelloResp) ProtoMessage()    {}
func (*HelloResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_61ef911816e0a8ce, []int{1}
}

func (m *HelloResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HelloResp.Unmarshal(m, b)
}
func (m *HelloResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HelloResp.Marshal(b, m, deterministic)
}
func (m *HelloResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HelloResp.Merge(m, src)
}
func (m *HelloResp) XXX_Size() int {
	return xxx_messageInfo_HelloResp.Size(m)
}
func (m *HelloResp) XXX_DiscardUnknown() {
	xxx_messageInfo_HelloResp.DiscardUnknown(m)
}

var xxx_messageInfo_HelloResp proto.InternalMessageInfo

func (m *HelloResp) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

// Basic Request
type GrpcReq struct {
	Data                 *any.Any `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GrpcReq) Reset()         { *m = GrpcReq{} }
func (m *GrpcReq) String() string { return proto.CompactTextString(m) }
func (*GrpcReq) ProtoMessage()    {}
func (*GrpcReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_61ef911816e0a8ce, []int{2}
}

func (m *GrpcReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GrpcReq.Unmarshal(m, b)
}
func (m *GrpcReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GrpcReq.Marshal(b, m, deterministic)
}
func (m *GrpcReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GrpcReq.Merge(m, src)
}
func (m *GrpcReq) XXX_Size() int {
	return xxx_messageInfo_GrpcReq.Size(m)
}
func (m *GrpcReq) XXX_DiscardUnknown() {
	xxx_messageInfo_GrpcReq.DiscardUnknown(m)
}

var xxx_messageInfo_GrpcReq proto.InternalMessageInfo

func (m *GrpcReq) GetData() *any.Any {
	if m != nil {
		return m.Data
	}
	return nil
}

// Basic Response
type GrpcResp struct {
	Code                 int32    `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Message              string   `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Data                 *any.Any `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GrpcResp) Reset()         { *m = GrpcResp{} }
func (m *GrpcResp) String() string { return proto.CompactTextString(m) }
func (*GrpcResp) ProtoMessage()    {}
func (*GrpcResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_61ef911816e0a8ce, []int{3}
}

func (m *GrpcResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GrpcResp.Unmarshal(m, b)
}
func (m *GrpcResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GrpcResp.Marshal(b, m, deterministic)
}
func (m *GrpcResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GrpcResp.Merge(m, src)
}
func (m *GrpcResp) XXX_Size() int {
	return xxx_messageInfo_GrpcResp.Size(m)
}
func (m *GrpcResp) XXX_DiscardUnknown() {
	xxx_messageInfo_GrpcResp.DiscardUnknown(m)
}

var xxx_messageInfo_GrpcResp proto.InternalMessageInfo

func (m *GrpcResp) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *GrpcResp) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *GrpcResp) GetData() *any.Any {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*HelloReq)(nil), "com.salty.protos.HelloReq")
	proto.RegisterType((*HelloResp)(nil), "com.salty.protos.HelloResp")
	proto.RegisterType((*GrpcReq)(nil), "com.salty.protos.GrpcReq")
	proto.RegisterType((*GrpcResp)(nil), "com.salty.protos.GrpcResp")
}

func init() {
	proto.RegisterFile("hello.proto", fileDescriptor_61ef911816e0a8ce)
}

var fileDescriptor_61ef911816e0a8ce = []byte{
	// 224 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x90, 0x3f, 0x4f, 0x80, 0x30,
	0x10, 0xc5, 0x45, 0x91, 0x3f, 0xc7, 0x62, 0x1a, 0x07, 0x60, 0x30, 0xa6, 0x89, 0x09, 0x53, 0x49,
	0xe0, 0x13, 0x18, 0x07, 0x5d, 0x5c, 0xea, 0xe0, 0x7c, 0x40, 0x45, 0x93, 0x42, 0x2b, 0xad, 0x03,
	0xdf, 0xde, 0xd0, 0x82, 0x51, 0x13, 0xe3, 0x76, 0x77, 0xef, 0xdd, 0xfb, 0xb5, 0x07, 0xd9, 0xab,
	0x90, 0x52, 0x31, 0xbd, 0x28, 0xab, 0xc8, 0x45, 0xaf, 0x26, 0x66, 0x50, 0xda, 0xd5, 0x0f, 0x4c,
	0x59, 0x8c, 0x4a, 0x8d, 0x52, 0xd4, 0xae, 0xed, 0x3e, 0x5e, 0x6a, 0x9c, 0x77, 0x8d, 0x5e, 0x41,
	0xf2, 0xb0, 0xed, 0x72, 0xf1, 0x4e, 0x08, 0x84, 0x33, 0x4e, 0x22, 0x0f, 0xae, 0x83, 0x2a, 0xe5,
	0xae, 0xa6, 0x37, 0x90, 0xee, 0xba, 0xd1, 0x24, 0x87, 0x78, 0x12, 0xc6, 0xe0, 0x78, 0x78, 0x8e,
	0x96, 0xb6, 0x10, 0xdf, 0x2f, 0xba, 0xdf, 0x52, 0x2a, 0x08, 0x07, 0xb4, 0xe8, 0x1c, 0x59, 0x73,
	0xc9, 0x3c, 0x9b, 0x1d, 0x6c, 0x76, 0x3b, 0xaf, 0xdc, 0x39, 0x68, 0x07, 0x89, 0x5f, 0x32, 0x7a,
	0x63, 0xf7, 0x6a, 0xf0, 0xb9, 0xe7, 0xdc, 0xd5, 0xdf, 0x71, 0xa7, 0x3f, 0x70, 0x5f, 0x8c, 0xb3,
	0xff, 0x18, 0xcd, 0x23, 0x44, 0xcf, 0xf8, 0x66, 0xc5, 0x42, 0xee, 0x20, 0x79, 0xc2, 0xd5, 0x7d,
	0x86, 0x14, 0xec, 0xf7, 0x8d, 0xd8, 0xfe, 0xfc, 0xb2, 0xfc, 0x4b, 0x32, 0x9a, 0x9e, 0x74, 0x91,
	0x1b, 0xb5, 0x9f, 0x01, 0x00, 0x00, 0xff, 0xff, 0xcc, 0x7f, 0xd0, 0xd7, 0x71, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// WaiterClient is the client API for Waiter service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type WaiterClient interface {
	SayHello(ctx context.Context, in *GrpcReq, opts ...grpc.CallOption) (*GrpcResp, error)
}

type waiterClient struct {
	cc grpc.ClientConnInterface
}

func NewWaiterClient(cc grpc.ClientConnInterface) WaiterClient {
	return &waiterClient{cc}
}

func (c *waiterClient) SayHello(ctx context.Context, in *GrpcReq, opts ...grpc.CallOption) (*GrpcResp, error) {
	out := new(GrpcResp)
	err := c.cc.Invoke(ctx, "/com.salty.protos.Waiter/SayHello", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// WaiterServer is the server API for Waiter service.
type WaiterServer interface {
	SayHello(context.Context, *GrpcReq) (*GrpcResp, error)
}

// UnimplementedWaiterServer can be embedded to have forward compatible implementations.
type UnimplementedWaiterServer struct {
}

func (*UnimplementedWaiterServer) SayHello(ctx context.Context, req *GrpcReq) (*GrpcResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SayHello not implemented")
}

func RegisterWaiterServer(s *grpc.Server, srv WaiterServer) {
	s.RegisterService(&_Waiter_serviceDesc, srv)
}

func _Waiter_SayHello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GrpcReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WaiterServer).SayHello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/com.salty.protos.Waiter/SayHello",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WaiterServer).SayHello(ctx, req.(*GrpcReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _Waiter_serviceDesc = grpc.ServiceDesc{
	ServiceName: "com.salty.protos.Waiter",
	HandlerType: (*WaiterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SayHello",
			Handler:    _Waiter_SayHello_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "hello.proto",
}
