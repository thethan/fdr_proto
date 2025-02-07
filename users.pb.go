// Code generated by protoc-gen-go. DO NOT EDIT.
// source: users.proto

// The package name determines the name of the directories that truss creates
// for `package echo;` truss will create the directory "echo-service".

package fdr_proto

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_struct "github.com/golang/protobuf/ptypes/struct"
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

type CreateUserRequest struct {
	User                 *User    `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateUserRequest) Reset()         { *m = CreateUserRequest{} }
func (m *CreateUserRequest) String() string { return proto.CompactTextString(m) }
func (*CreateUserRequest) ProtoMessage()    {}
func (*CreateUserRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_030765f334c86cea, []int{0}
}

func (m *CreateUserRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateUserRequest.Unmarshal(m, b)
}
func (m *CreateUserRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateUserRequest.Marshal(b, m, deterministic)
}
func (m *CreateUserRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateUserRequest.Merge(m, src)
}
func (m *CreateUserRequest) XXX_Size() int {
	return xxx_messageInfo_CreateUserRequest.Size(m)
}
func (m *CreateUserRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateUserRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateUserRequest proto.InternalMessageInfo

func (m *CreateUserRequest) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

type CreateUserResponse struct {
	Status               string   `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateUserResponse) Reset()         { *m = CreateUserResponse{} }
func (m *CreateUserResponse) String() string { return proto.CompactTextString(m) }
func (*CreateUserResponse) ProtoMessage()    {}
func (*CreateUserResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_030765f334c86cea, []int{1}
}

func (m *CreateUserResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateUserResponse.Unmarshal(m, b)
}
func (m *CreateUserResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateUserResponse.Marshal(b, m, deterministic)
}
func (m *CreateUserResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateUserResponse.Merge(m, src)
}
func (m *CreateUserResponse) XXX_Size() int {
	return xxx_messageInfo_CreateUserResponse.Size(m)
}
func (m *CreateUserResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateUserResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateUserResponse proto.InternalMessageInfo

func (m *CreateUserResponse) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

type LoginRequest struct {
	User                 *User    `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	OpenId               string   `protobuf:"bytes,2,opt,name=open_id,proto3" json:"open_id,omitempty"`
	Service              string   `protobuf:"bytes,3,opt,name=service,proto3" json:"service,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LoginRequest) Reset()         { *m = LoginRequest{} }
func (m *LoginRequest) String() string { return proto.CompactTextString(m) }
func (*LoginRequest) ProtoMessage()    {}
func (*LoginRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_030765f334c86cea, []int{2}
}

func (m *LoginRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoginRequest.Unmarshal(m, b)
}
func (m *LoginRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoginRequest.Marshal(b, m, deterministic)
}
func (m *LoginRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoginRequest.Merge(m, src)
}
func (m *LoginRequest) XXX_Size() int {
	return xxx_messageInfo_LoginRequest.Size(m)
}
func (m *LoginRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_LoginRequest.DiscardUnknown(m)
}

var xxx_messageInfo_LoginRequest proto.InternalMessageInfo

func (m *LoginRequest) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

func (m *LoginRequest) GetOpenId() string {
	if m != nil {
		return m.OpenId
	}
	return ""
}

func (m *LoginRequest) GetService() string {
	if m != nil {
		return m.Service
	}
	return ""
}

type LoginResponse struct {
	Token                *Token   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LoginResponse) Reset()         { *m = LoginResponse{} }
func (m *LoginResponse) String() string { return proto.CompactTextString(m) }
func (*LoginResponse) ProtoMessage()    {}
func (*LoginResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_030765f334c86cea, []int{3}
}

func (m *LoginResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoginResponse.Unmarshal(m, b)
}
func (m *LoginResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoginResponse.Marshal(b, m, deterministic)
}
func (m *LoginResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoginResponse.Merge(m, src)
}
func (m *LoginResponse) XXX_Size() int {
	return xxx_messageInfo_LoginResponse.Size(m)
}
func (m *LoginResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_LoginResponse.DiscardUnknown(m)
}

var xxx_messageInfo_LoginResponse proto.InternalMessageInfo

func (m *LoginResponse) GetToken() *Token {
	if m != nil {
		return m.Token
	}
	return nil
}

type CredentialRequest struct {
	Session              string   `protobuf:"bytes,1,opt,name=session,proto3" json:"session,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CredentialRequest) Reset()         { *m = CredentialRequest{} }
func (m *CredentialRequest) String() string { return proto.CompactTextString(m) }
func (*CredentialRequest) ProtoMessage()    {}
func (*CredentialRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_030765f334c86cea, []int{4}
}

func (m *CredentialRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CredentialRequest.Unmarshal(m, b)
}
func (m *CredentialRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CredentialRequest.Marshal(b, m, deterministic)
}
func (m *CredentialRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CredentialRequest.Merge(m, src)
}
func (m *CredentialRequest) XXX_Size() int {
	return xxx_messageInfo_CredentialRequest.Size(m)
}
func (m *CredentialRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CredentialRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CredentialRequest proto.InternalMessageInfo

func (m *CredentialRequest) GetSession() string {
	if m != nil {
		return m.Session
	}
	return ""
}

type CredentialResponse struct {
	Token                *Token   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CredentialResponse) Reset()         { *m = CredentialResponse{} }
func (m *CredentialResponse) String() string { return proto.CompactTextString(m) }
func (*CredentialResponse) ProtoMessage()    {}
func (*CredentialResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_030765f334c86cea, []int{5}
}

func (m *CredentialResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CredentialResponse.Unmarshal(m, b)
}
func (m *CredentialResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CredentialResponse.Marshal(b, m, deterministic)
}
func (m *CredentialResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CredentialResponse.Merge(m, src)
}
func (m *CredentialResponse) XXX_Size() int {
	return xxx_messageInfo_CredentialResponse.Size(m)
}
func (m *CredentialResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CredentialResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CredentialResponse proto.InternalMessageInfo

func (m *CredentialResponse) GetToken() *Token {
	if m != nil {
		return m.Token
	}
	return nil
}

type ListUserRequest struct {
	User                 *User    `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	OpenId               string   `protobuf:"bytes,2,opt,name=open_id,proto3" json:"open_id,omitempty"`
	Service              string   `protobuf:"bytes,3,opt,name=service,proto3" json:"service,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListUserRequest) Reset()         { *m = ListUserRequest{} }
func (m *ListUserRequest) String() string { return proto.CompactTextString(m) }
func (*ListUserRequest) ProtoMessage()    {}
func (*ListUserRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_030765f334c86cea, []int{6}
}

func (m *ListUserRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListUserRequest.Unmarshal(m, b)
}
func (m *ListUserRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListUserRequest.Marshal(b, m, deterministic)
}
func (m *ListUserRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListUserRequest.Merge(m, src)
}
func (m *ListUserRequest) XXX_Size() int {
	return xxx_messageInfo_ListUserRequest.Size(m)
}
func (m *ListUserRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListUserRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListUserRequest proto.InternalMessageInfo

func (m *ListUserRequest) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

func (m *ListUserRequest) GetOpenId() string {
	if m != nil {
		return m.OpenId
	}
	return ""
}

func (m *ListUserRequest) GetService() string {
	if m != nil {
		return m.Service
	}
	return ""
}

type ListUserResponse struct {
	User                 []*User   `protobuf:"bytes,1,rep,name=user,proto3" json:"user,omitempty"`
	Metadata             *Metadata `protobuf:"bytes,2,opt,name=metadata,proto3" json:"metadata,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *ListUserResponse) Reset()         { *m = ListUserResponse{} }
func (m *ListUserResponse) String() string { return proto.CompactTextString(m) }
func (*ListUserResponse) ProtoMessage()    {}
func (*ListUserResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_030765f334c86cea, []int{7}
}

func (m *ListUserResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListUserResponse.Unmarshal(m, b)
}
func (m *ListUserResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListUserResponse.Marshal(b, m, deterministic)
}
func (m *ListUserResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListUserResponse.Merge(m, src)
}
func (m *ListUserResponse) XXX_Size() int {
	return xxx_messageInfo_ListUserResponse.Size(m)
}
func (m *ListUserResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListUserResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListUserResponse proto.InternalMessageInfo

func (m *ListUserResponse) GetUser() []*User {
	if m != nil {
		return m.User
	}
	return nil
}

func (m *ListUserResponse) GetMetadata() *Metadata {
	if m != nil {
		return m.Metadata
	}
	return nil
}

type Token struct {
	AccessToken          string   `protobuf:"bytes,1,opt,name=access_token,proto3" json:"access_token,omitempty"`
	RefreshToken         string   `protobuf:"bytes,2,opt,name=refresh_token,proto3" json:"refresh_token,omitempty"`
	ExpiresIn            int32    `protobuf:"varint,3,opt,name=expires_in,proto3" json:"expires_in,omitempty"`
	Cookie               string   `protobuf:"bytes,4,opt,name=cookie,proto3" json:"cookie,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Token) Reset()         { *m = Token{} }
func (m *Token) String() string { return proto.CompactTextString(m) }
func (*Token) ProtoMessage()    {}
func (*Token) Descriptor() ([]byte, []int) {
	return fileDescriptor_030765f334c86cea, []int{8}
}

func (m *Token) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Token.Unmarshal(m, b)
}
func (m *Token) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Token.Marshal(b, m, deterministic)
}
func (m *Token) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Token.Merge(m, src)
}
func (m *Token) XXX_Size() int {
	return xxx_messageInfo_Token.Size(m)
}
func (m *Token) XXX_DiscardUnknown() {
	xxx_messageInfo_Token.DiscardUnknown(m)
}

var xxx_messageInfo_Token proto.InternalMessageInfo

func (m *Token) GetAccessToken() string {
	if m != nil {
		return m.AccessToken
	}
	return ""
}

func (m *Token) GetRefreshToken() string {
	if m != nil {
		return m.RefreshToken
	}
	return ""
}

func (m *Token) GetExpiresIn() int32 {
	if m != nil {
		return m.ExpiresIn
	}
	return 0
}

func (m *Token) GetCookie() string {
	if m != nil {
		return m.Cookie
	}
	return ""
}

type User struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Image                string   `protobuf:"bytes,3,opt,name=image,proto3" json:"image,omitempty"`
	Email                string   `protobuf:"bytes,4,opt,name=email,proto3" json:"email,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_030765f334c86cea, []int{9}
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

func (m *User) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *User) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *User) GetImage() string {
	if m != nil {
		return m.Image
	}
	return ""
}

func (m *User) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

type Metadata struct {
	Total                int32           `protobuf:"varint,1,opt,name=total,proto3" json:"total,omitempty"`
	Page                 int32           `protobuf:"varint,2,opt,name=page,proto3" json:"page,omitempty"`
	PerPage              int32           `protobuf:"varint,3,opt,name=per_page,json=perPage,proto3" json:"per_page,omitempty"`
	Filters              *_struct.Struct `protobuf:"bytes,4,opt,name=filters,proto3" json:"filters,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *Metadata) Reset()         { *m = Metadata{} }
func (m *Metadata) String() string { return proto.CompactTextString(m) }
func (*Metadata) ProtoMessage()    {}
func (*Metadata) Descriptor() ([]byte, []int) {
	return fileDescriptor_030765f334c86cea, []int{10}
}

func (m *Metadata) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Metadata.Unmarshal(m, b)
}
func (m *Metadata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Metadata.Marshal(b, m, deterministic)
}
func (m *Metadata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Metadata.Merge(m, src)
}
func (m *Metadata) XXX_Size() int {
	return xxx_messageInfo_Metadata.Size(m)
}
func (m *Metadata) XXX_DiscardUnknown() {
	xxx_messageInfo_Metadata.DiscardUnknown(m)
}

var xxx_messageInfo_Metadata proto.InternalMessageInfo

func (m *Metadata) GetTotal() int32 {
	if m != nil {
		return m.Total
	}
	return 0
}

func (m *Metadata) GetPage() int32 {
	if m != nil {
		return m.Page
	}
	return 0
}

func (m *Metadata) GetPerPage() int32 {
	if m != nil {
		return m.PerPage
	}
	return 0
}

func (m *Metadata) GetFilters() *_struct.Struct {
	if m != nil {
		return m.Filters
	}
	return nil
}

func init() {
	proto.RegisterType((*CreateUserRequest)(nil), "fdr.CreateUserRequest")
	proto.RegisterType((*CreateUserResponse)(nil), "fdr.CreateUserResponse")
	proto.RegisterType((*LoginRequest)(nil), "fdr.LoginRequest")
	proto.RegisterType((*LoginResponse)(nil), "fdr.LoginResponse")
	proto.RegisterType((*CredentialRequest)(nil), "fdr.CredentialRequest")
	proto.RegisterType((*CredentialResponse)(nil), "fdr.CredentialResponse")
	proto.RegisterType((*ListUserRequest)(nil), "fdr.ListUserRequest")
	proto.RegisterType((*ListUserResponse)(nil), "fdr.ListUserResponse")
	proto.RegisterType((*Token)(nil), "fdr.token")
	proto.RegisterType((*User)(nil), "fdr.User")
	proto.RegisterType((*Metadata)(nil), "fdr.metadata")
}

func init() {
	proto.RegisterFile("users.proto", fileDescriptor_030765f334c86cea)
}

var fileDescriptor_030765f334c86cea = []byte{
	// 552 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x53, 0x4f, 0x6f, 0xd3, 0x4e,
	0x10, 0x6d, 0xfe, 0x38, 0x69, 0x26, 0xcd, 0xef, 0x47, 0x57, 0xa5, 0x0d, 0x51, 0xa9, 0xaa, 0x15,
	0x07, 0x90, 0xc0, 0xa1, 0x41, 0x82, 0x03, 0x17, 0x04, 0xd7, 0x1e, 0x90, 0x2b, 0x38, 0x20, 0xa4,
	0x68, 0x6b, 0x8f, 0x93, 0x55, 0x63, 0xaf, 0xd9, 0x5d, 0x23, 0xce, 0xdc, 0xf8, 0xb8, 0x7c, 0x03,
	0xe4, 0xd9, 0x75, 0x9c, 0xa4, 0x1c, 0x22, 0x4e, 0xf6, 0xbc, 0x99, 0x9d, 0x37, 0x6f, 0x76, 0x1f,
	0x0c, 0x4b, 0x83, 0xda, 0x84, 0x85, 0x56, 0x56, 0xb1, 0x4e, 0x9a, 0xe8, 0xc9, 0xf9, 0x42, 0xa9,
	0xc5, 0x0a, 0xa7, 0x04, 0xdd, 0x96, 0xe9, 0xd4, 0x58, 0x5d, 0xc6, 0xd6, 0x95, 0xf0, 0x19, 0x1c,
	0x7f, 0xd0, 0x28, 0x2c, 0x7e, 0x32, 0xa8, 0x23, 0xfc, 0x56, 0xa2, 0xb1, 0xec, 0x31, 0x74, 0xab,
	0x36, 0xe3, 0xd6, 0x65, 0xeb, 0xe9, 0x70, 0x36, 0x08, 0xd3, 0x44, 0x87, 0x94, 0x27, 0x98, 0x3f,
	0x07, 0xb6, 0x79, 0xc6, 0x14, 0x2a, 0x37, 0xc8, 0x4e, 0xa1, 0x67, 0xac, 0xb0, 0xa5, 0xa1, 0x63,
	0x83, 0xc8, 0x47, 0x5c, 0xc0, 0xd1, 0xb5, 0x5a, 0xc8, 0x7c, 0xbf, 0xe6, 0x6c, 0x0c, 0x7d, 0x55,
	0x60, 0x3e, 0x97, 0xc9, 0xb8, 0x4d, 0x7d, 0xea, 0xb0, 0xca, 0x18, 0xd4, 0xdf, 0x65, 0x8c, 0xe3,
	0x8e, 0xcb, 0xf8, 0x90, 0x5f, 0xc1, 0xc8, 0x53, 0xf8, 0x59, 0x2e, 0x21, 0xb0, 0xea, 0x0e, 0x73,
	0x4f, 0x02, 0x44, 0x42, 0x48, 0xe4, 0x12, 0xfc, 0x05, 0xe9, 0x4e, 0x30, 0xb7, 0x52, 0xac, 0xea,
	0xd1, 0x88, 0xc1, 0x18, 0xa9, 0x72, 0xaf, 0xa1, 0x0e, 0xf9, 0x6b, 0x92, 0xbc, 0x2e, 0xdf, 0x9b,
	0x26, 0x81, 0xff, 0xaf, 0xa5, 0xb1, 0xfb, 0x2f, 0xf7, 0x9f, 0xf4, 0x7f, 0x85, 0x07, 0x0d, 0x8b,
	0x9f, 0xad, 0xa1, 0xe9, 0xfc, 0x8d, 0xe6, 0x19, 0x1c, 0x66, 0x68, 0x45, 0x22, 0xac, 0x20, 0x9e,
	0xe1, 0x6c, 0x44, 0x25, 0x35, 0x18, 0xad, 0xd3, 0xfc, 0x57, 0xcb, 0xcb, 0x64, 0x1c, 0x8e, 0x44,
	0x1c, 0xa3, 0x31, 0xf3, 0x46, 0xf6, 0x20, 0xda, 0xc2, 0xd8, 0x13, 0x18, 0x69, 0x4c, 0x35, 0x9a,
	0xa5, 0x2f, 0x72, 0x2a, 0xb6, 0x41, 0x76, 0x01, 0x80, 0x3f, 0x0a, 0xa9, 0xd1, 0xcc, 0x65, 0x4e,
	0x72, 0x82, 0x68, 0x03, 0xa9, 0x1e, 0x53, 0xac, 0xd4, 0x9d, 0xc4, 0x71, 0xd7, 0x3d, 0x26, 0x17,
	0xf1, 0xcf, 0xd0, 0xad, 0x44, 0xb0, 0xff, 0xa0, 0x2d, 0x13, 0xcf, 0xdf, 0x96, 0x09, 0x63, 0xd0,
	0xcd, 0x45, 0x86, 0x9e, 0x8c, 0xfe, 0xd9, 0x09, 0x04, 0x32, 0x13, 0x8b, 0x7a, 0x5b, 0x2e, 0xa8,
	0x50, 0xcc, 0x84, 0x5c, 0xf9, 0xc6, 0x2e, 0xe0, 0x3f, 0x5b, 0xcd, 0x3e, 0xaa, 0x12, 0xab, 0xac,
	0x58, 0x51, 0xff, 0x20, 0x72, 0x41, 0x45, 0x51, 0x54, 0xdd, 0xda, 0x04, 0xd2, 0x3f, 0x7b, 0x04,
	0x87, 0x05, 0xea, 0x79, 0x51, 0xb3, 0x04, 0x51, 0xbf, 0x40, 0xfd, 0xb1, 0x4a, 0x5d, 0x41, 0x3f,
	0x95, 0x2b, 0x8b, 0xda, 0x10, 0xd3, 0x70, 0x76, 0x16, 0x3a, 0x23, 0x86, 0xb5, 0x11, 0xc3, 0x1b,
	0x32, 0x62, 0x54, 0xd7, 0xcd, 0x7e, 0xb7, 0x20, 0xa8, 0xd4, 0x19, 0xf6, 0x16, 0x7a, 0xce, 0x61,
	0xec, 0x94, 0x6e, 0xe5, 0x9e, 0x45, 0x27, 0x67, 0xf7, 0x70, 0x77, 0xef, 0xfc, 0x80, 0xbd, 0x81,
	0xde, 0x0d, 0x0a, 0x1d, 0x2f, 0xd9, 0x09, 0x15, 0xed, 0x3c, 0xc0, 0xc9, 0xc3, 0x1d, 0x74, 0x7d,
	0xf0, 0x25, 0x04, 0x64, 0x23, 0x76, 0xec, 0x2a, 0x36, 0x5c, 0x3b, 0x61, 0x9b, 0xd0, 0xfa, 0xc4,
	0x3b, 0x18, 0x36, 0xb6, 0x30, 0xcd, 0xb0, 0xdb, 0xbe, 0x6a, 0x86, 0xdd, 0x31, 0x10, 0x3f, 0x78,
	0x7f, 0xf1, 0xe5, 0x7c, 0x21, 0xed, 0xb2, 0xbc, 0x0d, 0x63, 0x95, 0x4d, 0xed, 0x12, 0xed, 0x52,
	0xe4, 0xd3, 0x34, 0xd1, 0x73, 0xb7, 0xa6, 0x1e, 0x7d, 0x5e, 0xfd, 0x09, 0x00, 0x00, 0xff, 0xff,
	0x58, 0x3a, 0xb9, 0x85, 0xd8, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// UsersClient is the client API for Users service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type UsersClient interface {
	Create(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*CreateUserResponse, error)
	Search(ctx context.Context, in *ListUserRequest, opts ...grpc.CallOption) (*ListUserResponse, error)
	Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error)
	Credentials(ctx context.Context, in *CredentialRequest, opts ...grpc.CallOption) (*CredentialResponse, error)
}

type usersClient struct {
	cc grpc.ClientConnInterface
}

func NewUsersClient(cc grpc.ClientConnInterface) UsersClient {
	return &usersClient{cc}
}

func (c *usersClient) Create(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*CreateUserResponse, error) {
	out := new(CreateUserResponse)
	err := c.cc.Invoke(ctx, "/fdr.Users/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usersClient) Search(ctx context.Context, in *ListUserRequest, opts ...grpc.CallOption) (*ListUserResponse, error) {
	out := new(ListUserResponse)
	err := c.cc.Invoke(ctx, "/fdr.Users/Search", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usersClient) Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error) {
	out := new(LoginResponse)
	err := c.cc.Invoke(ctx, "/fdr.Users/Login", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usersClient) Credentials(ctx context.Context, in *CredentialRequest, opts ...grpc.CallOption) (*CredentialResponse, error) {
	out := new(CredentialResponse)
	err := c.cc.Invoke(ctx, "/fdr.Users/Credentials", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UsersServer is the server API for Users service.
type UsersServer interface {
	Create(context.Context, *CreateUserRequest) (*CreateUserResponse, error)
	Search(context.Context, *ListUserRequest) (*ListUserResponse, error)
	Login(context.Context, *LoginRequest) (*LoginResponse, error)
	Credentials(context.Context, *CredentialRequest) (*CredentialResponse, error)
}

// UnimplementedUsersServer can be embedded to have forward compatible implementations.
type UnimplementedUsersServer struct {
}

func (*UnimplementedUsersServer) Create(ctx context.Context, req *CreateUserRequest) (*CreateUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (*UnimplementedUsersServer) Search(ctx context.Context, req *ListUserRequest) (*ListUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Search not implemented")
}
func (*UnimplementedUsersServer) Login(ctx context.Context, req *LoginRequest) (*LoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (*UnimplementedUsersServer) Credentials(ctx context.Context, req *CredentialRequest) (*CredentialResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Credentials not implemented")
}

func RegisterUsersServer(s *grpc.Server, srv UsersServer) {
	s.RegisterService(&_Users_serviceDesc, srv)
}

func _Users_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsersServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fdr.Users/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsersServer).Create(ctx, req.(*CreateUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Users_Search_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsersServer).Search(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fdr.Users/Search",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsersServer).Search(ctx, req.(*ListUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Users_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsersServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fdr.Users/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsersServer).Login(ctx, req.(*LoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Users_Credentials_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CredentialRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsersServer).Credentials(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fdr.Users/Credentials",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsersServer).Credentials(ctx, req.(*CredentialRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Users_serviceDesc = grpc.ServiceDesc{
	ServiceName: "fdr.Users",
	HandlerType: (*UsersServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _Users_Create_Handler,
		},
		{
			MethodName: "Search",
			Handler:    _Users_Search_Handler,
		},
		{
			MethodName: "Login",
			Handler:    _Users_Login_Handler,
		},
		{
			MethodName: "Credentials",
			Handler:    _Users_Credentials_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "users.proto",
}
