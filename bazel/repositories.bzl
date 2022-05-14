load(":envoy_http_archive.bzl", "xds_http_archive")
load(":external_deps.bzl", "load_repository_locations")
load(":repository_locations.bzl", "REPOSITORY_LOCATIONS_SPEC")

REPOSITORY_LOCATIONS = load_repository_locations(REPOSITORY_LOCATIONS_SPEC)

# Use this macro to reference any HTTP archive from bazel/repository_locations.bzl.
def external_http_archive(name, **kwargs):
    xds_http_archive(
        name,
        locations = REPOSITORY_LOCATIONS,
        **kwargs
    )


def xds_api_dependencies():
    external_http_archive(
        name = "bazel_gazelle",
    )
    external_http_archive(
        name = "com_envoyproxy_protoc_gen_validate",
    )
    external_http_archive(
        name = "com_github_grpc_grpc",
    )
    external_http_archive(
        name = "com_google_googleapis",
    )
    external_http_archive(
        name = "com_google_protobuf",
    )
    external_http_archive(
        name = "io_bazel_rules_go",
    )
