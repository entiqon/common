package testkit

import (
	"io"
	"sort"
	"strings"
	"sync"
	"testing"
)

// Cases runs fn as a subtest for each named case, in sorted name order so
// output is stable.
//
//	testkit.Cases(t, map[string]tc{
//		"empty email": {input: "", wantErr: ErrInvalidEmail},
//		"valid email": {input: "a@b.co"},
//	}, func(t *testing.T, c tc) {
//		testkit.ErrorIs(t, ValidateEmail(c.input), c.wantErr)
//	})
func Cases[C any](
	t *testing.T,
	cases map[string]C,
	fn func(t *testing.T, c C),
) {
	t.Helper()

	names := make([]string, 0, len(cases))
	for name := range cases {
		names = append(names, name)
	}
	sort.Strings(names)

	for _, name := range names {
		c := cases[name]
		t.Run(name, func(t *testing.T) { fn(t, c) })
	}
}

// LogWriter returns an io.Writer that sends each write to t.Log, so log output
// only shows for failing tests or with -v. Writes after the test ends are
// dropped instead of panicking. Pass it to any logger:
//
//	logger := log.New(testkit.LogWriter(t), "", 0)
//	logger := slog.New(slog.NewTextHandler(testkit.LogWriter(t), nil)) // Go 1.21+
func LogWriter(t testing.TB) io.Writer {
	w := &testLogWriter{t: t}
	t.Cleanup(w.stop)

	return w
}

type testLogWriter struct {
	mu   sync.Mutex
	t    testing.TB
	done bool
}

func (w *testLogWriter) Write(p []byte) (int, error) {
	w.mu.Lock()
	defer w.mu.Unlock()

	if !w.done {
		w.t.Helper()
		w.t.Log(strings.TrimRight(string(p), "\n"))
	}

	return len(p), nil
}

func (w *testLogWriter) stop() {
	w.mu.Lock()
	defer w.mu.Unlock()

	w.done = true
}
