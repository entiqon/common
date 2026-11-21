package testkit

import (
	"reflect"
	"strings"
	"testing"
)

// Equal reports when expected and actual values differ.
func Equal[T comparable](
	t testing.TB,
	expected T,
	got T,
) {
	t.Helper()

	if expected != got {
		t.Errorf("expected %v, got %v", expected, got)
	}
}

// NotEqual reports when got equals the value it must differ from.
func NotEqual[T comparable](
	t testing.TB,
	notExpected T,
	got T,
) {
	t.Helper()

	if notExpected == got {
		t.Errorf("expected a value other than %v", got)
	}
}

// DeepEqual reports when expected and got are not deeply equal. Use it
// for slices, maps and structs that Equal cannot compare. Multi-line
// strings are shown as a line diff.
func DeepEqual[T any](
	t testing.TB,
	expected T,
	got T,
) {
	t.Helper()

	if reflect.DeepEqual(expected, got) {
		return
	}

	want, wantIsText := any(expected).(string)
	have, haveIsText := any(got).(string)
	if wantIsText && haveIsText && (strings.Contains(want, "\n") || strings.Contains(have, "\n")) {
		t.Errorf("values differ (-expected +got):\n%s", Diff(want, have))
		return
	}

	t.Errorf("values differ\nexpected: %#v\n     got: %#v", expected, got)
}

// True reports when cond is false.
func True(
	t testing.TB,
	cond bool,
) {
	t.Helper()

	if !cond {
		t.Error("expected true, got false")
	}
}

// False reports when cond is true.
func False(
	t testing.TB,
	cond bool,
) {
	t.Helper()

	if cond {
		t.Error("expected false, got true")
	}
}

// Nil reports when v is not nil. A nil pointer, map, slice, channel or
// func stored in an interface counts as nil.
func Nil(
	t testing.TB,
	v any,
) {
	t.Helper()

	if !isNil(v) {
		t.Errorf("expected nil, got %#v", v)
	}
}

// NotNil fails the test immediately when v is nil, since code after it
// usually dereferences v.
func NotNil(
	t testing.TB,
	v any,
) {
	t.Helper()

	if isNil(v) {
		t.Fatal("expected a non-nil value, got nil")
	}
}

// Pointer fails the test immediately when got is nil, and reports when
// the value it points to differs from expected.
func Pointer[T comparable](
	t testing.TB,
	expected T,
	got *T,
) {
	t.Helper()

	if got == nil {
		t.Fatalf("expected %v, got nil", expected)
	}

	Equal(t, expected, *got)
}

// Contains reports when s does not contain substr.
func Contains(
	t testing.TB,
	s string,
	substr string,
) {
	t.Helper()

	if !strings.Contains(s, substr) {
		t.Errorf("expected %q to contain %q", s, substr)
	}
}

// NotContains reports when s contains substr.
func NotContains(
	t testing.TB,
	s string,
	substr string,
) {
	t.Helper()

	if strings.Contains(s, substr) {
		t.Errorf("expected %q not to contain %q", s, substr)
	}
}

// Len reports when v (a string, slice, array, map or channel) does not
// have the expected length.
func Len(
	t testing.TB,
	expected int,
	v any,
) {
	t.Helper()

	rv := reflect.ValueOf(v)
	switch rv.Kind() {
	case reflect.String, reflect.Slice, reflect.Array, reflect.Map, reflect.Chan:
		if rv.Len() != expected {
			t.Errorf("expected length %d, got %d: %#v", expected, rv.Len(), v)
		}
	default:
		t.Fatalf("Len: %T has no length", v)
	}
}

// Empty reports when v is not nil, a zero value or of length 0.
func Empty(
	t testing.TB,
	v any,
) {
	t.Helper()

	if !isEmpty(v) {
		t.Errorf("expected empty, got %#v", v)
	}
}

// Match reports when expected and got do not hold the same
// elements, ignoring order. Duplicates are counted.
func Match[T any](
	t testing.TB,
	expected []T,
	got []T,
) {
	t.Helper()

	used := make([]bool, len(got))
	var missing []T
outer:
	for _, want := range expected {
		for i, have := range got {
			if !used[i] && reflect.DeepEqual(want, have) {
				used[i] = true
				continue outer
			}
		}
		missing = append(missing, want)
	}

	var extra []T
	for i, have := range got {
		if !used[i] {
			extra = append(extra, have)
		}
	}

	if len(missing) > 0 || len(extra) > 0 {
		t.Errorf("elements differ\nmissing: %#v\n  extra: %#v", missing, extra)
	}
}

// Panics reports when fn does not panic, and returns the recovered value.
func Panics(
	t testing.TB,
	fn func(),
) (recovered any) {
	t.Helper()

	panicked := true
	func() {
		defer func() { recovered = recover() }()
		fn()
		panicked = false
	}()

	if !panicked {
		t.Error("expected a panic, function returned normally")
	}

	return recovered
}

func isNil(v any) bool {
	if v == nil {
		return true
	}

	rv := reflect.ValueOf(v)
	switch rv.Kind() {
	case reflect.Chan, reflect.Func, reflect.Interface, reflect.Map, reflect.Pointer, reflect.Slice, reflect.UnsafePointer:
		return rv.IsNil()
	default:
		return false
	}
}

func isEmpty(v any) bool {
	if isNil(v) {
		return true
	}

	rv := reflect.ValueOf(v)
	switch rv.Kind() {
	case reflect.String, reflect.Slice, reflect.Array, reflect.Map, reflect.Chan:
		return rv.Len() == 0
	default:
		return rv.IsZero()
	}
}

// PointerOf returns a pointer to v, for building fixtures with optional fields.
func PointerOf[T any](v T) *T {
	p := &v // two statements on purpose: a bare `return &v` triggers gopls's newexpr hint on Go 1.26+

	return p
}
