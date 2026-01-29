.PHONY: help build test test-coverage lint clean install run fmt

# Variables
BINARY_NAME=skills-validate
CMD_PATH=./cmd/skills-validate
INSTALL_PATH=$(shell go env GOPATH)/bin
COVERAGE_FILE=coverage.out

help: ## Display this help message
	@echo "Available commands:"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2}'

build: ## Build the binary
	@echo "Building $(BINARY_NAME)..."
	go build -o $(BINARY_NAME) $(CMD_PATH)

install: ## Install the binary to GOPATH/bin
	@echo "Installing $(BINARY_NAME) to $(INSTALL_PATH)..."
	go install $(CMD_PATH)

test: ## Run all tests
	@echo "Running tests..."
	go test -v ./...

test-coverage: ## Run tests with coverage report
	@echo "Running tests with coverage..."
	go test -coverprofile=$(COVERAGE_FILE) ./...
	go tool cover -html=$(COVERAGE_FILE)

lint: ## Run linters
	@echo "Running linters..."
	@if ! command -v golangci-lint &> /dev/null; then \
		echo "golangci-lint not found. Install it with: brew install golangci-lint"; \
		exit 1; \
	fi
	golangci-lint run ./...

fmt: ## Format code
	@echo "Formatting code..."
	go fmt ./...
	@if command -v goimports &> /dev/null; then \
		goimports -w .; \
	else \
		echo "goimports not found. Run: go install golang.org/x/tools/cmd/goimports@latest"; \
	fi

run: build ## Build and run the binary
	./$(BINARY_NAME)

clean: ## Clean build artifacts
	@echo "Cleaning..."
	rm -f $(BINARY_NAME)
	rm -f $(COVERAGE_FILE)
	rm -rf dist/
	go clean

deps: ## Download dependencies
	@echo "Downloading dependencies..."
	go mod download
	go mod tidy

.DEFAULT_GOAL := help
