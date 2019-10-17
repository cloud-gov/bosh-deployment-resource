// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v1/services/shopping_performance_view_service.proto

package services

import (
	context "context"
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	resources "google.golang.org/genproto/googleapis/ads/googleads/v1/resources"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// Request message for
// [ShoppingPerformanceViewService.GetShoppingPerformanceView][google.ads.googleads.v1.services.ShoppingPerformanceViewService.GetShoppingPerformanceView].
type GetShoppingPerformanceViewRequest struct {
	// The resource name of the Shopping performance view to fetch.
	ResourceName         string   `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetShoppingPerformanceViewRequest) Reset()         { *m = GetShoppingPerformanceViewRequest{} }
func (m *GetShoppingPerformanceViewRequest) String() string { return proto.CompactTextString(m) }
func (*GetShoppingPerformanceViewRequest) ProtoMessage()    {}
func (*GetShoppingPerformanceViewRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_4d6c4cf22050d6d3, []int{0}
}

func (m *GetShoppingPerformanceViewRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetShoppingPerformanceViewRequest.Unmarshal(m, b)
}
func (m *GetShoppingPerformanceViewRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetShoppingPerformanceViewRequest.Marshal(b, m, deterministic)
}
func (m *GetShoppingPerformanceViewRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetShoppingPerformanceViewRequest.Merge(m, src)
}
func (m *GetShoppingPerformanceViewRequest) XXX_Size() int {
	return xxx_messageInfo_GetShoppingPerformanceViewRequest.Size(m)
}
func (m *GetShoppingPerformanceViewRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetShoppingPerformanceViewRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetShoppingPerformanceViewRequest proto.InternalMessageInfo

func (m *GetShoppingPerformanceViewRequest) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func init() {
	proto.RegisterType((*GetShoppingPerformanceViewRequest)(nil), "google.ads.googleads.v1.services.GetShoppingPerformanceViewRequest")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v1/services/shopping_performance_view_service.proto", fileDescriptor_4d6c4cf22050d6d3)
}

var fileDescriptor_4d6c4cf22050d6d3 = []byte{
	// 378 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x92, 0xb1, 0x4b, 0xe4, 0x40,
	0x14, 0xc6, 0x49, 0x0e, 0x0e, 0x2e, 0xdc, 0x35, 0xa9, 0x8e, 0x70, 0x1c, 0x7b, 0xbb, 0x5b, 0x1c,
	0x57, 0x4c, 0xc8, 0x59, 0x88, 0x23, 0x22, 0x59, 0x8b, 0xdd, 0x4a, 0x96, 0x5d, 0x48, 0x21, 0x81,
	0x30, 0x26, 0xcf, 0x18, 0xd8, 0xcc, 0xc4, 0x79, 0xd9, 0x6c, 0x21, 0x16, 0x5a, 0xda, 0xfa, 0x1f,
	0x58, 0xfa, 0xa7, 0xd8, 0xda, 0x5b, 0x59, 0xf9, 0x57, 0x48, 0x76, 0x32, 0x11, 0x85, 0xec, 0x76,
	0x1f, 0xc9, 0x37, 0xbf, 0x6f, 0xde, 0x37, 0xcf, 0x9a, 0xa4, 0x42, 0xa4, 0x0b, 0x70, 0x59, 0x82,
	0xae, 0x92, 0xb5, 0xaa, 0x3c, 0x17, 0x41, 0x56, 0x59, 0x0c, 0xe8, 0xe2, 0xb9, 0x28, 0x8a, 0x8c,
	0xa7, 0x51, 0x01, 0xf2, 0x4c, 0xc8, 0x9c, 0xf1, 0x18, 0xa2, 0x2a, 0x83, 0x55, 0xd4, 0x58, 0x48,
	0x21, 0x45, 0x29, 0xec, 0x9e, 0x3a, 0x4e, 0x58, 0x82, 0xa4, 0x25, 0x91, 0xca, 0x23, 0x9a, 0xe4,
	0xf8, 0x5d, 0x59, 0x12, 0x50, 0x2c, 0xe5, 0xc6, 0x30, 0x15, 0xe2, 0xfc, 0xd2, 0x88, 0x22, 0x73,
	0x19, 0xe7, 0xa2, 0x64, 0x65, 0x26, 0x38, 0xaa, 0xbf, 0xfd, 0x89, 0xf5, 0x67, 0x0c, 0xe5, 0xbc,
	0x61, 0x4c, 0xdf, 0x11, 0x41, 0x06, 0xab, 0x19, 0x5c, 0x2c, 0x01, 0x4b, 0x7b, 0x60, 0xfd, 0xd0,
	0x79, 0x11, 0x67, 0x39, 0xfc, 0x34, 0x7a, 0xc6, 0xdf, 0x6f, 0xb3, 0xef, 0xfa, 0xe3, 0x31, 0xcb,
	0xe1, 0xff, 0xb5, 0x69, 0xfd, 0xee, 0xe0, 0xcc, 0xd5, 0x38, 0xf6, 0xb3, 0x61, 0x39, 0xdd, 0x69,
	0xf6, 0x11, 0xd9, 0xd6, 0x07, 0xd9, 0x7a, 0x57, 0x87, 0x76, 0x42, 0xda, 0xca, 0x48, 0x07, 0xa2,
	0x7f, 0x78, 0xf3, 0xf4, 0x72, 0x67, 0xee, 0xd9, 0xbb, 0x75, 0xc3, 0x97, 0x1f, 0x46, 0x3e, 0x88,
	0x97, 0x58, 0x8a, 0x1c, 0x24, 0xba, 0xff, 0xda, 0xca, 0x3f, 0x9d, 0xbf, 0x1a, 0xdd, 0x9a, 0xd6,
	0x30, 0x16, 0xf9, 0xd6, 0x39, 0x46, 0x83, 0xcd, 0x4d, 0x4d, 0xeb, 0xb7, 0x99, 0x1a, 0x27, 0xcd,
	0xaa, 0x91, 0x54, 0x2c, 0x18, 0x4f, 0x89, 0x90, 0xa9, 0x9b, 0x02, 0x5f, 0xbf, 0x9c, 0x5e, 0x87,
	0x22, 0xc3, 0xee, 0x4d, 0xdc, 0xd7, 0xe2, 0xde, 0xfc, 0x32, 0xf6, 0xfd, 0x07, 0xb3, 0x37, 0x56,
	0x40, 0x3f, 0x41, 0xa2, 0x64, 0xad, 0x02, 0x8f, 0x34, 0xc1, 0xf8, 0xa8, 0x2d, 0xa1, 0x9f, 0x60,
	0xd8, 0x5a, 0xc2, 0xc0, 0x0b, 0xb5, 0xe5, 0xd5, 0x1c, 0xaa, 0xef, 0x94, 0xfa, 0x09, 0x52, 0xda,
	0x9a, 0x28, 0x0d, 0x3c, 0x4a, 0xb5, 0xed, 0xf4, 0xeb, 0xfa, 0x9e, 0x3b, 0x6f, 0x01, 0x00, 0x00,
	0xff, 0xff, 0x6b, 0x59, 0x13, 0x0f, 0x30, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ShoppingPerformanceViewServiceClient is the client API for ShoppingPerformanceViewService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ShoppingPerformanceViewServiceClient interface {
	// Returns the requested Shopping performance view in full detail.
	GetShoppingPerformanceView(ctx context.Context, in *GetShoppingPerformanceViewRequest, opts ...grpc.CallOption) (*resources.ShoppingPerformanceView, error)
}

type shoppingPerformanceViewServiceClient struct {
	cc *grpc.ClientConn
}

func NewShoppingPerformanceViewServiceClient(cc *grpc.ClientConn) ShoppingPerformanceViewServiceClient {
	return &shoppingPerformanceViewServiceClient{cc}
}

func (c *shoppingPerformanceViewServiceClient) GetShoppingPerformanceView(ctx context.Context, in *GetShoppingPerformanceViewRequest, opts ...grpc.CallOption) (*resources.ShoppingPerformanceView, error) {
	out := new(resources.ShoppingPerformanceView)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v1.services.ShoppingPerformanceViewService/GetShoppingPerformanceView", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ShoppingPerformanceViewServiceServer is the server API for ShoppingPerformanceViewService service.
type ShoppingPerformanceViewServiceServer interface {
	// Returns the requested Shopping performance view in full detail.
	GetShoppingPerformanceView(context.Context, *GetShoppingPerformanceViewRequest) (*resources.ShoppingPerformanceView, error)
}

// UnimplementedShoppingPerformanceViewServiceServer can be embedded to have forward compatible implementations.
type UnimplementedShoppingPerformanceViewServiceServer struct {
}

func (*UnimplementedShoppingPerformanceViewServiceServer) GetShoppingPerformanceView(ctx context.Context, req *GetShoppingPerformanceViewRequest) (*resources.ShoppingPerformanceView, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetShoppingPerformanceView not implemented")
}

func RegisterShoppingPerformanceViewServiceServer(s *grpc.Server, srv ShoppingPerformanceViewServiceServer) {
	s.RegisterService(&_ShoppingPerformanceViewService_serviceDesc, srv)
}

func _ShoppingPerformanceViewService_GetShoppingPerformanceView_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetShoppingPerformanceViewRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShoppingPerformanceViewServiceServer).GetShoppingPerformanceView(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v1.services.ShoppingPerformanceViewService/GetShoppingPerformanceView",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShoppingPerformanceViewServiceServer).GetShoppingPerformanceView(ctx, req.(*GetShoppingPerformanceViewRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ShoppingPerformanceViewService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v1.services.ShoppingPerformanceViewService",
	HandlerType: (*ShoppingPerformanceViewServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetShoppingPerformanceView",
			Handler:    _ShoppingPerformanceViewService_GetShoppingPerformanceView_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v1/services/shopping_performance_view_service.proto",
}
