// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.25.1
// source: reportcodec_v3.proto

package mercuryv3pb

import (
	pb "github.com/goplugin/plugin-common/pkg/loop/internal/pb"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// ReportFields is a gRPC adapter for the ReportFields struct [pkg/types/mercury/v3/ReportFields].
type ReportFields struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ValidFromTimestamp uint32     `protobuf:"varint,1,opt,name=validFromTimestamp,proto3" json:"validFromTimestamp,omitempty"`
	Timestamp          uint32     `protobuf:"varint,2,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	NativeFee          *pb.BigInt `protobuf:"bytes,3,opt,name=nativeFee,proto3" json:"nativeFee,omitempty"`
	LinkFee            *pb.BigInt `protobuf:"bytes,4,opt,name=linkFee,proto3" json:"linkFee,omitempty"`
	ExpiresAt          uint32     `protobuf:"varint,5,opt,name=expiresAt,proto3" json:"expiresAt,omitempty"`
	BenchmarkPrice     *pb.BigInt `protobuf:"bytes,6,opt,name=benchmarkPrice,proto3" json:"benchmarkPrice,omitempty"`
	Bid                *pb.BigInt `protobuf:"bytes,7,opt,name=bid,proto3" json:"bid,omitempty"`
	Ask                *pb.BigInt `protobuf:"bytes,8,opt,name=ask,proto3" json:"ask,omitempty"`
}

func (x *ReportFields) Reset() {
	*x = ReportFields{}
	if protoimpl.UnsafeEnabled {
		mi := &file_reportcodec_v3_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReportFields) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReportFields) ProtoMessage() {}

func (x *ReportFields) ProtoReflect() protoreflect.Message {
	mi := &file_reportcodec_v3_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReportFields.ProtoReflect.Descriptor instead.
func (*ReportFields) Descriptor() ([]byte, []int) {
	return file_reportcodec_v3_proto_rawDescGZIP(), []int{0}
}

func (x *ReportFields) GetValidFromTimestamp() uint32 {
	if x != nil {
		return x.ValidFromTimestamp
	}
	return 0
}

func (x *ReportFields) GetTimestamp() uint32 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

func (x *ReportFields) GetNativeFee() *pb.BigInt {
	if x != nil {
		return x.NativeFee
	}
	return nil
}

func (x *ReportFields) GetLinkFee() *pb.BigInt {
	if x != nil {
		return x.LinkFee
	}
	return nil
}

func (x *ReportFields) GetExpiresAt() uint32 {
	if x != nil {
		return x.ExpiresAt
	}
	return 0
}

func (x *ReportFields) GetBenchmarkPrice() *pb.BigInt {
	if x != nil {
		return x.BenchmarkPrice
	}
	return nil
}

func (x *ReportFields) GetBid() *pb.BigInt {
	if x != nil {
		return x.Bid
	}
	return nil
}

func (x *ReportFields) GetAsk() *pb.BigInt {
	if x != nil {
		return x.Ask
	}
	return nil
}

// BuildReportRequest is gRPC adapter for the inputs arguments of [pkg/types/mercury/v3/ReportCodec.BuildReport].
type BuildReportRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ReportFields *ReportFields `protobuf:"bytes,1,opt,name=reportFields,proto3" json:"reportFields,omitempty"`
}

func (x *BuildReportRequest) Reset() {
	*x = BuildReportRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_reportcodec_v3_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BuildReportRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BuildReportRequest) ProtoMessage() {}

func (x *BuildReportRequest) ProtoReflect() protoreflect.Message {
	mi := &file_reportcodec_v3_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BuildReportRequest.ProtoReflect.Descriptor instead.
func (*BuildReportRequest) Descriptor() ([]byte, []int) {
	return file_reportcodec_v3_proto_rawDescGZIP(), []int{1}
}

func (x *BuildReportRequest) GetReportFields() *ReportFields {
	if x != nil {
		return x.ReportFields
	}
	return nil
}

// BuildReportReply is gRPC adapter for the return values of [pkg/types/mercury/v3/ReportCodec.BuildReport].
type BuildReportReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Report []byte `protobuf:"bytes,1,opt,name=report,proto3" json:"report,omitempty"`
}

func (x *BuildReportReply) Reset() {
	*x = BuildReportReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_reportcodec_v3_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BuildReportReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BuildReportReply) ProtoMessage() {}

func (x *BuildReportReply) ProtoReflect() protoreflect.Message {
	mi := &file_reportcodec_v3_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BuildReportReply.ProtoReflect.Descriptor instead.
func (*BuildReportReply) Descriptor() ([]byte, []int) {
	return file_reportcodec_v3_proto_rawDescGZIP(), []int{2}
}

func (x *BuildReportReply) GetReport() []byte {
	if x != nil {
		return x.Report
	}
	return nil
}

// MaxReportLengthRequest is gRPC adapter for the input arguments of [github.com/goplugin/plugin-data-streams/mercury/v3/ReportCodec.MaxReportLength].
type MaxReportLengthRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NumOracles uint64 `protobuf:"varint,1,opt,name=numOracles,proto3" json:"numOracles,omitempty"`
}

func (x *MaxReportLengthRequest) Reset() {
	*x = MaxReportLengthRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_reportcodec_v3_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MaxReportLengthRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MaxReportLengthRequest) ProtoMessage() {}

func (x *MaxReportLengthRequest) ProtoReflect() protoreflect.Message {
	mi := &file_reportcodec_v3_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MaxReportLengthRequest.ProtoReflect.Descriptor instead.
func (*MaxReportLengthRequest) Descriptor() ([]byte, []int) {
	return file_reportcodec_v3_proto_rawDescGZIP(), []int{3}
}

func (x *MaxReportLengthRequest) GetNumOracles() uint64 {
	if x != nil {
		return x.NumOracles
	}
	return 0
}

// MaxReportLengthReply is gRPC adapter for the return values of [github.com/goplugin/plugin-data-streams/mercury/v3/ReportCodec.MaxReportLength].
type MaxReportLengthReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MaxReportLength uint64 `protobuf:"varint,1,opt,name=maxReportLength,proto3" json:"maxReportLength,omitempty"`
}

func (x *MaxReportLengthReply) Reset() {
	*x = MaxReportLengthReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_reportcodec_v3_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MaxReportLengthReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MaxReportLengthReply) ProtoMessage() {}

func (x *MaxReportLengthReply) ProtoReflect() protoreflect.Message {
	mi := &file_reportcodec_v3_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MaxReportLengthReply.ProtoReflect.Descriptor instead.
func (*MaxReportLengthReply) Descriptor() ([]byte, []int) {
	return file_reportcodec_v3_proto_rawDescGZIP(), []int{4}
}

func (x *MaxReportLengthReply) GetMaxReportLength() uint64 {
	if x != nil {
		return x.MaxReportLength
	}
	return 0
}

// ObservationTimestampFromReportRequest is gRPC adapter for the input arguments [github.com/goplugin/plugin-data-streams/mercury/v3/ReportCodec.ObservationTimestampFromReport].
type ObservationTimestampFromReportRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Report []byte `protobuf:"bytes,1,opt,name=report,proto3" json:"report,omitempty"`
}

func (x *ObservationTimestampFromReportRequest) Reset() {
	*x = ObservationTimestampFromReportRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_reportcodec_v3_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ObservationTimestampFromReportRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ObservationTimestampFromReportRequest) ProtoMessage() {}

func (x *ObservationTimestampFromReportRequest) ProtoReflect() protoreflect.Message {
	mi := &file_reportcodec_v3_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ObservationTimestampFromReportRequest.ProtoReflect.Descriptor instead.
func (*ObservationTimestampFromReportRequest) Descriptor() ([]byte, []int) {
	return file_reportcodec_v3_proto_rawDescGZIP(), []int{5}
}

func (x *ObservationTimestampFromReportRequest) GetReport() []byte {
	if x != nil {
		return x.Report
	}
	return nil
}

// ObservationTimestampFromReportReply is gRPC adapter for the return values of [github.com/goplugin/plugin-data-streams/mercury/v3/ReportCodec.ObservationTimestampFromReport].
type ObservationTimestampFromReportReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Timestamp uint32 `protobuf:"varint,1,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
}

func (x *ObservationTimestampFromReportReply) Reset() {
	*x = ObservationTimestampFromReportReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_reportcodec_v3_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ObservationTimestampFromReportReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ObservationTimestampFromReportReply) ProtoMessage() {}

func (x *ObservationTimestampFromReportReply) ProtoReflect() protoreflect.Message {
	mi := &file_reportcodec_v3_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ObservationTimestampFromReportReply.ProtoReflect.Descriptor instead.
func (*ObservationTimestampFromReportReply) Descriptor() ([]byte, []int) {
	return file_reportcodec_v3_proto_rawDescGZIP(), []int{6}
}

func (x *ObservationTimestampFromReportReply) GetTimestamp() uint32 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

var File_reportcodec_v3_proto protoreflect.FileDescriptor

var file_reportcodec_v3_proto_rawDesc = []byte{
	0x0a, 0x14, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x63, 0x5f, 0x76, 0x33,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1b, 0x6c, 0x6f, 0x6f, 0x70, 0x2e, 0x69, 0x6e, 0x74,
	0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x70, 0x62, 0x2e, 0x6d, 0x65, 0x72, 0x63, 0x75, 0x72, 0x79,
	0x2e, 0x76, 0x33, 0x1a, 0x0d, 0x72, 0x65, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0xc4, 0x02, 0x0a, 0x0c, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x46, 0x69, 0x65,
	0x6c, 0x64, 0x73, 0x12, 0x2e, 0x0a, 0x12, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x46, 0x72, 0x6f, 0x6d,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x12, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x46, 0x72, 0x6f, 0x6d, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x12, 0x2a, 0x0a, 0x09, 0x6e, 0x61, 0x74, 0x69, 0x76, 0x65, 0x46, 0x65, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x6c, 0x6f, 0x6f, 0x70, 0x2e, 0x42, 0x69, 0x67, 0x49,
	0x6e, 0x74, 0x52, 0x09, 0x6e, 0x61, 0x74, 0x69, 0x76, 0x65, 0x46, 0x65, 0x65, 0x12, 0x26, 0x0a,
	0x07, 0x6c, 0x69, 0x6e, 0x6b, 0x46, 0x65, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c,
	0x2e, 0x6c, 0x6f, 0x6f, 0x70, 0x2e, 0x42, 0x69, 0x67, 0x49, 0x6e, 0x74, 0x52, 0x07, 0x6c, 0x69,
	0x6e, 0x6b, 0x46, 0x65, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x73,
	0x41, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65,
	0x73, 0x41, 0x74, 0x12, 0x34, 0x0a, 0x0e, 0x62, 0x65, 0x6e, 0x63, 0x68, 0x6d, 0x61, 0x72, 0x6b,
	0x50, 0x72, 0x69, 0x63, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x6c, 0x6f,
	0x6f, 0x70, 0x2e, 0x42, 0x69, 0x67, 0x49, 0x6e, 0x74, 0x52, 0x0e, 0x62, 0x65, 0x6e, 0x63, 0x68,
	0x6d, 0x61, 0x72, 0x6b, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x1e, 0x0a, 0x03, 0x62, 0x69, 0x64,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x6c, 0x6f, 0x6f, 0x70, 0x2e, 0x42, 0x69,
	0x67, 0x49, 0x6e, 0x74, 0x52, 0x03, 0x62, 0x69, 0x64, 0x12, 0x1e, 0x0a, 0x03, 0x61, 0x73, 0x6b,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x6c, 0x6f, 0x6f, 0x70, 0x2e, 0x42, 0x69,
	0x67, 0x49, 0x6e, 0x74, 0x52, 0x03, 0x61, 0x73, 0x6b, 0x22, 0x63, 0x0a, 0x12, 0x42, 0x75, 0x69,
	0x6c, 0x64, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x4d, 0x0a, 0x0c, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x6c, 0x6f, 0x6f, 0x70, 0x2e, 0x69, 0x6e, 0x74,
	0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x70, 0x62, 0x2e, 0x6d, 0x65, 0x72, 0x63, 0x75, 0x72, 0x79,
	0x2e, 0x76, 0x33, 0x2e, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x73,
	0x52, 0x0c, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x22, 0x2a,
	0x0a, 0x10, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x06, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x22, 0x38, 0x0a, 0x16, 0x4d, 0x61,
	0x78, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x4c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x6e, 0x75, 0x6d, 0x4f, 0x72, 0x61, 0x63, 0x6c,
	0x65, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0a, 0x6e, 0x75, 0x6d, 0x4f, 0x72, 0x61,
	0x63, 0x6c, 0x65, 0x73, 0x22, 0x40, 0x0a, 0x14, 0x4d, 0x61, 0x78, 0x52, 0x65, 0x70, 0x6f, 0x72,
	0x74, 0x4c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x28, 0x0a, 0x0f,
	0x6d, 0x61, 0x78, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x4c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0f, 0x6d, 0x61, 0x78, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74,
	0x4c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x22, 0x3f, 0x0a, 0x25, 0x4f, 0x62, 0x73, 0x65, 0x72, 0x76,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x46, 0x72,
	0x6f, 0x6d, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x16, 0x0a, 0x06, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x06, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x22, 0x43, 0x0a, 0x23, 0x4f, 0x62, 0x73, 0x65, 0x72,
	0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x46,
	0x72, 0x6f, 0x6d, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x1c,
	0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x32, 0xa6, 0x03, 0x0a,
	0x0b, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x43, 0x6f, 0x64, 0x65, 0x63, 0x12, 0x6f, 0x0a, 0x0b,
	0x42, 0x75, 0x69, 0x6c, 0x64, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x2f, 0x2e, 0x6c, 0x6f,
	0x6f, 0x70, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x70, 0x62, 0x2e, 0x6d,
	0x65, 0x72, 0x63, 0x75, 0x72, 0x79, 0x2e, 0x76, 0x33, 0x2e, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x52,
	0x65, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2d, 0x2e, 0x6c,
	0x6f, 0x6f, 0x70, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x70, 0x62, 0x2e,
	0x6d, 0x65, 0x72, 0x63, 0x75, 0x72, 0x79, 0x2e, 0x76, 0x33, 0x2e, 0x42, 0x75, 0x69, 0x6c, 0x64,
	0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x7b, 0x0a,
	0x0f, 0x4d, 0x61, 0x78, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x4c, 0x65, 0x6e, 0x67, 0x74, 0x68,
	0x12, 0x33, 0x2e, 0x6c, 0x6f, 0x6f, 0x70, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c,
	0x2e, 0x70, 0x62, 0x2e, 0x6d, 0x65, 0x72, 0x63, 0x75, 0x72, 0x79, 0x2e, 0x76, 0x33, 0x2e, 0x4d,
	0x61, 0x78, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x4c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x31, 0x2e, 0x6c, 0x6f, 0x6f, 0x70, 0x2e, 0x69, 0x6e, 0x74,
	0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x70, 0x62, 0x2e, 0x6d, 0x65, 0x72, 0x63, 0x75, 0x72, 0x79,
	0x2e, 0x76, 0x33, 0x2e, 0x4d, 0x61, 0x78, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x4c, 0x65, 0x6e,
	0x67, 0x74, 0x68, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0xa8, 0x01, 0x0a, 0x1e, 0x4f,
	0x62, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x46, 0x72, 0x6f, 0x6d, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x42, 0x2e,
	0x6c, 0x6f, 0x6f, 0x70, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x70, 0x62,
	0x2e, 0x6d, 0x65, 0x72, 0x63, 0x75, 0x72, 0x79, 0x2e, 0x76, 0x33, 0x2e, 0x4f, 0x62, 0x73, 0x65,
	0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x46, 0x72, 0x6f, 0x6d, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x40, 0x2e, 0x6c, 0x6f, 0x6f, 0x70, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61,
	0x6c, 0x2e, 0x70, 0x62, 0x2e, 0x6d, 0x65, 0x72, 0x63, 0x75, 0x72, 0x79, 0x2e, 0x76, 0x33, 0x2e,
	0x4f, 0x62, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x46, 0x72, 0x6f, 0x6d, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x22, 0x00, 0x42, 0x5a, 0x5a, 0x58, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6d, 0x61, 0x72, 0x74, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63,
	0x74, 0x6b, 0x69, 0x74, 0x2f, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x6c, 0x69, 0x6e, 0x6b, 0x2d, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x6c, 0x6f, 0x6f, 0x70, 0x2f, 0x69,
	0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x70, 0x62, 0x2f, 0x6d, 0x65, 0x72, 0x63, 0x75,
	0x72, 0x79, 0x2f, 0x76, 0x33, 0x3b, 0x6d, 0x65, 0x72, 0x63, 0x75, 0x72, 0x79, 0x76, 0x33, 0x70,
	0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_reportcodec_v3_proto_rawDescOnce sync.Once
	file_reportcodec_v3_proto_rawDescData = file_reportcodec_v3_proto_rawDesc
)

func file_reportcodec_v3_proto_rawDescGZIP() []byte {
	file_reportcodec_v3_proto_rawDescOnce.Do(func() {
		file_reportcodec_v3_proto_rawDescData = protoimpl.X.CompressGZIP(file_reportcodec_v3_proto_rawDescData)
	})
	return file_reportcodec_v3_proto_rawDescData
}

var file_reportcodec_v3_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_reportcodec_v3_proto_goTypes = []interface{}{
	(*ReportFields)(nil),                          // 0: loop.internal.pb.mercury.v3.ReportFields
	(*BuildReportRequest)(nil),                    // 1: loop.internal.pb.mercury.v3.BuildReportRequest
	(*BuildReportReply)(nil),                      // 2: loop.internal.pb.mercury.v3.BuildReportReply
	(*MaxReportLengthRequest)(nil),                // 3: loop.internal.pb.mercury.v3.MaxReportLengthRequest
	(*MaxReportLengthReply)(nil),                  // 4: loop.internal.pb.mercury.v3.MaxReportLengthReply
	(*ObservationTimestampFromReportRequest)(nil), // 5: loop.internal.pb.mercury.v3.ObservationTimestampFromReportRequest
	(*ObservationTimestampFromReportReply)(nil),   // 6: loop.internal.pb.mercury.v3.ObservationTimestampFromReportReply
	(*pb.BigInt)(nil),                             // 7: loop.BigInt
}
var file_reportcodec_v3_proto_depIdxs = []int32{
	7, // 0: loop.internal.pb.mercury.v3.ReportFields.nativeFee:type_name -> loop.BigInt
	7, // 1: loop.internal.pb.mercury.v3.ReportFields.linkFee:type_name -> loop.BigInt
	7, // 2: loop.internal.pb.mercury.v3.ReportFields.benchmarkPrice:type_name -> loop.BigInt
	7, // 3: loop.internal.pb.mercury.v3.ReportFields.bid:type_name -> loop.BigInt
	7, // 4: loop.internal.pb.mercury.v3.ReportFields.ask:type_name -> loop.BigInt
	0, // 5: loop.internal.pb.mercury.v3.BuildReportRequest.reportFields:type_name -> loop.internal.pb.mercury.v3.ReportFields
	1, // 6: loop.internal.pb.mercury.v3.ReportCodec.BuildReport:input_type -> loop.internal.pb.mercury.v3.BuildReportRequest
	3, // 7: loop.internal.pb.mercury.v3.ReportCodec.MaxReportLength:input_type -> loop.internal.pb.mercury.v3.MaxReportLengthRequest
	5, // 8: loop.internal.pb.mercury.v3.ReportCodec.ObservationTimestampFromReport:input_type -> loop.internal.pb.mercury.v3.ObservationTimestampFromReportRequest
	2, // 9: loop.internal.pb.mercury.v3.ReportCodec.BuildReport:output_type -> loop.internal.pb.mercury.v3.BuildReportReply
	4, // 10: loop.internal.pb.mercury.v3.ReportCodec.MaxReportLength:output_type -> loop.internal.pb.mercury.v3.MaxReportLengthReply
	6, // 11: loop.internal.pb.mercury.v3.ReportCodec.ObservationTimestampFromReport:output_type -> loop.internal.pb.mercury.v3.ObservationTimestampFromReportReply
	9, // [9:12] is the sub-list for method output_type
	6, // [6:9] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_reportcodec_v3_proto_init() }
func file_reportcodec_v3_proto_init() {
	if File_reportcodec_v3_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_reportcodec_v3_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReportFields); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_reportcodec_v3_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BuildReportRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_reportcodec_v3_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BuildReportReply); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_reportcodec_v3_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MaxReportLengthRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_reportcodec_v3_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MaxReportLengthReply); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_reportcodec_v3_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ObservationTimestampFromReportRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_reportcodec_v3_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ObservationTimestampFromReportReply); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_reportcodec_v3_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_reportcodec_v3_proto_goTypes,
		DependencyIndexes: file_reportcodec_v3_proto_depIdxs,
		MessageInfos:      file_reportcodec_v3_proto_msgTypes,
	}.Build()
	File_reportcodec_v3_proto = out.File
	file_reportcodec_v3_proto_rawDesc = nil
	file_reportcodec_v3_proto_goTypes = nil
	file_reportcodec_v3_proto_depIdxs = nil
}