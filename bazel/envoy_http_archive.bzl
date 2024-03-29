load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")

def xds_http_archive(name, locations, **kwargs):
    # `existing_rule_keys` contains the names of repositories that have already
    # been defined in the Bazel workspace. By skipping repos with existing keys,
    # users can override dependency versions by using standard Bazel repository
    # rules in their WORKSPACE files.
    existing_rule_keys = native.existing_rules().keys()
    if name in existing_rule_keys:
        # This repository has already been defined, probably because the user
        # wants to override the version. Do nothing.
        return

    loc_key = kwargs.pop("repository_key", name)
    location = locations[loc_key]

    # HTTP tarball at a given URL. Add a BUILD file if requested.
    http_archive(
        name = name,
        urls = location["urls"],
        sha256 = location["sha256"],
        strip_prefix = location.get("strip_prefix", ""),
        **kwargs
    )

# Old name for backward compatibility.
# TODO(roth): Remove once all callers are changed to use the new name.
def udpa_http_archive(name, locations, **kwargs):
    xds_http_archive(name, locations, **kwargs)
