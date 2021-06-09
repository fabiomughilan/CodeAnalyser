// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.6.1
// source: protos/plugin/orm.proto

package plugin

import (
	helpers "code-analyser/protos/pb/helpers"
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

type OrmOutput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Used      bool           `protobuf:"varint,1,opt,name=used,proto3" json:"used,omitempty"`
	DbName    string         `protobuf:"bytes,2,opt,name=DbName,proto3" json:"DbName,omitempty"`
	DbVersion string         `protobuf:"bytes,3,opt,name=DbVersion,proto3" json:"DbVersion,omitempty"`
	Error     *helpers.Error `protobuf:"bytes,4,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *OrmOutput) Reset() {
	*x = OrmOutput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_plugin_orm_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrmOutput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrmOutput) ProtoMessage() {}

func (x *OrmOutput) ProtoReflect() protoreflect.Message {
	mi := &file_protos_plugin_orm_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrmOutput.ProtoReflect.Descriptor instead.
func (*OrmOutput) Descriptor() ([]byte, []int) {
	return file_protos_plugin_orm_proto_rawDescGZIP(), []int{0}
}

func (x *OrmOutput) GetUsed() bool {
	if x != nil {
		return x.Used
	}
	return false
}

func (x *OrmOutput) GetDbName() string {
	if x != nil {
		return x.DbName
	}
	return ""
}

func (x *OrmOutput) GetDbVersion() string {
	if x != nil {
		return x.DbVersion
	}
	return ""
}

func (x *OrmOutput) GetError() *helpers.Error {
	if x != nil {
		return x.Error
	}
	return nil
}

var File_protos_plugin_orm_proto protoreflect.FileDescriptor

var file_protos_plugin_orm_proto_rawDesc = []byte{
	0x0a, 0x17, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2f,
	0x6f, 0x72, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x68, 0x65, 0x6c, 0x70, 0x65, 0x72, 0x73,
	0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x82, 0x01,
	0x0a, 0x09, 0x4f, 0x72, 0x6d, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x75,
	0x73, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x04, 0x75, 0x73, 0x65, 0x64, 0x12,
	0x16, 0x0a, 0x06, 0x44, 0x62, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x44, 0x62, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x44, 0x62, 0x56, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x44, 0x62, 0x56, 0x65,
	0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x2b, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x68, 0x65,
	0x6c, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x05, 0x65, 0x72, 0x72,
	0x6f, 0x72, 0x32, 0xc0, 0x01, 0x0a, 0x03, 0x4f, 0x72, 0x6d, 0x12, 0x31, 0x0a, 0x06, 0x44, 0x65,
	0x74, 0x65, 0x63, 0x74, 0x12, 0x15, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x68, 0x65,
	0x6c, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x1a, 0x10, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x4f, 0x72, 0x6d, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x12, 0x3e, 0x0a,
	0x09, 0x49, 0x73, 0x4f, 0x52, 0x4d, 0x55, 0x73, 0x65, 0x64, 0x12, 0x15, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x73, 0x2e, 0x68, 0x65, 0x6c, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x49, 0x6e, 0x70, 0x75,
	0x74, 0x1a, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x68, 0x65, 0x6c, 0x70, 0x65,
	0x72, 0x73, 0x2e, 0x42, 0x6f, 0x6f, 0x6c, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x12, 0x46, 0x0a,
	0x10, 0x50, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x4f, 0x66, 0x4f, 0x52, 0x4d, 0x55, 0x73, 0x65,
	0x64, 0x12, 0x15, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x68, 0x65, 0x6c, 0x70, 0x65,
	0x72, 0x73, 0x2e, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x1a, 0x1b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x73, 0x2e, 0x68, 0x65, 0x6c, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x46, 0x6c, 0x6f, 0x61, 0x74, 0x4f,
	0x75, 0x74, 0x70, 0x75, 0x74, 0x42, 0x20, 0x5a, 0x1e, 0x63, 0x6f, 0x64, 0x65, 0x2d, 0x61, 0x6e,
	0x61, 0x6c, 0x79, 0x73, 0x65, 0x72, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x70, 0x62,
	0x2f, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protos_plugin_orm_proto_rawDescOnce sync.Once
	file_protos_plugin_orm_proto_rawDescData = file_protos_plugin_orm_proto_rawDesc
)

func file_protos_plugin_orm_proto_rawDescGZIP() []byte {
	file_protos_plugin_orm_proto_rawDescOnce.Do(func() {
		file_protos_plugin_orm_proto_rawDescData = protoimpl.X.CompressGZIP(file_protos_plugin_orm_proto_rawDescData)
	})
	return file_protos_plugin_orm_proto_rawDescData
}

var file_protos_plugin_orm_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_protos_plugin_orm_proto_goTypes = []interface{}{
	(*OrmOutput)(nil),           // 0: proto.OrmOutput
	(*helpers.Error)(nil),       // 1: protos.helpers.Error
	(*helpers.Input)(nil),       // 2: protos.helpers.Input
	(*helpers.BoolOutput)(nil),  // 3: protos.helpers.BoolOutput
	(*helpers.FloatOutput)(nil), // 4: protos.helpers.FloatOutput
}
var file_protos_plugin_orm_proto_depIdxs = []int32{
	1, // 0: proto.OrmOutput.error:type_name -> protos.helpers.Error
	2, // 1: proto.Orm.Detect:input_type -> protos.helpers.Input
	2, // 2: proto.Orm.IsORMUsed:input_type -> protos.helpers.Input
	2, // 3: proto.Orm.PercentOfORMUsed:input_type -> protos.helpers.Input
	0, // 4: proto.Orm.Detect:output_type -> proto.OrmOutput
	3, // 5: proto.Orm.IsORMUsed:output_type -> protos.helpers.BoolOutput
	4, // 6: proto.Orm.PercentOfORMUsed:output_type -> protos.helpers.FloatOutput
	4, // [4:7] is the sub-list for method output_type
	1, // [1:4] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_protos_plugin_orm_proto_init() }
func file_protos_plugin_orm_proto_init() {
	if File_protos_plugin_orm_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protos_plugin_orm_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrmOutput); i {
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
			RawDescriptor: file_protos_plugin_orm_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_protos_plugin_orm_proto_goTypes,
		DependencyIndexes: file_protos_plugin_orm_proto_depIdxs,
		MessageInfos:      file_protos_plugin_orm_proto_msgTypes,
	}.Build()
	File_protos_plugin_orm_proto = out.File
	file_protos_plugin_orm_proto_rawDesc = nil
	file_protos_plugin_orm_proto_goTypes = nil
	file_protos_plugin_orm_proto_depIdxs = nil
}
