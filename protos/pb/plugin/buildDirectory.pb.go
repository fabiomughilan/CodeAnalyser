// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.6.1
// source: protos/plugin/buildDirectory.proto

package plugin

import (
	helpers "code-analyser/protos/pb/helpers"
	languageSpecific "code-analyser/protos/pb/output/languageSpecific"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_protos_plugin_buildDirectory_proto protoreflect.FileDescriptor

var file_protos_plugin_buildDirectory_proto_rawDesc = []byte{
	0x0a, 0x22, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2f,
	0x62, 0x75, 0x69, 0x6c, 0x64, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x79, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x73, 0x2f, 0x68, 0x65, 0x6c, 0x70, 0x65, 0x72, 0x73, 0x2f, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x33, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73,
	0x2f, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x2f, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65,
	0x53, 0x70, 0x65, 0x63, 0x69, 0x66, 0x69, 0x63, 0x2f, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x44, 0x69,
	0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0x67, 0x0a,
	0x0e, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x79, 0x12,
	0x55, 0x0a, 0x06, 0x44, 0x65, 0x74, 0x65, 0x63, 0x74, 0x12, 0x15, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x73, 0x2e, 0x68, 0x65, 0x6c, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x49, 0x6e, 0x70, 0x75, 0x74,
	0x1a, 0x34, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74,
	0x2e, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x53, 0x70, 0x65, 0x63, 0x69, 0x66, 0x69,
	0x63, 0x2e, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x79,
	0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x42, 0x20, 0x5a, 0x1e, 0x63, 0x6f, 0x64, 0x65, 0x2d, 0x61,
	0x6e, 0x61, 0x6c, 0x79, 0x73, 0x65, 0x72, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x70,
	0x62, 0x2f, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_protos_plugin_buildDirectory_proto_goTypes = []interface{}{
	(*helpers.Input)(nil),                         // 0: protos.helpers.Input
	(*languageSpecific.BuildDirectoryOutput)(nil), // 1: protos.output.languageSpecific.BuildDirectoryOutput
}
var file_protos_plugin_buildDirectory_proto_depIdxs = []int32{
	0, // 0: proto.BuildDirectory.Detect:input_type -> protos.helpers.Input
	1, // 1: proto.BuildDirectory.Detect:output_type -> protos.output.languageSpecific.BuildDirectoryOutput
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_protos_plugin_buildDirectory_proto_init() }
func file_protos_plugin_buildDirectory_proto_init() {
	if File_protos_plugin_buildDirectory_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_protos_plugin_buildDirectory_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_protos_plugin_buildDirectory_proto_goTypes,
		DependencyIndexes: file_protos_plugin_buildDirectory_proto_depIdxs,
	}.Build()
	File_protos_plugin_buildDirectory_proto = out.File
	file_protos_plugin_buildDirectory_proto_rawDesc = nil
	file_protos_plugin_buildDirectory_proto_goTypes = nil
	file_protos_plugin_buildDirectory_proto_depIdxs = nil
}