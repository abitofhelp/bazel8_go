################################################################################
# Makefile for Bazel8_Go
# 
# This Makefile provides common commands for both Bazel and Go workflows.
# It maintains a clear separation between Bazel-specific, Go-specific, and shared targets.
#
# Usage: make [target]
# Run 'make help' for a list of available targets.
################################################################################

# Declare all targets as phony
.PHONY: all build build-opt build-dbg run test test-pkg test-coverage clean clean-all \
        lint fix deps verify query doc doc-server version help \
        bazel-build bazel-build-opt bazel-build-dbg bazel-run bazel-test bazel-test-pkg \
        bazel-test-coverage bazel-clean bazel-clean-all bazel-mod-tidy bazel-gazelle \
        bazel-query-deps bazel-query-all bazel-buildifier-check bazel-buildifier-fix \
        go-build go-run go-test go-test-pkg go-test-coverage go-clean go-mod-tidy \
        go-mod-vendor go-mod-verify go-doc go-doc-server go-version go-fmt go-vet go-generate-mocks

# Tool variables
BAZEL := bazel
GO := go
MOCKGEN := mockgen

# Project variables
APP_NAME := bgj
MAIN_PKG := ./cmd
BINARY_NAME := main
GO_VERSION := 1.24.4

# Output directories
BUILD_DIR := build
COVERAGE_DIR := coverage-reports

# Default target
all: build

# Bazel commands
# Build the main binary with Bazel
bazel-build:
	$(BAZEL) build //cmd:main

# Build the main binary with Bazel in optimized mode
bazel-build-opt:
	$(BAZEL) build --compilation_mode=opt //cmd:main

# Build the main binary with Bazel in debug mode
bazel-build-dbg:
	$(BAZEL) build --compilation_mode=dbg //cmd:main

# Run the main binary with Bazel
bazel-run:
	$(BAZEL) run //cmd:main

# Run all tests with Bazel
bazel-test:
	$(BAZEL) test //...

# Run tests in the pkg directory with Bazel
bazel-test-pkg:
	$(BAZEL) test //pkg/...

# Generate test coverage report with Bazel
bazel-test-coverage:
	@echo "Checking for required tools..."
	@if ! command -v gcov >/dev/null 2>&1; then \
		echo "GCov not found. Please install GCov."; \
		exit 1; \
	fi
	@if ! command -v genhtml >/dev/null 2>&1; then \
		echo "genhtml command not found. Installing LCOV..."; \
		if [ "$$(uname)" = "Darwin" ]; then \
			brew install lcov; \
		elif [ "$$(uname)" = "Linux" ]; then \
			sudo apt-get update && sudo apt-get install -y lcov; \
		else \
			echo "Unsupported OS. Please install LCOV manually."; \
			exit 1; \
		fi; \
	fi
	@mkdir -p $(COVERAGE_DIR)
	$(BAZEL) coverage --test_output=all --combined_report=lcov //...
	genhtml -o $(COVERAGE_DIR)/bazel "$(shell $(BAZEL) info output_path)/_coverage/_coverage_report.dat"
	@echo "Coverage report generated at $(COVERAGE_DIR)/bazel/index.html"

# Clean Bazel build artifacts
bazel-clean:
	$(BAZEL) clean

# Clean all Bazel cache (more thorough)
bazel-clean-all:
	$(BAZEL) clean --expunge

# Using bzlmod for dependency management
# Update Bazel dependencies using bzlmod
bazel-mod-tidy:
	$(BAZEL) mod tidy

# Run Gazelle to update BUILD files
bazel-gazelle:
	$(BAZEL) run //:gazelle

# Show dependencies of the main binary
bazel-query-deps:
	$(BAZEL) query "deps(//cmd:main)" --output label_kind

# Show all targets in the workspace
bazel-query-all:
	$(BAZEL) query "//..." --output label_kind

# Check BUILD files with buildifier
bazel-buildifier-check:
	$(BAZEL) run //:buildifier.check

# Fix BUILD files with buildifier
bazel-buildifier-fix:
	$(BAZEL) run //:buildifier.fix

# Go commands
# Build the main binary with Go
go-build:
	@mkdir -p $(BUILD_DIR)
	$(GO) build -o $(BUILD_DIR)/$(BINARY_NAME) $(MAIN_PKG)

# Run the main binary with Go
go-run:
	$(GO) run $(MAIN_PKG)

# Run all tests with Go
go-test:
	$(GO) test ./...

# Run tests in the pkg directory with Go
go-test-pkg:
	$(GO) test ./pkg/...

# Generate test coverage report with Go
go-test-coverage:
	@mkdir -p $(COVERAGE_DIR)/go
	$(GO) test ./... -coverprofile=$(COVERAGE_DIR)/go-coverage.out
	$(GO) tool cover -html=$(COVERAGE_DIR)/go-coverage.out -o $(COVERAGE_DIR)/go/index.html
	@echo "Coverage report generated at $(COVERAGE_DIR)/go/index.html"

# Format Go code
go-fmt:
	$(GO) fmt ./...

# Run Go static analysis
go-vet:
	$(GO) vet ./...

# Generate mocks for Go interfaces
go-generate-mocks:
	@echo "Generating mocks..."
	$(GO) install github.com/golang/mock/mockgen@latest
	@find . -name "*.go" -not -path "*/vendor/*" -not -path "*/mock/*" | xargs grep -l "//go:generate mockgen" | xargs -I{} dirname {} | sort -u | xargs -I{} sh -c 'echo "Processing {}..." && cd {} && $(GO) generate'
	@echo "Mocks generated successfully"

# Clean Go build artifacts
go-clean:
	rm -f $(BINARY_NAME)
	rm -rf $(BUILD_DIR)
	rm -rf $(COVERAGE_DIR)

# Update Go dependencies
go-mod-tidy:
	$(GO) mod tidy

# Vendor Go dependencies
go-mod-vendor:
	$(GO) mod vendor

# Verify Go module integrity
go-mod-verify:
	$(GO) mod verify

# Generate and display documentation
go-doc:
	$(GO) doc -all github.com/abitofhelp/bazel8_go/...

# Start a documentation server
go-doc-server:
	@echo "Starting Go documentation server at http://localhost:6060"
	@echo "Press Ctrl+C to stop"
	$(GO) install golang.org/x/tools/cmd/godoc@latest
	godoc -http=:6060

# Display Go version information
go-version:
	@echo "Go version: $(shell $(GO) version)"
	@echo "Required Go version: $(GO_VERSION)"

# Shared commands
# Build with both Bazel and Go
build: bazel-build go-build

# Build with optimized mode
build-opt: bazel-build-opt go-build

# Build with debug mode
build-dbg: bazel-build-dbg go-build

# Run the application with Bazel
run: bazel-run

# Run tests with Bazel
test: bazel-test

# Run tests only for pkg directory with both Bazel and Go
test-pkg: bazel-test-pkg go-test-pkg

# Generate and view test coverage
test-coverage: bazel-test-coverage go-test-coverage

# Clean build artifacts
clean: bazel-clean go-clean

# Clean all build cache (more thorough)
clean-all: bazel-clean-all go-clean

# Check code style and quality
lint: bazel-buildifier-check go-fmt go-vet

# Fix code style issues
fix: bazel-buildifier-fix go-fmt

# Update dependencies
deps: bazel-mod-tidy go-mod-tidy

# Verify module integrity
verify: go-mod-verify

# Show all targets in the workspace
query: bazel-query-all

# Generate and display documentation
doc: go-doc

# Start a documentation server
doc-server: go-doc-server

# Display version information
version: go-version

# Generate mocks for testing
mocks: go-generate-mocks

# Help command
help:
	@echo "################################################################################"
	@echo "#                             Bazel8_Go Makefile                               #"
	@echo "################################################################################"
	@echo ""
	@echo "Usage:"
	@echo "  make <target>"
	@echo ""
	@echo "Shared Targets:"
	@echo "  all               Build the application (default)"
	@echo "  build             Build with both Bazel and Go"
	@echo "  build-dbg         Build with debug mode"
	@echo "  build-opt         Build with optimized mode"
	@echo "  clean             Clean build artifacts"
	@echo "  clean-all         Clean all build cache (more thorough)"
	@echo "  deps              Update dependencies with Bazel mod tidy and Go mod tidy"
	@echo "  doc               Generate and display documentation"
	@echo "  doc-server        Start a documentation server at http://localhost:6060"
	@echo "  fix               Fix code style issues (Bazel buildifier, Go fmt)"
	@echo "  lint              Check code style and quality (Bazel buildifier, Go fmt, Go vet)"
	@echo "  mocks             Generate mocks for testing"
	@echo "  query             Show all targets in the workspace"
	@echo "  run               Run the application with Bazel"
	@echo "  test              Run tests with Bazel"
	@echo "  test-coverage     Generate and view test coverage"
	@echo "  test-pkg          Run tests only for pkg directory with both Bazel and Go"
	@echo "  verify            Verify Go module integrity"
	@echo "  version           Display Go version information"
	@echo ""
	@echo "Bazel Specific Targets:"
	@echo "  bazel-build       Build with Bazel"
	@echo "  bazel-build-dbg   Build with Bazel in debug mode"
	@echo "  bazel-build-opt   Build with Bazel in optimized mode"
	@echo "  bazel-buildifier-check  Check BUILD files"
	@echo "  bazel-buildifier-fix    Fix BUILD files"
	@echo "  bazel-clean       Clean Bazel artifacts"
	@echo "  bazel-clean-all   Clean all Bazel cache (more thorough)"
	@echo "  bazel-gazelle     Run Gazelle to update BUILD files"
	@echo "  bazel-mod-tidy    Update Bazel dependencies using bzlmod"
	@echo "  bazel-query-all   Show all targets in the workspace"
	@echo "  bazel-query-deps  Show dependencies of the main binary"
	@echo "  bazel-run         Run with Bazel"
	@echo "  bazel-test        Test all targets with Bazel"
	@echo "  bazel-test-coverage Generate and view test coverage with Bazel"
	@echo "  bazel-test-pkg    Test only pkg directory with Bazel"
	@echo ""
	@echo "Go Specific Targets:"
	@echo "  go-build          Build with Go"
	@echo "  go-clean          Clean Go artifacts"
	@echo "  go-doc            Generate and display documentation"
	@echo "  go-doc-server     Start a documentation server"
	@echo "  go-fmt            Format Go code"
	@echo "  go-generate-mocks Generate mocks for Go interfaces"
	@echo "  go-mod-tidy       Tidy Go modules"
	@echo "  go-mod-vendor     Vendor Go dependencies"
	@echo "  go-mod-verify     Verify Go module integrity"
	@echo "  go-run            Run with Go"
	@echo "  go-test           Test with Go"
	@echo "  go-test-coverage  Test with coverage report"
	@echo "  go-test-pkg       Test only pkg directory with Go"
	@echo "  go-version        Display Go version information"
	@echo "  go-vet            Run Go static analysis"
