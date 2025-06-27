package greeting

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

// customErrorContext is a custom implementation of context.Context that always returns a specific error
type customErrorContext struct {
	context.Context
	err error
}

func (c *customErrorContext) Err() error {
	return c.err
}

func (c *customErrorContext) Deadline() (deadline time.Time, ok bool) {
	return time.Time{}, false
}

func TestGreet(t *testing.T) {
	tests := []struct {
		name        string
		ctx         context.Context
		input       string
		expected    string
		expectedErr error
		setupCtx    func() (context.Context, context.CancelFunc)
		cancelCtx   bool
		cancelDelay time.Duration
	}{
		{
			name:        "valid input",
			setupCtx:    func() (context.Context, context.CancelFunc) { return context.Background(), func() {} },
			input:       "John",
			expected:    "Howdy John!\n",
			expectedErr: nil,
		},
		{
			name:        "empty name",
			setupCtx:    func() (context.Context, context.CancelFunc) { return context.Background(), func() {} },
			input:       "",
			expected:    "",
			expectedErr: ErrInvalidName,
		},
		{
			name:        "another valid input",
			setupCtx:    func() (context.Context, context.CancelFunc) { return context.Background(), func() {} },
			input:       "Alice",
			expected:    "Howdy Alice!\n",
			expectedErr: nil,
		},
		{
			name: "canceled context before call",
			setupCtx: func() (context.Context, context.CancelFunc) {
				ctx, cancel := context.WithCancel(context.Background())
				cancel()
				return ctx, cancel
			},
			input:       "John",
			expected:    "",
			expectedErr: ErrContextCanceled,
		},
		{
			name: "context with deadline exceeded before call",
			setupCtx: func() (context.Context, context.CancelFunc) {
				ctx, cancel := context.WithTimeout(context.Background(), -time.Second)
				return ctx, cancel
			},
			input:       "John",
			expected:    "",
			expectedErr: ErrContextDeadlineExceeded,
		},
		{
			name:        "context canceled during processing",
			setupCtx:    func() (context.Context, context.CancelFunc) { return context.WithCancel(context.Background()) },
			input:       "John",
			expected:    "",
			expectedErr: ErrContextCanceled,
			cancelCtx:   true,
			cancelDelay: 10 * time.Millisecond, // Cancel after a short delay
		},
		{
			name: "context deadline exceeded during processing",
			setupCtx: func() (context.Context, context.CancelFunc) {
				return context.WithTimeout(context.Background(), 50*time.Millisecond)
			},
			input:       "John",
			expected:    "",
			expectedErr: ErrContextDeadlineExceeded,
			cancelCtx:   false,
		},
		{
			name: "other context error",
			setupCtx: func() (context.Context, context.CancelFunc) {
				// Create a custom context with a custom error
				ctx := context.WithValue(context.Background(), struct{}{}, "value")
				// Wrap it with a custom context that returns a custom error
				ctx = &customErrorContext{ctx, errors.New("custom error")}
				return ctx, func() {}
			},
			input:       "John",
			expected:    "",
			expectedErr: errors.New("custom error"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx, cancel := tt.setupCtx()
			defer cancel()

			if tt.cancelCtx {
				go func() {
					time.Sleep(tt.cancelDelay)
					cancel()
				}()
			}

			result, err := Greet(ctx, tt.input)

			if tt.expectedErr != nil {
				assert.Error(t, err)
				if tt.name == "other context error" {
					// For the custom error case, just check the error message
					assert.Equal(t, tt.expectedErr.Error(), err.Error(), "Expected error %v, got %v", tt.expectedErr, err)
				} else {
					// For other cases, use errors.Is
					assert.True(t, errors.Is(err, tt.expectedErr), "Expected error %v, got %v", tt.expectedErr, err)
				}
			} else {
				assert.NoError(t, err)
			}

			assert.Equal(t, tt.expected, result)
		})
	}
}
