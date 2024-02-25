load("@com_google_googleapis//:repository_rules.bzl", "switched_rules_by_language")

def _extension_for_com_google_googleapis_impl(ctx):
    switched_rules_by_language(
        name = "com_google_googleapis_imports",
        cc = True,
        go = True,
        python = True,
        grpc = True,
    )

extension_for_com_google_googleapis = module_extension(implementation = _extension_for_com_google_googleapis_impl)
