package text_test

import (
	"testing"

	"github.com/entiqon/common/text"
)

func TestCapitalize(t *testing.T) {
	t.Run("Capitalize", func(t *testing.T) {
		tests := []struct {
			name     string
			value    string
			expected string
		}{
			{"Lowercase", "hello", "Hello"},
			{"Uppercase", "Hello", "Hello"},
			{"Unicode", "ángel", "Ángel"},
			{"Single", "h", "H"},
			{"Empty", "", ""},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				got := text.Capitalize(tt.value)
				if got != tt.expected {
					t.Fatalf("expected %q, got %q", tt.expected, got)
				}
			})
		}
	})

	t.Run("Uncapitalize", func(t *testing.T) {
		tests := []struct {
			name     string
			value    string
			expected string
		}{
			{"Uppercase", "Hello", "hello"},
			{"Lowercase", "hello", "hello"},
			{"Unicode", "Ángel", "ángel"},
			{"Single", "H", "h"},
			{"Empty", "", ""},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				got := text.Uncapitalize(tt.value)
				if got != tt.expected {
					t.Fatalf("expected %q, got %q", tt.expected, got)
				}
			})
		}
	})
}
