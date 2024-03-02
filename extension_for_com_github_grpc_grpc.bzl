load("@com_github_grpc_grpc//bazel:grpc_deps.bzl", "grpc_deps")
load("@com_github_grpc_grpc//bazel:grpc_extra_deps.bzl", "grpc_extra_deps")

def _extension_for_com_github_grpc_grpc_impl(ctx):
    grpc_deps()
    grpc_extra_deps()

extension_for_com_github_grpc_grpc = module_extension(implementation = _extension_for_com_github_grpc_grpc_impl)
