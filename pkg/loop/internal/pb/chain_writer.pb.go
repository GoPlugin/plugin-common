// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.25.1
// source: chain_writer.proto

package pb

import (
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

// TransactionStatus is an enum for the status of a transaction.
// This should always be a 1-1 mapping to: [github.com/goplugin/plugin-common/pkg/types.TransactionStatus].
type TransactionStatus int32

const (
	TransactionStatus_TRANSACTION_STATUS_UNKNOWN     TransactionStatus = 0
	TransactionStatus_TRANSACTION_STATUS_PENDING     TransactionStatus = 1
	TransactionStatus_TRANSACTION_STATUS_UNCONFIRMED TransactionStatus = 2
	TransactionStatus_TRANSACTION_STATUS_FINALIZED   TransactionStatus = 3
	TransactionStatus_TRANSACTION_STATUS_FAILED      TransactionStatus = 4
	TransactionStatus_TRANSACTION_STATUS_FATAL       TransactionStatus = 5
)

// Enum value maps for TransactionStatus.
var (
	TransactionStatus_name = map[int32]string{
		0: "TRANSACTION_STATUS_UNKNOWN",
		1: "TRANSACTION_STATUS_PENDING",
		2: "TRANSACTION_STATUS_UNCONFIRMED",
		3: "TRANSACTION_STATUS_FINALIZED",
		4: "TRANSACTION_STATUS_FAILED",
		5: "TRANSACTION_STATUS_FATAL",
	}
	TransactionStatus_value = map[string]int32{
		"TRANSACTION_STATUS_UNKNOWN":     0,
		"TRANSACTION_STATUS_PENDING":     1,
		"TRANSACTION_STATUS_UNCONFIRMED": 2,
		"TRANSACTION_STATUS_FINALIZED":   3,
		"TRANSACTION_STATUS_FAILED":      4,
		"TRANSACTION_STATUS_FATAL":       5,
	}
)

func (x TransactionStatus) Enum() *TransactionStatus {
	p := new(TransactionStatus)
	*p = x
	return p
}

func (x TransactionStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TransactionStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_chain_writer_proto_enumTypes[0].Descriptor()
}

func (TransactionStatus) Type() protoreflect.EnumType {
	return &file_chain_writer_proto_enumTypes[0]
}

func (x TransactionStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TransactionStatus.Descriptor instead.
func (TransactionStatus) EnumDescriptor() ([]byte, []int) {
	return file_chain_writer_proto_rawDescGZIP(), []int{0}
}

type SubmitTransactionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ContractName  string           `protobuf:"bytes,1,opt,name=contract_name,json=contractName,proto3" json:"contract_name,omitempty"`
	Method        string           `protobuf:"bytes,2,opt,name=method,proto3" json:"method,omitempty"`
	Params        *VersionedBytes  `protobuf:"bytes,3,opt,name=params,proto3" json:"params,omitempty"`
	TransactionId string           `protobuf:"bytes,4,opt,name=transaction_id,json=transactionId,proto3" json:"transaction_id,omitempty"`
	ToAddress     string           `protobuf:"bytes,5,opt,name=to_address,json=toAddress,proto3" json:"to_address,omitempty"`
	Meta          *TransactionMeta `protobuf:"bytes,6,opt,name=meta,proto3" json:"meta,omitempty"`
	Value         *BigInt          `protobuf:"bytes,7,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *SubmitTransactionRequest) Reset() {
	*x = SubmitTransactionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chain_writer_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubmitTransactionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubmitTransactionRequest) ProtoMessage() {}

func (x *SubmitTransactionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_chain_writer_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubmitTransactionRequest.ProtoReflect.Descriptor instead.
func (*SubmitTransactionRequest) Descriptor() ([]byte, []int) {
	return file_chain_writer_proto_rawDescGZIP(), []int{0}
}

func (x *SubmitTransactionRequest) GetContractName() string {
	if x != nil {
		return x.ContractName
	}
	return ""
}

func (x *SubmitTransactionRequest) GetMethod() string {
	if x != nil {
		return x.Method
	}
	return ""
}

func (x *SubmitTransactionRequest) GetParams() *VersionedBytes {
	if x != nil {
		return x.Params
	}
	return nil
}

func (x *SubmitTransactionRequest) GetTransactionId() string {
	if x != nil {
		return x.TransactionId
	}
	return ""
}

func (x *SubmitTransactionRequest) GetToAddress() string {
	if x != nil {
		return x.ToAddress
	}
	return ""
}

func (x *SubmitTransactionRequest) GetMeta() *TransactionMeta {
	if x != nil {
		return x.Meta
	}
	return nil
}

func (x *SubmitTransactionRequest) GetValue() *BigInt {
	if x != nil {
		return x.Value
	}
	return nil
}

type TransactionMeta struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	WorkflowExecutionId string  `protobuf:"bytes,1,opt,name=workflow_execution_id,json=workflowExecutionId,proto3" json:"workflow_execution_id,omitempty"`
	GasLimit            *BigInt `protobuf:"bytes,2,opt,name=gas_limit,json=gasLimit,proto3" json:"gas_limit,omitempty"`
}

func (x *TransactionMeta) Reset() {
	*x = TransactionMeta{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chain_writer_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TransactionMeta) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransactionMeta) ProtoMessage() {}

func (x *TransactionMeta) ProtoReflect() protoreflect.Message {
	mi := &file_chain_writer_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransactionMeta.ProtoReflect.Descriptor instead.
func (*TransactionMeta) Descriptor() ([]byte, []int) {
	return file_chain_writer_proto_rawDescGZIP(), []int{1}
}

func (x *TransactionMeta) GetWorkflowExecutionId() string {
	if x != nil {
		return x.WorkflowExecutionId
	}
	return ""
}

func (x *TransactionMeta) GetGasLimit() *BigInt {
	if x != nil {
		return x.GasLimit
	}
	return nil
}

// GetTransactionStatusRequest has arguments for [github.com/goplugin/plugin-common/pkg/types.ChainWriter.GetTransactionStatus].
type GetTransactionStatusRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TransactionId string `protobuf:"bytes,1,opt,name=transaction_id,json=transactionId,proto3" json:"transaction_id,omitempty"`
}

func (x *GetTransactionStatusRequest) Reset() {
	*x = GetTransactionStatusRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chain_writer_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTransactionStatusRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTransactionStatusRequest) ProtoMessage() {}

func (x *GetTransactionStatusRequest) ProtoReflect() protoreflect.Message {
	mi := &file_chain_writer_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTransactionStatusRequest.ProtoReflect.Descriptor instead.
func (*GetTransactionStatusRequest) Descriptor() ([]byte, []int) {
	return file_chain_writer_proto_rawDescGZIP(), []int{2}
}

func (x *GetTransactionStatusRequest) GetTransactionId() string {
	if x != nil {
		return x.TransactionId
	}
	return ""
}

// GetTransactionStatusReply has return arguments for [github.com/goplugin/plugin-common/pkg/types.ChainWriter.GetTransactionStatus].
type GetTransactionStatusReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TransactionStatus TransactionStatus `protobuf:"varint,1,opt,name=transaction_status,json=transactionStatus,proto3,enum=loop.TransactionStatus" json:"transaction_status,omitempty"`
}

func (x *GetTransactionStatusReply) Reset() {
	*x = GetTransactionStatusReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chain_writer_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTransactionStatusReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTransactionStatusReply) ProtoMessage() {}

func (x *GetTransactionStatusReply) ProtoReflect() protoreflect.Message {
	mi := &file_chain_writer_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTransactionStatusReply.ProtoReflect.Descriptor instead.
func (*GetTransactionStatusReply) Descriptor() ([]byte, []int) {
	return file_chain_writer_proto_rawDescGZIP(), []int{3}
}

func (x *GetTransactionStatusReply) GetTransactionStatus() TransactionStatus {
	if x != nil {
		return x.TransactionStatus
	}
	return TransactionStatus_TRANSACTION_STATUS_UNKNOWN
}

// GetFeeComponentsReply has return arguments for [github.com/goplugin/plugin-common/pkg/types.ChainWriter.GetFeeComponents].
type GetFeeComponentsReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ExecutionFee        *BigInt `protobuf:"bytes,1,opt,name=execution_fee,json=executionFee,proto3" json:"execution_fee,omitempty"`
	DataAvailabilityFee *BigInt `protobuf:"bytes,2,opt,name=data_availability_fee,json=dataAvailabilityFee,proto3" json:"data_availability_fee,omitempty"`
}

func (x *GetFeeComponentsReply) Reset() {
	*x = GetFeeComponentsReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chain_writer_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetFeeComponentsReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetFeeComponentsReply) ProtoMessage() {}

func (x *GetFeeComponentsReply) ProtoReflect() protoreflect.Message {
	mi := &file_chain_writer_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetFeeComponentsReply.ProtoReflect.Descriptor instead.
func (*GetFeeComponentsReply) Descriptor() ([]byte, []int) {
	return file_chain_writer_proto_rawDescGZIP(), []int{4}
}

func (x *GetFeeComponentsReply) GetExecutionFee() *BigInt {
	if x != nil {
		return x.ExecutionFee
	}
	return nil
}

func (x *GetFeeComponentsReply) GetDataAvailabilityFee() *BigInt {
	if x != nil {
		return x.DataAvailabilityFee
	}
	return nil
}

var File_chain_writer_proto protoreflect.FileDescriptor

var file_chain_writer_proto_rawDesc = []byte{
	0x0a, 0x12, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x5f, 0x77, 0x72, 0x69, 0x74, 0x65, 0x72, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x6c, 0x6f, 0x6f, 0x70, 0x1a, 0x0b, 0x63, 0x6f, 0x64, 0x65,
	0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0d, 0x72, 0x65, 0x6c, 0x61, 0x79, 0x65, 0x72,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x9a, 0x02, 0x0a, 0x18, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x54, 0x72,
	0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x23, 0x0a, 0x0d, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x5f, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63,
	0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x12, 0x2c, 0x0a,
	0x06, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e,
	0x6c, 0x6f, 0x6f, 0x70, 0x2e, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x65, 0x64, 0x42, 0x79,
	0x74, 0x65, 0x73, 0x52, 0x06, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x12, 0x25, 0x0a, 0x0e, 0x74,
	0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0d, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x74, 0x6f, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x74, 0x6f, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x12, 0x29, 0x0a, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x15, 0x2e, 0x6c, 0x6f, 0x6f, 0x70, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x4d, 0x65, 0x74, 0x61, 0x52, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x12, 0x22, 0x0a, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x6c, 0x6f,
	0x6f, 0x70, 0x2e, 0x42, 0x69, 0x67, 0x49, 0x6e, 0x74, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x22, 0x70, 0x0a, 0x0f, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4d,
	0x65, 0x74, 0x61, 0x12, 0x32, 0x0a, 0x15, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x5f,
	0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x13, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x45, 0x78, 0x65, 0x63,
	0x75, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x29, 0x0a, 0x09, 0x67, 0x61, 0x73, 0x5f, 0x6c,
	0x69, 0x6d, 0x69, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x6c, 0x6f, 0x6f,
	0x70, 0x2e, 0x42, 0x69, 0x67, 0x49, 0x6e, 0x74, 0x52, 0x08, 0x67, 0x61, 0x73, 0x4c, 0x69, 0x6d,
	0x69, 0x74, 0x22, 0x44, 0x0a, 0x1b, 0x47, 0x65, 0x74, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x25, 0x0a, 0x0e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x74, 0x72, 0x61, 0x6e, 0x73,
	0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x22, 0x63, 0x0a, 0x19, 0x47, 0x65, 0x74, 0x54,
	0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x46, 0x0a, 0x12, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x17, 0x2e, 0x6c, 0x6f, 0x6f, 0x70, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x11, 0x74, 0x72, 0x61, 0x6e,
	0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x8c, 0x01,
	0x0a, 0x15, 0x47, 0x65, 0x74, 0x46, 0x65, 0x65, 0x43, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e,
	0x74, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x31, 0x0a, 0x0d, 0x65, 0x78, 0x65, 0x63, 0x75,
	0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x66, 0x65, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c,
	0x2e, 0x6c, 0x6f, 0x6f, 0x70, 0x2e, 0x42, 0x69, 0x67, 0x49, 0x6e, 0x74, 0x52, 0x0c, 0x65, 0x78,
	0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x46, 0x65, 0x65, 0x12, 0x40, 0x0a, 0x15, 0x64, 0x61,
	0x74, 0x61, 0x5f, 0x61, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x5f,
	0x66, 0x65, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x6c, 0x6f, 0x6f, 0x70,
	0x2e, 0x42, 0x69, 0x67, 0x49, 0x6e, 0x74, 0x52, 0x13, 0x64, 0x61, 0x74, 0x61, 0x41, 0x76, 0x61,
	0x69, 0x6c, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x46, 0x65, 0x65, 0x2a, 0xd6, 0x01, 0x0a,
	0x11, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x12, 0x1e, 0x0a, 0x1a, 0x54, 0x52, 0x41, 0x4e, 0x53, 0x41, 0x43, 0x54, 0x49, 0x4f,
	0x4e, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e,
	0x10, 0x00, 0x12, 0x1e, 0x0a, 0x1a, 0x54, 0x52, 0x41, 0x4e, 0x53, 0x41, 0x43, 0x54, 0x49, 0x4f,
	0x4e, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x50, 0x45, 0x4e, 0x44, 0x49, 0x4e, 0x47,
	0x10, 0x01, 0x12, 0x22, 0x0a, 0x1e, 0x54, 0x52, 0x41, 0x4e, 0x53, 0x41, 0x43, 0x54, 0x49, 0x4f,
	0x4e, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x55, 0x4e, 0x43, 0x4f, 0x4e, 0x46, 0x49,
	0x52, 0x4d, 0x45, 0x44, 0x10, 0x02, 0x12, 0x20, 0x0a, 0x1c, 0x54, 0x52, 0x41, 0x4e, 0x53, 0x41,
	0x43, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x46, 0x49, 0x4e,
	0x41, 0x4c, 0x49, 0x5a, 0x45, 0x44, 0x10, 0x03, 0x12, 0x1d, 0x0a, 0x19, 0x54, 0x52, 0x41, 0x4e,
	0x53, 0x41, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x46,
	0x41, 0x49, 0x4c, 0x45, 0x44, 0x10, 0x04, 0x12, 0x1c, 0x0a, 0x18, 0x54, 0x52, 0x41, 0x4e, 0x53,
	0x41, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x46, 0x41,
	0x54, 0x41, 0x4c, 0x10, 0x05, 0x32, 0x85, 0x02, 0x0a, 0x0b, 0x43, 0x68, 0x61, 0x69, 0x6e, 0x57,
	0x72, 0x69, 0x74, 0x65, 0x72, 0x12, 0x4d, 0x0a, 0x11, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x54,
	0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1e, 0x2e, 0x6c, 0x6f, 0x6f,
	0x70, 0x2e, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x22, 0x00, 0x12, 0x5c, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x54, 0x72, 0x61, 0x6e, 0x73,
	0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x21, 0x2e, 0x6c,
	0x6f, 0x6f, 0x70, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x1f, 0x2e, 0x6c, 0x6f, 0x6f, 0x70, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x22, 0x00, 0x12, 0x49, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x46, 0x65, 0x65, 0x43, 0x6f, 0x6d, 0x70,
	0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x1b,
	0x2e, 0x6c, 0x6f, 0x6f, 0x70, 0x2e, 0x47, 0x65, 0x74, 0x46, 0x65, 0x65, 0x43, 0x6f, 0x6d, 0x70,
	0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x42, 0x43, 0x5a,
	0x41, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6d, 0x61, 0x72,
	0x74, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x6b, 0x69, 0x74, 0x2f, 0x63, 0x68, 0x61,
	0x69, 0x6e, 0x6c, 0x69, 0x6e, 0x6b, 0x2d, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x70, 0x6b,
	0x67, 0x2f, 0x6c, 0x6f, 0x6f, 0x70, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f,
	0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_chain_writer_proto_rawDescOnce sync.Once
	file_chain_writer_proto_rawDescData = file_chain_writer_proto_rawDesc
)

func file_chain_writer_proto_rawDescGZIP() []byte {
	file_chain_writer_proto_rawDescOnce.Do(func() {
		file_chain_writer_proto_rawDescData = protoimpl.X.CompressGZIP(file_chain_writer_proto_rawDescData)
	})
	return file_chain_writer_proto_rawDescData
}

var file_chain_writer_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_chain_writer_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_chain_writer_proto_goTypes = []interface{}{
	(TransactionStatus)(0),              // 0: loop.TransactionStatus
	(*SubmitTransactionRequest)(nil),    // 1: loop.SubmitTransactionRequest
	(*TransactionMeta)(nil),             // 2: loop.TransactionMeta
	(*GetTransactionStatusRequest)(nil), // 3: loop.GetTransactionStatusRequest
	(*GetTransactionStatusReply)(nil),   // 4: loop.GetTransactionStatusReply
	(*GetFeeComponentsReply)(nil),       // 5: loop.GetFeeComponentsReply
	(*VersionedBytes)(nil),              // 6: loop.VersionedBytes
	(*BigInt)(nil),                      // 7: loop.BigInt
	(*emptypb.Empty)(nil),               // 8: google.protobuf.Empty
}
var file_chain_writer_proto_depIdxs = []int32{
	6,  // 0: loop.SubmitTransactionRequest.params:type_name -> loop.VersionedBytes
	2,  // 1: loop.SubmitTransactionRequest.meta:type_name -> loop.TransactionMeta
	7,  // 2: loop.SubmitTransactionRequest.value:type_name -> loop.BigInt
	7,  // 3: loop.TransactionMeta.gas_limit:type_name -> loop.BigInt
	0,  // 4: loop.GetTransactionStatusReply.transaction_status:type_name -> loop.TransactionStatus
	7,  // 5: loop.GetFeeComponentsReply.execution_fee:type_name -> loop.BigInt
	7,  // 6: loop.GetFeeComponentsReply.data_availability_fee:type_name -> loop.BigInt
	1,  // 7: loop.ChainWriter.SubmitTransaction:input_type -> loop.SubmitTransactionRequest
	3,  // 8: loop.ChainWriter.GetTransactionStatus:input_type -> loop.GetTransactionStatusRequest
	8,  // 9: loop.ChainWriter.GetFeeComponents:input_type -> google.protobuf.Empty
	8,  // 10: loop.ChainWriter.SubmitTransaction:output_type -> google.protobuf.Empty
	4,  // 11: loop.ChainWriter.GetTransactionStatus:output_type -> loop.GetTransactionStatusReply
	5,  // 12: loop.ChainWriter.GetFeeComponents:output_type -> loop.GetFeeComponentsReply
	10, // [10:13] is the sub-list for method output_type
	7,  // [7:10] is the sub-list for method input_type
	7,  // [7:7] is the sub-list for extension type_name
	7,  // [7:7] is the sub-list for extension extendee
	0,  // [0:7] is the sub-list for field type_name
}

func init() { file_chain_writer_proto_init() }
func file_chain_writer_proto_init() {
	if File_chain_writer_proto != nil {
		return
	}
	file_codec_proto_init()
	file_relayer_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_chain_writer_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SubmitTransactionRequest); i {
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
		file_chain_writer_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TransactionMeta); i {
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
		file_chain_writer_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTransactionStatusRequest); i {
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
		file_chain_writer_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTransactionStatusReply); i {
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
		file_chain_writer_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetFeeComponentsReply); i {
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
			RawDescriptor: file_chain_writer_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_chain_writer_proto_goTypes,
		DependencyIndexes: file_chain_writer_proto_depIdxs,
		EnumInfos:         file_chain_writer_proto_enumTypes,
		MessageInfos:      file_chain_writer_proto_msgTypes,
	}.Build()
	File_chain_writer_proto = out.File
	file_chain_writer_proto_rawDesc = nil
	file_chain_writer_proto_goTypes = nil
	file_chain_writer_proto_depIdxs = nil
}