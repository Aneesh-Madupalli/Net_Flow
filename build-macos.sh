#!/bin/bash
echo "Building NetFlow for macOS..."

# Build for Intel (amd64)
echo "Building for Intel (amd64)..."
GOOS=darwin GOARCH=amd64 go build -ldflags="-s -w" -o netflow-macos-amd64
if [ $? -eq 0 ]; then
    echo "Build successful: netflow-macos-amd64"
else
    echo "Build failed for amd64!"
    exit 1
fi

# Build for Apple Silicon (arm64)
echo "Building for Apple Silicon (arm64)..."
GOOS=darwin GOARCH=arm64 go build -ldflags="-s -w" -o netflow-macos-arm64
if [ $? -eq 0 ]; then
    echo "Build successful: netflow-macos-arm64"
else
    echo "Build failed for arm64!"
    exit 1
fi

# Copy to release/ for README download links
mkdir -p release
cp -f netflow-macos-amd64 netflow-macos-arm64 release/ 2>/dev/null && chmod +x release/netflow-macos-*
echo "Copied to release/ for README download links."

echo "All macOS builds completed successfully!"

