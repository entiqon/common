package testkit

import (
	"context"
	"testing"
	"time"
)

// Eventually polls cond every tick until it returns true, and reports
// if it never does within timeout. Use it instead of time.Sleep for async code.
//
//	testkit.Eventually(t, worker.Done, time.Second, 10*time.Millisecond)
func Eventually(
	t testing.TB,
	cond func() bool,
	timeout time.Duration,
	tick time.Duration,
) {
	t.Helper()

	deadline := time.Now().Add(timeout)
	for !cond() {
		if time.Now().After(deadline) {
			t.Errorf("condition not met within %v", timeout)
			return
		}
		time.Sleep(tick)
	}
}

// Receive waits for a value on ch and returns it. It fails the test
// immediately if ch closes or nothing arrives within timeout.
func Receive[T any](
	t testing.TB,
	ch <-chan T,
	timeout time.Duration,
) T {
	t.Helper()

	timer := time.NewTimer(timeout)
	defer timer.Stop()

	var (
		value T
		open  bool
	)

	select {
	case value, open = <-ch:
		if !open {
			t.Fatal("expected a value, channel was closed")
		}
	case <-timer.C:
		t.Fatalf("expected a value within %v, got none", timeout)
	}

	return value
}

// Context returns a context that is cancelled when the test ends or after
// timeout, whichever comes first. A timeout <= 0 means no timeout.
// (On Go 1.24+, t.Context() covers the no-timeout case.)
func Context(
	t testing.TB,
	timeout time.Duration,
) context.Context {
	t.Helper()

	ctx, cancel := context.WithCancel(context.Background())
	if timeout > 0 {
		ctx, cancel = context.WithTimeout(context.Background(), timeout)
	}
	t.Cleanup(cancel)

	return ctx
}
