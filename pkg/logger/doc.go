// Package logger provides logging utilities for the application.
//
// This package offers a context-aware logging system that ensures all log messages
// include relevant context information. It provides a ContextLogger type with methods
// for logging at different severity levels (info, warning, error, fatal).
//
// The package also provides a DefaultLogger singleton instance that can be used
// throughout the application without needing to create a new logger instance.
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
// All logging methods require a context.Context as their first parameter, following
// the project's convention that any function requiring context must have it as the
// first parameter.
package logger