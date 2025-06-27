# Command-Line Applications

This directory contains the command-line applications for the Bazel8_Go project.

## Overview

The command-line applications serve as entry points to the functionality provided by the project's packages.

## Available Commands

### Main Application

The main application demonstrates the use of the greeting package and the go-humanize library.

#### Usage

##### Using Bazel

Build and run the application using Bazel:

```bash
bazel run //cmd:main
```

##### Using Go's Build System

Build and run the application using Go's native build tools:

```bash
# Run directly without creating an executable
go run cmd/main.go

# Or build an executable and then run it
go build -o main cmd/main.go
./main
```

Both methods will output a greeting message and a formatted large number.

## Implementation Details

The main application:

1. Imports the greeting package from the project
2. Uses the Greet function to generate a greeting message
3. Uses the go-humanize package to format a large number
4. Prints both to the console

## Dependencies

- `github.com/dustin/go-humanize` - For formatting large numbers
- `github.com/abitofhelp/bazel8_go/pkg/greeting` - For generating greeting messages
- `github.com/stretchr/testify/assert` - For assertions in tests

When using Go's build system, these dependencies are managed through the go.mod file at the root of the project.

## Testing

The command-line application includes both unit and integration tests to ensure its functionality works as expected.

### Unit Tests

Unit tests verify the behavior of individual components in isolation, using mocks for external dependencies.

Run the unit tests:

```bash
# Using Go's testing tools
go test -v ./cmd

# Using Bazel
bazel test //cmd:go_default_test
```

### Integration Tests

Integration tests verify the interaction between the command-line application and its dependencies, such as the greeting package.

Run the integration tests:

```bash
# Using Go's testing tools
go test -v ./cmd -tags=integration

# Using Bazel
bazel test //cmd:go_default_test
```

### Test Coverage

The tests aim to achieve at least 80% code coverage. You can check the coverage with:

```bash
go test -cover ./cmd
```
