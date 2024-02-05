load("@rules_proto//proto:defs.bzl", "proto_library")
load("@rules_buf//buf:defs.bzl", "buf_lint_test")
load("//bazel:api_build_system.bzl", "xds_proto_package")

licenses(["notice"])  # Apache 2

xds_proto_package()

proto_library(
    name = "v3_proto",
    srcs = [
        "migrate.proto",
        "security.proto",
        "sensitive.proto",
        "status.proto",
        "versioning.proto",
    ],
    visibility = ["//visibility:public"],
    deps = ["@com_google_protobuf//:descriptor_proto"],
)

buf_lint_test(
    name = "v3_proto_lint",
    targets = [":v3_proto"],
)
