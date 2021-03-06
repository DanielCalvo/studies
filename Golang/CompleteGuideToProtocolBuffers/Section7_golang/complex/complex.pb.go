// Code generated by protoc-gen-go. DO NOT EDIT.
// source: complex/complex.proto

package complexpb

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type ComplexMessage struct {
	OneDummy             *DummyMessage   `protobuf:"bytes,2,opt,name=one_dummy,json=oneDummy,proto3" json:"one_dummy,omitempty"`
	MultipleDummy        []*DummyMessage `protobuf:"bytes,3,rep,name=multiple_dummy,json=multipleDummy,proto3" json:"multiple_dummy,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *ComplexMessage) Reset()         { *m = ComplexMessage{} }
func (m *ComplexMessage) String() string { return proto.CompactTextString(m) }
func (*ComplexMessage) ProtoMessage()    {}
func (*ComplexMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_467978440e6d1735, []int{0}
}

func (m *ComplexMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ComplexMessage.Unmarshal(m, b)
}
func (m *ComplexMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ComplexMessage.Marshal(b, m, deterministic)
}
func (m *ComplexMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ComplexMessage.Merge(m, src)
}
func (m *ComplexMessage) XXX_Size() int {
	return xxx_messageInfo_ComplexMessage.Size(m)
}
func (m *ComplexMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_ComplexMessage.DiscardUnknown(m)
}

var xxx_messageInfo_ComplexMessage proto.InternalMessageInfo

func (m *ComplexMessage) GetOneDummy() *DummyMessage {
	if m != nil {
		return m.OneDummy
	}
	return nil
}

func (m *ComplexMessage) GetMultipleDummy() []*DummyMessage {
	if m != nil {
		return m.MultipleDummy
	}
	return nil
}

type DummyMessage struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DummyMessage) Reset()         { *m = DummyMessage{} }
func (m *DummyMessage) String() string { return proto.CompactTextString(m) }
func (*DummyMessage) ProtoMessage()    {}
func (*DummyMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_467978440e6d1735, []int{1}
}

func (m *DummyMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DummyMessage.Unmarshal(m, b)
}
func (m *DummyMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DummyMessage.Marshal(b, m, deterministic)
}
func (m *DummyMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DummyMessage.Merge(m, src)
}
func (m *DummyMessage) XXX_Size() int {
	return xxx_messageInfo_DummyMessage.Size(m)
}
func (m *DummyMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_DummyMessage.DiscardUnknown(m)
}

var xxx_messageInfo_DummyMessage proto.InternalMessageInfo

func (m *DummyMessage) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *DummyMessage) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func init() {
	proto.RegisterType((*ComplexMessage)(nil), "complexpb.ComplexMessage")
	proto.RegisterType((*DummyMessage)(nil), "complexpb.DummyMessage")
}

func init() { proto.RegisterFile("complex/complex.proto", fileDescriptor_467978440e6d1735) }

var fileDescriptor_467978440e6d1735 = []byte{
	// 163 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x4d, 0xce, 0xcf, 0x2d,
	0xc8, 0x49, 0xad, 0xd0, 0x87, 0xd2, 0x7a, 0x05, 0x45, 0xf9, 0x25, 0xf9, 0x42, 0x9c, 0x50, 0x6e,
	0x41, 0x92, 0x52, 0x1b, 0x23, 0x17, 0x9f, 0x33, 0x84, 0xe7, 0x9b, 0x5a, 0x5c, 0x9c, 0x98, 0x9e,
	0x2a, 0x64, 0xc2, 0xc5, 0x99, 0x9f, 0x97, 0x1a, 0x9f, 0x52, 0x9a, 0x9b, 0x5b, 0x29, 0xc1, 0xa4,
	0xc0, 0xa8, 0xc1, 0x6d, 0x24, 0xae, 0x07, 0xd7, 0xa1, 0xe7, 0x02, 0x12, 0x87, 0xaa, 0x0d, 0xe2,
	0xc8, 0xcf, 0x4b, 0x05, 0x0b, 0x08, 0xd9, 0x71, 0xf1, 0xe5, 0x96, 0xe6, 0x94, 0x64, 0x16, 0xe4,
	0xc0, 0xb4, 0x32, 0x2b, 0x30, 0xe3, 0xd3, 0xca, 0x0b, 0x53, 0x0e, 0x16, 0x55, 0x32, 0xe2, 0xe2,
	0x41, 0x96, 0x16, 0xe2, 0xe3, 0x62, 0xca, 0x4c, 0x91, 0x60, 0x54, 0x60, 0xd4, 0x60, 0x0d, 0x62,
	0xca, 0x4c, 0x11, 0x12, 0xe2, 0x62, 0xc9, 0x4b, 0xcc, 0x4d, 0x05, 0x3b, 0x88, 0x33, 0x08, 0xcc,
	0x4e, 0x62, 0x03, 0x7b, 0xc7, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0x83, 0x68, 0xac, 0xe7, 0xe7,
	0x00, 0x00, 0x00,
}
