// Code generated by protoc-gen-go. DO NOT EDIT.
// source: protos/outputs/languageSpecific/stack.proto

package protos

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type StackOutput struct {
}

func (m *StackOutput) Reset()                    { *m = StackOutput{} }
func (m *StackOutput) String() string            { return proto.CompactTextString(m) }
func (*StackOutput) ProtoMessage()               {}
func (*StackOutput) Descriptor() ([]byte, []int) { return fileDescriptor18, []int{0} }

func init() {
	proto.RegisterType((*StackOutput)(nil), "StackOutput")
}

func init() { proto.RegisterFile("protos/outputs/languageSpecific/stack.proto", fileDescriptor18) }

var fileDescriptor18 = []byte{
	// 88 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xd2, 0x2e, 0x28, 0xca, 0x2f,
	0xc9, 0x2f, 0xd6, 0xcf, 0x2f, 0x2d, 0x29, 0x28, 0x2d, 0x29, 0xd6, 0xcf, 0x49, 0xcc, 0x4b, 0x2f,
	0x4d, 0x4c, 0x4f, 0x0d, 0x2e, 0x48, 0x4d, 0xce, 0x4c, 0xcb, 0x4c, 0xd6, 0x2f, 0x2e, 0x49, 0x4c,
	0xce, 0xd6, 0x03, 0xab, 0x52, 0xe2, 0xe5, 0xe2, 0x0e, 0x06, 0x71, 0xfd, 0xc1, 0x8a, 0x9d, 0xb8,
	0xa2, 0x38, 0xf4, 0xf4, 0x21, 0xfa, 0x93, 0xd8, 0xc0, 0xb4, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff,
	0xd4, 0x7b, 0xf2, 0xef, 0x50, 0x00, 0x00, 0x00,
}
