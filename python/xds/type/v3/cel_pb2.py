# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# NO CHECKED-IN PROTOBUF GENCODE
# source: xds/type/v3/cel.proto
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
    'xds/type/v3/cel.proto'
)
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from google.api.expr.v1alpha1 import checked_pb2 as google_dot_api_dot_expr_dot_v1alpha1_dot_checked__pb2
from google.api.expr.v1alpha1 import syntax_pb2 as google_dot_api_dot_expr_dot_v1alpha1_dot_syntax__pb2
from cel.expr import checked_pb2 as cel_dot_expr_dot_checked__pb2
from cel.expr import syntax_pb2 as cel_dot_expr_dot_syntax__pb2
from google.protobuf import wrappers_pb2 as google_dot_protobuf_dot_wrappers__pb2
from xds.annotations.v3 import status_pb2 as xds_dot_annotations_dot_v3_dot_status__pb2
from validate import validate_pb2 as validate_dot_validate__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\x15xds/type/v3/cel.proto\x12\x0bxds.type.v3\x1a&google/api/expr/v1alpha1/checked.proto\x1a%google/api/expr/v1alpha1/syntax.proto\x1a\x16\x63\x65l/expr/checked.proto\x1a\x15\x63\x65l/expr/syntax.proto\x1a\x1egoogle/protobuf/wrappers.proto\x1a\x1fxds/annotations/v3/status.proto\x1a\x17validate/validate.proto\"\x9e\x02\n\rCelExpression\x12?\n\x0bparsed_expr\x18\x01 \x01(\x0b\x32$.google.api.expr.v1alpha1.ParsedExprB\x02\x18\x01H\x00\x12\x41\n\x0c\x63hecked_expr\x18\x02 \x01(\x0b\x32%.google.api.expr.v1alpha1.CheckedExprB\x02\x18\x01H\x00\x12-\n\x0f\x63\x65l_expr_parsed\x18\x03 \x01(\x0b\x32\x14.cel.expr.ParsedExpr\x12/\n\x10\x63\x65l_expr_checked\x18\x04 \x01(\x0b\x32\x15.cel.expr.CheckedExpr\x12\x17\n\x0f\x63\x65l_expr_string\x18\x05 \x01(\tB\x10\n\x0e\x65xpr_specifier\"\x83\x01\n\x10\x43\x65lExtractString\x12:\n\x0c\x65xpr_extract\x18\x01 \x01(\x0b\x32\x1a.xds.type.v3.CelExpressionB\x08\xfa\x42\x05\x8a\x01\x02\x10\x01\x12\x33\n\rdefault_value\x18\x02 \x01(\x0b\x32\x1c.google.protobuf.StringValueBP\n\x16\x63om.github.xds.type.v3B\x08\x43\x65lProtoP\x01Z\"github.com/cncf/xds/go/xds/type/v3\xd2\xc6\xa4\xe1\x06\x02\x08\x01\x62\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'xds.type.v3.cel_pb2', _globals)
if not _descriptor._USE_C_DESCRIPTORS:
  _globals['DESCRIPTOR']._loaded_options = None
  _globals['DESCRIPTOR']._serialized_options = b'\n\026com.github.xds.type.v3B\010CelProtoP\001Z\"github.com/cncf/xds/go/xds/type/v3\322\306\244\341\006\002\010\001'
  _globals['_CELEXPRESSION'].fields_by_name['parsed_expr']._loaded_options = None
  _globals['_CELEXPRESSION'].fields_by_name['parsed_expr']._serialized_options = b'\030\001'
  _globals['_CELEXPRESSION'].fields_by_name['checked_expr']._loaded_options = None
  _globals['_CELEXPRESSION'].fields_by_name['checked_expr']._serialized_options = b'\030\001'
  _globals['_CELEXTRACTSTRING'].fields_by_name['expr_extract']._loaded_options = None
  _globals['_CELEXTRACTSTRING'].fields_by_name['expr_extract']._serialized_options = b'\372B\005\212\001\002\020\001'
  _globals['_CELEXPRESSION']._serialized_start=255
  _globals['_CELEXPRESSION']._serialized_end=541
  _globals['_CELEXTRACTSTRING']._serialized_start=544
  _globals['_CELEXTRACTSTRING']._serialized_end=675
# @@protoc_insertion_point(module_scope)
