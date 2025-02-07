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

type DraftPlayerRequest struct {
	DraftResultID        int32    `protobuf:"varint,1,opt,name=DraftResultID,json=result_id,proto3" json:"DraftResultID,omitempty"`
	Season               *Season  `protobuf:"bytes,2,opt,name=season,proto3" json:"season,omitempty"`
	Player               *Player  `protobuf:"bytes,3,opt,name=player,proto3" json:"player,omitempty"`
	User                 *User    `protobuf:"bytes,4,opt,name=user,proto3" json:"user,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DraftPlayerRequest) Reset()         { *m = DraftPlayerRequest{} }
func (m *DraftPlayerRequest) String() string { return proto.CompactTextString(m) }
func (*DraftPlayerRequest) ProtoMessage()    {}
func (*DraftPlayerRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_446ed4cf1b75512d, []int{0}
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

func (m *DraftPlayerRequest) GetDraftResultID() int32 {
	if m != nil {
		return m.DraftResultID
	}
	return 0
}

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
	DraftResultID        int32                `protobuf:"varint,1,opt,name=DraftResultID,json=result_id,proto3" json:"DraftResultID,omitempty"`
	Season               *Season              `protobuf:"bytes,2,opt,name=season,proto3" json:"season,omitempty"`
	Player               *Player              `protobuf:"bytes,3,opt,name=player,proto3" json:"player,omitempty"`
	User                 *User                `protobuf:"bytes,4,opt,name=user,proto3" json:"user,omitempty"`
	Order                int32                `protobuf:"varint,5,opt,name=Order,json=draft_order,proto3" json:"Order,omitempty"`
	DraftTime            *timestamp.Timestamp `protobuf:"bytes,6,opt,name=draftTime,json=draft_time,proto3" json:"draftTime,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *DraftPlayerResponse) Reset()         { *m = DraftPlayerResponse{} }
func (m *DraftPlayerResponse) String() string { return proto.CompactTextString(m) }
func (*DraftPlayerResponse) ProtoMessage()    {}
func (*DraftPlayerResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_446ed4cf1b75512d, []int{1}
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

func (m *DraftPlayerResponse) GetDraftResultID() int32 {
	if m != nil {
		return m.DraftResultID
	}
	return 0
}

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
	return fileDescriptor_446ed4cf1b75512d, []int{2}
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
	proto.RegisterType((*DraftPlayerRequest)(nil), "fdr.DraftPlayerRequest")
	proto.RegisterType((*DraftPlayerResponse)(nil), "fdr.DraftPlayerResponse")
	proto.RegisterType((*DraftRequest)(nil), "fdr.DraftRequest")
}

func init() {
	proto.RegisterFile("draft_results.proto", fileDescriptor_446ed4cf1b75512d)
}

var fileDescriptor_446ed4cf1b75512d = []byte{
	// 349 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xcc, 0x52, 0x3d, 0x4f, 0xc3, 0x30,
	0x10, 0x25, 0x6d, 0x53, 0xa9, 0x4e, 0x19, 0x70, 0x07, 0xac, 0x88, 0x8f, 0x2a, 0x2c, 0x9d, 0x1c,
	0x54, 0x16, 0x58, 0xab, 0x0e, 0x30, 0x01, 0xa6, 0x2c, 0x2c, 0x51, 0x4a, 0x2e, 0x6d, 0xa4, 0xa6,
	0x0e, 0xb6, 0x33, 0xf0, 0x6b, 0xf8, 0x99, 0xac, 0x28, 0xe7, 0x84, 0xb6, 0x2a, 0x48, 0x8c, 0x4c,
	0xd6, 0x7b, 0xf7, 0xce, 0x7e, 0x77, 0xcf, 0x64, 0x90, 0xa8, 0x38, 0x35, 0x91, 0x02, 0x5d, 0xae,
	0x8c, 0xe6, 0x85, 0x92, 0x46, 0xd2, 0x76, 0x9a, 0x28, 0xff, 0x7c, 0x21, 0xe5, 0x62, 0x05, 0x21,
	0x52, 0xf3, 0x32, 0x0d, 0x4d, 0x96, 0x83, 0x36, 0x71, 0x5e, 0x58, 0x95, 0xdf, 0x2f, 0x56, 0xf1,
	0x3b, 0xa8, 0x1a, 0x79, 0xa5, 0x06, 0xa5, 0x1b, 0x80, 0xb7, 0x5a, 0x10, 0x7c, 0x38, 0x84, 0x4e,
	0x2b, 0xfc, 0x80, 0x7a, 0x01, 0x6f, 0x25, 0x68, 0x43, 0x87, 0xe4, 0x10, 0x59, 0x81, 0x4f, 0xdf,
	0x4d, 0x99, 0x33, 0x74, 0x46, 0xae, 0xe8, 0x59, 0x2b, 0x51, 0x96, 0xd0, 0x0b, 0xd2, 0xd5, 0x10,
	0x6b, 0xb9, 0x66, 0xad, 0xa1, 0x33, 0xf2, 0xc6, 0x1e, 0x4f, 0x13, 0xc5, 0x9f, 0x90, 0x12, 0x75,
	0xa9, 0x12, 0x59, 0x1f, 0xac, 0xbd, 0x25, 0xaa, 0x9f, 0xaa, 0x4b, 0xf4, 0x94, 0x74, 0x2a, 0x7b,
	0xac, 0x83, 0x92, 0x1e, 0x4a, 0x9e, 0x35, 0x28, 0x81, 0x74, 0xf0, 0xe9, 0x90, 0xc1, 0x8e, 0x43,
	0x5d, 0xc8, 0xb5, 0x86, 0x7f, 0x64, 0x91, 0xfa, 0xc4, 0xbd, 0x57, 0x09, 0x28, 0xe6, 0xa2, 0x05,
	0xbb, 0xe0, 0x48, 0x56, 0x14, 0xbd, 0x21, 0x3d, 0x84, 0xb3, 0x2c, 0x07, 0xd6, 0xc5, 0x7e, 0x9f,
	0xdb, 0xf4, 0x78, 0x93, 0x1e, 0x9f, 0x35, 0xe9, 0x09, 0x62, 0x7b, 0xab, 0x38, 0x03, 0x41, 0xfa,
	0xf5, 0x84, 0x36, 0x94, 0xcd, 0x3c, 0xce, 0xef, 0xf3, 0x34, 0x56, 0x5b, 0x3f, 0x5a, 0x1d, 0x3f,
	0x12, 0x6f, 0x6b, 0x99, 0x74, 0xb2, 0x0b, 0x8f, 0x51, 0xbe, 0xff, 0x1f, 0x7c, 0xb6, 0x5f, 0xb0,
	0x31, 0x04, 0x07, 0xe3, 0xdb, 0x6f, 0x9b, 0xf8, 0x4d, 0xe9, 0x35, 0x71, 0x11, 0xd3, 0xa3, 0x4d,
	0xd3, 0x1f, 0xee, 0xb9, 0x74, 0x26, 0x67, 0x2f, 0x27, 0x8b, 0xcc, 0x2c, 0xcb, 0x39, 0x7f, 0x95,
	0x79, 0x68, 0x96, 0x60, 0x96, 0xf1, 0x3a, 0x4c, 0x13, 0x15, 0xd9, 0x4d, 0x75, 0xf1, 0xb8, 0xfa,
	0x0a, 0x00, 0x00, 0xff, 0xff, 0x1a, 0x14, 0x55, 0xb5, 0x18, 0x03, 0x00, 0x00,
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

// DraftResultsClient is the client API for DraftResults service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type DraftResultsClient interface {
	Draft(ctx context.Context, in *DraftRequest, opts ...grpc.CallOption) (DraftResults_DraftClient, error)
}

type draftResultsClient struct {
	cc grpc.ClientConnInterface
}

func NewDraftResultsClient(cc grpc.ClientConnInterface) DraftResultsClient {
	return &draftResultsClient{cc}
}

func (c *draftResultsClient) Draft(ctx context.Context, in *DraftRequest, opts ...grpc.CallOption) (DraftResults_DraftClient, error) {
	stream, err := c.cc.NewStream(ctx, &_DraftResults_serviceDesc.Streams[0], "/fdr.DraftResults/Draft", opts...)
	if err != nil {
		return nil, err
	}
	x := &draftResultsDraftClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type DraftResults_DraftClient interface {
	Recv() (*DraftPlayerResponse, error)
	grpc.ClientStream
}

type draftResultsDraftClient struct {
	grpc.ClientStream
}

func (x *draftResultsDraftClient) Recv() (*DraftPlayerResponse, error) {
	m := new(DraftPlayerResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// DraftResultsServer is the server API for DraftResults service.
type DraftResultsServer interface {
	Draft(*DraftRequest, DraftResults_DraftServer) error
}

// UnimplementedDraftResultsServer can be embedded to have forward compatible implementations.
type UnimplementedDraftResultsServer struct {
}

func (*UnimplementedDraftResultsServer) Draft(req *DraftRequest, srv DraftResults_DraftServer) error {
	return status.Errorf(codes.Unimplemented, "method Draft not implemented")
}

func RegisterDraftResultsServer(s *grpc.Server, srv DraftResultsServer) {
	s.RegisterService(&_DraftResults_serviceDesc, srv)
}

func _DraftResults_Draft_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(DraftRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(DraftResultsServer).Draft(m, &draftResultsDraftServer{stream})
}

type DraftResults_DraftServer interface {
	Send(*DraftPlayerResponse) error
	grpc.ServerStream
}

type draftResultsDraftServer struct {
	grpc.ServerStream
}

func (x *draftResultsDraftServer) Send(m *DraftPlayerResponse) error {
	return x.ServerStream.SendMsg(m)
}

var _DraftResults_serviceDesc = grpc.ServiceDesc{
	ServiceName: "fdr.DraftResults",
	HandlerType: (*DraftResultsServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Draft",
			Handler:       _DraftResults_Draft_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "draft_results.proto",
}
