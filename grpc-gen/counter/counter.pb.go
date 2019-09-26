// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: counter.proto

package counter

import (
	context "context"
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
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
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type IncreaseRequest struct {
	UserID               int64    `protobuf:"varint,1,opt,name=userID,proto3" json:"userID,omitempty"`
	Amount               int64    `protobuf:"varint,2,opt,name=amount,proto3" json:"amount,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IncreaseRequest) Reset()         { *m = IncreaseRequest{} }
func (m *IncreaseRequest) String() string { return proto.CompactTextString(m) }
func (*IncreaseRequest) ProtoMessage()    {}
func (*IncreaseRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_75dcd656fce7132f, []int{0}
}
func (m *IncreaseRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IncreaseRequest.Unmarshal(m, b)
}
func (m *IncreaseRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IncreaseRequest.Marshal(b, m, deterministic)
}
func (m *IncreaseRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IncreaseRequest.Merge(m, src)
}
func (m *IncreaseRequest) XXX_Size() int {
	return xxx_messageInfo_IncreaseRequest.Size(m)
}
func (m *IncreaseRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_IncreaseRequest.DiscardUnknown(m)
}

var xxx_messageInfo_IncreaseRequest proto.InternalMessageInfo

func (m *IncreaseRequest) GetUserID() int64 {
	if m != nil {
		return m.UserID
	}
	return 0
}

func (m *IncreaseRequest) GetAmount() int64 {
	if m != nil {
		return m.Amount
	}
	return 0
}

type IncreaseResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IncreaseResponse) Reset()         { *m = IncreaseResponse{} }
func (m *IncreaseResponse) String() string { return proto.CompactTextString(m) }
func (*IncreaseResponse) ProtoMessage()    {}
func (*IncreaseResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_75dcd656fce7132f, []int{1}
}
func (m *IncreaseResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IncreaseResponse.Unmarshal(m, b)
}
func (m *IncreaseResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IncreaseResponse.Marshal(b, m, deterministic)
}
func (m *IncreaseResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IncreaseResponse.Merge(m, src)
}
func (m *IncreaseResponse) XXX_Size() int {
	return xxx_messageInfo_IncreaseResponse.Size(m)
}
func (m *IncreaseResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_IncreaseResponse.DiscardUnknown(m)
}

var xxx_messageInfo_IncreaseResponse proto.InternalMessageInfo

type DecreaseRequest struct {
	UserID               int64    `protobuf:"varint,1,opt,name=userID,proto3" json:"userID,omitempty"`
	Amount               int64    `protobuf:"varint,2,opt,name=amount,proto3" json:"amount,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DecreaseRequest) Reset()         { *m = DecreaseRequest{} }
func (m *DecreaseRequest) String() string { return proto.CompactTextString(m) }
func (*DecreaseRequest) ProtoMessage()    {}
func (*DecreaseRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_75dcd656fce7132f, []int{2}
}
func (m *DecreaseRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DecreaseRequest.Unmarshal(m, b)
}
func (m *DecreaseRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DecreaseRequest.Marshal(b, m, deterministic)
}
func (m *DecreaseRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DecreaseRequest.Merge(m, src)
}
func (m *DecreaseRequest) XXX_Size() int {
	return xxx_messageInfo_DecreaseRequest.Size(m)
}
func (m *DecreaseRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DecreaseRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DecreaseRequest proto.InternalMessageInfo

func (m *DecreaseRequest) GetUserID() int64 {
	if m != nil {
		return m.UserID
	}
	return 0
}

func (m *DecreaseRequest) GetAmount() int64 {
	if m != nil {
		return m.Amount
	}
	return 0
}

type DecreaseResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DecreaseResponse) Reset()         { *m = DecreaseResponse{} }
func (m *DecreaseResponse) String() string { return proto.CompactTextString(m) }
func (*DecreaseResponse) ProtoMessage()    {}
func (*DecreaseResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_75dcd656fce7132f, []int{3}
}
func (m *DecreaseResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DecreaseResponse.Unmarshal(m, b)
}
func (m *DecreaseResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DecreaseResponse.Marshal(b, m, deterministic)
}
func (m *DecreaseResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DecreaseResponse.Merge(m, src)
}
func (m *DecreaseResponse) XXX_Size() int {
	return xxx_messageInfo_DecreaseResponse.Size(m)
}
func (m *DecreaseResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DecreaseResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DecreaseResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*IncreaseRequest)(nil), "counter.IncreaseRequest")
	proto.RegisterType((*IncreaseResponse)(nil), "counter.IncreaseResponse")
	proto.RegisterType((*DecreaseRequest)(nil), "counter.DecreaseRequest")
	proto.RegisterType((*DecreaseResponse)(nil), "counter.DecreaseResponse")
}

func init() { proto.RegisterFile("counter.proto", fileDescriptor_75dcd656fce7132f) }

var fileDescriptor_75dcd656fce7132f = []byte{
	// 176 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4d, 0xce, 0x2f, 0xcd,
	0x2b, 0x49, 0x2d, 0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x87, 0x72, 0x95, 0x1c, 0xb9,
	0xf8, 0x3d, 0xf3, 0x92, 0x8b, 0x52, 0x13, 0x8b, 0x53, 0x83, 0x52, 0x0b, 0x4b, 0x53, 0x8b, 0x4b,
	0x84, 0xc4, 0xb8, 0xd8, 0x4a, 0x8b, 0x53, 0x8b, 0x3c, 0x5d, 0x24, 0x18, 0x15, 0x18, 0x35, 0x98,
	0x83, 0xa0, 0x3c, 0x90, 0x78, 0x62, 0x2e, 0x48, 0x9b, 0x04, 0x13, 0x44, 0x1c, 0xc2, 0x53, 0x12,
	0xe2, 0x12, 0x40, 0x18, 0x51, 0x5c, 0x90, 0x9f, 0x57, 0x9c, 0x0a, 0x32, 0xd6, 0x25, 0x95, 0x62,
	0x63, 0x11, 0x46, 0x40, 0x8c, 0x35, 0x5a, 0xc8, 0xc8, 0x25, 0xec, 0x0c, 0x71, 0x79, 0x40, 0x7e,
	0x66, 0x5e, 0x49, 0x70, 0x6a, 0x51, 0x59, 0x66, 0x72, 0xaa, 0x90, 0x0b, 0x17, 0x2f, 0xcc, 0x09,
	0x60, 0x71, 0x21, 0x09, 0x3d, 0x98, 0x7f, 0xd1, 0x7c, 0x27, 0x25, 0x89, 0x45, 0x06, 0x62, 0x3a,
	0xc8, 0x14, 0x98, 0x8d, 0xe8, 0xa6, 0xa0, 0x79, 0x06, 0xc9, 0x14, 0x74, 0x37, 0x26, 0xb1, 0x81,
	0x43, 0xd8, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0x2a, 0xa0, 0x35, 0x7a, 0x72, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// CounterPointServiceClient is the client API for CounterPointService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CounterPointServiceClient interface {
	IncreasePoint(ctx context.Context, in *IncreaseRequest, opts ...grpc.CallOption) (*IncreaseResponse, error)
	DecreasePoint(ctx context.Context, in *DecreaseRequest, opts ...grpc.CallOption) (*DecreaseResponse, error)
}

type counterPointServiceClient struct {
	cc *grpc.ClientConn
}

func NewCounterPointServiceClient(cc *grpc.ClientConn) CounterPointServiceClient {
	return &counterPointServiceClient{cc}
}

func (c *counterPointServiceClient) IncreasePoint(ctx context.Context, in *IncreaseRequest, opts ...grpc.CallOption) (*IncreaseResponse, error) {
	out := new(IncreaseResponse)
	err := c.cc.Invoke(ctx, "/counter.CounterPointService/IncreasePoint", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *counterPointServiceClient) DecreasePoint(ctx context.Context, in *DecreaseRequest, opts ...grpc.CallOption) (*DecreaseResponse, error) {
	out := new(DecreaseResponse)
	err := c.cc.Invoke(ctx, "/counter.CounterPointService/DecreasePoint", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CounterPointServiceServer is the server API for CounterPointService service.
type CounterPointServiceServer interface {
	IncreasePoint(context.Context, *IncreaseRequest) (*IncreaseResponse, error)
	DecreasePoint(context.Context, *DecreaseRequest) (*DecreaseResponse, error)
}

// UnimplementedCounterPointServiceServer can be embedded to have forward compatible implementations.
type UnimplementedCounterPointServiceServer struct {
}

func (*UnimplementedCounterPointServiceServer) IncreasePoint(ctx context.Context, req *IncreaseRequest) (*IncreaseResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IncreasePoint not implemented")
}
func (*UnimplementedCounterPointServiceServer) DecreasePoint(ctx context.Context, req *DecreaseRequest) (*DecreaseResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DecreasePoint not implemented")
}

func RegisterCounterPointServiceServer(s *grpc.Server, srv CounterPointServiceServer) {
	s.RegisterService(&_CounterPointService_serviceDesc, srv)
}

func _CounterPointService_IncreasePoint_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IncreaseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CounterPointServiceServer).IncreasePoint(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/counter.CounterPointService/IncreasePoint",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CounterPointServiceServer).IncreasePoint(ctx, req.(*IncreaseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CounterPointService_DecreasePoint_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DecreaseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CounterPointServiceServer).DecreasePoint(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/counter.CounterPointService/DecreasePoint",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CounterPointServiceServer).DecreasePoint(ctx, req.(*DecreaseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _CounterPointService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "counter.CounterPointService",
	HandlerType: (*CounterPointServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "IncreasePoint",
			Handler:    _CounterPointService_IncreasePoint_Handler,
		},
		{
			MethodName: "DecreasePoint",
			Handler:    _CounterPointService_DecreasePoint_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "counter.proto",
}
