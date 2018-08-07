// Code generated by protoc-gen-go. DO NOT EDIT.
// source: common/update.proto

package common

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type ActionOpEnum int32

const (
	ActionOpEnum_insert ActionOpEnum = 0
	ActionOpEnum_update ActionOpEnum = 1
	ActionOpEnum_delete ActionOpEnum = 2
)

var ActionOpEnum_name = map[int32]string{
	0: "insert",
	1: "update",
	2: "delete",
}
var ActionOpEnum_value = map[string]int32{
	"insert": 0,
	"update": 1,
	"delete": 2,
}

func (x ActionOpEnum) String() string {
	return proto.EnumName(ActionOpEnum_name, int32(x))
}
func (ActionOpEnum) EnumDescriptor() ([]byte, []int) { return fileDescriptor5, []int{0} }

type Action struct {
	Op ActionOpEnum `protobuf:"varint,1,opt,name=op,enum=common.ActionOpEnum" json:"op,omitempty"`
}

func (m *Action) Reset()                    { *m = Action{} }
func (m *Action) String() string            { return proto.CompactTextString(m) }
func (*Action) ProtoMessage()               {}
func (*Action) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{0} }

func (m *Action) GetOp() ActionOpEnum {
	if m != nil {
		return m.Op
	}
	return ActionOpEnum_insert
}

type UpdateValue struct {
	DoubleValue *DoubleValue `protobuf:"bytes,1,opt,name=doubleValue" json:"doubleValue,omitempty"`
	FloatValue  *FloatValue  `protobuf:"bytes,2,opt,name=floatValue" json:"floatValue,omitempty"`
	Int64Value  *Int64Value  `protobuf:"bytes,3,opt,name=int64Value" json:"int64Value,omitempty"`
	UInt64Value *UInt64Value `protobuf:"bytes,4,opt,name=uInt64Value" json:"uInt64Value,omitempty"`
	Int32Value  *Int32Value  `protobuf:"bytes,5,opt,name=int32Value" json:"int32Value,omitempty"`
	UInt32Value *UInt32Value `protobuf:"bytes,6,opt,name=uInt32Value" json:"uInt32Value,omitempty"`
	BoolValue   *BoolValue   `protobuf:"bytes,7,opt,name=boolValue" json:"boolValue,omitempty"`
	StringValue *StringValue `protobuf:"bytes,8,opt,name=stringValue" json:"stringValue,omitempty"`
	BytesValue  *BytesValue  `protobuf:"bytes,9,opt,name=bytesValue" json:"bytesValue,omitempty"`
}

func (m *UpdateValue) Reset()                    { *m = UpdateValue{} }
func (m *UpdateValue) String() string            { return proto.CompactTextString(m) }
func (*UpdateValue) ProtoMessage()               {}
func (*UpdateValue) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{1} }

func (m *UpdateValue) GetDoubleValue() *DoubleValue {
	if m != nil {
		return m.DoubleValue
	}
	return nil
}

func (m *UpdateValue) GetFloatValue() *FloatValue {
	if m != nil {
		return m.FloatValue
	}
	return nil
}

func (m *UpdateValue) GetInt64Value() *Int64Value {
	if m != nil {
		return m.Int64Value
	}
	return nil
}

func (m *UpdateValue) GetUInt64Value() *UInt64Value {
	if m != nil {
		return m.UInt64Value
	}
	return nil
}

func (m *UpdateValue) GetInt32Value() *Int32Value {
	if m != nil {
		return m.Int32Value
	}
	return nil
}

func (m *UpdateValue) GetUInt32Value() *UInt32Value {
	if m != nil {
		return m.UInt32Value
	}
	return nil
}

func (m *UpdateValue) GetBoolValue() *BoolValue {
	if m != nil {
		return m.BoolValue
	}
	return nil
}

func (m *UpdateValue) GetStringValue() *StringValue {
	if m != nil {
		return m.StringValue
	}
	return nil
}

func (m *UpdateValue) GetBytesValue() *BytesValue {
	if m != nil {
		return m.BytesValue
	}
	return nil
}

func init() {
	proto.RegisterType((*Action)(nil), "common.Action")
	proto.RegisterType((*UpdateValue)(nil), "common.UpdateValue")
	proto.RegisterEnum("common.ActionOpEnum", ActionOpEnum_name, ActionOpEnum_value)
}

func init() { proto.RegisterFile("common/update.proto", fileDescriptor5) }

var fileDescriptor5 = []byte{
	// 344 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x92, 0x51, 0x4b, 0xeb, 0x30,
	0x14, 0xc7, 0xef, 0xba, 0x7b, 0x7b, 0x5d, 0x26, 0x32, 0x33, 0x85, 0xe1, 0x93, 0x0c, 0x1f, 0x44,
	0x58, 0x8b, 0xdd, 0xf4, 0xdd, 0xa1, 0x82, 0x4f, 0x42, 0x64, 0x3e, 0xf8, 0x96, 0xb6, 0xb1, 0x2b,
	0xb4, 0x39, 0xa1, 0x4d, 0x11, 0xbf, 0xb2, 0x9f, 0x42, 0xd2, 0xb3, 0xac, 0x29, 0x7b, 0x4b, 0xfb,
	0xfb, 0xff, 0xf2, 0x3f, 0x81, 0x43, 0xa6, 0x09, 0x94, 0x25, 0xc8, 0xb0, 0x51, 0x29, 0xd7, 0x22,
	0x50, 0x15, 0x68, 0xa0, 0x3e, 0xfe, 0xbc, 0x38, 0xdf, 0xc1, 0xaf, 0x8a, 0x2b, 0x25, 0xaa, 0x1a,
	0xf1, 0x3c, 0x20, 0xfe, 0x43, 0xa2, 0x73, 0x90, 0xf4, 0x8a, 0x78, 0xa0, 0x66, 0x83, 0xcb, 0xc1,
	0xf5, 0x49, 0x74, 0x16, 0x60, 0x3a, 0x40, 0xf6, 0xaa, 0x9e, 0x64, 0x53, 0x32, 0x0f, 0xd4, 0xfc,
	0x67, 0x48, 0xc6, 0x9b, 0xf6, 0xfe, 0x77, 0x5e, 0x34, 0x82, 0xde, 0x91, 0x71, 0x0a, 0x4d, 0x5c,
	0xe0, 0x67, 0xab, 0x8f, 0xa3, 0xa9, 0xd5, 0x1f, 0x3b, 0xc4, 0xdc, 0x1c, 0x8d, 0x08, 0xf9, 0x2c,
	0x80, 0x6b, 0xb4, 0xbc, 0xd6, 0xa2, 0xd6, 0x7a, 0xde, 0x13, 0xe6, 0xa4, 0x8c, 0x93, 0x4b, 0x7d,
	0xbf, 0x42, 0x67, 0xd8, 0x77, 0x5e, 0xf6, 0x84, 0x39, 0x29, 0x33, 0x5e, 0xd3, 0xa1, 0xd9, 0xdf,
	0xfe, 0x78, 0x1b, 0xc7, 0x72, 0x73, 0xbb, 0xaa, 0x65, 0x84, 0xd6, 0xbf, 0x83, 0xaa, 0x1d, 0x61,
	0x4e, 0xca, 0x56, 0x59, 0xc9, 0x3f, 0xac, 0xb2, 0x96, 0x9b, 0xa3, 0x21, 0x19, 0xc5, 0x00, 0x05,
	0x4a, 0xff, 0x5b, 0xe9, 0xd4, 0x4a, 0x6b, 0x0b, 0x58, 0x97, 0x31, 0x3d, 0xb5, 0xae, 0x72, 0x99,
	0xa1, 0x72, 0xd4, 0xef, 0x79, 0xeb, 0x10, 0x73, 0x73, 0xe6, 0x49, 0xf1, 0xb7, 0x16, 0x35, 0x5a,
	0xa3, 0xfe, 0x93, 0xd6, 0x7b, 0xc2, 0x9c, 0xd4, 0x4d, 0x44, 0x8e, 0xdd, 0x05, 0xa0, 0x84, 0xf8,
	0xb9, 0xac, 0x45, 0xa5, 0x27, 0x7f, 0xcc, 0x19, 0xf7, 0x6c, 0x32, 0x30, 0xe7, 0x54, 0x14, 0x42,
	0x8b, 0x89, 0xb7, 0x5e, 0x7d, 0x44, 0x59, 0xae, 0xb7, 0x4d, 0x6c, 0xee, 0x0e, 0xd5, 0x6d, 0x22,
	0x43, 0xcd, 0xa5, 0xe6, 0x72, 0x91, 0x42, 0xc9, 0x73, 0xb9, 0xa8, 0x93, 0xad, 0x28, 0x79, 0x98,
	0x41, 0xc1, 0x65, 0x16, 0xe2, 0x00, 0xb1, 0xdf, 0x6e, 0xe3, 0xf2, 0x37, 0x00, 0x00, 0xff, 0xff,
	0xc1, 0x7c, 0x04, 0x01, 0xc3, 0x02, 0x00, 0x00,
}
