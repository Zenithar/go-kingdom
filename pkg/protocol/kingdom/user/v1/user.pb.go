// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: pkg/protocol/kingdom/user/v1/user.proto

package userv1

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
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

// User is an object group.
type User struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Label                string   `protobuf:"bytes,2,opt,name=label,proto3" json:"label,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_b70e813c2f6610d1, []int{0}
}

func (m *User) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_User.Unmarshal(m, b)
}

func (m *User) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_User.Marshal(b, m, deterministic)
}

func (m *User) XXX_Merge(src proto.Message) {
	xxx_messageInfo_User.Merge(m, src)
}

func (m *User) XXX_Size() int {
	return xxx_messageInfo_User.Size(m)
}

func (m *User) XXX_DiscardUnknown() {
	xxx_messageInfo_User.DiscardUnknown(m)
}

var xxx_messageInfo_User proto.InternalMessageInfo

func (m *User) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *User) GetLabel() string {
	if m != nil {
		return m.Label
	}
	return ""
}

func init() {
	proto.RegisterType((*User)(nil), "kingdom.user.v1.User")
}

func init() {
	proto.RegisterFile("pkg/protocol/kingdom/user/v1/user.proto", fileDescriptor_b70e813c2f6610d1)
}

var fileDescriptor_b70e813c2f6610d1 = []byte{
	// 172 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x2f, 0xc8, 0x4e, 0xd7,
	0x2f, 0x28, 0xca, 0x2f, 0xc9, 0x4f, 0xce, 0xcf, 0xd1, 0xcf, 0xce, 0xcc, 0x4b, 0x4f, 0xc9, 0xcf,
	0xd5, 0x2f, 0x2d, 0x4e, 0x2d, 0xd2, 0x2f, 0x33, 0x04, 0xd3, 0x7a, 0x60, 0x59, 0x21, 0x7e, 0xa8,
	0x9c, 0x1e, 0x58, 0xac, 0xcc, 0x50, 0x49, 0x87, 0x8b, 0x25, 0xb4, 0x38, 0xb5, 0x48, 0x88, 0x8f,
	0x8b, 0x29, 0x33, 0x45, 0x82, 0x51, 0x81, 0x51, 0x83, 0x33, 0x88, 0x29, 0x33, 0x45, 0x48, 0x84,
	0x8b, 0x35, 0x27, 0x31, 0x29, 0x35, 0x47, 0x82, 0x09, 0x2c, 0x04, 0xe1, 0x38, 0xc5, 0x72, 0xc9,
	0xe4, 0x17, 0xa5, 0xeb, 0x55, 0xa5, 0xe6, 0x65, 0x96, 0x64, 0x24, 0x16, 0xe9, 0xa1, 0x99, 0xe6,
	0xc4, 0x09, 0x32, 0x2b, 0x00, 0x64, 0x53, 0x00, 0x63, 0x14, 0x1b, 0x48, 0xb4, 0xcc, 0x70, 0x11,
	0x13, 0xb3, 0x77, 0x68, 0xc4, 0x2a, 0x26, 0x7e, 0x6f, 0xa8, 0x62, 0x90, 0x1a, 0xbd, 0x30, 0xc3,
	0x53, 0x70, 0x91, 0x18, 0x90, 0x48, 0x4c, 0x98, 0x61, 0x12, 0x1b, 0xd8, 0x91, 0xc6, 0x80, 0x00,
	0x00, 0x00, 0xff, 0xff, 0xa8, 0xe2, 0xc7, 0xa8, 0xcf, 0x00, 0x00, 0x00,
}