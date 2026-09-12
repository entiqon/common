package internal_test

import (
	"testing"
	"time"

	"github.com/entiqon/common/date/internal"
)

func TestInternal(t *testing.T) {
	t.Run("Methods", func(t *testing.T) {
		t.Run("ParseRFC3339", func(t *testing.T) {
			okCases := []string{
				"2023-11-14T22:13:20Z",
				"2023-11-14T22:13:20+02:00",
			}
			for _, in := range okCases {
				if got, ok := internal.ParseRFC3339(in); !ok {
					t.Errorf("expected success for %q", in)
				} else if got.IsZero() {
					t.Errorf("unexpected zero time for %q", in)
				}
			}
			if _, ok := internal.ParseRFC3339("not-a-date"); ok {
				t.Errorf("expected failure")
			}
		})

		t.Run("ParseYYYYMMDDPrefix", func(t *testing.T) {
			t.Run("EmptyInput", func(t *testing.T) {
				_, err := internal.ParseYYYYMMDDPrefix("")
				if err == nil {
					t.Fatalf("expected error for empty input, got nil")
				}

				_, err = internal.ParseYYYYMMDDPrefix("   ") // whitespace only
				if err == nil {
					t.Fatalf("expected error for whitespace input, got nil")
				}
			})

			t.Run("Valid", func(t *testing.T) {
				got, err := internal.ParseYYYYMMDDPrefix("20240229.suffix")
				if err != nil {
					t.Fatalf("unexpected error: %v", err)
				}
				want := time.Date(2024, 2, 29, 0, 0, 0, 0, time.UTC)
				if !got.Equal(want) {
					t.Errorf("want %s, got %s", want, got)
				}
			})

			t.Run("TooShort", func(t *testing.T) {
				if _, err := internal.ParseYYYYMMDDPrefix("2025"); err == nil {
					t.Errorf("expected error")
				}
			})

			t.Run("NonDigit", func(t *testing.T) {
				if _, err := internal.ParseYYYYMMDDPrefix("2025A507"); err == nil {
					t.Errorf("expected error")
				}
			})

			t.Run("InvalidMonth", func(t *testing.T) {
				if _, err := internal.ParseYYYYMMDDPrefix("20251301"); err == nil {
					t.Errorf("expected error")
				}
			})

			t.Run("InvalidDay", func(t *testing.T) {
				if _, err := internal.ParseYYYYMMDDPrefix("20250132"); err == nil {
					t.Errorf("expected error")
				}
			})

			t.Run("InvalidCalendar", func(t *testing.T) {
				if _, err := internal.ParseYYYYMMDDPrefix("20230229"); err == nil {
					t.Errorf("expected error")
				}
			})
		})

		t.Run("ParseEpoch", func(t *testing.T) {
			t.Run("Seconds", func(t *testing.T) {
				got, ok := internal.ParseEpoch("1700000000")
				if !ok || got.Year() != 2023 {
					t.Errorf("unexpected result %v %v", got, ok)
				}
			})

			t.Run("Milliseconds", func(t *testing.T) {
				got, ok := internal.ParseEpoch("1700000000000")
				if !ok || got.Year() != 2023 {
					t.Errorf("unexpected result %v %v", got, ok)
				}
			})

			t.Run("NotDigits", func(t *testing.T) {
				if _, ok := internal.ParseEpoch("12345abc"); ok {
					t.Errorf("expected failure")
				}
			})

			t.Run("WrongLength", func(t *testing.T) {
				if _, ok := internal.ParseEpoch("123456"); ok {
					t.Errorf("expected failure")
				}
			})
		})

		t.Run("ParseZoned", func(t *testing.T) {
			t.Run("Valid", func(t *testing.T) {
				in := "Tue, 14 Nov 2023 22:13:20 UTC"
				got, ok := internal.ParseZoned(time.RFC1123, in)
				if !ok || got.IsZero() {
					t.Errorf("expected success")
				}
			})

			t.Run("Invalid", func(t *testing.T) {
				if _, ok := internal.ParseZoned(time.RFC1123, "garbage"); ok {
					t.Errorf("expected failure")
				}
			})
		})

	})
}
