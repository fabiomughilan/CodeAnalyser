// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        (unknown)
// source: protos/outputs/decisionMaker.proto

package protos

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

type LanguageSpecificDetections struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name           string                  `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	RuntimeVersion string                  `protobuf:"bytes,2,opt,name=runtimeVersion,proto3" json:"runtimeVersion,omitempty"`
	Env            []*EnvOutput            `protobuf:"bytes,3,rep,name=env,proto3" json:"env,omitempty"`
	Framework      []*FrameworkOutput      `protobuf:"bytes,4,rep,name=framework,proto3" json:"framework,omitempty"`
	Db             *DBOutput               `protobuf:"bytes,5,opt,name=db,proto3" json:"db,omitempty"`
	Orm            *OrmOutput              `protobuf:"bytes,6,opt,name=orm,proto3" json:"orm,omitempty"`
	Dependencies   []*DependenciesOutput   `protobuf:"bytes,7,rep,name=dependencies,proto3" json:"dependencies,omitempty"`
	Libraries      []*LibrariesOutput      `protobuf:"bytes,8,rep,name=libraries,proto3" json:"libraries,omitempty"`
	StaticAssets   []*StaticAssetsOutput   `protobuf:"bytes,9,rep,name=staticAssets,proto3" json:"staticAssets,omitempty"`
	StackOutput    []*StackOutput          `protobuf:"bytes,10,rep,name=stackOutput,proto3" json:"stackOutput,omitempty"`
	Appserver      []*AppserverOutput      `protobuf:"bytes,11,rep,name=appserver,proto3" json:"appserver,omitempty"`
	BuildDirectory []*BuildDirectoryOutput `protobuf:"bytes,12,rep,name=buildDirectory,proto3" json:"buildDirectory,omitempty"`
	Commands       *Commands               `protobuf:"bytes,13,opt,name=commands,proto3" json:"commands,omitempty"` // todo: DetectTestCasesRunCommands
}

func (x *LanguageSpecificDetections) Reset() {
	*x = LanguageSpecificDetections{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_outputs_decisionMaker_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LanguageSpecificDetections) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LanguageSpecificDetections) ProtoMessage() {}

func (x *LanguageSpecificDetections) ProtoReflect() protoreflect.Message {
	mi := &file_protos_outputs_decisionMaker_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LanguageSpecificDetections.ProtoReflect.Descriptor instead.
func (*LanguageSpecificDetections) Descriptor() ([]byte, []int) {
	return file_protos_outputs_decisionMaker_proto_rawDescGZIP(), []int{0}
}

func (x *LanguageSpecificDetections) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *LanguageSpecificDetections) GetRuntimeVersion() string {
	if x != nil {
		return x.RuntimeVersion
	}
	return ""
}

func (x *LanguageSpecificDetections) GetEnv() []*EnvOutput {
	if x != nil {
		return x.Env
	}
	return nil
}

func (x *LanguageSpecificDetections) GetFramework() []*FrameworkOutput {
	if x != nil {
		return x.Framework
	}
	return nil
}

func (x *LanguageSpecificDetections) GetDb() *DBOutput {
	if x != nil {
		return x.Db
	}
	return nil
}

func (x *LanguageSpecificDetections) GetOrm() *OrmOutput {
	if x != nil {
		return x.Orm
	}
	return nil
}

func (x *LanguageSpecificDetections) GetDependencies() []*DependenciesOutput {
	if x != nil {
		return x.Dependencies
	}
	return nil
}

func (x *LanguageSpecificDetections) GetLibraries() []*LibrariesOutput {
	if x != nil {
		return x.Libraries
	}
	return nil
}

func (x *LanguageSpecificDetections) GetStaticAssets() []*StaticAssetsOutput {
	if x != nil {
		return x.StaticAssets
	}
	return nil
}

func (x *LanguageSpecificDetections) GetStackOutput() []*StackOutput {
	if x != nil {
		return x.StackOutput
	}
	return nil
}

func (x *LanguageSpecificDetections) GetAppserver() []*AppserverOutput {
	if x != nil {
		return x.Appserver
	}
	return nil
}

func (x *LanguageSpecificDetections) GetBuildDirectory() []*BuildDirectoryOutput {
	if x != nil {
		return x.BuildDirectory
	}
	return nil
}

func (x *LanguageSpecificDetections) GetCommands() *Commands {
	if x != nil {
		return x.Commands
	}
	return nil
}

type GlobalDetections struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProcFile          []*ProcFileOutput          `protobuf:"bytes,1,rep,name=procFile,proto3" json:"procFile,omitempty"`
	Makefile          []*MakefileOutput          `protobuf:"bytes,2,rep,name=makefile,proto3" json:"makefile,omitempty"`
	DockerFile        []*DockerFileOutput        `protobuf:"bytes,3,rep,name=dockerFile,proto3" json:"dockerFile,omitempty"`
	DockerComposeFile []*DockerComposeFileOutput `protobuf:"bytes,4,rep,name=dockerComposeFile,proto3" json:"dockerComposeFile,omitempty"`
}

func (x *GlobalDetections) Reset() {
	*x = GlobalDetections{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_outputs_decisionMaker_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GlobalDetections) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GlobalDetections) ProtoMessage() {}

func (x *GlobalDetections) ProtoReflect() protoreflect.Message {
	mi := &file_protos_outputs_decisionMaker_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GlobalDetections.ProtoReflect.Descriptor instead.
func (*GlobalDetections) Descriptor() ([]byte, []int) {
	return file_protos_outputs_decisionMaker_proto_rawDescGZIP(), []int{1}
}

func (x *GlobalDetections) GetProcFile() []*ProcFileOutput {
	if x != nil {
		return x.ProcFile
	}
	return nil
}

func (x *GlobalDetections) GetMakefile() []*MakefileOutput {
	if x != nil {
		return x.Makefile
	}
	return nil
}

func (x *GlobalDetections) GetDockerFile() []*DockerFileOutput {
	if x != nil {
		return x.DockerFile
	}
	return nil
}

func (x *GlobalDetections) GetDockerComposeFile() []*DockerComposeFileOutput {
	if x != nil {
		return x.DockerComposeFile
	}
	return nil
}

type Commands struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BuildCommands     []*BuildCommandsOutput     `protobuf:"bytes,1,rep,name=buildCommands,proto3" json:"buildCommands,omitempty"`
	StartUpCommands   []*StartUpCommandsOutput   `protobuf:"bytes,2,rep,name=startUpCommands,proto3" json:"startUpCommands,omitempty"`
	SeedCommands      []*SeedCommandsOutput      `protobuf:"bytes,3,rep,name=seedCommands,proto3" json:"seedCommands,omitempty"`
	MigrationCommands []*MigrationCommandsOutput `protobuf:"bytes,4,rep,name=MigrationCommands,json=migrationCommands,proto3" json:"MigrationCommands,omitempty"` //   todo:   DetectAdHocScripts
}

func (x *Commands) Reset() {
	*x = Commands{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_outputs_decisionMaker_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Commands) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Commands) ProtoMessage() {}

func (x *Commands) ProtoReflect() protoreflect.Message {
	mi := &file_protos_outputs_decisionMaker_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Commands.ProtoReflect.Descriptor instead.
func (*Commands) Descriptor() ([]byte, []int) {
	return file_protos_outputs_decisionMaker_proto_rawDescGZIP(), []int{2}
}

func (x *Commands) GetBuildCommands() []*BuildCommandsOutput {
	if x != nil {
		return x.BuildCommands
	}
	return nil
}

func (x *Commands) GetStartUpCommands() []*StartUpCommandsOutput {
	if x != nil {
		return x.StartUpCommands
	}
	return nil
}

func (x *Commands) GetSeedCommands() []*SeedCommandsOutput {
	if x != nil {
		return x.SeedCommands
	}
	return nil
}

func (x *Commands) GetMigrationCommands() []*MigrationCommandsOutput {
	if x != nil {
		return x.MigrationCommands
	}
	return nil
}

type DecisionMakerInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LanguageSpecificDetection []*LanguageSpecificDetections `protobuf:"bytes,1,rep,name=languageSpecificDetection,proto3" json:"languageSpecificDetection,omitempty"`
	// GlobalDetections for global files
	GloabalDetections *GlobalDetections `protobuf:"bytes,2,opt,name=gloabal_detections,json=gloabalDetections,proto3" json:"gloabal_detections,omitempty"`
}

func (x *DecisionMakerInput) Reset() {
	*x = DecisionMakerInput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_outputs_decisionMaker_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DecisionMakerInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DecisionMakerInput) ProtoMessage() {}

func (x *DecisionMakerInput) ProtoReflect() protoreflect.Message {
	mi := &file_protos_outputs_decisionMaker_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DecisionMakerInput.ProtoReflect.Descriptor instead.
func (*DecisionMakerInput) Descriptor() ([]byte, []int) {
	return file_protos_outputs_decisionMaker_proto_rawDescGZIP(), []int{3}
}

func (x *DecisionMakerInput) GetLanguageSpecificDetection() []*LanguageSpecificDetections {
	if x != nil {
		return x.LanguageSpecificDetection
	}
	return nil
}

func (x *DecisionMakerInput) GetGloabalDetections() *GlobalDetections {
	if x != nil {
		return x.GloabalDetections
	}
	return nil
}

var File_protos_outputs_decisionMaker_proto protoreflect.FileDescriptor

var file_protos_outputs_decisionMaker_proto_rawDesc = []byte{
	0x0a, 0x22, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x73,
	0x2f, 0x64, 0x65, 0x63, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x4d, 0x61, 0x6b, 0x65, 0x72, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x29, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x6f, 0x75, 0x74,
	0x70, 0x75, 0x74, 0x73, 0x2f, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x2f, 0x62, 0x75, 0x69, 0x6c,
	0x64, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x29, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x73, 0x2f,
	0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x2f, 0x64, 0x6f, 0x63, 0x6b, 0x65, 0x72, 0x63, 0x6f, 0x6d,
	0x70, 0x6f, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x26, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x73, 0x2f, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x73, 0x2f, 0x67, 0x6c, 0x6f, 0x62, 0x61,
	0x6c, 0x2f, 0x64, 0x6f, 0x63, 0x6b, 0x65, 0x72, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x24, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x6f, 0x75, 0x74, 0x70, 0x75,
	0x74, 0x73, 0x2f, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x2f, 0x6d, 0x61, 0x6b, 0x65, 0x66, 0x69,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73,
	0x2f, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x73, 0x2f, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x2f,
	0x6d, 0x69, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f,
	0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x73, 0x2f, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x2f, 0x70,
	0x72, 0x6f, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x28, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x73, 0x2f, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x73, 0x2f, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c,
	0x2f, 0x73, 0x65, 0x65, 0x64, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x2b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x6f, 0x75, 0x74, 0x70,
	0x75, 0x74, 0x73, 0x2f, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x2f, 0x73, 0x74, 0x61, 0x72, 0x74,
	0x55, 0x70, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x73,
	0x2f, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x53, 0x70, 0x65, 0x63, 0x69, 0x66, 0x69,
	0x63, 0x2f, 0x61, 0x70, 0x70, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x34, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74,
	0x73, 0x2f, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x53, 0x70, 0x65, 0x63, 0x69, 0x66,
	0x69, 0x63, 0x2f, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72,
	0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x28, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f,
	0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x73, 0x2f, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65,
	0x53, 0x70, 0x65, 0x63, 0x69, 0x66, 0x69, 0x63, 0x2f, 0x64, 0x62, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x32, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74,
	0x73, 0x2f, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x53, 0x70, 0x65, 0x63, 0x69, 0x66,
	0x69, 0x63, 0x2f, 0x64, 0x65, 0x70, 0x65, 0x6e, 0x64, 0x65, 0x6e, 0x63, 0x69, 0x65, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x29, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x6f, 0x75,
	0x74, 0x70, 0x75, 0x74, 0x73, 0x2f, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x53, 0x70,
	0x65, 0x63, 0x69, 0x66, 0x69, 0x63, 0x2f, 0x65, 0x6e, 0x76, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x73,
	0x2f, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x53, 0x70, 0x65, 0x63, 0x69, 0x66, 0x69,
	0x63, 0x2f, 0x66, 0x72, 0x61, 0x6d, 0x65, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74,
	0x73, 0x2f, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x53, 0x70, 0x65, 0x63, 0x69, 0x66,
	0x69, 0x63, 0x2f, 0x6c, 0x69, 0x62, 0x72, 0x61, 0x72, 0x69, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x29, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x6f, 0x75, 0x74, 0x70, 0x75,
	0x74, 0x73, 0x2f, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x53, 0x70, 0x65, 0x63, 0x69,
	0x66, 0x69, 0x63, 0x2f, 0x6f, 0x72, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2b, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x73, 0x2f, 0x6c, 0x61,
	0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x53, 0x70, 0x65, 0x63, 0x69, 0x66, 0x69, 0x63, 0x2f, 0x73,
	0x74, 0x61, 0x63, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x32, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x73, 0x2f, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x73, 0x2f, 0x6c, 0x61, 0x6e, 0x67, 0x75,
	0x61, 0x67, 0x65, 0x53, 0x70, 0x65, 0x63, 0x69, 0x66, 0x69, 0x63, 0x2f, 0x73, 0x74, 0x61, 0x74,
	0x69, 0x63, 0x41, 0x73, 0x73, 0x65, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc7,
	0x04, 0x0a, 0x1a, 0x4c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x53, 0x70, 0x65, 0x63, 0x69,
	0x66, 0x69, 0x63, 0x44, 0x65, 0x74, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x26, 0x0a, 0x0e, 0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x56, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x72, 0x75, 0x6e, 0x74, 0x69,
	0x6d, 0x65, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x1c, 0x0a, 0x03, 0x65, 0x6e, 0x76,
	0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x45, 0x6e, 0x76, 0x4f, 0x75, 0x74, 0x70,
	0x75, 0x74, 0x52, 0x03, 0x65, 0x6e, 0x76, 0x12, 0x2e, 0x0a, 0x09, 0x66, 0x72, 0x61, 0x6d, 0x65,
	0x77, 0x6f, 0x72, 0x6b, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x46, 0x72, 0x61,
	0x6d, 0x65, 0x77, 0x6f, 0x72, 0x6b, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x52, 0x09, 0x66, 0x72,
	0x61, 0x6d, 0x65, 0x77, 0x6f, 0x72, 0x6b, 0x12, 0x19, 0x0a, 0x02, 0x64, 0x62, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x44, 0x42, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x52, 0x02,
	0x64, 0x62, 0x12, 0x1c, 0x0a, 0x03, 0x6f, 0x72, 0x6d, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0a, 0x2e, 0x4f, 0x72, 0x6d, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x52, 0x03, 0x6f, 0x72, 0x6d,
	0x12, 0x37, 0x0a, 0x0c, 0x64, 0x65, 0x70, 0x65, 0x6e, 0x64, 0x65, 0x6e, 0x63, 0x69, 0x65, 0x73,
	0x18, 0x07, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x44, 0x65, 0x70, 0x65, 0x6e, 0x64, 0x65,
	0x6e, 0x63, 0x69, 0x65, 0x73, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x52, 0x0c, 0x64, 0x65, 0x70,
	0x65, 0x6e, 0x64, 0x65, 0x6e, 0x63, 0x69, 0x65, 0x73, 0x12, 0x2e, 0x0a, 0x09, 0x6c, 0x69, 0x62,
	0x72, 0x61, 0x72, 0x69, 0x65, 0x73, 0x18, 0x08, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x4c,
	0x69, 0x62, 0x72, 0x61, 0x72, 0x69, 0x65, 0x73, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x52, 0x09,
	0x6c, 0x69, 0x62, 0x72, 0x61, 0x72, 0x69, 0x65, 0x73, 0x12, 0x37, 0x0a, 0x0c, 0x73, 0x74, 0x61,
	0x74, 0x69, 0x63, 0x41, 0x73, 0x73, 0x65, 0x74, 0x73, 0x18, 0x09, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x13, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x69, 0x63, 0x41, 0x73, 0x73, 0x65, 0x74, 0x73, 0x4f, 0x75,
	0x74, 0x70, 0x75, 0x74, 0x52, 0x0c, 0x73, 0x74, 0x61, 0x74, 0x69, 0x63, 0x41, 0x73, 0x73, 0x65,
	0x74, 0x73, 0x12, 0x2e, 0x0a, 0x0b, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x4f, 0x75, 0x74, 0x70, 0x75,
	0x74, 0x18, 0x0a, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x53, 0x74, 0x61, 0x63, 0x6b, 0x4f,
	0x75, 0x74, 0x70, 0x75, 0x74, 0x52, 0x0b, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x4f, 0x75, 0x74, 0x70,
	0x75, 0x74, 0x12, 0x2e, 0x0a, 0x09, 0x61, 0x70, 0x70, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x18,
	0x0b, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x41, 0x70, 0x70, 0x73, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x52, 0x09, 0x61, 0x70, 0x70, 0x73, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x12, 0x3d, 0x0a, 0x0e, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x44, 0x69, 0x72, 0x65, 0x63,
	0x74, 0x6f, 0x72, 0x79, 0x18, 0x0c, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x42, 0x75, 0x69,
	0x6c, 0x64, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x79, 0x4f, 0x75, 0x74, 0x70, 0x75,
	0x74, 0x52, 0x0e, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72,
	0x79, 0x12, 0x25, 0x0a, 0x08, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x73, 0x18, 0x0d, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x73, 0x52, 0x08,
	0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x73, 0x22, 0xe7, 0x01, 0x0a, 0x10, 0x47, 0x6c, 0x6f,
	0x62, 0x61, 0x6c, 0x44, 0x65, 0x74, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x2b, 0x0a,
	0x08, 0x70, 0x72, 0x6f, 0x63, 0x46, 0x69, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x0f, 0x2e, 0x50, 0x72, 0x6f, 0x63, 0x46, 0x69, 0x6c, 0x65, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74,
	0x52, 0x08, 0x70, 0x72, 0x6f, 0x63, 0x46, 0x69, 0x6c, 0x65, 0x12, 0x2b, 0x0a, 0x08, 0x6d, 0x61,
	0x6b, 0x65, 0x66, 0x69, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x4d,
	0x61, 0x6b, 0x65, 0x66, 0x69, 0x6c, 0x65, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x52, 0x08, 0x6d,
	0x61, 0x6b, 0x65, 0x66, 0x69, 0x6c, 0x65, 0x12, 0x31, 0x0a, 0x0a, 0x64, 0x6f, 0x63, 0x6b, 0x65,
	0x72, 0x46, 0x69, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x44, 0x6f,
	0x63, 0x6b, 0x65, 0x72, 0x46, 0x69, 0x6c, 0x65, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x52, 0x0a,
	0x64, 0x6f, 0x63, 0x6b, 0x65, 0x72, 0x46, 0x69, 0x6c, 0x65, 0x12, 0x46, 0x0a, 0x11, 0x64, 0x6f,
	0x63, 0x6b, 0x65, 0x72, 0x43, 0x6f, 0x6d, 0x70, 0x6f, 0x73, 0x65, 0x46, 0x69, 0x6c, 0x65, 0x18,
	0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x44, 0x6f, 0x63, 0x6b, 0x65, 0x72, 0x43, 0x6f,
	0x6d, 0x70, 0x6f, 0x73, 0x65, 0x46, 0x69, 0x6c, 0x65, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x52,
	0x11, 0x64, 0x6f, 0x63, 0x6b, 0x65, 0x72, 0x43, 0x6f, 0x6d, 0x70, 0x6f, 0x73, 0x65, 0x46, 0x69,
	0x6c, 0x65, 0x22, 0x89, 0x02, 0x0a, 0x08, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x73, 0x12,
	0x3a, 0x0a, 0x0d, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x43, 0x6f,
	0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x73, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x52, 0x0d, 0x62, 0x75,
	0x69, 0x6c, 0x64, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x73, 0x12, 0x40, 0x0a, 0x0f, 0x73,
	0x74, 0x61, 0x72, 0x74, 0x55, 0x70, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x73, 0x18, 0x02,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x53, 0x74, 0x61, 0x72, 0x74, 0x55, 0x70, 0x43, 0x6f,
	0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x73, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x52, 0x0f, 0x73, 0x74,
	0x61, 0x72, 0x74, 0x55, 0x70, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x73, 0x12, 0x37, 0x0a,
	0x0c, 0x73, 0x65, 0x65, 0x64, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x73, 0x18, 0x03, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x53, 0x65, 0x65, 0x64, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e,
	0x64, 0x73, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x52, 0x0c, 0x73, 0x65, 0x65, 0x64, 0x43, 0x6f,
	0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x73, 0x12, 0x46, 0x0a, 0x11, 0x4d, 0x69, 0x67, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x18, 0x2e, 0x4d, 0x69, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6d,
	0x6d, 0x61, 0x6e, 0x64, 0x73, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x52, 0x11, 0x6d, 0x69, 0x67,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x73, 0x22, 0xb1,
	0x01, 0x0a, 0x12, 0x44, 0x65, 0x63, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x4d, 0x61, 0x6b, 0x65, 0x72,
	0x49, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x59, 0x0a, 0x19, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67,
	0x65, 0x53, 0x70, 0x65, 0x63, 0x69, 0x66, 0x69, 0x63, 0x44, 0x65, 0x74, 0x65, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x4c, 0x61, 0x6e, 0x67, 0x75,
	0x61, 0x67, 0x65, 0x53, 0x70, 0x65, 0x63, 0x69, 0x66, 0x69, 0x63, 0x44, 0x65, 0x74, 0x65, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x19, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x53,
	0x70, 0x65, 0x63, 0x69, 0x66, 0x69, 0x63, 0x44, 0x65, 0x74, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x40, 0x0a, 0x12, 0x67, 0x6c, 0x6f, 0x61, 0x62, 0x61, 0x6c, 0x5f, 0x64, 0x65, 0x74, 0x65,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x47,
	0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x44, 0x65, 0x74, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52,
	0x11, 0x67, 0x6c, 0x6f, 0x61, 0x62, 0x61, 0x6c, 0x44, 0x65, 0x74, 0x65, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protos_outputs_decisionMaker_proto_rawDescOnce sync.Once
	file_protos_outputs_decisionMaker_proto_rawDescData = file_protos_outputs_decisionMaker_proto_rawDesc
)

func file_protos_outputs_decisionMaker_proto_rawDescGZIP() []byte {
	file_protos_outputs_decisionMaker_proto_rawDescOnce.Do(func() {
		file_protos_outputs_decisionMaker_proto_rawDescData = protoimpl.X.CompressGZIP(file_protos_outputs_decisionMaker_proto_rawDescData)
	})
	return file_protos_outputs_decisionMaker_proto_rawDescData
}

var file_protos_outputs_decisionMaker_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_protos_outputs_decisionMaker_proto_goTypes = []interface{}{
	(*LanguageSpecificDetections)(nil), // 0: LanguageSpecificDetections
	(*GlobalDetections)(nil),           // 1: GlobalDetections
	(*Commands)(nil),                   // 2: Commands
	(*DecisionMakerInput)(nil),         // 3: DecisionMakerInput
	(*EnvOutput)(nil),                  // 4: EnvOutput
	(*FrameworkOutput)(nil),            // 5: FrameworkOutput
	(*DBOutput)(nil),                   // 6: DBOutput
	(*OrmOutput)(nil),                  // 7: OrmOutput
	(*DependenciesOutput)(nil),         // 8: DependenciesOutput
	(*LibrariesOutput)(nil),            // 9: LibrariesOutput
	(*StaticAssetsOutput)(nil),         // 10: StaticAssetsOutput
	(*StackOutput)(nil),                // 11: StackOutput
	(*AppserverOutput)(nil),            // 12: AppserverOutput
	(*BuildDirectoryOutput)(nil),       // 13: BuildDirectoryOutput
	(*ProcFileOutput)(nil),             // 14: ProcFileOutput
	(*MakefileOutput)(nil),             // 15: MakefileOutput
	(*DockerFileOutput)(nil),           // 16: DockerFileOutput
	(*DockerComposeFileOutput)(nil),    // 17: DockerComposeFileOutput
	(*BuildCommandsOutput)(nil),        // 18: BuildCommandsOutput
	(*StartUpCommandsOutput)(nil),      // 19: StartUpCommandsOutput
	(*SeedCommandsOutput)(nil),         // 20: SeedCommandsOutput
	(*MigrationCommandsOutput)(nil),    // 21: MigrationCommandsOutput
}
var file_protos_outputs_decisionMaker_proto_depIdxs = []int32{
	4,  // 0: LanguageSpecificDetections.env:type_name -> EnvOutput
	5,  // 1: LanguageSpecificDetections.framework:type_name -> FrameworkOutput
	6,  // 2: LanguageSpecificDetections.db:type_name -> DBOutput
	7,  // 3: LanguageSpecificDetections.orm:type_name -> OrmOutput
	8,  // 4: LanguageSpecificDetections.dependencies:type_name -> DependenciesOutput
	9,  // 5: LanguageSpecificDetections.libraries:type_name -> LibrariesOutput
	10, // 6: LanguageSpecificDetections.staticAssets:type_name -> StaticAssetsOutput
	11, // 7: LanguageSpecificDetections.stackOutput:type_name -> StackOutput
	12, // 8: LanguageSpecificDetections.appserver:type_name -> AppserverOutput
	13, // 9: LanguageSpecificDetections.buildDirectory:type_name -> BuildDirectoryOutput
	2,  // 10: LanguageSpecificDetections.commands:type_name -> Commands
	14, // 11: GlobalDetections.procFile:type_name -> ProcFileOutput
	15, // 12: GlobalDetections.makefile:type_name -> MakefileOutput
	16, // 13: GlobalDetections.dockerFile:type_name -> DockerFileOutput
	17, // 14: GlobalDetections.dockerComposeFile:type_name -> DockerComposeFileOutput
	18, // 15: Commands.buildCommands:type_name -> BuildCommandsOutput
	19, // 16: Commands.startUpCommands:type_name -> StartUpCommandsOutput
	20, // 17: Commands.seedCommands:type_name -> SeedCommandsOutput
	21, // 18: Commands.MigrationCommands:type_name -> MigrationCommandsOutput
	0,  // 19: DecisionMakerInput.languageSpecificDetection:type_name -> LanguageSpecificDetections
	1,  // 20: DecisionMakerInput.gloabal_detections:type_name -> GlobalDetections
	21, // [21:21] is the sub-list for method output_type
	21, // [21:21] is the sub-list for method input_type
	21, // [21:21] is the sub-list for extension type_name
	21, // [21:21] is the sub-list for extension extendee
	0,  // [0:21] is the sub-list for field type_name
}

func init() { file_protos_outputs_decisionMaker_proto_init() }
func file_protos_outputs_decisionMaker_proto_init() {
	if File_protos_outputs_decisionMaker_proto != nil {
		return
	}
	file_protos_outputs_global_buildCommands_proto_init()
	file_protos_outputs_global_dockercompose_proto_init()
	file_protos_outputs_global_dockerfile_proto_init()
	file_protos_outputs_global_makefile_proto_init()
	file_protos_outputs_global_migrationCommands_proto_init()
	file_protos_outputs_global_proc_proto_init()
	file_protos_outputs_global_seedCommands_proto_init()
	file_protos_outputs_global_startUpCommands_proto_init()
	file_protos_outputs_languageSpecific_appserver_proto_init()
	file_protos_outputs_languageSpecific_buildDirectory_proto_init()
	file_protos_outputs_languageSpecific_db_proto_init()
	file_protos_outputs_languageSpecific_dependencies_proto_init()
	file_protos_outputs_languageSpecific_env_proto_init()
	file_protos_outputs_languageSpecific_framework_proto_init()
	file_protos_outputs_languageSpecific_libraries_proto_init()
	file_protos_outputs_languageSpecific_orm_proto_init()
	file_protos_outputs_languageSpecific_stack_proto_init()
	file_protos_outputs_languageSpecific_staticAssets_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_protos_outputs_decisionMaker_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LanguageSpecificDetections); i {
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
		file_protos_outputs_decisionMaker_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GlobalDetections); i {
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
		file_protos_outputs_decisionMaker_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Commands); i {
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
		file_protos_outputs_decisionMaker_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DecisionMakerInput); i {
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
			RawDescriptor: file_protos_outputs_decisionMaker_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_protos_outputs_decisionMaker_proto_goTypes,
		DependencyIndexes: file_protos_outputs_decisionMaker_proto_depIdxs,
		MessageInfos:      file_protos_outputs_decisionMaker_proto_msgTypes,
	}.Build()
	File_protos_outputs_decisionMaker_proto = out.File
	file_protos_outputs_decisionMaker_proto_rawDesc = nil
	file_protos_outputs_decisionMaker_proto_goTypes = nil
	file_protos_outputs_decisionMaker_proto_depIdxs = nil
}
