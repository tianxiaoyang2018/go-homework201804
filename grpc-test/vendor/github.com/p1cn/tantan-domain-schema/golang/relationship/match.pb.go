// Code generated by protoc-gen-go. DO NOT EDIT.
// source: relationship/match.proto

package relationship

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type Match struct {
	UserId        string `protobuf:"bytes,1,opt,name=userId" json:"userId,omitempty"`
	OtherUserId   string `protobuf:"bytes,2,opt,name=otherUserId" json:"otherUserId,omitempty"`
	CreatedTime   int64  `protobuf:"varint,3,opt,name=createdTime" json:"createdTime,omitempty"`
	IsSecretCrush bool   `protobuf:"varint,4,opt,name=isSecretCrush" json:"isSecretCrush,omitempty"`
}

func (m *Match) Reset()                    { *m = Match{} }
func (m *Match) String() string            { return proto.CompactTextString(m) }
func (*Match) ProtoMessage()               {}
func (*Match) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

func (m *Match) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func (m *Match) GetOtherUserId() string {
	if m != nil {
		return m.OtherUserId
	}
	return ""
}

func (m *Match) GetCreatedTime() int64 {
	if m != nil {
		return m.CreatedTime
	}
	return 0
}

func (m *Match) GetIsSecretCrush() bool {
	if m != nil {
		return m.IsSecretCrush
	}
	return false
}

func init() {
	proto.RegisterType((*Match)(nil), "relationship.Match")
}

func init() { proto.RegisterFile("relationship/match.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
	// 198 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0xcf, 0xa1, 0x4f, 0xc5, 0x30,
	0x10, 0xc7, 0xf1, 0x94, 0xc1, 0x02, 0x05, 0x4c, 0x05, 0xa9, 0x6c, 0x08, 0x62, 0x66, 0x6b, 0x08,
	0x86, 0x20, 0x41, 0x21, 0x30, 0x03, 0x0c, 0xee, 0xd6, 0x5d, 0xd6, 0x26, 0x6b, 0xbb, 0xb4, 0xb7,
	0x3f, 0x82, 0xff, 0xfa, 0x65, 0x7b, 0x4f, 0xf4, 0xc9, 0xfb, 0x7e, 0xce, 0xfc, 0xb8, 0x4c, 0x38,
	0x03, 0xb9, 0x18, 0xb2, 0x75, 0x8b, 0xf6, 0x40, 0xc6, 0x76, 0x4b, 0x8a, 0x14, 0xc5, 0x5d, 0x29,
	0x8f, 0xff, 0x8c, 0x5f, 0x7d, 0x6d, 0x2a, 0x1e, 0x78, 0xbd, 0x66, 0x4c, 0x9f, 0xa3, 0x64, 0x8a,
	0x35, 0x37, 0xfd, 0xe9, 0x12, 0x8a, 0xdf, 0x46, 0xb2, 0x98, 0x7e, 0x8f, 0x78, 0xb1, 0x63, 0x99,
	0xb6, 0x0f, 0x93, 0x10, 0x08, 0xc7, 0x1f, 0xe7, 0x51, 0x56, 0x8a, 0x35, 0x55, 0x5f, 0x26, 0xf1,
	0xc4, 0xef, 0x5d, 0xfe, 0x46, 0x93, 0x90, 0x3e, 0xd2, 0x9a, 0xad, 0xbc, 0x54, 0xac, 0xb9, 0xee,
	0xcf, 0xe3, 0xfb, 0xdb, 0xdf, 0xeb, 0xe4, 0xc8, 0xae, 0x43, 0x67, 0xa2, 0xd7, 0xcb, 0xb3, 0x09,
	0x9a, 0x20, 0x10, 0x84, 0x76, 0x8c, 0x1e, 0x5c, 0x68, 0xb3, 0xb1, 0xe8, 0x41, 0x4f, 0x71, 0x86,
	0x30, 0xe9, 0x72, 0xc7, 0x50, 0xef, 0xe3, 0x5e, 0x0e, 0x01, 0x00, 0x00, 0xff, 0xff, 0x35, 0xfb,
	0xa9, 0xde, 0xf8, 0x00, 0x00, 0x00,
}