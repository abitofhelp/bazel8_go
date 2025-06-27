// Copyright (c) 2025 A Bit of Help, Inc.

// Package main provides the entry point for the application.
package main

import (
	"bytes"
	"context"
	"errors"
	"io"
	"os"
	"testing"

	"github.com/abitofhelp/bazel8_go/pkg/greeting"
	"github.com/stretchr/testify/assert"
)

// TestRun tests the run function's behavior with different greeting responses
func TestRun(t *testing.T) {
	// Save original functions and restore them after the test
	originalGreetFunc := greetFunc
	originalOsExit := osExit
	defer func() {
		greetFunc = originalGreetFunc
		osExit = originalOsExit
	}()

	tests := []struct {
		name           string
		mockGreetFunc  func(ctx context.Context, name string) (string, error)
		expectedOutput string
		expectedExit   int
	}{
		{
			name: "successful greeting",
			mockGreetFunc: func(ctx context.Context, name string) (string, error) {
				return "Hello, Mike!\n", nil
			},
			expectedOutput: "Hello, Mike!",
			expectedExit:   0,
		},
		{
			name: "invalid name error",
			mockGreetFunc: func(ctx context.Context, name string) (string, error) {
				return "", greeting.ErrInvalidName
			},
			expectedOutput: "",
			expectedExit:   1,
		},
		{
			name: "context canceled error",
			mockGreetFunc: func(ctx context.Context, name string) (string, error) {
				return "", greeting.ErrContextCanceled
			},
			expectedOutput: "",
			expectedExit:   2,
		},
		{
			name: "context deadline exceeded error",
			mockGreetFunc: func(ctx context.Context, name string) (string, error) {
				return "", greeting.ErrContextDeadlineExceeded
			},
			expectedOutput: "",
			expectedExit:   3,
		},
		{
			name: "unexpected error",
			mockGreetFunc: func(ctx context.Context, name string) (string, error) {
				return "", errors.New("unexpected error")
			},
			expectedOutput: "",
			expectedExit:   4,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Redirect stdout to capture output
			oldStdout := os.Stdout
			r, w, _ := os.Pipe()
			os.Stdout = w

			// Mock the greetFunc function
			greetFunc = tt.mockGreetFunc

			// Mock osExit to capture exit code instead of exiting
			var exitCode int
			osExit = func(code int) {
				exitCode = code
			}

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

			// Check results
			assert.Equal(t, tt.expectedExit, exitCode)
			if tt.expectedOutput != "" {
				assert.Contains(t, output, tt.expectedOutput)
			}
		})
	}
}

// TestContextCancellation tests the behavior when the context is canceled
func TestContextCancellation(t *testing.T) {
	// Save original functions and restore them after the test
	originalGreetFunc := greetFunc
	originalOsExit := osExit
	defer func() {
		greetFunc = originalGreetFunc
		osExit = originalOsExit
	}()

	// Mock greetFunc to return a context canceled error
	greetFunc = func(ctx context.Context, name string) (string, error) {
		return "", greeting.ErrContextCanceled
	}

	// Mock osExit to capture exit code instead of exiting
	var exitCode int
	osExit = func(code int) {
		exitCode = code
	}

	// Run the function
	run()

	// Check that the exit code is correct
	assert.Equal(t, 2, exitCode) // Context canceled error
}
