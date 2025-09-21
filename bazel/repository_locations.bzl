REPOSITORY_LOCATIONS = dict(
    bazel_gazelle = dict(
        sha256 = "b7387f72efb59f876e4daae42f1d3912d0d45563eac7cb23d1de0b094ab588cf",
        urls = [
            "https://mirror.bazel.build/github.com/bazelbuild/bazel-gazelle/releases/download/v0.34.0/bazel-gazelle-v0.34.0.tar.gz",
            "https://github.com/bazelbuild/bazel-gazelle/releases/download/v0.34.0/bazel-gazelle-v0.34.0.tar.gz",
        ],
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
        sha256 = "916f88a34f06b56432611aaa8c55befee96d0a7b7d7457733b9deeacbc016f99",
        strip_prefix = "grpc-1.59.1",
        urls = ["https://github.com/grpc/grpc/archive/refs/tags/v1.59.1.tar.gz"],
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
        sha256 = "3d32940e975c4ad9b8ba69640e78f5527075bae33ca2890275bf26b853c0962c",
        strip_prefix = "protobuf-29.1",
        urls = ["https://github.com/protocolbuffers/protobuf/archive/v29.1.tar.gz"],
    ),
    dev_cel = dict(
        sha256 = "3ee09eb69dbe77722e9dee23dc48dc2cd9f765869fcf5ffb1226587c81791a0b",
        strip_prefix = "cel-spec-0.15.0",
        urls = ["https://github.com/google/cel-spec/archive/refs/tags/v0.15.0.tar.gz"],
    ),
    io_bazel_rules_go = dict(
        sha256 = "80a98277ad1311dacd837f9b16db62887702e9f1d1c4c9f796d0121a46c8e184",
        urls = [
            "https://mirror.bazel.build/github.com/bazelbuild/rules_go/releases/download/v0.46.0/rules_go-v0.46.0.zip",
            "https://github.com/bazelbuild/rules_go/releases/download/v0.46.0/rules_go-v0.46.0.zip",
        ],
    ),
    rules_python = dict(
        sha256 = "5868e73107a8e85d8f323806e60cad7283f34b32163ea6ff1020cf27abef6036",
        strip_prefix = "rules_python-0.25.0",
        urls = ["https://github.com/bazelbuild/rules_python/releases/download/0.25.0/rules_python-0.25.0.tar.gz"],
    ),
)
