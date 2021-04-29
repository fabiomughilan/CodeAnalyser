// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.6.1
// source: protos/versions/language.proto

package versions

import (
	utils "code-analyser/protos/pb/output/utils"
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

type LanguageVersion struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Detectruntimecommand   string                               `protobuf:"bytes,7,opt,name=detectruntimecommand,proto3" json:"detectruntimecommand,omitempty"`
	Runtimeversions        map[string]*DependencyVersionDetails `protobuf:"bytes,1,rep,name=runtimeversions,proto3" json:"runtimeversions,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Framework              map[string]*DependencyDetails        `protobuf:"bytes,2,rep,name=framework,proto3" json:"framework,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Databases              map[string]*DependencyDetails        `protobuf:"bytes,3,rep,name=databases,proto3" json:"databases,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Orms                   map[string]*DependencyDetails        `protobuf:"bytes,4,rep,name=orms,proto3" json:"orms,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Libraries              map[string]*DependencyDetails        `protobuf:"bytes,5,rep,name=libraries,proto3" json:"libraries,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Dependencies           map[string]*DependencyDetails        `protobuf:"bytes,6,rep,name=dependencies,proto3" json:"dependencies,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	DetectEnvCommand       string                               `protobuf:"bytes,8,opt,name=detectEnvCommand,proto3" json:"detectEnvCommand,omitempty"`
	PreDetectCommand       string                               `protobuf:"bytes,9,opt,name=preDetectCommand,proto3" json:"preDetectCommand,omitempty"`
	StaticAssetsCommand    string                               `protobuf:"bytes,10,opt,name=staticAssetsCommand,proto3" json:"staticAssetsCommand,omitempty"`
	BuildDirectoryCommand  string                               `protobuf:"bytes,11,opt,name=buildDirectoryCommand,proto3" json:"buildDirectoryCommand,omitempty"`
	DetectTestCasesCommand string                               `protobuf:"bytes,12,opt,name=detectTestCasesCommand,proto3" json:"detectTestCasesCommand,omitempty"`
	Commands               string                               `protobuf:"bytes,13,opt,name=commands,proto3" json:"commands,omitempty"`
}

func (x *LanguageVersion) Reset() {
	*x = LanguageVersion{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_versions_language_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LanguageVersion) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LanguageVersion) ProtoMessage() {}

func (x *LanguageVersion) ProtoReflect() protoreflect.Message {
	mi := &file_protos_versions_language_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LanguageVersion.ProtoReflect.Descriptor instead.
func (*LanguageVersion) Descriptor() ([]byte, []int) {
	return file_protos_versions_language_proto_rawDescGZIP(), []int{0}
}

func (x *LanguageVersion) GetDetectruntimecommand() string {
	if x != nil {
		return x.Detectruntimecommand
	}
	return ""
}

func (x *LanguageVersion) GetRuntimeversions() map[string]*DependencyVersionDetails {
	if x != nil {
		return x.Runtimeversions
	}
	return nil
}

func (x *LanguageVersion) GetFramework() map[string]*DependencyDetails {
	if x != nil {
		return x.Framework
	}
	return nil
}

func (x *LanguageVersion) GetDatabases() map[string]*DependencyDetails {
	if x != nil {
		return x.Databases
	}
	return nil
}

func (x *LanguageVersion) GetOrms() map[string]*DependencyDetails {
	if x != nil {
		return x.Orms
	}
	return nil
}

func (x *LanguageVersion) GetLibraries() map[string]*DependencyDetails {
	if x != nil {
		return x.Libraries
	}
	return nil
}

func (x *LanguageVersion) GetDependencies() map[string]*DependencyDetails {
	if x != nil {
		return x.Dependencies
	}
	return nil
}

func (x *LanguageVersion) GetDetectEnvCommand() string {
	if x != nil {
		return x.DetectEnvCommand
	}
	return ""
}

func (x *LanguageVersion) GetPreDetectCommand() string {
	if x != nil {
		return x.PreDetectCommand
	}
	return ""
}

func (x *LanguageVersion) GetStaticAssetsCommand() string {
	if x != nil {
		return x.StaticAssetsCommand
	}
	return ""
}

func (x *LanguageVersion) GetBuildDirectoryCommand() string {
	if x != nil {
		return x.BuildDirectoryCommand
	}
	return ""
}

func (x *LanguageVersion) GetDetectTestCasesCommand() string {
	if x != nil {
		return x.DetectTestCasesCommand
	}
	return ""
}

func (x *LanguageVersion) GetCommands() string {
	if x != nil {
		return x.Commands
	}
	return ""
}

type DependencyDetails struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Version map[string]*DependencyVersionDetails `protobuf:"bytes,1,rep,name=version,proto3" json:"version,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *DependencyDetails) Reset() {
	*x = DependencyDetails{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_versions_language_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DependencyDetails) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DependencyDetails) ProtoMessage() {}

func (x *DependencyDetails) ProtoReflect() protoreflect.Message {
	mi := &file_protos_versions_language_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DependencyDetails.ProtoReflect.Descriptor instead.
func (*DependencyDetails) Descriptor() ([]byte, []int) {
	return file_protos_versions_language_proto_rawDescGZIP(), []int{1}
}

func (x *DependencyDetails) GetVersion() map[string]*DependencyVersionDetails {
	if x != nil {
		return x.Version
	}
	return nil
}

type DependencyVersionDetails struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Semver        []string         `protobuf:"bytes,2,rep,name=semver,proto3" json:"semver,omitempty"`
	Libraries     []*utils.Library `protobuf:"bytes,1,rep,name=libraries,proto3" json:"libraries,omitempty"`
	Plugincommand string           `protobuf:"bytes,3,opt,name=plugincommand,proto3" json:"plugincommand,omitempty"`
}

func (x *DependencyVersionDetails) Reset() {
	*x = DependencyVersionDetails{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_versions_language_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DependencyVersionDetails) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DependencyVersionDetails) ProtoMessage() {}

func (x *DependencyVersionDetails) ProtoReflect() protoreflect.Message {
	mi := &file_protos_versions_language_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DependencyVersionDetails.ProtoReflect.Descriptor instead.
func (*DependencyVersionDetails) Descriptor() ([]byte, []int) {
	return file_protos_versions_language_proto_rawDescGZIP(), []int{2}
}

func (x *DependencyVersionDetails) GetSemver() []string {
	if x != nil {
		return x.Semver
	}
	return nil
}

func (x *DependencyVersionDetails) GetLibraries() []*utils.Library {
	if x != nil {
		return x.Libraries
	}
	return nil
}

func (x *DependencyVersionDetails) GetPlugincommand() string {
	if x != nil {
		return x.Plugincommand
	}
	return ""
}

var File_protos_versions_language_proto protoreflect.FileDescriptor

var file_protos_versions_language_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x73, 0x2f, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x0f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x73, 0x1a, 0x24, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74,
	0x2f, 0x75, 0x74, 0x69, 0x6c, 0x73, 0x2f, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x59, 0x61, 0x6d,
	0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x96, 0x0b, 0x0a, 0x0f, 0x4c, 0x61, 0x6e, 0x67,
	0x75, 0x61, 0x67, 0x65, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x32, 0x0a, 0x14, 0x64,
	0x65, 0x74, 0x65, 0x63, 0x74, 0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x63, 0x6f, 0x6d, 0x6d,
	0x61, 0x6e, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x14, 0x64, 0x65, 0x74, 0x65, 0x63,
	0x74, 0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x12,
	0x5f, 0x0a, 0x0f, 0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x35, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x73, 0x2e, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x4c, 0x61, 0x6e, 0x67, 0x75,
	0x61, 0x67, 0x65, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x52, 0x75, 0x6e, 0x74, 0x69,
	0x6d, 0x65, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52,
	0x0f, 0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x73,
	0x12, 0x4d, 0x0a, 0x09, 0x66, 0x72, 0x61, 0x6d, 0x65, 0x77, 0x6f, 0x72, 0x6b, 0x18, 0x02, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x2f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x76, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x4c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x56, 0x65,
	0x72, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x46, 0x72, 0x61, 0x6d, 0x65, 0x77, 0x6f, 0x72, 0x6b, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x52, 0x09, 0x66, 0x72, 0x61, 0x6d, 0x65, 0x77, 0x6f, 0x72, 0x6b, 0x12,
	0x4d, 0x0a, 0x09, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x73, 0x18, 0x03, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x2f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x76, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x4c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x56, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x73, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x52, 0x09, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x73, 0x12, 0x3e,
	0x0a, 0x04, 0x6f, 0x72, 0x6d, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x4c,
	0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x4f,
	0x72, 0x6d, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x04, 0x6f, 0x72, 0x6d, 0x73, 0x12, 0x4d,
	0x0a, 0x09, 0x6c, 0x69, 0x62, 0x72, 0x61, 0x72, 0x69, 0x65, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x2f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x76, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x73, 0x2e, 0x4c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x56, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x2e, 0x4c, 0x69, 0x62, 0x72, 0x61, 0x72, 0x69, 0x65, 0x73, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x52, 0x09, 0x6c, 0x69, 0x62, 0x72, 0x61, 0x72, 0x69, 0x65, 0x73, 0x12, 0x56, 0x0a,
	0x0c, 0x64, 0x65, 0x70, 0x65, 0x6e, 0x64, 0x65, 0x6e, 0x63, 0x69, 0x65, 0x73, 0x18, 0x06, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x32, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x76, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x4c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x56, 0x65,
	0x72, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x44, 0x65, 0x70, 0x65, 0x6e, 0x64, 0x65, 0x6e, 0x63, 0x69,
	0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0c, 0x64, 0x65, 0x70, 0x65, 0x6e, 0x64, 0x65,
	0x6e, 0x63, 0x69, 0x65, 0x73, 0x12, 0x2a, 0x0a, 0x10, 0x64, 0x65, 0x74, 0x65, 0x63, 0x74, 0x45,
	0x6e, 0x76, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x10, 0x64, 0x65, 0x74, 0x65, 0x63, 0x74, 0x45, 0x6e, 0x76, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e,
	0x64, 0x12, 0x2a, 0x0a, 0x10, 0x70, 0x72, 0x65, 0x44, 0x65, 0x74, 0x65, 0x63, 0x74, 0x43, 0x6f,
	0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x70, 0x72, 0x65,
	0x44, 0x65, 0x74, 0x65, 0x63, 0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x12, 0x30, 0x0a,
	0x13, 0x73, 0x74, 0x61, 0x74, 0x69, 0x63, 0x41, 0x73, 0x73, 0x65, 0x74, 0x73, 0x43, 0x6f, 0x6d,
	0x6d, 0x61, 0x6e, 0x64, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x13, 0x73, 0x74, 0x61, 0x74,
	0x69, 0x63, 0x41, 0x73, 0x73, 0x65, 0x74, 0x73, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x12,
	0x34, 0x0a, 0x15, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72,
	0x79, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x15,
	0x62, 0x75, 0x69, 0x6c, 0x64, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x79, 0x43, 0x6f,
	0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x12, 0x36, 0x0a, 0x16, 0x64, 0x65, 0x74, 0x65, 0x63, 0x74, 0x54,
	0x65, 0x73, 0x74, 0x43, 0x61, 0x73, 0x65, 0x73, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x18,
	0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x16, 0x64, 0x65, 0x74, 0x65, 0x63, 0x74, 0x54, 0x65, 0x73,
	0x74, 0x43, 0x61, 0x73, 0x65, 0x73, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x12, 0x1a, 0x0a,
	0x08, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x73, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x73, 0x1a, 0x6d, 0x0a, 0x14, 0x52, 0x75, 0x6e,
	0x74, 0x69, 0x6d, 0x65, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x6b, 0x65, 0x79, 0x12, 0x3f, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x29, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x76, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x44, 0x65, 0x70, 0x65, 0x6e, 0x64, 0x65, 0x6e, 0x63, 0x79, 0x56,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x52, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x60, 0x0a, 0x0e, 0x46, 0x72, 0x61, 0x6d,
	0x65, 0x77, 0x6f, 0x72, 0x6b, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x38, 0x0a, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x44, 0x65,
	0x70, 0x65, 0x6e, 0x64, 0x65, 0x6e, 0x63, 0x79, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x52,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x60, 0x0a, 0x0e, 0x44, 0x61,
	0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03,
	0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x38,
	0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e,
	0x44, 0x65, 0x70, 0x65, 0x6e, 0x64, 0x65, 0x6e, 0x63, 0x79, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c,
	0x73, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x5b, 0x0a, 0x09,
	0x4f, 0x72, 0x6d, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x38, 0x0a, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x73, 0x2e, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x44, 0x65, 0x70,
	0x65, 0x6e, 0x64, 0x65, 0x6e, 0x63, 0x79, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x52, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x60, 0x0a, 0x0e, 0x4c, 0x69, 0x62,
	0x72, 0x61, 0x72, 0x69, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b,
	0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x38, 0x0a,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x44,
	0x65, 0x70, 0x65, 0x6e, 0x64, 0x65, 0x6e, 0x63, 0x79, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73,
	0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x63, 0x0a, 0x11, 0x44,
	0x65, 0x70, 0x65, 0x6e, 0x64, 0x65, 0x6e, 0x63, 0x69, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b,
	0x65, 0x79, 0x12, 0x38, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x22, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x76, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x73, 0x2e, 0x44, 0x65, 0x70, 0x65, 0x6e, 0x64, 0x65, 0x6e, 0x63, 0x79, 0x44, 0x65,
	0x74, 0x61, 0x69, 0x6c, 0x73, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01,
	0x22, 0xc5, 0x01, 0x0a, 0x11, 0x44, 0x65, 0x70, 0x65, 0x6e, 0x64, 0x65, 0x6e, 0x63, 0x79, 0x44,
	0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x12, 0x49, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73,
	0x2e, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x44, 0x65, 0x70, 0x65, 0x6e, 0x64,
	0x65, 0x6e, 0x63, 0x79, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x2e, 0x56, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x1a, 0x65, 0x0a, 0x0c, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x6b, 0x65, 0x79, 0x12, 0x3f, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x29, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x76, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x44, 0x65, 0x70, 0x65, 0x6e, 0x64, 0x65, 0x6e, 0x63, 0x79, 0x56,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x52, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x94, 0x01, 0x0a, 0x18, 0x44, 0x65, 0x70,
	0x65, 0x6e, 0x64, 0x65, 0x6e, 0x63, 0x79, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x44, 0x65,
	0x74, 0x61, 0x69, 0x6c, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x65, 0x6d, 0x76, 0x65, 0x72, 0x18,
	0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x73, 0x65, 0x6d, 0x76, 0x65, 0x72, 0x12, 0x3a, 0x0a,
	0x09, 0x6c, 0x69, 0x62, 0x72, 0x61, 0x72, 0x69, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x1c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74,
	0x2e, 0x75, 0x74, 0x69, 0x6c, 0x73, 0x2e, 0x4c, 0x69, 0x62, 0x72, 0x61, 0x72, 0x79, 0x52, 0x09,
	0x6c, 0x69, 0x62, 0x72, 0x61, 0x72, 0x69, 0x65, 0x73, 0x12, 0x24, 0x0a, 0x0d, 0x70, 0x6c, 0x75,
	0x67, 0x69, 0x6e, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0d, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x42,
	0x22, 0x5a, 0x20, 0x63, 0x6f, 0x64, 0x65, 0x2d, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x73, 0x65, 0x72,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x70, 0x62, 0x2f, 0x76, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protos_versions_language_proto_rawDescOnce sync.Once
	file_protos_versions_language_proto_rawDescData = file_protos_versions_language_proto_rawDesc
)

func file_protos_versions_language_proto_rawDescGZIP() []byte {
	file_protos_versions_language_proto_rawDescOnce.Do(func() {
		file_protos_versions_language_proto_rawDescData = protoimpl.X.CompressGZIP(file_protos_versions_language_proto_rawDescData)
	})
	return file_protos_versions_language_proto_rawDescData
}

var file_protos_versions_language_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_protos_versions_language_proto_goTypes = []interface{}{
	(*LanguageVersion)(nil),          // 0: protos.versions.LanguageVersion
	(*DependencyDetails)(nil),        // 1: protos.versions.DependencyDetails
	(*DependencyVersionDetails)(nil), // 2: protos.versions.DependencyVersionDetails
	nil,                              // 3: protos.versions.LanguageVersion.RuntimeversionsEntry
	nil,                              // 4: protos.versions.LanguageVersion.FrameworkEntry
	nil,                              // 5: protos.versions.LanguageVersion.DatabasesEntry
	nil,                              // 6: protos.versions.LanguageVersion.OrmsEntry
	nil,                              // 7: protos.versions.LanguageVersion.LibrariesEntry
	nil,                              // 8: protos.versions.LanguageVersion.DependenciesEntry
	nil,                              // 9: protos.versions.DependencyDetails.VersionEntry
	(*utils.Library)(nil),            // 10: protos.output.utils.Library
}
var file_protos_versions_language_proto_depIdxs = []int32{
	3,  // 0: protos.versions.LanguageVersion.runtimeversions:type_name -> protos.versions.LanguageVersion.RuntimeversionsEntry
	4,  // 1: protos.versions.LanguageVersion.framework:type_name -> protos.versions.LanguageVersion.FrameworkEntry
	5,  // 2: protos.versions.LanguageVersion.databases:type_name -> protos.versions.LanguageVersion.DatabasesEntry
	6,  // 3: protos.versions.LanguageVersion.orms:type_name -> protos.versions.LanguageVersion.OrmsEntry
	7,  // 4: protos.versions.LanguageVersion.libraries:type_name -> protos.versions.LanguageVersion.LibrariesEntry
	8,  // 5: protos.versions.LanguageVersion.dependencies:type_name -> protos.versions.LanguageVersion.DependenciesEntry
	9,  // 6: protos.versions.DependencyDetails.version:type_name -> protos.versions.DependencyDetails.VersionEntry
	10, // 7: protos.versions.DependencyVersionDetails.libraries:type_name -> protos.output.utils.Library
	2,  // 8: protos.versions.LanguageVersion.RuntimeversionsEntry.value:type_name -> protos.versions.DependencyVersionDetails
	1,  // 9: protos.versions.LanguageVersion.FrameworkEntry.value:type_name -> protos.versions.DependencyDetails
	1,  // 10: protos.versions.LanguageVersion.DatabasesEntry.value:type_name -> protos.versions.DependencyDetails
	1,  // 11: protos.versions.LanguageVersion.OrmsEntry.value:type_name -> protos.versions.DependencyDetails
	1,  // 12: protos.versions.LanguageVersion.LibrariesEntry.value:type_name -> protos.versions.DependencyDetails
	1,  // 13: protos.versions.LanguageVersion.DependenciesEntry.value:type_name -> protos.versions.DependencyDetails
	2,  // 14: protos.versions.DependencyDetails.VersionEntry.value:type_name -> protos.versions.DependencyVersionDetails
	15, // [15:15] is the sub-list for method output_type
	15, // [15:15] is the sub-list for method input_type
	15, // [15:15] is the sub-list for extension type_name
	15, // [15:15] is the sub-list for extension extendee
	0,  // [0:15] is the sub-list for field type_name
}

func init() { file_protos_versions_language_proto_init() }
func file_protos_versions_language_proto_init() {
	if File_protos_versions_language_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protos_versions_language_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LanguageVersion); i {
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
		file_protos_versions_language_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DependencyDetails); i {
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
		file_protos_versions_language_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DependencyVersionDetails); i {
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
			RawDescriptor: file_protos_versions_language_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_protos_versions_language_proto_goTypes,
		DependencyIndexes: file_protos_versions_language_proto_depIdxs,
		MessageInfos:      file_protos_versions_language_proto_msgTypes,
	}.Build()
	File_protos_versions_language_proto = out.File
	file_protos_versions_language_proto_rawDesc = nil
	file_protos_versions_language_proto_goTypes = nil
	file_protos_versions_language_proto_depIdxs = nil
}
