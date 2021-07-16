// Code generated by protoc-gen-go. DO NOT EDIT.
// source: xds/core/v3/extension.proto

package xds_core_v3

import (
	fmt "fmt"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
	any "github.com/golang/protobuf/ptypes/any"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type TypedExtensionConfig struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	TypedConfig          *any.Any `protobuf:"bytes,2,opt,name=typed_config,json=typedConfig,proto3" json:"typed_config,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TypedExtensionConfig) Reset()         { *m = TypedExtensionConfig{} }
func (m *TypedExtensionConfig) String() string { return proto.CompactTextString(m) }
func (*TypedExtensionConfig) ProtoMessage()    {}
func (*TypedExtensionConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_ce7d8620783f5d54, []int{0}
}

func (m *TypedExtensionConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TypedExtensionConfig.Unmarshal(m, b)
}
func (m *TypedExtensionConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TypedExtensionConfig.Marshal(b, m, deterministic)
}
func (m *TypedExtensionConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TypedExtensionConfig.Merge(m, src)
}
func (m *TypedExtensionConfig) XXX_Size() int {
	return xxx_messageInfo_TypedExtensionConfig.Size(m)
}
func (m *TypedExtensionConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_TypedExtensionConfig.DiscardUnknown(m)
}

var xxx_messageInfo_TypedExtensionConfig proto.InternalMessageInfo

func (m *TypedExtensionConfig) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *TypedExtensionConfig) GetTypedConfig() *any.Any {
	if m != nil {
		return m.TypedConfig
	}
	return nil
}

func init() {
	proto.RegisterType((*TypedExtensionConfig)(nil), "xds.core.v3.TypedExtensionConfig")
}

func init() { proto.RegisterFile("xds/core/v3/extension.proto", fileDescriptor_ce7d8620783f5d54) }

var fileDescriptor_ce7d8620783f5d54 = []byte{
	// 226 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x44, 0x8f, 0x51, 0x4b, 0xc3, 0x30,
	0x14, 0x85, 0x49, 0x99, 0x3a, 0x53, 0x11, 0x29, 0x03, 0xe7, 0xfa, 0x32, 0x7c, 0xda, 0xd3, 0x0d,
	0xda, 0x5f, 0xb0, 0x88, 0xef, 0x63, 0xf8, 0x2e, 0x69, 0x93, 0xc5, 0xc0, 0x96, 0x5b, 0xda, 0x34,
	0xb4, 0x7f, 0xc7, 0x9f, 0xd9, 0x27, 0x69, 0x42, 0xdd, 0x5b, 0x72, 0xcf, 0x39, 0xdf, 0xb9, 0x97,
	0xe6, 0xbd, 0x6c, 0x59, 0x85, 0x8d, 0x62, 0xbe, 0x60, 0xaa, 0x77, 0xca, 0xb6, 0x06, 0x2d, 0xd4,
	0x0d, 0x3a, 0xcc, 0xd2, 0x5e, 0xb6, 0x30, 0x89, 0xe0, 0x8b, 0xcd, 0xb3, 0x17, 0x67, 0x23, 0x85,
	0x53, 0x6c, 0x7e, 0x44, 0xd7, 0xe6, 0x45, 0x23, 0xea, 0xb3, 0x62, 0xe1, 0x57, 0x76, 0x27, 0x26,
	0xec, 0x10, 0xa5, 0x57, 0x4f, 0x57, 0x5f, 0x43, 0xad, 0xe4, 0xe7, 0x0c, 0xfe, 0x40, 0x7b, 0x32,
	0x3a, 0xcb, 0xe9, 0xc2, 0x8a, 0x8b, 0x5a, 0x93, 0x2d, 0xd9, 0xdd, 0xf3, 0xbb, 0x91, 0x2f, 0x9a,
	0xe4, 0x89, 0x1c, 0xc3, 0x30, 0xdb, 0xd3, 0x07, 0x37, 0x85, 0xbe, 0xab, 0x60, 0x5e, 0x27, 0x5b,
	0xb2, 0x4b, 0xdf, 0x57, 0x10, 0x6b, 0x60, 0xae, 0x81, 0xbd, 0x1d, 0xf8, 0x72, 0xe4, 0x37, 0xbf,
	0x24, 0x59, 0x92, 0x63, 0x1a, 0x32, 0x91, 0xcf, 0x19, 0xcd, 0x2b, 0xbc, 0x80, 0x36, 0xee, 0xa7,
	0x2b, 0xa1, 0x93, 0xb5, 0x80, 0xeb, 0x29, 0x6f, 0xfc, 0xf1, 0x7f, 0x9f, 0xc3, 0x04, 0x3b, 0x90,
	0xf2, 0x36, 0x50, 0x8b, 0xbf, 0x00, 0x00, 0x00, 0xff, 0xff, 0xc9, 0x93, 0xca, 0x73, 0x0f, 0x01,
	0x00, 0x00,
}