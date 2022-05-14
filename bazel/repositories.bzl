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


def _go_deps():
    external_http_archive(
        name = "io_bazel_rules_go",
        # TODO(wrowe, sunjayBhatia): remove when Windows RBE supports batch file invocation
        patch_args = ["-p1"],
        patches = ["//bazel:rules_go.patch"],
    )
    external_http_archive("bazel_gazelle")

def xds_api_dependencies():
    _go_deps()
    external_http_archive(
        name = "bazel_skylib",
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
        patches = ["//bazel:protobuf.patch"],
        patch_args = ["-p1"],
    )
