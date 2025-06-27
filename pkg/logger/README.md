# Logger Package

## Overview

The `logger` package provides logging utilities for the application, with a focus on context-aware logging. It implements a `ContextLogger` that ensures context information is properly included in each log entry, making it easier to trace and debug issues across different parts of the application.

## Features

- Context-aware logging that extracts and includes relevant context information
- Multiple log levels (Info, Warning, Error, Fatal)
- Default logger instance for easy use throughout the application
- Customizable underlying logger
- Helper functions for adding common values to context
- Support for request IDs and user IDs in logs

## Usage

### Basic Usage

```go
package main

import (
	"context"

	"github.com/abitofhelp/bazel8_go/pkg/logger"
)

func main() {
	// Create a context
	ctx := context.Background()

	// Log a message with the default logger
	logger.DefaultLogger.Info(ctx, "Application started")

	// Log a message with formatting
	logger.DefaultLogger.Info(ctx, "Processing item %d", 123)

	// Log a warning
	logger.DefaultLogger.Warning(ctx, "Resource usage is high: %d%%", 85)

	// Log an error
	logger.DefaultLogger.Error(ctx, "Failed to process item: %v", err)
}
```

### Adding Context Information

The logger extracts and includes context information in log messages. You can add values to the context using the provided helper functions:

```go
// Add a request ID to the context
ctx = logger.WithRequestID(ctx, "req-123")

// Add a user ID to the context
ctx = logger.WithUserID(ctx, "user-456")

// Log with context information
logger.DefaultLogger.Info(ctx, "Processing request")
// Output: INFO: [request_id=req-123 user_id=user-456] Processing request
```

### Custom Logger

You can create a custom logger with your own underlying logger:

```go
// Create a custom logger with specific flags
customLogger := log.New(os.Stdout, "CUSTOM: ", log.LstdFlags|log.Lshortfile)
ctxLogger := logger.NewContextLogger(customLogger)

// Use the custom logger
ctxLogger.Info(ctx, "Using custom logger")
```

## API Reference

### Types

#### `ContextLogger`

A logger that includes context information in log messages.

### Functions

#### `NewContextLogger(logger *log.Logger) *ContextLogger`

Creates a new ContextLogger with the provided logger. If logger is nil, it uses the default logger.

#### `WithRequestID(ctx context.Context, requestID string) context.Context`

Returns a new context with the given request ID.

#### `WithUserID(ctx context.Context, userID string) context.Context`

Returns a new context with the given user ID.

### Methods

#### `Info(ctx context.Context, format string, v ...interface{})`

Logs an informational message with context information.

#### `Warning(ctx context.Context, format string, v ...interface{})`

Logs a warning message with context information.

#### `Error(ctx context.Context, format string, v ...interface{})`

Logs an error message with context information.

#### `Fatal(ctx context.Context, format string, v ...interface{})`

Logs a fatal error message with context information and then exits the program.

### Variables

#### `DefaultLogger`

A singleton instance of ContextLogger that can be used throughout the application.

## Implementation Details

The logger extracts information from the context using predefined keys:

- `RequestIDKey`: Used to store and retrieve request IDs
- `UserIDKey`: Used to store and retrieve user IDs

When logging a message, the logger checks if these values are present in the context and includes them in the log message in the format `[key1=value1 key2=value2] message`.

## Testing

The package includes comprehensive tests to verify the functionality of the `ContextLogger`, including context value extraction and formatting. Run the tests using:

```bash
go test -v ./pkg/logger
```

or with Bazel:

```bash
bazel test //pkg/logger:go_default_test
```
