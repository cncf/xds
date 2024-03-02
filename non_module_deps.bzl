load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")
# -- load statements -- #

def _non_module_deps_impl(_):
  http_archive(
    name = "com_envoyproxy_protoc_gen_validate",
    urls = [
      "https://github.com/envoyproxy/protoc-gen-validate/archive/refs/tags/v0.6.1.tar.gz"
    ],
    sha256 = "c695fc5a2e5a1b52904cd8a58ce7a1c3a80f7f50719496fd606e551685c01101",
    strip_prefix = "protoc-gen-validate-0.6.1",
  )
  http_archive(
    name = "com_github_grpc_grpc",
    sha256 = "916f88a34f06b56432611aaa8c55befee96d0a7b7d7457733b9deeacbc016f99",
    strip_prefix = "grpc-1.59.1",
    urls = [
      "https://github.com/grpc/grpc/archive/refs/tags/v1.59.1.tar.gz",
    ],
  )
  http_archive(
    name = "com_google_googleapis",
    urls = [
      "https://github.com/googleapis/googleapis/archive/114a745b2841a044e98cdbb19358ed29fcf4a5f1.tar.gz"
    ],
    sha256 = "9b4e0d0a04a217c06b426aefd03b82581a9510ca766d2d1c70e52bb2ad4a0703",
    strip_prefix = "googleapis-114a745b2841a044e98cdbb19358ed29fcf4a5f1",
  )
  http_archive(
    name = "org_golang_x_tools",
    urls = [
      "https://mirror.bazel.build/github.com/golang/tools/archive/refs/tags/v0.7.0.zip",
      "https://github.com/golang/tools/archive/refs/tags/v0.7.0.zip"
    ],
    sha256 = "9f20a20f29f4008d797a8be882ef82b69cf8f7f2b96dbdfe3814c57d8280fa4b",
    strip_prefix = "tools-0.7.0",
    patches = [
      "@io_bazel_rules_go//third_party:org_golang_x_tools-deletegopls.patch",
      "@io_bazel_rules_go//third_party:org_golang_x_tools-gazelle.patch"
    ],
    patch_args = [
      "-p1"
    ],
  )
# -- repo definitions -- #

non_module_deps = module_extension(implementation = _non_module_deps_impl)
