// Copyright (c) 2025 A Bit of Help, Inc.

// Package main is the entry point for the application.
// It demonstrates the use of the greeting package to generate
// personalized greeting messages with formatted monetary amounts.
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

// main initializes and runs the application.
// It creates a greeting message for "Mike" with a winning amount
// and prints it to standard output.
// It also demonstrates proper context handling and signal handling.
// If an error occurs during greeting generation, it logs the error and exits.
func main() {
	// Create a context with cancellation
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Set up signal handling
	signalChan := make(chan os.Signal, 1)
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
	message, err := greeting.Greet(timeoutCtx, "Mike", 1234567890)

	// Handle different types of errors
	if err != nil {
		if errors.Is(err, greeting.ErrInvalidName) {
			logger.DefaultLogger.Error(ctx, "Invalid name provided: %v", err)
			os.Exit(1)
		} else if errors.Is(err, greeting.ErrContextCanceled) {
			logger.DefaultLogger.Warning(ctx, "Operation was canceled: %v", err)
			os.Exit(2)
		} else if errors.Is(err, greeting.ErrContextDeadlineExceeded) {
			logger.DefaultLogger.Warning(ctx, "Operation timed out: %v", err)
			os.Exit(3)
		} else {
			logger.DefaultLogger.Error(ctx, "Unexpected error: %v", err)
			os.Exit(4)
		}
	}

	// Print the greeting message
	fmt.Println(message)
}
