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
	
	// Test Info logging
	ctxLogger.Info(ctx, "test info message")
	assert.True(t, strings.Contains(buf.String(), "INFO: test info message"), "Info log should contain the message")
	buf.Reset()
	
	// Test Warning logging
	ctxLogger.Warning(ctx, "test warning message")
	assert.True(t, strings.Contains(buf.String(), "WARNING: test warning message"), "Warning log should contain the message")
	buf.Reset()
	
	// Test Error logging
	ctxLogger.Error(ctx, "test error message")
	assert.True(t, strings.Contains(buf.String(), "ERROR: test error message"), "Error log should contain the message")
	buf.Reset()
	
	// We can't easily test Fatal as it calls os.Exit
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