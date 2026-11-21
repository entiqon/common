package testkit

import (
	"errors"
	"testing"
)

// Error fails the test immediately when err is nil.
func Error(
	t testing.TB,
	err error,
) {
	t.Helper()

	if err == nil {
		t.Fatal("expected an error, got nil")
	}
}

// NoError fails the test immediately when err is not nil.
func NoError(
	t testing.TB,
	err error,
) {
	t.Helper()

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
}

// ErrorContains fails the test immediately when err is nil, and reports
// when its message does not contain substr.
func ErrorContains(
	t testing.TB,
	err error,
	substr string,
) {
	t.Helper()

	Error(t, err)
	Contains(t, err.Error(), substr)
}

// ErrorIs reports when err does not match target in its chain.
func ErrorIs(
	t testing.TB,
	err error,
	target error,
) {
	t.Helper()

	if !errors.Is(err, target) {
		t.Errorf("expected error %v in chain, got %v", target, err)
	}
}

// ErrorAs fails the test immediately when err's chain has no error of
// type E, and returns the match so its fields can be checked.
//
//	vErr := testkit.ErrorAs[*ValidationError](t, err)
//	testkit.Equal(t, "email", vErr.Field)
func ErrorAs[E error](
	t testing.TB,
	err error,
) E {
	t.Helper()

	var target E
	if !errors.As(err, &target) {
		t.Fatalf("expected a %T in error chain, got %v", target, err)
	}

	return target
}

// MustValue returns v, failing the test immediately when err is not nil.
// Use it to keep setup code short:
//
//	user := testkit.MustValue(repo.Create(ctx, input))(t)
func MustValue[T any](v T, err error) func(t testing.TB) T {
	return func(t testing.TB) T {
		t.Helper()

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		return v
	}
}
