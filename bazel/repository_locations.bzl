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
        sha256 = "13e7c6460cd979726e5b3b129bb01c34532f115883ac696a75eb7f1d6a9765ed",
        strip_prefix = "grpc-1.40.0",
        urls = ["https://github.com/grpc/grpc/archive/refs/tags/v1.40.0.tar.gz"],
    ),
    com_google_googleapis = dict(
        # TODO(dio): Consider writing a Skylark macro for importing Google API proto.
        sha256 = "9b4e0d0a04a217c06b426aefd03b82581a9510ca766d2d1c70e52bb2ad4a0703",
        strip_prefix = "googleapis-114a745b2841a044e98cdbb19358ed29fcf4a5f1",
        urls = ["https://github.com/googleapis/googleapis/archive/114a745b2841a044e98cdbb19358ed29fcf4a5f1.tar.gz"],
    ),
    com_google_protobuf = dict(
        sha256 = "bb1ddd8172b745cbdc75f06841bd9e7c9de0b3956397723d883423abfab8e176",
        strip_prefix = "protobuf-3.18.0",
        urls = ["https://github.com/protocolbuffers/protobuf/archive/refs/tags/v3.18.0.zip"],
    ),
    dev_cel = dict(
        sha256 = "3de60ea3a29b6246faf04d206b7ee4633c155042e040086205c83edf713aee29",
        strip_prefix = "cel-spec-2657e88023e5e7dfbb62e74de7a7cfd6e8284d7b",
        urls = ["https://github.com/google/cel-spec/archive/2657e88023e5e7dfbb62e74de7a7cfd6e8284d7b.zip"],
    ),
    io_bazel_rules_go = dict(
        sha256 = "6dc2da7ab4cf5d7bfc7c949776b1b7c733f05e56edc4bcd9022bb249d2e2a996",
        urls = [
            "https://mirror.bazel.build/github.com/bazelbuild/rules_go/releases/download/v0.39.1/rules_go-v0.39.1.zip",
            "https://github.com/bazelbuild/rules_go/releases/download/v0.39.1/rules_go-v0.39.1.zip",
        ],
    ),
    # used by dev_cel
    rules_python = dict(
        sha256 = "9d04041ac92a0985e344235f5d946f71ac543f1b1565f2cdbc9a2aaee8adf55b",
        strip_prefix = "rules_python-0.26.0",
        urls = ["https://github.com/bazelbuild/rules_python/releases/download/0.26.0/rules_python-0.26.0.tar.gz"],
    ),
    rules_proto = dict(
        sha256 = "dc3fb206a2cb3441b485eb1e423165b231235a1ea9b031b4433cf7bc1fa460dd",
        strip_prefix = "rules_proto-5.3.0-21.7",
        urls = ["https://github.com/bazelbuild/rules_proto/archive/refs/tags/5.3.0-21.7.tar.gz"],
    ),
)
