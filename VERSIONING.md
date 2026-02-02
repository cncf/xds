# Versioning Guide

This document describes the versioning and release process for the xDS project.

## Overview

The xDS project uses a unified version across all language ecosystems:

- **Python**: Published to PyPI as the `xds` package
- **Go**: Published as Go modules at `github.com/cncf/xds/go`
- **Bazel**: Published to Bazel Central Registry (BCR) as `xds` module

All ecosystems share the same version number and are released together.

## Version Format

All versions follow [Semantic Versioning 2.0.0](https://semver.org/):

- **MAJOR** version for incompatible API changes
- **MINOR** version for backwards-compatible functionality additions
- **PATCH** version for backwards-compatible bug fixes

## Tag Format

A single tag format is used for all ecosystems:

- Releases: `v1.2.3`

Pushing a `v*` tag triggers releases for Python (PyPI), Go (GitHub release), and Bazel (BCR) simultaneously.

## Release Process

### Manual Version Bump

Use the Version Bump workflow to update version numbers across all ecosystems:

1. Go to **Actions** → **Version Bump**
2. Click **Run workflow**
3. Enter the new version (e.g., `1.2.3`)
4. Choose whether to create a tag immediately
5. Click **Run workflow**

This will:
- Update the version in Python's `pyproject.toml`
- Update the version in Go's `VERSION` file
- Update the version in Bazel's `MODULE.bazel`
- Commit the changes
- Optionally create and push a version tag

### Automated Release

When a `v*` tag is pushed, all release workflows are triggered simultaneously:

#### Release Creation
1. A unified GitHub release is created with release notes
2. Release includes installation instructions for all ecosystems

#### Python
1. Builds the Python package
2. Publishes to PyPI
3. Uploads Python distribution artifacts to the release

#### Go
1. Verifies the Go module
2. Runs tests

#### Bazel
1. Automatically publishes to Bazel Central Registry using the `publish-to-bcr` reusable workflow
2. Generates BCR entry files (MODULE.bazel, source.json, presubmit.yml, metadata.json)
3. Creates attestations for security verification
4. Opens a pull request to bazelbuild/bazel-central-registry


### Creating a Release Manually

To create a unified release for all ecosystems:

```bash
# Update version in all ecosystem files
cd python
sed -i 's/version = ".*"/version = "X.Y.Z"/' pyproject.toml
cd ../go
echo "X.Y.Z" > VERSION
cd ..
sed -i 's/version = ".*"/version = "X.Y.Z"/' MODULE.bazel

# Commit changes
git add python/pyproject.toml go/VERSION MODULE.bazel
git commit -m "chore: bump version to X.Y.Z"
git push

# Create and push tag (this will trigger all release workflows)
git tag vX.Y.Z
git push origin vX.Y.Z

# The following happens automatically:
# - Python package is built and published to PyPI
# - Go module release is created on GitHub
# - publish-to-bcr workflow generates BCR entry and opens PR to bazel-central-registry
```

## Publishing Credentials

### Python (PyPI)

The Python release workflow uses OIDC trusted publishing. To configure:

1. Go to PyPI → Account Settings → Publishing
2. Add a new publisher:
   - PyPI Project Name: `xds`
   - Owner: `cncf`
   - Repository: `xds`
   - Workflow: `python-release.yml`
   - Environment: (leave empty)

No API tokens are needed with OIDC.

### Go Modules

Go modules are automatically published when tags are pushed. No additional setup required.

### Bazel (BCR)

Bazel modules are automatically published to the Bazel Central Registry using the `publish-to-bcr` reusable workflow:

#### Setup (One-time)

1. **Fork the BCR**: Fork [bazel-central-registry](https://github.com/bazelbuild/bazel-central-registry) to the CNCF organization or appropriate account
2. **Create a Personal Access Token (PAT)**:
   - Go to GitHub Settings → Developer settings → Personal access tokens → Tokens (classic)
   - Create a new token with `repo` and `workflow` scopes
   - Save it as a repository secret named `BCR_PUBLISH_TOKEN`
3. **Update the workflow**: Set `registry_fork` in `.github/workflows/publish-to-bcr.yml` to the CNCF fork (currently set to `cncf/bazel-central-registry`)

#### How It Works

When you push a `v*` tag:
1. The workflow automatically generates BCR entry files from `.bcr` templates
2. Creates security attestations
3. Opens a pull request to the BCR repository
4. BCR maintainers review and merge the PR

No manual file copying or PR creation needed!

See [publish-to-bcr documentation](https://github.com/bazel-contrib/publish-to-bcr) for more details.
