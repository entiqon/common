package boolean_test

import (
	"testing"

	"github.com/entiqon/common/boolean"
)

func TestBoolean(t *testing.T) {
	t.Run("From", func(t *testing.T) {
		check := func(t *testing.T, in any, want bool) {
			t.Helper()

			got, err := boolean.From(in)
			if err != nil {
				t.Errorf("From(%#v) unexpected error: %v", in, err)
				return
			}
			if got != want {
				t.Errorf("From(%#v) = %v; want %v", in, got, want)
			}
		}

		rejects := func(t *testing.T, in any) {
			t.Helper()

			if _, err := boolean.From(in); err == nil {
				t.Errorf("From(%#v) expected error, got nil", in)
			}
		}

		t.Run("Nil", func(t *testing.T) {
			rejects(t, nil)
		})

		t.Run("Bool", func(t *testing.T) {
			check(t, true, true)
			check(t, false, false)
		})

		t.Run("TrueStrings", func(t *testing.T) {
			for _, in := range []string{"1", "t", "true", "on", "y", "yes", " TRUE ", "Yes"} {
				check(t, in, true)
			}
		})

		t.Run("FalseStrings", func(t *testing.T) {
			for _, in := range []string{"0", "f", "false", "off", "n", "no", " FALSE ", "No"} {
				check(t, in, false)
			}
		})

		t.Run("InvalidStrings", func(t *testing.T) {
			for _, in := range []string{"enable", "disable", "ok", "", "  ", "sure", "nah"} {
				rejects(t, in)
			}
		})

		t.Run("Integers", func(t *testing.T) {
			check(t, 0, false)
			check(t, 1, true)
			check(t, -1, true)
			check(t, int8(0), false)
			check(t, int8(-128), true)
			check(t, int16(0), false)
			check(t, int16(32767), true)
			check(t, int32(0), false)
			check(t, int32(-2147483648), true)
			check(t, int64(0), false)
			check(t, int64(9223372036854775807), true)
		})

		t.Run("UnsignedIntegers", func(t *testing.T) {
			check(t, uint(0), false)
			check(t, uint(12345), true)
			check(t, uint8(0), false)
			check(t, uint8(255), true)
			check(t, uint16(0), false)
			check(t, uint16(65535), true)
			check(t, uint32(0), false)
			check(t, uint32(4294967295), true)
			check(t, uint64(0), false)
			check(t, uint64(18446744073709551615), true)
		})

		t.Run("Floats", func(t *testing.T) {
			check(t, float32(0), false)
			check(t, float32(-0.0), false)
			check(t, float32(1.5), true)
			check(t, float32(-2.3), true)
			check(t, 0.0, false)
			check(t, 3.14159, true)
			check(t, -123.456, true)
		})

		t.Run("UnsupportedTypes", func(t *testing.T) {
			rejects(t, struct{}{})
			rejects(t, []int{1, 2, 3})
			rejects(t, map[string]bool{"x": true})
			rejects(t, complex(1, 0))
			rejects(t, make(chan int))
			rejects(t, func() {})
		})
	})

	t.Run("FromOr", func(t *testing.T) {
		check := func(t *testing.T, in any, fallback, want bool) {
			t.Helper()

			if got := boolean.FromOr(in, fallback); got != want {
				t.Errorf("FromOr(%#v, %v) = %v; want %v", in, fallback, got, want)
			}
		}

		t.Run("Valid", func(t *testing.T) {
			check(t, "yes", false, true)
			check(t, "off", true, false)
			check(t, 1, false, true)
		})

		t.Run("Invalid", func(t *testing.T) {
			check(t, "maybe", true, true)
			check(t, "maybe", false, false)
		})

		t.Run("Nil", func(t *testing.T) {
			check(t, nil, true, true)
		})

		t.Run("UnsupportedType", func(t *testing.T) {
			check(t, []int{1}, false, false)
		})
	})

	t.Run("ToString", func(t *testing.T) {
		check := func(t *testing.T, b bool, trueStr, falseStr, want string) {
			t.Helper()

			if got := boolean.ToString(b, trueStr, falseStr); got != want {
				t.Errorf("ToString(%v, %q, %q) = %q; want %q", b, trueStr, falseStr, got, want)
			}
		}

		t.Run("True", func(t *testing.T) {
			check(t, true, "yes", "no", "yes")
			check(t, true, "enabled", "disabled", "enabled")
		})

		t.Run("False", func(t *testing.T) {
			check(t, false, "yes", "no", "no")
			check(t, false, "enabled", "disabled", "disabled")
		})
	})
}
