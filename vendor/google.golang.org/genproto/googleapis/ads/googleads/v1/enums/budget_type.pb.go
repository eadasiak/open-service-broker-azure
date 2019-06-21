// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v1/enums/budget_type.proto

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

// Possible Budget types.
type BudgetTypeEnum_BudgetType int32

const (
	// Not specified.
	BudgetTypeEnum_UNSPECIFIED BudgetTypeEnum_BudgetType = 0
	// Used for return value only. Represents value unknown in this version.
	BudgetTypeEnum_UNKNOWN BudgetTypeEnum_BudgetType = 1
	// Budget type for standard Google Ads usage.
	// Caps daily spend at two times the specified budget amount.
	// Full details: https://support.google.com/google-ads/answer/6385083
	BudgetTypeEnum_STANDARD BudgetTypeEnum_BudgetType = 2
	// Budget type for Hotels Ads commission program.
	// Full details: https://support.google.com/google-ads/answer/9243945
	//
	// This type is only supported by campaigns with
	// AdvertisingChannelType.HOTEL, BiddingStrategyType.COMMISSION and
	// PaymentMode.CONVERSION_VALUE.
	BudgetTypeEnum_HOTEL_ADS_COMMISSION BudgetTypeEnum_BudgetType = 3
	// Budget type with a fixed cost-per-acquisition (conversion).
	// Full details: https://support.google.com/google-ads/answer/7528254
	//
	// This type is only supported by campaigns with
	// AdvertisingChannelType.DISPLAY (excluding
	// AdvertisingChannelSubType.DISPLAY_GMAIL),
	// BiddingStrategyType.TARGET_CPA and PaymentMode.CONVERSIONS.
	BudgetTypeEnum_FIXED_CPA BudgetTypeEnum_BudgetType = 4
)

var BudgetTypeEnum_BudgetType_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "UNKNOWN",
	2: "STANDARD",
	3: "HOTEL_ADS_COMMISSION",
	4: "FIXED_CPA",
}

var BudgetTypeEnum_BudgetType_value = map[string]int32{
	"UNSPECIFIED":          0,
	"UNKNOWN":              1,
	"STANDARD":             2,
	"HOTEL_ADS_COMMISSION": 3,
	"FIXED_CPA":            4,
}

func (x BudgetTypeEnum_BudgetType) String() string {
	return proto.EnumName(BudgetTypeEnum_BudgetType_name, int32(x))
}

func (BudgetTypeEnum_BudgetType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_064eaff51e649410, []int{0, 0}
}

// Describes Budget types.
type BudgetTypeEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BudgetTypeEnum) Reset()         { *m = BudgetTypeEnum{} }
func (m *BudgetTypeEnum) String() string { return proto.CompactTextString(m) }
func (*BudgetTypeEnum) ProtoMessage()    {}
func (*BudgetTypeEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_064eaff51e649410, []int{0}
}

func (m *BudgetTypeEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BudgetTypeEnum.Unmarshal(m, b)
}
func (m *BudgetTypeEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BudgetTypeEnum.Marshal(b, m, deterministic)
}
func (m *BudgetTypeEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BudgetTypeEnum.Merge(m, src)
}
func (m *BudgetTypeEnum) XXX_Size() int {
	return xxx_messageInfo_BudgetTypeEnum.Size(m)
}
func (m *BudgetTypeEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_BudgetTypeEnum.DiscardUnknown(m)
}

var xxx_messageInfo_BudgetTypeEnum proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("google.ads.googleads.v1.enums.BudgetTypeEnum_BudgetType", BudgetTypeEnum_BudgetType_name, BudgetTypeEnum_BudgetType_value)
	proto.RegisterType((*BudgetTypeEnum)(nil), "google.ads.googleads.v1.enums.BudgetTypeEnum")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v1/enums/budget_type.proto", fileDescriptor_064eaff51e649410)
}

var fileDescriptor_064eaff51e649410 = []byte{
	// 325 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x50, 0xdd, 0x4e, 0xfa, 0x30,
	0x1c, 0xfd, 0x33, 0xfe, 0xf1, 0xa3, 0xa8, 0x2c, 0x8b, 0x17, 0xc6, 0xc8, 0x05, 0x3c, 0x40, 0x97,
	0xc5, 0xbb, 0x7a, 0xd5, 0xb1, 0x81, 0x8b, 0xb2, 0x2d, 0x0e, 0xd0, 0x98, 0x25, 0x4b, 0xb1, 0x4d,
	0x43, 0x02, 0xed, 0x42, 0x07, 0x09, 0xaf, 0xe3, 0xa5, 0x8f, 0xe2, 0x73, 0x78, 0xe5, 0x53, 0x98,
	0x75, 0x30, 0xae, 0xf4, 0xa6, 0x39, 0xf9, 0x9d, 0x8f, 0x9c, 0x1e, 0x60, 0x73, 0x29, 0xf9, 0x82,
	0xd9, 0x84, 0xaa, 0x1d, 0x2c, 0xd1, 0xc6, 0xb1, 0x99, 0x58, 0x2f, 0x95, 0x3d, 0x5b, 0x53, 0xce,
	0x8a, 0xac, 0xd8, 0xe6, 0x0c, 0xe6, 0x2b, 0x59, 0x48, 0xab, 0x53, 0xa9, 0x20, 0xa1, 0x0a, 0xd6,
	0x06, 0xb8, 0x71, 0xa0, 0x36, 0x5c, 0xdf, 0xec, 0xf3, 0xf2, 0xb9, 0x4d, 0x84, 0x90, 0x05, 0x29,
	0xe6, 0x52, 0xa8, 0xca, 0xdc, 0x53, 0xe0, 0xc2, 0xd5, 0x89, 0xe3, 0x6d, 0xce, 0x7c, 0xb1, 0x5e,
	0xf6, 0x08, 0x00, 0x87, 0x8b, 0xd5, 0x06, 0xad, 0x49, 0x98, 0xc4, 0x7e, 0x3f, 0x18, 0x04, 0xbe,
	0x67, 0xfe, 0xb3, 0x5a, 0xe0, 0x78, 0x12, 0x3e, 0x84, 0xd1, 0x73, 0x68, 0x36, 0xac, 0x33, 0x70,
	0x92, 0x8c, 0x71, 0xe8, 0xe1, 0x27, 0xcf, 0x34, 0xac, 0x2b, 0x70, 0x79, 0x1f, 0x8d, 0xfd, 0xc7,
	0x0c, 0x7b, 0x49, 0xd6, 0x8f, 0x46, 0xa3, 0x20, 0x49, 0x82, 0x28, 0x34, 0x9b, 0xd6, 0x39, 0x38,
	0x1d, 0x04, 0x2f, 0xbe, 0x97, 0xf5, 0x63, 0x6c, 0xfe, 0x77, 0xbf, 0x1a, 0xa0, 0xfb, 0x26, 0x97,
	0xf0, 0xcf, 0xe2, 0x6e, 0xfb, 0x50, 0x23, 0x2e, 0xbb, 0xc6, 0x8d, 0x57, 0x77, 0xe7, 0xe0, 0x72,
	0x41, 0x04, 0x87, 0x72, 0xc5, 0x6d, 0xce, 0x84, 0xfe, 0xc9, 0x7e, 0xab, 0x7c, 0xae, 0x7e, 0x99,
	0xee, 0x4e, 0xbf, 0xef, 0x46, 0x73, 0x88, 0xf1, 0x87, 0xd1, 0x19, 0x56, 0x51, 0x98, 0x2a, 0x58,
	0xc1, 0x12, 0x4d, 0x1d, 0x58, 0x8e, 0xa0, 0x3e, 0xf7, 0x7c, 0x8a, 0xa9, 0x4a, 0x6b, 0x3e, 0x9d,
	0x3a, 0xa9, 0xe6, 0xbf, 0x8d, 0x6e, 0x75, 0x44, 0x08, 0x53, 0x85, 0x50, 0xad, 0x40, 0x68, 0xea,
	0x20, 0xa4, 0x35, 0xb3, 0x23, 0x5d, 0xec, 0xf6, 0x27, 0x00, 0x00, 0xff, 0xff, 0x9a, 0xec, 0x70,
	0xb2, 0xd2, 0x01, 0x00, 0x00,
}