package text_test

import (
	"testing"

	"github.com/entiqon/common/text"
)

func TestReverse(t *testing.T) {
	t.Run("Reverse", func(t *testing.T) {
		tests := []struct {
			name     string
			value    string
			expected string
		}{
			{"Word", "Hello", "olleH"},
			{"Sentence", "Hello World", "dlroW olleH"},
			{"Unicode", "こんにちは", "はちにんこ"},
			{"Single", "A", "A"},
			{"Empty", "", ""},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				got := text.Reverse(tt.value)
				if got != tt.expected {
					t.Fatalf("expected %q, got %q", tt.expected, got)
				}
			})
		}
	})
}
