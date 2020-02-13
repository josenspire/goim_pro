// Code generated by protoc-gen-go. DO NOT EDIT.
// source: user.proto

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

type RegisterReq struct {
	Password             string       `protobuf:"bytes,1,opt,name=password,proto3" json:"password,omitempty"`
	VerificationCode     string       `protobuf:"bytes,2,opt,name=verificationCode,proto3" json:"verificationCode,omitempty"`
	Profile              *UserProfile `protobuf:"bytes,3,opt,name=profile,proto3" json:"profile,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *RegisterReq) Reset()         { *m = RegisterReq{} }
func (m *RegisterReq) String() string { return proto.CompactTextString(m) }
func (*RegisterReq) ProtoMessage()    {}
func (*RegisterReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{0}
}

func (m *RegisterReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegisterReq.Unmarshal(m, b)
}
func (m *RegisterReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegisterReq.Marshal(b, m, deterministic)
}
func (m *RegisterReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegisterReq.Merge(m, src)
}
func (m *RegisterReq) XXX_Size() int {
	return xxx_messageInfo_RegisterReq.Size(m)
}
func (m *RegisterReq) XXX_DiscardUnknown() {
	xxx_messageInfo_RegisterReq.DiscardUnknown(m)
}

var xxx_messageInfo_RegisterReq proto.InternalMessageInfo

func (m *RegisterReq) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *RegisterReq) GetVerificationCode() string {
	if m != nil {
		return m.VerificationCode
	}
	return ""
}

func (m *RegisterReq) GetProfile() *UserProfile {
	if m != nil {
		return m.Profile
	}
	return nil
}

type RegisterResp struct {
	Profile              *UserProfile `protobuf:"bytes,1,opt,name=profile,proto3" json:"profile,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *RegisterResp) Reset()         { *m = RegisterResp{} }
func (m *RegisterResp) String() string { return proto.CompactTextString(m) }
func (*RegisterResp) ProtoMessage()    {}
func (*RegisterResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{1}
}

func (m *RegisterResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegisterResp.Unmarshal(m, b)
}
func (m *RegisterResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegisterResp.Marshal(b, m, deterministic)
}
func (m *RegisterResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegisterResp.Merge(m, src)
}
func (m *RegisterResp) XXX_Size() int {
	return xxx_messageInfo_RegisterResp.Size(m)
}
func (m *RegisterResp) XXX_DiscardUnknown() {
	xxx_messageInfo_RegisterResp.DiscardUnknown(m)
}

var xxx_messageInfo_RegisterResp proto.InternalMessageInfo

func (m *RegisterResp) GetProfile() *UserProfile {
	if m != nil {
		return m.Profile
	}
	return nil
}

type LoginReq struct {
	Password         string `protobuf:"bytes,1,opt,name=password,proto3" json:"password,omitempty"`
	VerificationCode string `protobuf:"bytes,2,opt,name=verificationCode,proto3" json:"verificationCode,omitempty"`
	// Types that are valid to be assigned to TargetAccount:
	//	*LoginReq_Telephone
	//	*LoginReq_Email
	TargetAccount        isLoginReq_TargetAccount `protobuf_oneof:"targetAccount"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *LoginReq) Reset()         { *m = LoginReq{} }
func (m *LoginReq) String() string { return proto.CompactTextString(m) }
func (*LoginReq) ProtoMessage()    {}
func (*LoginReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{2}
}

func (m *LoginReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoginReq.Unmarshal(m, b)
}
func (m *LoginReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoginReq.Marshal(b, m, deterministic)
}
func (m *LoginReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoginReq.Merge(m, src)
}
func (m *LoginReq) XXX_Size() int {
	return xxx_messageInfo_LoginReq.Size(m)
}
func (m *LoginReq) XXX_DiscardUnknown() {
	xxx_messageInfo_LoginReq.DiscardUnknown(m)
}

var xxx_messageInfo_LoginReq proto.InternalMessageInfo

func (m *LoginReq) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *LoginReq) GetVerificationCode() string {
	if m != nil {
		return m.VerificationCode
	}
	return ""
}

type isLoginReq_TargetAccount interface {
	isLoginReq_TargetAccount()
}

type LoginReq_Telephone struct {
	Telephone string `protobuf:"bytes,3,opt,name=telephone,proto3,oneof"`
}

type LoginReq_Email struct {
	Email string `protobuf:"bytes,4,opt,name=email,proto3,oneof"`
}

func (*LoginReq_Telephone) isLoginReq_TargetAccount() {}

func (*LoginReq_Email) isLoginReq_TargetAccount() {}

func (m *LoginReq) GetTargetAccount() isLoginReq_TargetAccount {
	if m != nil {
		return m.TargetAccount
	}
	return nil
}

func (m *LoginReq) GetTelephone() string {
	if x, ok := m.GetTargetAccount().(*LoginReq_Telephone); ok {
		return x.Telephone
	}
	return ""
}

func (m *LoginReq) GetEmail() string {
	if x, ok := m.GetTargetAccount().(*LoginReq_Email); ok {
		return x.Email
	}
	return ""
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*LoginReq) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*LoginReq_Telephone)(nil),
		(*LoginReq_Email)(nil),
	}
}

type LoginResp struct {
	Token                string       `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	Profile              *UserProfile `protobuf:"bytes,2,opt,name=profile,proto3" json:"profile,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *LoginResp) Reset()         { *m = LoginResp{} }
func (m *LoginResp) String() string { return proto.CompactTextString(m) }
func (*LoginResp) ProtoMessage()    {}
func (*LoginResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{3}
}

func (m *LoginResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoginResp.Unmarshal(m, b)
}
func (m *LoginResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoginResp.Marshal(b, m, deterministic)
}
func (m *LoginResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoginResp.Merge(m, src)
}
func (m *LoginResp) XXX_Size() int {
	return xxx_messageInfo_LoginResp.Size(m)
}
func (m *LoginResp) XXX_DiscardUnknown() {
	xxx_messageInfo_LoginResp.DiscardUnknown(m)
}

var xxx_messageInfo_LoginResp proto.InternalMessageInfo

func (m *LoginResp) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *LoginResp) GetProfile() *UserProfile {
	if m != nil {
		return m.Profile
	}
	return nil
}

type ResetPasswordReq struct {
	OldPassword      string `protobuf:"bytes,1,opt,name=oldPassword,proto3" json:"oldPassword,omitempty"`
	NewPassword      string `protobuf:"bytes,2,opt,name=newPassword,proto3" json:"newPassword,omitempty"`
	VerificationCode string `protobuf:"bytes,3,opt,name=verificationCode,proto3" json:"verificationCode,omitempty"`
	// Types that are valid to be assigned to TargetAccount:
	//	*ResetPasswordReq_Telephone
	//	*ResetPasswordReq_Email
	TargetAccount        isResetPasswordReq_TargetAccount `protobuf_oneof:"targetAccount"`
	XXX_NoUnkeyedLiteral struct{}                         `json:"-"`
	XXX_unrecognized     []byte                           `json:"-"`
	XXX_sizecache        int32                            `json:"-"`
}

func (m *ResetPasswordReq) Reset()         { *m = ResetPasswordReq{} }
func (m *ResetPasswordReq) String() string { return proto.CompactTextString(m) }
func (*ResetPasswordReq) ProtoMessage()    {}
func (*ResetPasswordReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{4}
}

func (m *ResetPasswordReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ResetPasswordReq.Unmarshal(m, b)
}
func (m *ResetPasswordReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ResetPasswordReq.Marshal(b, m, deterministic)
}
func (m *ResetPasswordReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResetPasswordReq.Merge(m, src)
}
func (m *ResetPasswordReq) XXX_Size() int {
	return xxx_messageInfo_ResetPasswordReq.Size(m)
}
func (m *ResetPasswordReq) XXX_DiscardUnknown() {
	xxx_messageInfo_ResetPasswordReq.DiscardUnknown(m)
}

var xxx_messageInfo_ResetPasswordReq proto.InternalMessageInfo

func (m *ResetPasswordReq) GetOldPassword() string {
	if m != nil {
		return m.OldPassword
	}
	return ""
}

func (m *ResetPasswordReq) GetNewPassword() string {
	if m != nil {
		return m.NewPassword
	}
	return ""
}

func (m *ResetPasswordReq) GetVerificationCode() string {
	if m != nil {
		return m.VerificationCode
	}
	return ""
}

type isResetPasswordReq_TargetAccount interface {
	isResetPasswordReq_TargetAccount()
}

type ResetPasswordReq_Telephone struct {
	Telephone string `protobuf:"bytes,4,opt,name=telephone,proto3,oneof"`
}

type ResetPasswordReq_Email struct {
	Email string `protobuf:"bytes,5,opt,name=email,proto3,oneof"`
}

func (*ResetPasswordReq_Telephone) isResetPasswordReq_TargetAccount() {}

func (*ResetPasswordReq_Email) isResetPasswordReq_TargetAccount() {}

func (m *ResetPasswordReq) GetTargetAccount() isResetPasswordReq_TargetAccount {
	if m != nil {
		return m.TargetAccount
	}
	return nil
}

func (m *ResetPasswordReq) GetTelephone() string {
	if x, ok := m.GetTargetAccount().(*ResetPasswordReq_Telephone); ok {
		return x.Telephone
	}
	return ""
}

func (m *ResetPasswordReq) GetEmail() string {
	if x, ok := m.GetTargetAccount().(*ResetPasswordReq_Email); ok {
		return x.Email
	}
	return ""
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*ResetPasswordReq) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*ResetPasswordReq_Telephone)(nil),
		(*ResetPasswordReq_Email)(nil),
	}
}

type ResetPasswordResp struct {
	//是否需要重新登录，有的场景下重置完密码后需要重新登录，先保留这个字段
	IsNeedReLogin        bool     `protobuf:"varint,1,opt,name=isNeedReLogin,proto3" json:"isNeedReLogin,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ResetPasswordResp) Reset()         { *m = ResetPasswordResp{} }
func (m *ResetPasswordResp) String() string { return proto.CompactTextString(m) }
func (*ResetPasswordResp) ProtoMessage()    {}
func (*ResetPasswordResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{5}
}

func (m *ResetPasswordResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ResetPasswordResp.Unmarshal(m, b)
}
func (m *ResetPasswordResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ResetPasswordResp.Marshal(b, m, deterministic)
}
func (m *ResetPasswordResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResetPasswordResp.Merge(m, src)
}
func (m *ResetPasswordResp) XXX_Size() int {
	return xxx_messageInfo_ResetPasswordResp.Size(m)
}
func (m *ResetPasswordResp) XXX_DiscardUnknown() {
	xxx_messageInfo_ResetPasswordResp.DiscardUnknown(m)
}

var xxx_messageInfo_ResetPasswordResp proto.InternalMessageInfo

func (m *ResetPasswordResp) GetIsNeedReLogin() bool {
	if m != nil {
		return m.IsNeedReLogin
	}
	return false
}

type UpdateUserInfoReq struct {
	Profile              *UserProfile `protobuf:"bytes,1,opt,name=profile,proto3" json:"profile,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *UpdateUserInfoReq) Reset()         { *m = UpdateUserInfoReq{} }
func (m *UpdateUserInfoReq) String() string { return proto.CompactTextString(m) }
func (*UpdateUserInfoReq) ProtoMessage()    {}
func (*UpdateUserInfoReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{6}
}

func (m *UpdateUserInfoReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateUserInfoReq.Unmarshal(m, b)
}
func (m *UpdateUserInfoReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateUserInfoReq.Marshal(b, m, deterministic)
}
func (m *UpdateUserInfoReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateUserInfoReq.Merge(m, src)
}
func (m *UpdateUserInfoReq) XXX_Size() int {
	return xxx_messageInfo_UpdateUserInfoReq.Size(m)
}
func (m *UpdateUserInfoReq) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateUserInfoReq.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateUserInfoReq proto.InternalMessageInfo

func (m *UpdateUserInfoReq) GetProfile() *UserProfile {
	if m != nil {
		return m.Profile
	}
	return nil
}

type UpdateUserInfoResp struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateUserInfoResp) Reset()         { *m = UpdateUserInfoResp{} }
func (m *UpdateUserInfoResp) String() string { return proto.CompactTextString(m) }
func (*UpdateUserInfoResp) ProtoMessage()    {}
func (*UpdateUserInfoResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{7}
}

func (m *UpdateUserInfoResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateUserInfoResp.Unmarshal(m, b)
}
func (m *UpdateUserInfoResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateUserInfoResp.Marshal(b, m, deterministic)
}
func (m *UpdateUserInfoResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateUserInfoResp.Merge(m, src)
}
func (m *UpdateUserInfoResp) XXX_Size() int {
	return xxx_messageInfo_UpdateUserInfoResp.Size(m)
}
func (m *UpdateUserInfoResp) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateUserInfoResp.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateUserInfoResp proto.InternalMessageInfo

type GetUserInfoReq struct {
	UserId               string   `protobuf:"bytes,1,opt,name=userId,proto3" json:"userId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetUserInfoReq) Reset()         { *m = GetUserInfoReq{} }
func (m *GetUserInfoReq) String() string { return proto.CompactTextString(m) }
func (*GetUserInfoReq) ProtoMessage()    {}
func (*GetUserInfoReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{8}
}

func (m *GetUserInfoReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetUserInfoReq.Unmarshal(m, b)
}
func (m *GetUserInfoReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetUserInfoReq.Marshal(b, m, deterministic)
}
func (m *GetUserInfoReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetUserInfoReq.Merge(m, src)
}
func (m *GetUserInfoReq) XXX_Size() int {
	return xxx_messageInfo_GetUserInfoReq.Size(m)
}
func (m *GetUserInfoReq) XXX_DiscardUnknown() {
	xxx_messageInfo_GetUserInfoReq.DiscardUnknown(m)
}

var xxx_messageInfo_GetUserInfoReq proto.InternalMessageInfo

func (m *GetUserInfoReq) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

type GetUserInfoResp struct {
	Profile              *UserProfile `protobuf:"bytes,1,opt,name=profile,proto3" json:"profile,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *GetUserInfoResp) Reset()         { *m = GetUserInfoResp{} }
func (m *GetUserInfoResp) String() string { return proto.CompactTextString(m) }
func (*GetUserInfoResp) ProtoMessage()    {}
func (*GetUserInfoResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{9}
}

func (m *GetUserInfoResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetUserInfoResp.Unmarshal(m, b)
}
func (m *GetUserInfoResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetUserInfoResp.Marshal(b, m, deterministic)
}
func (m *GetUserInfoResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetUserInfoResp.Merge(m, src)
}
func (m *GetUserInfoResp) XXX_Size() int {
	return xxx_messageInfo_GetUserInfoResp.Size(m)
}
func (m *GetUserInfoResp) XXX_DiscardUnknown() {
	xxx_messageInfo_GetUserInfoResp.DiscardUnknown(m)
}

var xxx_messageInfo_GetUserInfoResp proto.InternalMessageInfo

func (m *GetUserInfoResp) GetProfile() *UserProfile {
	if m != nil {
		return m.Profile
	}
	return nil
}

type QueryUserInfoReq struct {
	// Types that are valid to be assigned to TargetAccount:
	//	*QueryUserInfoReq_Telephone
	//	*QueryUserInfoReq_Email
	TargetAccount        isQueryUserInfoReq_TargetAccount `protobuf_oneof:"targetAccount"`
	XXX_NoUnkeyedLiteral struct{}                         `json:"-"`
	XXX_unrecognized     []byte                           `json:"-"`
	XXX_sizecache        int32                            `json:"-"`
}

func (m *QueryUserInfoReq) Reset()         { *m = QueryUserInfoReq{} }
func (m *QueryUserInfoReq) String() string { return proto.CompactTextString(m) }
func (*QueryUserInfoReq) ProtoMessage()    {}
func (*QueryUserInfoReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{10}
}

func (m *QueryUserInfoReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_QueryUserInfoReq.Unmarshal(m, b)
}
func (m *QueryUserInfoReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_QueryUserInfoReq.Marshal(b, m, deterministic)
}
func (m *QueryUserInfoReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryUserInfoReq.Merge(m, src)
}
func (m *QueryUserInfoReq) XXX_Size() int {
	return xxx_messageInfo_QueryUserInfoReq.Size(m)
}
func (m *QueryUserInfoReq) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryUserInfoReq.DiscardUnknown(m)
}

var xxx_messageInfo_QueryUserInfoReq proto.InternalMessageInfo

type isQueryUserInfoReq_TargetAccount interface {
	isQueryUserInfoReq_TargetAccount()
}

type QueryUserInfoReq_Telephone struct {
	Telephone string `protobuf:"bytes,1,opt,name=telephone,proto3,oneof"`
}

type QueryUserInfoReq_Email struct {
	Email string `protobuf:"bytes,2,opt,name=email,proto3,oneof"`
}

func (*QueryUserInfoReq_Telephone) isQueryUserInfoReq_TargetAccount() {}

func (*QueryUserInfoReq_Email) isQueryUserInfoReq_TargetAccount() {}

func (m *QueryUserInfoReq) GetTargetAccount() isQueryUserInfoReq_TargetAccount {
	if m != nil {
		return m.TargetAccount
	}
	return nil
}

func (m *QueryUserInfoReq) GetTelephone() string {
	if x, ok := m.GetTargetAccount().(*QueryUserInfoReq_Telephone); ok {
		return x.Telephone
	}
	return ""
}

func (m *QueryUserInfoReq) GetEmail() string {
	if x, ok := m.GetTargetAccount().(*QueryUserInfoReq_Email); ok {
		return x.Email
	}
	return ""
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*QueryUserInfoReq) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*QueryUserInfoReq_Telephone)(nil),
		(*QueryUserInfoReq_Email)(nil),
	}
}

type QueryUserInfoResp struct {
	Profile              *UserProfile `protobuf:"bytes,1,opt,name=profile,proto3" json:"profile,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *QueryUserInfoResp) Reset()         { *m = QueryUserInfoResp{} }
func (m *QueryUserInfoResp) String() string { return proto.CompactTextString(m) }
func (*QueryUserInfoResp) ProtoMessage()    {}
func (*QueryUserInfoResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{11}
}

func (m *QueryUserInfoResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_QueryUserInfoResp.Unmarshal(m, b)
}
func (m *QueryUserInfoResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_QueryUserInfoResp.Marshal(b, m, deterministic)
}
func (m *QueryUserInfoResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryUserInfoResp.Merge(m, src)
}
func (m *QueryUserInfoResp) XXX_Size() int {
	return xxx_messageInfo_QueryUserInfoResp.Size(m)
}
func (m *QueryUserInfoResp) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryUserInfoResp.DiscardUnknown(m)
}

var xxx_messageInfo_QueryUserInfoResp proto.InternalMessageInfo

func (m *QueryUserInfoResp) GetProfile() *UserProfile {
	if m != nil {
		return m.Profile
	}
	return nil
}

func init() {
	proto.RegisterType((*RegisterReq)(nil), "com.salty.protos.RegisterReq")
	proto.RegisterType((*RegisterResp)(nil), "com.salty.protos.RegisterResp")
	proto.RegisterType((*LoginReq)(nil), "com.salty.protos.LoginReq")
	proto.RegisterType((*LoginResp)(nil), "com.salty.protos.LoginResp")
	proto.RegisterType((*ResetPasswordReq)(nil), "com.salty.protos.ResetPasswordReq")
	proto.RegisterType((*ResetPasswordResp)(nil), "com.salty.protos.ResetPasswordResp")
	proto.RegisterType((*UpdateUserInfoReq)(nil), "com.salty.protos.UpdateUserInfoReq")
	proto.RegisterType((*UpdateUserInfoResp)(nil), "com.salty.protos.UpdateUserInfoResp")
	proto.RegisterType((*GetUserInfoReq)(nil), "com.salty.protos.GetUserInfoReq")
	proto.RegisterType((*GetUserInfoResp)(nil), "com.salty.protos.GetUserInfoResp")
	proto.RegisterType((*QueryUserInfoReq)(nil), "com.salty.protos.QueryUserInfoReq")
	proto.RegisterType((*QueryUserInfoResp)(nil), "com.salty.protos.QueryUserInfoResp")
}

func init() { proto.RegisterFile("user.proto", fileDescriptor_116e343673f7ffaf) }

var fileDescriptor_116e343673f7ffaf = []byte{
	// 481 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x93, 0xcf, 0x6e, 0x13, 0x31,
	0x10, 0xc6, 0xeb, 0xa4, 0x29, 0xc9, 0x6c, 0xd3, 0x26, 0x56, 0x55, 0x85, 0x48, 0xa0, 0x68, 0xc5,
	0x21, 0xe2, 0x90, 0x43, 0x39, 0x20, 0x6e, 0xd0, 0x4a, 0xa4, 0x41, 0x11, 0x5a, 0x8c, 0x7a, 0x81,
	0xd3, 0xb2, 0x3b, 0x09, 0x16, 0x5b, 0xdb, 0xb5, 0x9d, 0x56, 0x7d, 0x09, 0x5e, 0x80, 0x47, 0xe1,
	0x11, 0x78, 0x29, 0xb4, 0xff, 0xda, 0xdd, 0x2d, 0xb4, 0x64, 0xc3, 0x71, 0x3e, 0x8f, 0x3e, 0xcf,
	0xf7, 0xf3, 0x18, 0x60, 0x65, 0x50, 0x4f, 0x94, 0x96, 0x56, 0xd2, 0x5e, 0x20, 0xcf, 0x27, 0xc6,
	0x8f, 0xec, 0x75, 0x2a, 0x98, 0xe1, 0x2e, 0x0a, 0xcb, 0xf3, 0xd2, 0xfd, 0x4e, 0xc0, 0x61, 0xb8,
	0xe4, 0xc6, 0xa2, 0x66, 0x78, 0x41, 0x87, 0xd0, 0x56, 0xbe, 0x31, 0x57, 0x52, 0x87, 0x03, 0x32,
	0x22, 0xe3, 0x0e, 0xbb, 0xa9, 0xe9, 0x73, 0xe8, 0x5d, 0xa2, 0xe6, 0x0b, 0x1e, 0xf8, 0x96, 0x4b,
	0x71, 0x22, 0x43, 0x1c, 0x34, 0x92, 0x9e, 0x3b, 0x3a, 0x7d, 0x09, 0x8f, 0x94, 0x96, 0x0b, 0x1e,
	0xe1, 0xa0, 0x39, 0x22, 0x63, 0xe7, 0xe8, 0xc9, 0xa4, 0x3a, 0xc9, 0xe4, 0xcc, 0xa0, 0xf6, 0xd2,
	0x26, 0x96, 0x77, 0xbb, 0x53, 0xd8, 0xbd, 0x9d, 0xc7, 0xa8, 0xa2, 0x11, 0x59, 0xcb, 0xe8, 0x07,
	0x81, 0xf6, 0x5c, 0x2e, 0xb9, 0xf8, 0x9f, 0xb1, 0x9e, 0x42, 0xc7, 0x62, 0x84, 0xea, 0xab, 0x14,
	0x69, 0xb0, 0xce, 0xe9, 0x16, 0xbb, 0x95, 0xe8, 0x21, 0xb4, 0xf0, 0xdc, 0xe7, 0xd1, 0x60, 0x3b,
	0x3b, 0x4b, 0xcb, 0xe3, 0x7d, 0xe8, 0x5a, 0x5f, 0x2f, 0xd1, 0xbe, 0x09, 0x02, 0xb9, 0x12, 0xd6,
	0xfd, 0x04, 0x9d, 0x6c, 0x38, 0xa3, 0xe8, 0x01, 0xb4, 0xac, 0xfc, 0x86, 0x22, 0x1b, 0x2d, 0x2d,
	0x8a, 0xc9, 0x1b, 0x6b, 0x25, 0xff, 0x45, 0xa0, 0xc7, 0xd0, 0xa0, 0xf5, 0xb2, 0x88, 0x31, 0x81,
	0x11, 0x38, 0x32, 0x0a, 0xbd, 0x32, 0x84, 0xa2, 0x14, 0x77, 0x08, 0xbc, 0xba, 0xe9, 0x48, 0x11,
	0x14, 0xa5, 0x3f, 0x92, 0x6a, 0xfe, 0x0b, 0xa9, 0xed, 0x7b, 0x48, 0xb5, 0x1e, 0x20, 0xf5, 0x0a,
	0xfa, 0x95, 0x30, 0x46, 0xd1, 0x67, 0xd0, 0xe5, 0xe6, 0x3d, 0x62, 0xc8, 0x30, 0xc1, 0x98, 0xe4,
	0x69, 0xb3, 0xb2, 0xe8, 0xce, 0xa1, 0x7f, 0xa6, 0x42, 0xdf, 0x62, 0x8c, 0x69, 0x26, 0x16, 0x32,
	0x06, 0x51, 0x7b, 0xa1, 0x0e, 0x80, 0x56, 0xdd, 0x8c, 0x72, 0xc7, 0xb0, 0x37, 0x45, 0x5b, 0xbc,
	0xe0, 0x10, 0x76, 0xe2, 0x0f, 0x38, 0xcb, 0x21, 0x67, 0x95, 0xfb, 0x0e, 0xf6, 0x4b, 0x9d, 0x9b,
	0x2c, 0xf7, 0x67, 0xe8, 0x7d, 0x58, 0xa1, 0xbe, 0x2e, 0xde, 0x5b, 0x22, 0x4e, 0xee, 0x21, 0xde,
	0x78, 0x80, 0xf8, 0x1c, 0xfa, 0x15, 0xf3, 0x0d, 0x46, 0x3d, 0xfa, 0xd9, 0x04, 0x27, 0x3e, 0xf8,
	0x88, 0xfa, 0x92, 0x07, 0x48, 0x4f, 0xa0, 0x9d, 0x7f, 0x70, 0xfa, 0xf8, 0xae, 0xc7, 0x54, 0xab,
	0x80, 0xe1, 0xc5, 0x70, 0xf8, 0xb7, 0x23, 0xa3, 0xdc, 0x2d, 0xfa, 0x1a, 0x5a, 0xc9, 0x13, 0xd7,
	0x77, 0x98, 0xc1, 0x5e, 0xf9, 0x35, 0xeb, 0x5b, 0x9d, 0x42, 0xb7, 0xb4, 0xa1, 0xf5, 0x9d, 0xde,
	0x82, 0x53, 0x58, 0x91, 0x8d, 0x26, 0x2a, 0xbd, 0x60, 0x6d, 0xa7, 0xe3, 0x86, 0x47, 0xbe, 0xec,
	0x24, 0xf2, 0x8b, 0xdf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x1f, 0x7c, 0x46, 0x37, 0x58, 0x06, 0x00,
	0x00,
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
	//注册，单纯的注册接口，注册后不会自动登录
	Register(ctx context.Context, in *GrpcReq, opts ...grpc.CallOption) (*GrpcResp, error)
	//登录
	Login(ctx context.Context, in *GrpcReq, opts ...grpc.CallOption) (*GrpcResp, error)
	//更新用户信息，目前不包含头像,更新头像的接口另外定义
	UpdateUserInfo(ctx context.Context, in *GrpcReq, opts ...grpc.CallOption) (*GrpcResp, error)
	//重置/找回密码，包含已登录和未登录状态下的重置/找回密码
	ResetPassword(ctx context.Context, in *GrpcReq, opts ...grpc.CallOption) (*GrpcResp, error)
	//通过UserId获取用户信息,包括获取自己或者别人
	GetUserInfo(ctx context.Context, in *GrpcReq, opts ...grpc.CallOption) (*GrpcResp, error)
	//通过账号获取用户信息,包括获取自己或者别人
	QueryUserInfo(ctx context.Context, in *GrpcReq, opts ...grpc.CallOption) (*GrpcResp, error)
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

func (c *userServiceClient) UpdateUserInfo(ctx context.Context, in *GrpcReq, opts ...grpc.CallOption) (*GrpcResp, error) {
	out := new(GrpcResp)
	err := c.cc.Invoke(ctx, "/com.salty.protos.UserService/UpdateUserInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) ResetPassword(ctx context.Context, in *GrpcReq, opts ...grpc.CallOption) (*GrpcResp, error) {
	out := new(GrpcResp)
	err := c.cc.Invoke(ctx, "/com.salty.protos.UserService/ResetPassword", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GetUserInfo(ctx context.Context, in *GrpcReq, opts ...grpc.CallOption) (*GrpcResp, error) {
	out := new(GrpcResp)
	err := c.cc.Invoke(ctx, "/com.salty.protos.UserService/GetUserInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) QueryUserInfo(ctx context.Context, in *GrpcReq, opts ...grpc.CallOption) (*GrpcResp, error) {
	out := new(GrpcResp)
	err := c.cc.Invoke(ctx, "/com.salty.protos.UserService/QueryUserInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServiceServer is the server API for UserService service.
type UserServiceServer interface {
	//注册，单纯的注册接口，注册后不会自动登录
	Register(context.Context, *GrpcReq) (*GrpcResp, error)
	//登录
	Login(context.Context, *GrpcReq) (*GrpcResp, error)
	//更新用户信息，目前不包含头像,更新头像的接口另外定义
	UpdateUserInfo(context.Context, *GrpcReq) (*GrpcResp, error)
	//重置/找回密码，包含已登录和未登录状态下的重置/找回密码
	ResetPassword(context.Context, *GrpcReq) (*GrpcResp, error)
	//通过UserId获取用户信息,包括获取自己或者别人
	GetUserInfo(context.Context, *GrpcReq) (*GrpcResp, error)
	//通过账号获取用户信息,包括获取自己或者别人
	QueryUserInfo(context.Context, *GrpcReq) (*GrpcResp, error)
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
func (*UnimplementedUserServiceServer) UpdateUserInfo(ctx context.Context, req *GrpcReq) (*GrpcResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUserInfo not implemented")
}
func (*UnimplementedUserServiceServer) ResetPassword(ctx context.Context, req *GrpcReq) (*GrpcResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ResetPassword not implemented")
}
func (*UnimplementedUserServiceServer) GetUserInfo(ctx context.Context, req *GrpcReq) (*GrpcResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserInfo not implemented")
}
func (*UnimplementedUserServiceServer) QueryUserInfo(ctx context.Context, req *GrpcReq) (*GrpcResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryUserInfo not implemented")
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

func _UserService_UpdateUserInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GrpcReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).UpdateUserInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/com.salty.protos.UserService/UpdateUserInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).UpdateUserInfo(ctx, req.(*GrpcReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_ResetPassword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GrpcReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).ResetPassword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/com.salty.protos.UserService/ResetPassword",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).ResetPassword(ctx, req.(*GrpcReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_GetUserInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GrpcReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetUserInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/com.salty.protos.UserService/GetUserInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetUserInfo(ctx, req.(*GrpcReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_QueryUserInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GrpcReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).QueryUserInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/com.salty.protos.UserService/QueryUserInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).QueryUserInfo(ctx, req.(*GrpcReq))
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
		{
			MethodName: "UpdateUserInfo",
			Handler:    _UserService_UpdateUserInfo_Handler,
		},
		{
			MethodName: "ResetPassword",
			Handler:    _UserService_ResetPassword_Handler,
		},
		{
			MethodName: "GetUserInfo",
			Handler:    _UserService_GetUserInfo_Handler,
		},
		{
			MethodName: "QueryUserInfo",
			Handler:    _UserService_QueryUserInfo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user.proto",
}