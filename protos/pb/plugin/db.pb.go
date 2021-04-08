// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.15.2
// source: protos/plugin/db.proto

package plugin

import (
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

type ServiceOutputBoolInt struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value    bool          `protobuf:"varint,1,opt,name=value,proto3" json:"value,omitempty"`
	IntValue int64         `protobuf:"varint,2,opt,name=intValue,proto3" json:"intValue,omitempty"`
	Error    *ServiceError `protobuf:"bytes,3,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *ServiceOutputBoolInt) Reset() {
	*x = ServiceOutputBoolInt{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_plugin_db_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServiceOutputBoolInt) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServiceOutputBoolInt) ProtoMessage() {}

func (x *ServiceOutputBoolInt) ProtoReflect() protoreflect.Message {
	mi := &file_protos_plugin_db_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServiceOutputBoolInt.ProtoReflect.Descriptor instead.
func (*ServiceOutputBoolInt) Descriptor() ([]byte, []int) {
	return file_protos_plugin_db_proto_rawDescGZIP(), []int{0}
}

func (x *ServiceOutputBoolInt) GetValue() bool {
	if x != nil {
		return x.Value
	}
	return false
}

func (x *ServiceOutputBoolInt) GetIntValue() int64 {
	if x != nil {
		return x.IntValue
	}
	return 0
}

func (x *ServiceOutputBoolInt) GetError() *ServiceError {
	if x != nil {
		return x.Error
	}
	return nil
}

var File_protos_plugin_db_proto protoreflect.FileDescriptor

var file_protos_plugin_db_proto_rawDesc = []byte{
	0x0a, 0x16, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2f,
	0x64, 0x62, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1a, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2f, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x73, 0x0a, 0x14, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x42, 0x6f, 0x6f, 0x6c,
	0x49, 0x6e, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x69, 0x6e, 0x74,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x69, 0x6e, 0x74,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x29, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72,
	0x32, 0xc5, 0x01, 0x0a, 0x09, 0x44, 0x62, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3a,
	0x0a, 0x06, 0x44, 0x65, 0x74, 0x65, 0x63, 0x74, 0x12, 0x13, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x1a, 0x1b, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4f, 0x75, 0x74,
	0x70, 0x75, 0x74, 0x42, 0x6f, 0x6f, 0x6c, 0x49, 0x6e, 0x74, 0x12, 0x39, 0x0a, 0x08, 0x49, 0x73,
	0x44, 0x62, 0x55, 0x73, 0x65, 0x64, 0x12, 0x13, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x1a, 0x18, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4f, 0x75, 0x74, 0x70, 0x75,
	0x74, 0x42, 0x6f, 0x6f, 0x6c, 0x12, 0x41, 0x0a, 0x0f, 0x50, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74,
	0x4f, 0x66, 0x44, 0x62, 0x55, 0x73, 0x65, 0x64, 0x12, 0x13, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x1a, 0x19, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4f, 0x75, 0x74,
	0x70, 0x75, 0x74, 0x46, 0x6c, 0x6f, 0x61, 0x74, 0x42, 0x20, 0x5a, 0x1e, 0x63, 0x6f, 0x64, 0x65,
	0x2d, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x73, 0x65, 0x72, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73,
	0x2f, 0x70, 0x62, 0x2f, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_protos_plugin_db_proto_rawDescOnce sync.Once
	file_protos_plugin_db_proto_rawDescData = file_protos_plugin_db_proto_rawDesc
)

func file_protos_plugin_db_proto_rawDescGZIP() []byte {
	file_protos_plugin_db_proto_rawDescOnce.Do(func() {
		file_protos_plugin_db_proto_rawDescData = protoimpl.X.CompressGZIP(file_protos_plugin_db_proto_rawDescData)
	})
	return file_protos_plugin_db_proto_rawDescData
}

var file_protos_plugin_db_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_protos_plugin_db_proto_goTypes = []interface{}{
	(*ServiceOutputBoolInt)(nil), // 0: proto.ServiceOutputBoolInt
	(*ServiceError)(nil),         // 1: proto.ServiceError
	(*ServiceInput)(nil),         // 2: proto.ServiceInput
	(*ServiceOutputBool)(nil),    // 3: proto.ServiceOutputBool
	(*ServiceOutputFloat)(nil),   // 4: proto.ServiceOutputFloat
}
var file_protos_plugin_db_proto_depIdxs = []int32{
	1, // 0: proto.ServiceOutputBoolInt.error:type_name -> proto.ServiceError
	2, // 1: proto.DbService.Detect:input_type -> proto.ServiceInput
	2, // 2: proto.DbService.IsDbUsed:input_type -> proto.ServiceInput
	2, // 3: proto.DbService.PercentOfDbUsed:input_type -> proto.ServiceInput
	0, // 4: proto.DbService.Detect:output_type -> proto.ServiceOutputBoolInt
	3, // 5: proto.DbService.IsDbUsed:output_type -> proto.ServiceOutputBool
	4, // 6: proto.DbService.PercentOfDbUsed:output_type -> proto.ServiceOutputFloat
	4, // [4:7] is the sub-list for method output_type
	1, // [1:4] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_protos_plugin_db_proto_init() }
func file_protos_plugin_db_proto_init() {
	if File_protos_plugin_db_proto != nil {
		return
	}
	file_protos_plugin_common_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_protos_plugin_db_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServiceOutputBoolInt); i {
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
			RawDescriptor: file_protos_plugin_db_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_protos_plugin_db_proto_goTypes,
		DependencyIndexes: file_protos_plugin_db_proto_depIdxs,
		MessageInfos:      file_protos_plugin_db_proto_msgTypes,
	}.Build()
	File_protos_plugin_db_proto = out.File
	file_protos_plugin_db_proto_rawDesc = nil
	file_protos_plugin_db_proto_goTypes = nil
	file_protos_plugin_db_proto_depIdxs = nil
}
