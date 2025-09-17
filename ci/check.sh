#!/usr/bin/env bash

set -e

bazel test --config=ci //...

find go/xds -name "*.pb.go" -delete
find go/udpa -name "*.pb.go" -delete
find python/xds -name "*_pb2.py" -delete
find python/udpa -name "*_pb2.py" -delete
find python/validate -name "*_pb2.py" -delete

tools/buf_generate.sh

git add go/xds go/udpa python/xds python/udpa python/validate

echo "If this check fails, apply following diff:"
git diff HEAD
git diff HEAD --quiet
