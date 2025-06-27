// Package logger provides logging utilities for the application.
// See doc.go for detailed package documentation.
package logger

import (
	"context"
	"fmt"
	"log"
)

// contextKey is a type for context keys to avoid collisions.
type contextKey string

// Context keys for storing and retrieving values from context.
const (
	// RequestIDKey is the key for request ID in context.
	RequestIDKey contextKey = "request_id"
	// UserIDKey is the key for user ID in context.
	UserIDKey contextKey = "user_id"
)

// ContextLogger is a logger that includes context information in log messages.
// It enhances the standard Go logger by automatically extracting and including
// relevant information from the context in each log message.
//
// # Features
//
// - Automatically includes context information in log messages
// - Provides methods for different log levels (Info, Warning, Error, Fatal)
// - Uses the standard Go logger underneath for actual logging
// - Thread-safe for concurrent use by multiple goroutines
//
// # Usage
//
// ContextLogger should be used instead of the standard log package when you want
// to include context information in your log messages. You can either use the
// DefaultLogger instance or create your own with NewContextLogger.
//
// See the package documentation for examples and more information.
type ContextLogger struct {
	// logger is the underlying standard Go logger used for actual logging
	logger *log.Logger
}

// NewContextLogger creates a new ContextLogger with the provided logger.
//
// # Parameters
//
// - logger: A pointer to a standard Go logger that will be used for actual logging.
//   If nil is provided, the function will use log.Default() instead.
//
// # Return Value
//
// - *ContextLogger: A new ContextLogger instance that wraps the provided logger
//   and adds context-aware logging capabilities.
//
// # Example
//
//	// Create a logger that writes to standard output with a custom prefix
//	stdLogger := log.New(os.Stdout, "APP: ", log.LstdFlags)
//	contextLogger := logger.NewContextLogger(stdLogger)
//
//	// Use the default logger
//	defaultContextLogger := logger.NewContextLogger(nil)
func NewContextLogger(logger *log.Logger) *ContextLogger {
	if logger == nil {
		logger = log.Default()
	}
	return &ContextLogger{
		logger: logger,
	}
}

// extractContextInfo extracts relevant information from the context and formats it.
func extractContextInfo(ctx context.Context) string {
	if ctx == nil {
		return ""
	}

	var contextInfo string

	// Extract request ID if present
	if requestID, ok := ctx.Value(RequestIDKey).(string); ok && requestID != "" {
		contextInfo += fmt.Sprintf("request_id=%s ", requestID)
	}

	// Extract user ID if present
	if userID, ok := ctx.Value(UserIDKey).(string); ok && userID != "" {
		contextInfo += fmt.Sprintf("user_id=%s ", userID)
	}

	// Add more context extractions as needed

	if contextInfo != "" {
		return "[" + contextInfo[:len(contextInfo)-1] + "] " // Remove trailing space and add brackets
	}
	return ""
}

// Info logs an informational message with context information.
//
// Use Info for general information about application progress and normal operations.
// These messages are helpful for understanding what the application is doing during
// normal operation.
//
// # Parameters
//
// - ctx: A context.Context that may contain values like request ID or user ID
//   that will be automatically included in the log message.
//
// - format: A format string, similar to fmt.Printf, that specifies the message format.
//
// - v: Optional arguments to be formatted according to the format string.
//
// # Example
//
//	ctx := context.Background()
//	logger.DefaultLogger.Info(ctx, "Processing item %d", itemID)
//	// Output: INFO: Processing item 123
//
//	// With context information
//	ctx = logger.WithRequestID(ctx, "req-123")
//	logger.DefaultLogger.Info(ctx, "Processing item %d", itemID)
//	// Output: INFO: [request_id=req-123] Processing item 123
func (l *ContextLogger) Info(ctx context.Context, format string, v ...interface{}) {
	contextInfo := extractContextInfo(ctx)
	l.logger.Printf("INFO: %s"+format, append([]interface{}{contextInfo}, v...)...)
}

// Warning logs a warning message with context information.
//
// Use Warning for unusual or unexpected situations that don't prevent the application
// from functioning but might indicate a problem. These messages indicate that something
// unusual happened, but the application can continue.
//
// # Parameters
//
// - ctx: A context.Context that may contain values like request ID or user ID
//   that will be automatically included in the log message.
//
// - format: A format string, similar to fmt.Printf, that specifies the message format.
//
// - v: Optional arguments to be formatted according to the format string.
//
// # Example
//
//	ctx := context.Background()
//	logger.DefaultLogger.Warning(ctx, "Unusual condition detected: %v", condition)
//	// Output: WARNING: Unusual condition detected: value out of range
func (l *ContextLogger) Warning(ctx context.Context, format string, v ...interface{}) {
	contextInfo := extractContextInfo(ctx)
	l.logger.Printf("WARNING: %s"+format, append([]interface{}{contextInfo}, v...)...)
}

// Error logs an error message with context information.
//
// Use Error for errors that prevent a specific operation from completing successfully
// but don't cause the application to stop. These messages indicate that something
// went wrong and the current operation failed.
//
// # Parameters
//
// - ctx: A context.Context that may contain values like request ID or user ID
//   that will be automatically included in the log message.
//
// - format: A format string, similar to fmt.Printf, that specifies the message format.
//
// - v: Optional arguments to be formatted according to the format string.
//
// # Example
//
//	ctx := context.Background()
//	logger.DefaultLogger.Error(ctx, "Failed to process item %d: %v", itemID, err)
//	// Output: ERROR: Failed to process item 123: file not found
func (l *ContextLogger) Error(ctx context.Context, format string, v ...interface{}) {
	contextInfo := extractContextInfo(ctx)
	l.logger.Printf("ERROR: %s"+format, append([]interface{}{contextInfo}, v...)...)
}

// Fatal logs a fatal error message with context information and then exits the program.
//
// Use Fatal for critical errors that prevent the application from continuing.
// After logging the message, this method will call os.Exit(1) to terminate the program.
//
// # Parameters
//
// - ctx: A context.Context that may contain values like request ID or user ID
//   that will be automatically included in the log message.
//
// - format: A format string, similar to fmt.Printf, that specifies the message format.
//
// - v: Optional arguments to be formatted according to the format string.
//
// # Example
//
//	ctx := context.Background()
//	logger.DefaultLogger.Fatal(ctx, "Configuration file not found: %s", configPath)
//	// Output: FATAL: Configuration file not found: /etc/app/config.json
//	// (Program will exit after this message)
//
// # Important Note
//
// This method will terminate the program. Use it only for errors that make it
// impossible for the application to continue running.
func (l *ContextLogger) Fatal(ctx context.Context, format string, v ...interface{}) {
	contextInfo := extractContextInfo(ctx)
	l.logger.Fatalf("FATAL: %s"+format, append([]interface{}{contextInfo}, v...)...)
}

// DefaultLogger is a singleton instance of ContextLogger that can be used throughout the application.
// This pre-configured logger is ready to use without any additional setup, making it
// convenient for quick logging needs.
//
// # Usage
//
//	import "github.com/abitofhelp/bazel8_go/pkg/logger"
//
//	func doSomething(ctx context.Context) {
//	    logger.DefaultLogger.Info(ctx, "Starting operation")
//	    // ... do work ...
//	    logger.DefaultLogger.Info(ctx, "Operation completed")
//	}
var DefaultLogger = NewContextLogger(nil)

// WithRequestID returns a new context with the given request ID.
//
// Request IDs are useful for tracking a request through different parts of an application,
// especially in distributed systems or microservices. When you log messages using a context
// with a request ID, the ID will be automatically included in the log messages.
//
// # Parameters
//
// - ctx: The parent context to which the request ID will be added.
//
// - requestID: A string identifier for the request, typically a unique ID like "req-123".
//
// # Return Value
//
// - context.Context: A new context that contains the request ID.
//
// # Example
//
//	// Add a request ID to the context
//	ctx = logger.WithRequestID(ctx, "req-123")
//
//	// Log with the context
//	logger.DefaultLogger.Info(ctx, "Processing request")
//	// Output: INFO: [request_id=req-123] Processing request
func WithRequestID(ctx context.Context, requestID string) context.Context {
	return context.WithValue(ctx, RequestIDKey, requestID)
}

// WithUserID returns a new context with the given user ID.
//
// User IDs are useful for tracking which user is associated with a particular operation
// or request. When you log messages using a context with a user ID, the ID will be
// automatically included in the log messages.
//
// # Parameters
//
// - ctx: The parent context to which the user ID will be added.
//
// - userID: A string identifier for the user, typically a unique ID like "user-456".
//
// # Return Value
//
// - context.Context: A new context that contains the user ID.
//
// # Example
//
//	// Add a user ID to the context
//	ctx = logger.WithUserID(ctx, "user-456")
//
//	// Log with the context
//	logger.DefaultLogger.Info(ctx, "User action performed")
//	// Output: INFO: [user_id=user-456] User action performed
func WithUserID(ctx context.Context, userID string) context.Context {
	return context.WithValue(ctx, UserIDKey, userID)
}
