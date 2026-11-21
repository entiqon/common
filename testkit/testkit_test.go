package testkit

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"sync/atomic"
	"testing"
	"time"
)

type validationError struct{ field string }

func (e *validationError) Error() string { return "invalid " + e.field }

func TestValueAssertions(t *testing.T) {
	var nilPtr *int

	Passes(t, func(tb testing.TB) { Equal(tb, 1, 1) })
	FailsSoft(t, func(tb testing.TB) { Equal(tb, 1, 2) })
	Passes(t, func(tb testing.TB) { NotEqual(tb, 1, 2) })
	FailsSoft(t, func(tb testing.TB) { NotEqual(tb, "a", "a") })

	Passes(t, func(tb testing.TB) { DeepEqual(tb, []int{1, 2}, []int{1, 2}) })
	FailsSoft(t, func(tb testing.TB) { DeepEqual(tb, map[string]int{"a": 1}, map[string]int{"a": 2}) })
	msg := FailsSoft(t, func(tb testing.TB) { DeepEqual(tb, "a\nb", "a\nc") })
	Contains(t, msg, "- b")
	Contains(t, msg, "+ c")

	Passes(t, func(tb testing.TB) { True(tb, true) })
	FailsSoft(t, func(tb testing.TB) { True(tb, false) })
	FailsSoft(t, func(tb testing.TB) { False(tb, true) })

	Passes(t, func(tb testing.TB) { Nil(tb, nilPtr) })
	FailsSoft(t, func(tb testing.TB) { Nil(tb, 1) })
	FailsHard(t, func(tb testing.TB) { NotNil(tb, nilPtr) })

	Passes(t, func(tb testing.TB) { Pointer(tb, "x", PointerOf("x")) })
	FailsHard(t, func(tb testing.TB) { Pointer(tb, 1, nilPtr) })
	FailsSoft(t, func(tb testing.TB) { Pointer(tb, 1, PointerOf(2)) })

	Passes(t, func(tb testing.TB) { Contains(tb, "hello", "ell") })
	FailsSoft(t, func(tb testing.TB) { Contains(tb, "hello", "xyz") })
	FailsSoft(t, func(tb testing.TB) { NotContains(tb, "hello", "ell") })

	Passes(t, func(tb testing.TB) { Len(tb, 2, []int{1, 2}) })
	FailsSoft(t, func(tb testing.TB) { Len(tb, 1, "ab") })
	FailsHard(t, func(tb testing.TB) { Len(tb, 0, 5) })

	Passes(t, func(tb testing.TB) {
		Empty(tb, "")
		Empty(tb, []int{})
		Empty(tb, nil)
		Empty(tb, 0)
	})
	FailsSoft(t, func(tb testing.TB) { Empty(tb, []int{1}) })

	Passes(t, func(tb testing.TB) { Match(tb, []int{1, 1, 2}, []int{2, 1, 1}) })
	msg = FailsSoft(t, func(tb testing.TB) { Match(tb, []int{1, 3}, []int{1, 2}) })
	Contains(t, msg, "missing: []int{3}")

	Passes(t, func(tb testing.TB) {
		DeepEqual[any](tb, "boom", Panics(tb, func() { panic("boom") }))
	})
	FailsSoft(t, func(tb testing.TB) { Panics(tb, func() {}) })
}

func TestErrorAssertions(t *testing.T) {
	base := errors.New("not found")
	wrapped := fmt.Errorf("load lead: %w", base)

	Passes(t, func(tb testing.TB) { Error(tb, base) })
	FailsHard(t, func(tb testing.TB) { Error(tb, nil) })
	Passes(t, func(tb testing.TB) { NoError(tb, nil) })
	FailsHard(t, func(tb testing.TB) { NoError(tb, base) })

	Passes(t, func(tb testing.TB) { ErrorContains(tb, wrapped, "not found") })
	FailsHard(t, func(tb testing.TB) { ErrorContains(tb, nil, "x") })
	FailsSoft(t, func(tb testing.TB) { ErrorContains(tb, wrapped, "timeout") })

	Passes(t, func(tb testing.TB) { ErrorIs(tb, wrapped, base) })
	FailsSoft(t, func(tb testing.TB) { ErrorIs(tb, errors.New("other"), base) })

	Passes(t, func(tb testing.TB) {
		vErr := ErrorAs[*validationError](tb, fmt.Errorf("create: %w", &validationError{"email"}))
		Equal(tb, "email", vErr.field)
	})
	FailsHard(t, func(tb testing.TB) { ErrorAs[*validationError](tb, base) })

	Passes(t, func(tb testing.TB) { Equal(tb, 7, MustValue(7, nil)(tb)) })
	FailsHard(t, func(tb testing.TB) { MustValue(0, base)(tb) })
}

func TestJSON(t *testing.T) {
	Passes(t, func(tb testing.TB) { JSONEqual(tb, `{"a":1,"b":[1,2]}`, `{ "b":[1,2], "a":1 }`) })
	FailsSoft(t, func(tb testing.TB) { JSONEqual(tb, `{"a":1}`, `{"a":2}`) })
	FailsHard(t, func(tb testing.TB) { JSONEqual(tb, `{"a":1}`, `{`) })

	type lead struct {
		Name string `json:"name"`
	}
	Equal(t, "Ana", DecodeJSON[lead](t, `{"name":"Ana"}`).Name)
	FailsHard(t, func(tb testing.TB) { DecodeJSON[lead](tb, `{`) })
}

func TestHTTP(t *testing.T) {
	type lead struct {
		Name string `json:"name"`
	}

	h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("Content-Type") != "application/json" || r.Header.Get("Authorization") != "Bearer x" {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		in := DecodeJSON[lead](t, r.Body)
		w.WriteHeader(http.StatusCreated)
		_, _ = w.Write(MustJSON(t, lead{Name: strings.ToUpper(in.Name)}))
	})

	rec := Serve(t, h, http.MethodPost, "/leads", lead{Name: "ana"}, "Authorization", "Bearer x")
	Equal(t, http.StatusCreated, rec.Code)
	Equal(t, "ANA", DecodeJSON[lead](t, rec).Name)

	srv := Server(t, h)
	resp := MustValue(http.Get(srv.URL))(t)
	_ = resp.Body.Close()
	Equal(t, http.StatusBadRequest, resp.StatusCode)

	client := &http.Client{Transport: RoundTripFunc(func(*http.Request) (*http.Response, error) {
		return Response(http.StatusOK, `{"name":"stub"}`), nil
	})}
	resp = MustValue(client.Get("http://crm.invalid/leads/1"))(t)
	Equal(t, "stub", DecodeJSON[lead](t, resp).Name)
}

func TestAsync(t *testing.T) {
	var n int32
	Eventually(t, func() bool { return atomic.AddInt32(&n, 1) >= 3 }, time.Second, time.Millisecond)
	FailsSoft(t, func(tb testing.TB) {
		Eventually(tb, func() bool { return false }, 20*time.Millisecond, 5*time.Millisecond)
	})

	ch := make(chan int, 1)
	ch <- 5
	Equal(t, 5, Receive(t, ch, time.Second))
	FailsHard(t, func(tb testing.TB) { Receive(tb, ch, 10*time.Millisecond) })
	close(ch)
	FailsHard(t, func(tb testing.TB) { Receive(tb, ch, time.Second) })

	ctx := Context(t, time.Millisecond)
	<-ctx.Done()
}

func TestFiles(t *testing.T) {
	dir := WriteFiles(t, map[string]string{"a.txt": "A", "sub/b.txt": "B"})
	Equal(t, "B", ReadFile(t, filepath.Join(dir, "sub", "b.txt")))

	wd := MustValue(os.Getwd())(t)
	NoError(t, os.Chdir(t.TempDir()))
	t.Cleanup(func() { _ = os.Chdir(wd) })

	t.Setenv("UPDATE_FILES", "1")
	EqualFile(t, "out/lead.json", []byte(`{"id":1}`))
	t.Setenv("UPDATE_FILES", "0")
	EqualFile(t, "out/lead.json", []byte(`{"id":1}`))
	FailsSoft(t, func(tb testing.TB) { EqualFile(tb, "out/lead.json", []byte(`{"id":2}`)) })
	FailsHard(t, func(tb testing.TB) { EqualFile(tb, "nope.txt", nil) })
}

func TestFakeClock(t *testing.T) {
	clock := NewFakeClock(time.Time{})
	start := clock.Now()
	Equal(t, 2024, start.Year())
	Equal(t, start.Add(time.Hour), clock.Advance(time.Hour))
	clock.Set(start)
	Equal(t, start, clock.Now())
}

func TestCasesAndLogger(t *testing.T) {
	var order []string
	Cases(t, map[string]string{"b": "B", "a": "A"}, func(t *testing.T, want string) {
		order = append(order, want)
		log.New(LogWriter(t), "", 0).Printf("running case %s", want)
	})
	DeepEqual(t, []string{"A", "B"}, order)
}

func TestDiff(t *testing.T) {
	Equal(t, "", Diff("same", "same"))
	Equal(t, "  a\n- b\n+ c\n", Diff("a\nb", "a\nc"))
}

func TestRecorder(t *testing.T) {
	r := Record(t, func(tb testing.TB) {
		tb.Error("first")
		tb.Errorf("second %d", 2)
		tb.Fail()
		tb.Log("goes to the real test log")
	})
	True(t, r.Failed())
	False(t, r.Stopped())
	Equal(t, "first\nsecond 2", r.Message())

	r = Record(t, func(tb testing.TB) { tb.FailNow() })
	True(t, r.Stopped())
	Equal(t, "", r.Message())

	for name, skip := range map[string]func(tb testing.TB){
		"Skip":    func(tb testing.TB) { tb.Skip("reason") },
		"Skipf":   func(tb testing.TB) { tb.Skipf("reason %d", 1) },
		"SkipNow": func(tb testing.TB) { tb.SkipNow() },
	} {
		skip := skip
		t.Run(name, func(t *testing.T) {
			r := Record(t, skip)
			True(t, r.Skipped())
			False(t, r.Failed())
		})
	}

	// The checks themselves report when fn does not behave as expected.
	FailsSoft(t, func(tb testing.TB) { Passes(tb, func(inner testing.TB) { inner.Error("x") }) })
	FailsSoft(t, func(tb testing.TB) { FailsSoft(tb, func(testing.TB) {}) })
	FailsSoft(t, func(tb testing.TB) { FailsSoft(tb, func(inner testing.TB) { inner.FailNow() }) })
	FailsSoft(t, func(tb testing.TB) { FailsHard(tb, func(inner testing.TB) { inner.Error("x") }) })
}

func TestErrorPaths(t *testing.T) {
	// A path below a regular file cannot be created.
	blocker := filepath.Join(WriteFiles(t, map[string]string{"file": "x"}), "file")

	FailsHard(t, func(tb testing.TB) { ReadFile(tb, filepath.Join(blocker, "missing")) })

	wd := MustValue(os.Getwd())(t)
	NoError(t, os.Chdir(filepath.Dir(blocker)))
	t.Cleanup(func() { _ = os.Chdir(wd) })

	t.Setenv("UPDATE_FILES", "1")
	FailsHard(t, func(tb testing.TB) { EqualFile(tb, "../file/out.txt", nil) })
	t.Setenv("UPDATE_FILES", "0")

	FailsHard(t, func(tb testing.TB) { JSONEqual(tb, `{`, `{}`) })
	FailsHard(t, func(tb testing.TB) { MustJSON(tb, make(chan int)) })
	FailsHard(t, func(tb testing.TB) { DecodeJSON[int](tb, 42) })
	Equal(t, 7, DecodeJSON[int](t, []byte("7")))
	Equal(t, 8, DecodeJSON[int](t, strings.NewReader("8")))

	h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body := MustValue(io.ReadAll(r.Body))(t)
		_, _ = w.Write(body)
	})
	FailsHard(t, func(tb testing.TB) { Serve(tb, h, http.MethodGet, "/", nil, "only-key") })
	Equal(t, "text", Serve(t, h, http.MethodPost, "/", "text").Body.String())
	Equal(t, "bytes", Serve(t, h, http.MethodPost, "/", []byte("bytes")).Body.String())
	Equal(t, "reader", Serve(t, h, http.MethodPost, "/", strings.NewReader("reader")).Body.String())

	resp := Response(http.StatusAccepted, "", "X-Request-ID", "abc")
	Equal(t, "", Serve(t, h, http.MethodGet, "/", nil).Body.String())
	Equal(t, "abc", resp.Header.Get("X-Request-ID"))
}

func TestWriteFilesError(t *testing.T) {
	// A NUL byte makes the directory invalid; "." makes the file path a directory.
	FailsHard(t, func(tb testing.TB) { WriteFiles(tb, map[string]string{"bad\x00dir/file": "x"}) })
	FailsHard(t, func(tb testing.TB) { WriteFiles(tb, map[string]string{".": "x"}) })
}

func TestEqualFileUpdateError(t *testing.T) {
	wd := MustValue(os.Getwd())(t)
	NoError(t, os.Chdir(t.TempDir()))
	t.Cleanup(func() { _ = os.Chdir(wd) })
	t.Setenv("UPDATE_FILES", "1")

	FailsHard(t, func(tb testing.TB) { EqualFile(tb, "bad\x00dir/out.txt", nil) })

	EqualFile(t, "dir/out.txt", []byte("x"))
	FailsHard(t, func(tb testing.TB) { EqualFile(tb, "dir", []byte("x")) })
}
