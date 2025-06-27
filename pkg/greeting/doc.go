// Package greeting provides functionality for generating personalized greeting messages.
//
// This package is designed to create friendly greeting messages that can include
// personalized information such as the recipient's name and a monetary amount.
// It handles validation of inputs and provides clear error messages when invalid
// inputs are provided.
//
// Basic usage:
//
//	message, err := greeting.Greet("John", 1234567)
//	if err != nil {
//	    log.Fatalf("Error generating greeting: %v", err)
//	}
//	fmt.Println(message)
//
// The package uses the go-humanize library to format monetary amounts in a
// human-readable way, with proper comma separation and decimal places.
package greeting