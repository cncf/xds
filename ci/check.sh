#!/usr/bin/env bash

set -e

# Check buf format - this runs buf format in diff mode to verify formatting
echo "Checking proto file formatting with buf..."
bazel run //:buf_format -- -d || {
    echo "ERROR: Proto files are not properly formatted."
    echo "Run 'bazel run //:buf_format' to fix formatting issues."
    exit 1
}
echo "Proto formatting check passed."

# Run all tests including buf_lint_test via Bazel
bazel test --config=ci //...

rm -rf go/xds go/udpa

tools/generate_lang_files_from_protos.py

git add go/xds go/udpa

echo "If this check fails, apply following diff:"
git diff HEAD
git diff HEAD --quiet
