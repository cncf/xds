load("@rules_proto//proto:defs.bzl", "proto_library")
load("@rules_buf//buf:defs.bzl", "buf_lint_test")
load("//bazel:api_build_system.bzl", "xds_proto_package")

licenses(["notice"])  # Apache 2

xds_proto_package(
    deps = [
        "//xds/annotations/v3:pkg",
        "@com_google_googleapis//google/api/expr/v1alpha1:checked_proto",
        "@com_google_googleapis//google/api/expr/v1alpha1:syntax_proto",
    ],
)

proto_library(
    name = "v3_proto",
    srcs = [
        "cel.proto",
        "range.proto",
        "typed_struct.proto",
    ],
    visibility = ["//visibility:public"],
    deps = [
        "//xds/annotations/v3:v3_proto",
        "@buf_deps//validate:validate_proto",
        "@com_google_protobuf//:struct_proto",
        "@com_google_protobuf//:wrappers_proto",
        "@go_googleapis//google/api/expr/v1alpha1:expr_proto",
    ],
)

buf_lint_test(
    name = "v3_proto_lint",
    targets = [":v3_proto"],
)
