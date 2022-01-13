// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.18.0
// source: xds/type/matcher/v3/range.proto

package v3

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type Int64Range struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Start int64 `protobuf:"varint,1,opt,name=start,proto3" json:"start,omitempty"`
	End   int64 `protobuf:"varint,2,opt,name=end,proto3" json:"end,omitempty"`
}

func (x *Int64Range) Reset() {
	*x = Int64Range{}
	if protoimpl.UnsafeEnabled {
		mi := &file_xds_type_matcher_v3_range_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Int64Range) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Int64Range) ProtoMessage() {}

func (x *Int64Range) ProtoReflect() protoreflect.Message {
	mi := &file_xds_type_matcher_v3_range_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Int64Range.ProtoReflect.Descriptor instead.
func (*Int64Range) Descriptor() ([]byte, []int) {
	return file_xds_type_matcher_v3_range_proto_rawDescGZIP(), []int{0}
}

func (x *Int64Range) GetStart() int64 {
	if x != nil {
		return x.Start
	}
	return 0
}

func (x *Int64Range) GetEnd() int64 {
	if x != nil {
		return x.End
	}
	return 0
}

type Int64RangeMatcher struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ranges []*Int64Range `protobuf:"bytes,1,rep,name=ranges,proto3" json:"ranges,omitempty"`
}

func (x *Int64RangeMatcher) Reset() {
	*x = Int64RangeMatcher{}
	if protoimpl.UnsafeEnabled {
		mi := &file_xds_type_matcher_v3_range_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Int64RangeMatcher) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Int64RangeMatcher) ProtoMessage() {}

func (x *Int64RangeMatcher) ProtoReflect() protoreflect.Message {
	mi := &file_xds_type_matcher_v3_range_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Int64RangeMatcher.ProtoReflect.Descriptor instead.
func (*Int64RangeMatcher) Descriptor() ([]byte, []int) {
	return file_xds_type_matcher_v3_range_proto_rawDescGZIP(), []int{1}
}

func (x *Int64RangeMatcher) GetRanges() []*Int64Range {
	if x != nil {
		return x.Ranges
	}
	return nil
}

type Int32Range struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Start int32 `protobuf:"varint,1,opt,name=start,proto3" json:"start,omitempty"`
	End   int32 `protobuf:"varint,2,opt,name=end,proto3" json:"end,omitempty"`
}

func (x *Int32Range) Reset() {
	*x = Int32Range{}
	if protoimpl.UnsafeEnabled {
		mi := &file_xds_type_matcher_v3_range_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Int32Range) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Int32Range) ProtoMessage() {}

func (x *Int32Range) ProtoReflect() protoreflect.Message {
	mi := &file_xds_type_matcher_v3_range_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Int32Range.ProtoReflect.Descriptor instead.
func (*Int32Range) Descriptor() ([]byte, []int) {
	return file_xds_type_matcher_v3_range_proto_rawDescGZIP(), []int{2}
}

func (x *Int32Range) GetStart() int32 {
	if x != nil {
		return x.Start
	}
	return 0
}

func (x *Int32Range) GetEnd() int32 {
	if x != nil {
		return x.End
	}
	return 0
}

type Int32RangeMatcher struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ranges []*Int32Range `protobuf:"bytes,1,rep,name=ranges,proto3" json:"ranges,omitempty"`
}

func (x *Int32RangeMatcher) Reset() {
	*x = Int32RangeMatcher{}
	if protoimpl.UnsafeEnabled {
		mi := &file_xds_type_matcher_v3_range_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Int32RangeMatcher) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Int32RangeMatcher) ProtoMessage() {}

func (x *Int32RangeMatcher) ProtoReflect() protoreflect.Message {
	mi := &file_xds_type_matcher_v3_range_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Int32RangeMatcher.ProtoReflect.Descriptor instead.
func (*Int32RangeMatcher) Descriptor() ([]byte, []int) {
	return file_xds_type_matcher_v3_range_proto_rawDescGZIP(), []int{3}
}

func (x *Int32RangeMatcher) GetRanges() []*Int32Range {
	if x != nil {
		return x.Ranges
	}
	return nil
}

type DoubleRange struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Start float64 `protobuf:"fixed64,1,opt,name=start,proto3" json:"start,omitempty"`
	End   float64 `protobuf:"fixed64,2,opt,name=end,proto3" json:"end,omitempty"`
}

func (x *DoubleRange) Reset() {
	*x = DoubleRange{}
	if protoimpl.UnsafeEnabled {
		mi := &file_xds_type_matcher_v3_range_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DoubleRange) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DoubleRange) ProtoMessage() {}

func (x *DoubleRange) ProtoReflect() protoreflect.Message {
	mi := &file_xds_type_matcher_v3_range_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DoubleRange.ProtoReflect.Descriptor instead.
func (*DoubleRange) Descriptor() ([]byte, []int) {
	return file_xds_type_matcher_v3_range_proto_rawDescGZIP(), []int{4}
}

func (x *DoubleRange) GetStart() float64 {
	if x != nil {
		return x.Start
	}
	return 0
}

func (x *DoubleRange) GetEnd() float64 {
	if x != nil {
		return x.End
	}
	return 0
}

type DoubleRangeMatcher struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ranges []*DoubleRange `protobuf:"bytes,1,rep,name=ranges,proto3" json:"ranges,omitempty"`
}

func (x *DoubleRangeMatcher) Reset() {
	*x = DoubleRangeMatcher{}
	if protoimpl.UnsafeEnabled {
		mi := &file_xds_type_matcher_v3_range_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DoubleRangeMatcher) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DoubleRangeMatcher) ProtoMessage() {}

func (x *DoubleRangeMatcher) ProtoReflect() protoreflect.Message {
	mi := &file_xds_type_matcher_v3_range_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DoubleRangeMatcher.ProtoReflect.Descriptor instead.
func (*DoubleRangeMatcher) Descriptor() ([]byte, []int) {
	return file_xds_type_matcher_v3_range_proto_rawDescGZIP(), []int{5}
}

func (x *DoubleRangeMatcher) GetRanges() []*DoubleRange {
	if x != nil {
		return x.Ranges
	}
	return nil
}

var File_xds_type_matcher_v3_range_proto protoreflect.FileDescriptor

var file_xds_type_matcher_v3_range_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x78, 0x64, 0x73, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x2f, 0x6d, 0x61, 0x74, 0x63, 0x68,
	0x65, 0x72, 0x2f, 0x76, 0x33, 0x2f, 0x72, 0x61, 0x6e, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x13, 0x78, 0x64, 0x73, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x6d, 0x61, 0x74, 0x63,
	0x68, 0x65, 0x72, 0x2e, 0x76, 0x33, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65,
	0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x34, 0x0a, 0x0a, 0x49, 0x6e, 0x74, 0x36, 0x34, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x12, 0x14, 0x0a,
	0x05, 0x73, 0x74, 0x61, 0x72, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x73, 0x74,
	0x61, 0x72, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x65, 0x6e, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x03, 0x65, 0x6e, 0x64, 0x22, 0x56, 0x0a, 0x11, 0x49, 0x6e, 0x74, 0x36, 0x34, 0x52, 0x61,
	0x6e, 0x67, 0x65, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x12, 0x41, 0x0a, 0x06, 0x72, 0x61,
	0x6e, 0x67, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x78, 0x64, 0x73,
	0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x2e, 0x76, 0x33,
	0x2e, 0x49, 0x6e, 0x74, 0x36, 0x34, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x42, 0x08, 0xfa, 0x42, 0x05,
	0x92, 0x01, 0x02, 0x08, 0x01, 0x52, 0x06, 0x72, 0x61, 0x6e, 0x67, 0x65, 0x73, 0x22, 0x34, 0x0a,
	0x0a, 0x49, 0x6e, 0x74, 0x33, 0x32, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x73,
	0x74, 0x61, 0x72, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x73, 0x74, 0x61, 0x72,
	0x74, 0x12, 0x10, 0x0a, 0x03, 0x65, 0x6e, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03,
	0x65, 0x6e, 0x64, 0x22, 0x56, 0x0a, 0x11, 0x49, 0x6e, 0x74, 0x33, 0x32, 0x52, 0x61, 0x6e, 0x67,
	0x65, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x12, 0x41, 0x0a, 0x06, 0x72, 0x61, 0x6e, 0x67,
	0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x78, 0x64, 0x73, 0x2e, 0x74,
	0x79, 0x70, 0x65, 0x2e, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x2e, 0x76, 0x33, 0x2e, 0x49,
	0x6e, 0x74, 0x33, 0x32, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x92, 0x01,
	0x02, 0x08, 0x01, 0x52, 0x06, 0x72, 0x61, 0x6e, 0x67, 0x65, 0x73, 0x22, 0x35, 0x0a, 0x0b, 0x44,
	0x6f, 0x75, 0x62, 0x6c, 0x65, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x74,
	0x61, 0x72, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x01, 0x52, 0x05, 0x73, 0x74, 0x61, 0x72, 0x74,
	0x12, 0x10, 0x0a, 0x03, 0x65, 0x6e, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x03, 0x65,
	0x6e, 0x64, 0x22, 0x58, 0x0a, 0x12, 0x44, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x52, 0x61, 0x6e, 0x67,
	0x65, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x12, 0x42, 0x0a, 0x06, 0x72, 0x61, 0x6e, 0x67,
	0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x78, 0x64, 0x73, 0x2e, 0x74,
	0x79, 0x70, 0x65, 0x2e, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x2e, 0x76, 0x33, 0x2e, 0x44,
	0x6f, 0x75, 0x62, 0x6c, 0x65, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x92,
	0x01, 0x02, 0x08, 0x01, 0x52, 0x06, 0x72, 0x61, 0x6e, 0x67, 0x65, 0x73, 0x42, 0x5a, 0x0a, 0x1e,
	0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x78, 0x64, 0x73, 0x2e, 0x74,
	0x79, 0x70, 0x65, 0x2e, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x2e, 0x76, 0x33, 0x42, 0x0a,
	0x52, 0x61, 0x6e, 0x67, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x2a, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6e, 0x63, 0x66, 0x2f, 0x78, 0x64,
	0x73, 0x2f, 0x67, 0x6f, 0x2f, 0x78, 0x64, 0x73, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x2f, 0x6d, 0x61,
	0x74, 0x63, 0x68, 0x65, 0x72, 0x2f, 0x76, 0x33, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_xds_type_matcher_v3_range_proto_rawDescOnce sync.Once
	file_xds_type_matcher_v3_range_proto_rawDescData = file_xds_type_matcher_v3_range_proto_rawDesc
)

func file_xds_type_matcher_v3_range_proto_rawDescGZIP() []byte {
	file_xds_type_matcher_v3_range_proto_rawDescOnce.Do(func() {
		file_xds_type_matcher_v3_range_proto_rawDescData = protoimpl.X.CompressGZIP(file_xds_type_matcher_v3_range_proto_rawDescData)
	})
	return file_xds_type_matcher_v3_range_proto_rawDescData
}

var file_xds_type_matcher_v3_range_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_xds_type_matcher_v3_range_proto_goTypes = []interface{}{
	(*Int64Range)(nil),         // 0: xds.type.matcher.v3.Int64Range
	(*Int64RangeMatcher)(nil),  // 1: xds.type.matcher.v3.Int64RangeMatcher
	(*Int32Range)(nil),         // 2: xds.type.matcher.v3.Int32Range
	(*Int32RangeMatcher)(nil),  // 3: xds.type.matcher.v3.Int32RangeMatcher
	(*DoubleRange)(nil),        // 4: xds.type.matcher.v3.DoubleRange
	(*DoubleRangeMatcher)(nil), // 5: xds.type.matcher.v3.DoubleRangeMatcher
}
var file_xds_type_matcher_v3_range_proto_depIdxs = []int32{
	0, // 0: xds.type.matcher.v3.Int64RangeMatcher.ranges:type_name -> xds.type.matcher.v3.Int64Range
	2, // 1: xds.type.matcher.v3.Int32RangeMatcher.ranges:type_name -> xds.type.matcher.v3.Int32Range
	4, // 2: xds.type.matcher.v3.DoubleRangeMatcher.ranges:type_name -> xds.type.matcher.v3.DoubleRange
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_xds_type_matcher_v3_range_proto_init() }
func file_xds_type_matcher_v3_range_proto_init() {
	if File_xds_type_matcher_v3_range_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_xds_type_matcher_v3_range_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Int64Range); i {
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
		file_xds_type_matcher_v3_range_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Int64RangeMatcher); i {
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
		file_xds_type_matcher_v3_range_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Int32Range); i {
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
		file_xds_type_matcher_v3_range_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Int32RangeMatcher); i {
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
		file_xds_type_matcher_v3_range_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DoubleRange); i {
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
		file_xds_type_matcher_v3_range_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DoubleRangeMatcher); i {
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
			RawDescriptor: file_xds_type_matcher_v3_range_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_xds_type_matcher_v3_range_proto_goTypes,
		DependencyIndexes: file_xds_type_matcher_v3_range_proto_depIdxs,
		MessageInfos:      file_xds_type_matcher_v3_range_proto_msgTypes,
	}.Build()
	File_xds_type_matcher_v3_range_proto = out.File
	file_xds_type_matcher_v3_range_proto_rawDesc = nil
	file_xds_type_matcher_v3_range_proto_goTypes = nil
	file_xds_type_matcher_v3_range_proto_depIdxs = nil
}
