# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# NO CHECKED-IN PROTOBUF GENCODE
# source: xds/core/v3/context_params.proto
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
    'xds/core/v3/context_params.proto'
)
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from xds.annotations.v3 import status_pb2 as xds_dot_annotations_dot_v3_dot_status__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n xds/core/v3/context_params.proto\x12\x0bxds.core.v3\x1a\x1fxds/annotations/v3/status.proto\"v\n\rContextParams\x12\x36\n\x06params\x18\x01 \x03(\x0b\x32&.xds.core.v3.ContextParams.ParamsEntry\x1a-\n\x0bParamsEntry\x12\x0b\n\x03key\x18\x01 \x01(\t\x12\r\n\x05value\x18\x02 \x01(\t:\x02\x38\x01\x42Z\n\x16\x63om.github.xds.core.v3B\x12\x43ontextParamsProtoP\x01Z\"github.com/cncf/xds/go/xds/core/v3\xd2\xc6\xa4\xe1\x06\x02\x08\x01\x62\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'xds.core.v3.context_params_pb2', _globals)
if not _descriptor._USE_C_DESCRIPTORS:
  _globals['DESCRIPTOR']._loaded_options = None
  _globals['DESCRIPTOR']._serialized_options = b'\n\026com.github.xds.core.v3B\022ContextParamsProtoP\001Z\"github.com/cncf/xds/go/xds/core/v3\322\306\244\341\006\002\010\001'
  _globals['_CONTEXTPARAMS_PARAMSENTRY']._loaded_options = None
  _globals['_CONTEXTPARAMS_PARAMSENTRY']._serialized_options = b'8\001'
  _globals['_CONTEXTPARAMS']._serialized_start=82
  _globals['_CONTEXTPARAMS']._serialized_end=200
  _globals['_CONTEXTPARAMS_PARAMSENTRY']._serialized_start=155
  _globals['_CONTEXTPARAMS_PARAMSENTRY']._serialized_end=200
# @@protoc_insertion_point(module_scope)
