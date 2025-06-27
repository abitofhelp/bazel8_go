package greeting

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGreet(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "empty name",
			input:    "",
			expected: "Hello, World!",
		},
		{
			name:     "with name",
			input:    "John",
			expected: "Hello, John!",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Greet(tt.input)
			assert.Equal(t, tt.expected, result)
		})
	}
}