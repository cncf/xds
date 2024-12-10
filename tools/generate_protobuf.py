#!/usr/bin/env python3

from subprocess import check_output
import glob
import os
import shutil

def generate_go_protobufs(output,language):
  bazel_bin = check_output(['bazel', 'info', 'bazel-bin']).decode().strip()

  if language == "go":
    protos = check_output([
        'bazel',
        'query',
        'kind("go_proto_library", ...)',
    ]).split()
    output_dir = 'github.com/cncf/xds/python'
  if language == "python":
    protos = check_output([
        'bazel',
        'query',
        'kind("py_proto_library", ...)',
    ]).split()
    output_dir = 'github.com/cncf/xds/go'
  check_output(['bazel', 'build', '-c', 'fastbuild'] + protos)

  for rule in protos:
    rule_dir  = rule.decode()[2:].rsplit(':')[0]
    if language == "go":
      input_dir = os.path.join(bazel_bin, rule_dir, 'pkg_go_proto_',
                              'github.com/cncf/xds/go', rule_dir)
      input_files = glob.glob(os.path.join(input_dir, '*.go'))
    elif language == "python":
      input_dir = os.path.join(bazel_bin, rule_dir)
      input_files = glob.glob(os.path.join(input_dir, '*_pb2.py'))
    
    output_dir = os.path.join(output, rule_dir)

    # Ensure the output directory exists
    os.makedirs(output_dir, 0o755, exist_ok=True)
    for generated_file in input_files:
      output_file = shutil.copy(generated_file, output_dir)
      os.chmod(output_file, 0o644)


if __name__ == "__main__":
  workspace = check_output(['bazel', 'info', 'workspace']).decode().strip()
  output = os.path.join(workspace, 'go')
  generate_go_protobufs(output, "go")
  output = os.path.join(workspace, 'python')
  generate_go_protobufs(output, "python")