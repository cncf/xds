load("@com_google_protobuf//:protobuf_deps.bzl", "protobuf_deps")
load("@bazel_gazelle//:deps.bzl", "gazelle_dependencies", "go_repository")
load("@com_google_googleapis//:repository_rules.bzl", "switched_rules_by_language")
load("@io_bazel_rules_go//go:deps.bzl", "go_register_toolchains", "go_rules_dependencies")
load("@com_envoyproxy_protoc_gen_validate//bazel:repositories.bzl", "pgv_dependencies")
load("@rules_proto//proto:repositories.bzl", "rules_proto_dependencies", "rules_proto_toolchains")

# go version for rules_go
GO_VERSION = "1.20.2"

def xds_dependency_imports(go_version = GO_VERSION):
    rules_proto_dependencies()
    rules_proto_toolchains()
    protobuf_deps()
    go_rules_dependencies()
    go_register_toolchains(go_version)
    gazelle_dependencies()
    pgv_dependencies()

    # Needed for grpc's @com_github_grpc_grpc//bazel:python_rules.bzl
    # Used in place of calling grpc_deps() because it needs to be called before
    # loading `grpc_extra_deps.bzl` - which is not allowed in a method def context.
    native.bind(
        name = "protocol_compiler",
        actual = "@com_google_protobuf//:protoc",
    )

    switched_rules_by_language(
        name = "com_google_googleapis_imports",
        cc = True,
        go = True,
        python = True,
        grpc = True,
    )

    # These dependencies, like most of the Go in this repository, exist only for the API.
    # These repos also have transient dependencies - `build_external` allows them to use them.
    # TODO(phlax): remove `build_external` and pin all transients
    go_repository(
        name = "com_github_iancoleman_strcase",
        importpath = "github.com/iancoleman/strcase",
        sum = "h1:ux/56T2xqZO/3cP1I2F86qpeoYPCOzk+KF/UH/Ar+lk=",
        version = "v0.0.0-20180726023541-3605ed457bf7",
        build_external = "external",
        # project_url = "https://pkg.go.dev/github.com/iancoleman/strcase",
        # last_update = "2020-11-22"
        # use_category = ["api"],
        # source = "https://github.com/bufbuild/protoc-gen-validate/blob/v0.6.1/dependencies.bzl#L23-L28"
    )
    go_repository(
        name = "org_golang_x_net",
        importpath = "golang.org/x/net",
        sum = "h1:0mm1VjtFUOIlE1SbDlwjYaDxZVDP2S5ou6y0gSgXHu8=",
        version = "v0.0.0-20200226121028-0de0cce0169b",
        build_external = "external",
        # project_url = "https://pkg.go.dev/golang.org/x/net",
        # last_update = "2020-02-26"
        # use_category = ["api"],
        # source = "https://github.com/bufbuild/protoc-gen-validate/blob/v0.6.1/dependencies.bzl#L129-L134"
    )
    go_repository(
        name = "org_golang_x_text",
        importpath = "golang.org/x/text",
        sum = "h1:cokOdA+Jmi5PJGXLlLllQSgYigAEfHXJAERHVMaCc2k=",
        version = "v0.3.3",
        build_external = "external",
        # project_url = "https://pkg.go.dev/golang.org/x/text",
        # last_update = "2021-06-16"
        # use_category = ["api"],
        # source = "https://github.com/bufbuild/protoc-gen-validate/blob/v0.6.1/dependencies.bzl#L148-L153"
    )
    go_repository(
        name = "com_github_spf13_afero",
        importpath = "github.com/spf13/afero",
        sum = "h1:8q6vk3hthlpb2SouZcnBVKboxWQWMDNF38bwholZrJc=",
        version = "v1.3.4",
        build_external = "external",
        # project_url = "https://pkg.go.dev/github.com/spf13/afero",
        # last_update = "2021-03-20"
        # use_category = ["api"],
        # source = "https://github.com/bufbuild/protoc-gen-validate/blob/v0.6.1/dependencies.bzl#L60-L65"
    )
    go_repository(
        name = "com_github_lyft_protoc_gen_star",
        importpath = "github.com/lyft/protoc-gen-star/v2",
        sum = "h1:keaAo8hRuAT0O3DfJ/wM3rufbAjGeJ1lAtWZHDjKGB0=",
        version = "v2.0.1",
        build_external = "external",
        # project_url = "https://pkg.go.dev/github.com/lyft/protoc-gen-star",
        # last_update = "2023-01-06"
        # use_category = ["api"],
        # source = "https://github.com/bufbuild/protoc-gen-validate/blob/v0.10.1/dependencies.bzl#L35-L40"
    )
    go_repository(
        name = "org_golang_google_protobuf",
        importpath = "google.golang.org/protobuf",
        sum = "h1:d0NfwRgPtno5B1Wa6L2DAG+KivqkdutMf1UhdNx175w=",
        version = "v1.28.1",
        build_external = "external",
    )
    go_repository(
        name = "org_golang_google_grpc",
        build_file_proto_mode = "disable",
        importpath = "google.golang.org/grpc",
        sum = "h1:AGJ0Ih4mHjSeibYkFGh1dD9KJ/eOtZ93I6hoHhukQ5Q=",
        version = "v1.40.0",
    )


# Old name for backward compatibility.
# TODO(roth): Remove this once callers are migrated to the new name.
def udpa_dependency_imports(go_version = GO_VERSION):
  xds_dependency_imports(go_version=go_version)
