// Code generated by protoc-gen-go. DO NOT EDIT.
// source: draft_results.proto

// The package name determines the name of the directories that truss creates
// for `package echo;` truss will create the directory "echo-service".

package fdr_proto

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

type League int32

const (
	League_League_ALL League = 0
	League_League_NFL League = 1
	League_League_MLB League = 2
	League_League_NBA League = 3
	League_League_NHL League = 4
	League_League_MLS League = 5
)

var League_name = map[int32]string{
	0: "League_ALL",
	1: "League_NFL",
	2: "League_MLB",
	3: "League_NBA",
	4: "League_NHL",
	5: "League_MLS",
}

var League_value = map[string]int32{
	"League_ALL": 0,
	"League_NFL": 1,
	"League_MLB": 2,
	"League_NBA": 3,
	"League_NHL": 4,
	"League_MLS": 5,
}

func (x League) String() string {
	return proto.EnumName(League_name, int32(x))
}

func (League) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_446ed4cf1b75512d, []int{0}
}

type DraftType int32

const (
	DraftType_DraftType_Snake       DraftType = 0
	DraftType_DraftType_Traditional DraftType = 1
)

var DraftType_name = map[int32]string{
	0: "DraftType_Snake",
	1: "DraftType_Traditional",
}

var DraftType_value = map[string]int32{
	"DraftType_Snake":       0,
	"DraftType_Traditional": 1,
}

func (x DraftType) String() string {
	return proto.EnumName(DraftType_name, int32(x))
}

func (DraftType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_446ed4cf1b75512d, []int{1}
}

type Season struct {
	ID                   string               `protobuf:"bytes,1,opt,name=ID,json=id,proto3" json:"ID,omitempty"`
	Year                 string               `protobuf:"bytes,2,opt,name=Year,json=year,proto3" json:"Year,omitempty"`
	League               League               `protobuf:"varint,3,opt,name=league,proto3,enum=fdr.League" json:"league,omitempty"`
	DraftType            DraftType            `protobuf:"varint,4,opt,name=draft_type,proto3,enum=fdr.DraftType" json:"draft_type,omitempty"`
	DraftTime            *timestamp.Timestamp `protobuf:"bytes,5,opt,name=draftTime,json=draft_time,proto3" json:"draftTime,omitempty"`
	Users                []*User              `protobuf:"bytes,6,rep,name=users,proto3" json:"users,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Season) Reset()         { *m = Season{} }
func (m *Season) String() string { return proto.CompactTextString(m) }
func (*Season) ProtoMessage()    {}
func (*Season) Descriptor() ([]byte, []int) {
	return fileDescriptor_446ed4cf1b75512d, []int{0}
}

func (m *Season) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Season.Unmarshal(m, b)
}
func (m *Season) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Season.Marshal(b, m, deterministic)
}
func (m *Season) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Season.Merge(m, src)
}
func (m *Season) XXX_Size() int {
	return xxx_messageInfo_Season.Size(m)
}
func (m *Season) XXX_DiscardUnknown() {
	xxx_messageInfo_Season.DiscardUnknown(m)
}

var xxx_messageInfo_Season proto.InternalMessageInfo

func (m *Season) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

func (m *Season) GetYear() string {
	if m != nil {
		return m.Year
	}
	return ""
}

func (m *Season) GetLeague() League {
	if m != nil {
		return m.League
	}
	return League_League_ALL
}

func (m *Season) GetDraftType() DraftType {
	if m != nil {
		return m.DraftType
	}
	return DraftType_DraftType_Snake
}

func (m *Season) GetDraftTime() *timestamp.Timestamp {
	if m != nil {
		return m.DraftTime
	}
	return nil
}

func (m *Season) GetUsers() []*User {
	if m != nil {
		return m.Users
	}
	return nil
}

type DraftPlayerRequest struct {
	Season               *Season  `protobuf:"bytes,1,opt,name=season,proto3" json:"season,omitempty"`
	Player               *Player  `protobuf:"bytes,2,opt,name=player,proto3" json:"player,omitempty"`
	User                 *User    `protobuf:"bytes,3,opt,name=user,proto3" json:"user,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DraftPlayerRequest) Reset()         { *m = DraftPlayerRequest{} }
func (m *DraftPlayerRequest) String() string { return proto.CompactTextString(m) }
func (*DraftPlayerRequest) ProtoMessage()    {}
func (*DraftPlayerRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_446ed4cf1b75512d, []int{1}
}

func (m *DraftPlayerRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DraftPlayerRequest.Unmarshal(m, b)
}
func (m *DraftPlayerRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DraftPlayerRequest.Marshal(b, m, deterministic)
}
func (m *DraftPlayerRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DraftPlayerRequest.Merge(m, src)
}
func (m *DraftPlayerRequest) XXX_Size() int {
	return xxx_messageInfo_DraftPlayerRequest.Size(m)
}
func (m *DraftPlayerRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DraftPlayerRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DraftPlayerRequest proto.InternalMessageInfo

func (m *DraftPlayerRequest) GetSeason() *Season {
	if m != nil {
		return m.Season
	}
	return nil
}

func (m *DraftPlayerRequest) GetPlayer() *Player {
	if m != nil {
		return m.Player
	}
	return nil
}

func (m *DraftPlayerRequest) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

type DraftPlayerResponse struct {
	Season               *Season              `protobuf:"bytes,1,opt,name=season,proto3" json:"season,omitempty"`
	Player               *Player              `protobuf:"bytes,2,opt,name=player,proto3" json:"player,omitempty"`
	User                 *User                `protobuf:"bytes,3,opt,name=user,proto3" json:"user,omitempty"`
	Order                int32                `protobuf:"varint,4,opt,name=Order,json=draft_order,proto3" json:"Order,omitempty"`
	DraftTime            *timestamp.Timestamp `protobuf:"bytes,5,opt,name=draftTime,json=draft_time,proto3" json:"draftTime,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *DraftPlayerResponse) Reset()         { *m = DraftPlayerResponse{} }
func (m *DraftPlayerResponse) String() string { return proto.CompactTextString(m) }
func (*DraftPlayerResponse) ProtoMessage()    {}
func (*DraftPlayerResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_446ed4cf1b75512d, []int{2}
}

func (m *DraftPlayerResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DraftPlayerResponse.Unmarshal(m, b)
}
func (m *DraftPlayerResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DraftPlayerResponse.Marshal(b, m, deterministic)
}
func (m *DraftPlayerResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DraftPlayerResponse.Merge(m, src)
}
func (m *DraftPlayerResponse) XXX_Size() int {
	return xxx_messageInfo_DraftPlayerResponse.Size(m)
}
func (m *DraftPlayerResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DraftPlayerResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DraftPlayerResponse proto.InternalMessageInfo

func (m *DraftPlayerResponse) GetSeason() *Season {
	if m != nil {
		return m.Season
	}
	return nil
}

func (m *DraftPlayerResponse) GetPlayer() *Player {
	if m != nil {
		return m.Player
	}
	return nil
}

func (m *DraftPlayerResponse) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

func (m *DraftPlayerResponse) GetOrder() int32 {
	if m != nil {
		return m.Order
	}
	return 0
}

func (m *DraftPlayerResponse) GetDraftTime() *timestamp.Timestamp {
	if m != nil {
		return m.DraftTime
	}
	return nil
}

type DraftRequest struct {
	Season               *Season  `protobuf:"bytes,1,opt,name=season,proto3" json:"season,omitempty"`
	User                 *User    `protobuf:"bytes,2,opt,name=user,proto3" json:"user,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DraftRequest) Reset()         { *m = DraftRequest{} }
func (m *DraftRequest) String() string { return proto.CompactTextString(m) }
func (*DraftRequest) ProtoMessage()    {}
func (*DraftRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_446ed4cf1b75512d, []int{3}
}

func (m *DraftRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DraftRequest.Unmarshal(m, b)
}
func (m *DraftRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DraftRequest.Marshal(b, m, deterministic)
}
func (m *DraftRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DraftRequest.Merge(m, src)
}
func (m *DraftRequest) XXX_Size() int {
	return xxx_messageInfo_DraftRequest.Size(m)
}
func (m *DraftRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DraftRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DraftRequest proto.InternalMessageInfo

func (m *DraftRequest) GetSeason() *Season {
	if m != nil {
		return m.Season
	}
	return nil
}

func (m *DraftRequest) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

func init() {
	proto.RegisterEnum("fdr.League", League_name, League_value)
	proto.RegisterEnum("fdr.DraftType", DraftType_name, DraftType_value)
	proto.RegisterType((*Season)(nil), "fdr.Season")
	proto.RegisterType((*DraftPlayerRequest)(nil), "fdr.DraftPlayerRequest")
	proto.RegisterType((*DraftPlayerResponse)(nil), "fdr.DraftPlayerResponse")
	proto.RegisterType((*DraftRequest)(nil), "fdr.DraftRequest")
}

func init() {
	proto.RegisterFile("draft_results.proto", fileDescriptor_446ed4cf1b75512d)
}

var fileDescriptor_446ed4cf1b75512d = []byte{
	// 486 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x52, 0x4d, 0xaf, 0x93, 0x40,
	0x14, 0x2d, 0x6d, 0x21, 0xe9, 0xe5, 0xa5, 0xe2, 0x34, 0x46, 0x24, 0xea, 0x6b, 0x70, 0xd3, 0xbc,
	0x05, 0x35, 0xb8, 0xd1, 0xb8, 0x6a, 0xf3, 0x62, 0x34, 0xc1, 0x2f, 0x5a, 0x17, 0xba, 0x69, 0xe6,
	0xc9, 0xa5, 0x25, 0xf2, 0xe5, 0xcc, 0xb0, 0xe8, 0xc2, 0x9f, 0xe9, 0xd2, 0xff, 0x62, 0x98, 0x81,
	0x16, 0xa2, 0x0b, 0x8d, 0x89, 0x2b, 0xe6, 0x9c, 0x7b, 0x98, 0x7b, 0xcf, 0x99, 0x0b, 0xb3, 0x88,
	0xd1, 0x58, 0xec, 0x18, 0xf2, 0x2a, 0x15, 0xdc, 0x2b, 0x59, 0x21, 0x0a, 0x32, 0x8a, 0x23, 0xe6,
	0x5c, 0xee, 0x8b, 0x62, 0x9f, 0xe2, 0x52, 0x52, 0x37, 0x55, 0xbc, 0x14, 0x49, 0x86, 0x5c, 0xd0,
	0xac, 0x54, 0x2a, 0xe7, 0xa2, 0x4c, 0xe9, 0x11, 0x59, 0x83, 0xcc, 0x8a, 0x23, 0x6b, 0x2e, 0x70,
	0x7f, 0x68, 0x60, 0x6c, 0x90, 0xf2, 0x22, 0x27, 0x53, 0x18, 0xbe, 0xba, 0xb6, 0xb5, 0xb9, 0xb6,
	0x98, 0x84, 0xc3, 0x24, 0x22, 0x04, 0xc6, 0x1f, 0x91, 0x32, 0x7b, 0x28, 0x99, 0xf1, 0x11, 0x29,
	0x23, 0x8f, 0xc0, 0x48, 0x91, 0xee, 0x2b, 0xb4, 0x47, 0x73, 0x6d, 0x31, 0xf5, 0x4d, 0x2f, 0x8e,
	0x98, 0x17, 0x48, 0x2a, 0x6c, 0x4a, 0xc4, 0x03, 0x50, 0xb3, 0x8a, 0x63, 0x89, 0xf6, 0x58, 0x0a,
	0xa7, 0x52, 0x78, 0x5d, 0xd3, 0xdb, 0x63, 0x89, 0x61, 0x47, 0x41, 0x9e, 0xc1, 0x44, 0xa2, 0x6d,
	0x92, 0xa1, 0xad, 0xcf, 0xb5, 0x85, 0xe9, 0x3b, 0x9e, 0xf2, 0xe4, 0xb5, 0x9e, 0xbc, 0x6d, 0xeb,
	0xe9, 0xf4, 0x6b, 0x92, 0x21, 0xb9, 0x04, 0x5d, 0xba, 0xb1, 0x8d, 0xf9, 0x68, 0x61, 0xfa, 0x13,
	0xd9, 0xe5, 0x03, 0x47, 0x16, 0x2a, 0xde, 0xfd, 0x06, 0x44, 0x36, 0x7d, 0x27, 0x13, 0x08, 0xf1,
	0x6b, 0x85, 0x5c, 0xd4, 0x36, 0xb8, 0x34, 0x2d, 0xed, 0x9a, 0x8d, 0x0d, 0x95, 0x43, 0xd8, 0x94,
	0x6a, 0x91, 0xca, 0x4d, 0x26, 0xd0, 0x8a, 0x9a, 0x8b, 0x9a, 0x12, 0x79, 0x00, 0xe3, 0xba, 0x91,
	0x8c, 0xa3, 0xd7, 0x5f, 0xd2, 0xee, 0x77, 0x0d, 0x66, 0xbd, 0xfe, 0xbc, 0x2c, 0x72, 0x8e, 0xff,
	0x6d, 0x00, 0xe2, 0x80, 0xfe, 0x96, 0x45, 0xc8, 0xe4, 0x33, 0xe8, 0xa1, 0xa9, 0xb2, 0x2b, 0x6a,
	0xea, 0x1f, 0x72, 0x77, 0x43, 0xb8, 0x90, 0xb6, 0xfe, 0x2a, 0xd0, 0x76, 0xd4, 0xe1, 0x6f, 0x47,
	0xbd, 0x3a, 0x80, 0xa1, 0x16, 0x89, 0x4c, 0x01, 0xd4, 0x69, 0xb7, 0x0a, 0x02, 0x6b, 0xd0, 0xc1,
	0x6f, 0x5e, 0x04, 0x96, 0xd6, 0xc1, 0xaf, 0x83, 0xb5, 0x35, 0xec, 0xd6, 0xd7, 0x2b, 0x6b, 0xd4,
	0xc5, 0x2f, 0x03, 0x6b, 0xdc, 0xd3, 0x6f, 0x2c, 0xfd, 0xea, 0x39, 0x4c, 0x4e, 0x9b, 0x48, 0x66,
	0x70, 0xeb, 0x04, 0x76, 0x9b, 0x9c, 0x7e, 0x41, 0x6b, 0x40, 0xee, 0xc1, 0x9d, 0x33, 0xb9, 0x65,
	0x34, 0x4a, 0x44, 0x52, 0xe4, 0x34, 0xb5, 0x34, 0xff, 0x3d, 0x98, 0x9d, 0x17, 0x25, 0xeb, 0x3e,
	0xbc, 0x7b, 0xde, 0xf3, 0xde, 0xca, 0x39, 0xf6, 0xaf, 0x05, 0xb5, 0x0b, 0xee, 0xc0, 0x5f, 0x81,
	0x2e, 0x0b, 0xe4, 0x69, 0x7b, 0xb8, 0x7d, 0x56, 0xff, 0xc1, 0x05, 0x8f, 0xb5, 0xf5, 0xc3, 0x4f,
	0xf7, 0xf7, 0x89, 0x38, 0x54, 0x37, 0xde, 0xe7, 0x22, 0x5b, 0x8a, 0x03, 0x8a, 0x03, 0xcd, 0x97,
	0x71, 0xc4, 0x76, 0xea, 0x25, 0x0d, 0xf9, 0x79, 0xf2, 0x33, 0x00, 0x00, 0xff, 0xff, 0xee, 0x5b,
	0x5b, 0x39, 0x46, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// DraftPlayerClient is the client API for DraftPlayer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type DraftPlayerClient interface {
	DraftPlayer(ctx context.Context, in *DraftPlayerRequest, opts ...grpc.CallOption) (*DraftPlayerResponse, error)
}

type draftPlayerClient struct {
	cc grpc.ClientConnInterface
}

func NewDraftPlayerClient(cc grpc.ClientConnInterface) DraftPlayerClient {
	return &draftPlayerClient{cc}
}

func (c *draftPlayerClient) DraftPlayer(ctx context.Context, in *DraftPlayerRequest, opts ...grpc.CallOption) (*DraftPlayerResponse, error) {
	out := new(DraftPlayerResponse)
	err := c.cc.Invoke(ctx, "/fdr.DraftPlayer/DraftPlayer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DraftPlayerServer is the server API for DraftPlayer service.
type DraftPlayerServer interface {
	DraftPlayer(context.Context, *DraftPlayerRequest) (*DraftPlayerResponse, error)
}

// UnimplementedDraftPlayerServer can be embedded to have forward compatible implementations.
type UnimplementedDraftPlayerServer struct {
}

func (*UnimplementedDraftPlayerServer) DraftPlayer(ctx context.Context, req *DraftPlayerRequest) (*DraftPlayerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DraftPlayer not implemented")
}

func RegisterDraftPlayerServer(s *grpc.Server, srv DraftPlayerServer) {
	s.RegisterService(&_DraftPlayer_serviceDesc, srv)
}

func _DraftPlayer_DraftPlayer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DraftPlayerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DraftPlayerServer).DraftPlayer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fdr.DraftPlayer/DraftPlayer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DraftPlayerServer).DraftPlayer(ctx, req.(*DraftPlayerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _DraftPlayer_serviceDesc = grpc.ServiceDesc{
	ServiceName: "fdr.DraftPlayer",
	HandlerType: (*DraftPlayerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "DraftPlayer",
			Handler:    _DraftPlayer_DraftPlayer_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "draft_results.proto",
}

// DraftClient is the client API for Draft service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type DraftClient interface {
	Draft(ctx context.Context, in *DraftRequest, opts ...grpc.CallOption) (Draft_DraftClient, error)
}

type draftClient struct {
	cc grpc.ClientConnInterface
}

func NewDraftClient(cc grpc.ClientConnInterface) DraftClient {
	return &draftClient{cc}
}

func (c *draftClient) Draft(ctx context.Context, in *DraftRequest, opts ...grpc.CallOption) (Draft_DraftClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Draft_serviceDesc.Streams[0], "/fdr.Draft/Draft", opts...)
	if err != nil {
		return nil, err
	}
	x := &draftDraftClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Draft_DraftClient interface {
	Recv() (*DraftPlayerResponse, error)
	grpc.ClientStream
}

type draftDraftClient struct {
	grpc.ClientStream
}

func (x *draftDraftClient) Recv() (*DraftPlayerResponse, error) {
	m := new(DraftPlayerResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// DraftServer is the server API for Draft service.
type DraftServer interface {
	Draft(*DraftRequest, Draft_DraftServer) error
}

// UnimplementedDraftServer can be embedded to have forward compatible implementations.
type UnimplementedDraftServer struct {
}

func (*UnimplementedDraftServer) Draft(req *DraftRequest, srv Draft_DraftServer) error {
	return status.Errorf(codes.Unimplemented, "method Draft not implemented")
}

func RegisterDraftServer(s *grpc.Server, srv DraftServer) {
	s.RegisterService(&_Draft_serviceDesc, srv)
}

func _Draft_Draft_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(DraftRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(DraftServer).Draft(m, &draftDraftServer{stream})
}

type Draft_DraftServer interface {
	Send(*DraftPlayerResponse) error
	grpc.ServerStream
}

type draftDraftServer struct {
	grpc.ServerStream
}

func (x *draftDraftServer) Send(m *DraftPlayerResponse) error {
	return x.ServerStream.SendMsg(m)
}

var _Draft_serviceDesc = grpc.ServiceDesc{
	ServiceName: "fdr.Draft",
	HandlerType: (*DraftServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Draft",
			Handler:       _Draft_Draft_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "draft_results.proto",
}
