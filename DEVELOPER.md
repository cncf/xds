# Developer documentation

## Synchronize generated files

Run the following command to update the generated files and commit them with your change:

```sh
bazel build //...
tools/generate_go_protobuf.py
```

## Versioning and Releases

For information about versioning and the release process, see [VERSIONING.md](VERSIONING.md).

### Quick Release Guide

To create a new release:

1. Use the GitHub Actions **Version Bump** workflow to update the version
2. Or manually create a tag: `git tag python/v1.2.3` or `git tag go/v1.2.3`
3. Push the tag: `git push origin <tag-name>`
4. The release workflow will automatically build and publish the package

