// Package greeting provides functionality for generating personalized greeting messages.
//
// # Overview
//
// This package is designed to create friendly greeting messages that can include
// personalized information such as the recipient's name and a monetary amount.
// It handles validation of inputs and provides clear error messages when invalid
// inputs are provided. The package is context-aware and supports context cancellation
// and timeout.
//
// # Key Features
//
// - Personalized greeting messages with the recipient's name
// - Formatting of monetary amounts in a human-readable way
// - Context-aware operations with support for cancellation and timeouts
// - Comprehensive error handling with specific error types
// - Logging of operations for debugging and monitoring
//
// # Basic Usage
//
// Here's a simple example of how to use the greeting package:
//
//	ctx := context.Background()
//	message, err := greeting.Greet(ctx, "John", 1234567)
//	if err != nil {
//	    log.Fatalf("Error generating greeting: %v", err)
//	}
//	fmt.Println(message)
//	// Output: Howdy John! You have won $12,345.67 USD!
//
// # Advanced Usage with Timeout
//
// For operations that should not run indefinitely, you can use a timeout context:
//
//	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
//	defer cancel() // Always call cancel to release resources
//	message, err := greeting.Greet(ctx, "John", 1234567)
//	if err != nil {
//	    if errors.Is(err, greeting.ErrContextDeadlineExceeded) {
//	        log.Println("Operation timed out")
//	    } else {
//	        log.Fatalf("Error generating greeting: %v", err)
//	    }
//	}
//
// # Error Handling
//
// The package defines several error types to help with error handling:
//
// - ErrInvalidName: Returned when the provided name is empty
// - ErrInvalidWinnings: Returned when the provided winnings amount is negative
// - ErrContextCanceled: Returned when the context is canceled during processing
// - ErrContextDeadlineExceeded: Returned when the context deadline is exceeded
//
// It's recommended to use errors.Is() to check for specific error types:
//
//	if errors.Is(err, greeting.ErrInvalidName) {
//	    // Handle invalid name error
//	}
//
// # Monetary Amount Formatting
//
// The package uses the go-humanize library to format monetary amounts in a
// human-readable way, with proper comma separation and decimal places.
// For example, a winnings value of 1234567 (cents) will be formatted as $12,345.67.
//
// # Parameters
//
// The Greet function takes the following parameters:
//
// - ctx (context.Context): The context for the operation, which can be used for cancellation and timeouts
// - name (string): The name of the person to greet (must not be empty)
// - winnings (uint64): The amount of winnings in cents (e.g., 1234567 = $12,345.67)
package greeting
