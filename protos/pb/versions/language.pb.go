// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.15.2
// source: protos/versions/language.proto

package versions

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

type LanguageVersion struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Detectruntimecommand string                               `protobuf:"bytes,7,opt,name=detectruntimecommand,proto3" json:"detectruntimecommand,omitempty"`
	Runtimeversions      map[string]*DependencyVersionDetails `protobuf:"bytes,1,rep,name=runtimeversions,proto3" json:"runtimeversions,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Framework            map[string]*DependencyDetails        `protobuf:"bytes,2,rep,name=framework,proto3" json:"framework,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Databases            map[string]*DependencyDetails        `protobuf:"bytes,3,rep,name=databases,proto3" json:"databases,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Orms                 map[string]*DependencyDetails        `protobuf:"bytes,4,rep,name=orms,proto3" json:"orms,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Libraries            map[string]*DependencyDetails        `protobuf:"bytes,5,rep,name=libraries,proto3" json:"libraries,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Dependencies         map[string]*DependencyDetails        `protobuf:"bytes,6,rep,name=dependencies,proto3" json:"dependencies,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	DetectEnvCommand     string                               `protobuf:"bytes,8,opt,name=detectEnvCommand,proto3" json:"detectEnvCommand,omitempty"`
	PreDetectCommand     string                               `protobuf:"bytes,9,opt,name=preDetectCommand,proto3" json:"preDetectCommand,omitempty"`
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

	Semver        []string `protobuf:"bytes,2,rep,name=semver,proto3" json:"semver,omitempty"`
	Plugincommand string   `protobuf:"bytes,4,opt,name=plugincommand,proto3" json:"plugincommand,omitempty"`
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
	0x22, 0x9a, 0x08, 0x0a, 0x0f, 0x4c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x56, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x12, 0x32, 0x0a, 0x14, 0x64, 0x65, 0x74, 0x65, 0x63, 0x74, 0x72, 0x75,
	0x6e, 0x74, 0x69, 0x6d, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x14, 0x64, 0x65, 0x74, 0x65, 0x63, 0x74, 0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d,
	0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x12, 0x4f, 0x0a, 0x0f, 0x72, 0x75, 0x6e, 0x74,
	0x69, 0x6d, 0x65, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x25, 0x2e, 0x4c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x56, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x2e, 0x52, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x76, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0f, 0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d,
	0x65, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x3d, 0x0a, 0x09, 0x66, 0x72, 0x61,
	0x6d, 0x65, 0x77, 0x6f, 0x72, 0x6b, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x4c,
	0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x46,
	0x72, 0x61, 0x6d, 0x65, 0x77, 0x6f, 0x72, 0x6b, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x09, 0x66,
	0x72, 0x61, 0x6d, 0x65, 0x77, 0x6f, 0x72, 0x6b, 0x12, 0x3d, 0x0a, 0x09, 0x64, 0x61, 0x74, 0x61,
	0x62, 0x61, 0x73, 0x65, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x4c, 0x61,
	0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x44, 0x61,
	0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x09, 0x64, 0x61,
	0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x73, 0x12, 0x2e, 0x0a, 0x04, 0x6f, 0x72, 0x6d, 0x73, 0x18,
	0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x4c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65,
	0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x4f, 0x72, 0x6d, 0x73, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x52, 0x04, 0x6f, 0x72, 0x6d, 0x73, 0x12, 0x3d, 0x0a, 0x09, 0x6c, 0x69, 0x62, 0x72, 0x61,
	0x72, 0x69, 0x65, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x4c, 0x61, 0x6e,
	0x67, 0x75, 0x61, 0x67, 0x65, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x4c, 0x69, 0x62,
	0x72, 0x61, 0x72, 0x69, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x09, 0x6c, 0x69, 0x62,
	0x72, 0x61, 0x72, 0x69, 0x65, 0x73, 0x12, 0x46, 0x0a, 0x0c, 0x64, 0x65, 0x70, 0x65, 0x6e, 0x64,
	0x65, 0x6e, 0x63, 0x69, 0x65, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x4c,
	0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x44,
	0x65, 0x70, 0x65, 0x6e, 0x64, 0x65, 0x6e, 0x63, 0x69, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x52, 0x0c, 0x64, 0x65, 0x70, 0x65, 0x6e, 0x64, 0x65, 0x6e, 0x63, 0x69, 0x65, 0x73, 0x12, 0x2a,
	0x0a, 0x10, 0x64, 0x65, 0x74, 0x65, 0x63, 0x74, 0x45, 0x6e, 0x76, 0x43, 0x6f, 0x6d, 0x6d, 0x61,
	0x6e, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x64, 0x65, 0x74, 0x65, 0x63, 0x74,
	0x45, 0x6e, 0x76, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x12, 0x2a, 0x0a, 0x10, 0x70, 0x72,
	0x65, 0x44, 0x65, 0x74, 0x65, 0x63, 0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x18, 0x09,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x70, 0x72, 0x65, 0x44, 0x65, 0x74, 0x65, 0x63, 0x74, 0x43,
	0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x1a, 0x5d, 0x0a, 0x14, 0x52, 0x75, 0x6e, 0x74, 0x69, 0x6d,
	0x65, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10,
	0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79,
	0x12, 0x2f, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x19, 0x2e, 0x44, 0x65, 0x70, 0x65, 0x6e, 0x64, 0x65, 0x6e, 0x63, 0x79, 0x56, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x50, 0x0a, 0x0e, 0x46, 0x72, 0x61, 0x6d, 0x65, 0x77, 0x6f,
	0x72, 0x6b, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x28, 0x0a, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x44, 0x65, 0x70, 0x65, 0x6e,
	0x64, 0x65, 0x6e, 0x63, 0x79, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x52, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x50, 0x0a, 0x0e, 0x44, 0x61, 0x74, 0x61, 0x62,
	0x61, 0x73, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x28, 0x0a, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x44, 0x65, 0x70,
	0x65, 0x6e, 0x64, 0x65, 0x6e, 0x63, 0x79, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x52, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x4b, 0x0a, 0x09, 0x4f, 0x72, 0x6d,
	0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x28, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x44, 0x65, 0x70, 0x65, 0x6e, 0x64,
	0x65, 0x6e, 0x63, 0x79, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x52, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x50, 0x0a, 0x0e, 0x4c, 0x69, 0x62, 0x72, 0x61, 0x72,
	0x69, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x28, 0x0a, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x44, 0x65, 0x70, 0x65,
	0x6e, 0x64, 0x65, 0x6e, 0x63, 0x79, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x52, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x53, 0x0a, 0x11, 0x44, 0x65, 0x70, 0x65,
	0x6e, 0x64, 0x65, 0x6e, 0x63, 0x69, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a,
	0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12,
	0x28, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12,
	0x2e, 0x44, 0x65, 0x70, 0x65, 0x6e, 0x64, 0x65, 0x6e, 0x63, 0x79, 0x44, 0x65, 0x74, 0x61, 0x69,
	0x6c, 0x73, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0xa5, 0x01,
	0x0a, 0x11, 0x44, 0x65, 0x70, 0x65, 0x6e, 0x64, 0x65, 0x6e, 0x63, 0x79, 0x44, 0x65, 0x74, 0x61,
	0x69, 0x6c, 0x73, 0x12, 0x39, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x44, 0x65, 0x70, 0x65, 0x6e, 0x64, 0x65, 0x6e, 0x63,
	0x79, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x2e, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x1a, 0x55,
	0x0a, 0x0c, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10,
	0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79,
	0x12, 0x2f, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x19, 0x2e, 0x44, 0x65, 0x70, 0x65, 0x6e, 0x64, 0x65, 0x6e, 0x63, 0x79, 0x56, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x58, 0x0a, 0x18, 0x44, 0x65, 0x70, 0x65, 0x6e, 0x64, 0x65,
	0x6e, 0x63, 0x79, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c,
	0x73, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x65, 0x6d, 0x76, 0x65, 0x72, 0x18, 0x02, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x06, 0x73, 0x65, 0x6d, 0x76, 0x65, 0x72, 0x12, 0x24, 0x0a, 0x0d, 0x70, 0x6c, 0x75,
	0x67, 0x69, 0x6e, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
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
	(*LanguageVersion)(nil),          // 0: LanguageVersion
	(*DependencyDetails)(nil),        // 1: DependencyDetails
	(*DependencyVersionDetails)(nil), // 2: DependencyVersionDetails
	nil,                              // 3: LanguageVersion.RuntimeversionsEntry
	nil,                              // 4: LanguageVersion.FrameworkEntry
	nil,                              // 5: LanguageVersion.DatabasesEntry
	nil,                              // 6: LanguageVersion.OrmsEntry
	nil,                              // 7: LanguageVersion.LibrariesEntry
	nil,                              // 8: LanguageVersion.DependenciesEntry
	nil,                              // 9: DependencyDetails.VersionEntry
}
var file_protos_versions_language_proto_depIdxs = []int32{
	3,  // 0: LanguageVersion.runtimeversions:type_name -> LanguageVersion.RuntimeversionsEntry
	4,  // 1: LanguageVersion.framework:type_name -> LanguageVersion.FrameworkEntry
	5,  // 2: LanguageVersion.databases:type_name -> LanguageVersion.DatabasesEntry
	6,  // 3: LanguageVersion.orms:type_name -> LanguageVersion.OrmsEntry
	7,  // 4: LanguageVersion.libraries:type_name -> LanguageVersion.LibrariesEntry
	8,  // 5: LanguageVersion.dependencies:type_name -> LanguageVersion.DependenciesEntry
	9,  // 6: DependencyDetails.version:type_name -> DependencyDetails.VersionEntry
	2,  // 7: LanguageVersion.RuntimeversionsEntry.value:type_name -> DependencyVersionDetails
	1,  // 8: LanguageVersion.FrameworkEntry.value:type_name -> DependencyDetails
	1,  // 9: LanguageVersion.DatabasesEntry.value:type_name -> DependencyDetails
	1,  // 10: LanguageVersion.OrmsEntry.value:type_name -> DependencyDetails
	1,  // 11: LanguageVersion.LibrariesEntry.value:type_name -> DependencyDetails
	1,  // 12: LanguageVersion.DependenciesEntry.value:type_name -> DependencyDetails
	2,  // 13: DependencyDetails.VersionEntry.value:type_name -> DependencyVersionDetails
	14, // [14:14] is the sub-list for method output_type
	14, // [14:14] is the sub-list for method input_type
	14, // [14:14] is the sub-list for extension type_name
	14, // [14:14] is the sub-list for extension extendee
	0,  // [0:14] is the sub-list for field type_name
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
