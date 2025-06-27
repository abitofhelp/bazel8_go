// Package main provides the entry point for the bazel8_go application.
//
// This application demonstrates the use of the greeting package to generate
// personalized greeting messages with formatted monetary amounts. It serves
// as an example of how to use the greeting package in a real application.
//
// The main package:
// - Imports the greeting package
// - Calls the Greet function with a name and winning amount
// - Handles any errors that might occur
// - Prints the resulting greeting message to standard output
//
// Usage:
//
//	go run cmd/main.go
//
// Or with Bazel:
//
//	bazel run //cmd:main
//
// The application is built using Bazel, which provides a consistent build
// system across different platforms and environments.
package main
