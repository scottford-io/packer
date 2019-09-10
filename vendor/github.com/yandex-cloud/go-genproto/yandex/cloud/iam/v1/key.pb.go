// Code generated by protoc-gen-go. DO NOT EDIT.
// source: yandex/cloud/iam/v1/key.proto

package iam

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

type Key_Algorithm int32

const (
	Key_ALGORITHM_UNSPECIFIED Key_Algorithm = 0
	// RSA with a 2048-bit key size. Default value.
	Key_RSA_2048 Key_Algorithm = 1
	// RSA with a 4096-bit key size.
	Key_RSA_4096 Key_Algorithm = 2
)

var Key_Algorithm_name = map[int32]string{
	0: "ALGORITHM_UNSPECIFIED",
	1: "RSA_2048",
	2: "RSA_4096",
}

var Key_Algorithm_value = map[string]int32{
	"ALGORITHM_UNSPECIFIED": 0,
	"RSA_2048":              1,
	"RSA_4096":              2,
}

func (x Key_Algorithm) String() string {
	return proto.EnumName(Key_Algorithm_name, int32(x))
}

func (Key_Algorithm) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_d3c4378eab1afe9e, []int{0, 0}
}

// A Key resource. For more information, see [Authorized keys](/docs/iam/concepts/authorization/key).
type Key struct {
	// ID of the Key resource.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Types that are valid to be assigned to Subject:
	//	*Key_UserAccountId
	//	*Key_ServiceAccountId
	Subject isKey_Subject `protobuf_oneof:"subject"`
	// Creation timestamp.
	CreatedAt *timestamp.Timestamp `protobuf:"bytes,4,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	// Description of the Key resource. 0-256 characters long.
	Description string `protobuf:"bytes,5,opt,name=description,proto3" json:"description,omitempty"`
	// An algorithm used to generate a key pair of the Key resource.
	KeyAlgorithm Key_Algorithm `protobuf:"varint,6,opt,name=key_algorithm,json=keyAlgorithm,proto3,enum=yandex.cloud.iam.v1.Key_Algorithm" json:"key_algorithm,omitempty"`
	// A public key of the Key resource.
	PublicKey            string   `protobuf:"bytes,7,opt,name=public_key,json=publicKey,proto3" json:"public_key,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Key) Reset()         { *m = Key{} }
func (m *Key) String() string { return proto.CompactTextString(m) }
func (*Key) ProtoMessage()    {}
func (*Key) Descriptor() ([]byte, []int) {
	return fileDescriptor_d3c4378eab1afe9e, []int{0}
}

func (m *Key) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Key.Unmarshal(m, b)
}
func (m *Key) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Key.Marshal(b, m, deterministic)
}
func (m *Key) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Key.Merge(m, src)
}
func (m *Key) XXX_Size() int {
	return xxx_messageInfo_Key.Size(m)
}
func (m *Key) XXX_DiscardUnknown() {
	xxx_messageInfo_Key.DiscardUnknown(m)
}

var xxx_messageInfo_Key proto.InternalMessageInfo

func (m *Key) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type isKey_Subject interface {
	isKey_Subject()
}

type Key_UserAccountId struct {
	UserAccountId string `protobuf:"bytes,2,opt,name=user_account_id,json=userAccountId,proto3,oneof"`
}

type Key_ServiceAccountId struct {
	ServiceAccountId string `protobuf:"bytes,3,opt,name=service_account_id,json=serviceAccountId,proto3,oneof"`
}

func (*Key_UserAccountId) isKey_Subject() {}

func (*Key_ServiceAccountId) isKey_Subject() {}

func (m *Key) GetSubject() isKey_Subject {
	if m != nil {
		return m.Subject
	}
	return nil
}

func (m *Key) GetUserAccountId() string {
	if x, ok := m.GetSubject().(*Key_UserAccountId); ok {
		return x.UserAccountId
	}
	return ""
}

func (m *Key) GetServiceAccountId() string {
	if x, ok := m.GetSubject().(*Key_ServiceAccountId); ok {
		return x.ServiceAccountId
	}
	return ""
}

func (m *Key) GetCreatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

func (m *Key) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Key) GetKeyAlgorithm() Key_Algorithm {
	if m != nil {
		return m.KeyAlgorithm
	}
	return Key_ALGORITHM_UNSPECIFIED
}

func (m *Key) GetPublicKey() string {
	if m != nil {
		return m.PublicKey
	}
	return ""
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*Key) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*Key_UserAccountId)(nil),
		(*Key_ServiceAccountId)(nil),
	}
}

func init() {
	proto.RegisterEnum("yandex.cloud.iam.v1.Key_Algorithm", Key_Algorithm_name, Key_Algorithm_value)
	proto.RegisterType((*Key)(nil), "yandex.cloud.iam.v1.Key")
}

func init() { proto.RegisterFile("yandex/cloud/iam/v1/key.proto", fileDescriptor_d3c4378eab1afe9e) }

var fileDescriptor_d3c4378eab1afe9e = []byte{
	// 384 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x92, 0x4b, 0x6f, 0xd4, 0x30,
	0x14, 0x85, 0x9b, 0x0c, 0xb4, 0xc4, 0x7d, 0x10, 0x19, 0x21, 0x85, 0x4a, 0x15, 0xd1, 0xac, 0xb2,
	0xa9, 0xdd, 0x0e, 0x15, 0xa2, 0xaa, 0x58, 0x64, 0xa0, 0xb4, 0xd1, 0xf0, 0x52, 0x5a, 0x36, 0x6c,
	0x22, 0xc7, 0xbe, 0xa4, 0x26, 0x0f, 0x47, 0x89, 0x33, 0xc2, 0x6b, 0xfe, 0x38, 0x6a, 0x32, 0x61,
	0x66, 0x31, 0xcb, 0x7b, 0xbe, 0xcf, 0xf2, 0xd1, 0xb5, 0xd1, 0x89, 0x61, 0x95, 0x80, 0x3f, 0x94,
	0x17, 0xaa, 0x13, 0x54, 0xb2, 0x92, 0x2e, 0xcf, 0x69, 0x0e, 0x86, 0xd4, 0x8d, 0xd2, 0x0a, 0xbf,
	0x18, 0x30, 0xe9, 0x31, 0x91, 0xac, 0x24, 0xcb, 0xf3, 0xe3, 0xd7, 0x99, 0x52, 0x59, 0x01, 0xb4,
	0x57, 0xd2, 0xee, 0x17, 0xd5, 0xb2, 0x84, 0x56, 0xb3, 0xb2, 0x1e, 0x4e, 0x4d, 0xff, 0x4e, 0xd0,
	0x64, 0x01, 0x06, 0x1f, 0x21, 0x5b, 0x0a, 0xcf, 0xf2, 0xad, 0xc0, 0x89, 0x6d, 0x29, 0x70, 0x80,
	0x9e, 0x77, 0x2d, 0x34, 0x09, 0xe3, 0x5c, 0x75, 0x95, 0x4e, 0xa4, 0xf0, 0xec, 0x47, 0x78, 0xbb,
	0x13, 0x1f, 0x3e, 0x82, 0x70, 0xc8, 0x23, 0x81, 0x09, 0xc2, 0x2d, 0x34, 0x4b, 0xc9, 0x61, 0x53,
	0x9e, 0xac, 0x64, 0x77, 0xc5, 0xd6, 0xfe, 0x25, 0x42, 0xbc, 0x01, 0xa6, 0x41, 0x24, 0x4c, 0x7b,
	0x4f, 0x7c, 0x2b, 0xd8, 0x9f, 0x1d, 0x93, 0xa1, 0x27, 0x19, 0x7b, 0x92, 0xfb, 0xb1, 0x67, 0xec,
	0xac, 0xec, 0x50, 0x63, 0x1f, 0xed, 0x0b, 0x68, 0x79, 0x23, 0x6b, 0x2d, 0x55, 0xe5, 0x3d, 0xed,
	0xdb, 0x6e, 0x46, 0xf8, 0x06, 0x1d, 0xe6, 0x60, 0x12, 0x56, 0x64, 0xaa, 0x91, 0xfa, 0xa1, 0xf4,
	0x76, 0x7d, 0x2b, 0x38, 0x9a, 0x4d, 0xc9, 0x96, 0xe5, 0x90, 0x05, 0x18, 0x12, 0x8e, 0x66, 0x7c,
	0x90, 0x83, 0xf9, 0x3f, 0xe1, 0x13, 0x84, 0xea, 0x2e, 0x2d, 0x24, 0x4f, 0x72, 0x30, 0xde, 0x5e,
	0x7f, 0x93, 0x33, 0x24, 0x0b, 0x30, 0xd3, 0x39, 0x72, 0xd6, 0xee, 0x2b, 0xf4, 0x32, 0xfc, 0x7c,
	0xf3, 0x2d, 0x8e, 0xee, 0x6f, 0xbf, 0x24, 0x3f, 0xbe, 0xde, 0x7d, 0xbf, 0xfe, 0x10, 0x7d, 0x8a,
	0xae, 0x3f, 0xba, 0x3b, 0xf8, 0x00, 0x3d, 0x8b, 0xef, 0xc2, 0x64, 0x76, 0x76, 0xf1, 0xce, 0xb5,
	0xc6, 0xe9, 0xe2, 0xec, 0xf2, 0xad, 0x6b, 0xcf, 0x1d, 0xb4, 0xd7, 0x76, 0xe9, 0x6f, 0xe0, 0x7a,
	0xfe, 0xfe, 0xe7, 0x55, 0x26, 0xf5, 0x43, 0x97, 0x12, 0xae, 0x4a, 0x3a, 0x74, 0x3d, 0x1d, 0xde,
	0x39, 0x53, 0xa7, 0x19, 0x54, 0xfd, 0x5e, 0xe8, 0x96, 0x0f, 0x70, 0x25, 0x59, 0x99, 0xee, 0xf6,
	0xf8, 0xcd, 0xbf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x41, 0xba, 0x65, 0x21, 0x22, 0x02, 0x00, 0x00,
}
