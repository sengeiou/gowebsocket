// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v2/enums/feed_status.proto

package enums

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

// Possible statuses of a feed.
type FeedStatusEnum_FeedStatus int32

const (
	// Not specified.
	FeedStatusEnum_UNSPECIFIED FeedStatusEnum_FeedStatus = 0
	// Used for return value only. Represents value unknown in this version.
	FeedStatusEnum_UNKNOWN FeedStatusEnum_FeedStatus = 1
	// Feed is enabled.
	FeedStatusEnum_ENABLED FeedStatusEnum_FeedStatus = 2
	// Feed has been removed.
	FeedStatusEnum_REMOVED FeedStatusEnum_FeedStatus = 3
)

var FeedStatusEnum_FeedStatus_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "UNKNOWN",
	2: "ENABLED",
	3: "REMOVED",
}

var FeedStatusEnum_FeedStatus_value = map[string]int32{
	"UNSPECIFIED": 0,
	"UNKNOWN":     1,
	"ENABLED":     2,
	"REMOVED":     3,
}

func (x FeedStatusEnum_FeedStatus) String() string {
	return proto.EnumName(FeedStatusEnum_FeedStatus_name, int32(x))
}

func (FeedStatusEnum_FeedStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_ef5097fa2e289d1d, []int{0, 0}
}

// Container for enum describing possible statuses of a feed.
type FeedStatusEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FeedStatusEnum) Reset()         { *m = FeedStatusEnum{} }
func (m *FeedStatusEnum) String() string { return proto.CompactTextString(m) }
func (*FeedStatusEnum) ProtoMessage()    {}
func (*FeedStatusEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_ef5097fa2e289d1d, []int{0}
}

func (m *FeedStatusEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FeedStatusEnum.Unmarshal(m, b)
}
func (m *FeedStatusEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FeedStatusEnum.Marshal(b, m, deterministic)
}
func (m *FeedStatusEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FeedStatusEnum.Merge(m, src)
}
func (m *FeedStatusEnum) XXX_Size() int {
	return xxx_messageInfo_FeedStatusEnum.Size(m)
}
func (m *FeedStatusEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_FeedStatusEnum.DiscardUnknown(m)
}

var xxx_messageInfo_FeedStatusEnum proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("google.ads.googleads.v2.enums.FeedStatusEnum_FeedStatus", FeedStatusEnum_FeedStatus_name, FeedStatusEnum_FeedStatus_value)
	proto.RegisterType((*FeedStatusEnum)(nil), "google.ads.googleads.v2.enums.FeedStatusEnum")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v2/enums/feed_status.proto", fileDescriptor_ef5097fa2e289d1d)
}

var fileDescriptor_ef5097fa2e289d1d = []byte{
	// 293 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x50, 0xdd, 0x4a, 0xf3, 0x30,
	0x18, 0xfe, 0xd6, 0xc1, 0x27, 0x64, 0xe0, 0x4a, 0x0f, 0xc5, 0x1d, 0x6c, 0x17, 0x90, 0x40, 0x3d,
	0x8b, 0x47, 0xa9, 0xcd, 0xc6, 0x50, 0xbb, 0xe2, 0x58, 0x05, 0x29, 0x48, 0x34, 0x31, 0x14, 0xd6,
	0xa4, 0x2c, 0xed, 0x2e, 0xc8, 0x43, 0x2f, 0xc5, 0xeb, 0xf0, 0xc8, 0xab, 0x90, 0xa4, 0x3f, 0x3b,
	0xd2, 0x93, 0xf2, 0xbc, 0xef, 0xf3, 0xd3, 0x27, 0x2f, 0x40, 0x52, 0x6b, 0xb9, 0x17, 0x88, 0x71,
	0xd3, 0x41, 0x8b, 0x8e, 0x21, 0x12, 0xaa, 0x29, 0x0d, 0x7a, 0x13, 0x82, 0x3f, 0x9b, 0x9a, 0xd5,
	0x8d, 0x81, 0xd5, 0x41, 0xd7, 0x3a, 0x98, 0xb5, 0x2a, 0xc8, 0xb8, 0x81, 0x83, 0x01, 0x1e, 0x43,
	0xe8, 0x0c, 0x17, 0x97, 0x7d, 0x5e, 0x55, 0x20, 0xa6, 0x94, 0xae, 0x59, 0x5d, 0x68, 0xd5, 0x99,
	0x17, 0x19, 0x38, 0x5f, 0x0a, 0xc1, 0xb7, 0x2e, 0x90, 0xaa, 0xa6, 0x5c, 0xc4, 0x00, 0x9c, 0x36,
	0xc1, 0x14, 0x4c, 0x76, 0xc9, 0x36, 0xa5, 0x37, 0xeb, 0xe5, 0x9a, 0xc6, 0xfe, 0xbf, 0x60, 0x02,
	0xce, 0x76, 0xc9, 0x6d, 0xb2, 0x79, 0x4c, 0xfc, 0x91, 0x1d, 0x68, 0x42, 0xa2, 0x3b, 0x1a, 0xfb,
	0x9e, 0x1d, 0x1e, 0xe8, 0xfd, 0x26, 0xa3, 0xb1, 0x3f, 0x8e, 0xbe, 0x46, 0x60, 0xfe, 0xaa, 0x4b,
	0xf8, 0x67, 0xb7, 0x68, 0x7a, 0xfa, 0x53, 0x6a, 0xeb, 0xa4, 0xa3, 0xa7, 0xa8, 0x73, 0x48, 0xbd,
	0x67, 0x4a, 0x42, 0x7d, 0x90, 0x48, 0x0a, 0xe5, 0xca, 0xf6, 0xe7, 0xa8, 0x0a, 0xf3, 0xcb, 0x75,
	0xae, 0xdd, 0xf7, 0xdd, 0x1b, 0xaf, 0x08, 0xf9, 0xf0, 0x66, 0xab, 0x36, 0x8a, 0x70, 0x03, 0x5b,
	0x68, 0x51, 0x16, 0x42, 0xfb, 0x4e, 0xf3, 0xd9, 0xf3, 0x39, 0xe1, 0x26, 0x1f, 0xf8, 0x3c, 0x0b,
	0x73, 0xc7, 0x7f, 0x7b, 0xf3, 0x76, 0x89, 0x31, 0xe1, 0x06, 0xe3, 0x41, 0x81, 0x71, 0x16, 0x62,
	0xec, 0x34, 0x2f, 0xff, 0x5d, 0xb1, 0xab, 0x9f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x4a, 0x3c, 0x03,
	0x66, 0xb5, 0x01, 0x00, 0x00,
}
