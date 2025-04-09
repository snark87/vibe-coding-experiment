# Makefile for Quantum Circuit Editor Go Backend

# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOTEST=$(GOCMD) test
GOVET=$(GOCMD) vet
GOMODTIDY=$(GOCMD) mod tidy
GOFMT=$(GOCMD) fmt
GORUN=$(GOCMD) run
BINARY_NAME=quantum-circuit-editor-backend
BINARY_UNIX=$(BINARY_NAME)_unix

# Linting tools
GOLANGCI_LINT=golangci-lint

# Build directory
BUILD_DIR=build

# Test flags
TEST_FLAGS=-v -race

# Main package path
MAIN_PACKAGE=.

.PHONY: all build clean test lint vet fmt run help tidy install-tools coverage

# Default target
all: tidy fmt vet lint test build

# Build the application
build:
	mkdir -p $(BUILD_DIR)
	$(GOBUILD) -o $(BUILD_DIR)/$(BINARY_NAME) $(MAIN_PACKAGE)

# Build for unix/linux
build-linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) -o $(BUILD_DIR)/$(BINARY_UNIX) $(MAIN_PACKAGE)

# Clean build files
clean:
	rm -rf $(BUILD_DIR)
	rm -f coverage.out

# Run tests
test:
	$(GOTEST) $(TEST_FLAGS) ./...

# Run tests with coverage
coverage:
	$(GOTEST) $(TEST_FLAGS) -coverprofile=coverage.out ./...
	$(GOCMD) tool cover -html=coverage.out

# Run code linting
lint:
	$(GOLANGCI_LINT) run ./...

# Run go vet
vet:
	$(GOVET) ./...

# Format code
fmt:
	$(GOFMT) ./...

# Run the application
run:
	$(GORUN) $(MAIN_PACKAGE)

# Update dependencies
tidy:
	$(GOMODTIDY)

# Install required tools
install-tools:
	@echo "Installing golangci-lint..."
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/main/install.sh | sh -s -- -b $(shell go env GOPATH)/bin v2.0.2
	@echo "Tools installed successfully"

# Show help
help:
	@echo "Available targets:"
	@echo "  all           - Run tidy, fmt, vet, lint, test, and build"
	@echo "  build         - Build the application"
	@echo "  build-linux   - Build the application for Linux"
	@echo "  clean         - Remove build artifacts"
	@echo "  test          - Run tests with race detection"
	@echo "  coverage      - Run tests with coverage report"
	@echo "  lint          - Run linter"
	@echo "  vet           - Run go vet"
	@echo "  fmt           - Format code"
	@echo "  run           - Run the application"
	@echo "  tidy          - Tidy and verify go.mod"
	@echo "  install-tools - Install required tools"