// Package logger provides logging utilities for the application.
//
// # Overview
//
// This package offers a context-aware logging system that ensures all log messages
// include relevant context information. It provides a ContextLogger type with methods
// for logging at different severity levels (info, warning, error, fatal).
//
// The package also provides a DefaultLogger singleton instance that can be used
// throughout the application without needing to create a new logger instance.
//
// # Key Features
//
// - Context-aware logging that automatically includes context information
// - Multiple log levels for different types of messages
// - Support for formatted log messages (similar to fmt.Printf)
// - Easy creation of custom loggers with different configurations
// - Helper functions to add values to the context for logging
//
// # Log Levels
//
// The logger supports the following log levels, in order of increasing severity:
//
// 1. Info: For general information about application progress and normal operations.
//    Use this for messages that are helpful for understanding what the application is doing.
//    Example: "Processing item 123", "Request completed successfully"
//
// 2. Warning: For unusual or unexpected situations that don't prevent the application
//    from functioning but might indicate a problem.
//    Example: "Unusual condition detected", "Using fallback method"
//
// 3. Error: For errors that prevent a specific operation from completing successfully
//    but don't cause the application to stop.
//    Example: "Failed to process item 123", "Database connection failed"
//
// 4. Fatal: For critical errors that prevent the application from continuing.
//    This level will terminate the application after logging the message.
//    Example: "Configuration file not found", "Required service unavailable"
//
// # Log Message Format
//
// Log messages follow this format:
//
//	LEVEL: [context_info] message
//
// Where:
// - LEVEL is the log level (INFO, WARNING, ERROR, FATAL)
// - context_info is information extracted from the context (if available)
// - message is the actual log message
//
// Example:
//
//	INFO: [request_id=req-123 user_id=user-456] Processing request
//
// # Basic Usage
//
//	ctx := context.Background()
//	logger.DefaultLogger.Info(ctx, "Processing item %d", itemID)
//
//	// Or create a custom logger
//	customLogger := logger.NewContextLogger(log.New(os.Stdout, "CUSTOM: ", log.LstdFlags))
//	customLogger.Warning(ctx, "Unusual condition detected: %v", condition)
//
// # Adding Context Values
//
// You can add values to the context that will be automatically included in log messages:
//
//	// Add a request ID to the context
//	ctx = logger.WithRequestID(ctx, "req-123")
//
//	// Add a user ID to the context
//	ctx = logger.WithUserID(ctx, "user-456")
//
//	// Log with context information
//	logger.DefaultLogger.Info(ctx, "Processing request")
//	// Output: INFO: [request_id=req-123 user_id=user-456] Processing request
//
// # Context Convention
//
// All logging methods require a context.Context as their first parameter, following
// the project's convention that any function requiring context must have it as the
// first parameter.
//
// # Thread Safety
//
// The logger is safe for concurrent use by multiple goroutines. You can use the
// DefaultLogger from different parts of your application without worrying about
// race conditions.
package logger
