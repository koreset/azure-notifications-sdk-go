.PHONY: default all build test clean lint vet fmt check coverage help install-tools

# Default target
default: all

# Build all
all: fmt vet lint staticcheck test build

# Build the project
build:
	@echo "Building..."
	@go build -v ./...

# Run tests
test:
	@echo "Running tests..."
	@go test -v -race -coverprofile=coverage.out -covermode=atomic ./...

# Run tests with coverage
coverage:
	@echo "Running tests with coverage..."
	@go test -v -race -coverprofile=coverage.out -covermode=atomic ./...
	@go tool cover -html=coverage.out -o coverage.html
	@echo "Coverage report generated: coverage.html"

# Clean build artifacts
clean:
	@echo "Cleaning..."
	@go clean -cache -testcache
	@rm -f coverage.out coverage.html

# Format code
fmt:
	@echo "Formatting code..."
	@go fmt ./...

# Run go vet
vet:
	@echo "Running go vet..."
	@go vet ./...

# Run linter
lint:
	@echo "Running golint..."
	@$(shell go env GOPATH)/bin/golint -set_exit_status ./...

# Run staticcheck
staticcheck:
	@echo "Running staticcheck..."
	@$(shell go env GOPATH)/bin/staticcheck ./...

# Run all checks
check: fmt vet lint staticcheck

# Install development tools
install-tools:
	@echo "Installing development tools..."
	@go install golang.org/x/lint/golint@latest
	@go install honnef.co/go/tools/cmd/staticcheck@latest
	@go install golang.org/x/tools/cmd/goimports@latest

# Run pre-commit checks
pre-commit: fmt vet lint staticcheck test
	@echo "Pre-commit checks passed!"

# Show help
help:
	@echo "Available targets:"
	@echo "  build         - Build the project"
	@echo "  test          - Run tests"
	@echo "  coverage      - Run tests with coverage report"
	@echo "  clean         - Clean build artifacts"
	@echo "  fmt           - Format code"
	@echo "  vet           - Run go vet"
	@echo "  lint          - Run golint"
	@echo "  staticcheck   - Run staticcheck"
	@echo "  check         - Run all checks"
	@echo "  install-tools - Install development tools"
	@echo "  pre-commit    - Run pre-commit checks"
	@echo "  help          - Show this help"
