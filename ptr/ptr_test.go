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
		t.Run("EmptyString", func(t *testing.T) {
			got := ptr.EmptyString()

			if got == nil {
				t.Fatal("EmptyString returned nil")
			}

			if *got != "" {
				t.Errorf(`EmptyString() = %q; want ""`, *got)
			}

			// Ensure each call returns a distinct pointer.
			other := ptr.EmptyString()
			if got == other {
				t.Error("EmptyString returned the same pointer instance")
			}
		})

		t.Run("Of", func(t *testing.T) {
			t.Run("Int", func(t *testing.T) {
				got := ptr.Of(42)

				if got == nil {
					t.Fatal("Of returned nil")
				}

				if *got != 42 {
					t.Errorf("Of(42) = %d; want 42", *got)
				}
			})

			t.Run("String", func(t *testing.T) {
				got := ptr.Of("hello")

				if got == nil {
					t.Fatal("Of returned nil")
				}

				if *got != "hello" {
					t.Errorf(`Of("hello") = %q; want "hello"`, *got)
				}
			})

			t.Run("Struct", func(t *testing.T) {
				want := sample{
					A: 7,
					B: "seven",
				}

				got := ptr.Of(want)

				if got == nil {
					t.Fatal("Of returned nil")
				}

				if *got != want {
					t.Errorf("Of(%v) = %v; want %v", want, *got, want)
				}
			})
		})

		t.Run("Deref", func(t *testing.T) {
			t.Run("Pointer", func(t *testing.T) {
				value := "hello"

				got := ptr.Deref(&value, "fallback")

				if got != value {
					t.Errorf("Deref() = %q; want %q", got, value)
				}
			})

			t.Run("Fallback", func(t *testing.T) {
				var missing *string

				got := ptr.Deref(missing, "fallback")

				if got != "fallback" {
					t.Errorf(`Deref(nil) = %q; want "fallback"`, got)
				}
			})

			t.Run("Struct", func(t *testing.T) {
				want := sample{
					A: 7,
					B: "seven",
				}

				got := ptr.Deref(
					ptr.Of(want),
					sample{},
				)

				if got != want {
					t.Errorf("Deref() = %v; want %v", got, want)
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

				got := ptr.Deref(
					ptr.Of(want),
					time.Time{},
				)

				if !got.Equal(want) {
					t.Errorf(
						"Deref() = %s; want %s",
						got.Format(time.RFC3339Nano),
						want.Format(time.RFC3339Nano),
					)
				}
			})
		})
	})

	t.Run("Deprecated", func(t *testing.T) {
		t.Run("ToPtr", func(t *testing.T) {
			got := ptr.ToPtr(42)

			if got == nil {
				t.Fatal("ToPtr returned nil")
			}

			if *got != 42 {
				t.Errorf("ToPtr(42) = %d; want 42", *got)
			}
		})

		t.Run("FromPtr", func(t *testing.T) {
			var missing *string

			got := ptr.FromPtr(missing, "fallback")

			if got != "fallback" {
				t.Errorf(`FromPtr(nil) = %q; want "fallback"`, got)
			}
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

		for i := range tests {
			tc := tests[i]

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

		for i := range tests {
			tc := tests[i]

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
