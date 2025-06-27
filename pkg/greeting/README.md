# Greeting Package

This package provides greeting functionality for the Bazel8_Go project.

## Overview

The greeting package offers a simple way to generate personalized greeting messages. It's designed to be easy to use and integrate into other applications.

## Usage

Here's how to use the greeting package:

```
// In your Go file:
package main

import (
    "fmt"
    "github.com/abitofhelp/bazel8_go/pkg/greeting"
)

func main() {
    // Get a greeting for a specific name
    message := greeting.Greet("John")
    fmt.Println(message) // Outputs: Hello, John!

    // Get a default greeting
    message = greeting.Greet("")
    fmt.Println(message) // Outputs: Hello, World!
}
```

## API Reference

### Functions

#### `Greet(name string) string`

Returns a greeting message for the given name.

**Parameters:**
- `name` (string): The name to include in the greeting. If empty, defaults to "World".

**Returns:**
- (string): A greeting message in the format "Hello, {name}!"

## Implementation Details

The package implements a simple greeting function that:
1. Takes a name parameter
2. Checks if the name is empty, and if so, uses "World" as the default
3. Returns a formatted greeting message

## Testing

The package includes unit tests to verify the functionality:

```bash
bazel test //pkg/greeting:greeting_test
```
