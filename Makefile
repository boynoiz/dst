# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GOMOD=$(GOCMD) mod
GOTOOL=$(GOCMD) tool
GOLIST=$(GOCMD) list
GO_SOURCE=.
GOPATH=$(shell $(GOCMD) env GOPATH)

# Test variables
TEST_COVERAGE_DIR=coverage
TEST_COVERAGE_FILE=$(TEST_COVERAGE_DIR)/coverage.out
TEST_COVERAGE_HTML=$(TEST_COVERAGE_DIR)/coverage.html

# Color output
RED=\033[0;31m
GREEN=\033[0;32m
YELLOW=\033[0;33m
BLUE=\033[0;34m
NC=\033[0m # No Color

# Run tests with coverage
.PHONY: test
test:
	@echo "Running tests with coverage..."
	@mkdir -p $(TEST_COVERAGE_DIR)
	$(GOTEST) -v -race -coverprofile=$(TEST_COVERAGE_FILE) -covermode=atomic $(GO_SOURCE)/...
	@$(GOCMD) tool cover -html=$(TEST_COVERAGE_FILE) -o $(TEST_COVERAGE_HTML)
	@$(GOCMD) tool cover -func=$(TEST_COVERAGE_FILE)
	@$(MAKE) --no-print-directory update-coverage-badge
	@echo "Tests completed. Coverage report: $(TEST_COVERAGE_HTML)"

# Run unit tests only (exclude integration tests)
.PHONY: test-unit
test-unit:
	@echo "Running unit tests with coverage..."
	@mkdir -p $(TEST_COVERAGE_DIR)
	ENV=test $(GOTEST) -v -race -coverprofile=$(TEST_COVERAGE_FILE) -covermode=atomic -short $(GO_SOURCE)/...
	@$(GOCMD) tool cover -html=$(TEST_COVERAGE_FILE) -o $(TEST_COVERAGE_HTML)
	@$(GOCMD) tool cover -func=$(TEST_COVERAGE_FILE)
	@echo "----------------------------------------"
	@$(GOCMD) tool cover -func=$(TEST_COVERAGE_FILE) | grep "total:" | awk '{print "Coverage: " $$3}'
	@echo "----------------------------------------"
	@if command -v gocover-cobertura > /dev/null 2>&1; then \
		gocover-cobertura < $(TEST_COVERAGE_FILE) > $(TEST_COVERAGE_DIR)/coverage.xml; \
	fi
	@$(MAKE) --no-print-directory update-coverage-badge
	@echo "Unit tests completed. Coverage report: $(TEST_COVERAGE_HTML)"

# Update coverage badge in README.md from latest coverage.out
.PHONY: update-coverage-badge
update-coverage-badge:
	@if [ -f "$(TEST_COVERAGE_FILE)" ]; then \
		pct=$$($(GOCMD) tool cover -func=$(TEST_COVERAGE_FILE) | grep "^total:" | awk '{print $$3}' | tr -d '%'); \
		sed -i "s|coverage-[0-9.]*%25-|coverage-$${pct}%25-|g" README.md; \
		echo "Coverage badge updated: $${pct}%"; \
	fi

# Install dependencies
.PHONY: deps
deps:
	@echo "Downloading dependencies..."
	$(GOMOD) download
	@echo "Dependencies downloaded"

# Tidy go.mod and go.sum
.PHONY: tidy
tidy:
	@echo "Tidying dependencies..."
	$(GOMOD) tidy
	@echo "Dependencies tidied"

# Install the application to GOPATH/bin
.PHONY: install
install: build
	@echo "Installing $(APP_BINARY_NAME)..."
	@cp $(APP_BUILD_DIR)/$(APP_BINARY_NAME) $(GOPATH)/bin/
	@echo "Installation complete: $(GOPATH)/bin/$(APP_BINARY_NAME)"

# Run the application
.PHONY: run
run: build
	@echo "Running $(APP_BINARY_NAME)..."
	@$(APP_BUILD_DIR)/$(APP_BINARY_NAME)

# Install development tools
.PHONY: tools-install
tools-install:
	@echo "Installing development tools..."
	@echo "Installing golangci-lint"
	@$(GOCMD) install github.com/golangci/golangci-lint/v2/cmd/golangci-lint@latest

	@echo "Installing govulncheck"
	@$(GOCMD) install golang.org/x/vuln/cmd/govulncheck@latest

	@echo "Install gocover-cobertura"
	@$(GOCMD) install github.com/boumenot/gocover-cobertura@latest

	@echo "All tools installed successfully!"

# Format code
.PHONY: fmt
fmt:
	@echo "Formatting code..."
	@gofumpt -l -w . 2>/dev/null || true
	@golangci-lint fmt $(GO_SOURCE)/...
	@echo "Formatting completed"

# Lint and fix code
.PHONY: lint
lint:
	@echo "Linting code..."
	@golangci-lint run $(GO_SOURCE)/...
	@echo "Linting completed"

# Lint and fix code
.PHONY: lint-fix
lint-fix:
	@echo "Linting and fixing code..."
	@golangci-lint run --fix $(GO_SOURCE)/... | true
	@echo "Linting and fixing completed"

# Vet code
.PHONY: vet
vet:
	@echo "Vetting code..."
	@$(GOCMD) vet $(GO_SOURCE)/... | true
	@echo "Vetting completed"

# Check vulnerabilities
.PHONY: vuln-check
vuln-check:
	@echo "Go code vulnerabilities check..."
	@govulncheck $(GO_SOURCE)/... | true
	@echo "Go code vulnerabilities check completed"

# Fix code using go fix
.PHONY: fix
fix:
	@echo "Running go fix..."
	@$(GOCMD) fix $(GO_SOURCE)/...
	@echo "Fix completed"

# Complete quality assurance
.PHONY: qa
qa: tidy fmt fix lint-fix vet vuln-check test

# CI-specific quality checks (no fixing)
.PHONY: qa-ci
qa-ci: fmt-check lint vet vuln-check test-ci

# Check formatting (for CI)
.PHONY: fmt-check
fmt-check:
	@echo "Checking code formatting..."
	@golangci-lint fmt --diff $(GO_SOURCE)/...
	@echo "Code formatting is correct"


# Make sure these aren't treated as files
.PHONY: test test-unit test-all test-ci deps tidy install run tools-install fmt fmt-check lint lint-fix vet fix vuln-check qa qa-ci
