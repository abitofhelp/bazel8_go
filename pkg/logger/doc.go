// Package logger provides logging utilities for the application.
//
// This package offers a context-aware logging system that ensures all log messages
// include relevant context information. It provides a ContextLogger type with methods
// for logging at different severity levels (info, warning, error, fatal).
//
// The package also provides a DefaultLogger singleton instance that can be used
// throughout the application without needing to create a new logger instance.
//
// Context Information Extraction:
//
// The logger automatically extracts and includes relevant information from the context
// in log messages, such as request IDs and user IDs. This makes it easier to trace
// and debug issues across different parts of the application.
//
// Basic usage:
//
//	ctx := context.Background()
//	logger.DefaultLogger.Info(ctx, "Processing item %d", itemID)
//
//	// Or create a custom logger
//	customLogger := logger.NewContextLogger(log.New(os.Stdout, "CUSTOM: ", log.LstdFlags))
//	customLogger.Warning(ctx, "Unusual condition detected: %v", condition)
//
// Adding context values:
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
// All logging methods require a context.Context as their first parameter, following
// the project's convention that any function requiring context must have it as the
// first parameter.
package logger
