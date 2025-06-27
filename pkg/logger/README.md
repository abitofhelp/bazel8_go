# Logger Package

## Overview

The `logger` package provides logging utilities for the application, with a focus on context-aware logging. It implements a `ContextLogger` that ensures context information is properly included in each log entry.

## Features

- Context-aware logging
- Multiple log levels (Info, Warning, Error, Fatal)
- Default logger instance for easy use throughout the application
- Customizable underlying logger

## Usage

See the godoc documentation for detailed usage examples.

## Testing

The package includes tests to verify the functionality of the `ContextLogger`. Run the tests using:

```
go test -v ./pkg/logger
```

or with Bazel:

```
bazel test //pkg/logger:go_default_test
```