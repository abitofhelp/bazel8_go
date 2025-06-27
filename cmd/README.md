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

When using Go's build system, these dependencies are managed through the go.mod file at the root of the project.
