# archiTerm Makefile
# Cross-platform build for Windows, Linux, and macOS

APP_NAME := architerm
VERSION := 0.1.0
BUILD_DIR := build
GO := go

# Build flags
LDFLAGS := -ldflags "-s -w -X 'github.com/architerm/architerm/cmd.version=$(VERSION)'"

# Determine OS
UNAME_S := $(shell uname -s 2>/dev/null || echo Windows)

.PHONY: all build clean test run install deps help

# Default target
all: build

# Install dependencies
deps:
	$(GO) mod download
	$(GO) mod tidy

# Build for current platform
build: deps
	$(GO) build $(LDFLAGS) -o $(BUILD_DIR)/$(APP_NAME) .

# Run the application
run: build
	./$(BUILD_DIR)/$(APP_NAME)

# Build for all platforms
build-all: build-linux build-darwin build-windows

# Build for Linux (amd64 and arm64)
build-linux:
	@echo "Building for Linux amd64..."
	GOOS=linux GOARCH=amd64 $(GO) build $(LDFLAGS) -o $(BUILD_DIR)/$(APP_NAME)-linux-amd64 .
	@echo "Building for Linux arm64..."
	GOOS=linux GOARCH=arm64 $(GO) build $(LDFLAGS) -o $(BUILD_DIR)/$(APP_NAME)-linux-arm64 .

# Build for macOS (amd64 and arm64)
build-darwin:
	@echo "Building for macOS amd64..."
	GOOS=darwin GOARCH=amd64 $(GO) build $(LDFLAGS) -o $(BUILD_DIR)/$(APP_NAME)-darwin-amd64 .
	@echo "Building for macOS arm64 (Apple Silicon)..."
	GOOS=darwin GOARCH=arm64 $(GO) build $(LDFLAGS) -o $(BUILD_DIR)/$(APP_NAME)-darwin-arm64 .

# Build for Windows (amd64)
build-windows:
	@echo "Building for Windows amd64..."
	GOOS=windows GOARCH=amd64 $(GO) build $(LDFLAGS) -o $(BUILD_DIR)/$(APP_NAME)-windows-amd64.exe .

# Run tests
test:
	$(GO) test -v ./...

# Run tests with coverage
test-coverage:
	$(GO) test -v -coverprofile=coverage.out ./...
	$(GO) tool cover -html=coverage.out -o coverage.html

# Clean build artifacts
clean:
	rm -rf $(BUILD_DIR)
	rm -f coverage.out coverage.html

# Install to GOPATH/bin
install: build
	$(GO) install .

# Create release archives
release: build-all
	@echo "Creating release archives..."
	@mkdir -p $(BUILD_DIR)/release
	@cd $(BUILD_DIR) && tar -czf release/$(APP_NAME)-$(VERSION)-linux-amd64.tar.gz $(APP_NAME)-linux-amd64
	@cd $(BUILD_DIR) && tar -czf release/$(APP_NAME)-$(VERSION)-linux-arm64.tar.gz $(APP_NAME)-linux-arm64
	@cd $(BUILD_DIR) && tar -czf release/$(APP_NAME)-$(VERSION)-darwin-amd64.tar.gz $(APP_NAME)-darwin-amd64
	@cd $(BUILD_DIR) && tar -czf release/$(APP_NAME)-$(VERSION)-darwin-arm64.tar.gz $(APP_NAME)-darwin-arm64
	@cd $(BUILD_DIR) && zip -q release/$(APP_NAME)-$(VERSION)-windows-amd64.zip $(APP_NAME)-windows-amd64.exe
	@echo "Release archives created in $(BUILD_DIR)/release/"

# Format code
fmt:
	$(GO) fmt ./...

# Lint code (requires golangci-lint)
lint:
	golangci-lint run

# Help
help:
	@echo "archiTerm Build System"
	@echo ""
	@echo "Usage:"
	@echo "  make              Build for current platform"
	@echo "  make build        Build for current platform"
	@echo "  make run          Build and run"
	@echo "  make build-all    Build for all platforms"
	@echo "  make build-linux  Build for Linux"
	@echo "  make build-darwin Build for macOS"
	@echo "  make build-windows Build for Windows"
	@echo "  make test         Run tests"
	@echo "  make clean        Clean build artifacts"
	@echo "  make install      Install to GOPATH/bin"
	@echo "  make release      Create release archives"
	@echo "  make deps         Download dependencies"
	@echo "  make fmt          Format code"
	@echo "  make lint         Lint code"
	@echo "  make help         Show this help"
