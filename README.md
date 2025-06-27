# Bazel8_Go

An example application demonstrating the use of both Go and Bazel8 build systems.

## Overview

This project showcases how to set up and use both Go's native build system and Bazel8 for building and testing Go applications. It includes a simple greeting package and a command-line application that uses this package. The Go's go.mod file is the source of truth for dependencies in Bazel.

## Features

- Go's native build system integration
- Bazel8 build system integration
- Simple greeting functionality
- Code coverage reporting
- Clean architecture principles

## Getting Started

### Prerequisites

- Go 1.24 or later
- Bazel

### Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/abitofhelp/bazel8_go.git
   cd bazel8_go
   ```

2. Build and run the project using Bazel:
   ```bash
   # Build the project
   bazel build //...

   # Run the application
   bazel run //cmd:main
   ```

3. Alternatively, build and run using Go's native build tools:
   ```bash
   # Run directly without creating an executable
   go run cmd/main.go

   # Or build an executable and then run it
   go build -o main cmd/main.go
   ./main
   ```

## Using the Makefile

This project includes a comprehensive Makefile that simplifies common development tasks. It's recommended to use the Makefile for most operations as it provides a consistent interface for both Bazel and Go workflows.

```bash
# Display all available commands
make help

# Build the project with both Bazel and Go
make build

# Run the application
make run

# Run all tests
make test

# Generate test coverage reports
make test-coverage

# Format and lint code
make lint

# Fix code style issues
make fix

# Clean build artifacts
make clean
```

The Makefile includes many more commands for specific tasks. Run `make help` to see the complete list of available commands.

## Project Structure

- [`/cmd`](./cmd/README.md) - Command-line applications
- [`/pkg`](./pkg/README.md) - Reusable packages
- [`/DOCS`](./DOCS/README.md) - Project documentation
- [`/tools`](./tools/README.md) - Utility scripts

## Documentation

For more detailed documentation, see the [DOCS](./DOCS/README.md) directory.

## Testing

### Using Bazel

Run the tests with Bazel:

```bash
bazel test //...
```

Generate code coverage reports:

```bash
./tools/run_coverage.sh
```

### Using Go's Native Testing Tools

Run the tests with Go's testing tools:

```bash
go test ./...
```

Generate code coverage reports:

```bash
go test ./... -coverprofile=coverage.out
go tool cover -html=coverage.out
```

For more information on code coverage, see the [coverage documentation](./DOCS/COVERAGE.md).

## License

This project is licensed under the terms specified in the [LICENSE](./LICENSE) file.
