// Code generated by protoc-gen-go. DO NOT EDIT.
// source: feast/core/DatasetService.proto

package core // import "github.com/gojek/feast/protos/generated/go/feast/core"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import timestamp "github.com/golang/protobuf/ptypes/timestamp"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type DatasetServiceTypes struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DatasetServiceTypes) Reset()         { *m = DatasetServiceTypes{} }
func (m *DatasetServiceTypes) String() string { return proto.CompactTextString(m) }
func (*DatasetServiceTypes) ProtoMessage()    {}
func (*DatasetServiceTypes) Descriptor() ([]byte, []int) {
	return fileDescriptor_DatasetService_9ba96186e7cec93b, []int{0}
}
func (m *DatasetServiceTypes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DatasetServiceTypes.Unmarshal(m, b)
}
func (m *DatasetServiceTypes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DatasetServiceTypes.Marshal(b, m, deterministic)
}
func (dst *DatasetServiceTypes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DatasetServiceTypes.Merge(dst, src)
}
func (m *DatasetServiceTypes) XXX_Size() int {
	return xxx_messageInfo_DatasetServiceTypes.Size(m)
}
func (m *DatasetServiceTypes) XXX_DiscardUnknown() {
	xxx_messageInfo_DatasetServiceTypes.DiscardUnknown(m)
}

var xxx_messageInfo_DatasetServiceTypes proto.InternalMessageInfo

type DatasetServiceTypes_CreateDatasetRequest struct {
	// set of features for which its training data should be created
	FeatureSet *FeatureSet `protobuf:"bytes,1,opt,name=featureSet,proto3" json:"featureSet,omitempty"`
	// start date of the training data (inclusive)
	StartDate *timestamp.Timestamp `protobuf:"bytes,2,opt,name=startDate,proto3" json:"startDate,omitempty"`
	// end date of the training data (inclusive)
	EndDate *timestamp.Timestamp `protobuf:"bytes,3,opt,name=endDate,proto3" json:"endDate,omitempty"`
	// (optional) number of row that should be generated
	// (default: none)
	Limit int64 `protobuf:"varint,4,opt,name=limit,proto3" json:"limit,omitempty"`
	// (optional) prefix for dataset name
	NamePrefix           string   `protobuf:"bytes,5,opt,name=namePrefix,proto3" json:"namePrefix,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DatasetServiceTypes_CreateDatasetRequest) Reset() {
	*m = DatasetServiceTypes_CreateDatasetRequest{}
}
func (m *DatasetServiceTypes_CreateDatasetRequest) String() string { return proto.CompactTextString(m) }
func (*DatasetServiceTypes_CreateDatasetRequest) ProtoMessage()    {}
func (*DatasetServiceTypes_CreateDatasetRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_DatasetService_9ba96186e7cec93b, []int{0, 0}
}
func (m *DatasetServiceTypes_CreateDatasetRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DatasetServiceTypes_CreateDatasetRequest.Unmarshal(m, b)
}
func (m *DatasetServiceTypes_CreateDatasetRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DatasetServiceTypes_CreateDatasetRequest.Marshal(b, m, deterministic)
}
func (dst *DatasetServiceTypes_CreateDatasetRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DatasetServiceTypes_CreateDatasetRequest.Merge(dst, src)
}
func (m *DatasetServiceTypes_CreateDatasetRequest) XXX_Size() int {
	return xxx_messageInfo_DatasetServiceTypes_CreateDatasetRequest.Size(m)
}
func (m *DatasetServiceTypes_CreateDatasetRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DatasetServiceTypes_CreateDatasetRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DatasetServiceTypes_CreateDatasetRequest proto.InternalMessageInfo

func (m *DatasetServiceTypes_CreateDatasetRequest) GetFeatureSet() *FeatureSet {
	if m != nil {
		return m.FeatureSet
	}
	return nil
}

func (m *DatasetServiceTypes_CreateDatasetRequest) GetStartDate() *timestamp.Timestamp {
	if m != nil {
		return m.StartDate
	}
	return nil
}

func (m *DatasetServiceTypes_CreateDatasetRequest) GetEndDate() *timestamp.Timestamp {
	if m != nil {
		return m.EndDate
	}
	return nil
}

func (m *DatasetServiceTypes_CreateDatasetRequest) GetLimit() int64 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func (m *DatasetServiceTypes_CreateDatasetRequest) GetNamePrefix() string {
	if m != nil {
		return m.NamePrefix
	}
	return ""
}

type DatasetServiceTypes_CreateDatasetResponse struct {
	// information of the created training dataset
	DatasetInfo          *DatasetInfo `protobuf:"bytes,1,opt,name=datasetInfo,proto3" json:"datasetInfo,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *DatasetServiceTypes_CreateDatasetResponse) Reset() {
	*m = DatasetServiceTypes_CreateDatasetResponse{}
}
func (m *DatasetServiceTypes_CreateDatasetResponse) String() string { return proto.CompactTextString(m) }
func (*DatasetServiceTypes_CreateDatasetResponse) ProtoMessage()    {}
func (*DatasetServiceTypes_CreateDatasetResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_DatasetService_9ba96186e7cec93b, []int{0, 1}
}
func (m *DatasetServiceTypes_CreateDatasetResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DatasetServiceTypes_CreateDatasetResponse.Unmarshal(m, b)
}
func (m *DatasetServiceTypes_CreateDatasetResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DatasetServiceTypes_CreateDatasetResponse.Marshal(b, m, deterministic)
}
func (dst *DatasetServiceTypes_CreateDatasetResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DatasetServiceTypes_CreateDatasetResponse.Merge(dst, src)
}
func (m *DatasetServiceTypes_CreateDatasetResponse) XXX_Size() int {
	return xxx_messageInfo_DatasetServiceTypes_CreateDatasetResponse.Size(m)
}
func (m *DatasetServiceTypes_CreateDatasetResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DatasetServiceTypes_CreateDatasetResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DatasetServiceTypes_CreateDatasetResponse proto.InternalMessageInfo

func (m *DatasetServiceTypes_CreateDatasetResponse) GetDatasetInfo() *DatasetInfo {
	if m != nil {
		return m.DatasetInfo
	}
	return nil
}

// Represent a collection of feature having same entity name
type FeatureSet struct {
	// entity related to this feature set
	EntityName string `protobuf:"bytes,1,opt,name=entityName,proto3" json:"entityName,omitempty"`
	// list of feature id in this feature set
	FeatureIds           []string `protobuf:"bytes,2,rep,name=featureIds,proto3" json:"featureIds,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FeatureSet) Reset()         { *m = FeatureSet{} }
func (m *FeatureSet) String() string { return proto.CompactTextString(m) }
func (*FeatureSet) ProtoMessage()    {}
func (*FeatureSet) Descriptor() ([]byte, []int) {
	return fileDescriptor_DatasetService_9ba96186e7cec93b, []int{1}
}
func (m *FeatureSet) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FeatureSet.Unmarshal(m, b)
}
func (m *FeatureSet) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FeatureSet.Marshal(b, m, deterministic)
}
func (dst *FeatureSet) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FeatureSet.Merge(dst, src)
}
func (m *FeatureSet) XXX_Size() int {
	return xxx_messageInfo_FeatureSet.Size(m)
}
func (m *FeatureSet) XXX_DiscardUnknown() {
	xxx_messageInfo_FeatureSet.DiscardUnknown(m)
}

var xxx_messageInfo_FeatureSet proto.InternalMessageInfo

func (m *FeatureSet) GetEntityName() string {
	if m != nil {
		return m.EntityName
	}
	return ""
}

func (m *FeatureSet) GetFeatureIds() []string {
	if m != nil {
		return m.FeatureIds
	}
	return nil
}

// Representation of training dataset information
type DatasetInfo struct {
	// name of dataset
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// URL to table location of the training dataset
	TableUrl             string   `protobuf:"bytes,2,opt,name=tableUrl,proto3" json:"tableUrl,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DatasetInfo) Reset()         { *m = DatasetInfo{} }
func (m *DatasetInfo) String() string { return proto.CompactTextString(m) }
func (*DatasetInfo) ProtoMessage()    {}
func (*DatasetInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_DatasetService_9ba96186e7cec93b, []int{2}
}
func (m *DatasetInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DatasetInfo.Unmarshal(m, b)
}
func (m *DatasetInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DatasetInfo.Marshal(b, m, deterministic)
}
func (dst *DatasetInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DatasetInfo.Merge(dst, src)
}
func (m *DatasetInfo) XXX_Size() int {
	return xxx_messageInfo_DatasetInfo.Size(m)
}
func (m *DatasetInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_DatasetInfo.DiscardUnknown(m)
}

var xxx_messageInfo_DatasetInfo proto.InternalMessageInfo

func (m *DatasetInfo) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *DatasetInfo) GetTableUrl() string {
	if m != nil {
		return m.TableUrl
	}
	return ""
}

func init() {
	proto.RegisterType((*DatasetServiceTypes)(nil), "feast.core.DatasetServiceTypes")
	proto.RegisterType((*DatasetServiceTypes_CreateDatasetRequest)(nil), "feast.core.DatasetServiceTypes.CreateDatasetRequest")
	proto.RegisterType((*DatasetServiceTypes_CreateDatasetResponse)(nil), "feast.core.DatasetServiceTypes.CreateDatasetResponse")
	proto.RegisterType((*FeatureSet)(nil), "feast.core.FeatureSet")
	proto.RegisterType((*DatasetInfo)(nil), "feast.core.DatasetInfo")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// DatasetServiceClient is the client API for DatasetService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type DatasetServiceClient interface {
	// Create training dataset for a feature set
	CreateDataset(ctx context.Context, in *DatasetServiceTypes_CreateDatasetRequest, opts ...grpc.CallOption) (*DatasetServiceTypes_CreateDatasetResponse, error)
}

type datasetServiceClient struct {
	cc *grpc.ClientConn
}

func NewDatasetServiceClient(cc *grpc.ClientConn) DatasetServiceClient {
	return &datasetServiceClient{cc}
}

func (c *datasetServiceClient) CreateDataset(ctx context.Context, in *DatasetServiceTypes_CreateDatasetRequest, opts ...grpc.CallOption) (*DatasetServiceTypes_CreateDatasetResponse, error) {
	out := new(DatasetServiceTypes_CreateDatasetResponse)
	err := c.cc.Invoke(ctx, "/feast.core.DatasetService/CreateDataset", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DatasetServiceServer is the server API for DatasetService service.
type DatasetServiceServer interface {
	// Create training dataset for a feature set
	CreateDataset(context.Context, *DatasetServiceTypes_CreateDatasetRequest) (*DatasetServiceTypes_CreateDatasetResponse, error)
}

func RegisterDatasetServiceServer(s *grpc.Server, srv DatasetServiceServer) {
	s.RegisterService(&_DatasetService_serviceDesc, srv)
}

func _DatasetService_CreateDataset_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DatasetServiceTypes_CreateDatasetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DatasetServiceServer).CreateDataset(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/feast.core.DatasetService/CreateDataset",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DatasetServiceServer).CreateDataset(ctx, req.(*DatasetServiceTypes_CreateDatasetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _DatasetService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "feast.core.DatasetService",
	HandlerType: (*DatasetServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateDataset",
			Handler:    _DatasetService_CreateDataset_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "feast/core/DatasetService.proto",
}

func init() {
	proto.RegisterFile("feast/core/DatasetService.proto", fileDescriptor_DatasetService_9ba96186e7cec93b)
}

var fileDescriptor_DatasetService_9ba96186e7cec93b = []byte{
	// 414 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x53, 0xc1, 0x6e, 0xd4, 0x30,
	0x10, 0x25, 0xbb, 0x2d, 0x90, 0x59, 0xc1, 0xc1, 0x14, 0x88, 0x72, 0xa0, 0x51, 0x4e, 0x7b, 0xb2,
	0xa5, 0xd2, 0x22, 0x38, 0x70, 0x29, 0x2b, 0xa4, 0x4a, 0x08, 0x55, 0x6e, 0x91, 0x10, 0x37, 0x67,
	0x33, 0x09, 0x81, 0x24, 0x0e, 0xf6, 0x04, 0xd1, 0x0b, 0xdf, 0xc0, 0x27, 0xf2, 0x11, 0x7c, 0x00,
	0x8a, 0xb3, 0xdb, 0x78, 0xab, 0x4a, 0x88, 0x9b, 0x3d, 0xef, 0xcd, 0x78, 0xde, 0xf8, 0x0d, 0x1c,
	0x16, 0xa8, 0x2c, 0x89, 0xb5, 0x36, 0x28, 0x56, 0x8a, 0x94, 0x45, 0xba, 0x40, 0xf3, 0xbd, 0x5a,
	0x23, 0xef, 0x8c, 0x26, 0xcd, 0xc0, 0x11, 0xf8, 0x40, 0x88, 0x0f, 0x4b, 0xad, 0xcb, 0x1a, 0x85,
	0x43, 0xb2, 0xbe, 0x10, 0x54, 0x35, 0x68, 0x49, 0x35, 0xdd, 0x48, 0x4e, 0x7f, 0xcf, 0xe0, 0xd1,
	0x6e, 0x95, 0xcb, 0xab, 0x0e, 0x6d, 0xfc, 0x27, 0x80, 0x83, 0x37, 0x06, 0x15, 0xe1, 0x06, 0x95,
	0xf8, 0xad, 0x47, 0x4b, 0xec, 0x05, 0x0c, 0xf5, 0xa9, 0x37, 0x78, 0x81, 0x14, 0x05, 0x49, 0xb0,
	0x5c, 0x1c, 0x3d, 0xe1, 0xd3, 0x93, 0xfc, 0xed, 0x35, 0x2a, 0x3d, 0x26, 0x7b, 0x09, 0xa1, 0x25,
	0x65, 0x68, 0xa5, 0x08, 0xa3, 0x99, 0x4b, 0x8b, 0xf9, 0xd8, 0x1d, 0xdf, 0x76, 0xc7, 0x2f, 0xb7,
	0xdd, 0xc9, 0x89, 0xcc, 0x8e, 0xe1, 0x1e, 0xb6, 0xb9, 0xcb, 0x9b, 0xff, 0x33, 0x6f, 0x4b, 0x65,
	0x07, 0xb0, 0x5f, 0x57, 0x4d, 0x45, 0xd1, 0x5e, 0x12, 0x2c, 0xe7, 0x72, 0xbc, 0xb0, 0x67, 0x00,
	0xad, 0x6a, 0xf0, 0xdc, 0x60, 0x51, 0xfd, 0x88, 0xf6, 0x93, 0x60, 0x19, 0x4a, 0x2f, 0x12, 0x4b,
	0x78, 0x7c, 0x43, 0xb5, 0xed, 0x74, 0x6b, 0x91, 0xbd, 0x82, 0x45, 0x3e, 0x86, 0xce, 0xda, 0x42,
	0x6f, 0x74, 0x3f, 0xf5, 0x75, 0xaf, 0x26, 0x58, 0xfa, 0xdc, 0xf4, 0x1d, 0xc0, 0x34, 0x93, 0xa1,
	0x03, 0x6c, 0xa9, 0xa2, 0xab, 0xf7, 0xaa, 0x41, 0x57, 0x27, 0x94, 0x5e, 0x64, 0xc0, 0x37, 0x53,
	0x3b, 0xcb, 0x6d, 0x34, 0x4b, 0xe6, 0x03, 0x3e, 0x45, 0xd2, 0xd7, 0xb0, 0xf0, 0x5e, 0x62, 0x0c,
	0xf6, 0xda, 0xa9, 0x90, 0x3b, 0xb3, 0x18, 0xee, 0x93, 0xca, 0x6a, 0xfc, 0x60, 0x6a, 0x37, 0xe9,
	0x50, 0x5e, 0xdf, 0x8f, 0x7e, 0x05, 0xf0, 0x70, 0xf7, 0xbf, 0xd9, 0x4f, 0x78, 0xb0, 0xa3, 0x99,
	0x1d, 0xdf, 0x22, 0xcb, 0x37, 0x07, 0xbf, 0xcd, 0x18, 0xf1, 0xc9, 0x7f, 0x66, 0x8d, 0x83, 0x4d,
	0xef, 0x9c, 0x7e, 0x04, 0xcf, 0xb1, 0xa7, 0x37, 0xdc, 0x78, 0x3e, 0x7c, 0xf1, 0xa7, 0x93, 0xb2,
	0xa2, 0xcf, 0x7d, 0xc6, 0xd7, 0xba, 0x11, 0xa5, 0xfe, 0x82, 0x5f, 0xc5, 0xb8, 0x04, 0xce, 0x00,
	0x56, 0x94, 0xd8, 0xa2, 0x51, 0x84, 0xb9, 0x28, 0xb5, 0x98, 0xd6, 0x23, 0xbb, 0xeb, 0xf0, 0xe7,
	0x7f, 0x03, 0x00, 0x00, 0xff, 0xff, 0xbb, 0x9e, 0x4d, 0xe0, 0x33, 0x03, 0x00, 0x00,
}
