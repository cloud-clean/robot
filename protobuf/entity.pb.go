// Code generated by protoc-gen-go. DO NOT EDIT.
// source: entity.proto

package entity

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type MsgType int32

const (
	MsgType_PING MsgType = 0
	MsgType_PONG MsgType = 1
	MsgType_CMD  MsgType = 2
	MsgType_MSG  MsgType = 3
)

var MsgType_name = map[int32]string{
	0: "PING",
	1: "PONG",
	2: "CMD",
	3: "MSG",
}
var MsgType_value = map[string]int32{
	"PING": 0,
	"PONG": 1,
	"CMD":  2,
	"MSG":  3,
}

func (x MsgType) String() string {
	return proto.EnumName(MsgType_name, int32(x))
}
func (MsgType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_entity_79ec8738b853913d, []int{0}
}

type CmdMsg struct {
	Magic                uint32   `protobuf:"varint,1,opt,name=magic,proto3" json:"magic,omitempty"`
	Type                 MsgType  `protobuf:"varint,2,opt,name=type,proto3,enum=entity.MsgType" json:"type,omitempty"`
	Data                 []byte   `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CmdMsg) Reset()         { *m = CmdMsg{} }
func (m *CmdMsg) String() string { return proto.CompactTextString(m) }
func (*CmdMsg) ProtoMessage()    {}
func (*CmdMsg) Descriptor() ([]byte, []int) {
	return fileDescriptor_entity_79ec8738b853913d, []int{0}
}
func (m *CmdMsg) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CmdMsg.Unmarshal(m, b)
}
func (m *CmdMsg) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CmdMsg.Marshal(b, m, deterministic)
}
func (dst *CmdMsg) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CmdMsg.Merge(dst, src)
}
func (m *CmdMsg) XXX_Size() int {
	return xxx_messageInfo_CmdMsg.Size(m)
}
func (m *CmdMsg) XXX_DiscardUnknown() {
	xxx_messageInfo_CmdMsg.DiscardUnknown(m)
}

var xxx_messageInfo_CmdMsg proto.InternalMessageInfo

func (m *CmdMsg) GetMagic() uint32 {
	if m != nil {
		return m.Magic
	}
	return 0
}

func (m *CmdMsg) GetType() MsgType {
	if m != nil {
		return m.Type
	}
	return MsgType_PING
}

func (m *CmdMsg) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*CmdMsg)(nil), "entity.CmdMsg")
	proto.RegisterEnum("entity.MsgType", MsgType_name, MsgType_value)
}

func init() { proto.RegisterFile("entity.proto", fileDescriptor_entity_79ec8738b853913d) }

var fileDescriptor_entity_79ec8738b853913d = []byte{
	// 159 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x49, 0xcd, 0x2b, 0xc9,
	0x2c, 0xa9, 0xd4, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x83, 0xf0, 0x94, 0xc2, 0xb9, 0xd8,
	0x9c, 0x73, 0x53, 0x7c, 0x8b, 0xd3, 0x85, 0x44, 0xb8, 0x58, 0x73, 0x13, 0xd3, 0x33, 0x93, 0x25,
	0x18, 0x15, 0x18, 0x35, 0x78, 0x83, 0x20, 0x1c, 0x21, 0x65, 0x2e, 0x96, 0x92, 0xca, 0x82, 0x54,
	0x09, 0x26, 0x05, 0x46, 0x0d, 0x3e, 0x23, 0x7e, 0x3d, 0xa8, 0x21, 0xbe, 0xc5, 0xe9, 0x21, 0x95,
	0x05, 0xa9, 0x41, 0x60, 0x49, 0x21, 0x21, 0x2e, 0x96, 0x94, 0xc4, 0x92, 0x44, 0x09, 0x66, 0x05,
	0x46, 0x0d, 0x9e, 0x20, 0x30, 0x5b, 0x4b, 0x9f, 0x8b, 0x1d, 0xaa, 0x48, 0x88, 0x83, 0x8b, 0x25,
	0xc0, 0xd3, 0xcf, 0x5d, 0x80, 0x01, 0xcc, 0xf2, 0xf7, 0x73, 0x17, 0x60, 0x14, 0x62, 0xe7, 0x62,
	0x76, 0xf6, 0x75, 0x11, 0x60, 0x02, 0x31, 0x7c, 0x83, 0xdd, 0x05, 0x98, 0x93, 0xd8, 0xc0, 0x0e,
	0x33, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0xcb, 0xb3, 0xdf, 0x56, 0xa8, 0x00, 0x00, 0x00,
}
