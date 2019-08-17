// Code generated by protoc-gen-go. DO NOT EDIT.
// source: core.proto

package core

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

//  protoc -I ~/go/src/awesomeProject2/proto ~/go/src/awesomeProject2/proto/test.proto --go_out=plugins=grpc:proto
type CheckAuthResponse struct {
	IsAuth               bool     `protobuf:"varint,1,opt,name=isAuth,proto3" json:"isAuth,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CheckAuthResponse) Reset()         { *m = CheckAuthResponse{} }
func (m *CheckAuthResponse) String() string { return proto.CompactTextString(m) }
func (*CheckAuthResponse) ProtoMessage()    {}
func (*CheckAuthResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_f7e43720d1edc0fe, []int{0}
}

func (m *CheckAuthResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CheckAuthResponse.Unmarshal(m, b)
}
func (m *CheckAuthResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CheckAuthResponse.Marshal(b, m, deterministic)
}
func (m *CheckAuthResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CheckAuthResponse.Merge(m, src)
}
func (m *CheckAuthResponse) XXX_Size() int {
	return xxx_messageInfo_CheckAuthResponse.Size(m)
}
func (m *CheckAuthResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CheckAuthResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CheckAuthResponse proto.InternalMessageInfo

func (m *CheckAuthResponse) GetIsAuth() bool {
	if m != nil {
		return m.IsAuth
	}
	return false
}

type CheckAuthRequest struct {
	Session              string   `protobuf:"bytes,1,opt,name=session,proto3" json:"session,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CheckAuthRequest) Reset()         { *m = CheckAuthRequest{} }
func (m *CheckAuthRequest) String() string { return proto.CompactTextString(m) }
func (*CheckAuthRequest) ProtoMessage()    {}
func (*CheckAuthRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_f7e43720d1edc0fe, []int{1}
}

func (m *CheckAuthRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CheckAuthRequest.Unmarshal(m, b)
}
func (m *CheckAuthRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CheckAuthRequest.Marshal(b, m, deterministic)
}
func (m *CheckAuthRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CheckAuthRequest.Merge(m, src)
}
func (m *CheckAuthRequest) XXX_Size() int {
	return xxx_messageInfo_CheckAuthRequest.Size(m)
}
func (m *CheckAuthRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CheckAuthRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CheckAuthRequest proto.InternalMessageInfo

func (m *CheckAuthRequest) GetSession() string {
	if m != nil {
		return m.Session
	}
	return ""
}

func init() {
	proto.RegisterType((*CheckAuthResponse)(nil), "CheckAuthResponse")
	proto.RegisterType((*CheckAuthRequest)(nil), "CheckAuthRequest")
}

func init() { proto.RegisterFile("core.proto", fileDescriptor_f7e43720d1edc0fe) }

var fileDescriptor_f7e43720d1edc0fe = []byte{
	// 134 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4a, 0xce, 0x2f, 0x4a,
	0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x57, 0xd2, 0xe6, 0x12, 0x74, 0xce, 0x48, 0x4d, 0xce, 0x76,
	0x2c, 0x2d, 0xc9, 0x08, 0x4a, 0x2d, 0x2e, 0xc8, 0xcf, 0x2b, 0x4e, 0x15, 0x12, 0xe3, 0x62, 0xcb,
	0x2c, 0x06, 0x89, 0x48, 0x30, 0x2a, 0x30, 0x6a, 0x70, 0x04, 0x41, 0x79, 0x4a, 0x3a, 0x5c, 0x02,
	0x48, 0x8a, 0x0b, 0x4b, 0x53, 0x8b, 0x4b, 0x84, 0x24, 0xb8, 0xd8, 0x8b, 0x53, 0x8b, 0x8b, 0x33,
	0xf3, 0xf3, 0xc0, 0x8a, 0x39, 0x83, 0x60, 0x5c, 0x23, 0x37, 0x24, 0xd5, 0xc1, 0xa9, 0x45, 0x65,
	0x99, 0xc9, 0xa9, 0x42, 0x46, 0x5c, 0x9c, 0x70, 0x31, 0x21, 0x41, 0x3d, 0x74, 0xd3, 0xa4, 0x84,
	0xf4, 0x30, 0x5c, 0x93, 0xc4, 0x06, 0x76, 0xa9, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0x5b, 0x32,
	0x2e, 0xc8, 0xb7, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// CheckAuthServiceClient is the client API for CheckAuthService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CheckAuthServiceClient interface {
	CheckAuth(ctx context.Context, in *CheckAuthRequest, opts ...grpc.CallOption) (*CheckAuthResponse, error)
}

type checkAuthServiceClient struct {
	cc *grpc.ClientConn
}

func NewCheckAuthServiceClient(cc *grpc.ClientConn) CheckAuthServiceClient {
	return &checkAuthServiceClient{cc}
}

func (c *checkAuthServiceClient) CheckAuth(ctx context.Context, in *CheckAuthRequest, opts ...grpc.CallOption) (*CheckAuthResponse, error) {
	out := new(CheckAuthResponse)
	err := c.cc.Invoke(ctx, "/CheckAuthService/CheckAuth", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CheckAuthServiceServer is the server API for CheckAuthService service.
type CheckAuthServiceServer interface {
	CheckAuth(context.Context, *CheckAuthRequest) (*CheckAuthResponse, error)
}

// UnimplementedCheckAuthServiceServer can be embedded to have forward compatible implementations.
type UnimplementedCheckAuthServiceServer struct {
}

func (*UnimplementedCheckAuthServiceServer) CheckAuth(ctx context.Context, req *CheckAuthRequest) (*CheckAuthResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckAuth not implemented")
}

func RegisterCheckAuthServiceServer(s *grpc.Server, srv CheckAuthServiceServer) {
	s.RegisterService(&_CheckAuthService_serviceDesc, srv)
}

func _CheckAuthService_CheckAuth_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckAuthRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CheckAuthServiceServer).CheckAuth(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/CheckAuthService/CheckAuth",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CheckAuthServiceServer).CheckAuth(ctx, req.(*CheckAuthRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _CheckAuthService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "CheckAuthService",
	HandlerType: (*CheckAuthServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CheckAuth",
			Handler:    _CheckAuthService_CheckAuth_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "core.proto",
}
