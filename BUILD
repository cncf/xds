load("@rules_buf//buf:defs.bzl", "buf_format", "buf_lint_test")

# Buf configuration files
exports_files([
    "buf.yaml",
    "buf.lock",
])

# Buf lint test for the entire repository
buf_lint_test(
    name = "buf_lint_test",
    config = ":buf.yaml",
    targets = [
        "//udpa/annotations:pkg",
        "//udpa/data/orca/v1:pkg",
        "//udpa/service/orca/v1:pkg",
        "//udpa/type/v1:pkg",
        "//xds/annotations/v3:pkg",
        "//xds/core/v3:pkg",
        "//xds/data/orca/v3:pkg",
        "//xds/service/orca/v3:pkg",
        "//xds/type/v3:pkg",
        "//xds/type/matcher/v3:pkg",
    ],
)

# Buf format rule - run with: bazel run //:buf_format
# This formats all proto files in the workspace
buf_format(
    name = "buf_format",
)
