load(":envoy_http_archive.bzl", "xds_http_archive")
load(":repository_locations.bzl", "REPOSITORY_LOCATIONS")

def xds_api_dependencies():
    xds_http_archive(
        "bazel_gazelle",
        locations = REPOSITORY_LOCATIONS,
    )
    xds_http_archive(
        "com_github_bufbuild_buf",
        locations = REPOSITORY_LOCATIONS,
        build_file_content = BUF_BUILD_CONTENT,
        tags = ["manual"],
    )
    xds_http_archive(
        "com_envoyproxy_protoc_gen_validate",
        locations = REPOSITORY_LOCATIONS,
    )
    xds_http_archive(
        name = "com_github_grpc_grpc",
        locations = REPOSITORY_LOCATIONS,
    )
    xds_http_archive(
        name = "com_google_googleapis",
        locations = REPOSITORY_LOCATIONS,
    )
    xds_http_archive(
        "com_google_protobuf",
        locations = REPOSITORY_LOCATIONS,
    )
    xds_http_archive(
        "io_bazel_rules_go",
        locations = REPOSITORY_LOCATIONS,
    )

# Old name for backward compatibility.
# TODO(roth): Remove once all callers are updated to use the new name.
def udpa_api_dependencies():
  xds_api_dependencies()

BUF_BUILD_CONTENT = """
package(
    default_visibility = ["//visibility:public"],
)

filegroup(
    name = "buf",
    srcs = [
        "@com_github_bufbuild_buf//:bin/buf",
    ],
)
"""
