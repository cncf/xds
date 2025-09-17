# Developer documentation

## Synchronize generated files

Run the following command to update the generated files and commit them with your change:

```sh
tools/buf_generate.sh
```

This uses **buf gen v2** with remote plugins to generate Go and Python protobuf files, replacing the previous Bazel-based approach.

### Buf Generation Details

The new buf-based generation provides:

- **Modern protobuf generation**: Uses buf with remote plugins (no local dependencies required)
- **Separate language configs**: `buf.gen.go.yaml` for Go, `buf.gen.python.yaml` for Python
- **Full validation support**: protoc-gen-validate for both languages
- **Version alignment**: All plugin versions match original Bazel configuration

#### Generated Files
- **Go**: `go/` directory with all `.pb.go`, `_grpc.pb.go`, and `.pb.validate.go` files
- **Python**: `python/` directory with all `_pb2.py` files and validation support

#### API Build System Compatibility

The existing `xds_proto_package()` macro continues to work unchanged. The implementation acts as a façade that internally coordinates with buf-generated files while maintaining full backward compatibility.

No changes to existing BUILD files in proto directories are required.

### Alternative: Buf + Gazelle Integration (Advanced)

For advanced workflows, you can use the integrated approach:

```sh
tools/update_buf_and_gazelle.sh
```

This provides automatic BUILD file management for Go packages via Gazelle.

### Requirements

- [Buf CLI](https://docs.buf.build/installation) for protobuf generation
- Go 1.20+ for Go module builds
- Python 3.7+ with protobuf package for Python imports

### Configuration

Protobuf generation is configured via:
- `buf.yaml` - Module configuration and dependencies
- `buf.gen.go.yaml` - Go-specific generation with remote plugins
- `buf.gen.python.yaml` - Python-specific generation with module inputs

All dependency versions are pinned to match the original Bazel configuration exactly.
