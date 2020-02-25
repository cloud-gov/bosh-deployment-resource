// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/bytestream/bytestream.proto

package bytestream

import (
	context "context"
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	_ "github.com/golang/protobuf/ptypes/wrappers"
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

// Request object for ByteStream.Read.
type ReadRequest struct {
	// The name of the resource to read.
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	// The offset for the first byte to return in the read, relative to the start
	// of the resource.
	//
	// A `read_offset` that is negative or greater than the size of the resource
	// will cause an `OUT_OF_RANGE` error.
	ReadOffset int64 `protobuf:"varint,2,opt,name=read_offset,json=readOffset,proto3" json:"read_offset,omitempty"`
	// The maximum number of `data` bytes the server is allowed to return in the
	// sum of all `ReadResponse` messages. A `read_limit` of zero indicates that
	// there is no limit, and a negative `read_limit` will cause an error.
	//
	// If the stream returns fewer bytes than allowed by the `read_limit` and no
	// error occurred, the stream includes all data from the `read_offset` to the
	// end of the resource.
	ReadLimit            int64    `protobuf:"varint,3,opt,name=read_limit,json=readLimit,proto3" json:"read_limit,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReadRequest) Reset()         { *m = ReadRequest{} }
func (m *ReadRequest) String() string { return proto.CompactTextString(m) }
func (*ReadRequest) ProtoMessage()    {}
func (*ReadRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_657cab877f44cd08, []int{0}
}

func (m *ReadRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReadRequest.Unmarshal(m, b)
}
func (m *ReadRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReadRequest.Marshal(b, m, deterministic)
}
func (m *ReadRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReadRequest.Merge(m, src)
}
func (m *ReadRequest) XXX_Size() int {
	return xxx_messageInfo_ReadRequest.Size(m)
}
func (m *ReadRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ReadRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ReadRequest proto.InternalMessageInfo

func (m *ReadRequest) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func (m *ReadRequest) GetReadOffset() int64 {
	if m != nil {
		return m.ReadOffset
	}
	return 0
}

func (m *ReadRequest) GetReadLimit() int64 {
	if m != nil {
		return m.ReadLimit
	}
	return 0
}

// Response object for ByteStream.Read.
type ReadResponse struct {
	// A portion of the data for the resource. The service **may** leave `data`
	// empty for any given `ReadResponse`. This enables the service to inform the
	// client that the request is still live while it is running an operation to
	// generate more data.
	Data                 []byte   `protobuf:"bytes,10,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReadResponse) Reset()         { *m = ReadResponse{} }
func (m *ReadResponse) String() string { return proto.CompactTextString(m) }
func (*ReadResponse) ProtoMessage()    {}
func (*ReadResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_657cab877f44cd08, []int{1}
}

func (m *ReadResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReadResponse.Unmarshal(m, b)
}
func (m *ReadResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReadResponse.Marshal(b, m, deterministic)
}
func (m *ReadResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReadResponse.Merge(m, src)
}
func (m *ReadResponse) XXX_Size() int {
	return xxx_messageInfo_ReadResponse.Size(m)
}
func (m *ReadResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ReadResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ReadResponse proto.InternalMessageInfo

func (m *ReadResponse) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

// Request object for ByteStream.Write.
type WriteRequest struct {
	// The name of the resource to write. This **must** be set on the first
	// `WriteRequest` of each `Write()` action. If it is set on subsequent calls,
	// it **must** match the value of the first request.
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	// The offset from the beginning of the resource at which the data should be
	// written. It is required on all `WriteRequest`s.
	//
	// In the first `WriteRequest` of a `Write()` action, it indicates
	// the initial offset for the `Write()` call. The value **must** be equal to
	// the `committed_size` that a call to `QueryWriteStatus()` would return.
	//
	// On subsequent calls, this value **must** be set and **must** be equal to
	// the sum of the first `write_offset` and the sizes of all `data` bundles
	// sent previously on this stream.
	//
	// An incorrect value will cause an error.
	WriteOffset int64 `protobuf:"varint,2,opt,name=write_offset,json=writeOffset,proto3" json:"write_offset,omitempty"`
	// If `true`, this indicates that the write is complete. Sending any
	// `WriteRequest`s subsequent to one in which `finish_write` is `true` will
	// cause an error.
	FinishWrite bool `protobuf:"varint,3,opt,name=finish_write,json=finishWrite,proto3" json:"finish_write,omitempty"`
	// A portion of the data for the resource. The client **may** leave `data`
	// empty for any given `WriteRequest`. This enables the client to inform the
	// service that the request is still live while it is running an operation to
	// generate more data.
	Data                 []byte   `protobuf:"bytes,10,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *WriteRequest) Reset()         { *m = WriteRequest{} }
func (m *WriteRequest) String() string { return proto.CompactTextString(m) }
func (*WriteRequest) ProtoMessage()    {}
func (*WriteRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_657cab877f44cd08, []int{2}
}

func (m *WriteRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WriteRequest.Unmarshal(m, b)
}
func (m *WriteRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WriteRequest.Marshal(b, m, deterministic)
}
func (m *WriteRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WriteRequest.Merge(m, src)
}
func (m *WriteRequest) XXX_Size() int {
	return xxx_messageInfo_WriteRequest.Size(m)
}
func (m *WriteRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_WriteRequest.DiscardUnknown(m)
}

var xxx_messageInfo_WriteRequest proto.InternalMessageInfo

func (m *WriteRequest) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func (m *WriteRequest) GetWriteOffset() int64 {
	if m != nil {
		return m.WriteOffset
	}
	return 0
}

func (m *WriteRequest) GetFinishWrite() bool {
	if m != nil {
		return m.FinishWrite
	}
	return false
}

func (m *WriteRequest) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

// Response object for ByteStream.Write.
type WriteResponse struct {
	// The number of bytes that have been processed for the given resource.
	CommittedSize        int64    `protobuf:"varint,1,opt,name=committed_size,json=committedSize,proto3" json:"committed_size,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *WriteResponse) Reset()         { *m = WriteResponse{} }
func (m *WriteResponse) String() string { return proto.CompactTextString(m) }
func (*WriteResponse) ProtoMessage()    {}
func (*WriteResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_657cab877f44cd08, []int{3}
}

func (m *WriteResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WriteResponse.Unmarshal(m, b)
}
func (m *WriteResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WriteResponse.Marshal(b, m, deterministic)
}
func (m *WriteResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WriteResponse.Merge(m, src)
}
func (m *WriteResponse) XXX_Size() int {
	return xxx_messageInfo_WriteResponse.Size(m)
}
func (m *WriteResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_WriteResponse.DiscardUnknown(m)
}

var xxx_messageInfo_WriteResponse proto.InternalMessageInfo

func (m *WriteResponse) GetCommittedSize() int64 {
	if m != nil {
		return m.CommittedSize
	}
	return 0
}

// Request object for ByteStream.QueryWriteStatus.
type QueryWriteStatusRequest struct {
	// The name of the resource whose write status is being requested.
	ResourceName         string   `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *QueryWriteStatusRequest) Reset()         { *m = QueryWriteStatusRequest{} }
func (m *QueryWriteStatusRequest) String() string { return proto.CompactTextString(m) }
func (*QueryWriteStatusRequest) ProtoMessage()    {}
func (*QueryWriteStatusRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_657cab877f44cd08, []int{4}
}

func (m *QueryWriteStatusRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_QueryWriteStatusRequest.Unmarshal(m, b)
}
func (m *QueryWriteStatusRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_QueryWriteStatusRequest.Marshal(b, m, deterministic)
}
func (m *QueryWriteStatusRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryWriteStatusRequest.Merge(m, src)
}
func (m *QueryWriteStatusRequest) XXX_Size() int {
	return xxx_messageInfo_QueryWriteStatusRequest.Size(m)
}
func (m *QueryWriteStatusRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryWriteStatusRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryWriteStatusRequest proto.InternalMessageInfo

func (m *QueryWriteStatusRequest) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

// Response object for ByteStream.QueryWriteStatus.
type QueryWriteStatusResponse struct {
	// The number of bytes that have been processed for the given resource.
	CommittedSize int64 `protobuf:"varint,1,opt,name=committed_size,json=committedSize,proto3" json:"committed_size,omitempty"`
	// `complete` is `true` only if the client has sent a `WriteRequest` with
	// `finish_write` set to true, and the server has processed that request.
	Complete             bool     `protobuf:"varint,2,opt,name=complete,proto3" json:"complete,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *QueryWriteStatusResponse) Reset()         { *m = QueryWriteStatusResponse{} }
func (m *QueryWriteStatusResponse) String() string { return proto.CompactTextString(m) }
func (*QueryWriteStatusResponse) ProtoMessage()    {}
func (*QueryWriteStatusResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_657cab877f44cd08, []int{5}
}

func (m *QueryWriteStatusResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_QueryWriteStatusResponse.Unmarshal(m, b)
}
func (m *QueryWriteStatusResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_QueryWriteStatusResponse.Marshal(b, m, deterministic)
}
func (m *QueryWriteStatusResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryWriteStatusResponse.Merge(m, src)
}
func (m *QueryWriteStatusResponse) XXX_Size() int {
	return xxx_messageInfo_QueryWriteStatusResponse.Size(m)
}
func (m *QueryWriteStatusResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryWriteStatusResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryWriteStatusResponse proto.InternalMessageInfo

func (m *QueryWriteStatusResponse) GetCommittedSize() int64 {
	if m != nil {
		return m.CommittedSize
	}
	return 0
}

func (m *QueryWriteStatusResponse) GetComplete() bool {
	if m != nil {
		return m.Complete
	}
	return false
}

func init() {
	proto.RegisterType((*ReadRequest)(nil), "google.bytestream.ReadRequest")
	proto.RegisterType((*ReadResponse)(nil), "google.bytestream.ReadResponse")
	proto.RegisterType((*WriteRequest)(nil), "google.bytestream.WriteRequest")
	proto.RegisterType((*WriteResponse)(nil), "google.bytestream.WriteResponse")
	proto.RegisterType((*QueryWriteStatusRequest)(nil), "google.bytestream.QueryWriteStatusRequest")
	proto.RegisterType((*QueryWriteStatusResponse)(nil), "google.bytestream.QueryWriteStatusResponse")
}

func init() { proto.RegisterFile("google/bytestream/bytestream.proto", fileDescriptor_657cab877f44cd08) }

var fileDescriptor_657cab877f44cd08 = []byte{
	// 446 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x53, 0x5b, 0x8b, 0x13, 0x31,
	0x14, 0x66, 0xb6, 0xab, 0x74, 0x4f, 0xa7, 0x5e, 0x02, 0xe2, 0x30, 0xe8, 0x6e, 0x77, 0x44, 0x28,
	0x0a, 0x33, 0xa2, 0xe0, 0xcb, 0x82, 0x0f, 0x7d, 0x13, 0x16, 0x2f, 0xd9, 0x07, 0x41, 0x90, 0x21,
	0x6d, 0x4f, 0xc7, 0x60, 0x93, 0x8c, 0x49, 0x86, 0xa5, 0xfb, 0x1f, 0x7c, 0xf1, 0x17, 0x4b, 0x92,
	0xb1, 0x1d, 0x6d, 0x0b, 0xdb, 0xb7, 0xe4, 0xbb, 0xcc, 0xf9, 0xe6, 0xe4, 0x1c, 0xc8, 0x2a, 0xa5,
	0xaa, 0x25, 0x16, 0xd3, 0x95, 0x45, 0x63, 0x35, 0x32, 0xd1, 0x39, 0xe6, 0xb5, 0x56, 0x56, 0x91,
	0x87, 0x41, 0x93, 0x6f, 0x88, 0xf4, 0x49, 0x6b, 0x63, 0x35, 0x2f, 0x98, 0x94, 0xca, 0x32, 0xcb,
	0x95, 0x34, 0xc1, 0x90, 0x9e, 0xb6, 0xac, 0xbf, 0x4d, 0x9b, 0x45, 0x71, 0xad, 0x59, 0x5d, 0xa3,
	0x6e, 0xf9, 0x4c, 0xc3, 0x80, 0x22, 0x9b, 0x53, 0xfc, 0xd9, 0xa0, 0xb1, 0xe4, 0x19, 0x0c, 0x35,
	0x1a, 0xd5, 0xe8, 0x19, 0x96, 0x92, 0x09, 0x4c, 0xa2, 0x51, 0x34, 0x3e, 0xa1, 0xf1, 0x5f, 0xf0,
	0x03, 0x13, 0x48, 0xce, 0x60, 0xa0, 0x91, 0xcd, 0x4b, 0xb5, 0x58, 0x18, 0xb4, 0xc9, 0xd1, 0x28,
	0x1a, 0xf7, 0x28, 0x38, 0xe8, 0xa3, 0x47, 0xc8, 0x53, 0xf0, 0xb7, 0x72, 0xc9, 0x05, 0xb7, 0x49,
	0xcf, 0xf3, 0x27, 0x0e, 0xb9, 0x74, 0x40, 0x96, 0x41, 0x1c, 0x6a, 0x9a, 0x5a, 0x49, 0x83, 0x84,
	0xc0, 0xf1, 0x9c, 0x59, 0x96, 0xc0, 0x28, 0x1a, 0xc7, 0xd4, 0x9f, 0xb3, 0x5f, 0x11, 0xc4, 0x5f,
	0x34, 0xb7, 0x78, 0x50, 0xb2, 0x73, 0x88, 0xaf, 0x9d, 0xe9, 0xdf, 0x68, 0x03, 0x8f, 0xb5, 0xd9,
	0xce, 0x21, 0x5e, 0x70, 0xc9, 0xcd, 0xf7, 0xd2, 0xa3, 0x3e, 0x5d, 0x9f, 0x0e, 0x02, 0xe6, 0x2b,
	0xee, 0xcc, 0xf3, 0x16, 0x86, 0x6d, 0x9c, 0x36, 0xf4, 0x73, 0xb8, 0x37, 0x53, 0x42, 0x70, 0x6b,
	0x71, 0x5e, 0x1a, 0x7e, 0x13, 0x02, 0xf5, 0xe8, 0x70, 0x8d, 0x5e, 0xf1, 0x1b, 0xcc, 0xde, 0xc1,
	0xe3, 0xcf, 0x0d, 0xea, 0x95, 0x37, 0x5f, 0x59, 0x66, 0x1b, 0x73, 0xc8, 0x1f, 0x65, 0xdf, 0x20,
	0xd9, 0xf6, 0x1f, 0x14, 0x81, 0xa4, 0xd0, 0x9f, 0x29, 0x51, 0x2f, 0xd1, 0xa2, 0x6f, 0x48, 0x9f,
	0xae, 0xef, 0xaf, 0x7f, 0x1f, 0x01, 0x4c, 0x56, 0xee, 0xcb, 0x6e, 0x96, 0xc8, 0x7b, 0x38, 0x76,
	0x2f, 0x43, 0x4e, 0xf3, 0xad, 0x39, 0xcb, 0x3b, 0x63, 0x92, 0x9e, 0xed, 0xe5, 0x43, 0xb4, 0x57,
	0x11, 0xb9, 0x84, 0x3b, 0xa1, 0x9b, 0xbb, 0xb4, 0xdd, 0x97, 0x4d, 0x47, 0xfb, 0x05, 0xe1, 0x6b,
	0xe3, 0x88, 0xfc, 0x80, 0x07, 0xff, 0xb7, 0x81, 0xbc, 0xd8, 0xe1, 0xdb, 0xd3, 0xeb, 0xf4, 0xe5,
	0xad, 0xb4, 0xa1, 0xdc, 0x04, 0xe1, 0xd1, 0x4c, 0x89, 0x6d, 0xc7, 0xe4, 0xfe, 0xa6, 0x55, 0x9f,
	0xdc, 0xf6, 0x7c, 0xbd, 0x68, 0x35, 0x95, 0x5a, 0x32, 0x59, 0xe5, 0x4a, 0x57, 0x45, 0x85, 0xd2,
	0x6f, 0x56, 0x11, 0x28, 0x56, 0x73, 0xd3, 0x59, 0xe3, 0x8b, 0xcd, 0x71, 0x7a, 0xd7, 0xeb, 0xde,
	0xfc, 0x09, 0x00, 0x00, 0xff, 0xff, 0x8f, 0x91, 0x09, 0xd4, 0xf8, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ByteStreamClient is the client API for ByteStream service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ByteStreamClient interface {
	// `Read()` is used to retrieve the contents of a resource as a sequence
	// of bytes. The bytes are returned in a sequence of responses, and the
	// responses are delivered as the results of a server-side streaming RPC.
	Read(ctx context.Context, in *ReadRequest, opts ...grpc.CallOption) (ByteStream_ReadClient, error)
	// `Write()` is used to send the contents of a resource as a sequence of
	// bytes. The bytes are sent in a sequence of request protos of a client-side
	// streaming RPC.
	//
	// A `Write()` action is resumable. If there is an error or the connection is
	// broken during the `Write()`, the client should check the status of the
	// `Write()` by calling `QueryWriteStatus()` and continue writing from the
	// returned `committed_size`. This may be less than the amount of data the
	// client previously sent.
	//
	// Calling `Write()` on a resource name that was previously written and
	// finalized could cause an error, depending on whether the underlying service
	// allows over-writing of previously written resources.
	//
	// When the client closes the request channel, the service will respond with
	// a `WriteResponse`. The service will not view the resource as `complete`
	// until the client has sent a `WriteRequest` with `finish_write` set to
	// `true`. Sending any requests on a stream after sending a request with
	// `finish_write` set to `true` will cause an error. The client **should**
	// check the `WriteResponse` it receives to determine how much data the
	// service was able to commit and whether the service views the resource as
	// `complete` or not.
	Write(ctx context.Context, opts ...grpc.CallOption) (ByteStream_WriteClient, error)
	// `QueryWriteStatus()` is used to find the `committed_size` for a resource
	// that is being written, which can then be used as the `write_offset` for
	// the next `Write()` call.
	//
	// If the resource does not exist (i.e., the resource has been deleted, or the
	// first `Write()` has not yet reached the service), this method returns the
	// error `NOT_FOUND`.
	//
	// The client **may** call `QueryWriteStatus()` at any time to determine how
	// much data has been processed for this resource. This is useful if the
	// client is buffering data and needs to know which data can be safely
	// evicted. For any sequence of `QueryWriteStatus()` calls for a given
	// resource name, the sequence of returned `committed_size` values will be
	// non-decreasing.
	QueryWriteStatus(ctx context.Context, in *QueryWriteStatusRequest, opts ...grpc.CallOption) (*QueryWriteStatusResponse, error)
}

type byteStreamClient struct {
	cc grpc.ClientConnInterface
}

func NewByteStreamClient(cc grpc.ClientConnInterface) ByteStreamClient {
	return &byteStreamClient{cc}
}

func (c *byteStreamClient) Read(ctx context.Context, in *ReadRequest, opts ...grpc.CallOption) (ByteStream_ReadClient, error) {
	stream, err := c.cc.NewStream(ctx, &_ByteStream_serviceDesc.Streams[0], "/google.bytestream.ByteStream/Read", opts...)
	if err != nil {
		return nil, err
	}
	x := &byteStreamReadClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ByteStream_ReadClient interface {
	Recv() (*ReadResponse, error)
	grpc.ClientStream
}

type byteStreamReadClient struct {
	grpc.ClientStream
}

func (x *byteStreamReadClient) Recv() (*ReadResponse, error) {
	m := new(ReadResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *byteStreamClient) Write(ctx context.Context, opts ...grpc.CallOption) (ByteStream_WriteClient, error) {
	stream, err := c.cc.NewStream(ctx, &_ByteStream_serviceDesc.Streams[1], "/google.bytestream.ByteStream/Write", opts...)
	if err != nil {
		return nil, err
	}
	x := &byteStreamWriteClient{stream}
	return x, nil
}

type ByteStream_WriteClient interface {
	Send(*WriteRequest) error
	CloseAndRecv() (*WriteResponse, error)
	grpc.ClientStream
}

type byteStreamWriteClient struct {
	grpc.ClientStream
}

func (x *byteStreamWriteClient) Send(m *WriteRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *byteStreamWriteClient) CloseAndRecv() (*WriteResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(WriteResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *byteStreamClient) QueryWriteStatus(ctx context.Context, in *QueryWriteStatusRequest, opts ...grpc.CallOption) (*QueryWriteStatusResponse, error) {
	out := new(QueryWriteStatusResponse)
	err := c.cc.Invoke(ctx, "/google.bytestream.ByteStream/QueryWriteStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ByteStreamServer is the server API for ByteStream service.
type ByteStreamServer interface {
	// `Read()` is used to retrieve the contents of a resource as a sequence
	// of bytes. The bytes are returned in a sequence of responses, and the
	// responses are delivered as the results of a server-side streaming RPC.
	Read(*ReadRequest, ByteStream_ReadServer) error
	// `Write()` is used to send the contents of a resource as a sequence of
	// bytes. The bytes are sent in a sequence of request protos of a client-side
	// streaming RPC.
	//
	// A `Write()` action is resumable. If there is an error or the connection is
	// broken during the `Write()`, the client should check the status of the
	// `Write()` by calling `QueryWriteStatus()` and continue writing from the
	// returned `committed_size`. This may be less than the amount of data the
	// client previously sent.
	//
	// Calling `Write()` on a resource name that was previously written and
	// finalized could cause an error, depending on whether the underlying service
	// allows over-writing of previously written resources.
	//
	// When the client closes the request channel, the service will respond with
	// a `WriteResponse`. The service will not view the resource as `complete`
	// until the client has sent a `WriteRequest` with `finish_write` set to
	// `true`. Sending any requests on a stream after sending a request with
	// `finish_write` set to `true` will cause an error. The client **should**
	// check the `WriteResponse` it receives to determine how much data the
	// service was able to commit and whether the service views the resource as
	// `complete` or not.
	Write(ByteStream_WriteServer) error
	// `QueryWriteStatus()` is used to find the `committed_size` for a resource
	// that is being written, which can then be used as the `write_offset` for
	// the next `Write()` call.
	//
	// If the resource does not exist (i.e., the resource has been deleted, or the
	// first `Write()` has not yet reached the service), this method returns the
	// error `NOT_FOUND`.
	//
	// The client **may** call `QueryWriteStatus()` at any time to determine how
	// much data has been processed for this resource. This is useful if the
	// client is buffering data and needs to know which data can be safely
	// evicted. For any sequence of `QueryWriteStatus()` calls for a given
	// resource name, the sequence of returned `committed_size` values will be
	// non-decreasing.
	QueryWriteStatus(context.Context, *QueryWriteStatusRequest) (*QueryWriteStatusResponse, error)
}

// UnimplementedByteStreamServer can be embedded to have forward compatible implementations.
type UnimplementedByteStreamServer struct {
}

func (*UnimplementedByteStreamServer) Read(req *ReadRequest, srv ByteStream_ReadServer) error {
	return status.Errorf(codes.Unimplemented, "method Read not implemented")
}
func (*UnimplementedByteStreamServer) Write(srv ByteStream_WriteServer) error {
	return status.Errorf(codes.Unimplemented, "method Write not implemented")
}
func (*UnimplementedByteStreamServer) QueryWriteStatus(ctx context.Context, req *QueryWriteStatusRequest) (*QueryWriteStatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryWriteStatus not implemented")
}

func RegisterByteStreamServer(s *grpc.Server, srv ByteStreamServer) {
	s.RegisterService(&_ByteStream_serviceDesc, srv)
}

func _ByteStream_Read_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ReadRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ByteStreamServer).Read(m, &byteStreamReadServer{stream})
}

type ByteStream_ReadServer interface {
	Send(*ReadResponse) error
	grpc.ServerStream
}

type byteStreamReadServer struct {
	grpc.ServerStream
}

func (x *byteStreamReadServer) Send(m *ReadResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _ByteStream_Write_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ByteStreamServer).Write(&byteStreamWriteServer{stream})
}

type ByteStream_WriteServer interface {
	SendAndClose(*WriteResponse) error
	Recv() (*WriteRequest, error)
	grpc.ServerStream
}

type byteStreamWriteServer struct {
	grpc.ServerStream
}

func (x *byteStreamWriteServer) SendAndClose(m *WriteResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *byteStreamWriteServer) Recv() (*WriteRequest, error) {
	m := new(WriteRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _ByteStream_QueryWriteStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryWriteStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ByteStreamServer).QueryWriteStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.bytestream.ByteStream/QueryWriteStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ByteStreamServer).QueryWriteStatus(ctx, req.(*QueryWriteStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ByteStream_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.bytestream.ByteStream",
	HandlerType: (*ByteStreamServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "QueryWriteStatus",
			Handler:    _ByteStream_QueryWriteStatus_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Read",
			Handler:       _ByteStream_Read_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "Write",
			Handler:       _ByteStream_Write_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "google/bytestream/bytestream.proto",
}
