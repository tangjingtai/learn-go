// Code generated by protoc-gen-go. DO NOT EDIT.
// source: test.proto

//包名，通过protoc生成时go文件时

package test

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

// 手机号码类型
type PhoneType int32

const (
	PhoneType_HOME PhoneType = 0
	PhoneType_WORK PhoneType = 1
)

var PhoneType_name = map[int32]string{
	0: "HOME",
	1: "WORK",
}

var PhoneType_value = map[string]int32{
	"HOME": 0,
	"WORK": 1,
}

func (x PhoneType) String() string {
	return proto.EnumName(PhoneType_name, int32(x))
}

func (PhoneType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_c161fcfdc0c3ff1e, []int{0}
}

// 手机号码
type Phone struct {
	Type                 PhoneType `protobuf:"varint,1,opt,name=type,proto3,enum=test.PhoneType" json:"type,omitempty"`
	Number               string    `protobuf:"bytes,2,opt,name=number,proto3" json:"number,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *Phone) Reset()         { *m = Phone{} }
func (m *Phone) String() string { return proto.CompactTextString(m) }
func (*Phone) ProtoMessage()    {}
func (*Phone) Descriptor() ([]byte, []int) {
	return fileDescriptor_c161fcfdc0c3ff1e, []int{0}
}

func (m *Phone) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Phone.Unmarshal(m, b)
}
func (m *Phone) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Phone.Marshal(b, m, deterministic)
}
func (m *Phone) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Phone.Merge(m, src)
}
func (m *Phone) XXX_Size() int {
	return xxx_messageInfo_Phone.Size(m)
}
func (m *Phone) XXX_DiscardUnknown() {
	xxx_messageInfo_Phone.DiscardUnknown(m)
}

var xxx_messageInfo_Phone proto.InternalMessageInfo

func (m *Phone) GetType() PhoneType {
	if m != nil {
		return m.Type
	}
	return PhoneType_HOME
}

func (m *Phone) GetNumber() string {
	if m != nil {
		return m.Number
	}
	return ""
}

// 联系人
type Person struct {
	Id     int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name   string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Phones []*Phone `protobuf:"bytes,3,rep,name=phones,proto3" json:"phones,omitempty"`
	// 测试负数
	Id2                  int32    `protobuf:"zigzag32,4,opt,name=id2,proto3" json:"id2,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Person) Reset()         { *m = Person{} }
func (m *Person) String() string { return proto.CompactTextString(m) }
func (*Person) ProtoMessage()    {}
func (*Person) Descriptor() ([]byte, []int) {
	return fileDescriptor_c161fcfdc0c3ff1e, []int{1}
}

func (m *Person) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Person.Unmarshal(m, b)
}
func (m *Person) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Person.Marshal(b, m, deterministic)
}
func (m *Person) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Person.Merge(m, src)
}
func (m *Person) XXX_Size() int {
	return xxx_messageInfo_Person.Size(m)
}
func (m *Person) XXX_DiscardUnknown() {
	xxx_messageInfo_Person.DiscardUnknown(m)
}

var xxx_messageInfo_Person proto.InternalMessageInfo

func (m *Person) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Person) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Person) GetPhones() []*Phone {
	if m != nil {
		return m.Phones
	}
	return nil
}

func (m *Person) GetId2() int32 {
	if m != nil {
		return m.Id2
	}
	return 0
}

// 通讯录
type ContactBook struct {
	Persons              []*Person `protobuf:"bytes,1,rep,name=persons,proto3" json:"persons,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *ContactBook) Reset()         { *m = ContactBook{} }
func (m *ContactBook) String() string { return proto.CompactTextString(m) }
func (*ContactBook) ProtoMessage()    {}
func (*ContactBook) Descriptor() ([]byte, []int) {
	return fileDescriptor_c161fcfdc0c3ff1e, []int{2}
}

func (m *ContactBook) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ContactBook.Unmarshal(m, b)
}
func (m *ContactBook) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ContactBook.Marshal(b, m, deterministic)
}
func (m *ContactBook) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ContactBook.Merge(m, src)
}
func (m *ContactBook) XXX_Size() int {
	return xxx_messageInfo_ContactBook.Size(m)
}
func (m *ContactBook) XXX_DiscardUnknown() {
	xxx_messageInfo_ContactBook.DiscardUnknown(m)
}

var xxx_messageInfo_ContactBook proto.InternalMessageInfo

func (m *ContactBook) GetPersons() []*Person {
	if m != nil {
		return m.Persons
	}
	return nil
}

func init() {
	proto.RegisterEnum("test.PhoneType", PhoneType_name, PhoneType_value)
	proto.RegisterType((*Phone)(nil), "test.Phone")
	proto.RegisterType((*Person)(nil), "test.Person")
	proto.RegisterType((*ContactBook)(nil), "test.ContactBook")
}

func init() { proto.RegisterFile("test.proto", fileDescriptor_c161fcfdc0c3ff1e) }

var fileDescriptor_c161fcfdc0c3ff1e = []byte{
	// 226 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x90, 0x4d, 0x4b, 0x03, 0x31,
	0x10, 0x40, 0xcd, 0x6e, 0x1a, 0xed, 0xac, 0xd4, 0x75, 0x0e, 0x92, 0x9b, 0x61, 0x0b, 0x12, 0x3c,
	0xf4, 0xb0, 0xe2, 0x1f, 0xf0, 0x03, 0x04, 0x91, 0x96, 0x20, 0x78, 0x6e, 0xbb, 0x01, 0x17, 0x69,
	0x12, 0x36, 0xf1, 0xd0, 0x7f, 0x2f, 0x19, 0xa2, 0x78, 0x7b, 0xc9, 0x1b, 0xde, 0x90, 0x00, 0x24,
	0x1b, 0xd3, 0x2a, 0x4c, 0x3e, 0x79, 0xe4, 0x99, 0xbb, 0x27, 0x98, 0x6d, 0x3e, 0xbd, 0xb3, 0xb8,
	0x04, 0x9e, 0x8e, 0xc1, 0x4a, 0xa6, 0x98, 0x5e, 0xf4, 0x17, 0x2b, 0x9a, 0x24, 0xf5, 0x7e, 0x0c,
	0xd6, 0x90, 0xc4, 0x2b, 0x10, 0xee, 0xfb, 0xb0, 0xb3, 0x93, 0xac, 0x14, 0xd3, 0x73, 0x53, 0x4e,
	0xdd, 0x1e, 0xc4, 0xc6, 0x4e, 0xd1, 0x3b, 0x5c, 0x40, 0x35, 0x0e, 0x14, 0x99, 0x99, 0x6a, 0x1c,
	0x10, 0x81, 0xbb, 0xed, 0xc1, 0x96, 0x79, 0x62, 0x5c, 0x82, 0x08, 0x39, 0x1c, 0x65, 0xad, 0x6a,
	0xdd, 0xf4, 0xcd, 0xbf, 0x65, 0xa6, 0x28, 0x6c, 0xa1, 0x1e, 0x87, 0x5e, 0x72, 0xc5, 0xf4, 0xa5,
	0xc9, 0xd8, 0xdd, 0x43, 0xf3, 0xe8, 0x5d, 0xda, 0xee, 0xd3, 0x83, 0xf7, 0x5f, 0x78, 0x03, 0xa7,
	0x81, 0x76, 0x46, 0xc9, 0x28, 0x73, 0x5e, 0x32, 0x74, 0x69, 0x7e, 0xe5, 0xed, 0x35, 0xcc, 0xff,
	0x9e, 0x81, 0x67, 0xc0, 0x5f, 0xd6, 0x6f, 0xcf, 0xed, 0x49, 0xa6, 0x8f, 0xb5, 0x79, 0x6d, 0xd9,
	0x4e, 0xd0, 0x7f, 0xdc, 0xfd, 0x04, 0x00, 0x00, 0xff, 0xff, 0x28, 0x9e, 0x40, 0xe3, 0x1d, 0x01,
	0x00, 0x00,
}
