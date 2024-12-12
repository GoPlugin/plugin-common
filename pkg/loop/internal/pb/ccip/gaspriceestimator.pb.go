// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.25.1
// source: gaspriceestimator.proto

package ccippb

import (
	pb "github.com/goplugin/plugin-common/pkg/loop/internal/pb"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// GetGasPriceResponse returns the current gas price. It is a gRPC adpater for the return values of
// [github.com/goplugin/plugin-common/pkg/types/ccip.GasPriceEstimatorExec.GetGasPrice]
type GetGasPriceResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GasPrice *pb.BigInt `protobuf:"bytes,1,opt,name=gas_price,json=gasPrice,proto3" json:"gas_price,omitempty"`
}

func (x *GetGasPriceResponse) Reset() {
	*x = GetGasPriceResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gaspriceestimator_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetGasPriceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetGasPriceResponse) ProtoMessage() {}

func (x *GetGasPriceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_gaspriceestimator_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetGasPriceResponse.ProtoReflect.Descriptor instead.
func (*GetGasPriceResponse) Descriptor() ([]byte, []int) {
	return file_gaspriceestimator_proto_rawDescGZIP(), []int{0}
}

func (x *GetGasPriceResponse) GetGasPrice() *pb.BigInt {
	if x != nil {
		return x.GasPrice
	}
	return nil
}

// DenoteInUSDRequest is a gRPC adapter for the input values of
// [github.com/goplugin/plugin-common/pkg/types/ccip.GasPriceEstimatorExec.DenoteInUSD]
type DenoteInUSDRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	P                  *pb.BigInt `protobuf:"bytes,1,opt,name=p,proto3" json:"p,omitempty"`
	WrappedNativePrice *pb.BigInt `protobuf:"bytes,2,opt,name=wrapped_native_price,json=wrappedNativePrice,proto3" json:"wrapped_native_price,omitempty"`
}

func (x *DenoteInUSDRequest) Reset() {
	*x = DenoteInUSDRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gaspriceestimator_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DenoteInUSDRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DenoteInUSDRequest) ProtoMessage() {}

func (x *DenoteInUSDRequest) ProtoReflect() protoreflect.Message {
	mi := &file_gaspriceestimator_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DenoteInUSDRequest.ProtoReflect.Descriptor instead.
func (*DenoteInUSDRequest) Descriptor() ([]byte, []int) {
	return file_gaspriceestimator_proto_rawDescGZIP(), []int{1}
}

func (x *DenoteInUSDRequest) GetP() *pb.BigInt {
	if x != nil {
		return x.P
	}
	return nil
}

func (x *DenoteInUSDRequest) GetWrappedNativePrice() *pb.BigInt {
	if x != nil {
		return x.WrappedNativePrice
	}
	return nil
}

// DenoteInUSDResponse returns the price in USD. It is a gRPC adapter for the return values of
// [github.com/goplugin/plugin-common/pkg/types/ccip.GasPriceEstimatorExec.DenoteInUSD]
type DenoteInUSDResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UsdPrice *pb.BigInt `protobuf:"bytes,1,opt,name=usd_price,json=usdPrice,proto3" json:"usd_price,omitempty"`
}

func (x *DenoteInUSDResponse) Reset() {
	*x = DenoteInUSDResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gaspriceestimator_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DenoteInUSDResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DenoteInUSDResponse) ProtoMessage() {}

func (x *DenoteInUSDResponse) ProtoReflect() protoreflect.Message {
	mi := &file_gaspriceestimator_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DenoteInUSDResponse.ProtoReflect.Descriptor instead.
func (*DenoteInUSDResponse) Descriptor() ([]byte, []int) {
	return file_gaspriceestimator_proto_rawDescGZIP(), []int{2}
}

func (x *DenoteInUSDResponse) GetUsdPrice() *pb.BigInt {
	if x != nil {
		return x.UsdPrice
	}
	return nil
}

// EstimateMsgCostUSDRequest is a gRPC adapter for the input values of
// [github.com/goplugin/plugin-common/pkg/types/ccip.GasPriceEstimatorExec.EstimateMsgCostUSD]
type EstimateMsgCostUSDRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	P                  *pb.BigInt                              `protobuf:"bytes,1,opt,name=p,proto3" json:"p,omitempty"`
	WrappedNativePrice *pb.BigInt                              `protobuf:"bytes,2,opt,name=wrapped_native_price,json=wrappedNativePrice,proto3" json:"wrapped_native_price,omitempty"`
	Msg                *EVM2EVMOnRampCCIPSendRequestedWithMeta `protobuf:"bytes,3,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *EstimateMsgCostUSDRequest) Reset() {
	*x = EstimateMsgCostUSDRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gaspriceestimator_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EstimateMsgCostUSDRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EstimateMsgCostUSDRequest) ProtoMessage() {}

func (x *EstimateMsgCostUSDRequest) ProtoReflect() protoreflect.Message {
	mi := &file_gaspriceestimator_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EstimateMsgCostUSDRequest.ProtoReflect.Descriptor instead.
func (*EstimateMsgCostUSDRequest) Descriptor() ([]byte, []int) {
	return file_gaspriceestimator_proto_rawDescGZIP(), []int{3}
}

func (x *EstimateMsgCostUSDRequest) GetP() *pb.BigInt {
	if x != nil {
		return x.P
	}
	return nil
}

func (x *EstimateMsgCostUSDRequest) GetWrappedNativePrice() *pb.BigInt {
	if x != nil {
		return x.WrappedNativePrice
	}
	return nil
}

func (x *EstimateMsgCostUSDRequest) GetMsg() *EVM2EVMOnRampCCIPSendRequestedWithMeta {
	if x != nil {
		return x.Msg
	}
	return nil
}

// EstimateMsgCostUSDResponse returns the estimated cost in USD. It is a gRPC adapter for the return values of
// [github.com/goplugin/plugin-common/pkg/types/ccip.GasPriceEstimatorExec.EstimateMsgCostUSD]
type EstimateMsgCostUSDResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UsdCost *pb.BigInt `protobuf:"bytes,1,opt,name=usd_cost,json=usdCost,proto3" json:"usd_cost,omitempty"`
}

func (x *EstimateMsgCostUSDResponse) Reset() {
	*x = EstimateMsgCostUSDResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gaspriceestimator_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EstimateMsgCostUSDResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EstimateMsgCostUSDResponse) ProtoMessage() {}

func (x *EstimateMsgCostUSDResponse) ProtoReflect() protoreflect.Message {
	mi := &file_gaspriceestimator_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EstimateMsgCostUSDResponse.ProtoReflect.Descriptor instead.
func (*EstimateMsgCostUSDResponse) Descriptor() ([]byte, []int) {
	return file_gaspriceestimator_proto_rawDescGZIP(), []int{4}
}

func (x *EstimateMsgCostUSDResponse) GetUsdCost() *pb.BigInt {
	if x != nil {
		return x.UsdCost
	}
	return nil
}

// MedianRequest is a gRPC adapter for the input values of
// [github.com/goplugin/plugin-common/pkg/types/ccip.CommonGasPriceEstimator.Median]
type MedianRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GasPrices []*pb.BigInt `protobuf:"bytes,1,rep,name=gas_prices,json=gasPrices,proto3" json:"gas_prices,omitempty"`
}

func (x *MedianRequest) Reset() {
	*x = MedianRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gaspriceestimator_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MedianRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MedianRequest) ProtoMessage() {}

func (x *MedianRequest) ProtoReflect() protoreflect.Message {
	mi := &file_gaspriceestimator_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MedianRequest.ProtoReflect.Descriptor instead.
func (*MedianRequest) Descriptor() ([]byte, []int) {
	return file_gaspriceestimator_proto_rawDescGZIP(), []int{5}
}

func (x *MedianRequest) GetGasPrices() []*pb.BigInt {
	if x != nil {
		return x.GasPrices
	}
	return nil
}

// MedianResponse returns the median gas price. It is a gRPC adapter for the return values of
// [github.com/goplugin/plugin-common/pkg/types/ccip.CommonGasPriceEstimator.Median]
type MedianResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GasPrice *pb.BigInt `protobuf:"bytes,1,opt,name=gas_price,json=gasPrice,proto3" json:"gas_price,omitempty"`
}

func (x *MedianResponse) Reset() {
	*x = MedianResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gaspriceestimator_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MedianResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MedianResponse) ProtoMessage() {}

func (x *MedianResponse) ProtoReflect() protoreflect.Message {
	mi := &file_gaspriceestimator_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MedianResponse.ProtoReflect.Descriptor instead.
func (*MedianResponse) Descriptor() ([]byte, []int) {
	return file_gaspriceestimator_proto_rawDescGZIP(), []int{6}
}

func (x *MedianResponse) GetGasPrice() *pb.BigInt {
	if x != nil {
		return x.GasPrice
	}
	return nil
}

// DeviatesRequest is a gRPC adapter for the input values of
// [github.com/goplugin/plugin-common/pkg/types/ccip.GasPriceEstimatorCommit.Deviates]
type DeviatesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	P1 *pb.BigInt `protobuf:"bytes,1,opt,name=p1,proto3" json:"p1,omitempty"`
	P2 *pb.BigInt `protobuf:"bytes,2,opt,name=p2,proto3" json:"p2,omitempty"`
}

func (x *DeviatesRequest) Reset() {
	*x = DeviatesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gaspriceestimator_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeviatesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeviatesRequest) ProtoMessage() {}

func (x *DeviatesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_gaspriceestimator_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeviatesRequest.ProtoReflect.Descriptor instead.
func (*DeviatesRequest) Descriptor() ([]byte, []int) {
	return file_gaspriceestimator_proto_rawDescGZIP(), []int{7}
}

func (x *DeviatesRequest) GetP1() *pb.BigInt {
	if x != nil {
		return x.P1
	}
	return nil
}

func (x *DeviatesRequest) GetP2() *pb.BigInt {
	if x != nil {
		return x.P2
	}
	return nil
}

// DeviatesResponse returns the deviation between two gas prices. It is a gRPC adapter for the return values of
// [github.com/goplugin/plugin-common/pkg/types/ccip.GasPriceEstimatorCommit.Deviates]
type DeviatesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Deviates bool `protobuf:"varint,1,opt,name=deviates,proto3" json:"deviates,omitempty"`
}

func (x *DeviatesResponse) Reset() {
	*x = DeviatesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gaspriceestimator_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeviatesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeviatesResponse) ProtoMessage() {}

func (x *DeviatesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_gaspriceestimator_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeviatesResponse.ProtoReflect.Descriptor instead.
func (*DeviatesResponse) Descriptor() ([]byte, []int) {
	return file_gaspriceestimator_proto_rawDescGZIP(), []int{8}
}

func (x *DeviatesResponse) GetDeviates() bool {
	if x != nil {
		return x.Deviates
	}
	return false
}

var File_gaspriceestimator_proto protoreflect.FileDescriptor

var file_gaspriceestimator_proto_rawDesc = []byte{
	0x0a, 0x17, 0x67, 0x61, 0x73, 0x70, 0x72, 0x69, 0x63, 0x65, 0x65, 0x73, 0x74, 0x69, 0x6d, 0x61,
	0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x15, 0x6c, 0x6f, 0x6f, 0x70, 0x2e,
	0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x70, 0x62, 0x2e, 0x63, 0x63, 0x69, 0x70,
	0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0d, 0x72,
	0x65, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0c, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x40, 0x0a, 0x13, 0x47, 0x65,
	0x74, 0x47, 0x61, 0x73, 0x50, 0x72, 0x69, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x29, 0x0a, 0x09, 0x67, 0x61, 0x73, 0x5f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x6c, 0x6f, 0x6f, 0x70, 0x2e, 0x42, 0x69, 0x67, 0x49,
	0x6e, 0x74, 0x52, 0x08, 0x67, 0x61, 0x73, 0x50, 0x72, 0x69, 0x63, 0x65, 0x22, 0x70, 0x0a, 0x12,
	0x44, 0x65, 0x6e, 0x6f, 0x74, 0x65, 0x49, 0x6e, 0x55, 0x53, 0x44, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x1a, 0x0a, 0x01, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e,
	0x6c, 0x6f, 0x6f, 0x70, 0x2e, 0x42, 0x69, 0x67, 0x49, 0x6e, 0x74, 0x52, 0x01, 0x70, 0x12, 0x3e,
	0x0a, 0x14, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x64, 0x5f, 0x6e, 0x61, 0x74, 0x69, 0x76, 0x65,
	0x5f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x6c,
	0x6f, 0x6f, 0x70, 0x2e, 0x42, 0x69, 0x67, 0x49, 0x6e, 0x74, 0x52, 0x12, 0x77, 0x72, 0x61, 0x70,
	0x70, 0x65, 0x64, 0x4e, 0x61, 0x74, 0x69, 0x76, 0x65, 0x50, 0x72, 0x69, 0x63, 0x65, 0x22, 0x40,
	0x0a, 0x13, 0x44, 0x65, 0x6e, 0x6f, 0x74, 0x65, 0x49, 0x6e, 0x55, 0x53, 0x44, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x29, 0x0a, 0x09, 0x75, 0x73, 0x64, 0x5f, 0x70, 0x72, 0x69,
	0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x6c, 0x6f, 0x6f, 0x70, 0x2e,
	0x42, 0x69, 0x67, 0x49, 0x6e, 0x74, 0x52, 0x08, 0x75, 0x73, 0x64, 0x50, 0x72, 0x69, 0x63, 0x65,
	0x22, 0xc8, 0x01, 0x0a, 0x19, 0x45, 0x73, 0x74, 0x69, 0x6d, 0x61, 0x74, 0x65, 0x4d, 0x73, 0x67,
	0x43, 0x6f, 0x73, 0x74, 0x55, 0x53, 0x44, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a,
	0x0a, 0x01, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x6c, 0x6f, 0x6f, 0x70,
	0x2e, 0x42, 0x69, 0x67, 0x49, 0x6e, 0x74, 0x52, 0x01, 0x70, 0x12, 0x3e, 0x0a, 0x14, 0x77, 0x72,
	0x61, 0x70, 0x70, 0x65, 0x64, 0x5f, 0x6e, 0x61, 0x74, 0x69, 0x76, 0x65, 0x5f, 0x70, 0x72, 0x69,
	0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x6c, 0x6f, 0x6f, 0x70, 0x2e,
	0x42, 0x69, 0x67, 0x49, 0x6e, 0x74, 0x52, 0x12, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x64, 0x4e,
	0x61, 0x74, 0x69, 0x76, 0x65, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x4f, 0x0a, 0x03, 0x6d, 0x73,
	0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x3d, 0x2e, 0x6c, 0x6f, 0x6f, 0x70, 0x2e, 0x69,
	0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x70, 0x62, 0x2e, 0x63, 0x63, 0x69, 0x70, 0x2e,
	0x45, 0x56, 0x4d, 0x32, 0x45, 0x56, 0x4d, 0x4f, 0x6e, 0x52, 0x61, 0x6d, 0x70, 0x43, 0x43, 0x49,
	0x50, 0x53, 0x65, 0x6e, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x65, 0x64, 0x57, 0x69,
	0x74, 0x68, 0x4d, 0x65, 0x74, 0x61, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x22, 0x45, 0x0a, 0x1a, 0x45,
	0x73, 0x74, 0x69, 0x6d, 0x61, 0x74, 0x65, 0x4d, 0x73, 0x67, 0x43, 0x6f, 0x73, 0x74, 0x55, 0x53,
	0x44, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x27, 0x0a, 0x08, 0x75, 0x73, 0x64,
	0x5f, 0x63, 0x6f, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x6c, 0x6f,
	0x6f, 0x70, 0x2e, 0x42, 0x69, 0x67, 0x49, 0x6e, 0x74, 0x52, 0x07, 0x75, 0x73, 0x64, 0x43, 0x6f,
	0x73, 0x74, 0x22, 0x3c, 0x0a, 0x0d, 0x4d, 0x65, 0x64, 0x69, 0x61, 0x6e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x2b, 0x0a, 0x0a, 0x67, 0x61, 0x73, 0x5f, 0x70, 0x72, 0x69, 0x63, 0x65,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x6c, 0x6f, 0x6f, 0x70, 0x2e, 0x42,
	0x69, 0x67, 0x49, 0x6e, 0x74, 0x52, 0x09, 0x67, 0x61, 0x73, 0x50, 0x72, 0x69, 0x63, 0x65, 0x73,
	0x22, 0x3b, 0x0a, 0x0e, 0x4d, 0x65, 0x64, 0x69, 0x61, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x29, 0x0a, 0x09, 0x67, 0x61, 0x73, 0x5f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x6c, 0x6f, 0x6f, 0x70, 0x2e, 0x42, 0x69, 0x67,
	0x49, 0x6e, 0x74, 0x52, 0x08, 0x67, 0x61, 0x73, 0x50, 0x72, 0x69, 0x63, 0x65, 0x22, 0x4d, 0x0a,
	0x0f, 0x44, 0x65, 0x76, 0x69, 0x61, 0x74, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x1c, 0x0a, 0x02, 0x70, 0x31, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x6c,
	0x6f, 0x6f, 0x70, 0x2e, 0x42, 0x69, 0x67, 0x49, 0x6e, 0x74, 0x52, 0x02, 0x70, 0x31, 0x12, 0x1c,
	0x0a, 0x02, 0x70, 0x32, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x6c, 0x6f, 0x6f,
	0x70, 0x2e, 0x42, 0x69, 0x67, 0x49, 0x6e, 0x74, 0x52, 0x02, 0x70, 0x32, 0x22, 0x2e, 0x0a, 0x10,
	0x44, 0x65, 0x76, 0x69, 0x61, 0x74, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x1a, 0x0a, 0x08, 0x64, 0x65, 0x76, 0x69, 0x61, 0x74, 0x65, 0x73, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x08, 0x64, 0x65, 0x76, 0x69, 0x61, 0x74, 0x65, 0x73, 0x32, 0xa2, 0x03, 0x0a,
	0x15, 0x47, 0x61, 0x73, 0x50, 0x72, 0x69, 0x63, 0x65, 0x45, 0x73, 0x74, 0x69, 0x6d, 0x61, 0x74,
	0x6f, 0x72, 0x45, 0x78, 0x65, 0x63, 0x12, 0x51, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x47, 0x61, 0x73,
	0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x2a, 0x2e,
	0x6c, 0x6f, 0x6f, 0x70, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x70, 0x62,
	0x2e, 0x63, 0x63, 0x69, 0x70, 0x2e, 0x47, 0x65, 0x74, 0x47, 0x61, 0x73, 0x50, 0x72, 0x69, 0x63,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x64, 0x0a, 0x0b, 0x44, 0x65, 0x6e,
	0x6f, 0x74, 0x65, 0x49, 0x6e, 0x55, 0x53, 0x44, 0x12, 0x29, 0x2e, 0x6c, 0x6f, 0x6f, 0x70, 0x2e,
	0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x70, 0x62, 0x2e, 0x63, 0x63, 0x69, 0x70,
	0x2e, 0x44, 0x65, 0x6e, 0x6f, 0x74, 0x65, 0x49, 0x6e, 0x55, 0x53, 0x44, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x2a, 0x2e, 0x6c, 0x6f, 0x6f, 0x70, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72,
	0x6e, 0x61, 0x6c, 0x2e, 0x70, 0x62, 0x2e, 0x63, 0x63, 0x69, 0x70, 0x2e, 0x44, 0x65, 0x6e, 0x6f,
	0x74, 0x65, 0x49, 0x6e, 0x55, 0x53, 0x44, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x79, 0x0a, 0x12, 0x45, 0x73, 0x74, 0x69, 0x6d, 0x61, 0x74, 0x65, 0x4d, 0x73, 0x67, 0x43, 0x6f,
	0x73, 0x74, 0x55, 0x53, 0x44, 0x12, 0x30, 0x2e, 0x6c, 0x6f, 0x6f, 0x70, 0x2e, 0x69, 0x6e, 0x74,
	0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x70, 0x62, 0x2e, 0x63, 0x63, 0x69, 0x70, 0x2e, 0x45, 0x73,
	0x74, 0x69, 0x6d, 0x61, 0x74, 0x65, 0x4d, 0x73, 0x67, 0x43, 0x6f, 0x73, 0x74, 0x55, 0x53, 0x44,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x31, 0x2e, 0x6c, 0x6f, 0x6f, 0x70, 0x2e, 0x69,
	0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x70, 0x62, 0x2e, 0x63, 0x63, 0x69, 0x70, 0x2e,
	0x45, 0x73, 0x74, 0x69, 0x6d, 0x61, 0x74, 0x65, 0x4d, 0x73, 0x67, 0x43, 0x6f, 0x73, 0x74, 0x55,
	0x53, 0x44, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x55, 0x0a, 0x06, 0x4d, 0x65,
	0x64, 0x69, 0x61, 0x6e, 0x12, 0x24, 0x2e, 0x6c, 0x6f, 0x6f, 0x70, 0x2e, 0x69, 0x6e, 0x74, 0x65,
	0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x70, 0x62, 0x2e, 0x63, 0x63, 0x69, 0x70, 0x2e, 0x4d, 0x65, 0x64,
	0x69, 0x61, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e, 0x6c, 0x6f, 0x6f,
	0x70, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x70, 0x62, 0x2e, 0x63, 0x63,
	0x69, 0x70, 0x2e, 0x4d, 0x65, 0x64, 0x69, 0x61, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x32, 0x86, 0x03, 0x0a, 0x17, 0x47, 0x61, 0x73, 0x50, 0x72, 0x69, 0x63, 0x65, 0x45, 0x73,
	0x74, 0x69, 0x6d, 0x61, 0x74, 0x6f, 0x72, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x12, 0x51, 0x0a,
	0x0b, 0x47, 0x65, 0x74, 0x47, 0x61, 0x73, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x16, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x1a, 0x2a, 0x2e, 0x6c, 0x6f, 0x6f, 0x70, 0x2e, 0x69, 0x6e, 0x74, 0x65,
	0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x70, 0x62, 0x2e, 0x63, 0x63, 0x69, 0x70, 0x2e, 0x47, 0x65, 0x74,
	0x47, 0x61, 0x73, 0x50, 0x72, 0x69, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x64, 0x0a, 0x0b, 0x44, 0x65, 0x6e, 0x6f, 0x74, 0x65, 0x49, 0x6e, 0x55, 0x53, 0x44, 0x12,
	0x29, 0x2e, 0x6c, 0x6f, 0x6f, 0x70, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e,
	0x70, 0x62, 0x2e, 0x63, 0x63, 0x69, 0x70, 0x2e, 0x44, 0x65, 0x6e, 0x6f, 0x74, 0x65, 0x49, 0x6e,
	0x55, 0x53, 0x44, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2a, 0x2e, 0x6c, 0x6f, 0x6f,
	0x70, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x70, 0x62, 0x2e, 0x63, 0x63,
	0x69, 0x70, 0x2e, 0x44, 0x65, 0x6e, 0x6f, 0x74, 0x65, 0x49, 0x6e, 0x55, 0x53, 0x44, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x55, 0x0a, 0x06, 0x4d, 0x65, 0x64, 0x69, 0x61, 0x6e,
	0x12, 0x24, 0x2e, 0x6c, 0x6f, 0x6f, 0x70, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c,
	0x2e, 0x70, 0x62, 0x2e, 0x63, 0x63, 0x69, 0x70, 0x2e, 0x4d, 0x65, 0x64, 0x69, 0x61, 0x6e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e, 0x6c, 0x6f, 0x6f, 0x70, 0x2e, 0x69, 0x6e,
	0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x70, 0x62, 0x2e, 0x63, 0x63, 0x69, 0x70, 0x2e, 0x4d,
	0x65, 0x64, 0x69, 0x61, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x5b, 0x0a,
	0x08, 0x44, 0x65, 0x76, 0x69, 0x61, 0x74, 0x65, 0x73, 0x12, 0x26, 0x2e, 0x6c, 0x6f, 0x6f, 0x70,
	0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x70, 0x62, 0x2e, 0x63, 0x63, 0x69,
	0x70, 0x2e, 0x44, 0x65, 0x76, 0x69, 0x61, 0x74, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x27, 0x2e, 0x6c, 0x6f, 0x6f, 0x70, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61,
	0x6c, 0x2e, 0x70, 0x62, 0x2e, 0x63, 0x63, 0x69, 0x70, 0x2e, 0x44, 0x65, 0x76, 0x69, 0x61, 0x74,
	0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x4f, 0x5a, 0x4d, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6d, 0x61, 0x72, 0x74, 0x63, 0x6f,
	0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x6b, 0x69, 0x74, 0x2f, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x6c,
	0x69, 0x6e, 0x6b, 0x2d, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x6c,
	0x6f, 0x6f, 0x70, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x70, 0x62, 0x2f,
	0x63, 0x63, 0x69, 0x70, 0x3b, 0x63, 0x63, 0x69, 0x70, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_gaspriceestimator_proto_rawDescOnce sync.Once
	file_gaspriceestimator_proto_rawDescData = file_gaspriceestimator_proto_rawDesc
)

func file_gaspriceestimator_proto_rawDescGZIP() []byte {
	file_gaspriceestimator_proto_rawDescOnce.Do(func() {
		file_gaspriceestimator_proto_rawDescData = protoimpl.X.CompressGZIP(file_gaspriceestimator_proto_rawDescData)
	})
	return file_gaspriceestimator_proto_rawDescData
}

var file_gaspriceestimator_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_gaspriceestimator_proto_goTypes = []interface{}{
	(*GetGasPriceResponse)(nil),                    // 0: loop.internal.pb.ccip.GetGasPriceResponse
	(*DenoteInUSDRequest)(nil),                     // 1: loop.internal.pb.ccip.DenoteInUSDRequest
	(*DenoteInUSDResponse)(nil),                    // 2: loop.internal.pb.ccip.DenoteInUSDResponse
	(*EstimateMsgCostUSDRequest)(nil),              // 3: loop.internal.pb.ccip.EstimateMsgCostUSDRequest
	(*EstimateMsgCostUSDResponse)(nil),             // 4: loop.internal.pb.ccip.EstimateMsgCostUSDResponse
	(*MedianRequest)(nil),                          // 5: loop.internal.pb.ccip.MedianRequest
	(*MedianResponse)(nil),                         // 6: loop.internal.pb.ccip.MedianResponse
	(*DeviatesRequest)(nil),                        // 7: loop.internal.pb.ccip.DeviatesRequest
	(*DeviatesResponse)(nil),                       // 8: loop.internal.pb.ccip.DeviatesResponse
	(*pb.BigInt)(nil),                              // 9: loop.BigInt
	(*EVM2EVMOnRampCCIPSendRequestedWithMeta)(nil), // 10: loop.internal.pb.ccip.EVM2EVMOnRampCCIPSendRequestedWithMeta
	(*emptypb.Empty)(nil),                          // 11: google.protobuf.Empty
}
var file_gaspriceestimator_proto_depIdxs = []int32{
	9,  // 0: loop.internal.pb.ccip.GetGasPriceResponse.gas_price:type_name -> loop.BigInt
	9,  // 1: loop.internal.pb.ccip.DenoteInUSDRequest.p:type_name -> loop.BigInt
	9,  // 2: loop.internal.pb.ccip.DenoteInUSDRequest.wrapped_native_price:type_name -> loop.BigInt
	9,  // 3: loop.internal.pb.ccip.DenoteInUSDResponse.usd_price:type_name -> loop.BigInt
	9,  // 4: loop.internal.pb.ccip.EstimateMsgCostUSDRequest.p:type_name -> loop.BigInt
	9,  // 5: loop.internal.pb.ccip.EstimateMsgCostUSDRequest.wrapped_native_price:type_name -> loop.BigInt
	10, // 6: loop.internal.pb.ccip.EstimateMsgCostUSDRequest.msg:type_name -> loop.internal.pb.ccip.EVM2EVMOnRampCCIPSendRequestedWithMeta
	9,  // 7: loop.internal.pb.ccip.EstimateMsgCostUSDResponse.usd_cost:type_name -> loop.BigInt
	9,  // 8: loop.internal.pb.ccip.MedianRequest.gas_prices:type_name -> loop.BigInt
	9,  // 9: loop.internal.pb.ccip.MedianResponse.gas_price:type_name -> loop.BigInt
	9,  // 10: loop.internal.pb.ccip.DeviatesRequest.p1:type_name -> loop.BigInt
	9,  // 11: loop.internal.pb.ccip.DeviatesRequest.p2:type_name -> loop.BigInt
	11, // 12: loop.internal.pb.ccip.GasPriceEstimatorExec.GetGasPrice:input_type -> google.protobuf.Empty
	1,  // 13: loop.internal.pb.ccip.GasPriceEstimatorExec.DenoteInUSD:input_type -> loop.internal.pb.ccip.DenoteInUSDRequest
	3,  // 14: loop.internal.pb.ccip.GasPriceEstimatorExec.EstimateMsgCostUSD:input_type -> loop.internal.pb.ccip.EstimateMsgCostUSDRequest
	5,  // 15: loop.internal.pb.ccip.GasPriceEstimatorExec.Median:input_type -> loop.internal.pb.ccip.MedianRequest
	11, // 16: loop.internal.pb.ccip.GasPriceEstimatorCommit.GetGasPrice:input_type -> google.protobuf.Empty
	1,  // 17: loop.internal.pb.ccip.GasPriceEstimatorCommit.DenoteInUSD:input_type -> loop.internal.pb.ccip.DenoteInUSDRequest
	5,  // 18: loop.internal.pb.ccip.GasPriceEstimatorCommit.Median:input_type -> loop.internal.pb.ccip.MedianRequest
	7,  // 19: loop.internal.pb.ccip.GasPriceEstimatorCommit.Deviates:input_type -> loop.internal.pb.ccip.DeviatesRequest
	0,  // 20: loop.internal.pb.ccip.GasPriceEstimatorExec.GetGasPrice:output_type -> loop.internal.pb.ccip.GetGasPriceResponse
	2,  // 21: loop.internal.pb.ccip.GasPriceEstimatorExec.DenoteInUSD:output_type -> loop.internal.pb.ccip.DenoteInUSDResponse
	4,  // 22: loop.internal.pb.ccip.GasPriceEstimatorExec.EstimateMsgCostUSD:output_type -> loop.internal.pb.ccip.EstimateMsgCostUSDResponse
	6,  // 23: loop.internal.pb.ccip.GasPriceEstimatorExec.Median:output_type -> loop.internal.pb.ccip.MedianResponse
	0,  // 24: loop.internal.pb.ccip.GasPriceEstimatorCommit.GetGasPrice:output_type -> loop.internal.pb.ccip.GetGasPriceResponse
	2,  // 25: loop.internal.pb.ccip.GasPriceEstimatorCommit.DenoteInUSD:output_type -> loop.internal.pb.ccip.DenoteInUSDResponse
	6,  // 26: loop.internal.pb.ccip.GasPriceEstimatorCommit.Median:output_type -> loop.internal.pb.ccip.MedianResponse
	8,  // 27: loop.internal.pb.ccip.GasPriceEstimatorCommit.Deviates:output_type -> loop.internal.pb.ccip.DeviatesResponse
	20, // [20:28] is the sub-list for method output_type
	12, // [12:20] is the sub-list for method input_type
	12, // [12:12] is the sub-list for extension type_name
	12, // [12:12] is the sub-list for extension extendee
	0,  // [0:12] is the sub-list for field type_name
}

func init() { file_gaspriceestimator_proto_init() }
func file_gaspriceestimator_proto_init() {
	if File_gaspriceestimator_proto != nil {
		return
	}
	file_models_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_gaspriceestimator_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetGasPriceResponse); i {
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
		file_gaspriceestimator_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DenoteInUSDRequest); i {
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
		file_gaspriceestimator_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DenoteInUSDResponse); i {
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
		file_gaspriceestimator_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EstimateMsgCostUSDRequest); i {
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
		file_gaspriceestimator_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EstimateMsgCostUSDResponse); i {
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
		file_gaspriceestimator_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MedianRequest); i {
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
		file_gaspriceestimator_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MedianResponse); i {
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
		file_gaspriceestimator_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeviatesRequest); i {
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
		file_gaspriceestimator_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeviatesResponse); i {
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
			RawDescriptor: file_gaspriceestimator_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_gaspriceestimator_proto_goTypes,
		DependencyIndexes: file_gaspriceestimator_proto_depIdxs,
		MessageInfos:      file_gaspriceestimator_proto_msgTypes,
	}.Build()
	File_gaspriceestimator_proto = out.File
	file_gaspriceestimator_proto_rawDesc = nil
	file_gaspriceestimator_proto_goTypes = nil
	file_gaspriceestimator_proto_depIdxs = nil
}
