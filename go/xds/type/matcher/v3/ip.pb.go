// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.4
// source: xds/type/matcher/v3/ip.proto

package v3

import (
	_ "github.com/cncf/xds/go/xds/annotations/v3"
	v3 "github.com/cncf/xds/go/xds/core/v3"
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

type IPMatcher struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RangeMatchers []*IPMatcher_IPRangeMatcher `protobuf:"bytes,1,rep,name=range_matchers,json=rangeMatchers,proto3" json:"range_matchers,omitempty"`
}

func (x *IPMatcher) Reset() {
	*x = IPMatcher{}
	if protoimpl.UnsafeEnabled {
		mi := &file_xds_type_matcher_v3_ip_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IPMatcher) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IPMatcher) ProtoMessage() {}

func (x *IPMatcher) ProtoReflect() protoreflect.Message {
	mi := &file_xds_type_matcher_v3_ip_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IPMatcher.ProtoReflect.Descriptor instead.
func (*IPMatcher) Descriptor() ([]byte, []int) {
	return file_xds_type_matcher_v3_ip_proto_rawDescGZIP(), []int{0}
}

func (x *IPMatcher) GetRangeMatchers() []*IPMatcher_IPRangeMatcher {
	if x != nil {
		return x.RangeMatchers
	}
	return nil
}

type IPMatcher_IPRangeMatcher struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ranges    []*v3.CidrRange  `protobuf:"bytes,1,rep,name=ranges,proto3" json:"ranges,omitempty"`
	OnMatch   *Matcher_OnMatch `protobuf:"bytes,2,opt,name=on_match,json=onMatch,proto3" json:"on_match,omitempty"`
	Exclusive bool             `protobuf:"varint,3,opt,name=exclusive,proto3" json:"exclusive,omitempty"`
}

func (x *IPMatcher_IPRangeMatcher) Reset() {
	*x = IPMatcher_IPRangeMatcher{}
	if protoimpl.UnsafeEnabled {
		mi := &file_xds_type_matcher_v3_ip_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IPMatcher_IPRangeMatcher) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IPMatcher_IPRangeMatcher) ProtoMessage() {}

func (x *IPMatcher_IPRangeMatcher) ProtoReflect() protoreflect.Message {
	mi := &file_xds_type_matcher_v3_ip_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IPMatcher_IPRangeMatcher.ProtoReflect.Descriptor instead.
func (*IPMatcher_IPRangeMatcher) Descriptor() ([]byte, []int) {
	return file_xds_type_matcher_v3_ip_proto_rawDescGZIP(), []int{0, 0}
}

func (x *IPMatcher_IPRangeMatcher) GetRanges() []*v3.CidrRange {
	if x != nil {
		return x.Ranges
	}
	return nil
}

func (x *IPMatcher_IPRangeMatcher) GetOnMatch() *Matcher_OnMatch {
	if x != nil {
		return x.OnMatch
	}
	return nil
}

func (x *IPMatcher_IPRangeMatcher) GetExclusive() bool {
	if x != nil {
		return x.Exclusive
	}
	return false
}

var File_xds_type_matcher_v3_ip_proto protoreflect.FileDescriptor

var file_xds_type_matcher_v3_ip_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x78, 0x64, 0x73, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x2f, 0x6d, 0x61, 0x74, 0x63, 0x68,
	0x65, 0x72, 0x2f, 0x76, 0x33, 0x2f, 0x69, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x13,
	0x78, 0x64, 0x73, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72,
	0x2e, 0x76, 0x33, 0x1a, 0x1f, 0x78, 0x64, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x76, 0x33, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x16, 0x78, 0x64, 0x73, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x76,
	0x33, 0x2f, 0x63, 0x69, 0x64, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x21, 0x78, 0x64,
	0x73, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x2f, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x2f, 0x76,
	0x33, 0x2f, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61,
	0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8d, 0x02, 0x0a, 0x09, 0x49, 0x50, 0x4d,
	0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x12, 0x54, 0x0a, 0x0e, 0x72, 0x61, 0x6e, 0x67, 0x65, 0x5f,
	0x6d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2d,
	0x2e, 0x78, 0x64, 0x73, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x65,
	0x72, 0x2e, 0x76, 0x33, 0x2e, 0x49, 0x50, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x2e, 0x49,
	0x50, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x52, 0x0d, 0x72,
	0x61, 0x6e, 0x67, 0x65, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x73, 0x1a, 0xa9, 0x01, 0x0a,
	0x0e, 0x49, 0x50, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x12,
	0x38, 0x0a, 0x06, 0x72, 0x61, 0x6e, 0x67, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x16, 0x2e, 0x78, 0x64, 0x73, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x33, 0x2e, 0x43, 0x69,
	0x64, 0x72, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x92, 0x01, 0x02, 0x08,
	0x01, 0x52, 0x06, 0x72, 0x61, 0x6e, 0x67, 0x65, 0x73, 0x12, 0x3f, 0x0a, 0x08, 0x6f, 0x6e, 0x5f,
	0x6d, 0x61, 0x74, 0x63, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x78, 0x64,
	0x73, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x2e, 0x76,
	0x33, 0x2e, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x2e, 0x4f, 0x6e, 0x4d, 0x61, 0x74, 0x63,
	0x68, 0x52, 0x07, 0x6f, 0x6e, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x12, 0x1c, 0x0a, 0x09, 0x65, 0x78,
	0x63, 0x6c, 0x75, 0x73, 0x69, 0x76, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x65,
	0x78, 0x63, 0x6c, 0x75, 0x73, 0x69, 0x76, 0x65, 0x42, 0x66, 0x0a, 0x1e, 0x63, 0x6f, 0x6d, 0x2e,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x78, 0x64, 0x73, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e,
	0x6d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x2e, 0x76, 0x33, 0x42, 0x0e, 0x49, 0x50, 0x4d, 0x61,
	0x74, 0x63, 0x68, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x2a, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6e, 0x63, 0x66, 0x2f, 0x78, 0x64,
	0x73, 0x2f, 0x67, 0x6f, 0x2f, 0x78, 0x64, 0x73, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x2f, 0x6d, 0x61,
	0x74, 0x63, 0x68, 0x65, 0x72, 0x2f, 0x76, 0x33, 0xd2, 0xc6, 0xa4, 0xe1, 0x06, 0x02, 0x08, 0x01,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_xds_type_matcher_v3_ip_proto_rawDescOnce sync.Once
	file_xds_type_matcher_v3_ip_proto_rawDescData = file_xds_type_matcher_v3_ip_proto_rawDesc
)

func file_xds_type_matcher_v3_ip_proto_rawDescGZIP() []byte {
	file_xds_type_matcher_v3_ip_proto_rawDescOnce.Do(func() {
		file_xds_type_matcher_v3_ip_proto_rawDescData = protoimpl.X.CompressGZIP(file_xds_type_matcher_v3_ip_proto_rawDescData)
	})
	return file_xds_type_matcher_v3_ip_proto_rawDescData
}

var file_xds_type_matcher_v3_ip_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_xds_type_matcher_v3_ip_proto_goTypes = []interface{}{
	(*IPMatcher)(nil),                // 0: xds.type.matcher.v3.IPMatcher
	(*IPMatcher_IPRangeMatcher)(nil), // 1: xds.type.matcher.v3.IPMatcher.IPRangeMatcher
	(*v3.CidrRange)(nil),             // 2: xds.core.v3.CidrRange
	(*Matcher_OnMatch)(nil),          // 3: xds.type.matcher.v3.Matcher.OnMatch
}
var file_xds_type_matcher_v3_ip_proto_depIdxs = []int32{
	1, // 0: xds.type.matcher.v3.IPMatcher.range_matchers:type_name -> xds.type.matcher.v3.IPMatcher.IPRangeMatcher
	2, // 1: xds.type.matcher.v3.IPMatcher.IPRangeMatcher.ranges:type_name -> xds.core.v3.CidrRange
	3, // 2: xds.type.matcher.v3.IPMatcher.IPRangeMatcher.on_match:type_name -> xds.type.matcher.v3.Matcher.OnMatch
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_xds_type_matcher_v3_ip_proto_init() }
func file_xds_type_matcher_v3_ip_proto_init() {
	if File_xds_type_matcher_v3_ip_proto != nil {
		return
	}
	file_xds_type_matcher_v3_matcher_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_xds_type_matcher_v3_ip_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IPMatcher); i {
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
		file_xds_type_matcher_v3_ip_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IPMatcher_IPRangeMatcher); i {
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
			RawDescriptor: file_xds_type_matcher_v3_ip_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_xds_type_matcher_v3_ip_proto_goTypes,
		DependencyIndexes: file_xds_type_matcher_v3_ip_proto_depIdxs,
		MessageInfos:      file_xds_type_matcher_v3_ip_proto_msgTypes,
	}.Build()
	File_xds_type_matcher_v3_ip_proto = out.File
	file_xds_type_matcher_v3_ip_proto_rawDesc = nil
	file_xds_type_matcher_v3_ip_proto_goTypes = nil
	file_xds_type_matcher_v3_ip_proto_depIdxs = nil
}
