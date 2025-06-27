// Package logger provides logging utilities for the application.
package logger

import (
	"context"
	"log"
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

// Info logs an informational message with context information.
func (l *ContextLogger) Info(ctx context.Context, format string, v ...interface{}) {
	l.logger.Printf("INFO: "+format, v...)
}

// Warning logs a warning message with context information.
func (l *ContextLogger) Warning(ctx context.Context, format string, v ...interface{}) {
	l.logger.Printf("WARNING: "+format, v...)
}

// Error logs an error message with context information.
func (l *ContextLogger) Error(ctx context.Context, format string, v ...interface{}) {
	l.logger.Printf("ERROR: "+format, v...)
}

// Fatal logs a fatal error message with context information and then exits the program.
func (l *ContextLogger) Fatal(ctx context.Context, format string, v ...interface{}) {
	l.logger.Fatalf("FATAL: "+format, v...)
}

// DefaultLogger is a singleton instance of ContextLogger that can be used throughout the application.
var DefaultLogger = NewContextLogger(nil)