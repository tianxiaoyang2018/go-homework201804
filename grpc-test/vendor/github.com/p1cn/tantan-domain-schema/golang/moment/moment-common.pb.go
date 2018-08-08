// Code generated by protoc-gen-go. DO NOT EDIT.
// source: moment/moment-common.proto

package moment

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type MomentStatus int32

const (
	MomentStatus_MOMENT_STATUS_DEFAULT     MomentStatus = 0
	MomentStatus_MOMENT_STATUS_DELETED     MomentStatus = 1
	MomentStatus_MOMENT_STATUS_HIDDEN      MomentStatus = 2
	MomentStatus_MOMENT_STATUS_INACTIVATED MomentStatus = 3
)

var MomentStatus_name = map[int32]string{
	0: "MOMENT_STATUS_DEFAULT",
	1: "MOMENT_STATUS_DELETED",
	2: "MOMENT_STATUS_HIDDEN",
	3: "MOMENT_STATUS_INACTIVATED",
}
var MomentStatus_value = map[string]int32{
	"MOMENT_STATUS_DEFAULT":     0,
	"MOMENT_STATUS_DELETED":     1,
	"MOMENT_STATUS_HIDDEN":      2,
	"MOMENT_STATUS_INACTIVATED": 3,
}

func (x MomentStatus) String() string {
	return proto.EnumName(MomentStatus_name, int32(x))
}
func (MomentStatus) EnumDescriptor() ([]byte, []int) { return fileDescriptor4, []int{0} }

func init() {
	proto.RegisterEnum("moment.MomentStatus", MomentStatus_name, MomentStatus_value)
}

func init() { proto.RegisterFile("moment/moment-common.proto", fileDescriptor4) }

var fileDescriptor4 = []byte{
	// 186 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0xca, 0xcd, 0xcf, 0x4d,
	0xcd, 0x2b, 0xd1, 0x87, 0x50, 0xba, 0xc9, 0xf9, 0xb9, 0xb9, 0xf9, 0x79, 0x7a, 0x05, 0x45, 0xf9,
	0x25, 0xf9, 0x42, 0x6c, 0x10, 0x41, 0xad, 0x5a, 0x2e, 0x1e, 0x5f, 0x30, 0x2b, 0xb8, 0x24, 0xb1,
	0xa4, 0xb4, 0x58, 0x48, 0x92, 0x4b, 0xd4, 0xd7, 0xdf, 0xd7, 0xd5, 0x2f, 0x24, 0x3e, 0x38, 0xc4,
	0x31, 0x24, 0x34, 0x38, 0xde, 0xc5, 0xd5, 0xcd, 0x31, 0xd4, 0x27, 0x44, 0x80, 0x01, 0x9b, 0x94,
	0x8f, 0x6b, 0x88, 0xab, 0x8b, 0x00, 0xa3, 0x90, 0x04, 0x97, 0x08, 0xaa, 0x94, 0x87, 0xa7, 0x8b,
	0x8b, 0xab, 0x9f, 0x00, 0x93, 0x90, 0x2c, 0x97, 0x24, 0xaa, 0x8c, 0xa7, 0x9f, 0xa3, 0x73, 0x88,
	0x67, 0x98, 0x23, 0x48, 0x23, 0xb3, 0x93, 0x49, 0x94, 0x51, 0x7a, 0x66, 0x49, 0x46, 0x69, 0x92,
	0x5e, 0x72, 0x7e, 0xae, 0x7e, 0x81, 0x61, 0x72, 0x9e, 0x7e, 0x49, 0x62, 0x5e, 0x49, 0x62, 0x9e,
	0x6e, 0x4a, 0x7e, 0x6e, 0x62, 0x66, 0x9e, 0x6e, 0x71, 0x72, 0x46, 0x6a, 0x6e, 0xa2, 0x7e, 0x7a,
	0x7e, 0x4e, 0x62, 0x5e, 0x3a, 0xd4, 0x27, 0x49, 0x6c, 0x60, 0x3f, 0x18, 0x03, 0x02, 0x00, 0x00,
	0xff, 0xff, 0xe2, 0x50, 0xe7, 0xe8, 0xe1, 0x00, 0x00, 0x00,
}