load("@bazel_gazelle//internal:go_repository.bzl", "go_repository")

def _extension_for_bazel_gazelle_impl(_):
    go_repository(
        name = "com_github_lyft_protoc_gen_star",
        importpath = "github.com/lyft/protoc-gen-star/v2",
        version = "v2.0.1",
        sum = "h1:keaAo8hRuAT0O3DfJ/wM3rufbAjGeJ1lAtWZHDjKGB0=",
        build_external = "external",
    )
    go_repository(
        name = "com_github_spf13_afero",
        importpath = "github.com/spf13/afero",
        version = "v1.3.4",
        sum = "h1:8q6vk3hthlpb2SouZcnBVKboxWQWMDNF38bwholZrJc=",
        build_external = "external",
    )

extension_for_bazel_gazelle = module_extension(implementation = _extension_for_bazel_gazelle_impl)
