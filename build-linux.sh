#!/bin/bash
echo "Building NetFlow for Linux..."

GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o netflow-linux-amd64
if [ $? -eq 0 ]; then
    echo "Build successful: netflow-linux-amd64"
else
    echo "Build failed!"
    exit 1
fi

