load("@bazel_gazelle//:def.bzl", "gazelle", "gazelle_binary")

gazelle_binary(
    name = "gazelle-buf",
    languages = [
        # Loads the native proto extension
        "@bazel_gazelle//language/proto:go_default_library",
        # Loads the Buf extension
        "@rules_buf//gazelle/buf:buf",
        # NOTE: This needs to be loaded after the proto language
    ],
)

gazelle(
    name = "gazelle",
    gazelle = ":gazelle-buf",
)

gazelle(
    name = "gazelle-update-repos",
    args = [
        # This can also be `buf.yaml` and `buf.lock`.
        "--from_file=buf.yaml",
        # This is optional but recommended, if absent gazelle
        # will add the rules directly to WORKSPACE
        "-to_macro=buf_deps.bzl%buf_deps",
        # Deletes outdated repo rules
        "-prune",
    ],
    command = "update-repos",
    gazelle = ":gazelle-buf",
)

exports_files(["buf.yaml"])
