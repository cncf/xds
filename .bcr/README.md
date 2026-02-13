# Bazel Central Registry (BCR) Publishing Configuration

This directory contains the configuration files for publishing the xds module to the [Bazel Central Registry](https://github.com/bazelbuild/bazel-central-registry).

## Overview

The BCR publishing process is automated via GitHub Actions. When a new release is created, the workflow will automatically create a pull request to the Bazel Central Registry with the new version.

## Configuration Files

### metadata.template.json
Contains metadata about the xds module including:
- Homepage URL
- Maintainers list
- Repository references

The `versions` array will be automatically populated by the BCR publishing workflow.

### source.template.json
Template for the `source.json` file that will be created in the BCR. It specifies:
- Archive URL pattern: `https://github.com/cncf/xds/archive/{TAG}.tar.gz`
- Strip prefix pattern for extracting the archive
- Integrity hash (automatically computed)

### presubmit.yml
Defines the tests that BCR will run to validate the module before accepting it. The tests run on multiple platforms (Debian, macOS, Ubuntu, Windows) with different Bazel versions. The verification builds all targets in the xds module using `@xds//...`.

## GitHub Actions Workflow

The `.github/workflows/publish.yaml` workflow automates the publishing process. It:

1. Triggers on new releases or manual workflow dispatch
2. Uses the `bazel-contrib/publish-to-bcr` reusable workflow
3. Opens a pull request to the Bazel Central Registry

## Personal Access Token (PAT) Setup

To enable automated BCR publishing, you need to configure a GitHub Personal Access Token:

### Creating the PAT

1. Go to GitHub Settings → Developer settings → Personal access tokens → Tokens (classic)
2. Click "Generate new token (classic)"
3. Give it a descriptive name (e.g., "BCR Publish Token for xds")
4. Set an appropriate expiration (recommend 1 year and renew regularly)
5. Select the following scopes:
   - `repo` (Full control of private repositories)
   - `workflow` (Update GitHub Action workflows)
6. Click "Generate token" and copy the token value

### Configuring the Secret

1. Go to the xds repository settings: `https://github.com/cncf/xds/settings/secrets/actions`
2. Click "New repository secret"
3. Name: `BCR_PUBLISH_TOKEN`
4. Value: Paste the PAT you created
5. Click "Add secret"

### Registry Fork

The workflow requires a fork of the Bazel Central Registry. When the workflow runs, it will:

1. Push the new module version to your fork of the BCR (specified in `registry_fork`)
2. Create a new branch in your fork with the module version
3. Open a pull request from your fork to `bazelbuild/bazel-central-registry` (the main BCR)

**Important**: You must update the `registry_fork` parameter in `.github/workflows/publish.yaml`:
- Change `REPLACE_WITH_YOUR_FORK/bazel-central-registry` to your actual fork
- Example: `adisuissa/bazel-central-registry` if adisuissa owns the PAT

To create a fork:
1. Go to https://github.com/bazelbuild/bazel-central-registry
2. Click "Fork" button
3. Create the fork under your personal account or organization
4. Update the workflow file with your fork name

## Manual Publishing

If needed, you can manually trigger the publishing workflow:

1. Go to the Actions tab in the repository
2. Select "Publish to BCR" workflow
3. Click "Run workflow"
4. Enter the tag name (e.g., `v1.0.0`)
5. Click "Run workflow"

## Troubleshooting

### Common Issues

1. **PAT expired**: Renew the token and update the `BCR_PUBLISH_TOKEN` secret
2. **Permission denied**: Ensure the PAT has `repo` and `workflow` scopes
3. **Build failures**: Check the BCR presubmit results and fix any issues in the test module

### Verifying the Configuration

You can verify the templates are valid by checking:
- metadata.template.json has valid JSON syntax
- source.template.json uses correct URL patterns
- presubmit.yml has valid YAML syntax

## References

- [Bazel Central Registry](https://github.com/bazelbuild/bazel-central-registry)
- [publish-to-bcr GitHub Action](https://github.com/bazel-contrib/publish-to-bcr)
- [Bzlmod User Guide](https://bazel.build/docs/bzlmod)
