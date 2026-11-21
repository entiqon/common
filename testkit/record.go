package testkit

import (
	"fmt"
	"runtime"
	"strings"
	"sync"
	"testing"
)

// Recorder is a testing.TB that records failures instead of reporting them.
// Use it to test your own test helpers: Passes, FailsSoft and FailsHard cover
// the common cases, and Record gives you the Recorder for anything else.
//
// Error, Errorf and Fail are recorded and execution continues. Fatal, Fatalf
// and FailNow are recorded and stop the helper, like the real ones. Skip,
// Skipf and SkipNow are recorded and stop the helper. Every other method,
// such as Log, Cleanup and TempDir, goes to the test that created it.
type Recorder struct {
	testing.TB

	mu       sync.Mutex
	failed   bool
	stopped  bool
	skipped  bool
	messages []string
}

// Record runs fn with a Recorder in place of t and returns it once fn has
// finished or stopped.
//
//	r := testkit.Record(t, func(tb testing.TB) { MyHelper(tb, input) })
//	testkit.True(t, r.Failed())
func Record(
	t testing.TB,
	fn func(tb testing.TB),
) *Recorder {
	t.Helper()

	r := &Recorder{TB: t}
	done := make(chan struct{})

	go func() {
		defer close(done)
		fn(r)
	}()
	<-done

	return r
}

// Passes reports when fn records any failure.
//
//	testkit.Passes(t, func(tb testing.TB) { testkit.Equal(tb, 1, 1) })
func Passes(
	t testing.TB,
	fn func(tb testing.TB),
) {
	t.Helper()

	if r := Record(t, fn); r.Failed() {
		t.Errorf("expected no failure, got: %s", r.Message())
	}
}

// FailsSoft reports unless fn records a failure that lets it keep running
// (Error, Errorf or Fail). It returns the failure message.
//
//	msg := testkit.FailsSoft(t, func(tb testing.TB) { testkit.Equal(tb, 1, 2) })
//	testkit.Contains(t, msg, "expected 1, got 2")
func FailsSoft(
	t testing.TB,
	fn func(tb testing.TB),
) string {
	t.Helper()

	r := Record(t, fn)
	if !r.Failed() || r.Stopped() {
		t.Errorf("expected a failure that continues (Errorf), got failed=%v stopped=%v", r.Failed(), r.Stopped())
	}

	return r.Message()
}

// FailsHard reports unless fn records a failure that stops it (Fatal, Fatalf
// or FailNow). It returns the failure message.
//
//	msg := testkit.FailsHard(t, func(tb testing.TB) { testkit.NoError(tb, err) })
func FailsHard(
	t testing.TB,
	fn func(tb testing.TB),
) string {
	t.Helper()

	r := Record(t, fn)
	if !r.Stopped() {
		t.Errorf("expected a failure that stops (Fatal), got failed=%v", r.Failed())
	}

	return r.Message()
}

// Failed reports whether any failure was recorded.
func (r *Recorder) Failed() bool {
	r.mu.Lock()
	defer r.mu.Unlock()

	return r.failed
}

// Stopped reports whether Fatal, Fatalf or FailNow was called.
func (r *Recorder) Stopped() bool {
	r.mu.Lock()
	defer r.mu.Unlock()

	return r.stopped
}

// Skipped reports whether Skip, Skipf or SkipNow was called.
func (r *Recorder) Skipped() bool {
	r.mu.Lock()
	defer r.mu.Unlock()

	return r.skipped
}

// Message returns every recorded failure message, one per line.
func (r *Recorder) Message() string {
	r.mu.Lock()
	defer r.mu.Unlock()

	return strings.Join(r.messages, "\n")
}

// Helper does nothing: recorded failures have no file and line to adjust.
func (r *Recorder) Helper() {}

// Error records a failure and continues.
func (r *Recorder) Error(args ...any) { r.record(false, fmt.Sprint(args...)) }

// Errorf records a failure and continues.
func (r *Recorder) Errorf(format string, args ...any) {
	r.record(false, fmt.Sprintf(format, args...))
}

// Fail records a failure and continues.
func (r *Recorder) Fail() { r.record(false, "") }

// Fatal records a failure and stops the helper.
func (r *Recorder) Fatal(args ...any) {
	r.record(true, fmt.Sprint(args...))
	runtime.Goexit()
}

// Fatalf records a failure and stops the helper.
func (r *Recorder) Fatalf(format string, args ...any) {
	r.record(true, fmt.Sprintf(format, args...))
	runtime.Goexit()
}

// FailNow records a failure and stops the helper.
func (r *Recorder) FailNow() {
	r.record(true, "")
	runtime.Goexit()
}

// Skip records a skip and stops the helper.
func (r *Recorder) Skip(...any) {
	r.skip()
}

// Skipf records a skip and stops the helper.
func (r *Recorder) Skipf(string, ...any) {
	r.skip()
}

// SkipNow records a skip and stops the helper.
func (r *Recorder) SkipNow() {
	r.skip()
}

func (r *Recorder) record(
	stop bool,
	msg string,
) {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.failed = true
	if stop {
		r.stopped = true
	}
	if msg != "" {
		r.messages = append(r.messages, msg)
	}
}

func (r *Recorder) skip() {
	r.mu.Lock()
	r.skipped = true
	r.mu.Unlock()

	runtime.Goexit()
}
