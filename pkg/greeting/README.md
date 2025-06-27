# Greeting Package

This package provides greeting functionality for the Bazel8_Go project.

## Overview

The greeting package offers a way to generate personalized greeting messages. It's designed to be context-aware, supporting cancellation and timeouts, and provides comprehensive error handling.

## Features

- Context-aware operations with support for cancellation and timeouts
- Personalized greeting messages
- Comprehensive error handling with custom error types
- Input validation

## Usage

Here's how to use the greeting package:

```go
// In your Go file:
package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/abitofhelp/bazel8_go/pkg/greeting" // Import the greeting package
)

func main() {
	// Create a context
	ctx := context.Background()

	// Get a greeting for a specific name
	message, err := greeting.Greet(ctx, "John")
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	fmt.Println(message) // Outputs: Howdy John!

	// Using a context with timeout
	timeoutCtx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	message, err = greeting.Greet(timeoutCtx, "Alice")
	if err != nil {
		if errors.Is(err, greeting.ErrContextDeadlineExceeded) {
			log.Println("Operation timed out")
		} else if errors.Is(err, greeting.ErrInvalidName) {
			log.Println("Invalid name provided")
		} else {
			log.Fatalf("Unexpected error: %v", err)
		}
		return
	}
	fmt.Println(message)
}
```

## API Reference

### Functions

#### `Greet(ctx context.Context, name string) (string, error)`

Returns a personalized greeting message for the given name.

**Parameters:**
- `ctx` (context.Context): The context for the operation, supporting cancellation and timeouts.
- `name` (string): The name to include in the greeting. Cannot be empty.

**Returns:**
- (string): A greeting message in the format "Howdy {name}!"
- (error): An error if the operation failed. Possible errors include:
  - `ErrInvalidName`: If the name is empty.
  - `ErrContextCanceled`: If the context was canceled during processing.
  - `ErrContextDeadlineExceeded`: If the context deadline was exceeded.

### Error Types

- `ErrInvalidName`: Returned when the provided name is empty.
- `ErrContextCanceled`: Returned when the context is canceled during processing.
- `ErrContextDeadlineExceeded`: Returned when the context deadline is exceeded during processing.

## Implementation Details

The package implements a greeting function that:
1. Takes a context and name as parameters
2. Validates the input parameters (name cannot be empty)
3. Handles context cancellation and timeouts
4. Returns a formatted greeting message and any error that occurred

## Testing

The package includes comprehensive unit tests to verify the functionality, including edge cases and error handling:

```bash
bazel test //pkg/greeting:greeting_test
```

Or using Go's native testing tools:

```bash
go test -v ./pkg/greeting
```
