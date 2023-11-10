// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: xds/core/v3/resource_name.proto

package v3

import (
	_ "github.com/cncf/xds/go/xds/annotations/v3"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
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

// xDS resource name. This has a canonical xdstp:// URI representation:
//
//   xdstp://{authority}/{type_url}/{id}?{context_params}
//
// where context_params take the form of URI query parameters.
//
// A xDS resource name fully identifies a network resource for transport
// purposes. xDS resource names in this form appear only in discovery
// request/response messages used with the xDS transport.
type ResourceName struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Opaque identifier for the resource. Any '/' will not be escaped during URI
	// encoding and will form part of the URI path.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Logical authority for resource (not necessarily transport network address).
	// Authorities are opaque in the xDS API, data-plane load balancers will map
	// them to concrete network transports such as an xDS management server.
	Authority string `protobuf:"bytes,2,opt,name=authority,proto3" json:"authority,omitempty"`
	// Fully qualified resource type (as in type URL without types.googleapis.com/
	// prefix).
	ResourceType string `protobuf:"bytes,3,opt,name=resource_type,json=resourceType,proto3" json:"resource_type,omitempty"`
	// Additional parameters that can be used to select resource variants.
	Context *ContextParams `protobuf:"bytes,4,opt,name=context,proto3" json:"context,omitempty"`
}

func (x *ResourceName) Reset() {
	*x = ResourceName{}
	if protoimpl.UnsafeEnabled {
		mi := &file_xds_core_v3_resource_name_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResourceName) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResourceName) ProtoMessage() {}

func (x *ResourceName) ProtoReflect() protoreflect.Message {
	mi := &file_xds_core_v3_resource_name_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResourceName.ProtoReflect.Descriptor instead.
func (*ResourceName) Descriptor() ([]byte, []int) {
	return file_xds_core_v3_resource_name_proto_rawDescGZIP(), []int{0}
}

func (x *ResourceName) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ResourceName) GetAuthority() string {
	if x != nil {
		return x.Authority
	}
	return ""
}

func (x *ResourceName) GetResourceType() string {
	if x != nil {
		return x.ResourceType
	}
	return ""
}

func (x *ResourceName) GetContext() *ContextParams {
	if x != nil {
		return x.Context
	}
	return nil
}

var File_xds_core_v3_resource_name_proto protoreflect.FileDescriptor

var file_xds_core_v3_resource_name_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x78, 0x64, 0x73, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x76, 0x33, 0x2f, 0x72, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x0b, 0x78, 0x64, 0x73, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x33, 0x1a, 0x1f,
	0x78, 0x64, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f,
	0x76, 0x33, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x20, 0x78, 0x64, 0x73, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x76, 0x33, 0x2f, 0x63, 0x6f, 0x6e,
	0x74, 0x65, 0x78, 0x74, 0x5f, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69,
	0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa0, 0x01, 0x0a, 0x0c, 0x52,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x61,
	0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x12, 0x2c, 0x0a, 0x0d, 0x72, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x01, 0x52, 0x0c, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x34, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65,
	0x78, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x78, 0x64, 0x73, 0x2e, 0x63,
	0x6f, 0x72, 0x65, 0x2e, 0x76, 0x33, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x50, 0x61,
	0x72, 0x61, 0x6d, 0x73, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x42, 0x9e, 0x01,
	0x0a, 0x0f, 0x63, 0x6f, 0x6d, 0x2e, 0x78, 0x64, 0x73, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x76,
	0x33, 0x42, 0x11, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x22, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x63, 0x6e, 0x63, 0x66, 0x2f, 0x78, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x2f, 0x78,
	0x64, 0x73, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x76, 0x33, 0xa2, 0x02, 0x03, 0x58, 0x43, 0x58,
	0xaa, 0x02, 0x0b, 0x58, 0x64, 0x73, 0x2e, 0x43, 0x6f, 0x72, 0x65, 0x2e, 0x56, 0x33, 0xca, 0x02,
	0x0b, 0x58, 0x64, 0x73, 0x5c, 0x43, 0x6f, 0x72, 0x65, 0x5c, 0x56, 0x33, 0xe2, 0x02, 0x17, 0x58,
	0x64, 0x73, 0x5c, 0x43, 0x6f, 0x72, 0x65, 0x5c, 0x56, 0x33, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x0d, 0x58, 0x64, 0x73, 0x3a, 0x3a, 0x43, 0x6f,
	0x72, 0x65, 0x3a, 0x3a, 0x56, 0x33, 0xd2, 0xc6, 0xa4, 0xe1, 0x06, 0x02, 0x08, 0x01, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_xds_core_v3_resource_name_proto_rawDescOnce sync.Once
	file_xds_core_v3_resource_name_proto_rawDescData = file_xds_core_v3_resource_name_proto_rawDesc
)

func file_xds_core_v3_resource_name_proto_rawDescGZIP() []byte {
	file_xds_core_v3_resource_name_proto_rawDescOnce.Do(func() {
		file_xds_core_v3_resource_name_proto_rawDescData = protoimpl.X.CompressGZIP(file_xds_core_v3_resource_name_proto_rawDescData)
	})
	return file_xds_core_v3_resource_name_proto_rawDescData
}

var file_xds_core_v3_resource_name_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_xds_core_v3_resource_name_proto_goTypes = []interface{}{
	(*ResourceName)(nil),  // 0: xds.core.v3.ResourceName
	(*ContextParams)(nil), // 1: xds.core.v3.ContextParams
}
var file_xds_core_v3_resource_name_proto_depIdxs = []int32{
	1, // 0: xds.core.v3.ResourceName.context:type_name -> xds.core.v3.ContextParams
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_xds_core_v3_resource_name_proto_init() }
func file_xds_core_v3_resource_name_proto_init() {
	if File_xds_core_v3_resource_name_proto != nil {
		return
	}
	file_xds_core_v3_context_params_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_xds_core_v3_resource_name_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResourceName); i {
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
			RawDescriptor: file_xds_core_v3_resource_name_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_xds_core_v3_resource_name_proto_goTypes,
		DependencyIndexes: file_xds_core_v3_resource_name_proto_depIdxs,
		MessageInfos:      file_xds_core_v3_resource_name_proto_msgTypes,
	}.Build()
	File_xds_core_v3_resource_name_proto = out.File
	file_xds_core_v3_resource_name_proto_rawDesc = nil
	file_xds_core_v3_resource_name_proto_goTypes = nil
	file_xds_core_v3_resource_name_proto_depIdxs = nil
}
