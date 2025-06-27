// Package greeting provides functionality for generating personalized greeting messages.
// See doc.go for detailed package documentation.
package greeting

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/abitofhelp/bazel8_go/pkg/logger"
)

// Error definitions for the greeting package.
var (
	// ErrInvalidName is returned when the provided name is empty.
	ErrInvalidName = errors.New("name cannot be empty")

	// ErrContextCanceled is returned when the context is canceled during processing.
	ErrContextCanceled = errors.New("operation was canceled by context")

	// ErrContextDeadlineExceeded is returned when the context deadline is exceeded during processing.
	ErrContextDeadlineExceeded = errors.New("operation timed out")
)

// Greet generates a personalized greeting message for the given name.
//
// # Function Purpose
//
// This function creates a friendly greeting message that includes the recipient's name.
// It's designed to be used in applications that need to generate personalized messages,
// such as welcome notifications.
//
// # Parameters
//
//   - ctx: A context.Context that can be used to cancel the operation or set a deadline.
//     This follows the project convention of having context as the first parameter.
//
//   - name: The name of the person to greet. This must be a non-empty string.
//     If an empty string is provided, ErrInvalidName will be returned.
//
// # Return Values
//
//   - string: The formatted greeting message if successful.
//     Example: "Howdy John!"
//
// - error: An error if something went wrong. Possible errors include:
//   - ErrInvalidName: If the name parameter is empty
//   - ErrContextCanceled: If the context was canceled during processing
//   - ErrContextDeadlineExceeded: If the context deadline was exceeded
//
// # Example Usage
//
//	ctx := context.Background()
//	message, err := greeting.Greet(ctx, "John")
//	if err != nil {
//	    log.Fatalf("Error: %v", err)
//	}
//	fmt.Println(message)
//	// Output: Howdy John!
//
// # Error Handling
//
// It's recommended to use errors.Is() to check for specific error types:
//
//	if errors.Is(err, greeting.ErrInvalidName) {
//	    // Handle invalid name error
//	} else if errors.Is(err, greeting.ErrContextCanceled) {
//	    // Handle context cancellation
//	} else if errors.Is(err, greeting.ErrContextDeadlineExceeded) {
//	    // Handle timeout
//	} else if err != nil {
//	    // Handle other errors
//	}
func Greet(ctx context.Context, name string) (string, error) {
	// Check if context is already canceled or deadline exceeded
	if ctx.Err() != nil {
		if errors.Is(ctx.Err(), context.Canceled) {
			logger.DefaultLogger.Warning(ctx, "Context was canceled before processing: %v", ctx.Err())
			return "", ErrContextCanceled
		}
		if errors.Is(ctx.Err(), context.DeadlineExceeded) {
			logger.DefaultLogger.Warning(ctx, "Context deadline exceeded before processing: %v", ctx.Err())
			return "", ErrContextDeadlineExceeded
		}
		// For any other context error
		logger.DefaultLogger.Error(ctx, "Context error: %v", ctx.Err())
		return "", ctx.Err()
	}

	// Validate input parameters
	if name == "" {
		logger.DefaultLogger.Warning(ctx, "Invalid name provided: empty string")
		return "", ErrInvalidName
	}

	logger.DefaultLogger.Info(ctx, "Generating greeting for '%s'", name)

	// Simulate some processing time to demonstrate context handling
	select {
	case <-ctx.Done():
		if errors.Is(ctx.Err(), context.Canceled) {
			logger.DefaultLogger.Warning(ctx, "Context was canceled during processing: %v", ctx.Err())
			return "", ErrContextCanceled
		}
		if errors.Is(ctx.Err(), context.DeadlineExceeded) {
			logger.DefaultLogger.Warning(ctx, "Context deadline exceeded during processing: %v", ctx.Err())
			return "", ErrContextDeadlineExceeded
		}
		// For any other context error
		logger.DefaultLogger.Error(ctx, "Context error during processing: %v", ctx.Err())
		return "", ctx.Err()
	case <-time.After(100 * time.Millisecond): // Simulate a short processing time
		// Continue processing
	}

	message := fmt.Sprintf("Howdy %s!\n", name)

	logger.DefaultLogger.Info(ctx, "Generated greeting: %s", message)
	return message, nil
}
