BINARY_NAME=mytodo
VERSION=$(shell git describe --tags --always)
BUILD_TIME=$(shell date -u '+%Y-%m-%d_%H:%M:%S')
COMMIT_HASH=$(shell git rev-parse --short HEAD)

# Go Setup
GO=go
GOFLAGS=-trimpath
LDFLAGS=-X 'main.version=$(VERSION)' \
		-X 'main.buildTime=$(BUILD_TIME)' \
		-X 'main.commitHash=$(COMMIT_HASH)' \
		-s -w

# Support Platforms
PLATFORMS=darwin/amd64 darwin/arm64 linux/amd64 linux/arm64 windows/amd64

# Clean
.PHONY: clean
clean:
	@echo "Cleaning..."
	@rm -f $(BINARY_NAME)
	@go clean

# Build
.PHONY: build
build:
	@echo "Building $(BINARY_NAME)..."
	@$(GO) build $(GOFLAGS) -ldflags "$(LDFLAGS)" -o $(BINARY_NAME)

# Help
.PHONY: help
help:
	@echo "Usage: make <target>"
	@echo ""
	@echo "Targets:"
	@echo "  clean   - Clean up"
	@echo "  build   - Build the binary"
	@echo "  help    - Show this help message"

# Default
.DEFAULT_GOAL := help