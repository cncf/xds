load("@rules_proto//proto:defs.bzl", "proto_library")
load("@rules_buf//buf:defs.bzl", "buf_lint_test")
load("//bazel:api_build_system.bzl", "xds_proto_package")

licenses(["notice"])  # Apache 2

xds_proto_package()

proto_library(
    name = "v1_proto",
    srcs = ["orca_load_report.proto"],
    visibility = ["//visibility:public"],
    deps = ["@buf_deps//validate:validate_proto"],
)

buf_lint_test(
    name = "v1_proto_lint",
    targets = [":v1_proto"],
)
