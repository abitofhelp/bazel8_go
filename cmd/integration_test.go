// Copyright (c) 2025 A Bit of Help, Inc.

// Package main provides the entry point for the application.
package main

import (
	"bytes"
	"context"
	"errors"
	"io"
	"log"
	"os"
	"testing"

	"github.com/abitofhelp/bazel8_go/pkg/greeting"
	"github.com/abitofhelp/bazel8_go/pkg/logger"
	"github.com/stretchr/testify/assert"
)

// TestIntegrationWithGreeting tests the integration between the main package and the greeting package.
// This is an integration test that verifies the actual greeting.Greet function is called correctly.
func TestIntegrationWithGreeting(t *testing.T) {
	// Save original functions and restore them after the test
	originalGreetFunc := greetFunc
	originalOsExit := osExit
	defer func() {
		greetFunc = originalGreetFunc
		osExit = originalOsExit
	}()

	// Use the actual greeting.Greet function
	greetFunc = greeting.Greet

	// Mock osExit to prevent the test from exiting
	osExit = func(code int) {
		// Do nothing
	}

	// Redirect stdout to capture output
	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	// Run the function
	run()

	// Restore stdout
	w.Close()
	os.Stdout = oldStdout

	// Read captured output
	var buf bytes.Buffer
	_, err := io.Copy(&buf, r)
	assert.NoError(t, err)
	output := buf.String()

	// Verify the output contains the expected greeting
	assert.Contains(t, output, "Howdy Mike!")
}

// TestIntegrationWithLogger tests the integration between the main package and the logger package.
// This is an integration test that verifies the logger is used correctly when an error occurs.
func TestIntegrationWithLogger(t *testing.T) {
	// Save original functions and restore them after the test
	originalGreetFunc := greetFunc
	originalOsExit := osExit
	originalLogger := logger.DefaultLogger
	defer func() {
		greetFunc = originalGreetFunc
		osExit = originalOsExit
		logger.DefaultLogger = originalLogger
	}()

	// Create a custom logger for testing
	var logOutput bytes.Buffer
	customLogger := logger.NewContextLogger(log.New(&logOutput, "", 0))
	logger.DefaultLogger = customLogger

	// Mock greetFunc to return an error
	greetFunc = func(ctx context.Context, name string) (string, error) {
		return "", greeting.ErrInvalidName
	}

	// Mock osExit to prevent the test from exiting
	var exitCode int
	osExit = func(code int) {
		exitCode = code
	}

	// Run the function
	run()

	// Check that the exit code is correct
	assert.Equal(t, 1, exitCode)

	// Verify that the logger was used to log the error
	assert.Contains(t, logOutput.String(), "Invalid name provided")
}

// TestIntegrationWithContextTimeout tests the behavior when a context timeout occurs.
// This is an integration test that verifies the application handles context timeouts correctly.
func TestIntegrationWithContextTimeout(t *testing.T) {
	// Save original functions and restore them after the test
	originalGreetFunc := greetFunc
	originalOsExit := osExit
	defer func() {
		greetFunc = originalGreetFunc
		osExit = originalOsExit
	}()

	// Mock greetFunc to simulate a timeout
	greetFunc = func(ctx context.Context, name string) (string, error) {
		// Wait for the context to be canceled or timeout
		<-ctx.Done()
		return "", errors.New("context deadline exceeded: " + ctx.Err().Error())
	}

	// Mock osExit to capture the exit code
	var exitCode int
	osExit = func(code int) {
		exitCode = code
	}

	// Run the function
	run()

	// Check that the exit code is correct
	assert.Equal(t, 4, exitCode) // Unexpected error
}
