package ptr_test

import (
	"testing"
	"time"

	"github.com/entiqon/common/ptr"
)

type sample struct {
	A int
	B string
}

func TestPtrSuite(t *testing.T) {
	t.Run("ptr", func(t *testing.T) {
		t.Run("ToPtr", func(t *testing.T) {
			t.Run("Int", func(t *testing.T) {
				got := ptr.ToPtr(42)

				if got == nil {
					t.Fatal("ToPtr returned nil")
				}

				if *got != 42 {
					t.Errorf("ToPtr(42) = %d; want 42", *got)
				}
			})

			t.Run("String", func(t *testing.T) {
				got := ptr.ToPtr("hello")

				if got == nil {
					t.Fatal("ToPtr returned nil")
				}

				if *got != "hello" {
					t.Errorf(`ToPtr("hello") = %q; want "hello"`, *got)
				}
			})

			t.Run("Struct", func(t *testing.T) {
				want := sample{
					A: 7,
					B: "seven",
				}

				got := ptr.ToPtr(want)

				if got == nil {
					t.Fatal("ToPtr returned nil")
				}

				if *got != want {
					t.Errorf("ToPtr(%v) = %v; want %v", want, *got, want)
				}
			})
		})

		t.Run("FromPtr", func(t *testing.T) {
			t.Run("ReturnsPointerValue", func(t *testing.T) {
				value := "hello"

				got := ptr.FromPtr(&value, "fallback")

				if got != value {
					t.Errorf("FromPtr() = %q; want %q", got, value)
				}
			})

			t.Run("ReturnsDefaultForNilPointer", func(t *testing.T) {
				got := ptr.FromPtr[string](nil, "fallback")

				if got != "fallback" {
					t.Errorf(`FromPtr(nil) = %q; want "fallback"`, got)
				}
			})

			t.Run("SupportsStruct", func(t *testing.T) {
				want := sample{
					A: 7,
					B: "seven",
				}

				got := ptr.FromPtr(
					ptr.ToPtr(want),
					sample{},
				)

				if got != want {
					t.Errorf("FromPtr() = %v; want %v", got, want)
				}
			})

			t.Run("Time", func(t *testing.T) {
				want := time.Date(
					2026,
					time.July,
					31,
					13,
					57,
					0,
					0,
					time.UTC,
				)

				got := ptr.FromPtr(
					ptr.ToPtr(want),
					time.Time{},
				)

				if !got.Equal(want) {
					t.Errorf(
						"FromPtr() = %s; want %s",
						got.Format(time.RFC3339Nano),
						want.Format(time.RFC3339Nano),
					)
				}
			})
		})
	})

	t.Run("NilIfEmpty", func(t *testing.T) {
		tests := []struct {
			name    string
			input   string
			wantNil bool
		}{
			{
				name:    "Empty",
				input:   "",
				wantNil: true,
			},
			{
				name:    "Whitespace",
				input:   "   ",
				wantNil: false,
			},
			{
				name:    "Value",
				input:   "hello",
				wantNil: false,
			},
		}

		for _, tc := range tests {
			tc := tc

			t.Run(tc.name, func(t *testing.T) {
				got := ptr.NilIfEmpty(tc.input)

				if tc.wantNil {
					if got != nil {
						t.Errorf(
							"NilIfEmpty(%q) = %q; want nil",
							tc.input,
							*got,
						)
					}

					return
				}

				if got == nil {
					t.Fatalf(
						"NilIfEmpty(%q) returned nil",
						tc.input,
					)
				}

				if *got != tc.input {
					t.Errorf(
						"NilIfEmpty(%q) = %q; want %q",
						tc.input,
						*got,
						tc.input,
					)
				}
			})
		}
	})

	t.Run("NilIfBlank", func(t *testing.T) {
		tests := []struct {
			name    string
			input   string
			wantNil bool
		}{
			{
				name:    "Empty",
				input:   "",
				wantNil: true,
			},
			{
				name:    "Spaces",
				input:   "   ",
				wantNil: true,
			},
			{
				name:    "Tab",
				input:   "\t",
				wantNil: true,
			},
			{
				name:    "NewLine",
				input:   "\n",
				wantNil: true,
			},
			{
				name:    "Value",
				input:   "hello",
				wantNil: false,
			},
			{
				name:    "ValueWithWhitespace",
				input:   "  hello  ",
				wantNil: false,
			},
		}

		for _, tc := range tests {
			tc := tc

			t.Run(tc.name, func(t *testing.T) {
				got := ptr.NilIfBlank(tc.input)

				if tc.wantNil {
					if got != nil {
						t.Errorf(
							"NilIfBlank(%q) = %q; want nil",
							tc.input,
							*got,
						)
					}

					return
				}

				if got == nil {
					t.Fatalf(
						"NilIfBlank(%q) returned nil",
						tc.input,
					)
				}

				if *got != tc.input {
					t.Errorf(
						"NilIfBlank(%q) = %q; want %q",
						tc.input,
						*got,
						tc.input,
					)
				}
			})
		}
	})
}
