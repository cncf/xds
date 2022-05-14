// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.4
// source: xds/core/v3/context_params.proto

package v3

import (
	_ "github.com/cncf/xds/go/xds/annotations/v3"
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

type ContextParams struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Params map[string]string `protobuf:"bytes,1,rep,name=params,proto3" json:"params,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *ContextParams) Reset() {
	*x = ContextParams{}
	if protoimpl.UnsafeEnabled {
		mi := &file_xds_core_v3_context_params_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ContextParams) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ContextParams) ProtoMessage() {}

func (x *ContextParams) ProtoReflect() protoreflect.Message {
	mi := &file_xds_core_v3_context_params_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ContextParams.ProtoReflect.Descriptor instead.
func (*ContextParams) Descriptor() ([]byte, []int) {
	return file_xds_core_v3_context_params_proto_rawDescGZIP(), []int{0}
}

func (x *ContextParams) GetParams() map[string]string {
	if x != nil {
		return x.Params
	}
	return nil
}

var File_xds_core_v3_context_params_proto protoreflect.FileDescriptor

var file_xds_core_v3_context_params_proto_rawDesc = []byte{
	0x0a, 0x20, 0x78, 0x64, 0x73, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x76, 0x33, 0x2f, 0x63, 0x6f,
	0x6e, 0x74, 0x65, 0x78, 0x74, 0x5f, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x0b, 0x78, 0x64, 0x73, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x33, 0x1a,
	0x1f, 0x78, 0x64, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2f, 0x76, 0x33, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x8a, 0x01, 0x0a, 0x0d, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x50, 0x61, 0x72, 0x61,
	0x6d, 0x73, 0x12, 0x3e, 0x0a, 0x06, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x26, 0x2e, 0x78, 0x64, 0x73, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x33,
	0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x2e, 0x50,
	0x61, 0x72, 0x61, 0x6d, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x06, 0x70, 0x61, 0x72, 0x61,
	0x6d, 0x73, 0x1a, 0x39, 0x0a, 0x0b, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x42, 0x5a, 0x0a,
	0x16, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x78, 0x64, 0x73, 0x2e,
	0x63, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x33, 0x42, 0x12, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74,
	0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x22, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6e, 0x63, 0x66, 0x2f, 0x78,
	0x64, 0x73, 0x2f, 0x67, 0x6f, 0x2f, 0x78, 0x64, 0x73, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x76,
	0x33, 0xd2, 0xc6, 0xa4, 0xe1, 0x06, 0x02, 0x08, 0x01, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_xds_core_v3_context_params_proto_rawDescOnce sync.Once
	file_xds_core_v3_context_params_proto_rawDescData = file_xds_core_v3_context_params_proto_rawDesc
)

func file_xds_core_v3_context_params_proto_rawDescGZIP() []byte {
	file_xds_core_v3_context_params_proto_rawDescOnce.Do(func() {
		file_xds_core_v3_context_params_proto_rawDescData = protoimpl.X.CompressGZIP(file_xds_core_v3_context_params_proto_rawDescData)
	})
	return file_xds_core_v3_context_params_proto_rawDescData
}

var file_xds_core_v3_context_params_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_xds_core_v3_context_params_proto_goTypes = []interface{}{
	(*ContextParams)(nil), // 0: xds.core.v3.ContextParams
	nil,                   // 1: xds.core.v3.ContextParams.ParamsEntry
}
var file_xds_core_v3_context_params_proto_depIdxs = []int32{
	1, // 0: xds.core.v3.ContextParams.params:type_name -> xds.core.v3.ContextParams.ParamsEntry
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_xds_core_v3_context_params_proto_init() }
func file_xds_core_v3_context_params_proto_init() {
	if File_xds_core_v3_context_params_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_xds_core_v3_context_params_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ContextParams); i {
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
			RawDescriptor: file_xds_core_v3_context_params_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_xds_core_v3_context_params_proto_goTypes,
		DependencyIndexes: file_xds_core_v3_context_params_proto_depIdxs,
		MessageInfos:      file_xds_core_v3_context_params_proto_msgTypes,
	}.Build()
	File_xds_core_v3_context_params_proto = out.File
	file_xds_core_v3_context_params_proto_rawDesc = nil
	file_xds_core_v3_context_params_proto_goTypes = nil
	file_xds_core_v3_context_params_proto_depIdxs = nil
}
