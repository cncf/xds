load("@rules_proto//proto:defs.bzl", "proto_library")
load("@rules_buf//buf:defs.bzl", "buf_lint_test")
load("//bazel:api_build_system.bzl", "xds_proto_package")

licenses(["notice"])  #Apache 2

xds_proto_package(
    deps = [
        "//xds/annotations/v3:pkg",
        "//xds/core/v3:pkg",
        "//xds/type/v3:pkg",
    ],
)

proto_library(
    name = "v3_proto",
    srcs = [
        "cel.proto",
        "domain.proto",
        "http_inputs.proto",
        "ip.proto",
        "matcher.proto",
        "range.proto",
        "regex.proto",
        "string.proto",
    ],
    visibility = ["//visibility:public"],
    deps = [
        "//xds/annotations/v3:v3_proto",
        "//xds/core/v3:v3_proto",
        "//xds/type/v3:v3_proto",
        "@buf_deps//validate:validate_proto",
    ],
)

buf_lint_test(
    name = "v3_proto_lint",
    targets = [":v3_proto"],
)
