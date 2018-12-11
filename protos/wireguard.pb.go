// Code generated by protoc-gen-go. DO NOT EDIT.
// source: wireguard.proto

package wireguard

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

// The request message containing the user's name.
type ShowRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ShowRequest) Reset()         { *m = ShowRequest{} }
func (m *ShowRequest) String() string { return proto.CompactTextString(m) }
func (*ShowRequest) ProtoMessage()    {}
func (*ShowRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_wireguard_9673e8d1941d2075, []int{0}
}
func (m *ShowRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ShowRequest.Unmarshal(m, b)
}
func (m *ShowRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ShowRequest.Marshal(b, m, deterministic)
}
func (dst *ShowRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ShowRequest.Merge(dst, src)
}
func (m *ShowRequest) XXX_Size() int {
	return xxx_messageInfo_ShowRequest.Size(m)
}
func (m *ShowRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ShowRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ShowRequest proto.InternalMessageInfo

func (m *ShowRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

// The response message containing the greetings
type ShowReply struct {
	Message              string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ShowReply) Reset()         { *m = ShowReply{} }
func (m *ShowReply) String() string { return proto.CompactTextString(m) }
func (*ShowReply) ProtoMessage()    {}
func (*ShowReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_wireguard_9673e8d1941d2075, []int{1}
}
func (m *ShowReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ShowReply.Unmarshal(m, b)
}
func (m *ShowReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ShowReply.Marshal(b, m, deterministic)
}
func (dst *ShowReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ShowReply.Merge(dst, src)
}
func (m *ShowReply) XXX_Size() int {
	return xxx_messageInfo_ShowReply.Size(m)
}
func (m *ShowReply) XXX_DiscardUnknown() {
	xxx_messageInfo_ShowReply.DiscardUnknown(m)
}

var xxx_messageInfo_ShowReply proto.InternalMessageInfo

func (m *ShowReply) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*ShowRequest)(nil), "wireguard.ShowRequest")
	proto.RegisterType((*ShowReply)(nil), "wireguard.ShowReply")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// WireGuardClient is the client API for WireGuard service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type WireGuardClient interface {
	// Sends a greeting
	Show(ctx context.Context, in *ShowRequest, opts ...grpc.CallOption) (*ShowReply, error)
}

type wireGuardClient struct {
	cc *grpc.ClientConn
}

func NewWireGuardClient(cc *grpc.ClientConn) WireGuardClient {
	return &wireGuardClient{cc}
}

func (c *wireGuardClient) Show(ctx context.Context, in *ShowRequest, opts ...grpc.CallOption) (*ShowReply, error) {
	out := new(ShowReply)
	err := c.cc.Invoke(ctx, "/wireguard.WireGuard/Show", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// WireGuardServer is the server API for WireGuard service.
type WireGuardServer interface {
	// Sends a greeting
	Show(context.Context, *ShowRequest) (*ShowReply, error)
}

func RegisterWireGuardServer(s *grpc.Server, srv WireGuardServer) {
	s.RegisterService(&_WireGuard_serviceDesc, srv)
}

func _WireGuard_Show_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ShowRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WireGuardServer).Show(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/wireguard.WireGuard/Show",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WireGuardServer).Show(ctx, req.(*ShowRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _WireGuard_serviceDesc = grpc.ServiceDesc{
	ServiceName: "wireguard.WireGuard",
	HandlerType: (*WireGuardServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Show",
			Handler:    _WireGuard_Show_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "wireguard.proto",
}

func init() { proto.RegisterFile("wireguard.proto", fileDescriptor_wireguard_9673e8d1941d2075) }

var fileDescriptor_wireguard_9673e8d1941d2075 = []byte{
	// 173 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2f, 0xcf, 0x2c, 0x4a,
	0x4d, 0x2f, 0x4d, 0x2c, 0x4a, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x84, 0x0b, 0x28,
	0x29, 0x72, 0x71, 0x07, 0x67, 0xe4, 0x97, 0x07, 0xa5, 0x16, 0x96, 0xa6, 0x16, 0x97, 0x08, 0x09,
	0x71, 0xb1, 0xe4, 0x25, 0xe6, 0xa6, 0x4a, 0x30, 0x2a, 0x30, 0x6a, 0x70, 0x06, 0x81, 0xd9, 0x4a,
	0xaa, 0x5c, 0x9c, 0x10, 0x25, 0x05, 0x39, 0x95, 0x42, 0x12, 0x5c, 0xec, 0xb9, 0xa9, 0xc5, 0xc5,
	0x89, 0xe9, 0x30, 0x35, 0x30, 0xae, 0x91, 0x33, 0x17, 0x67, 0x78, 0x66, 0x51, 0xaa, 0x3b, 0xc8,
	0x58, 0x21, 0x33, 0x2e, 0x16, 0x90, 0x1e, 0x21, 0x31, 0x3d, 0x84, 0xdd, 0x48, 0xf6, 0x48, 0x89,
	0x60, 0x88, 0x17, 0xe4, 0x54, 0x2a, 0x31, 0x38, 0xe9, 0x73, 0x49, 0x25, 0xe7, 0xe7, 0xea, 0x95,
	0xe6, 0x25, 0x67, 0x24, 0x16, 0x95, 0xa4, 0xa6, 0x14, 0x67, 0x57, 0x22, 0x54, 0x3a, 0xf1, 0x83,
	0x2c, 0x28, 0x02, 0xdb, 0x10, 0x00, 0xf2, 0x48, 0x00, 0x63, 0x12, 0x1b, 0xd8, 0x47, 0xc6, 0x80,
	0x00, 0x00, 0x00, 0xff, 0xff, 0xf4, 0x6b, 0x96, 0x62, 0xe4, 0x00, 0x00, 0x00,
}
