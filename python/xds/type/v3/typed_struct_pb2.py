# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# NO CHECKED-IN PROTOBUF GENCODE
# source: xds/type/v3/typed_struct.proto
# Protobuf Python Version: 5.29.1
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import runtime_version as _runtime_version
from google.protobuf import symbol_database as _symbol_database
from google.protobuf.internal import builder as _builder
_runtime_version.ValidateProtobufRuntimeVersion(
    _runtime_version.Domain.PUBLIC,
    5,
    29,
    1,
    '',
    'xds/type/v3/typed_struct.proto'
)
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from google.protobuf import struct_pb2 as google_dot_protobuf_dot_struct__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\x1exds/type/v3/typed_struct.proto\x12\x0bxds.type.v3\x1a\x1cgoogle/protobuf/struct.proto\"G\n\x0bTypedStruct\x12\x10\n\x08type_url\x18\x01 \x01(\t\x12&\n\x05value\x18\x02 \x01(\x0b\x32\x17.google.protobuf.StructBP\n\x16\x63om.github.xds.type.v3B\x10TypedStructProtoP\x01Z\"github.com/cncf/xds/go/xds/type/v3b\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'xds.type.v3.typed_struct_pb2', _globals)
if not _descriptor._USE_C_DESCRIPTORS:
  _globals['DESCRIPTOR']._loaded_options = None
  _globals['DESCRIPTOR']._serialized_options = b'\n\026com.github.xds.type.v3B\020TypedStructProtoP\001Z\"github.com/cncf/xds/go/xds/type/v3'
  _globals['_TYPEDSTRUCT']._serialized_start=77
  _globals['_TYPEDSTRUCT']._serialized_end=148
# @@protoc_insertion_point(module_scope)
