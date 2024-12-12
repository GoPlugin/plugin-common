// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.25.1
// source: oraclefactory.proto

package oraclefactory

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	durationpb "google.golang.org/protobuf/types/known/durationpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type LocalConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BlockchainTimeout                  *durationpb.Duration `protobuf:"bytes,1,opt,name=blockchain_timeout,json=blockchainTimeout,proto3" json:"blockchain_timeout,omitempty"`
	ContractConfigConfirmations        uint32               `protobuf:"varint,2,opt,name=contract_config_confirmations,json=contractConfigConfirmations,proto3" json:"contract_config_confirmations,omitempty"`
	SkipContractConfigConfirmations    bool                 `protobuf:"varint,3,opt,name=skip_contract_config_confirmations,json=skipContractConfigConfirmations,proto3" json:"skip_contract_config_confirmations,omitempty"`
	ContractConfigTrackerPollInterval  *durationpb.Duration `protobuf:"bytes,4,opt,name=contract_config_tracker_poll_interval,json=contractConfigTrackerPollInterval,proto3" json:"contract_config_tracker_poll_interval,omitempty"`
	ContractTransmitterTransmitTimeout *durationpb.Duration `protobuf:"bytes,5,opt,name=contract_transmitter_transmit_timeout,json=contractTransmitterTransmitTimeout,proto3" json:"contract_transmitter_transmit_timeout,omitempty"`
	DatabaseTimeout                    *durationpb.Duration `protobuf:"bytes,6,opt,name=database_timeout,json=databaseTimeout,proto3" json:"database_timeout,omitempty"`
	MinOcr2MaxDurationQuery            *durationpb.Duration `protobuf:"bytes,7,opt,name=min_ocr2_max_duration_query,json=minOcr2MaxDurationQuery,proto3" json:"min_ocr2_max_duration_query,omitempty"`
	DevelopmentMode                    string               `protobuf:"bytes,8,opt,name=development_mode,json=developmentMode,proto3" json:"development_mode,omitempty"`
}

func (x *LocalConfig) Reset() {
	*x = LocalConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_oraclefactory_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LocalConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LocalConfig) ProtoMessage() {}

func (x *LocalConfig) ProtoReflect() protoreflect.Message {
	mi := &file_oraclefactory_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LocalConfig.ProtoReflect.Descriptor instead.
func (*LocalConfig) Descriptor() ([]byte, []int) {
	return file_oraclefactory_proto_rawDescGZIP(), []int{0}
}

func (x *LocalConfig) GetBlockchainTimeout() *durationpb.Duration {
	if x != nil {
		return x.BlockchainTimeout
	}
	return nil
}

func (x *LocalConfig) GetContractConfigConfirmations() uint32 {
	if x != nil {
		return x.ContractConfigConfirmations
	}
	return 0
}

func (x *LocalConfig) GetSkipContractConfigConfirmations() bool {
	if x != nil {
		return x.SkipContractConfigConfirmations
	}
	return false
}

func (x *LocalConfig) GetContractConfigTrackerPollInterval() *durationpb.Duration {
	if x != nil {
		return x.ContractConfigTrackerPollInterval
	}
	return nil
}

func (x *LocalConfig) GetContractTransmitterTransmitTimeout() *durationpb.Duration {
	if x != nil {
		return x.ContractTransmitterTransmitTimeout
	}
	return nil
}

func (x *LocalConfig) GetDatabaseTimeout() *durationpb.Duration {
	if x != nil {
		return x.DatabaseTimeout
	}
	return nil
}

func (x *LocalConfig) GetMinOcr2MaxDurationQuery() *durationpb.Duration {
	if x != nil {
		return x.MinOcr2MaxDurationQuery
	}
	return nil
}

func (x *LocalConfig) GetDevelopmentMode() string {
	if x != nil {
		return x.DevelopmentMode
	}
	return ""
}

type NewOracleRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LocalConfig                     *LocalConfig `protobuf:"bytes,1,opt,name=local_config,json=localConfig,proto3" json:"local_config,omitempty"`
	ReportingPluginFactoryServiceId uint32       `protobuf:"varint,2,opt,name=reporting_plugin_factory_service_id,json=reportingPluginFactoryServiceId,proto3" json:"reporting_plugin_factory_service_id,omitempty"`
	ContractConfigTrackerId         uint32       `protobuf:"varint,3,opt,name=contract_config_tracker_id,json=contractConfigTrackerId,proto3" json:"contract_config_tracker_id,omitempty"`
	ContractTransmitterId           uint32       `protobuf:"varint,4,opt,name=contract_transmitter_id,json=contractTransmitterId,proto3" json:"contract_transmitter_id,omitempty"`
	OffchainConfigDigesterId        uint32       `protobuf:"varint,5,opt,name=offchain_config_digester_id,json=offchainConfigDigesterId,proto3" json:"offchain_config_digester_id,omitempty"`
}

func (x *NewOracleRequest) Reset() {
	*x = NewOracleRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_oraclefactory_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NewOracleRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NewOracleRequest) ProtoMessage() {}

func (x *NewOracleRequest) ProtoReflect() protoreflect.Message {
	mi := &file_oraclefactory_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NewOracleRequest.ProtoReflect.Descriptor instead.
func (*NewOracleRequest) Descriptor() ([]byte, []int) {
	return file_oraclefactory_proto_rawDescGZIP(), []int{1}
}

func (x *NewOracleRequest) GetLocalConfig() *LocalConfig {
	if x != nil {
		return x.LocalConfig
	}
	return nil
}

func (x *NewOracleRequest) GetReportingPluginFactoryServiceId() uint32 {
	if x != nil {
		return x.ReportingPluginFactoryServiceId
	}
	return 0
}

func (x *NewOracleRequest) GetContractConfigTrackerId() uint32 {
	if x != nil {
		return x.ContractConfigTrackerId
	}
	return 0
}

func (x *NewOracleRequest) GetContractTransmitterId() uint32 {
	if x != nil {
		return x.ContractTransmitterId
	}
	return 0
}

func (x *NewOracleRequest) GetOffchainConfigDigesterId() uint32 {
	if x != nil {
		return x.OffchainConfigDigesterId
	}
	return 0
}

type NewOracleReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OracleId uint32 `protobuf:"varint,1,opt,name=oracle_id,json=oracleId,proto3" json:"oracle_id,omitempty"`
}

func (x *NewOracleReply) Reset() {
	*x = NewOracleReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_oraclefactory_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NewOracleReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NewOracleReply) ProtoMessage() {}

func (x *NewOracleReply) ProtoReflect() protoreflect.Message {
	mi := &file_oraclefactory_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NewOracleReply.ProtoReflect.Descriptor instead.
func (*NewOracleReply) Descriptor() ([]byte, []int) {
	return file_oraclefactory_proto_rawDescGZIP(), []int{2}
}

func (x *NewOracleReply) GetOracleId() uint32 {
	if x != nil {
		return x.OracleId
	}
	return 0
}

var File_oraclefactory_proto protoreflect.FileDescriptor

var file_oraclefactory_proto_rawDesc = []byte{
	0x0a, 0x13, 0x6f, 0x72, 0x61, 0x63, 0x6c, 0x65, 0x66, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x79, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x6c, 0x6f, 0x6f, 0x70, 0x1a, 0x1e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x75, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8d, 0x05, 0x0a, 0x0b,
	0x4c, 0x6f, 0x63, 0x61, 0x6c, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x48, 0x0a, 0x12, 0x62,
	0x6c, 0x6f, 0x63, 0x6b, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75,
	0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x11, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x54, 0x69,
	0x6d, 0x65, 0x6f, 0x75, 0x74, 0x12, 0x42, 0x0a, 0x1d, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63,
	0x74, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x1b, 0x63, 0x6f,
	0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x4b, 0x0a, 0x22, 0x73, 0x6b, 0x69,
	0x70, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x1f, 0x73, 0x6b, 0x69, 0x70, 0x43, 0x6f, 0x6e, 0x74, 0x72,
	0x61, 0x63, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x6b, 0x0a, 0x25, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61,
	0x63, 0x74, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x5f, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x65,
	0x72, 0x5f, 0x70, 0x6f, 0x6c, 0x6c, 0x5f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x21, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x54, 0x72, 0x61, 0x63, 0x6b, 0x65, 0x72, 0x50, 0x6f, 0x6c, 0x6c, 0x49, 0x6e, 0x74, 0x65, 0x72,
	0x76, 0x61, 0x6c, 0x12, 0x6c, 0x0a, 0x25, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x5f,
	0x74, 0x72, 0x61, 0x6e, 0x73, 0x6d, 0x69, 0x74, 0x74, 0x65, 0x72, 0x5f, 0x74, 0x72, 0x61, 0x6e,
	0x73, 0x6d, 0x69, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x22, 0x63,
	0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x6d, 0x69, 0x74, 0x74,
	0x65, 0x72, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x6d, 0x69, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x6f, 0x75,
	0x74, 0x12, 0x44, 0x0a, 0x10, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x74, 0x69,
	0x6d, 0x65, 0x6f, 0x75, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0f, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65,
	0x54, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x12, 0x57, 0x0a, 0x1b, 0x6d, 0x69, 0x6e, 0x5f, 0x6f,
	0x63, 0x72, 0x32, 0x5f, 0x6d, 0x61, 0x78, 0x5f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x5f, 0x71, 0x75, 0x65, 0x72, 0x79, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44,
	0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x17, 0x6d, 0x69, 0x6e, 0x4f, 0x63, 0x72, 0x32,
	0x4d, 0x61, 0x78, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x51, 0x75, 0x65, 0x72, 0x79,
	0x12, 0x29, 0x0a, 0x10, 0x64, 0x65, 0x76, 0x65, 0x6c, 0x6f, 0x70, 0x6d, 0x65, 0x6e, 0x74, 0x5f,
	0x6d, 0x6f, 0x64, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x64, 0x65, 0x76, 0x65,
	0x6c, 0x6f, 0x70, 0x6d, 0x65, 0x6e, 0x74, 0x4d, 0x6f, 0x64, 0x65, 0x22, 0xca, 0x02, 0x0a, 0x10,
	0x4e, 0x65, 0x77, 0x4f, 0x72, 0x61, 0x63, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x34, 0x0a, 0x0c, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x6c, 0x6f, 0x6f, 0x70, 0x2e, 0x4c, 0x6f,
	0x63, 0x61, 0x6c, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x0b, 0x6c, 0x6f, 0x63, 0x61, 0x6c,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x4c, 0x0a, 0x23, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74,
	0x69, 0x6e, 0x67, 0x5f, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x5f, 0x66, 0x61, 0x63, 0x74, 0x6f,
	0x72, 0x79, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x1f, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x50, 0x6c,
	0x75, 0x67, 0x69, 0x6e, 0x46, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x79, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x49, 0x64, 0x12, 0x3b, 0x0a, 0x1a, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74,
	0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x5f, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x65, 0x72, 0x5f,
	0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x17, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61,
	0x63, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x54, 0x72, 0x61, 0x63, 0x6b, 0x65, 0x72, 0x49,
	0x64, 0x12, 0x36, 0x0a, 0x17, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x5f, 0x74, 0x72,
	0x61, 0x6e, 0x73, 0x6d, 0x69, 0x74, 0x74, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x15, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x54, 0x72, 0x61, 0x6e,
	0x73, 0x6d, 0x69, 0x74, 0x74, 0x65, 0x72, 0x49, 0x64, 0x12, 0x3d, 0x0a, 0x1b, 0x6f, 0x66, 0x66,
	0x63, 0x68, 0x61, 0x69, 0x6e, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x5f, 0x64, 0x69, 0x67,
	0x65, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x18,
	0x6f, 0x66, 0x66, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x44, 0x69,
	0x67, 0x65, 0x73, 0x74, 0x65, 0x72, 0x49, 0x64, 0x22, 0x2d, 0x0a, 0x0e, 0x4e, 0x65, 0x77, 0x4f,
	0x72, 0x61, 0x63, 0x6c, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x1b, 0x0a, 0x09, 0x6f, 0x72,
	0x61, 0x63, 0x6c, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x6f,
	0x72, 0x61, 0x63, 0x6c, 0x65, 0x49, 0x64, 0x32, 0x4c, 0x0a, 0x0d, 0x4f, 0x72, 0x61, 0x63, 0x6c,
	0x65, 0x46, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x79, 0x12, 0x3b, 0x0a, 0x09, 0x4e, 0x65, 0x77, 0x4f,
	0x72, 0x61, 0x63, 0x6c, 0x65, 0x12, 0x16, 0x2e, 0x6c, 0x6f, 0x6f, 0x70, 0x2e, 0x4e, 0x65, 0x77,
	0x4f, 0x72, 0x61, 0x63, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e,
	0x6c, 0x6f, 0x6f, 0x70, 0x2e, 0x4e, 0x65, 0x77, 0x4f, 0x72, 0x61, 0x63, 0x6c, 0x65, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x22, 0x00, 0x42, 0x4e, 0x5a, 0x4c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6d, 0x61, 0x72, 0x74, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63,
	0x74, 0x6b, 0x69, 0x74, 0x2f, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x6c, 0x69, 0x6e, 0x6b, 0x2d, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x6c, 0x6f, 0x6f, 0x70, 0x2f, 0x69,
	0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x6f, 0x72, 0x61, 0x63, 0x6c, 0x65, 0x66, 0x61,
	0x63, 0x74, 0x6f, 0x72, 0x79, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_oraclefactory_proto_rawDescOnce sync.Once
	file_oraclefactory_proto_rawDescData = file_oraclefactory_proto_rawDesc
)

func file_oraclefactory_proto_rawDescGZIP() []byte {
	file_oraclefactory_proto_rawDescOnce.Do(func() {
		file_oraclefactory_proto_rawDescData = protoimpl.X.CompressGZIP(file_oraclefactory_proto_rawDescData)
	})
	return file_oraclefactory_proto_rawDescData
}

var file_oraclefactory_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_oraclefactory_proto_goTypes = []interface{}{
	(*LocalConfig)(nil),         // 0: loop.LocalConfig
	(*NewOracleRequest)(nil),    // 1: loop.NewOracleRequest
	(*NewOracleReply)(nil),      // 2: loop.NewOracleReply
	(*durationpb.Duration)(nil), // 3: google.protobuf.Duration
}
var file_oraclefactory_proto_depIdxs = []int32{
	3, // 0: loop.LocalConfig.blockchain_timeout:type_name -> google.protobuf.Duration
	3, // 1: loop.LocalConfig.contract_config_tracker_poll_interval:type_name -> google.protobuf.Duration
	3, // 2: loop.LocalConfig.contract_transmitter_transmit_timeout:type_name -> google.protobuf.Duration
	3, // 3: loop.LocalConfig.database_timeout:type_name -> google.protobuf.Duration
	3, // 4: loop.LocalConfig.min_ocr2_max_duration_query:type_name -> google.protobuf.Duration
	0, // 5: loop.NewOracleRequest.local_config:type_name -> loop.LocalConfig
	1, // 6: loop.OracleFactory.NewOracle:input_type -> loop.NewOracleRequest
	2, // 7: loop.OracleFactory.NewOracle:output_type -> loop.NewOracleReply
	7, // [7:8] is the sub-list for method output_type
	6, // [6:7] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_oraclefactory_proto_init() }
func file_oraclefactory_proto_init() {
	if File_oraclefactory_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_oraclefactory_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LocalConfig); i {
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
		file_oraclefactory_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NewOracleRequest); i {
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
		file_oraclefactory_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NewOracleReply); i {
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
			RawDescriptor: file_oraclefactory_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_oraclefactory_proto_goTypes,
		DependencyIndexes: file_oraclefactory_proto_depIdxs,
		MessageInfos:      file_oraclefactory_proto_msgTypes,
	}.Build()
	File_oraclefactory_proto = out.File
	file_oraclefactory_proto_rawDesc = nil
	file_oraclefactory_proto_goTypes = nil
	file_oraclefactory_proto_depIdxs = nil
}