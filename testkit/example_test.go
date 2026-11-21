package testkit_test

import (
	"fmt"
	"io"
	"io/fs"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"sync/atomic"
	"testing"
	"time"

	"github.com/entiqon/common/testkit"
)

func ExamplePointerOf() {
	score := testkit.PointerOf(80)

	fmt.Println(*score)

	// Output:
	// 80
}

func ExampleDiff() {
	fmt.Print(testkit.Diff("id: 1\nname: Ana", "id: 1\nname: Eva"))

	// Output:
	//   id: 1
	// - name: Ana
	// + name: Eva
}

func ExampleNewFakeClock() {
	clock := testkit.NewFakeClock(time.Time{})

	fmt.Println(clock.Now().Format(time.RFC3339))

	clock.Advance(90 * time.Minute)

	fmt.Println(clock.Now().Format(time.RFC3339))

	// Output:
	// 2024-01-01T00:00:00Z
	// 2024-01-01T01:30:00Z
}

func ExampleRoundTripFunc() {
	client := &http.Client{Transport: testkit.RoundTripFunc(
		func(r *http.Request) (*http.Response, error) {
			return testkit.Response(http.StatusOK, `{"ok":true}`), nil
		},
	)}

	resp, err := client.Get("https://api.example.com/health")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)

	fmt.Println(resp.StatusCode, string(body))

	// Output:
	// 200 {"ok":true}
}

// The examples below show test code. They are compiled with every test run,
// so they always match the API, but they are not executed because they need
// the *testing.T of a real test.

func ExampleEqual() {
	var t *testing.T // the *testing.T of your test

	got := strings.ToUpper("ana")

	testkit.Equal(t, "ANA", got)
	testkit.DeepEqual(t, []string{"a", "b"}, strings.Split("a,b", ","))
	testkit.Match(t, []int{3, 1, 2}, []int{1, 2, 3})
}

func ExampleErrorAs() {
	var t *testing.T // the *testing.T of your test

	_, err := os.Open("missing.txt")

	testkit.ErrorIs(t, err, fs.ErrNotExist)

	pathErr := testkit.ErrorAs[*fs.PathError](t, err)
	testkit.Equal(t, "missing.txt", pathErr.Path)
}

func ExampleMustValue() {
	var t *testing.T // the *testing.T of your test

	// Stops the test if os.ReadFile returns an error.
	data := testkit.MustValue(os.ReadFile("testdata/input.json"))(t)

	testkit.True(t, len(data) > 0)
}

func ExampleServe() {
	var t *testing.T // the *testing.T of your test

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusCreated)
		_, _ = w.Write([]byte(`{"name":"Ana"}`))
	})

	rec := testkit.Serve(t, handler, http.MethodPost, "/users",
		map[string]string{"name": "Ana"},
		"Authorization", "Bearer token")

	testkit.Equal(t, http.StatusCreated, rec.Code)
	testkit.JSONEqual(t, `{"name":"Ana"}`, rec.Body.String())
}

func ExampleEventually() {
	var t *testing.T // the *testing.T of your test

	done := make(chan string, 1)
	go func() { done <- "finished" }()

	testkit.Equal(t, "finished", testkit.Receive(t, done, time.Second))

	var ready atomic.Value
	ready.Store(true)
	testkit.Eventually(t, func() bool { return ready.Load() == true }, time.Second, 10*time.Millisecond)
}

func ExampleEqualFile() {
	var t *testing.T // the *testing.T of your test

	dir := testkit.WriteFiles(t, map[string]string{"in.txt": "hello"})
	out := strings.ToUpper(testkit.ReadFile(t, filepath.Join(dir, "in.txt")))

	// Compares with testdata/upper.txt; run with UPDATE_FILES=1 to create it.
	testkit.EqualFile(t, "upper.txt", []byte(out))
}

func ExampleCases() {
	var t *testing.T // the *testing.T of your test

	testkit.Cases(t, map[string]struct{ in, want string }{
		"lower": {"ana", "ANA"},
		"mixed": {"aNa", "ANA"},
	}, func(t *testing.T, c struct{ in, want string }) {
		testkit.Equal(t, c.want, strings.ToUpper(c.in))
	})
}

func ExampleLogWriter() {
	var t *testing.T // the *testing.T of your test

	// Log output only shows when the test fails or runs with -v.
	logger := log.New(testkit.LogWriter(t), "", 0)
	logger.Printf("processing %d items", 3)
}

func ExamplePasses() {
	var t *testing.T // the *testing.T of your test

	// Your own helper.
	nonEmpty := func(tb testing.TB, s string) {
		tb.Helper()
		if s == "" {
			tb.Fatal("string is empty")
		}
	}

	testkit.Passes(t, func(tb testing.TB) { nonEmpty(tb, "ok") })

	msg := testkit.FailsHard(t, func(tb testing.TB) { nonEmpty(tb, "") })
	testkit.Contains(t, msg, "empty")
}

func ExampleFailsSoft() {
	var t *testing.T // the *testing.T of your test

	msg := testkit.FailsSoft(t, func(tb testing.TB) {
		testkit.Equal(tb, 1, 2)
	})

	testkit.Contains(t, msg, "expected 1, got 2")
}

func ExampleRecord() {
	var t *testing.T // the *testing.T of your test

	r := testkit.Record(t, func(tb testing.TB) {
		tb.Error("first problem")
		tb.Error("second problem")
	})

	testkit.False(t, r.Stopped())
	testkit.Equal(t, "first problem\nsecond problem", r.Message())

	r = testkit.Record(t, func(tb testing.TB) { tb.Skip("no database") })
	testkit.True(t, r.Skipped())
}
