load("//bazel:api_build_system.bzl", "xds_cc_test", "xds_go_test")

licenses(["notice"])  # Apache 2

xds_cc_test(
    name = "build_test",
    srcs = ["build_test.cc"],
    deps = [
        "//xds/service/orca/v3:pkg_cc_proto",
    ],
)

xds_go_test(
    name = "go_build_test",
    srcs = ["go_build_test.go"],
    importpath = "go_build_test",
    deps = [
        "//xds/data/orca/v3:pkg_go_proto",
        "//xds/service/orca/v3:pkg_go_proto",
        "//xds/type/v3:pkg_go_proto",
    ],
)
