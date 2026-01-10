#!/usr/bin/env python3

from subprocess import check_output
import glob
import os
import shutil
import fnmatch
import stat
from language_config import GoConfig, PythonConfig, LanguageConfig

workspace = check_output(["bazel", "info", "workspace"]).decode().strip()
LANG_CONFIGS = {config.language: config for config in [GoConfig(), PythonConfig()]}

def _realpath(path: str) -> str:
    return os.path.realpath(os.path.abspath(path))

def _is_within(path: str, root: str) -> bool:
    path_r = _realpath(path)
    root_r = _realpath(root)
    try:
        return os.path.commonpath([path_r, root_r]) == root_r
    except ValueError:
        return False

def _safe_copy_generated_file(src_path: str, dst_dir: str, *, bazel_bin: str, expected_input_dir: str) -> str:
    # Fail-closed security checks: treat bazel-bin outputs as untrusted filesystem inputs.
    if not _is_within(src_path, expected_input_dir):
        raise RuntimeError(f"Refusing to copy file outside expected input dir: {src_path}")
    if not _is_within(src_path, bazel_bin):
        raise RuntimeError(f"Refusing to copy file outside bazel-bin: {src_path}")

    # lstat() so we can detect symlinks without following them.
    st1 = os.stat(src_path, follow_symlinks=False)

    if stat.S_ISLNK(st1.st_mode):
        raise RuntimeError(f"Refusing to copy symlink output (sandbox escape risk): {src_path}")
    if not stat.S_ISREG(st1.st_mode):
        raise RuntimeError(f"Refusing to copy non-regular output: {src_path}")
    # Prevent hardlink-based exfiltration (e.g., link to /etc/passwd on same filesystem).
    if st1.st_nlink != 1:
        raise RuntimeError(f"Refusing to copy linked file (nlink={st1.st_nlink}): {src_path}")

    os.makedirs(dst_dir, 0o755, exist_ok=True)
    dst_path = os.path.join(dst_dir, os.path.basename(src_path))

    # Open source/dest with O_NOFOLLOW when available to prevent symlink races.
    src_flags = os.O_RDONLY
    if hasattr(os, "O_NOFOLLOW"):
        src_flags |= os.O_NOFOLLOW

    dst_flags = os.O_WRONLY | os.O_CREAT | os.O_TRUNC
    if hasattr(os, "O_NOFOLLOW"):
        dst_flags |= os.O_NOFOLLOW

    src_fd = os.open(src_path, src_flags)
    try:
        st2 = os.fstat(src_fd)
        # Ensure TOCTOU safety: same inode/dev as lstat result.
        if (st1.st_ino, st1.st_dev, st1.st_nlink) != (st2.st_ino, st2.st_dev, st2.st_nlink):
            raise RuntimeError(f"Refusing to copy due to source race/replace: {src_path}")

        dst_fd = os.open(dst_path, dst_flags, 0o644)
        try:
            with os.fdopen(src_fd, "rb", closefd=False) as src_f, os.fdopen(dst_fd, "wb", closefd=False) as dst_f:
                shutil.copyfileobj(src_f, dst_f)
        finally:
            os.close(dst_fd)

        os.chmod(dst_path, 0o644)
        return dst_path
    finally:
        os.close(src_fd)

def generate_lang_files_from_protos(language_config: LanguageConfig):
    language = language_config.language
    print(f"Generating proto code in language {language}")
    output = os.path.join(workspace, language)
    bazel_bin = check_output(["bazel", "info", "bazel-bin"]).decode().strip()

    protos = check_output(
        [
            "bazel",
            "query",
            language_config.bazel_query_kind,
        ]
    ).split()
    output_dir = f"github.com/cncf/xds/{language}"
    check_output(["bazel", "build", "-c", "fastbuild"] + protos)
    print(f"Found {len(protos)} bazel rules to generate code")
    for rule in protos:
        rule_dir = rule.decode()[2:].rsplit(":")[0]
        input_dir = language_config.get_input_dir(bazel_bin, rule_dir)

        if not _is_within(input_dir, bazel_bin):
            raise RuntimeError(f"Refusing to read inputs outside bazel-bin: {input_dir}")

        # Prefer directory listing + pattern match to keep the search scope explicit.
        try:
            names = os.listdir(input_dir)
        except FileNotFoundError:
            # Some rules may not emit language outputs in expected locations; skip safely.
            print(f"Skipping missing input dir: {input_dir}")
            continue

        input_files = [
            os.path.join(input_dir, name)
            for name in names
            if fnmatch.fnmatch(name, language_config.generated_file_pattern)
        ]

        output_dir = os.path.join(output, rule_dir)
        print(f"Moving {len(input_files)} generated files from {input_dir} to output_dir {output_dir}")
        os.makedirs(output_dir, 0o755, exist_ok=True)

        for generated_file in input_files:
            copied = _safe_copy_generated_file(
                generated_file,
                output_dir,
                bazel_bin=bazel_bin,
                expected_input_dir=input_dir,
            )
            # Ensure the output directory exists
            os.makedirs(output_dir, 0o755, exist_ok=True)
            for generated_file in input_files:
                output_file = shutil.copy(generated_file, output_dir)
                os.chmod(output_file, 0o644)


if __name__ == "__main__":
    for config in LANG_CONFIGS.values():
        generate_lang_files_from_protos(language_config=config)