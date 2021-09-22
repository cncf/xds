load(":envoy_http_archive.bzl", "udpa_http_archive")
load(":repository_locations.bzl", "REPOSITORY_LOCATIONS")

def udpa_api_dependencies():
    udpa_http_archive(
        "bazel_gazelle",
        locations = REPOSITORY_LOCATIONS,
    )
    udpa_http_archive(
        "com_envoyproxy_protoc_gen_validate",
        locations = REPOSITORY_LOCATIONS,
    )
    udpa_http_archive(
        name = "com_github_grpc_grpc",
        locations = REPOSITORY_LOCATIONS,
    )
    udpa_http_archive(
        name = "com_google_googleapis",
        locations = REPOSITORY_LOCATIONS,
    )
    udpa_http_archive(
        "com_google_protobuf",
        locations = REPOSITORY_LOCATIONS,
    )
    udpa_http_archive(
        "io_bazel_rules_go",
        locations = REPOSITORY_LOCATIONS,
    )
