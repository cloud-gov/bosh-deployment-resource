// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v1/services/mutate_job_service.proto

package services

import (
	context "context"
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	resources "google.golang.org/genproto/googleapis/ads/googleads/v1/resources"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	longrunning "google.golang.org/genproto/googleapis/longrunning"
	status "google.golang.org/genproto/googleapis/rpc/status"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status1 "google.golang.org/grpc/status"
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

// Request message for [MutateJobService.CreateMutateJobRequest][]
type CreateMutateJobRequest struct {
	// The ID of the customer for which to create a mutate job.
	CustomerId           string   `protobuf:"bytes,1,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateMutateJobRequest) Reset()         { *m = CreateMutateJobRequest{} }
func (m *CreateMutateJobRequest) String() string { return proto.CompactTextString(m) }
func (*CreateMutateJobRequest) ProtoMessage()    {}
func (*CreateMutateJobRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_4b20a6cc488e9516, []int{0}
}

func (m *CreateMutateJobRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateMutateJobRequest.Unmarshal(m, b)
}
func (m *CreateMutateJobRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateMutateJobRequest.Marshal(b, m, deterministic)
}
func (m *CreateMutateJobRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateMutateJobRequest.Merge(m, src)
}
func (m *CreateMutateJobRequest) XXX_Size() int {
	return xxx_messageInfo_CreateMutateJobRequest.Size(m)
}
func (m *CreateMutateJobRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateMutateJobRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateMutateJobRequest proto.InternalMessageInfo

func (m *CreateMutateJobRequest) GetCustomerId() string {
	if m != nil {
		return m.CustomerId
	}
	return ""
}

// Response message for [MutateJobService.CreateMutateJobResponse][]
type CreateMutateJobResponse struct {
	// The resource name of the MutateJob.
	ResourceName         string   `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateMutateJobResponse) Reset()         { *m = CreateMutateJobResponse{} }
func (m *CreateMutateJobResponse) String() string { return proto.CompactTextString(m) }
func (*CreateMutateJobResponse) ProtoMessage()    {}
func (*CreateMutateJobResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_4b20a6cc488e9516, []int{1}
}

func (m *CreateMutateJobResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateMutateJobResponse.Unmarshal(m, b)
}
func (m *CreateMutateJobResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateMutateJobResponse.Marshal(b, m, deterministic)
}
func (m *CreateMutateJobResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateMutateJobResponse.Merge(m, src)
}
func (m *CreateMutateJobResponse) XXX_Size() int {
	return xxx_messageInfo_CreateMutateJobResponse.Size(m)
}
func (m *CreateMutateJobResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateMutateJobResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateMutateJobResponse proto.InternalMessageInfo

func (m *CreateMutateJobResponse) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

// Request message for [MutateJobService.GetMutateJob][google.ads.googleads.v1.services.MutateJobService.GetMutateJob]
type GetMutateJobRequest struct {
	// The resource name of the MutateJob to get.
	ResourceName         string   `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetMutateJobRequest) Reset()         { *m = GetMutateJobRequest{} }
func (m *GetMutateJobRequest) String() string { return proto.CompactTextString(m) }
func (*GetMutateJobRequest) ProtoMessage()    {}
func (*GetMutateJobRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_4b20a6cc488e9516, []int{2}
}

func (m *GetMutateJobRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetMutateJobRequest.Unmarshal(m, b)
}
func (m *GetMutateJobRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetMutateJobRequest.Marshal(b, m, deterministic)
}
func (m *GetMutateJobRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetMutateJobRequest.Merge(m, src)
}
func (m *GetMutateJobRequest) XXX_Size() int {
	return xxx_messageInfo_GetMutateJobRequest.Size(m)
}
func (m *GetMutateJobRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetMutateJobRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetMutateJobRequest proto.InternalMessageInfo

func (m *GetMutateJobRequest) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

// Request message for [MutateJobService.RunMutateJob][google.ads.googleads.v1.services.MutateJobService.RunMutateJob]
type RunMutateJobRequest struct {
	// The resource name of the MutateJob to run.
	ResourceName         string   `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RunMutateJobRequest) Reset()         { *m = RunMutateJobRequest{} }
func (m *RunMutateJobRequest) String() string { return proto.CompactTextString(m) }
func (*RunMutateJobRequest) ProtoMessage()    {}
func (*RunMutateJobRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_4b20a6cc488e9516, []int{3}
}

func (m *RunMutateJobRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RunMutateJobRequest.Unmarshal(m, b)
}
func (m *RunMutateJobRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RunMutateJobRequest.Marshal(b, m, deterministic)
}
func (m *RunMutateJobRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RunMutateJobRequest.Merge(m, src)
}
func (m *RunMutateJobRequest) XXX_Size() int {
	return xxx_messageInfo_RunMutateJobRequest.Size(m)
}
func (m *RunMutateJobRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RunMutateJobRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RunMutateJobRequest proto.InternalMessageInfo

func (m *RunMutateJobRequest) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

// Request message for [MutateJobService.AddMutateJobOperations][google.ads.googleads.v1.services.MutateJobService.AddMutateJobOperations]
type AddMutateJobOperationsRequest struct {
	// The resource name of the MutateJob.
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	// A token used to enforce sequencing.
	//
	// The first AddMutateJobOperations request for a MutateJob should not set
	// sequence_token. Subsequent requests must set sequence_token to the value of
	// next_sequence_token received in the previous AddMutateJobOperations
	// response.
	SequenceToken string `protobuf:"bytes,2,opt,name=sequence_token,json=sequenceToken,proto3" json:"sequence_token,omitempty"`
	// The list of mutates being added.
	//
	// Operations can use negative integers as temp ids to signify dependencies
	// between entities created in this MutateJob. For example, a customer with
	// id = 1234 can create a campaign and an ad group in that same campaign by
	// creating a campaign in the first operation with the resource name
	// explicitly set to "customers/1234/campaigns/-1", and creating an ad group
	// in the second operation with the campaign field also set to
	// "customers/1234/campaigns/-1".
	MutateOperations     []*MutateOperation `protobuf:"bytes,3,rep,name=mutate_operations,json=mutateOperations,proto3" json:"mutate_operations,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *AddMutateJobOperationsRequest) Reset()         { *m = AddMutateJobOperationsRequest{} }
func (m *AddMutateJobOperationsRequest) String() string { return proto.CompactTextString(m) }
func (*AddMutateJobOperationsRequest) ProtoMessage()    {}
func (*AddMutateJobOperationsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_4b20a6cc488e9516, []int{4}
}

func (m *AddMutateJobOperationsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddMutateJobOperationsRequest.Unmarshal(m, b)
}
func (m *AddMutateJobOperationsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddMutateJobOperationsRequest.Marshal(b, m, deterministic)
}
func (m *AddMutateJobOperationsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddMutateJobOperationsRequest.Merge(m, src)
}
func (m *AddMutateJobOperationsRequest) XXX_Size() int {
	return xxx_messageInfo_AddMutateJobOperationsRequest.Size(m)
}
func (m *AddMutateJobOperationsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AddMutateJobOperationsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AddMutateJobOperationsRequest proto.InternalMessageInfo

func (m *AddMutateJobOperationsRequest) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func (m *AddMutateJobOperationsRequest) GetSequenceToken() string {
	if m != nil {
		return m.SequenceToken
	}
	return ""
}

func (m *AddMutateJobOperationsRequest) GetMutateOperations() []*MutateOperation {
	if m != nil {
		return m.MutateOperations
	}
	return nil
}

// Response message for [MutateJobService.AddMutateJobOperations][google.ads.googleads.v1.services.MutateJobService.AddMutateJobOperations]
type AddMutateJobOperationsResponse struct {
	// The total number of operations added so far for this job.
	TotalOperations int64 `protobuf:"varint,1,opt,name=total_operations,json=totalOperations,proto3" json:"total_operations,omitempty"`
	// The sequence token to be used when calling AddMutateJobOperations again if
	// more operations need to be added. The next AddMutateJobOperations request
	// must set the sequence_token field to the value of this field.
	NextSequenceToken    string   `protobuf:"bytes,2,opt,name=next_sequence_token,json=nextSequenceToken,proto3" json:"next_sequence_token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddMutateJobOperationsResponse) Reset()         { *m = AddMutateJobOperationsResponse{} }
func (m *AddMutateJobOperationsResponse) String() string { return proto.CompactTextString(m) }
func (*AddMutateJobOperationsResponse) ProtoMessage()    {}
func (*AddMutateJobOperationsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_4b20a6cc488e9516, []int{5}
}

func (m *AddMutateJobOperationsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddMutateJobOperationsResponse.Unmarshal(m, b)
}
func (m *AddMutateJobOperationsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddMutateJobOperationsResponse.Marshal(b, m, deterministic)
}
func (m *AddMutateJobOperationsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddMutateJobOperationsResponse.Merge(m, src)
}
func (m *AddMutateJobOperationsResponse) XXX_Size() int {
	return xxx_messageInfo_AddMutateJobOperationsResponse.Size(m)
}
func (m *AddMutateJobOperationsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AddMutateJobOperationsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AddMutateJobOperationsResponse proto.InternalMessageInfo

func (m *AddMutateJobOperationsResponse) GetTotalOperations() int64 {
	if m != nil {
		return m.TotalOperations
	}
	return 0
}

func (m *AddMutateJobOperationsResponse) GetNextSequenceToken() string {
	if m != nil {
		return m.NextSequenceToken
	}
	return ""
}

// Request message for [MutateJobService.ListMutateJobResults][google.ads.googleads.v1.services.MutateJobService.ListMutateJobResults].
type ListMutateJobResultsRequest struct {
	// The resource name of the MutateJob whose results are being listed.
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	// Token of the page to retrieve. If not specified, the first
	// page of results will be returned. Use the value obtained from
	// `next_page_token` in the previous response in order to request
	// the next page of results.
	PageToken string `protobuf:"bytes,2,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"`
	// Number of elements to retrieve in a single page.
	// When a page request is too large, the server may decide to
	// further limit the number of returned resources.
	PageSize             int32    `protobuf:"varint,3,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListMutateJobResultsRequest) Reset()         { *m = ListMutateJobResultsRequest{} }
func (m *ListMutateJobResultsRequest) String() string { return proto.CompactTextString(m) }
func (*ListMutateJobResultsRequest) ProtoMessage()    {}
func (*ListMutateJobResultsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_4b20a6cc488e9516, []int{6}
}

func (m *ListMutateJobResultsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListMutateJobResultsRequest.Unmarshal(m, b)
}
func (m *ListMutateJobResultsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListMutateJobResultsRequest.Marshal(b, m, deterministic)
}
func (m *ListMutateJobResultsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListMutateJobResultsRequest.Merge(m, src)
}
func (m *ListMutateJobResultsRequest) XXX_Size() int {
	return xxx_messageInfo_ListMutateJobResultsRequest.Size(m)
}
func (m *ListMutateJobResultsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListMutateJobResultsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListMutateJobResultsRequest proto.InternalMessageInfo

func (m *ListMutateJobResultsRequest) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func (m *ListMutateJobResultsRequest) GetPageToken() string {
	if m != nil {
		return m.PageToken
	}
	return ""
}

func (m *ListMutateJobResultsRequest) GetPageSize() int32 {
	if m != nil {
		return m.PageSize
	}
	return 0
}

// Response message for [MutateJobService.ListMutateJobResults][google.ads.googleads.v1.services.MutateJobService.ListMutateJobResults].
type ListMutateJobResultsResponse struct {
	// The list of rows that matched the query.
	Results []*MutateJobResult `protobuf:"bytes,1,rep,name=results,proto3" json:"results,omitempty"`
	// Pagination token used to retrieve the next page of results.
	// Pass the content of this string as the `page_token` attribute of
	// the next request. `next_page_token` is not returned for the last
	// page.
	NextPageToken        string   `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListMutateJobResultsResponse) Reset()         { *m = ListMutateJobResultsResponse{} }
func (m *ListMutateJobResultsResponse) String() string { return proto.CompactTextString(m) }
func (*ListMutateJobResultsResponse) ProtoMessage()    {}
func (*ListMutateJobResultsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_4b20a6cc488e9516, []int{7}
}

func (m *ListMutateJobResultsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListMutateJobResultsResponse.Unmarshal(m, b)
}
func (m *ListMutateJobResultsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListMutateJobResultsResponse.Marshal(b, m, deterministic)
}
func (m *ListMutateJobResultsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListMutateJobResultsResponse.Merge(m, src)
}
func (m *ListMutateJobResultsResponse) XXX_Size() int {
	return xxx_messageInfo_ListMutateJobResultsResponse.Size(m)
}
func (m *ListMutateJobResultsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListMutateJobResultsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListMutateJobResultsResponse proto.InternalMessageInfo

func (m *ListMutateJobResultsResponse) GetResults() []*MutateJobResult {
	if m != nil {
		return m.Results
	}
	return nil
}

func (m *ListMutateJobResultsResponse) GetNextPageToken() string {
	if m != nil {
		return m.NextPageToken
	}
	return ""
}

// MutateJob result.
type MutateJobResult struct {
	// Index of the mutate operation.
	OperationIndex int64 `protobuf:"varint,1,opt,name=operation_index,json=operationIndex,proto3" json:"operation_index,omitempty"`
	// Response for the mutate.
	// May be empty if errors occurred.
	MutateOperationResponse *MutateOperationResponse `protobuf:"bytes,2,opt,name=mutate_operation_response,json=mutateOperationResponse,proto3" json:"mutate_operation_response,omitempty"`
	// Details of the errors when processing the operation.
	Status               *status.Status `protobuf:"bytes,3,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *MutateJobResult) Reset()         { *m = MutateJobResult{} }
func (m *MutateJobResult) String() string { return proto.CompactTextString(m) }
func (*MutateJobResult) ProtoMessage()    {}
func (*MutateJobResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_4b20a6cc488e9516, []int{8}
}

func (m *MutateJobResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MutateJobResult.Unmarshal(m, b)
}
func (m *MutateJobResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MutateJobResult.Marshal(b, m, deterministic)
}
func (m *MutateJobResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MutateJobResult.Merge(m, src)
}
func (m *MutateJobResult) XXX_Size() int {
	return xxx_messageInfo_MutateJobResult.Size(m)
}
func (m *MutateJobResult) XXX_DiscardUnknown() {
	xxx_messageInfo_MutateJobResult.DiscardUnknown(m)
}

var xxx_messageInfo_MutateJobResult proto.InternalMessageInfo

func (m *MutateJobResult) GetOperationIndex() int64 {
	if m != nil {
		return m.OperationIndex
	}
	return 0
}

func (m *MutateJobResult) GetMutateOperationResponse() *MutateOperationResponse {
	if m != nil {
		return m.MutateOperationResponse
	}
	return nil
}

func (m *MutateJobResult) GetStatus() *status.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func init() {
	proto.RegisterType((*CreateMutateJobRequest)(nil), "google.ads.googleads.v1.services.CreateMutateJobRequest")
	proto.RegisterType((*CreateMutateJobResponse)(nil), "google.ads.googleads.v1.services.CreateMutateJobResponse")
	proto.RegisterType((*GetMutateJobRequest)(nil), "google.ads.googleads.v1.services.GetMutateJobRequest")
	proto.RegisterType((*RunMutateJobRequest)(nil), "google.ads.googleads.v1.services.RunMutateJobRequest")
	proto.RegisterType((*AddMutateJobOperationsRequest)(nil), "google.ads.googleads.v1.services.AddMutateJobOperationsRequest")
	proto.RegisterType((*AddMutateJobOperationsResponse)(nil), "google.ads.googleads.v1.services.AddMutateJobOperationsResponse")
	proto.RegisterType((*ListMutateJobResultsRequest)(nil), "google.ads.googleads.v1.services.ListMutateJobResultsRequest")
	proto.RegisterType((*ListMutateJobResultsResponse)(nil), "google.ads.googleads.v1.services.ListMutateJobResultsResponse")
	proto.RegisterType((*MutateJobResult)(nil), "google.ads.googleads.v1.services.MutateJobResult")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v1/services/mutate_job_service.proto", fileDescriptor_4b20a6cc488e9516)
}

var fileDescriptor_4b20a6cc488e9516 = []byte{
	// 915 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x56, 0x4f, 0x6f, 0xdc, 0x44,
	0x14, 0xd7, 0xec, 0x8a, 0xd2, 0xbe, 0x24, 0xdd, 0x74, 0x4a, 0x9b, 0xc5, 0x69, 0x60, 0xe5, 0x16,
	0x08, 0xab, 0xca, 0x66, 0x97, 0x56, 0x6a, 0x8c, 0xd2, 0x76, 0x8b, 0xaa, 0xd0, 0x42, 0x21, 0x72,
	0x50, 0x0e, 0x28, 0xc2, 0x9a, 0xb5, 0x07, 0xcb, 0xb0, 0x9e, 0x31, 0x9e, 0x71, 0x54, 0x1a, 0x95,
	0x03, 0x5f, 0xa1, 0x7c, 0x02, 0x6e, 0xf0, 0x31, 0x38, 0x86, 0x23, 0x07, 0x4e, 0xbd, 0x20, 0x4e,
	0x7c, 0x05, 0x38, 0x54, 0xb6, 0x67, 0xbc, 0x7f, 0xb2, 0xe9, 0x76, 0xf7, 0x36, 0xfb, 0x9b, 0xf7,
	0xe7, 0xf7, 0x7b, 0xef, 0xf9, 0xcd, 0xc2, 0x56, 0xc8, 0x79, 0x38, 0xa0, 0x36, 0x09, 0x84, 0x5d,
	0x1e, 0xf3, 0xd3, 0x61, 0xc7, 0x16, 0x34, 0x3d, 0x8c, 0x7c, 0x2a, 0xec, 0x38, 0x93, 0x44, 0x52,
	0xef, 0x5b, 0xde, 0xf7, 0x14, 0x66, 0x25, 0x29, 0x97, 0x1c, 0xb7, 0x4a, 0x7b, 0x8b, 0x04, 0xc2,
	0xaa, 0x5c, 0xad, 0xc3, 0x8e, 0xa5, 0x5d, 0x8d, 0xee, 0x69, 0xc1, 0x53, 0x2a, 0x78, 0x96, 0x8e,
	0x47, 0x2f, 0xa3, 0x1a, 0xb3, 0x09, 0x95, 0xa8, 0x47, 0x02, 0x31, 0x4e, 0xc8, 0xb8, 0xa2, 0x5d,
	0x93, 0xc8, 0x26, 0x8c, 0x71, 0x49, 0x64, 0xc4, 0x99, 0x50, 0xb7, 0x57, 0xd5, 0xed, 0x80, 0xb3,
	0x30, 0xcd, 0x18, 0x8b, 0x58, 0x68, 0xf3, 0x84, 0xa6, 0x63, 0x46, 0x6b, 0xca, 0x28, 0x4d, 0x7c,
	0x5b, 0x48, 0x22, 0xb3, 0xc9, 0x8b, 0x3c, 0xb6, 0x3f, 0x88, 0x28, 0x93, 0xe5, 0x85, 0xb9, 0x05,
	0x97, 0x3f, 0x4e, 0x29, 0x91, 0xf4, 0x51, 0xa1, 0xe4, 0x21, 0xef, 0xbb, 0xf4, 0xfb, 0x8c, 0x0a,
	0x89, 0xdf, 0x86, 0x25, 0x3f, 0x13, 0x92, 0xc7, 0x34, 0xf5, 0xa2, 0xa0, 0x89, 0x5a, 0x68, 0xf3,
	0x9c, 0x0b, 0x1a, 0x7a, 0x10, 0x98, 0xb7, 0x61, 0xed, 0x84, 0xab, 0x48, 0x38, 0x13, 0x14, 0x5f,
	0x85, 0x15, 0x5d, 0x23, 0x8f, 0x91, 0x98, 0x2a, 0xef, 0x65, 0x0d, 0x7e, 0x4e, 0x62, 0x6a, 0x3a,
	0x70, 0x71, 0x87, 0xca, 0x13, 0x79, 0x5f, 0xd5, 0xd7, 0xcd, 0xd8, 0x62, 0xbe, 0x7f, 0x20, 0xd8,
	0xe8, 0x05, 0x41, 0xe5, 0xfc, 0x45, 0x55, 0xc5, 0x79, 0xc2, 0xe0, 0x77, 0xe0, 0xbc, 0xc8, 0xed,
	0x99, 0x4f, 0x3d, 0xc9, 0xbf, 0xa3, 0xac, 0x59, 0x2b, 0xac, 0x56, 0x34, 0xfa, 0x65, 0x0e, 0xe2,
	0xaf, 0xe1, 0x82, 0x1a, 0x92, 0x61, 0xb7, 0x9a, 0xf5, 0x56, 0x7d, 0x73, 0xa9, 0xdb, 0xb1, 0x66,
	0x8d, 0xa0, 0x55, 0x92, 0xac, 0x18, 0xba, 0xab, 0xf1, 0x38, 0x20, 0xcc, 0x23, 0x78, 0xeb, 0x34,
	0x31, 0xaa, 0x19, 0xef, 0xc3, 0xaa, 0xe4, 0x92, 0x0c, 0x46, 0x09, 0xe4, 0x82, 0xea, 0x6e, 0xa3,
	0xc0, 0x87, 0x2e, 0xd8, 0x82, 0x8b, 0x8c, 0x3e, 0x96, 0xde, 0x54, 0x61, 0x17, 0xf2, 0xab, 0xbd,
	0x51, 0x71, 0xe6, 0x8f, 0xb0, 0xfe, 0x59, 0x24, 0x46, 0x7b, 0x28, 0xb2, 0x81, 0x9c, 0xaf, 0x8e,
	0x1b, 0x00, 0x09, 0x09, 0xc7, 0x53, 0x9d, 0xcb, 0x91, 0xb2, 0x7e, 0xeb, 0x50, 0xfc, 0xf0, 0x44,
	0xf4, 0x84, 0x36, 0xeb, 0x2d, 0xb4, 0xf9, 0x9a, 0x7b, 0x36, 0x07, 0xf6, 0xa2, 0x27, 0xd4, 0x7c,
	0x86, 0xe0, 0xca, 0x74, 0x02, 0x4a, 0xfb, 0xa7, 0xf0, 0x7a, 0x5a, 0x42, 0x4d, 0x34, 0x5f, 0xcd,
	0xab, 0x60, 0xae, 0x8e, 0x80, 0xdf, 0x85, 0x46, 0x51, 0x9d, 0x13, 0x74, 0x57, 0x72, 0x78, 0x57,
	0x53, 0x36, 0x9f, 0x23, 0x68, 0x4c, 0x04, 0xc1, 0xef, 0x41, 0xa3, 0x2a, 0xbf, 0x17, 0xb1, 0x80,
	0x3e, 0x56, 0x3d, 0x38, 0x5f, 0xc1, 0x0f, 0x72, 0x14, 0x67, 0xf0, 0xe6, 0xe4, 0xbc, 0x78, 0xa9,
	0x92, 0x53, 0xa4, 0x5b, 0xea, 0x6e, 0xcd, 0x3f, 0x37, 0x2a, 0x80, 0xbb, 0x16, 0x4f, 0xbf, 0xc0,
	0x6d, 0x38, 0x53, 0x2e, 0x8c, 0xa2, 0xc6, 0x4b, 0x5d, 0xac, 0x73, 0xa4, 0x89, 0x6f, 0xed, 0x15,
	0x37, 0xae, 0xb2, 0xe8, 0x3e, 0x3f, 0x0b, 0xab, 0x95, 0xbe, 0xbd, 0x32, 0x25, 0xfe, 0x1d, 0x41,
	0x63, 0x62, 0x1d, 0xe0, 0x5b, 0xb3, 0x89, 0x4e, 0x5f, 0x3e, 0xc6, 0xd6, 0x02, 0x9e, 0xa5, 0x12,
	0xd3, 0xf9, 0xe9, 0xcf, 0x7f, 0x9e, 0xd5, 0x6e, 0x98, 0x76, 0xbe, 0x75, 0xf5, 0xba, 0x12, 0xf6,
	0xd1, 0xc8, 0x32, 0xdb, 0x6e, 0x3f, 0x55, 0x9b, 0xfb, 0x21, 0xef, 0x0b, 0xc7, 0x2f, 0x22, 0x39,
	0xa8, 0x8d, 0x7f, 0x45, 0xb0, 0x3c, 0xba, 0x93, 0xf0, 0xcd, 0xd9, 0x3c, 0xa6, 0xec, 0x30, 0xe3,
	0xfa, 0xa9, 0x6e, 0xd5, 0xd3, 0x31, 0x1c, 0x33, 0xf3, 0x46, 0xc1, 0xd8, 0xc2, 0xd7, 0x73, 0xc6,
	0x47, 0x63, 0x5f, 0xcc, 0xf6, 0x50, 0x40, 0x7b, 0x84, 0xb2, 0xdd, 0x7e, 0x8a, 0xff, 0x42, 0xf0,
	0xc6, 0xb4, 0xd9, 0xc7, 0xdb, 0xb3, 0x39, 0xbf, 0xe4, 0xa3, 0x35, 0x6e, 0x2f, 0xea, 0xae, 0xea,
	0x7f, 0xb7, 0x50, 0xe3, 0xe0, 0x5b, 0xf3, 0xa8, 0x71, 0x06, 0x91, 0x90, 0x5a, 0xc0, 0x7f, 0x08,
	0x96, 0x47, 0xb7, 0xfb, 0xab, 0x74, 0x61, 0xca, 0x6b, 0x60, 0x6c, 0x68, 0xb7, 0x91, 0x37, 0xd3,
	0xaa, 0x46, 0xdf, 0xfc, 0x19, 0x1d, 0xf7, 0xf6, 0xe1, 0x92, 0xb2, 0x29, 0x9e, 0xc3, 0x7e, 0xf6,
	0x8d, 0x75, 0x3f, 0x4e, 0xe4, 0x0f, 0x78, 0x7b, 0x8e, 0x06, 0x0e, 0x4f, 0x8f, 0xa8, 0x24, 0x01,
	0x91, 0xa4, 0x28, 0xc1, 0x4d, 0xf3, 0x83, 0xb9, 0x4a, 0x90, 0x66, 0x2c, 0x9f, 0xc1, 0xbf, 0x11,
	0x5c, 0x9e, 0xbe, 0xd1, 0xf1, 0x9d, 0xd9, 0x75, 0x78, 0xe9, 0xc3, 0x66, 0xdc, 0x5d, 0x3c, 0x80,
	0xea, 0xee, 0xfd, 0x42, 0xda, 0x1d, 0xd3, 0x99, 0x4b, 0x1a, 0x09, 0x82, 0x61, 0x2c, 0x07, 0xb5,
	0x8d, 0xf5, 0xe3, 0x5e, 0x73, 0x98, 0x5f, 0x9d, 0x92, 0x48, 0x58, 0x3e, 0x8f, 0xef, 0xfd, 0x8f,
	0xe0, 0x9a, 0xcf, 0xe3, 0x99, 0x5c, 0xef, 0x5d, 0x9a, 0xdc, 0x42, 0xbb, 0x79, 0x13, 0x77, 0xd1,
	0x57, 0x9f, 0x28, 0xd7, 0x90, 0x0f, 0x08, 0x0b, 0x2d, 0x9e, 0x86, 0x76, 0x48, 0x59, 0xd1, 0x62,
	0x7b, 0x98, 0xec, 0xf4, 0x3f, 0x69, 0x1f, 0xe9, 0xc3, 0x2f, 0xb5, 0xfa, 0x4e, 0xaf, 0xf7, 0x5b,
	0xad, 0xb5, 0x53, 0x06, 0xec, 0x05, 0xc2, 0x2a, 0x8f, 0xf9, 0x69, 0xbf, 0x63, 0xa9, 0xc4, 0xe2,
	0x58, 0x9b, 0x1c, 0xf4, 0x02, 0x71, 0x50, 0x99, 0x1c, 0xec, 0x77, 0x0e, 0xb4, 0xc9, 0xbf, 0xb5,
	0x6b, 0x25, 0xee, 0x38, 0xbd, 0x40, 0x38, 0x4e, 0x65, 0xe4, 0x38, 0xfb, 0x1d, 0xc7, 0xd1, 0x66,
	0xfd, 0x33, 0x05, 0xcf, 0x0f, 0x5f, 0x04, 0x00, 0x00, 0xff, 0xff, 0x80, 0x11, 0x01, 0x42, 0xdc,
	0x0a, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// MutateJobServiceClient is the client API for MutateJobService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MutateJobServiceClient interface {
	// Creates a mutate job.
	CreateMutateJob(ctx context.Context, in *CreateMutateJobRequest, opts ...grpc.CallOption) (*CreateMutateJobResponse, error)
	// Returns the mutate job.
	GetMutateJob(ctx context.Context, in *GetMutateJobRequest, opts ...grpc.CallOption) (*resources.MutateJob, error)
	// Returns the results of the mutate job. The job must be done.
	// Supports standard list paging.
	ListMutateJobResults(ctx context.Context, in *ListMutateJobResultsRequest, opts ...grpc.CallOption) (*ListMutateJobResultsResponse, error)
	// Runs the mutate job.
	//
	// The Operation.metadata field type is MutateJobMetadata. When finished, the
	// long running operation will not contain errors or a response. Instead, use
	// ListMutateJobResults to get the results of the job.
	RunMutateJob(ctx context.Context, in *RunMutateJobRequest, opts ...grpc.CallOption) (*longrunning.Operation, error)
	// Add operations to the mutate job.
	AddMutateJobOperations(ctx context.Context, in *AddMutateJobOperationsRequest, opts ...grpc.CallOption) (*AddMutateJobOperationsResponse, error)
}

type mutateJobServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMutateJobServiceClient(cc grpc.ClientConnInterface) MutateJobServiceClient {
	return &mutateJobServiceClient{cc}
}

func (c *mutateJobServiceClient) CreateMutateJob(ctx context.Context, in *CreateMutateJobRequest, opts ...grpc.CallOption) (*CreateMutateJobResponse, error) {
	out := new(CreateMutateJobResponse)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v1.services.MutateJobService/CreateMutateJob", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mutateJobServiceClient) GetMutateJob(ctx context.Context, in *GetMutateJobRequest, opts ...grpc.CallOption) (*resources.MutateJob, error) {
	out := new(resources.MutateJob)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v1.services.MutateJobService/GetMutateJob", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mutateJobServiceClient) ListMutateJobResults(ctx context.Context, in *ListMutateJobResultsRequest, opts ...grpc.CallOption) (*ListMutateJobResultsResponse, error) {
	out := new(ListMutateJobResultsResponse)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v1.services.MutateJobService/ListMutateJobResults", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mutateJobServiceClient) RunMutateJob(ctx context.Context, in *RunMutateJobRequest, opts ...grpc.CallOption) (*longrunning.Operation, error) {
	out := new(longrunning.Operation)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v1.services.MutateJobService/RunMutateJob", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mutateJobServiceClient) AddMutateJobOperations(ctx context.Context, in *AddMutateJobOperationsRequest, opts ...grpc.CallOption) (*AddMutateJobOperationsResponse, error) {
	out := new(AddMutateJobOperationsResponse)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v1.services.MutateJobService/AddMutateJobOperations", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MutateJobServiceServer is the server API for MutateJobService service.
type MutateJobServiceServer interface {
	// Creates a mutate job.
	CreateMutateJob(context.Context, *CreateMutateJobRequest) (*CreateMutateJobResponse, error)
	// Returns the mutate job.
	GetMutateJob(context.Context, *GetMutateJobRequest) (*resources.MutateJob, error)
	// Returns the results of the mutate job. The job must be done.
	// Supports standard list paging.
	ListMutateJobResults(context.Context, *ListMutateJobResultsRequest) (*ListMutateJobResultsResponse, error)
	// Runs the mutate job.
	//
	// The Operation.metadata field type is MutateJobMetadata. When finished, the
	// long running operation will not contain errors or a response. Instead, use
	// ListMutateJobResults to get the results of the job.
	RunMutateJob(context.Context, *RunMutateJobRequest) (*longrunning.Operation, error)
	// Add operations to the mutate job.
	AddMutateJobOperations(context.Context, *AddMutateJobOperationsRequest) (*AddMutateJobOperationsResponse, error)
}

// UnimplementedMutateJobServiceServer can be embedded to have forward compatible implementations.
type UnimplementedMutateJobServiceServer struct {
}

func (*UnimplementedMutateJobServiceServer) CreateMutateJob(ctx context.Context, req *CreateMutateJobRequest) (*CreateMutateJobResponse, error) {
	return nil, status1.Errorf(codes.Unimplemented, "method CreateMutateJob not implemented")
}
func (*UnimplementedMutateJobServiceServer) GetMutateJob(ctx context.Context, req *GetMutateJobRequest) (*resources.MutateJob, error) {
	return nil, status1.Errorf(codes.Unimplemented, "method GetMutateJob not implemented")
}
func (*UnimplementedMutateJobServiceServer) ListMutateJobResults(ctx context.Context, req *ListMutateJobResultsRequest) (*ListMutateJobResultsResponse, error) {
	return nil, status1.Errorf(codes.Unimplemented, "method ListMutateJobResults not implemented")
}
func (*UnimplementedMutateJobServiceServer) RunMutateJob(ctx context.Context, req *RunMutateJobRequest) (*longrunning.Operation, error) {
	return nil, status1.Errorf(codes.Unimplemented, "method RunMutateJob not implemented")
}
func (*UnimplementedMutateJobServiceServer) AddMutateJobOperations(ctx context.Context, req *AddMutateJobOperationsRequest) (*AddMutateJobOperationsResponse, error) {
	return nil, status1.Errorf(codes.Unimplemented, "method AddMutateJobOperations not implemented")
}

func RegisterMutateJobServiceServer(s *grpc.Server, srv MutateJobServiceServer) {
	s.RegisterService(&_MutateJobService_serviceDesc, srv)
}

func _MutateJobService_CreateMutateJob_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateMutateJobRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MutateJobServiceServer).CreateMutateJob(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v1.services.MutateJobService/CreateMutateJob",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MutateJobServiceServer).CreateMutateJob(ctx, req.(*CreateMutateJobRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MutateJobService_GetMutateJob_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMutateJobRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MutateJobServiceServer).GetMutateJob(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v1.services.MutateJobService/GetMutateJob",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MutateJobServiceServer).GetMutateJob(ctx, req.(*GetMutateJobRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MutateJobService_ListMutateJobResults_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListMutateJobResultsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MutateJobServiceServer).ListMutateJobResults(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v1.services.MutateJobService/ListMutateJobResults",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MutateJobServiceServer).ListMutateJobResults(ctx, req.(*ListMutateJobResultsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MutateJobService_RunMutateJob_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RunMutateJobRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MutateJobServiceServer).RunMutateJob(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v1.services.MutateJobService/RunMutateJob",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MutateJobServiceServer).RunMutateJob(ctx, req.(*RunMutateJobRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MutateJobService_AddMutateJobOperations_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddMutateJobOperationsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MutateJobServiceServer).AddMutateJobOperations(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v1.services.MutateJobService/AddMutateJobOperations",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MutateJobServiceServer).AddMutateJobOperations(ctx, req.(*AddMutateJobOperationsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _MutateJobService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v1.services.MutateJobService",
	HandlerType: (*MutateJobServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateMutateJob",
			Handler:    _MutateJobService_CreateMutateJob_Handler,
		},
		{
			MethodName: "GetMutateJob",
			Handler:    _MutateJobService_GetMutateJob_Handler,
		},
		{
			MethodName: "ListMutateJobResults",
			Handler:    _MutateJobService_ListMutateJobResults_Handler,
		},
		{
			MethodName: "RunMutateJob",
			Handler:    _MutateJobService_RunMutateJob_Handler,
		},
		{
			MethodName: "AddMutateJobOperations",
			Handler:    _MutateJobService_AddMutateJobOperations_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v1/services/mutate_job_service.proto",
}
