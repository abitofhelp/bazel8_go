package logger

import (
	"bytes"
	"context"
	"log"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestContextLogger(t *testing.T) {
	// Create a buffer to capture log output
	var buf bytes.Buffer
	logger := log.New(&buf, "", 0)

	// Create a context logger with our custom logger
	ctxLogger := NewContextLogger(logger)

	// Create a context
	ctx := context.Background()

	// Test Info logging with empty context
	ctxLogger.Info(ctx, "test info message")
	assert.True(t, strings.Contains(buf.String(), "INFO: test info message"), "Info log should contain the message")
	buf.Reset()

	// Test Warning logging with empty context
	ctxLogger.Warning(ctx, "test warning message")
	assert.True(t, strings.Contains(buf.String(), "WARNING: test warning message"), "Warning log should contain the message")
	buf.Reset()

	// Test Error logging with empty context
	ctxLogger.Error(ctx, "test error message")
	assert.True(t, strings.Contains(buf.String(), "ERROR: test error message"), "Error log should contain the message")
	buf.Reset()

	// We can't easily test Fatal as it calls os.Exit
}

func TestContextLoggerWithContextValues(t *testing.T) {
	// Create a buffer to capture log output
	var buf bytes.Buffer
	logger := log.New(&buf, "", 0)

	// Create a context logger with our custom logger
	ctxLogger := NewContextLogger(logger)

	// Create a context with request ID
	ctx := context.Background()
	ctx = WithRequestID(ctx, "req-123")

	// Test Info logging with request ID
	ctxLogger.Info(ctx, "test info message")
	logOutput := buf.String()
	assert.True(t, strings.Contains(logOutput, "INFO: [request_id=req-123] test info message"),
		"Info log should contain the request ID: %s", logOutput)
	buf.Reset()

	// Add user ID to context
	ctx = WithUserID(ctx, "user-456")

	// Test Warning logging with request ID and user ID
	ctxLogger.Warning(ctx, "test warning message")
	logOutput = buf.String()
	assert.True(t, strings.Contains(logOutput, "WARNING: [request_id=req-123 user_id=user-456] test warning message") ||
		strings.Contains(logOutput, "WARNING: [user_id=user-456 request_id=req-123] test warning message"),
		"Warning log should contain both request ID and user ID: %s", logOutput)
	buf.Reset()

	// Test Error logging with request ID and user ID
	ctxLogger.Error(ctx, "test error message")
	logOutput = buf.String()
	assert.True(t, strings.Contains(logOutput, "ERROR: [request_id=req-123 user_id=user-456] test error message") ||
		strings.Contains(logOutput, "ERROR: [user_id=user-456 request_id=req-123] test error message"),
		"Error log should contain both request ID and user ID: %s", logOutput)
	buf.Reset()
}

func TestExtractContextInfo(t *testing.T) {
	// Test with nil context
	assert.Equal(t, "", extractContextInfo(nil), "Nil context should return empty string")

	// Test with empty context
	ctx := context.Background()
	assert.Equal(t, "", extractContextInfo(ctx), "Empty context should return empty string")

	// Test with request ID only
	ctx = WithRequestID(ctx, "req-123")
	assert.Equal(t, "[request_id=req-123] ", extractContextInfo(ctx), "Context with request ID should return formatted string")

	// Test with user ID only
	ctx = context.Background()
	ctx = WithUserID(ctx, "user-456")
	assert.Equal(t, "[user_id=user-456] ", extractContextInfo(ctx), "Context with user ID should return formatted string")

	// Test with both request ID and user ID
	ctx = context.Background()
	ctx = WithRequestID(ctx, "req-123")
	ctx = WithUserID(ctx, "user-456")
	contextInfo := extractContextInfo(ctx)
	assert.True(t, contextInfo == "[request_id=req-123 user_id=user-456] " ||
		contextInfo == "[user_id=user-456 request_id=req-123] ",
		"Context with both IDs should return formatted string with both values: %s", contextInfo)

	// Test with empty values
	ctx = context.Background()
	ctx = WithRequestID(ctx, "")
	ctx = WithUserID(ctx, "")
	assert.Equal(t, "", extractContextInfo(ctx), "Context with empty values should return empty string")
}

func TestWithRequestID(t *testing.T) {
	ctx := context.Background()
	ctx = WithRequestID(ctx, "req-123")

	// Verify the value was stored in the context
	requestID, ok := ctx.Value(RequestIDKey).(string)
	assert.True(t, ok, "RequestID should be retrievable from context")
	assert.Equal(t, "req-123", requestID, "RequestID should match the value set")
}

func TestWithUserID(t *testing.T) {
	ctx := context.Background()
	ctx = WithUserID(ctx, "user-456")

	// Verify the value was stored in the context
	userID, ok := ctx.Value(UserIDKey).(string)
	assert.True(t, ok, "UserID should be retrievable from context")
	assert.Equal(t, "user-456", userID, "UserID should match the value set")
}

func TestNewContextLogger_Default(t *testing.T) {
	// Test that NewContextLogger uses the default logger when nil is provided
	ctxLogger := NewContextLogger(nil)
	assert.NotNil(t, ctxLogger, "ContextLogger should not be nil when created with nil logger")
	assert.NotNil(t, ctxLogger.logger, "Internal logger should not be nil when created with nil logger")
}

func TestDefaultLogger(t *testing.T) {
	// Test that DefaultLogger is not nil
	assert.NotNil(t, DefaultLogger, "DefaultLogger should not be nil")
}
