#!/usr/bin/env bash

# Script to generate protobuf files using Buf
# This script is used for local development and CI

set -euo pipefail

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
NC='\033[0m' # No Color

# Function to print colored output
print_status() {
    echo -e "${GREEN}[INFO]${NC} $1"
}

print_warning() {
    echo -e "${YELLOW}[WARN]${NC} $1"
}

print_error() {
    echo -e "${RED}[ERROR]${NC} $1"
}

# Check if Buf is installed
if ! command -v buf &> /dev/null; then
    print_error "Buf CLI is not installed."
    echo ""
    echo "Please install Buf using one of the following methods:"
    echo ""
    echo "macOS:"
    echo "  brew install bufbuild/buf/buf"
    echo ""
    echo "Linux/Windows:"
    echo "  curl -sSL \"https://github.com/bufbuild/buf/releases/latest/download/buf-\$(uname -s)-\$(uname -m)\" -o /usr/local/bin/buf"
    echo "  chmod +x /usr/local/bin/buf"
    echo ""
    echo "For more options, see: https://docs.buf.build/installation"
    exit 1
fi

print_status "Using Buf version: $(buf --version)"

# Navigate to repository root
SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
REPO_ROOT="$(cd "$SCRIPT_DIR/.." && pwd)"
cd "$REPO_ROOT"

# Validate buf.yaml and buf.gen.yaml exist
if [[ ! -f "buf.yaml" ]]; then
    print_error "buf.yaml not found in repository root"
    exit 1
fi

if [[ ! -f "buf.gen.go.yaml" ]]; then
    print_error "buf.gen.go.yaml not found in repository root"
    exit 1
fi

if [[ ! -f "buf.gen.python.yaml" ]]; then
    print_error "buf.gen.python.yaml not found in repository root"
    exit 1
fi

print_status "Validating Buf configuration..."
if ! buf config ls-modules; then
    print_error "Invalid Buf configuration"
    exit 1
fi

# Update dependencies
print_status "Updating Buf dependencies..."
if ! buf dep update; then
    print_error "Failed to update Buf dependencies"
    exit 1
fi

# Clean existing generated files
print_status "Cleaning existing generated files..."
find go -type f -name "*.go" -delete
find python -type f -name "*_pb2.py" -delete


# Generate protobuf files
print_status "Generating Go protobuf files..."
if ! buf generate --template buf.gen.go.yaml; then
    print_error "Failed to generate Go protobuf files"
    exit 1
fi

print_status "Generating Python protobuf files..."
if ! buf generate --template buf.gen.python.yaml; then
    print_error "Failed to generate Python protobuf files"
    exit 1
fi

# Create missing __init__.py files for Python packages
echo "Creating missing __init__.py files for Python packages..."
find python/udpa -type d -exec sh -c 'for dir; do [ ! -f "$dir/__init__.py" ] && touch "$dir/__init__.py"; done' sh {} +
find python/xds -type d -exec sh -c 'for dir; do [ ! -f "$dir/__init__.py" ] && touch "$dir/__init__.py"; done' sh {} +

# Verify generated files exist
GO_FILES=$(find go -name "*.pb.go" 2>/dev/null | wc -l)
PYTHON_FILES=$(find python -name "*_pb2.py" 2>/dev/null | wc -l)

print_status "Generated $GO_FILES Go protobuf files"
print_status "Generated $PYTHON_FILES Python protobuf files"

if [[ $GO_FILES -eq 0 || $PYTHON_FILES -eq 0 ]]; then
    print_warning "No files generated. Check your proto files and buf.gen.yaml configuration."
    exit 1
fi

print_status "Protobuf code generation completed successfully!"
print_status "Remember to commit both your proto changes and the generated files."