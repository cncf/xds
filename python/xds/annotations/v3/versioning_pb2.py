# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# NO CHECKED-IN PROTOBUF GENCODE
# source: xds/annotations/v3/versioning.proto
# Protobuf Python Version: 5.28.0
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import runtime_version as _runtime_version
from google.protobuf import symbol_database as _symbol_database
from google.protobuf.internal import builder as _builder
_runtime_version.ValidateProtobufRuntimeVersion(
    _runtime_version.Domain.PUBLIC,
    5,
    28,
    0,
    '',
    'xds/annotations/v3/versioning.proto'
)
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from google.protobuf import descriptor_pb2 as google_dot_protobuf_dot_descriptor__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n#xds/annotations/v3/versioning.proto\x12\x12xds.annotations.v3\x1a google/protobuf/descriptor.proto\"5\n\x14VersioningAnnotation\x12\x1d\n\x15previous_message_type\x18\x01 \x01(\t:`\n\nversioning\x12\x1f.google.protobuf.MessageOptions\x18\x93\xfd\x86, \x01(\x0b\x32(.xds.annotations.v3.VersioningAnnotationB+Z)github.com/cncf/xds/go/xds/annotations/v3b\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'xds.annotations.v3.versioning_pb2', _globals)
if not _descriptor._USE_C_DESCRIPTORS:
  _globals['DESCRIPTOR']._loaded_options = None
  _globals['DESCRIPTOR']._serialized_options = b'Z)github.com/cncf/xds/go/xds/annotations/v3'
  _globals['_VERSIONINGANNOTATION']._serialized_start=93
  _globals['_VERSIONINGANNOTATION']._serialized_end=146
# @@protoc_insertion_point(module_scope)
