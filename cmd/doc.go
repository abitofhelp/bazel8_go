// Package main provides the entry point for the bazel8_go application.
//
// This application demonstrates the use of the greeting package to generate
// personalized greeting messages. It serves as an example of how to use the
// greeting package in a real application.
//
// # What This Application Does
//
// The application takes a user's name ("Mike" in this case) and generates
// a friendly greeting message. The output will be: "Howdy Mike!"
//
// # Key Components
//
// The main package:
// - Imports the greeting package from pkg/greeting
// - Imports the logger package from pkg/logger for context-aware logging
// - Sets up proper context handling with cancellation and timeout
// - Implements signal handling for graceful shutdown
// - Calls the Greet function with a name
// - Handles different types of errors that might occur
// - Prints the resulting greeting message to standard output
//
// # Usage
//
// Using Go's standard tools:
//
//	go run cmd/main.go
//
// Using Bazel:
//
//	bazel run //cmd:main
//
// # Testing
//
// Using Go's standard tools:
//
//	go test -v ./cmd
//
// Using Bazel:
//
//	bazel test //cmd:go_default_test
//
// # Build System
//
// The application is built using Bazel, which provides a consistent build
// system across different platforms and environments. It includes both unit
// and integration tests to ensure its functionality works as expected.
//
// # Error Handling
//
// The application demonstrates proper error handling techniques, including:
// - Checking for invalid inputs
// - Handling context cancellation
// - Managing timeouts
// - Providing clear error messages
package main
