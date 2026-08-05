package text_test

import (
	"testing"

	"github.com/entiqon/common/text"
)

func TestTruncate(t *testing.T) {
	tests := []struct {
		name     string
		value    string
		max      int
		expected string
	}{
		{"Truncate", "Hello World", 5, "Hello"},
		{"Equal", "Hello", 5, "Hello"},
		{"Shorter", "Hi", 5, "Hi"},
		{"Unicode", "こんにちは世界", 4, "こんにち"},
		{"Zero", "Hello", 0, ""},
		{"Negative", "Hello", -1, ""},
		{"Empty", "", 5, ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := text.Truncate(tt.value, tt.max)
			if got != tt.expected {
				t.Fatalf("expected %q, got %q", tt.expected, got)
			}
		})
	}
}
