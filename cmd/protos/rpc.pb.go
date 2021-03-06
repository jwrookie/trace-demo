// Code generated by protoc-gen-go. DO NOT EDIT.
// source: rpc.proto

package protos

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

type PingReq struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PingReq) Reset()         { *m = PingReq{} }
func (m *PingReq) String() string { return proto.CompactTextString(m) }
func (*PingReq) ProtoMessage()    {}
func (*PingReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_77a6da22d6a3feb1, []int{0}
}

func (m *PingReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PingReq.Unmarshal(m, b)
}
func (m *PingReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PingReq.Marshal(b, m, deterministic)
}
func (m *PingReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PingReq.Merge(m, src)
}
func (m *PingReq) XXX_Size() int {
	return xxx_messageInfo_PingReq.Size(m)
}
func (m *PingReq) XXX_DiscardUnknown() {
	xxx_messageInfo_PingReq.DiscardUnknown(m)
}

var xxx_messageInfo_PingReq proto.InternalMessageInfo

type PingRes struct {
	Res                  string   `protobuf:"bytes,1,opt,name=res,proto3" json:"res,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PingRes) Reset()         { *m = PingRes{} }
func (m *PingRes) String() string { return proto.CompactTextString(m) }
func (*PingRes) ProtoMessage()    {}
func (*PingRes) Descriptor() ([]byte, []int) {
	return fileDescriptor_77a6da22d6a3feb1, []int{1}
}

func (m *PingRes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PingRes.Unmarshal(m, b)
}
func (m *PingRes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PingRes.Marshal(b, m, deterministic)
}
func (m *PingRes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PingRes.Merge(m, src)
}
func (m *PingRes) XXX_Size() int {
	return xxx_messageInfo_PingRes.Size(m)
}
func (m *PingRes) XXX_DiscardUnknown() {
	xxx_messageInfo_PingRes.DiscardUnknown(m)
}

var xxx_messageInfo_PingRes proto.InternalMessageInfo

func (m *PingRes) GetRes() string {
	if m != nil {
		return m.Res
	}
	return ""
}

func init() {
	proto.RegisterType((*PingReq)(nil), "protos.pingReq")
	proto.RegisterType((*PingRes)(nil), "protos.pingRes")
}

func init() { proto.RegisterFile("rpc.proto", fileDescriptor_77a6da22d6a3feb1) }

var fileDescriptor_77a6da22d6a3feb1 = []byte{
	// 108 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2c, 0x2a, 0x48, 0xd6,
	0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x03, 0x53, 0xc5, 0x4a, 0x9c, 0x5c, 0xec, 0x05, 0x99,
	0x79, 0xe9, 0x41, 0xa9, 0x85, 0x4a, 0xd2, 0x30, 0x66, 0xb1, 0x90, 0x00, 0x17, 0x73, 0x51, 0x6a,
	0xb1, 0x04, 0xa3, 0x02, 0xa3, 0x06, 0x67, 0x10, 0x88, 0x69, 0x64, 0x0e, 0xd6, 0x1c, 0x9c, 0x5a,
	0x54, 0x96, 0x5a, 0x24, 0xa4, 0xc5, 0xc5, 0x12, 0x90, 0x99, 0x97, 0x2e, 0xc4, 0x0f, 0x31, 0xac,
	0x58, 0x0f, 0x6a, 0x84, 0x14, 0x9a, 0x40, 0xb1, 0x12, 0x43, 0x12, 0xc4, 0x22, 0x63, 0x40, 0x00,
	0x00, 0x00, 0xff, 0xff, 0xd5, 0x71, 0x46, 0x64, 0x7c, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// RpcServerClient is the client API for RpcServer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type RpcServerClient interface {
	Ping(ctx context.Context, in *PingReq, opts ...grpc.CallOption) (*PingRes, error)
}

type rpcServerClient struct {
	cc grpc.ClientConnInterface
}

func NewRpcServerClient(cc grpc.ClientConnInterface) RpcServerClient {
	return &rpcServerClient{cc}
}

func (c *rpcServerClient) Ping(ctx context.Context, in *PingReq, opts ...grpc.CallOption) (*PingRes, error) {
	out := new(PingRes)
	err := c.cc.Invoke(ctx, "/protos.rpcServer/Ping", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RpcServerServer is the server API for RpcServer service.
type RpcServerServer interface {
	Ping(context.Context, *PingReq) (*PingRes, error)
}

// UnimplementedRpcServerServer can be embedded to have forward compatible implementations.
type UnimplementedRpcServerServer struct {
}

func (*UnimplementedRpcServerServer) Ping(ctx context.Context, req *PingReq) (*PingRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ping not implemented")
}

func RegisterRpcServerServer(s *grpc.Server, srv RpcServerServer) {
	s.RegisterService(&_RpcServer_serviceDesc, srv)
}

func _RpcServer_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PingReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RpcServerServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.rpcServer/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RpcServerServer).Ping(ctx, req.(*PingReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _RpcServer_serviceDesc = grpc.ServiceDesc{
	ServiceName: "protos.rpcServer",
	HandlerType: (*RpcServerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ping",
			Handler:    _RpcServer_Ping_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "rpc.proto",
}
