REPOSITORY_LOCATIONS = dict(
    bazel_gazelle = dict(
        sha256 = "b8b6d75de6e4bf7c41b7737b183523085f56283f6db929b86c5e7e1f09cf59c9",
        urls = ["https://github.com/bazelbuild/bazel-gazelle/releases/download/v0.31.1/bazel-gazelle-v0.31.1.tar.gz"],
    ),
    bazel_skylib = dict(
        sha256 = "1c531376ac7e5a180e0237938a2536de0c54d93f5c278634818e0efc952dd56c",
        urls = ["https://mirror.bazel.build/github.com/bazelbuild/bazel-skylib/releases/download/1.0.3/bazel-skylib-1.0.3.tar.gz"],
    ),
    com_envoyproxy_protoc_gen_validate = dict(
        sha256 = "c695fc5a2e5a1b52904cd8a58ce7a1c3a80f7f50719496fd606e551685c01101",
        strip_prefix = "protoc-gen-validate-0.6.1",
        urls = ["https://github.com/envoyproxy/protoc-gen-validate/archive/refs/tags/v0.6.1.tar.gz"],
    ),
    com_github_grpc_grpc = dict(
        sha256 = "916f88a34f06b56432611aaa8c55befee96d0a7b7d7457733b9deeacbc016f99",
        strip_prefix = "grpc-1.59.1",
        urls = ["https://github.com/grpc/grpc/archive/refs/tags/v1.59.1.tar.gz"],
    ),
    com_google_googleapis = dict(
        # TODO(dio): Consider writing a Skylark macro for importing Google API proto.
        sha256 = "9b4e0d0a04a217c06b426aefd03b82581a9510ca766d2d1c70e52bb2ad4a0703",
        strip_prefix = "googleapis-114a745b2841a044e98cdbb19358ed29fcf4a5f1",
        urls = ["https://github.com/googleapis/googleapis/archive/114a745b2841a044e98cdbb19358ed29fcf4a5f1.tar.gz"],
    ),
    com_google_protobuf = dict(
        sha256 = "8242327e5df8c80ba49e4165250b8f79a76bd11765facefaaecfca7747dc8da2",
        strip_prefix = "protobuf-3.21.5",
        urls = ["https://github.com/protocolbuffers/protobuf/archive/v3.21.5.zip"],
    ),
    io_bazel_rules_go = dict(
        sha256 = "6dc2da7ab4cf5d7bfc7c949776b1b7c733f05e56edc4bcd9022bb249d2e2a996",
        urls = [
            "https://mirror.bazel.build/github.com/bazelbuild/rules_go/releases/download/v0.39.1/rules_go-v0.39.1.zip",
            "https://github.com/bazelbuild/rules_go/releases/download/v0.39.1/rules_go-v0.39.1.zip",
        ],
    ),
    rules_proto = dict(
        sha256 = "80d3a4ec17354cccc898bfe32118edd934f851b03029d63ef3fc7c8663a7415c",
        strip_prefix = "rules_proto-5.3.0-21.5",
        urls = ["https://github.com/bazelbuild/rules_proto/archive/refs/tags/5.3.0-21.5.tar.gz",],
    ),
)
