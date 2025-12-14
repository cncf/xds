# Developer documentation

## Synchronize generated files

Run the following command to update the generated files and commit them with your change:

```sh
bazel build //...
tools/generate_go_protobuf.py
```

## Buf tools

### Format proto files

To format all proto files in the workspace:

```sh
bazel run //:buf_format
```

To check if proto files are formatted correctly (used in CI):

```sh
bazel run //:buf_format -- -d
```

### Linting

Buf linting is automatically run as part of the test suite:

```sh
bazel test //...
```
