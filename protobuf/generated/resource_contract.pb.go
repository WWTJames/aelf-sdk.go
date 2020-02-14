// Code generated by protoc-gen-go. DO NOT EDIT.
// source: resource_contract.proto

package client

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

type ResourceType int32

const (
	ResourceType_UNDEFINED_RESOURCE_TYPE ResourceType = 0
	ResourceType_RAM                     ResourceType = 1
	ResourceType_CPU                     ResourceType = 2
	ResourceType_NET                     ResourceType = 3
	ResourceType_STO                     ResourceType = 4
)

var ResourceType_name = map[int32]string{
	0: "UNDEFINED_RESOURCE_TYPE",
	1: "RAM",
	2: "CPU",
	3: "NET",
	4: "STO",
}

var ResourceType_value = map[string]int32{
	"UNDEFINED_RESOURCE_TYPE": 0,
	"RAM":                     1,
	"CPU":                     2,
	"NET":                     3,
	"STO":                     4,
}

func (x ResourceType) String() string {
	return proto.EnumName(ResourceType_name, int32(x))
}

func (ResourceType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_7b254ff96e3bb431, []int{0}
}

type ResourceId struct {
	Type                 ResourceType `protobuf:"varint,1,opt,name=type,proto3,enum=client.ResourceType" json:"type" form:"type"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *ResourceId) Reset()         { *m = ResourceId{} }
func (m *ResourceId) String() string { return proto.CompactTextString(m) }
func (*ResourceId) ProtoMessage()    {}
func (*ResourceId) Descriptor() ([]byte, []int) {
	return fileDescriptor_7b254ff96e3bb431, []int{0}
}

func (m *ResourceId) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ResourceId.Unmarshal(m, b)
}
func (m *ResourceId) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ResourceId.Marshal(b, m, deterministic)
}
func (m *ResourceId) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResourceId.Merge(m, src)
}
func (m *ResourceId) XXX_Size() int {
	return xxx_messageInfo_ResourceId.Size(m)
}
func (m *ResourceId) XXX_DiscardUnknown() {
	xxx_messageInfo_ResourceId.DiscardUnknown(m)
}

var xxx_messageInfo_ResourceId proto.InternalMessageInfo

func (m *ResourceId) GetType() ResourceType {
	if m != nil {
		return m.Type
	}
	return ResourceType_UNDEFINED_RESOURCE_TYPE
}

type Converter struct {
	ResBalance           int64        `protobuf:"zigzag64,1,opt,name=res_balance,json=resBalance,proto3" json:"res_balance" form:"res_balance"`
	ElfBalance           int64        `protobuf:"zigzag64,2,opt,name=elf_balance,json=elfBalance,proto3" json:"elf_balance" form:"elf_balance"`
	ResWeight            int64        `protobuf:"zigzag64,3,opt,name=res_weight,json=resWeight,proto3" json:"res_weight" form:"res_weight"`
	ElfWeight            int64        `protobuf:"zigzag64,4,opt,name=elf_weight,json=elfWeight,proto3" json:"elf_weight" form:"elf_weight"`
	Type                 ResourceType `protobuf:"varint,5,opt,name=type,proto3,enum=client.ResourceType" json:"type" form:"type"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *Converter) Reset()         { *m = Converter{} }
func (m *Converter) String() string { return proto.CompactTextString(m) }
func (*Converter) ProtoMessage()    {}
func (*Converter) Descriptor() ([]byte, []int) {
	return fileDescriptor_7b254ff96e3bb431, []int{1}
}

func (m *Converter) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Converter.Unmarshal(m, b)
}
func (m *Converter) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Converter.Marshal(b, m, deterministic)
}
func (m *Converter) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Converter.Merge(m, src)
}
func (m *Converter) XXX_Size() int {
	return xxx_messageInfo_Converter.Size(m)
}
func (m *Converter) XXX_DiscardUnknown() {
	xxx_messageInfo_Converter.DiscardUnknown(m)
}

var xxx_messageInfo_Converter proto.InternalMessageInfo

func (m *Converter) GetResBalance() int64 {
	if m != nil {
		return m.ResBalance
	}
	return 0
}

func (m *Converter) GetElfBalance() int64 {
	if m != nil {
		return m.ElfBalance
	}
	return 0
}

func (m *Converter) GetResWeight() int64 {
	if m != nil {
		return m.ResWeight
	}
	return 0
}

func (m *Converter) GetElfWeight() int64 {
	if m != nil {
		return m.ElfWeight
	}
	return 0
}

func (m *Converter) GetType() ResourceType {
	if m != nil {
		return m.Type
	}
	return ResourceType_UNDEFINED_RESOURCE_TYPE
}

type UserResourceId struct {
	Address              *Address     `protobuf:"bytes,1,opt,name=address,proto3" json:"address" form:"address"`
	Type                 ResourceType `protobuf:"varint,2,opt,name=type,proto3,enum=client.ResourceType" json:"type" form:"type"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *UserResourceId) Reset()         { *m = UserResourceId{} }
func (m *UserResourceId) String() string { return proto.CompactTextString(m) }
func (*UserResourceId) ProtoMessage()    {}
func (*UserResourceId) Descriptor() ([]byte, []int) {
	return fileDescriptor_7b254ff96e3bb431, []int{2}
}

func (m *UserResourceId) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserResourceId.Unmarshal(m, b)
}
func (m *UserResourceId) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserResourceId.Marshal(b, m, deterministic)
}
func (m *UserResourceId) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserResourceId.Merge(m, src)
}
func (m *UserResourceId) XXX_Size() int {
	return xxx_messageInfo_UserResourceId.Size(m)
}
func (m *UserResourceId) XXX_DiscardUnknown() {
	xxx_messageInfo_UserResourceId.DiscardUnknown(m)
}

var xxx_messageInfo_UserResourceId proto.InternalMessageInfo

func (m *UserResourceId) GetAddress() *Address {
	if m != nil {
		return m.Address
	}
	return nil
}

func (m *UserResourceId) GetType() ResourceType {
	if m != nil {
		return m.Type
	}
	return ResourceType_UNDEFINED_RESOURCE_TYPE
}

func init() {
	proto.RegisterEnum("client.ResourceType", ResourceType_name, ResourceType_value)
	proto.RegisterType((*ResourceId)(nil), "client.ResourceId")
	proto.RegisterType((*Converter)(nil), "client.Converter")
	proto.RegisterType((*UserResourceId)(nil), "client.UserResourceId")
}

func init() { proto.RegisterFile("resource_contract.proto", fileDescriptor_7b254ff96e3bb431) }

var fileDescriptor_7b254ff96e3bb431 = []byte{
	// 293 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x91, 0x4d, 0x4b, 0xc3, 0x40,
	0x10, 0x86, 0xcd, 0x87, 0x2d, 0x9d, 0x96, 0x1a, 0x16, 0xa1, 0x45, 0x11, 0xa5, 0xa7, 0xea, 0xa1,
	0x87, 0x0a, 0xde, 0x6b, 0xba, 0x42, 0x0f, 0x26, 0x65, 0x9b, 0x20, 0x9e, 0x42, 0x9a, 0x4c, 0xb4,
	0x10, 0x92, 0xb0, 0xbb, 0x2a, 0xfd, 0x61, 0xfe, 0xbf, 0xb2, 0xd9, 0xa4, 0xf4, 0xd6, 0x53, 0x86,
	0xf7, 0x79, 0x5e, 0x66, 0xc2, 0xc2, 0x88, 0xa3, 0x28, 0x7f, 0x78, 0x82, 0x51, 0x52, 0x16, 0x92,
	0xc7, 0x89, 0x9c, 0x55, 0xbc, 0x94, 0x25, 0xe9, 0x24, 0xf9, 0x0e, 0x0b, 0x79, 0x33, 0xd0, 0x5f,
	0x9d, 0x4e, 0x5e, 0x00, 0x58, 0x53, 0x58, 0xa5, 0x64, 0x0a, 0xb6, 0xdc, 0x57, 0x38, 0x36, 0x1e,
	0x8c, 0xe9, 0x70, 0x7e, 0x3d, 0x6b, 0xd4, 0xd6, 0x08, 0xf6, 0x15, 0xb2, 0xda, 0x98, 0xfc, 0x1b,
	0xd0, 0x73, 0xcb, 0xe2, 0x17, 0xb9, 0x44, 0x4e, 0xee, 0xa1, 0xcf, 0x51, 0x44, 0xdb, 0x38, 0x8f,
	0x8b, 0x44, 0xd7, 0x09, 0x03, 0x8e, 0xe2, 0x55, 0x27, 0x4a, 0xc0, 0x3c, 0x3b, 0x0a, 0xa6, 0x16,
	0x30, 0xcf, 0x5a, 0xe1, 0x0e, 0x94, 0x1e, 0xfd, 0xe1, 0xee, 0xeb, 0x5b, 0x8e, 0xad, 0x9a, 0xf7,
	0x38, 0x8a, 0x8f, 0x3a, 0x50, 0x58, 0xf5, 0x1b, 0x6c, 0x6b, 0x8c, 0x79, 0xd6, 0xe0, 0xf6, 0xee,
	0xcb, 0xb3, 0x77, 0x23, 0x0c, 0x43, 0x81, 0xfc, 0xe4, 0x9f, 0x1f, 0xa1, 0x1b, 0xa7, 0x29, 0x47,
	0x21, 0xea, 0xbb, 0xfb, 0xf3, 0xab, 0xb6, 0xbe, 0xd0, 0x31, 0x6b, 0xf9, 0x71, 0x8d, 0x79, 0x6e,
	0xcd, 0x93, 0x0f, 0x83, 0xd3, 0x94, 0xdc, 0xc2, 0x28, 0xf4, 0x96, 0xf4, 0x6d, 0xe5, 0xd1, 0x65,
	0xc4, 0xe8, 0xc6, 0x0f, 0x99, 0x4b, 0xa3, 0xe0, 0x73, 0x4d, 0x9d, 0x0b, 0xd2, 0x05, 0x8b, 0x2d,
	0xde, 0x1d, 0x43, 0x0d, 0xee, 0x3a, 0x74, 0x4c, 0x35, 0x78, 0x34, 0x70, 0x2c, 0x35, 0x6c, 0x02,
	0xdf, 0xb1, 0xb7, 0x9d, 0xfa, 0xb9, 0x9e, 0x0f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x94, 0x19, 0xa1,
	0x4f, 0xdf, 0x01, 0x00, 0x00,
}