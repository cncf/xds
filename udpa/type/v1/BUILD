load("@rules_proto//proto:defs.bzl", "proto_library")
load("@rules_buf//buf:defs.bzl", "buf_lint_test")
load("//bazel:api_build_system.bzl", "xds_proto_package")

licenses(["notice"])  # Apache 2

xds_proto_package()

proto_library(
    name = "v1_proto",
    srcs = ["typed_struct.proto"],
    visibility = ["//visibility:public"],
    deps = ["@com_google_protobuf//:struct_proto"],
)

buf_lint_test(
    name = "v1_proto_lint",
    targets = [":v1_proto"],
)
