load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")

def _non_module_dependencies_impl(_ctx):
    http_archive(
        name = "com_google_googleapis",
        sha256 = "9d1a930e767c93c825398b8f8692eca3fe353b9aaadedfbcf1fca2282c85df88",
        strip_prefix = "googleapis-64926d52febbf298cb82a8f472ade4a3969ba922",
        urls = [
            "https://github.com/googleapis/googleapis/archive/64926d52febbf298cb82a8f472ade4a3969ba922.zip",
        ],
    )

non_module_dependencies = module_extension(
    implementation = _non_module_dependencies_impl,
)
