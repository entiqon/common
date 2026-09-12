package test

import "testing"

// AssertEqual reports when expected and actual values differ.
func AssertEqual[T comparable](
	t *testing.T,
	expected T,
	got T,
) {
	t.Helper()

	if expected != got {
		t.Errorf("expected %v, got %v", expected, got)
	}
}
