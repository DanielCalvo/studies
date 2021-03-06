// Code generated by protoc-gen-go. DO NOT EDIT.
// source: Section8_grpcbidirectionalstreaming/greet/greetpb/greet.proto

package greetpb

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

type Greeting struct {
	FirstName            string   `protobuf:"bytes,1,opt,name=first_name,json=firstName,proto3" json:"first_name,omitempty"`
	LastName             string   `protobuf:"bytes,2,opt,name=last_name,json=lastName,proto3" json:"last_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Greeting) Reset()         { *m = Greeting{} }
func (m *Greeting) String() string { return proto.CompactTextString(m) }
func (*Greeting) ProtoMessage()    {}
func (*Greeting) Descriptor() ([]byte, []int) {
	return fileDescriptor_08e0af01c6905008, []int{0}
}

func (m *Greeting) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Greeting.Unmarshal(m, b)
}
func (m *Greeting) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Greeting.Marshal(b, m, deterministic)
}
func (m *Greeting) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Greeting.Merge(m, src)
}
func (m *Greeting) XXX_Size() int {
	return xxx_messageInfo_Greeting.Size(m)
}
func (m *Greeting) XXX_DiscardUnknown() {
	xxx_messageInfo_Greeting.DiscardUnknown(m)
}

var xxx_messageInfo_Greeting proto.InternalMessageInfo

func (m *Greeting) GetFirstName() string {
	if m != nil {
		return m.FirstName
	}
	return ""
}

func (m *Greeting) GetLastName() string {
	if m != nil {
		return m.LastName
	}
	return ""
}

type GreetEveryoneRequest struct {
	Greeting             *Greeting `protobuf:"bytes,1,opt,name=greeting,proto3" json:"greeting,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *GreetEveryoneRequest) Reset()         { *m = GreetEveryoneRequest{} }
func (m *GreetEveryoneRequest) String() string { return proto.CompactTextString(m) }
func (*GreetEveryoneRequest) ProtoMessage()    {}
func (*GreetEveryoneRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_08e0af01c6905008, []int{1}
}

func (m *GreetEveryoneRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GreetEveryoneRequest.Unmarshal(m, b)
}
func (m *GreetEveryoneRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GreetEveryoneRequest.Marshal(b, m, deterministic)
}
func (m *GreetEveryoneRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GreetEveryoneRequest.Merge(m, src)
}
func (m *GreetEveryoneRequest) XXX_Size() int {
	return xxx_messageInfo_GreetEveryoneRequest.Size(m)
}
func (m *GreetEveryoneRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GreetEveryoneRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GreetEveryoneRequest proto.InternalMessageInfo

func (m *GreetEveryoneRequest) GetGreeting() *Greeting {
	if m != nil {
		return m.Greeting
	}
	return nil
}

type GreetEveryoneResponse struct {
	Result               string   `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GreetEveryoneResponse) Reset()         { *m = GreetEveryoneResponse{} }
func (m *GreetEveryoneResponse) String() string { return proto.CompactTextString(m) }
func (*GreetEveryoneResponse) ProtoMessage()    {}
func (*GreetEveryoneResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_08e0af01c6905008, []int{2}
}

func (m *GreetEveryoneResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GreetEveryoneResponse.Unmarshal(m, b)
}
func (m *GreetEveryoneResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GreetEveryoneResponse.Marshal(b, m, deterministic)
}
func (m *GreetEveryoneResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GreetEveryoneResponse.Merge(m, src)
}
func (m *GreetEveryoneResponse) XXX_Size() int {
	return xxx_messageInfo_GreetEveryoneResponse.Size(m)
}
func (m *GreetEveryoneResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GreetEveryoneResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GreetEveryoneResponse proto.InternalMessageInfo

func (m *GreetEveryoneResponse) GetResult() string {
	if m != nil {
		return m.Result
	}
	return ""
}

func init() {
	proto.RegisterType((*Greeting)(nil), "greet.Greeting")
	proto.RegisterType((*GreetEveryoneRequest)(nil), "greet.GreetEveryoneRequest")
	proto.RegisterType((*GreetEveryoneResponse)(nil), "greet.GreetEveryoneResponse")
}

func init() {
	proto.RegisterFile("Section8_grpcbidirectionalstreaming/greet/greetpb/greet.proto", fileDescriptor_08e0af01c6905008)
}

var fileDescriptor_08e0af01c6905008 = []byte{
	// 247 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x90, 0x41, 0x4b, 0x03, 0x31,
	0x10, 0x85, 0x5d, 0xc1, 0xba, 0x19, 0x15, 0x21, 0xa8, 0x88, 0x55, 0x90, 0x9c, 0x0a, 0x42, 0x2b,
	0xf5, 0xe2, 0xc5, 0x8b, 0xa2, 0xde, 0x44, 0xb6, 0x37, 0x2f, 0x35, 0xbb, 0x8e, 0x21, 0xb0, 0x9b,
	0xc4, 0x49, 0x5a, 0xf0, 0xdf, 0xcb, 0x26, 0xa9, 0x68, 0xe9, 0x65, 0x37, 0xf3, 0x5e, 0xf2, 0xbd,
	0xc7, 0xc0, 0xdd, 0x0c, 0x9b, 0xa0, 0xad, 0xb9, 0x9d, 0x2b, 0x72, 0x4d, 0xad, 0x3f, 0x34, 0x25,
	0x41, 0xb6, 0x3e, 0x10, 0xca, 0x4e, 0x1b, 0x35, 0x51, 0x84, 0x18, 0xd2, 0xd7, 0xd5, 0xe9, 0x3f,
	0x76, 0x64, 0x83, 0xe5, 0x3b, 0x71, 0x10, 0x4f, 0x50, 0x3e, 0xf7, 0x07, 0x6d, 0x14, 0xbf, 0x00,
	0xf8, 0xd4, 0xe4, 0xc3, 0xdc, 0xc8, 0x0e, 0x4f, 0x8b, 0xcb, 0x62, 0xc4, 0x2a, 0x16, 0x95, 0x17,
	0xd9, 0x21, 0x1f, 0x02, 0x6b, 0xe5, 0xca, 0xdd, 0x8e, 0x6e, 0xd9, 0x0b, 0xbd, 0x29, 0x1e, 0xe0,
	0x28, 0x72, 0x1e, 0x97, 0x48, 0xdf, 0xd6, 0x60, 0x85, 0x5f, 0x0b, 0xf4, 0x81, 0x5f, 0x41, 0xa9,
	0x32, 0x3f, 0x12, 0xf7, 0xa6, 0x87, 0xe3, 0x54, 0x63, 0x15, 0x5b, 0xfd, 0x5e, 0x10, 0x13, 0x38,
	0x5e, 0x83, 0x78, 0x67, 0x8d, 0x47, 0x7e, 0x02, 0x03, 0x42, 0xbf, 0x68, 0x43, 0x6e, 0x95, 0xa7,
	0xe9, 0x3b, 0xec, 0xc7, 0x07, 0x33, 0xa4, 0xa5, 0x6e, 0x90, 0xbf, 0xc2, 0xc1, 0x3f, 0x00, 0x1f,
	0xfe, 0x0d, 0x5b, 0xeb, 0x76, 0x76, 0xbe, 0xd9, 0x4c, 0x99, 0x62, 0x6b, 0x54, 0x5c, 0x17, 0xf7,
	0xec, 0x6d, 0x37, 0x6f, 0xaf, 0x1e, 0xc4, 0xc5, 0xdd, 0xfc, 0x04, 0x00, 0x00, 0xff, 0xff, 0x59,
	0x89, 0xb9, 0xe8, 0x79, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// GreetServiceClient is the client API for GreetService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type GreetServiceClient interface {
	//BiDi Streaming
	GreetEveryone(ctx context.Context, opts ...grpc.CallOption) (GreetService_GreetEveryoneClient, error)
}

type greetServiceClient struct {
	cc *grpc.ClientConn
}

func NewGreetServiceClient(cc *grpc.ClientConn) GreetServiceClient {
	return &greetServiceClient{cc}
}

func (c *greetServiceClient) GreetEveryone(ctx context.Context, opts ...grpc.CallOption) (GreetService_GreetEveryoneClient, error) {
	stream, err := c.cc.NewStream(ctx, &_GreetService_serviceDesc.Streams[0], "/greet.GreetService/GreetEveryone", opts...)
	if err != nil {
		return nil, err
	}
	x := &greetServiceGreetEveryoneClient{stream}
	return x, nil
}

type GreetService_GreetEveryoneClient interface {
	Send(*GreetEveryoneRequest) error
	Recv() (*GreetEveryoneResponse, error)
	grpc.ClientStream
}

type greetServiceGreetEveryoneClient struct {
	grpc.ClientStream
}

func (x *greetServiceGreetEveryoneClient) Send(m *GreetEveryoneRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *greetServiceGreetEveryoneClient) Recv() (*GreetEveryoneResponse, error) {
	m := new(GreetEveryoneResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// GreetServiceServer is the server API for GreetService service.
type GreetServiceServer interface {
	//BiDi Streaming
	GreetEveryone(GreetService_GreetEveryoneServer) error
}

// UnimplementedGreetServiceServer can be embedded to have forward compatible implementations.
type UnimplementedGreetServiceServer struct {
}

func (*UnimplementedGreetServiceServer) GreetEveryone(srv GreetService_GreetEveryoneServer) error {
	return status.Errorf(codes.Unimplemented, "method GreetEveryone not implemented")
}

func RegisterGreetServiceServer(s *grpc.Server, srv GreetServiceServer) {
	s.RegisterService(&_GreetService_serviceDesc, srv)
}

func _GreetService_GreetEveryone_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(GreetServiceServer).GreetEveryone(&greetServiceGreetEveryoneServer{stream})
}

type GreetService_GreetEveryoneServer interface {
	Send(*GreetEveryoneResponse) error
	Recv() (*GreetEveryoneRequest, error)
	grpc.ServerStream
}

type greetServiceGreetEveryoneServer struct {
	grpc.ServerStream
}

func (x *greetServiceGreetEveryoneServer) Send(m *GreetEveryoneResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *greetServiceGreetEveryoneServer) Recv() (*GreetEveryoneRequest, error) {
	m := new(GreetEveryoneRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _GreetService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "greet.GreetService",
	HandlerType: (*GreetServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GreetEveryone",
			Handler:       _GreetService_GreetEveryone_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "Section8_grpcbidirectionalstreaming/greet/greetpb/greet.proto",
}
