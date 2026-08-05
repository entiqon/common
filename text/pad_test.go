package text_test

import (
	"testing"

	"github.com/entiqon/common/text"
)

func TestPad(t *testing.T) {
	t.Run("Pad", func(t *testing.T) {
		t.Run("Left", func(t *testing.T) {
			tests := []struct {
				name     string
				value    string
				length   int
				pad      rune
				expected string
			}{
				{"Pad", "42", 5, '0', "00042"},
				{"SameLength", "Hello", 5, '0', "Hello"},
				{"Longer", "Hello", 3, '0', "Hello"},
				{"Empty", "", 3, '*', "***"},
			}

			for _, tt := range tests {
				t.Run(tt.name, func(t *testing.T) {
					got := text.PadLeft(tt.value, tt.length, tt.pad)
					if got != tt.expected {
						t.Fatalf("expected %q, got %q", tt.expected, got)
					}
				})
			}
		})

		t.Run("Center", func(t *testing.T) {
			tests := []struct {
				name     string
				value    string
				length   int
				pad      rune
				expected string
			}{
				{"Even", "Go", 6, ' ', "  Go  "},
				{"Odd", "Go", 5, ' ', " Go  "},
				{"SameLength", "Hello", 5, '.', "Hello"},
				{"Longer", "Hello", 3, '.', "Hello"},
				{"Empty", "", 4, '*', "****"},
			}

			for _, tt := range tests {
				t.Run(tt.name, func(t *testing.T) {
					got := text.PadCenter(tt.value, tt.length, tt.pad)
					if got != tt.expected {
						t.Fatalf("expected %q, got %q", tt.expected, got)
					}
				})
			}
		})

		t.Run("Right", func(t *testing.T) {
			tests := []struct {
				name     string
				value    string
				length   int
				pad      rune
				expected string
			}{
				{"Pad", "42", 5, '.', "42..."},
				{"SameLength", "Hello", 5, '.', "Hello"},
				{"Longer", "Hello", 3, '.', "Hello"},
				{"Empty", "", 3, '*', "***"},
			}

			for _, tt := range tests {
				t.Run(tt.name, func(t *testing.T) {
					got := text.PadRight(tt.value, tt.length, tt.pad)
					if got != tt.expected {
						t.Fatalf("expected %q, got %q", tt.expected, got)
					}
				})
			}
		})

		t.Run("EdgeCases", func(t *testing.T) {
			t.Run("InvalidDirection", func(t *testing.T) {

			})
		})
	})
}
