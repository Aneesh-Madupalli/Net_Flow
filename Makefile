.PHONY: build deps clean test fmt lint

# Build for current platform
build:
	go build -ldflags="-s -w" -o netflow

# On Windows, use: go build -ldflags="-s -w -H windowsgui" -o netflow.exe

# Install dependencies
deps:
	go mod download
	go mod tidy

# Clean build artifacts
clean:
	rm -f netflow netflow.exe netflow-*.exe netflow-*-amd64 netflow-*-arm64

# Run tests
test:
	go test ./...

# Format code
fmt:
	go fmt ./...

# Run linter
lint:
	golangci-lint run || true
