load("@rules_proto//proto:defs.bzl", "proto_library")
load("@rules_buf//buf:defs.bzl", "buf_lint_test")
load("//bazel:api_build_system.bzl", "xds_proto_package")

licenses(["notice"])  # Apache 2

xds_proto_package(deps = ["//xds/annotations/v3:pkg"])

proto_library(
    name = "v3_proto",
    srcs = [
        "authority.proto",
        "cidr.proto",
        "collection_entry.proto",
        "context_params.proto",
        "extension.proto",
        "resource.proto",
        "resource_locator.proto",
        "resource_name.proto",
    ],
    visibility = ["//visibility:public"],
    deps = [
        "//xds/annotations/v3:v3_proto",
        "@buf_deps//validate:validate_proto",
        "@com_google_protobuf//:any_proto",
        "@com_google_protobuf//:wrappers_proto",
    ],
)

buf_lint_test(
    name = "v3_proto_lint",
    targets = [":v3_proto"],
)
