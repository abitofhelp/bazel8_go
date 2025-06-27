// Package greeting provides functionality for generating personalized greeting messages.
//
// This package is designed to create friendly greeting messages that can include
// personalized information such as the recipient's name and a monetary amount.
// It handles validation of inputs and provides clear error messages when invalid
// inputs are provided. The package is context-aware and supports context cancellation
// and timeout.
//
// Basic usage:
//
//	ctx := context.Background()
//	message, err := greeting.Greet(ctx, "John", 1234567)
//	if err != nil {
//	    log.Fatalf("Error generating greeting: %v", err)
//	}
//	fmt.Println(message)
//
// With timeout:
//
//	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
//	defer cancel()
//	message, err := greeting.Greet(ctx, "John", 1234567)
//	if err != nil {
//	    if errors.Is(err, greeting.ErrContextDeadlineExceeded) {
//	        log.Println("Operation timed out")
//	    } else {
//	        log.Fatalf("Error generating greeting: %v", err)
//	    }
//	}
//
// The package uses the go-humanize library to format monetary amounts in a
// human-readable way, with proper comma separation and decimal places.
// For example, a winnings value of 1234567 (cents) will be formatted as $12,345.67.
package greeting
