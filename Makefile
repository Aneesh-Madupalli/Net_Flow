.PHONY: build build-windows build-macos build-linux build-all clean deps test

# Build for current platform
build:
	go build -ldflags="-s -w" -o netflow

# Build for Windows (no console window; GUI/tray only)
build-windows:
	@echo "Building for Windows..."
	GOOS=windows GOARCH=amd64 go build -ldflags="-s -w -H windowsgui" -o netflow-windows-amd64.exe

# Build for macOS (both architectures)
build-macos:
	@echo "Building for macOS (Intel)..."
	GOOS=darwin GOARCH=amd64 go build -ldflags="-s -w" -o netflow-macos-amd64
	@echo "Building for macOS (Apple Silicon)..."
	GOOS=darwin GOARCH=arm64 go build -ldflags="-s -w" -o netflow-macos-arm64

# Build for Linux
build-linux:
	@echo "Building for Linux..."
	GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o netflow-linux-amd64

# Build for all platforms
build-all: build-windows build-macos build-linux
	@echo "All builds completed!"

# Install dependencies
deps:
	go mod download
	go mod tidy

# Clean build artifacts
clean:
	rm -f netflow netflow-*.exe netflow-*-amd64 netflow-*-arm64

# Run tests (if any)
test:
	go test ./...

# Format code
fmt:
	go fmt ./...

# Run linter
lint:
	golangci-lint run || true

