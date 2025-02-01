#!/usr/bin/env bash

set -e

bazel test --config=ci //...

find xds/ -name *.go -type f -exec rm -rf {} \;
find udpa/ -name *.go -type f -exec rm -rf {} \;

tools/generate_lang_files_from_protos.py

git add xds udpa

echo "If this check fails, apply following diff:"
git diff HEAD
git diff HEAD --quiet
