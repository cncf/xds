#!/usr/bin/env bash

set -e

bazel run @com_github_bufbuild_buf//:bin/buf lint -- --path xds
bazel run @com_github_bufbuild_buf//:bin/buf lint -- --path udpa

bazel test --config=ci //...

rm -rf go/xds go/udpa

tools/generate_go_protobuf.py

git add go/xds go/udpa

echo "If this check fails, apply following diff:"
git diff HEAD
git diff HEAD --quiet
