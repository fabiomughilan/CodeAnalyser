// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.6.1
// source: protos/output/global/proc.proto

package global

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

type ProcFileOutput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// check ProcCommands
	Used bool `protobuf:"varint,1,opt,name=used,proto3" json:"used,omitempty"`
	// list of all ProcCommands
	FilePath string              `protobuf:"bytes,2,opt,name=filePath,proto3" json:"filePath,omitempty"`
	Commands map[string]*Command `protobuf:"bytes,3,rep,name=commands,proto3" json:"commands,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *ProcFileOutput) Reset() {
	*x = ProcFileOutput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_output_global_proc_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProcFileOutput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProcFileOutput) ProtoMessage() {}

func (x *ProcFileOutput) ProtoReflect() protoreflect.Message {
	mi := &file_protos_output_global_proc_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProcFileOutput.ProtoReflect.Descriptor instead.
func (*ProcFileOutput) Descriptor() ([]byte, []int) {
	return file_protos_output_global_proc_proto_rawDescGZIP(), []int{0}
}

func (x *ProcFileOutput) GetUsed() bool {
	if x != nil {
		return x.Used
	}
	return false
}

func (x *ProcFileOutput) GetFilePath() string {
	if x != nil {
		return x.FilePath
	}
	return ""
}

func (x *ProcFileOutput) GetCommands() map[string]*Command {
	if x != nil {
		return x.Commands
	}
	return nil
}

var File_protos_output_global_proc_proto protoreflect.FileDescriptor

var file_protos_output_global_proc_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x2f,
	0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x2f, 0x70, 0x72, 0x6f, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x14, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74,
	0x2e, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x1a, 0x22, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f,
	0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x2f, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x2f, 0x63, 0x6f,
	0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xec, 0x01, 0x0a, 0x0e,
	0x50, 0x72, 0x6f, 0x63, 0x46, 0x69, 0x6c, 0x65, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x12, 0x12,
	0x0a, 0x04, 0x75, 0x73, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x04, 0x75, 0x73,
	0x65, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x50, 0x61, 0x74, 0x68, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x50, 0x61, 0x74, 0x68, 0x12, 0x4e,
	0x0a, 0x08, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x32, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74,
	0x2e, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x2e, 0x50, 0x72, 0x6f, 0x63, 0x46, 0x69, 0x6c, 0x65,
	0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x73, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x52, 0x08, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x73, 0x1a, 0x5a,
	0x0a, 0x0d, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12,
	0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65,
	0x79, 0x12, 0x33, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74,
	0x2e, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x52,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x42, 0x27, 0x5a, 0x25, 0x63, 0x6f,
	0x64, 0x65, 0x2d, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x73, 0x65, 0x72, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x73, 0x2f, 0x70, 0x62, 0x2f, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x2f, 0x67, 0x6c, 0x6f,
	0x62, 0x61, 0x6c, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protos_output_global_proc_proto_rawDescOnce sync.Once
	file_protos_output_global_proc_proto_rawDescData = file_protos_output_global_proc_proto_rawDesc
)

func file_protos_output_global_proc_proto_rawDescGZIP() []byte {
	file_protos_output_global_proc_proto_rawDescOnce.Do(func() {
		file_protos_output_global_proc_proto_rawDescData = protoimpl.X.CompressGZIP(file_protos_output_global_proc_proto_rawDescData)
	})
	return file_protos_output_global_proc_proto_rawDescData
}

var file_protos_output_global_proc_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_protos_output_global_proc_proto_goTypes = []interface{}{
	(*ProcFileOutput)(nil), // 0: protos.output.global.ProcFileOutput
	nil,                    // 1: protos.output.global.ProcFileOutput.CommandsEntry
	(*Command)(nil),        // 2: protos.output.global.Command
}
var file_protos_output_global_proc_proto_depIdxs = []int32{
	1, // 0: protos.output.global.ProcFileOutput.commands:type_name -> protos.output.global.ProcFileOutput.CommandsEntry
	2, // 1: protos.output.global.ProcFileOutput.CommandsEntry.value:type_name -> protos.output.global.Command
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_protos_output_global_proc_proto_init() }
func file_protos_output_global_proc_proto_init() {
	if File_protos_output_global_proc_proto != nil {
		return
	}
	file_protos_output_global_command_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_protos_output_global_proc_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProcFileOutput); i {
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
			RawDescriptor: file_protos_output_global_proc_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_protos_output_global_proc_proto_goTypes,
		DependencyIndexes: file_protos_output_global_proc_proto_depIdxs,
		MessageInfos:      file_protos_output_global_proc_proto_msgTypes,
	}.Build()
	File_protos_output_global_proc_proto = out.File
	file_protos_output_global_proc_proto_rawDesc = nil
	file_protos_output_global_proc_proto_goTypes = nil
	file_protos_output_global_proc_proto_depIdxs = nil
}
