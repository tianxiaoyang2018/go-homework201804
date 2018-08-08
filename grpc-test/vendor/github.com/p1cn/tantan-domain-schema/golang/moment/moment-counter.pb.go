// Code generated by protoc-gen-go. DO NOT EDIT.
// source: moment/moment-counter.proto

package moment

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type MomentCounter struct {
	UserId               string `protobuf:"bytes,1,opt,name=userId" json:"userId,omitempty"`
	Moments              int32  `protobuf:"varint,2,opt,name=moments" json:"moments,omitempty"`
	Followers            int32  `protobuf:"varint,3,opt,name=followers" json:"followers,omitempty"`
	Followings           int32  `protobuf:"varint,4,opt,name=followings" json:"followings,omitempty"`
	ReceiveMomentLikes   int32  `protobuf:"varint,5,opt,name=receiveMomentLikes" json:"receiveMomentLikes,omitempty"`
	UnReadFollowers      int32  `protobuf:"varint,6,opt,name=unReadFollowers" json:"unReadFollowers,omitempty"`
	UnReadMomentLikes    int32  `protobuf:"varint,7,opt,name=unReadMomentLikes" json:"unReadMomentLikes,omitempty"`
	UnReadMomentComments int32  `protobuf:"varint,8,opt,name=unReadMomentComments" json:"unReadMomentComments,omitempty"`
	CreateTime           int64  `protobuf:"varint,9,opt,name=createTime" json:"createTime,omitempty"`
	UpdateTime           int64  `protobuf:"varint,10,opt,name=updateTime" json:"updateTime,omitempty"`
}

func (m *MomentCounter) Reset()                    { *m = MomentCounter{} }
func (m *MomentCounter) String() string            { return proto.CompactTextString(m) }
func (*MomentCounter) ProtoMessage()               {}
func (*MomentCounter) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{0} }

func (m *MomentCounter) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func (m *MomentCounter) GetMoments() int32 {
	if m != nil {
		return m.Moments
	}
	return 0
}

func (m *MomentCounter) GetFollowers() int32 {
	if m != nil {
		return m.Followers
	}
	return 0
}

func (m *MomentCounter) GetFollowings() int32 {
	if m != nil {
		return m.Followings
	}
	return 0
}

func (m *MomentCounter) GetReceiveMomentLikes() int32 {
	if m != nil {
		return m.ReceiveMomentLikes
	}
	return 0
}

func (m *MomentCounter) GetUnReadFollowers() int32 {
	if m != nil {
		return m.UnReadFollowers
	}
	return 0
}

func (m *MomentCounter) GetUnReadMomentLikes() int32 {
	if m != nil {
		return m.UnReadMomentLikes
	}
	return 0
}

func (m *MomentCounter) GetUnReadMomentComments() int32 {
	if m != nil {
		return m.UnReadMomentComments
	}
	return 0
}

func (m *MomentCounter) GetCreateTime() int64 {
	if m != nil {
		return m.CreateTime
	}
	return 0
}

func (m *MomentCounter) GetUpdateTime() int64 {
	if m != nil {
		return m.UpdateTime
	}
	return 0
}

func init() {
	proto.RegisterType((*MomentCounter)(nil), "moment.MomentCounter")
}

func init() { proto.RegisterFile("moment/moment-counter.proto", fileDescriptor5) }

var fileDescriptor5 = []byte{
	// 278 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x91, 0x5f, 0x4b, 0xf3, 0x30,
	0x14, 0x87, 0xe9, 0xf6, 0xae, 0x7b, 0x7b, 0x40, 0xc4, 0x20, 0x12, 0x50, 0xa4, 0x78, 0xd5, 0x0b,
	0xdb, 0xe2, 0xf4, 0x13, 0x38, 0x10, 0x04, 0xbd, 0x29, 0x5e, 0x79, 0x97, 0xa5, 0xc7, 0x2e, 0xd8,
	0x24, 0x25, 0x7f, 0xf4, 0xc3, 0x7b, 0x23, 0x4b, 0x36, 0x57, 0x74, 0x57, 0xed, 0x79, 0x9e, 0xdf,
	0x49, 0x72, 0x38, 0x70, 0x2e, 0xb5, 0x44, 0xe5, 0xea, 0xf8, 0x29, 0xb9, 0xf6, 0xca, 0xa1, 0xa9,
	0x06, 0xa3, 0x9d, 0x26, 0x69, 0xa4, 0x57, 0x5f, 0x13, 0x38, 0x7a, 0x0e, 0xbf, 0xcb, 0xe8, 0xc9,
	0x19, 0xa4, 0xde, 0xa2, 0x79, 0x6c, 0x69, 0x92, 0x27, 0x45, 0xd6, 0x6c, 0x2b, 0x42, 0x61, 0x1e,
	0x7b, 0x2c, 0x9d, 0xe4, 0x49, 0x31, 0x6b, 0x76, 0x25, 0xb9, 0x80, 0xec, 0x4d, 0xf7, 0xbd, 0xfe,
	0x44, 0x63, 0xe9, 0x34, 0xb8, 0x3d, 0x20, 0x97, 0x00, 0xb1, 0x10, 0xaa, 0xb3, 0xf4, 0x5f, 0xd0,
	0x23, 0x42, 0x2a, 0x20, 0x06, 0x39, 0x8a, 0x0f, 0x8c, 0xef, 0x78, 0x12, 0xef, 0x68, 0xe9, 0x2c,
	0xe4, 0x0e, 0x18, 0x52, 0xc0, 0xb1, 0x57, 0x0d, 0xb2, 0xf6, 0xe1, 0xe7, 0xce, 0x34, 0x84, 0x7f,
	0x63, 0x72, 0x0d, 0x27, 0x11, 0x8d, 0x0f, 0x9e, 0x87, 0xec, 0x5f, 0x41, 0x16, 0x70, 0x3a, 0x86,
	0x4b, 0x2d, 0xe3, 0xb0, 0xff, 0x43, 0xc3, 0x41, 0xb7, 0x99, 0x8d, 0x1b, 0x64, 0x0e, 0x5f, 0x84,
	0x44, 0x9a, 0xe5, 0x49, 0x31, 0x6d, 0x46, 0x64, 0xe3, 0xfd, 0xd0, 0xee, 0x3c, 0x44, 0xbf, 0x27,
	0xf7, 0x77, 0xaf, 0x8b, 0x4e, 0xb8, 0xb5, 0x5f, 0x55, 0x5c, 0xcb, 0x7a, 0xb8, 0xe1, 0xaa, 0x76,
	0x4c, 0x39, 0xa6, 0xca, 0x56, 0x4b, 0x26, 0x54, 0x69, 0xf9, 0x1a, 0x25, 0xab, 0x3b, 0xdd, 0x33,
	0xd5, 0x6d, 0x37, 0xb9, 0x4a, 0xc3, 0x0a, 0x6f, 0xbf, 0x03, 0x00, 0x00, 0xff, 0xff, 0x29, 0x82,
	0xd6, 0xfe, 0xe1, 0x01, 0x00, 0x00,
}