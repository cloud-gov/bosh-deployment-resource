// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v1/resources/ad_parameter.proto

package resources

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
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

// An ad parameter that is used to update numeric values (such as prices or
// inventory levels) in any text line of an ad (including URLs). There can
// be a maximum of two AdParameters per ad group criterion. (One with
// parameter_index = 1 and one with parameter_index = 2.)
// In the ad the parameters are referenced by a placeholder of the form
// "{param#:value}". E.g. "{param1:$17}"
type AdParameter struct {
	// The resource name of the ad parameter.
	// Ad parameter resource names have the form:
	//
	// `customers/{customer_id}/adParameters/{ad_group_id}~{criterion_id}~{parameter_index}`
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	// The ad group criterion that this ad parameter belongs to.
	AdGroupCriterion *wrappers.StringValue `protobuf:"bytes,2,opt,name=ad_group_criterion,json=adGroupCriterion,proto3" json:"ad_group_criterion,omitempty"`
	// The unique index of this ad parameter. Must be either 1 or 2.
	ParameterIndex *wrappers.Int64Value `protobuf:"bytes,3,opt,name=parameter_index,json=parameterIndex,proto3" json:"parameter_index,omitempty"`
	// Numeric value to insert into the ad text. The following restrictions
	//  apply:
	//  - Can use comma or period as a separator, with an optional period or
	//    comma (respectively) for fractional values. For example, 1,000,000.00
	//    and 2.000.000,10 are valid.
	//  - Can be prepended or appended with a currency symbol. For example,
	//    $99.99 is valid.
	//  - Can be prepended or appended with a currency code. For example, 99.99USD
	//    and EUR200 are valid.
	//  - Can use '%'. For example, 1.0% and 1,0% are valid.
	//  - Can use plus or minus. For example, -10.99 and 25+ are valid.
	//  - Can use '/' between two numbers. For example 4/1 and 0.95/0.45 are
	//    valid.
	InsertionText        *wrappers.StringValue `protobuf:"bytes,4,opt,name=insertion_text,json=insertionText,proto3" json:"insertion_text,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *AdParameter) Reset()         { *m = AdParameter{} }
func (m *AdParameter) String() string { return proto.CompactTextString(m) }
func (*AdParameter) ProtoMessage()    {}
func (*AdParameter) Descriptor() ([]byte, []int) {
	return fileDescriptor_80d3f7afedc49596, []int{0}
}

func (m *AdParameter) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AdParameter.Unmarshal(m, b)
}
func (m *AdParameter) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AdParameter.Marshal(b, m, deterministic)
}
func (m *AdParameter) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AdParameter.Merge(m, src)
}
func (m *AdParameter) XXX_Size() int {
	return xxx_messageInfo_AdParameter.Size(m)
}
func (m *AdParameter) XXX_DiscardUnknown() {
	xxx_messageInfo_AdParameter.DiscardUnknown(m)
}

var xxx_messageInfo_AdParameter proto.InternalMessageInfo

func (m *AdParameter) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func (m *AdParameter) GetAdGroupCriterion() *wrappers.StringValue {
	if m != nil {
		return m.AdGroupCriterion
	}
	return nil
}

func (m *AdParameter) GetParameterIndex() *wrappers.Int64Value {
	if m != nil {
		return m.ParameterIndex
	}
	return nil
}

func (m *AdParameter) GetInsertionText() *wrappers.StringValue {
	if m != nil {
		return m.InsertionText
	}
	return nil
}

func init() {
	proto.RegisterType((*AdParameter)(nil), "google.ads.googleads.v1.resources.AdParameter")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v1/resources/ad_parameter.proto", fileDescriptor_80d3f7afedc49596)
}

var fileDescriptor_80d3f7afedc49596 = []byte{
	// 385 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x91, 0xdf, 0x6a, 0xdb, 0x30,
	0x18, 0xc5, 0xb1, 0x33, 0x06, 0x73, 0x96, 0x2c, 0xf8, 0xca, 0x64, 0x61, 0x24, 0x1b, 0x81, 0x5c,
	0xc9, 0x78, 0x0b, 0xbb, 0xd0, 0xae, 0x9c, 0x0c, 0x42, 0x72, 0x31, 0x42, 0x56, 0x7c, 0x51, 0x0c,
	0x46, 0x89, 0x54, 0x21, 0x88, 0x25, 0x23, 0xc9, 0x69, 0x5e, 0xa1, 0xaf, 0xd1, 0xcb, 0x3e, 0x4a,
	0x1f, 0xa5, 0xef, 0x50, 0x28, 0xfe, 0x23, 0x51, 0x28, 0xb4, 0xbd, 0x3b, 0x48, 0xe7, 0x77, 0xbe,
	0x23, 0x7d, 0xde, 0x9c, 0x0a, 0x41, 0x8f, 0x24, 0x44, 0x58, 0x85, 0x8d, 0xac, 0xd4, 0x29, 0x0a,
	0x25, 0x51, 0xa2, 0x94, 0x07, 0xa2, 0x42, 0x84, 0xb3, 0x02, 0x49, 0x94, 0x13, 0x4d, 0x24, 0x28,
	0xa4, 0xd0, 0xc2, 0x9f, 0x34, 0x56, 0x80, 0xb0, 0x02, 0x96, 0x02, 0xa7, 0x08, 0x58, 0x6a, 0xf8,
	0xad, 0x0d, 0xae, 0x81, 0x7d, 0x79, 0x15, 0x5e, 0x4b, 0x54, 0x14, 0x44, 0xaa, 0x26, 0x62, 0x38,
	0x32, 0x83, 0x0b, 0x16, 0x22, 0xce, 0x85, 0x46, 0x9a, 0x09, 0xde, 0xde, 0x7e, 0xbf, 0x71, 0xbd,
	0x6e, 0x8c, 0xb7, 0x66, 0xac, 0xff, 0xc3, 0xeb, 0x99, 0xe8, 0x8c, 0xa3, 0x9c, 0x04, 0xce, 0xd8,
	0x99, 0x7d, 0xda, 0x7d, 0x36, 0x87, 0xff, 0x50, 0x4e, 0xfc, 0x8d, 0xe7, 0x23, 0x9c, 0x51, 0x29,
	0xca, 0x22, 0x3b, 0x48, 0xa6, 0x89, 0x64, 0x82, 0x07, 0xee, 0xd8, 0x99, 0x75, 0x7f, 0x8e, 0xda,
	0x9e, 0xc0, 0xf4, 0x01, 0xff, 0xb5, 0x64, 0x9c, 0x26, 0xe8, 0x58, 0x92, 0xdd, 0x00, 0xe1, 0x55,
	0x85, 0x2d, 0x0d, 0xe5, 0xff, 0xf5, 0xbe, 0xd8, 0x47, 0x67, 0x8c, 0x63, 0x72, 0x0e, 0x3a, 0x75,
	0xd0, 0xd7, 0x17, 0x41, 0x6b, 0xae, 0x7f, 0xcf, 0x9b, 0x9c, 0xbe, 0x65, 0xd6, 0x15, 0xe2, 0x2f,
	0xbd, 0x3e, 0xe3, 0x8a, 0xc8, 0xea, 0x69, 0x99, 0x26, 0x67, 0x1d, 0x7c, 0x78, 0x47, 0x9b, 0x9e,
	0x65, 0x2e, 0xc8, 0x59, 0x2f, 0x1e, 0x1d, 0x6f, 0x7a, 0x10, 0x39, 0x78, 0xf3, 0xcf, 0x17, 0x83,
	0x67, 0x5f, 0xb6, 0xad, 0x92, 0xb7, 0xce, 0xe5, 0xa6, 0xc5, 0xa8, 0x38, 0x22, 0x4e, 0x81, 0x90,
	0x34, 0xa4, 0x84, 0xd7, 0x73, 0xcd, 0xc2, 0x0b, 0xa6, 0x5e, 0xd9, 0xff, 0x1f, 0xab, 0x6e, 0xdd,
	0xce, 0x2a, 0x8e, 0xef, 0xdc, 0xc9, 0xaa, 0x89, 0x8c, 0xb1, 0x02, 0x8d, 0xac, 0x54, 0x12, 0x81,
	0x9d, 0x71, 0xde, 0x1b, 0x4f, 0x1a, 0x63, 0x95, 0x5a, 0x4f, 0x9a, 0x44, 0xa9, 0xf5, 0x3c, 0xb8,
	0xd3, 0xe6, 0x02, 0xc2, 0x18, 0x2b, 0x08, 0xad, 0x0b, 0xc2, 0x24, 0x82, 0xd0, 0xfa, 0xf6, 0x1f,
	0xeb, 0xb2, 0xbf, 0x9e, 0x02, 0x00, 0x00, 0xff, 0xff, 0xe7, 0xad, 0x70, 0x5d, 0xab, 0x02, 0x00,
	0x00,
}
