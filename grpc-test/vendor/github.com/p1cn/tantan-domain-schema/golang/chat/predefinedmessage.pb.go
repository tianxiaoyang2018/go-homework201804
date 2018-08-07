// Code generated by protoc-gen-go. DO NOT EDIT.
// source: chat/predefinedmessage.proto

package chat

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type PredefinedMessage struct {
	Id                string  `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Key               string  `protobuf:"bytes,2,opt,name=key" json:"key,omitempty"`
	Language          string  `protobuf:"bytes,3,opt,name=language" json:"language,omitempty"`
	Text              string  `protobuf:"bytes,4,opt,name=text" json:"text,omitempty"`
	PictureIdentifier string  `protobuf:"bytes,5,opt,name=pictureIdentifier" json:"pictureIdentifier,omitempty"`
	PictureWidth      int64   `protobuf:"varint,6,opt,name=pictureWidth" json:"pictureWidth,omitempty"`
	PictureHeight     int64   `protobuf:"varint,7,opt,name=pictureHeight" json:"pictureHeight,omitempty"`
	VideoIdentifier   string  `protobuf:"bytes,8,opt,name=videoIdentifier" json:"videoIdentifier,omitempty"`
	VideoWidth        int64   `protobuf:"varint,9,opt,name=videoWidth" json:"videoWidth,omitempty"`
	VideoHeight       int64   `protobuf:"varint,10,opt,name=videoHeight" json:"videoHeight,omitempty"`
	VideoDuration     float32 `protobuf:"fixed32,11,opt,name=videoDuration" json:"videoDuration,omitempty"`
	AudioIdentifier   string  `protobuf:"bytes,12,opt,name=audioIdentifier" json:"audioIdentifier,omitempty"`
	AudioDuration     float32 `protobuf:"fixed32,13,opt,name=audioDuration" json:"audioDuration,omitempty"`
	CreatedTime       int64   `protobuf:"varint,14,opt,name=CreatedTime" json:"CreatedTime,omitempty"`
	UpdatedTime       int64   `protobuf:"varint,15,opt,name=UpdatedTime" json:"UpdatedTime,omitempty"`
}

func (m *PredefinedMessage) Reset()                    { *m = PredefinedMessage{} }
func (m *PredefinedMessage) String() string            { return proto.CompactTextString(m) }
func (*PredefinedMessage) ProtoMessage()               {}
func (*PredefinedMessage) Descriptor() ([]byte, []int) { return fileDescriptor6, []int{0} }

func (m *PredefinedMessage) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *PredefinedMessage) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *PredefinedMessage) GetLanguage() string {
	if m != nil {
		return m.Language
	}
	return ""
}

func (m *PredefinedMessage) GetText() string {
	if m != nil {
		return m.Text
	}
	return ""
}

func (m *PredefinedMessage) GetPictureIdentifier() string {
	if m != nil {
		return m.PictureIdentifier
	}
	return ""
}

func (m *PredefinedMessage) GetPictureWidth() int64 {
	if m != nil {
		return m.PictureWidth
	}
	return 0
}

func (m *PredefinedMessage) GetPictureHeight() int64 {
	if m != nil {
		return m.PictureHeight
	}
	return 0
}

func (m *PredefinedMessage) GetVideoIdentifier() string {
	if m != nil {
		return m.VideoIdentifier
	}
	return ""
}

func (m *PredefinedMessage) GetVideoWidth() int64 {
	if m != nil {
		return m.VideoWidth
	}
	return 0
}

func (m *PredefinedMessage) GetVideoHeight() int64 {
	if m != nil {
		return m.VideoHeight
	}
	return 0
}

func (m *PredefinedMessage) GetVideoDuration() float32 {
	if m != nil {
		return m.VideoDuration
	}
	return 0
}

func (m *PredefinedMessage) GetAudioIdentifier() string {
	if m != nil {
		return m.AudioIdentifier
	}
	return ""
}

func (m *PredefinedMessage) GetAudioDuration() float32 {
	if m != nil {
		return m.AudioDuration
	}
	return 0
}

func (m *PredefinedMessage) GetCreatedTime() int64 {
	if m != nil {
		return m.CreatedTime
	}
	return 0
}

func (m *PredefinedMessage) GetUpdatedTime() int64 {
	if m != nil {
		return m.UpdatedTime
	}
	return 0
}

func init() {
	proto.RegisterType((*PredefinedMessage)(nil), "chat.PredefinedMessage")
}

func init() { proto.RegisterFile("chat/predefinedmessage.proto", fileDescriptor6) }

var fileDescriptor6 = []byte{
	// 336 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x92, 0xcd, 0x4a, 0xf3, 0x40,
	0x14, 0x86, 0x49, 0x9a, 0xaf, 0x5f, 0x3b, 0xfd, 0xb3, 0xb3, 0x1a, 0x44, 0x24, 0x14, 0x17, 0x59,
	0xd8, 0xc6, 0x9f, 0x3b, 0x50, 0x17, 0xba, 0x10, 0xa4, 0x28, 0x82, 0xbb, 0x69, 0xe6, 0x34, 0x39,
	0x68, 0x66, 0x42, 0x3a, 0x11, 0x5d, 0x7b, 0xe3, 0x32, 0x27, 0x25, 0x24, 0x75, 0x37, 0xe7, 0x99,
	0x87, 0xf3, 0xbe, 0x30, 0xc3, 0x4e, 0x92, 0x4c, 0xda, 0xb8, 0x28, 0x41, 0xc1, 0x16, 0x35, 0xa8,
	0x1c, 0x76, 0x3b, 0x99, 0xc2, 0xaa, 0x28, 0x8d, 0x35, 0x3c, 0x70, 0xb7, 0x8b, 0x9f, 0x80, 0xcd,
	0x9f, 0x1a, 0xe3, 0xb1, 0x36, 0xf8, 0x94, 0xf9, 0xa8, 0x84, 0x17, 0x7a, 0xd1, 0x70, 0xed, 0xa3,
	0xe2, 0x47, 0xac, 0xf7, 0x0e, 0xdf, 0xc2, 0x27, 0xe0, 0x8e, 0xfc, 0x98, 0x0d, 0x3e, 0xa4, 0x4e,
	0x2b, 0x99, 0x82, 0xe8, 0x11, 0x6e, 0x66, 0xce, 0x59, 0x60, 0xe1, 0xcb, 0x8a, 0x80, 0x38, 0x9d,
	0xf9, 0x39, 0x9b, 0x17, 0x98, 0xd8, 0xaa, 0x84, 0x07, 0x05, 0xda, 0xe2, 0x16, 0xa1, 0x14, 0xff,
	0x48, 0xf8, 0x7b, 0xc1, 0x17, 0x6c, 0xbc, 0x87, 0xaf, 0xa8, 0x6c, 0x26, 0xfa, 0xa1, 0x17, 0xf5,
	0xd6, 0x1d, 0xc6, 0xcf, 0xd8, 0x64, 0x3f, 0xdf, 0x03, 0xa6, 0x99, 0x15, 0xff, 0x49, 0xea, 0x42,
	0x1e, 0xb1, 0xd9, 0x27, 0x2a, 0x30, 0xad, 0xd4, 0x01, 0xa5, 0x1e, 0x62, 0x7e, 0xca, 0x18, 0xa1,
	0x3a, 0x71, 0x48, 0xcb, 0x5a, 0x84, 0x87, 0x6c, 0x44, 0xd3, 0x3e, 0x8d, 0x91, 0xd0, 0x46, 0xae,
	0x11, 0x8d, 0x77, 0x55, 0x29, 0x2d, 0x1a, 0x2d, 0x46, 0xa1, 0x17, 0xf9, 0xeb, 0x2e, 0x74, 0x8d,
	0x64, 0xa5, 0xb0, 0xdd, 0x68, 0x5c, 0x37, 0x3a, 0xc0, 0x6e, 0x1f, 0xa1, 0x66, 0xdf, 0xa4, 0xde,
	0xd7, 0x81, 0xae, 0xd7, 0x6d, 0x09, 0xd2, 0x82, 0x7a, 0xc6, 0x1c, 0xc4, 0xb4, 0xee, 0xd5, 0x42,
	0xce, 0x78, 0x29, 0x54, 0x63, 0xcc, 0x6a, 0xa3, 0x85, 0x6e, 0xae, 0xde, 0x2e, 0x52, 0xb4, 0x59,
	0xb5, 0x59, 0x25, 0x26, 0x8f, 0x8b, 0xcb, 0x44, 0xc7, 0x56, 0x6a, 0x2b, 0xf5, 0x52, 0x99, 0x5c,
	0xa2, 0x5e, 0xee, 0x92, 0x0c, 0x72, 0x19, 0xa7, 0xc6, 0x3d, 0x72, 0xec, 0x7e, 0xce, 0xa6, 0x4f,
	0xdf, 0xe8, 0xfa, 0x37, 0x00, 0x00, 0xff, 0xff, 0xb6, 0x60, 0xfd, 0xfb, 0x66, 0x02, 0x00, 0x00,
}
