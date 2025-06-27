// Package greeting provides greeting functionality.
package greeting

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/dustin/go-humanize"

	"github.com/abitofhelp/bazel8_go/pkg/logger"
)

// Error definitions for the greeting package.
var (
	// ErrInvalidName is returned when the provided name is empty.
	ErrInvalidName = errors.New("name cannot be empty")

	// ErrInvalidWinnings is returned when the provided winnings amount is negative.
	ErrInvalidWinnings = errors.New("winnings amount cannot be negative")

	// ErrContextCanceled is returned when the context is canceled during processing.
	ErrContextCanceled = errors.New("operation was canceled by context")

	// ErrContextDeadlineExceeded is returned when the context deadline is exceeded during processing.
	ErrContextDeadlineExceeded = errors.New("operation timed out")
)

// Greet returns a personalized greeting message for the given name and winning amount.
// The function validates that the name is not empty and formats the monetary amount
// using the go-humanize library to make it more readable.
//
// The winnings amount is expected to be in cents (e.g., 1234567 = $12,345.67).
//
// If the name is empty, ErrInvalidName is returned.
// If the context is canceled, ErrContextCanceled is returned.
// If the context deadline is exceeded, ErrContextDeadlineExceeded is returned.
//
// Example:
//
//	ctx := context.Background()
//	message, err := greeting.Greet(ctx, "John", 1234567)
//	if err != nil {
//	    log.Fatalf("Error: %v", err)
//	}
//	// Output: Howdy John! You have won $12,345.67 USD!
func Greet(ctx context.Context, name string, winnings uint64) (string, error) {
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

	logger.DefaultLogger.Info(ctx, "Generating greeting for %s with winnings of %d cents", name, winnings)

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

	// Format currency with two decimal places
	amount := humanize.CommafWithDigits(float64(winnings)/100.0, 2)
	message := fmt.Sprintf("Howdy %s! You have won $%s USD!\n", name, amount)

	logger.DefaultLogger.Info(ctx, "Generated greeting: %s", message)
	return message, nil
}
