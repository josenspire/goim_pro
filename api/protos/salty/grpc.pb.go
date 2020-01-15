// Code generated by protoc-gen-go. DO NOT EDIT.
// source: grpc.proto

package com_salty_protos

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

func init() { proto.RegisterFile("grpc.proto", fileDescriptor_bedfbfc9b54e5600) }

var fileDescriptor_bedfbfc9b54e5600 = []byte{
	// 155 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4a, 0x2f, 0x2a, 0x48,
	0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0x48, 0xce, 0xcf, 0xd5, 0x2b, 0x4e, 0xcc, 0x29,
	0xa9, 0x84, 0x08, 0x14, 0x4b, 0xf1, 0xa4, 0xe6, 0x95, 0x64, 0xc2, 0xb8, 0x46, 0x53, 0x18, 0xb9,
	0xb8, 0x43, 0x8b, 0x53, 0x8b, 0x82, 0x53, 0x8b, 0xca, 0x32, 0x93, 0x53, 0x85, 0x9c, 0xb9, 0x38,
	0x82, 0x52, 0xd3, 0x33, 0x8b, 0x4b, 0x52, 0x8b, 0x84, 0x24, 0xf5, 0xd0, 0x35, 0xeb, 0xb9, 0x17,
	0x15, 0x24, 0x07, 0xa5, 0x16, 0x4a, 0x49, 0xe1, 0x92, 0x2a, 0x2e, 0x50, 0x62, 0x10, 0x72, 0xe0,
	0x62, 0xf5, 0xc9, 0x4f, 0xcf, 0xcc, 0x23, 0xdb, 0x04, 0xa3, 0x30, 0x2e, 0xae, 0x60, 0xdf, 0x60,
	0x98, 0xa3, 0x3c, 0xb8, 0x78, 0xfd, 0x93, 0x4a, 0x12, 0x33, 0xf3, 0x82, 0x7d, 0x83, 0x9d, 0xf3,
	0x53, 0x52, 0xc9, 0x36, 0xd7, 0x89, 0x29, 0x80, 0x31, 0x89, 0x0d, 0x2c, 0x6c, 0x0c, 0x08, 0x00,
	0x00, 0xff, 0xff, 0x28, 0x1e, 0xa4, 0x71, 0x27, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// UserServiceClient is the client API for UserService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type UserServiceClient interface {
	Register(ctx context.Context, in *GrpcReq, opts ...grpc.CallOption) (*GrpcResp, error)
	Login(ctx context.Context, in *GrpcReq, opts ...grpc.CallOption) (*GrpcResp, error)
}

type userServiceClient struct {
	cc *grpc.ClientConn
}

func NewUserServiceClient(cc *grpc.ClientConn) UserServiceClient {
	return &userServiceClient{cc}
}

func (c *userServiceClient) Register(ctx context.Context, in *GrpcReq, opts ...grpc.CallOption) (*GrpcResp, error) {
	out := new(GrpcResp)
	err := c.cc.Invoke(ctx, "/com.salty.protos.UserService/Register", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) Login(ctx context.Context, in *GrpcReq, opts ...grpc.CallOption) (*GrpcResp, error) {
	out := new(GrpcResp)
	err := c.cc.Invoke(ctx, "/com.salty.protos.UserService/Login", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServiceServer is the server API for UserService service.
type UserServiceServer interface {
	Register(context.Context, *GrpcReq) (*GrpcResp, error)
	Login(context.Context, *GrpcReq) (*GrpcResp, error)
}

// UnimplementedUserServiceServer can be embedded to have forward compatible implementations.
type UnimplementedUserServiceServer struct {
}

func (*UnimplementedUserServiceServer) Register(ctx context.Context, req *GrpcReq) (*GrpcResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Register not implemented")
}
func (*UnimplementedUserServiceServer) Login(ctx context.Context, req *GrpcReq) (*GrpcResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}

func RegisterUserServiceServer(s *grpc.Server, srv UserServiceServer) {
	s.RegisterService(&_UserService_serviceDesc, srv)
}

func _UserService_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GrpcReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/com.salty.protos.UserService/Register",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).Register(ctx, req.(*GrpcReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GrpcReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/com.salty.protos.UserService/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).Login(ctx, req.(*GrpcReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _UserService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "com.salty.protos.UserService",
	HandlerType: (*UserServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Register",
			Handler:    _UserService_Register_Handler,
		},
		{
			MethodName: "Login",
			Handler:    _UserService_Login_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "grpc.proto",
}

// SMSServiceClient is the client API for SMSService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SMSServiceClient interface {
	ObtainSMSCode(ctx context.Context, in *GrpcReq, opts ...grpc.CallOption) (*GrpcResp, error)
}

type sMSServiceClient struct {
	cc *grpc.ClientConn
}

func NewSMSServiceClient(cc *grpc.ClientConn) SMSServiceClient {
	return &sMSServiceClient{cc}
}

func (c *sMSServiceClient) ObtainSMSCode(ctx context.Context, in *GrpcReq, opts ...grpc.CallOption) (*GrpcResp, error) {
	out := new(GrpcResp)
	err := c.cc.Invoke(ctx, "/com.salty.protos.SMSService/ObtainSMSCode", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SMSServiceServer is the server API for SMSService service.
type SMSServiceServer interface {
	ObtainSMSCode(context.Context, *GrpcReq) (*GrpcResp, error)
}

// UnimplementedSMSServiceServer can be embedded to have forward compatible implementations.
type UnimplementedSMSServiceServer struct {
}

func (*UnimplementedSMSServiceServer) ObtainSMSCode(ctx context.Context, req *GrpcReq) (*GrpcResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ObtainSMSCode not implemented")
}

func RegisterSMSServiceServer(s *grpc.Server, srv SMSServiceServer) {
	s.RegisterService(&_SMSService_serviceDesc, srv)
}

func _SMSService_ObtainSMSCode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GrpcReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SMSServiceServer).ObtainSMSCode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/com.salty.protos.SMSService/ObtainSMSCode",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SMSServiceServer).ObtainSMSCode(ctx, req.(*GrpcReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _SMSService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "com.salty.protos.SMSService",
	HandlerType: (*SMSServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ObtainSMSCode",
			Handler:    _SMSService_ObtainSMSCode_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "grpc.proto",
}