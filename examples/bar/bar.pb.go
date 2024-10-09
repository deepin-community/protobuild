// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/containerd/protobuild/examples/bar/bar.proto

package bar

import (
	context "context"
	fmt "fmt"
	internal "github.com/containerd/protobuild/internal"
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

type DoRequest struct {
	Timestamp            *timestamp.Timestamp `protobuf:"bytes,3,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *DoRequest) Reset()         { *m = DoRequest{} }
func (m *DoRequest) String() string { return proto.CompactTextString(m) }
func (*DoRequest) ProtoMessage()    {}
func (*DoRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_bdef534bf6116654, []int{0}
}

func (m *DoRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DoRequest.Unmarshal(m, b)
}
func (m *DoRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DoRequest.Marshal(b, m, deterministic)
}
func (m *DoRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DoRequest.Merge(m, src)
}
func (m *DoRequest) XXX_Size() int {
	return xxx_messageInfo_DoRequest.Size(m)
}
func (m *DoRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DoRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DoRequest proto.InternalMessageInfo

func (m *DoRequest) GetTimestamp() *timestamp.Timestamp {
	if m != nil {
		return m.Timestamp
	}
	return nil
}

func init() {
	proto.RegisterType((*DoRequest)(nil), "protobuild.example.bar.DoRequest")
}

func init() {
	proto.RegisterFile("github.com/containerd/protobuild/examples/bar/bar.proto", fileDescriptor_bdef534bf6116654)
}

var fileDescriptor_bdef534bf6116654 = []byte{
	// 205 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x32, 0x4f, 0xcf, 0x2c, 0xc9,
	0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0x4f, 0xce, 0xcf, 0x2b, 0x49, 0xcc, 0xcc, 0x4b, 0x2d,
	0x4a, 0xd1, 0x2f, 0x28, 0xca, 0x2f, 0xc9, 0x4f, 0x2a, 0xcd, 0xcc, 0x49, 0xd1, 0x4f, 0xad, 0x48,
	0xcc, 0x2d, 0xc8, 0x49, 0x2d, 0xd6, 0x4f, 0x4a, 0x2c, 0x02, 0x61, 0x3d, 0xb0, 0x9c, 0x90, 0x18,
	0x42, 0x89, 0x1e, 0x54, 0x89, 0x5e, 0x52, 0x62, 0x91, 0x94, 0x74, 0x7a, 0x7e, 0x7e, 0x7a, 0x4e,
	0x2a, 0xcc, 0x84, 0x34, 0xfd, 0xd4, 0xdc, 0x82, 0x92, 0x4a, 0x88, 0x26, 0x29, 0x79, 0x74, 0xc9,
	0x92, 0xcc, 0xdc, 0xd4, 0xe2, 0x92, 0xc4, 0xdc, 0x02, 0x88, 0x02, 0x25, 0x57, 0x2e, 0x4e, 0x97,
	0xfc, 0xa0, 0xd4, 0xc2, 0xd2, 0xd4, 0xe2, 0x12, 0x21, 0x0b, 0x2e, 0x4e, 0xb8, 0xbc, 0x04, 0xb3,
	0x02, 0xa3, 0x06, 0xb7, 0x91, 0x94, 0x1e, 0xc4, 0x04, 0x3d, 0x98, 0x09, 0x7a, 0x21, 0x30, 0x15,
	0x41, 0x08, 0xc5, 0x46, 0x6e, 0x5c, 0xcc, 0x4e, 0x89, 0x45, 0x42, 0xf6, 0x5c, 0x4c, 0x2e, 0xf9,
	0x42, 0x8a, 0x7a, 0xd8, 0x9d, 0xaa, 0x07, 0xb7, 0x49, 0x4a, 0x0c, 0xc3, 0x58, 0x57, 0x90, 0xab,
	0x9d, 0x4c, 0xa2, 0x8c, 0x48, 0x0b, 0x1f, 0xeb, 0xa4, 0xc4, 0xa2, 0x24, 0x36, 0xb0, 0xac, 0x31,
	0x20, 0x00, 0x00, 0xff, 0xff, 0x2a, 0xa0, 0x00, 0x73, 0x5c, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// BarClient is the client API for Bar service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type BarClient interface {
	Do(ctx context.Context, in *DoRequest, opts ...grpc.CallOption) (*internal.Empty, error)
}

type barClient struct {
	cc grpc.ClientConnInterface
}

func NewBarClient(cc grpc.ClientConnInterface) BarClient {
	return &barClient{cc}
}

func (c *barClient) Do(ctx context.Context, in *DoRequest, opts ...grpc.CallOption) (*internal.Empty, error) {
	out := new(internal.Empty)
	err := c.cc.Invoke(ctx, "/protobuild.example.bar.Bar/Do", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BarServer is the server API for Bar service.
type BarServer interface {
	Do(context.Context, *DoRequest) (*internal.Empty, error)
}

// UnimplementedBarServer can be embedded to have forward compatible implementations.
type UnimplementedBarServer struct {
}

func (*UnimplementedBarServer) Do(ctx context.Context, req *DoRequest) (*internal.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Do not implemented")
}

func RegisterBarServer(s *grpc.Server, srv BarServer) {
	s.RegisterService(&_Bar_serviceDesc, srv)
}

func _Bar_Do_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BarServer).Do(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protobuild.example.bar.Bar/Do",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BarServer).Do(ctx, req.(*DoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Bar_serviceDesc = grpc.ServiceDesc{
	ServiceName: "protobuild.example.bar.Bar",
	HandlerType: (*BarServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Do",
			Handler:    _Bar_Do_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "github.com/containerd/protobuild/examples/bar/bar.proto",
}