// Copyright (c) 2025 A Bit of Help, Inc.

// Package main is the entry point for the bazel8_go application.
// See doc.go for detailed package documentation.
package main

import (
	"context"
	"errors"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/abitofhelp/bazel8_go/pkg/greeting"
	"github.com/abitofhelp/bazel8_go/pkg/logger"
)

// Variables to allow mocking in tests.
// These variables are defined at package level to facilitate unit testing
// by allowing test code to replace them with mock implementations.
var (
	// greetFunc is a function variable that wraps greeting.Greet.
	// In production code, it calls the actual greeting.Greet function.
	// In tests, it can be replaced with a mock implementation to avoid
	// dependencies on the real greeting package.
	greetFunc = greeting.Greet

	// osExit is a function variable that wraps os.Exit.
	// This allows tests to prevent the program from actually exiting
	// when testing error scenarios.
	osExit = os.Exit

	// signalChan is the channel used for signal handling.
	// It receives OS signals like SIGINT (Ctrl+C) and SIGTERM
	// to enable graceful shutdown of the application.
	signalChan = make(chan os.Signal, 1)
)

// main initializes and runs the application.
// This function serves as the entry point for the application when executed.
// It delegates all work to the run() function, which contains the actual
// application logic. This separation allows the run() function to be tested
// without actually starting the application.
//
// The main function doesn't return any values and doesn't take any parameters.
// When the application completes successfully, it exits with status code 0.
// If an error occurs, it exits with a non-zero status code.
func main() {
	run()
}

// run contains the main logic of the application, extracted for testability.
// This function:
// 1. Sets up context with cancellation for proper resource management
// 2. Configures signal handling to enable graceful shutdown
// 3. Creates a timeout context to prevent hanging operations
// 4. Calls the greeting function with a name and winning amount
// 5. Handles different types of errors that might occur
// 6. Prints the resulting greeting message to standard output
//
// By extracting this logic from main(), we can unit test it without
// actually running the application, which makes testing more reliable.
func run() {
	// Create a context with cancellation
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Set up signal handling
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)

	// Handle signals in a separate goroutine
	go func() {
		sig := <-signalChan
		logger.DefaultLogger.Info(ctx, "Received signal: %v", sig)
		logger.DefaultLogger.Info(ctx, "Shutting down gracefully...")
		cancel() // Cancel the context
	}()

	// Create a context with timeout
	timeoutCtx, timeoutCancel := context.WithTimeout(ctx, 5*time.Second)
	defer timeoutCancel()

	// Generate greeting with proper error handling
	message, err := greetFunc(timeoutCtx, "Mike")

	// Handle different types of errors
	if err != nil {
		if errors.Is(err, greeting.ErrInvalidName) {
			logger.DefaultLogger.Error(ctx, "Invalid name provided: %v", err)
			osExit(1)
		} else if errors.Is(err, greeting.ErrContextCanceled) {
			logger.DefaultLogger.Warning(ctx, "Operation was canceled: %v", err)
			osExit(2)
		} else if errors.Is(err, greeting.ErrContextDeadlineExceeded) {
			logger.DefaultLogger.Warning(ctx, "Operation timed out: %v", err)
			osExit(3)
		} else {
			logger.DefaultLogger.Error(ctx, "Unexpected error: %v", err)
			osExit(4)
		}
	}

	// Print the greeting message
	fmt.Println(message)
}
