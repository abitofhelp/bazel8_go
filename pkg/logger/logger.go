// Package logger provides logging utilities for the application.
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
// It provides methods for logging at different levels (info, warning, error)
// and ensures that context information is properly included in each log entry.
type ContextLogger struct {
	logger *log.Logger
}

// NewContextLogger creates a new ContextLogger with the provided logger.
// If logger is nil, it uses the default logger.
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
func (l *ContextLogger) Info(ctx context.Context, format string, v ...interface{}) {
	contextInfo := extractContextInfo(ctx)
	l.logger.Printf("INFO: %s"+format, append([]interface{}{contextInfo}, v...)...)
}

// Warning logs a warning message with context information.
func (l *ContextLogger) Warning(ctx context.Context, format string, v ...interface{}) {
	contextInfo := extractContextInfo(ctx)
	l.logger.Printf("WARNING: %s"+format, append([]interface{}{contextInfo}, v...)...)
}

// Error logs an error message with context information.
func (l *ContextLogger) Error(ctx context.Context, format string, v ...interface{}) {
	contextInfo := extractContextInfo(ctx)
	l.logger.Printf("ERROR: %s"+format, append([]interface{}{contextInfo}, v...)...)
}

// Fatal logs a fatal error message with context information and then exits the program.
func (l *ContextLogger) Fatal(ctx context.Context, format string, v ...interface{}) {
	contextInfo := extractContextInfo(ctx)
	l.logger.Fatalf("FATAL: %s"+format, append([]interface{}{contextInfo}, v...)...)
}

// DefaultLogger is a singleton instance of ContextLogger that can be used throughout the application.
var DefaultLogger = NewContextLogger(nil)

// WithRequestID returns a new context with the given request ID.
func WithRequestID(ctx context.Context, requestID string) context.Context {
	return context.WithValue(ctx, RequestIDKey, requestID)
}

// WithUserID returns a new context with the given user ID.
func WithUserID(ctx context.Context, userID string) context.Context {
	return context.WithValue(ctx, UserIDKey, userID)
}
