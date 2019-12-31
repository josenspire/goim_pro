// Code generated by protoc-gen-go. DO NOT EDIT.
// source: salty.proto

package protos_salty

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

type User_Role int32

const (
	User_DEFAULT     User_Role = 0
	User_ADMIN       User_Role = 10
	User_SUPER_ADMIN User_Role = 99
)

var User_Role_name = map[int32]string{
	0:  "DEFAULT",
	10: "ADMIN",
	99: "SUPER_ADMIN",
}

var User_Role_value = map[string]int32{
	"DEFAULT":     0,
	"ADMIN":       10,
	"SUPER_ADMIN": 99,
}

func (x User_Role) String() string {
	return proto.EnumName(User_Role_name, int32(x))
}

func (User_Role) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_f0b0587afcdc51fc, []int{3, 0}
}

type User_Sex int32

const (
	User_MALE   User_Sex = 0
	User_FEMALE User_Sex = 1
)

var User_Sex_name = map[int32]string{
	0: "MALE",
	1: "FEMALE",
}

var User_Sex_value = map[string]int32{
	"MALE":   0,
	"FEMALE": 1,
}

func (x User_Sex) String() string {
	return proto.EnumName(User_Sex_name, int32(x))
}

func (User_Sex) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_f0b0587afcdc51fc, []int{3, 1}
}

type User_Status int32

const (
	User_ACTIVE   User_Status = 0
	User_INACTIVE User_Status = 1
)

var User_Status_name = map[int32]string{
	0: "ACTIVE",
	1: "INACTIVE",
}

var User_Status_value = map[string]int32{
	"ACTIVE":   0,
	"INACTIVE": 1,
}

func (x User_Status) String() string {
	return proto.EnumName(User_Status_name, int32(x))
}

func (User_Status) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_f0b0587afcdc51fc, []int{3, 2}
}

// Request
type UserServiceRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserServiceRequest) Reset()         { *m = UserServiceRequest{} }
func (m *UserServiceRequest) String() string { return proto.CompactTextString(m) }
func (*UserServiceRequest) ProtoMessage()    {}
func (*UserServiceRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_f0b0587afcdc51fc, []int{0}
}

func (m *UserServiceRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserServiceRequest.Unmarshal(m, b)
}
func (m *UserServiceRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserServiceRequest.Marshal(b, m, deterministic)
}
func (m *UserServiceRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserServiceRequest.Merge(m, src)
}
func (m *UserServiceRequest) XXX_Size() int {
	return xxx_messageInfo_UserServiceRequest.Size(m)
}
func (m *UserServiceRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UserServiceRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UserServiceRequest proto.InternalMessageInfo

// Response
type UserServiceResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserServiceResponse) Reset()         { *m = UserServiceResponse{} }
func (m *UserServiceResponse) String() string { return proto.CompactTextString(m) }
func (*UserServiceResponse) ProtoMessage()    {}
func (*UserServiceResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_f0b0587afcdc51fc, []int{1}
}

func (m *UserServiceResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserServiceResponse.Unmarshal(m, b)
}
func (m *UserServiceResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserServiceResponse.Marshal(b, m, deterministic)
}
func (m *UserServiceResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserServiceResponse.Merge(m, src)
}
func (m *UserServiceResponse) XXX_Size() int {
	return xxx_messageInfo_UserServiceResponse.Size(m)
}
func (m *UserServiceResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UserServiceResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UserServiceResponse proto.InternalMessageInfo

type ServerCommonResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ServerCommonResponse) Reset()         { *m = ServerCommonResponse{} }
func (m *ServerCommonResponse) String() string { return proto.CompactTextString(m) }
func (*ServerCommonResponse) ProtoMessage()    {}
func (*ServerCommonResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_f0b0587afcdc51fc, []int{2}
}

func (m *ServerCommonResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ServerCommonResponse.Unmarshal(m, b)
}
func (m *ServerCommonResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ServerCommonResponse.Marshal(b, m, deterministic)
}
func (m *ServerCommonResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ServerCommonResponse.Merge(m, src)
}
func (m *ServerCommonResponse) XXX_Size() int {
	return xxx_messageInfo_ServerCommonResponse.Size(m)
}
func (m *ServerCommonResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ServerCommonResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ServerCommonResponse proto.InternalMessageInfo

// user information
type User struct {
	UserId      int64             `protobuf:"varint,1,opt,name=userId,proto3" json:"userId,omitempty"`
	Role        User_Role         `protobuf:"varint,2,opt,name=role,proto3,enum=protos.salty.User_Role" json:"role,omitempty"`
	UserProfile *User_UserProfile `protobuf:"bytes,3,opt,name=userProfile,proto3" json:"userProfile,omitempty"`
	Password    string            `protobuf:"bytes,4,opt,name=password,proto3" json:"password,omitempty"`
	Status      User_Status       `protobuf:"varint,5,opt,name=status,proto3,enum=protos.salty.User_Status" json:"status,omitempty"`
	//    repeated Address address = 6;
	CreatedAt            *timestamp.Timestamp `protobuf:"bytes,6,opt,name=CreatedAt,proto3" json:"CreatedAt,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_f0b0587afcdc51fc, []int{3}
}

func (m *User) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_User.Unmarshal(m, b)
}
func (m *User) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_User.Marshal(b, m, deterministic)
}
func (m *User) XXX_Merge(src proto.Message) {
	xxx_messageInfo_User.Merge(m, src)
}
func (m *User) XXX_Size() int {
	return xxx_messageInfo_User.Size(m)
}
func (m *User) XXX_DiscardUnknown() {
	xxx_messageInfo_User.DiscardUnknown(m)
}

var xxx_messageInfo_User proto.InternalMessageInfo

func (m *User) GetUserId() int64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *User) GetRole() User_Role {
	if m != nil {
		return m.Role
	}
	return User_DEFAULT
}

func (m *User) GetUserProfile() *User_UserProfile {
	if m != nil {
		return m.UserProfile
	}
	return nil
}

func (m *User) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *User) GetStatus() User_Status {
	if m != nil {
		return m.Status
	}
	return User_ACTIVE
}

func (m *User) GetCreatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

type User_UserProfile struct {
	Telephone            string               `protobuf:"bytes,1,opt,name=telephone,proto3" json:"telephone,omitempty"`
	Email                string               `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
	Username             string               `protobuf:"bytes,3,opt,name=username,proto3" json:"username,omitempty"`
	Nickname             string               `protobuf:"bytes,4,opt,name=nickname,proto3" json:"nickname,omitempty"`
	Avatar               string               `protobuf:"bytes,5,opt,name=avatar,proto3" json:"avatar,omitempty"`
	Signature            string               `protobuf:"bytes,6,opt,name=signature,proto3" json:"signature,omitempty"`
	Sex                  User_Sex             `protobuf:"varint,7,opt,name=sex,proto3,enum=protos.salty.User_Sex" json:"sex,omitempty"`
	Birthday             *timestamp.Timestamp `protobuf:"bytes,8,opt,name=Birthday,proto3" json:"Birthday,omitempty"`
	Location             string               `protobuf:"bytes,9,opt,name=location,proto3" json:"location,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *User_UserProfile) Reset()         { *m = User_UserProfile{} }
func (m *User_UserProfile) String() string { return proto.CompactTextString(m) }
func (*User_UserProfile) ProtoMessage()    {}
func (*User_UserProfile) Descriptor() ([]byte, []int) {
	return fileDescriptor_f0b0587afcdc51fc, []int{3, 0}
}

func (m *User_UserProfile) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_User_UserProfile.Unmarshal(m, b)
}
func (m *User_UserProfile) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_User_UserProfile.Marshal(b, m, deterministic)
}
func (m *User_UserProfile) XXX_Merge(src proto.Message) {
	xxx_messageInfo_User_UserProfile.Merge(m, src)
}
func (m *User_UserProfile) XXX_Size() int {
	return xxx_messageInfo_User_UserProfile.Size(m)
}
func (m *User_UserProfile) XXX_DiscardUnknown() {
	xxx_messageInfo_User_UserProfile.DiscardUnknown(m)
}

var xxx_messageInfo_User_UserProfile proto.InternalMessageInfo

func (m *User_UserProfile) GetTelephone() string {
	if m != nil {
		return m.Telephone
	}
	return ""
}

func (m *User_UserProfile) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *User_UserProfile) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *User_UserProfile) GetNickname() string {
	if m != nil {
		return m.Nickname
	}
	return ""
}

func (m *User_UserProfile) GetAvatar() string {
	if m != nil {
		return m.Avatar
	}
	return ""
}

func (m *User_UserProfile) GetSignature() string {
	if m != nil {
		return m.Signature
	}
	return ""
}

func (m *User_UserProfile) GetSex() User_Sex {
	if m != nil {
		return m.Sex
	}
	return User_MALE
}

func (m *User_UserProfile) GetBirthday() *timestamp.Timestamp {
	if m != nil {
		return m.Birthday
	}
	return nil
}

func (m *User_UserProfile) GetLocation() string {
	if m != nil {
		return m.Location
	}
	return ""
}

func init() {
	proto.RegisterEnum("protos.salty.User_Role", User_Role_name, User_Role_value)
	proto.RegisterEnum("protos.salty.User_Sex", User_Sex_name, User_Sex_value)
	proto.RegisterEnum("protos.salty.User_Status", User_Status_name, User_Status_value)
	proto.RegisterType((*UserServiceRequest)(nil), "protos.salty.UserServiceRequest")
	proto.RegisterType((*UserServiceResponse)(nil), "protos.salty.UserServiceResponse")
	proto.RegisterType((*ServerCommonResponse)(nil), "protos.salty.ServerCommonResponse")
	proto.RegisterType((*User)(nil), "protos.salty.User")
	proto.RegisterType((*User_UserProfile)(nil), "protos.salty.User.UserProfile")
}

func init() { proto.RegisterFile("salty.proto", fileDescriptor_f0b0587afcdc51fc) }

var fileDescriptor_f0b0587afcdc51fc = []byte{
	// 523 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x53, 0x5d, 0x6f, 0xda, 0x30,
	0x14, 0x25, 0x05, 0xd2, 0xe4, 0xa6, 0xda, 0x90, 0xc7, 0x58, 0xc6, 0xa6, 0x0d, 0xe5, 0x09, 0x69,
	0x52, 0x50, 0x99, 0x34, 0xed, 0x71, 0x0c, 0xa8, 0x84, 0x44, 0xab, 0xca, 0xc0, 0x5e, 0x27, 0x17,
	0x6e, 0x69, 0xb4, 0x24, 0x66, 0xb6, 0xd3, 0xd1, 0x3f, 0xb0, 0x7f, 0xb3, 0xf7, 0xfd, 0xbc, 0xc9,
	0x4e, 0xf8, 0xa8, 0x16, 0xa9, 0x4f, 0x70, 0x8e, 0xcf, 0xb5, 0xce, 0x39, 0xd7, 0x01, 0x4f, 0xb2,
	0x58, 0x3d, 0x84, 0x1b, 0xc1, 0x15, 0x27, 0x67, 0xe6, 0x47, 0x86, 0x86, 0x6b, 0xbf, 0x5f, 0x73,
	0xbe, 0x8e, 0xb1, 0x67, 0xc8, 0x9b, 0xec, 0xb6, 0xa7, 0xa2, 0x04, 0xa5, 0x62, 0xc9, 0x26, 0x97,
	0x07, 0x4d, 0x20, 0x0b, 0x89, 0x62, 0x86, 0xe2, 0x3e, 0x5a, 0x22, 0xc5, 0x9f, 0x19, 0x4a, 0x15,
	0xbc, 0x84, 0x17, 0x8f, 0x58, 0xb9, 0xe1, 0xa9, 0xc4, 0xa0, 0x05, 0x4d, 0x4d, 0xa1, 0x18, 0xf2,
	0x24, 0xe1, 0xe9, 0x9e, 0xff, 0x5b, 0x87, 0x9a, 0xd6, 0x93, 0x16, 0xd8, 0x99, 0x44, 0x31, 0x59,
	0xf9, 0x56, 0xc7, 0xea, 0x56, 0x69, 0x81, 0xc8, 0x07, 0xa8, 0x09, 0x1e, 0xa3, 0x7f, 0xd2, 0xb1,
	0xba, 0xcf, 0xfa, 0xaf, 0xc2, 0x63, 0x8f, 0xa1, 0x9e, 0x0c, 0x29, 0x8f, 0x91, 0x1a, 0x11, 0xf9,
	0x02, 0x9e, 0x1e, 0xbb, 0x16, 0xfc, 0x36, 0x8a, 0xd1, 0xaf, 0x76, 0xac, 0xae, 0xd7, 0x7f, 0x57,
	0x32, 0xb3, 0x38, 0xa8, 0xe8, 0xf1, 0x08, 0x69, 0x83, 0xb3, 0x61, 0x52, 0xfe, 0xe2, 0x62, 0xe5,
	0xd7, 0x3a, 0x56, 0xd7, 0xa5, 0x7b, 0x4c, 0xce, 0xc1, 0x96, 0x8a, 0xa9, 0x4c, 0xfa, 0x75, 0x63,
	0xe6, 0x75, 0xc9, 0xc5, 0x33, 0x23, 0xa0, 0x85, 0x90, 0x7c, 0x06, 0x77, 0x28, 0x90, 0x29, 0x5c,
	0x0d, 0x94, 0x6f, 0x1b, 0x3b, 0xed, 0x30, 0x2f, 0x36, 0xdc, 0x15, 0x1b, 0xce, 0x77, 0xc5, 0xd2,
	0x83, 0xb8, 0xfd, 0xe7, 0x04, 0xbc, 0x23, 0x97, 0xe4, 0x2d, 0xb8, 0x0a, 0x63, 0xdc, 0xdc, 0xf1,
	0x14, 0x4d, 0x45, 0x2e, 0x3d, 0x10, 0xa4, 0x09, 0x75, 0x4c, 0x58, 0x14, 0x9b, 0x9a, 0x5c, 0x9a,
	0x03, 0x1d, 0x46, 0x67, 0x4b, 0x59, 0x92, 0x77, 0xe1, 0xd2, 0x3d, 0xd6, 0x67, 0x69, 0xb4, 0xfc,
	0x61, 0xce, 0x8a, 0xa0, 0x3b, 0xac, 0x77, 0xc1, 0xee, 0x99, 0x62, 0xc2, 0x04, 0x75, 0x69, 0x81,
	0xb4, 0x07, 0x19, 0xad, 0x53, 0xa6, 0x32, 0x81, 0x26, 0x8d, 0x4b, 0x0f, 0x04, 0xe9, 0x42, 0x55,
	0xe2, 0xd6, 0x3f, 0x35, 0xdd, 0xb4, 0xca, 0xba, 0xc1, 0x2d, 0xd5, 0x12, 0xf2, 0x09, 0x9c, 0xaf,
	0x91, 0x50, 0x77, 0x2b, 0xf6, 0xe0, 0x3b, 0x4f, 0x96, 0xb2, 0xd7, 0x6a, 0xcf, 0x31, 0x5f, 0x32,
	0x15, 0xf1, 0xd4, 0x77, 0x73, 0xcf, 0x3b, 0x1c, 0xf4, 0xa0, 0xa6, 0x1f, 0x02, 0xf1, 0xe0, 0x74,
	0x34, 0xbe, 0x18, 0x2c, 0xa6, 0xf3, 0x46, 0x85, 0xb8, 0x50, 0x1f, 0x8c, 0x2e, 0x27, 0x57, 0x0d,
	0x20, 0xcf, 0xc1, 0x9b, 0x2d, 0xae, 0xc7, 0xf4, 0x7b, 0x4e, 0x2c, 0x83, 0x37, 0x50, 0x9d, 0xe1,
	0x96, 0x38, 0x50, 0xbb, 0x1c, 0x4c, 0xc7, 0x8d, 0x0a, 0x01, 0xb0, 0x2f, 0xc6, 0xe6, 0xbf, 0x15,
	0x04, 0x60, 0xe7, 0x9b, 0xd4, 0xec, 0x60, 0x38, 0x9f, 0x7c, 0xd3, 0x8a, 0x33, 0x70, 0x26, 0x57,
	0x05, 0xb2, 0xfa, 0xbf, 0xad, 0x7c, 0x43, 0xc5, 0x53, 0x27, 0x23, 0x70, 0x28, 0xae, 0x23, 0xa9,
	0x50, 0x10, 0xf2, 0x7f, 0xfc, 0x76, 0xf0, 0x98, 0x2b, 0xfd, 0x1c, 0x2a, 0xe4, 0x1c, 0xea, 0x53,
	0xbe, 0x8e, 0xd2, 0xd2, 0x2b, 0x4a, 0xb8, 0xa0, 0x72, 0x63, 0x1b, 0xf2, 0xe3, 0xbf, 0x00, 0x00,
	0x00, 0xff, 0xff, 0x6c, 0x38, 0x88, 0xe6, 0xcd, 0x03, 0x00, 0x00,
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
	// user register service
	Register(ctx context.Context, in *User, opts ...grpc.CallOption) (*ServerCommonResponse, error)
	Login(ctx context.Context, in *User, opts ...grpc.CallOption) (*User, error)
}

type userServiceClient struct {
	cc *grpc.ClientConn
}

func NewUserServiceClient(cc *grpc.ClientConn) UserServiceClient {
	return &userServiceClient{cc}
}

func (c *userServiceClient) Register(ctx context.Context, in *User, opts ...grpc.CallOption) (*ServerCommonResponse, error) {
	out := new(ServerCommonResponse)
	err := c.cc.Invoke(ctx, "/protos.salty.UserService/Register", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) Login(ctx context.Context, in *User, opts ...grpc.CallOption) (*User, error) {
	out := new(User)
	err := c.cc.Invoke(ctx, "/protos.salty.UserService/Login", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServiceServer is the server API for UserService service.
type UserServiceServer interface {
	// user register service
	Register(context.Context, *User) (*ServerCommonResponse, error)
	Login(context.Context, *User) (*User, error)
}

// UnimplementedUserServiceServer can be embedded to have forward compatible implementations.
type UnimplementedUserServiceServer struct {
}

func (*UnimplementedUserServiceServer) Register(ctx context.Context, req *User) (*ServerCommonResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Register not implemented")
}
func (*UnimplementedUserServiceServer) Login(ctx context.Context, req *User) (*User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}

func RegisterUserServiceServer(s *grpc.Server, srv UserServiceServer) {
	s.RegisterService(&_UserService_serviceDesc, srv)
}

func _UserService_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(User)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.salty.UserService/Register",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).Register(ctx, req.(*User))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(User)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.salty.UserService/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).Login(ctx, req.(*User))
	}
	return interceptor(ctx, in, info, handler)
}

var _UserService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "protos.salty.UserService",
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
	Metadata: "salty.proto",
}
