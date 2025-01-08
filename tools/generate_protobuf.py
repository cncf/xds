#!/usr/bin/env python3

from subprocess import check_output
import glob
import os
import shutil

workspace = check_output(["bazel", "info", "workspace"]).decode().strip()
SUPPORTED_LANGUAGES = ["go", "python"]


def generate_lang_files_from_protos(language: str):
    output = os.path.join(workspace, language)
    bazel_bin = check_output(["bazel", "info", "bazel-bin"]).decode().strip()

    protos = (
        check_output(
            [
                "bazel",
                "query",
                f'kind("{language}_proto_library", ...)',
            ]
        )
        .decode()
        .split()
    )
    output_dir = f"github.com/cncf/xds/{language}"
    check_output(["bazel", "build", "-c", "fastbuild"] + protos)

    for rule in protos:
        rule_dir = rule.decode()[2:].rsplit(":")[0]
        input_dir = {
            "go": os.path.join(
                bazel_bin, rule_dir, "pkg_go_proto_", "github.com/cncf/xds/go", rule_dir
            ),
            "python": os.path.join(bazel_bin, rule_dir),
        }[language]
        input_files = {
            "go": glob.glob(os.path.join(input_dir, "*.go")),
            "python": glob.glob(os.path.join(input_dir, "*_pb2.py")),
        }[language]
        if language == "go":
            input_dir = os.path.join(
                bazel_bin, rule_dir, "pkg_go_proto_", "github.com/cncf/xds/go", rule_dir
            )
            input_files = glob.glob(os.path.join(input_dir, "*.go"))
        elif language == "python":
            input_dir = os.path.join(bazel_bin, rule_dir)
            input_files = glob.glob(os.path.join(input_dir, "*_pb2.py"))

        output_dir = os.path.join(output, rule_dir)

        # Ensure the output directory exists
        os.makedirs(output_dir, 0o755, exist_ok=True)
        for generated_file in input_files:
            output_file = shutil.copy(generated_file, output_dir)
            os.chmod(output_file, 0o644)


if __name__ == "__main__":
    for language in SUPPORTED_LANGUAGES:
        generate_lang_files_from_protos(language=language)
