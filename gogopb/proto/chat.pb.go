// Code generated by protoc-gen-go. DO NOT EDIT.
// source: chat.proto

package proto

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

type ChatReq struct {
	Content              string   `protobuf:"bytes,1,opt,name=content,proto3" json:"content,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ChatReq) Reset()         { *m = ChatReq{} }
func (m *ChatReq) String() string { return proto.CompactTextString(m) }
func (*ChatReq) ProtoMessage()    {}
func (*ChatReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_8c585a45e2093e54, []int{0}
}

func (m *ChatReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ChatReq.Unmarshal(m, b)
}
func (m *ChatReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ChatReq.Marshal(b, m, deterministic)
}
func (m *ChatReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChatReq.Merge(m, src)
}
func (m *ChatReq) XXX_Size() int {
	return xxx_messageInfo_ChatReq.Size(m)
}
func (m *ChatReq) XXX_DiscardUnknown() {
	xxx_messageInfo_ChatReq.DiscardUnknown(m)
}

var xxx_messageInfo_ChatReq proto.InternalMessageInfo

func (m *ChatReq) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

type ChatAck struct {
	Status               int32    `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	Content              string   `protobuf:"bytes,2,opt,name=content,proto3" json:"content,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ChatAck) Reset()         { *m = ChatAck{} }
func (m *ChatAck) String() string { return proto.CompactTextString(m) }
func (*ChatAck) ProtoMessage()    {}
func (*ChatAck) Descriptor() ([]byte, []int) {
	return fileDescriptor_8c585a45e2093e54, []int{1}
}

func (m *ChatAck) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ChatAck.Unmarshal(m, b)
}
func (m *ChatAck) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ChatAck.Marshal(b, m, deterministic)
}
func (m *ChatAck) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChatAck.Merge(m, src)
}
func (m *ChatAck) XXX_Size() int {
	return xxx_messageInfo_ChatAck.Size(m)
}
func (m *ChatAck) XXX_DiscardUnknown() {
	xxx_messageInfo_ChatAck.DiscardUnknown(m)
}

var xxx_messageInfo_ChatAck proto.InternalMessageInfo

func (m *ChatAck) GetStatus() int32 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *ChatAck) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

type BytesReq struct {
	Data                 []byte   `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BytesReq) Reset()         { *m = BytesReq{} }
func (m *BytesReq) String() string { return proto.CompactTextString(m) }
func (*BytesReq) ProtoMessage()    {}
func (*BytesReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_8c585a45e2093e54, []int{2}
}

func (m *BytesReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BytesReq.Unmarshal(m, b)
}
func (m *BytesReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BytesReq.Marshal(b, m, deterministic)
}
func (m *BytesReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BytesReq.Merge(m, src)
}
func (m *BytesReq) XXX_Size() int {
	return xxx_messageInfo_BytesReq.Size(m)
}
func (m *BytesReq) XXX_DiscardUnknown() {
	xxx_messageInfo_BytesReq.DiscardUnknown(m)
}

var xxx_messageInfo_BytesReq proto.InternalMessageInfo

func (m *BytesReq) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

type BytesAck struct {
	Status               int32    `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	Data                 []byte   `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BytesAck) Reset()         { *m = BytesAck{} }
func (m *BytesAck) String() string { return proto.CompactTextString(m) }
func (*BytesAck) ProtoMessage()    {}
func (*BytesAck) Descriptor() ([]byte, []int) {
	return fileDescriptor_8c585a45e2093e54, []int{3}
}

func (m *BytesAck) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BytesAck.Unmarshal(m, b)
}
func (m *BytesAck) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BytesAck.Marshal(b, m, deterministic)
}
func (m *BytesAck) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BytesAck.Merge(m, src)
}
func (m *BytesAck) XXX_Size() int {
	return xxx_messageInfo_BytesAck.Size(m)
}
func (m *BytesAck) XXX_DiscardUnknown() {
	xxx_messageInfo_BytesAck.DiscardUnknown(m)
}

var xxx_messageInfo_BytesAck proto.InternalMessageInfo

func (m *BytesAck) GetStatus() int32 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *BytesAck) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*ChatReq)(nil), "proto.ChatReq")
	proto.RegisterType((*ChatAck)(nil), "proto.ChatAck")
	proto.RegisterType((*BytesReq)(nil), "proto.BytesReq")
	proto.RegisterType((*BytesAck)(nil), "proto.BytesAck")
}

func init() { proto.RegisterFile("chat.proto", fileDescriptor_8c585a45e2093e54) }

var fileDescriptor_8c585a45e2093e54 = []byte{
	// 138 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4a, 0xce, 0x48, 0x2c,
	0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x05, 0x53, 0x4a, 0xca, 0x5c, 0xec, 0xce, 0x19,
	0x89, 0x25, 0x41, 0xa9, 0x85, 0x42, 0x12, 0x5c, 0xec, 0xc9, 0xf9, 0x79, 0x25, 0xa9, 0x79, 0x25,
	0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0x9c, 0x41, 0x30, 0xae, 0x92, 0x35, 0x44, 0x91, 0x63, 0x72, 0xb6,
	0x90, 0x18, 0x17, 0x5b, 0x71, 0x49, 0x62, 0x49, 0x69, 0x31, 0x58, 0x0d, 0x6b, 0x10, 0x94, 0x87,
	0xac, 0x99, 0x09, 0x55, 0xb3, 0x1c, 0x17, 0x87, 0x53, 0x65, 0x49, 0x6a, 0x31, 0xc8, 0x0a, 0x21,
	0x2e, 0x96, 0x94, 0xc4, 0x92, 0x44, 0xb0, 0x5e, 0x9e, 0x20, 0x30, 0x5b, 0xc9, 0x0c, 0x2a, 0x8f,
	0xcf, 0x74, 0x98, 0x3e, 0x26, 0x84, 0xbe, 0x24, 0x36, 0xb0, 0x07, 0x8c, 0x01, 0x01, 0x00, 0x00,
	0xff, 0xff, 0xb6, 0x1a, 0x36, 0x4a, 0xd5, 0x00, 0x00, 0x00,
}