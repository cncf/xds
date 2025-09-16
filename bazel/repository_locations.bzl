REPOSITORY_LOCATIONS = dict(
    bazel_gazelle = dict(
        sha256 = "e467b801046b6598c657309b45d2426dc03513777bd1092af2c62eebf990aca5",
        urls = [
            "https://mirror.bazel.build/github.com/bazelbuild/bazel-gazelle/releases/download/v0.45.0/bazel-gazelle-v0.45.0.tar.gz",
            "https://github.com/bazelbuild/bazel-gazelle/releases/download/v0.45.0/bazel-gazelle-v0.45.0.tar.gz",
        ],
    ),
    bazel_features = dict(
        sha256 = "c41853e3b636c533b86bf5ab4658064e6cc9db0a3bce52cbff0629e094344ca9",
        strip_prefix = "bazel_features-1.33.0",
        urls = ["https://github.com/bazel-contrib/bazel_features/releases/download/v1.33.0/bazel_features-v1.33.0.tar.gz"],
    ),
    bazel_skylib = dict(
        sha256 = "1c531376ac7e5a180e0237938a2536de0c54d93f5c278634818e0efc952dd56c",
        urls = ["https://mirror.bazel.build/github.com/bazelbuild/bazel-skylib/releases/download/1.0.3/bazel-skylib-1.0.3.tar.gz"],
    ),
    com_envoyproxy_protoc_gen_validate = dict(
        sha256 = "92e29c2150675ce954c965bcaa559ca944704b75711533cfe03ce541dcf5a1dd",
        strip_prefix = "protoc-gen-validate-1.0.4",
        urls = ["https://github.com/envoyproxy/protoc-gen-validate/archive/refs/tags/v1.0.4.tar.gz"],
    ),
    com_github_grpc_grpc = dict(
        sha256 = "4a8aa99d5e24f80ea6b7ec95463e16af5bd91aa805e26c661ef6491ae3d2d23c",
        strip_prefix = "grpc-1.72.0",
        urls = ["https://github.com/grpc/grpc/archive/refs/tags/v1.72.0.tar.gz"],
    ),
    com_google_googleapis = dict(
        # TODO(dio): Consider writing a Starlark macro for importing Google API proto.
        sha256 = "9d1a930e767c93c825398b8f8692eca3fe353b9aaadedfbcf1fca2282c85df88",
        strip_prefix = "googleapis-64926d52febbf298cb82a8f472ade4a3969ba922",
        urls = [
            "https://github.com/googleapis/googleapis/archive/64926d52febbf298cb82a8f472ade4a3969ba922.zip",
        ],
    ),
    com_google_protobuf = dict(
        sha256 = "85803e01f347141e16a2f770213a496f808fff9f0138c7c0e0c9dfa708b0da92",
        strip_prefix = "protobuf-29.3",
        urls = ["https://github.com/protocolbuffers/protobuf/archive/v29.3.zip"],
    ),
    dev_cel = dict(
        sha256 = "5cba6b0029e727d1f4d8fd134de4e747cecc0bc293d026017d7edc48058d09f7",
        strip_prefix = "cel-spec-0.24.0",
        urls = ["https://github.com/google/cel-spec/archive/refs/tags/v0.24.0.tar.gz"],
    ),
    io_bazel_rules_go = dict(
        sha256 = "b78f77458e77162f45b4564d6b20b6f92f56431ed59eaaab09e7819d1d850313",
        urls = [
            "https://mirror.bazel.build/github.com/bazelbuild/rules_go/releases/download/v0.53.0/rules_go-v0.53.0.zip",
            "https://github.com/bazelbuild/rules_go/releases/download/v0.53.0/rules_go-v0.53.0.zip",
        ],
    ),
    rules_proto = dict(
        sha256 = "80d3a4ec17354cccc898bfe32118edd934f851b03029d63ef3fc7c8663a7415c",
        strip_prefix = "rules_proto-5.3.0-21.5",
        urls = ["https://github.com/bazelbuild/rules_proto/archive/refs/tags/5.3.0-21.5.tar.gz"],
    ),
)
