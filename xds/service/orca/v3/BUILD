load("@rules_proto//proto:defs.bzl", "proto_library")
load("@rules_buf//buf:defs.bzl", "buf_lint_test")
load("//bazel:api_build_system.bzl", "xds_proto_package")

licenses(["notice"])  # Apache 2

xds_proto_package(
    has_services = True,
    deps = [
        "//xds/data/orca/v3:pkg",
    ],
)

proto_library(
    name = "v3_proto",
    srcs = ["orca.proto"],
    visibility = ["//visibility:public"],
    deps = [
        "//xds/data/orca/v3:v3_proto",
        "@com_google_protobuf//:duration_proto",
    ],
)

buf_lint_test(
    name = "v3_proto_lint",
    targets = [":v3_proto"],
)
