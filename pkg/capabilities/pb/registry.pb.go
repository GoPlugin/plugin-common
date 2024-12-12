// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.25.1
// source: capabilities/pb/registry.proto

package pb

import (
	pb "github.com/goplugin/plugin-common/pkg/values/pb"
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

type RemoteTriggerConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RegistrationRefresh     *durationpb.Duration `protobuf:"bytes,1,opt,name=registrationRefresh,proto3" json:"registrationRefresh,omitempty"`
	RegistrationExpiry      *durationpb.Duration `protobuf:"bytes,2,opt,name=registrationExpiry,proto3" json:"registrationExpiry,omitempty"`
	MinResponsesToAggregate uint32               `protobuf:"varint,3,opt,name=minResponsesToAggregate,proto3" json:"minResponsesToAggregate,omitempty"`
	MessageExpiry           *durationpb.Duration `protobuf:"bytes,4,opt,name=messageExpiry,proto3" json:"messageExpiry,omitempty"`
	MaxBatchSize            uint32               `protobuf:"varint,5,opt,name=maxBatchSize,proto3" json:"maxBatchSize,omitempty"`
	BatchCollectionPeriod   *durationpb.Duration `protobuf:"bytes,6,opt,name=batchCollectionPeriod,proto3" json:"batchCollectionPeriod,omitempty"`
}

func (x *RemoteTriggerConfig) Reset() {
	*x = RemoteTriggerConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_capabilities_pb_registry_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoteTriggerConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoteTriggerConfig) ProtoMessage() {}

func (x *RemoteTriggerConfig) ProtoReflect() protoreflect.Message {
	mi := &file_capabilities_pb_registry_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoteTriggerConfig.ProtoReflect.Descriptor instead.
func (*RemoteTriggerConfig) Descriptor() ([]byte, []int) {
	return file_capabilities_pb_registry_proto_rawDescGZIP(), []int{0}
}

func (x *RemoteTriggerConfig) GetRegistrationRefresh() *durationpb.Duration {
	if x != nil {
		return x.RegistrationRefresh
	}
	return nil
}

func (x *RemoteTriggerConfig) GetRegistrationExpiry() *durationpb.Duration {
	if x != nil {
		return x.RegistrationExpiry
	}
	return nil
}

func (x *RemoteTriggerConfig) GetMinResponsesToAggregate() uint32 {
	if x != nil {
		return x.MinResponsesToAggregate
	}
	return 0
}

func (x *RemoteTriggerConfig) GetMessageExpiry() *durationpb.Duration {
	if x != nil {
		return x.MessageExpiry
	}
	return nil
}

func (x *RemoteTriggerConfig) GetMaxBatchSize() uint32 {
	if x != nil {
		return x.MaxBatchSize
	}
	return 0
}

func (x *RemoteTriggerConfig) GetBatchCollectionPeriod() *durationpb.Duration {
	if x != nil {
		return x.BatchCollectionPeriod
	}
	return nil
}

type RemoteTargetConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// A collection of dot seperated paths to attributes that should be excluded from the request sent to the remote target
	// when calculating the hash of the request.  This is useful for excluding attributes that are not deterministic to ensure
	// that the hash of logically identical requests is consistent.
	RequestHashExcludedAttributes []string `protobuf:"bytes,1,rep,name=requestHashExcludedAttributes,proto3" json:"requestHashExcludedAttributes,omitempty"`
}

func (x *RemoteTargetConfig) Reset() {
	*x = RemoteTargetConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_capabilities_pb_registry_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoteTargetConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoteTargetConfig) ProtoMessage() {}

func (x *RemoteTargetConfig) ProtoReflect() protoreflect.Message {
	mi := &file_capabilities_pb_registry_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoteTargetConfig.ProtoReflect.Descriptor instead.
func (*RemoteTargetConfig) Descriptor() ([]byte, []int) {
	return file_capabilities_pb_registry_proto_rawDescGZIP(), []int{1}
}

func (x *RemoteTargetConfig) GetRequestHashExcludedAttributes() []string {
	if x != nil {
		return x.RequestHashExcludedAttributes
	}
	return nil
}

type CapabilityConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DefaultConfig *pb.Map `protobuf:"bytes,1,opt,name=default_config,json=defaultConfig,proto3" json:"default_config,omitempty"`
	// Types that are assignable to RemoteConfig:
	//
	//	*CapabilityConfig_RemoteTriggerConfig
	//	*CapabilityConfig_RemoteTargetConfig
	RemoteConfig isCapabilityConfig_RemoteConfig `protobuf_oneof:"remote_config"`
}

func (x *CapabilityConfig) Reset() {
	*x = CapabilityConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_capabilities_pb_registry_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CapabilityConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CapabilityConfig) ProtoMessage() {}

func (x *CapabilityConfig) ProtoReflect() protoreflect.Message {
	mi := &file_capabilities_pb_registry_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CapabilityConfig.ProtoReflect.Descriptor instead.
func (*CapabilityConfig) Descriptor() ([]byte, []int) {
	return file_capabilities_pb_registry_proto_rawDescGZIP(), []int{2}
}

func (x *CapabilityConfig) GetDefaultConfig() *pb.Map {
	if x != nil {
		return x.DefaultConfig
	}
	return nil
}

func (m *CapabilityConfig) GetRemoteConfig() isCapabilityConfig_RemoteConfig {
	if m != nil {
		return m.RemoteConfig
	}
	return nil
}

func (x *CapabilityConfig) GetRemoteTriggerConfig() *RemoteTriggerConfig {
	if x, ok := x.GetRemoteConfig().(*CapabilityConfig_RemoteTriggerConfig); ok {
		return x.RemoteTriggerConfig
	}
	return nil
}

func (x *CapabilityConfig) GetRemoteTargetConfig() *RemoteTargetConfig {
	if x, ok := x.GetRemoteConfig().(*CapabilityConfig_RemoteTargetConfig); ok {
		return x.RemoteTargetConfig
	}
	return nil
}

type isCapabilityConfig_RemoteConfig interface {
	isCapabilityConfig_RemoteConfig()
}

type CapabilityConfig_RemoteTriggerConfig struct {
	RemoteTriggerConfig *RemoteTriggerConfig `protobuf:"bytes,2,opt,name=remote_trigger_config,json=remoteTriggerConfig,proto3,oneof"`
}

type CapabilityConfig_RemoteTargetConfig struct {
	RemoteTargetConfig *RemoteTargetConfig `protobuf:"bytes,3,opt,name=remote_target_config,json=remoteTargetConfig,proto3,oneof"`
}

func (*CapabilityConfig_RemoteTriggerConfig) isCapabilityConfig_RemoteConfig() {}

func (*CapabilityConfig_RemoteTargetConfig) isCapabilityConfig_RemoteConfig() {}

var File_capabilities_pb_registry_proto protoreflect.FileDescriptor

var file_capabilities_pb_registry_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x63, 0x61, 0x70, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2f, 0x70,
	0x62, 0x2f, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x04, 0x6c, 0x6f, 0x6f, 0x70, 0x1a, 0x16, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x2f, 0x70,
	0x62, 0x2f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x9d,
	0x03, 0x0a, 0x13, 0x52, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x54, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x4b, 0x0a, 0x13, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x13,
	0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x66, 0x72,
	0x65, 0x73, 0x68, 0x12, 0x49, 0x0a, 0x12, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x45, 0x78, 0x70, 0x69, 0x72, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x12, 0x72, 0x65, 0x67, 0x69,
	0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x45, 0x78, 0x70, 0x69, 0x72, 0x79, 0x12, 0x38,
	0x0a, 0x17, 0x6d, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x73, 0x54, 0x6f,
	0x41, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x17, 0x6d, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x73, 0x54, 0x6f, 0x41,
	0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x12, 0x3f, 0x0a, 0x0d, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x45, 0x78, 0x70, 0x69, 0x72, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0d, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x45, 0x78, 0x70, 0x69, 0x72, 0x79, 0x12, 0x22, 0x0a, 0x0c, 0x6d, 0x61, 0x78,
	0x42, 0x61, 0x74, 0x63, 0x68, 0x53, 0x69, 0x7a, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x0c, 0x6d, 0x61, 0x78, 0x42, 0x61, 0x74, 0x63, 0x68, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x4f, 0x0a,
	0x15, 0x62, 0x61, 0x74, 0x63, 0x68, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x50, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44,
	0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x15, 0x62, 0x61, 0x74, 0x63, 0x68, 0x43, 0x6f,
	0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x22, 0x5a,
	0x0a, 0x12, 0x52, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x12, 0x44, 0x0a, 0x1d, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x48,
	0x61, 0x73, 0x68, 0x45, 0x78, 0x63, 0x6c, 0x75, 0x64, 0x65, 0x64, 0x41, 0x74, 0x74, 0x72, 0x69,
	0x62, 0x75, 0x74, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x1d, 0x72, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x48, 0x61, 0x73, 0x68, 0x45, 0x78, 0x63, 0x6c, 0x75, 0x64, 0x65, 0x64,
	0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x22, 0xf6, 0x01, 0x0a, 0x10, 0x43,
	0x61, 0x70, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12,
	0x32, 0x0a, 0x0e, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73,
	0x2e, 0x4d, 0x61, 0x70, 0x52, 0x0d, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x12, 0x4f, 0x0a, 0x15, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x5f, 0x74, 0x72,
	0x69, 0x67, 0x67, 0x65, 0x72, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x19, 0x2e, 0x6c, 0x6f, 0x6f, 0x70, 0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x74, 0x65,
	0x54, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x48, 0x00, 0x52,
	0x13, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x54, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x12, 0x4c, 0x0a, 0x14, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x5f, 0x74,
	0x61, 0x72, 0x67, 0x65, 0x74, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x18, 0x2e, 0x6c, 0x6f, 0x6f, 0x70, 0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x74, 0x65,
	0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x48, 0x00, 0x52, 0x12,
	0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x42, 0x0f, 0x0a, 0x0d, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x5f, 0x63, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x42, 0x42, 0x5a, 0x40, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x73, 0x6d, 0x61, 0x72, 0x74, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x6b,
	0x69, 0x74, 0x2f, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x6c, 0x69, 0x6e, 0x6b, 0x2d, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x63, 0x61, 0x70, 0x61, 0x62, 0x69, 0x6c, 0x69,
	0x74, 0x69, 0x65, 0x73, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_capabilities_pb_registry_proto_rawDescOnce sync.Once
	file_capabilities_pb_registry_proto_rawDescData = file_capabilities_pb_registry_proto_rawDesc
)

func file_capabilities_pb_registry_proto_rawDescGZIP() []byte {
	file_capabilities_pb_registry_proto_rawDescOnce.Do(func() {
		file_capabilities_pb_registry_proto_rawDescData = protoimpl.X.CompressGZIP(file_capabilities_pb_registry_proto_rawDescData)
	})
	return file_capabilities_pb_registry_proto_rawDescData
}

var file_capabilities_pb_registry_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_capabilities_pb_registry_proto_goTypes = []interface{}{
	(*RemoteTriggerConfig)(nil), // 0: loop.RemoteTriggerConfig
	(*RemoteTargetConfig)(nil),  // 1: loop.RemoteTargetConfig
	(*CapabilityConfig)(nil),    // 2: loop.CapabilityConfig
	(*durationpb.Duration)(nil), // 3: google.protobuf.Duration
	(*pb.Map)(nil),              // 4: values.Map
}
var file_capabilities_pb_registry_proto_depIdxs = []int32{
	3, // 0: loop.RemoteTriggerConfig.registrationRefresh:type_name -> google.protobuf.Duration
	3, // 1: loop.RemoteTriggerConfig.registrationExpiry:type_name -> google.protobuf.Duration
	3, // 2: loop.RemoteTriggerConfig.messageExpiry:type_name -> google.protobuf.Duration
	3, // 3: loop.RemoteTriggerConfig.batchCollectionPeriod:type_name -> google.protobuf.Duration
	4, // 4: loop.CapabilityConfig.default_config:type_name -> values.Map
	0, // 5: loop.CapabilityConfig.remote_trigger_config:type_name -> loop.RemoteTriggerConfig
	1, // 6: loop.CapabilityConfig.remote_target_config:type_name -> loop.RemoteTargetConfig
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_capabilities_pb_registry_proto_init() }
func file_capabilities_pb_registry_proto_init() {
	if File_capabilities_pb_registry_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_capabilities_pb_registry_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemoteTriggerConfig); i {
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
		file_capabilities_pb_registry_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemoteTargetConfig); i {
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
		file_capabilities_pb_registry_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CapabilityConfig); i {
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
	file_capabilities_pb_registry_proto_msgTypes[2].OneofWrappers = []interface{}{
		(*CapabilityConfig_RemoteTriggerConfig)(nil),
		(*CapabilityConfig_RemoteTargetConfig)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_capabilities_pb_registry_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_capabilities_pb_registry_proto_goTypes,
		DependencyIndexes: file_capabilities_pb_registry_proto_depIdxs,
		MessageInfos:      file_capabilities_pb_registry_proto_msgTypes,
	}.Build()
	File_capabilities_pb_registry_proto = out.File
	file_capabilities_pb_registry_proto_rawDesc = nil
	file_capabilities_pb_registry_proto_goTypes = nil
	file_capabilities_pb_registry_proto_depIdxs = nil
}
