# testkit

<h6 align="left">
  Part of the
  <a href="https://github.com/entiqon/entiqon">entiqon</a>::<a href="https://github.com/entiqon/common">common</a>
  toolkit.
</h6>

The `testkit` package provides dependency-free helpers for writing Go tests: assertions, error checks, JSON and HTTP helpers, polling for asynchronous code, temporary files, expected-output files, a fake clock, and table tests.

The package requires **Go 1.18 or later** because it uses generics.

## Installation

```bash
go get github.com/entiqon/common/testkit
```

Import it only from `_test.go` files, so it never ends up in production builds.

## Conventions

- Arguments go **expected first, then got**: `testkit.Equal(t, want, got)`.
- **Value checks** report with `t.Errorf` and let the test keep going, so one run shows every mismatch.
- **Checks whose failure makes the rest of the test meaningless** (a missing error, a nil pointer, a helper that must return a value) stop the test with `t.Fatal`.
- Every helper takes `testing.TB`, so it works in tests, benchmarks, and fuzz tests.

## Usage

### Values and errors

```go
func TestCreateLead(t *testing.T) {
lead, err := service.Create(ctx, input)
testkit.NoError(t, err)

testkit.Equal(t, "Ana", lead.Name)
testkit.Pointer(t, 80, lead.Score)
testkit.DeepEqual(t, []string{"new", "web"}, lead.Tags)
}

func TestCreateLeadValidation(t *testing.T) {
_, err := service.Create(ctx, Input{})

testkit.ErrorIs(t, err, ErrInvalidInput)

vErr := testkit.ErrorAs[*ValidationError](t, err)
testkit.Equal(t, "email", vErr.Field)
}
```

### HTTP handlers and clients

```go
func TestCreateLeadHandler(t *testing.T) {
rec := testkit.Serve(t, handler, http.MethodPost, "/leads",
CreateLead{Name: "Ana"},
"Authorization", "Bearer token")

testkit.Equal(t, http.StatusCreated, rec.Code)

lead := testkit.DecodeJSON[LeadResponse](t, rec)
testkit.Equal(t, "Ana", lead.Name)
}

func TestCRMClient(t *testing.T) {
client := &http.Client{Transport: testkit.RoundTripFunc(
func(r *http.Request) (*http.Response, error) {
return testkit.Response(http.StatusOK, `{"id":1}`), nil
},
)}

// pass client to the code under test...
}
```

### Asynchronous code

```go
func TestWorker(t *testing.T) {
ctx := testkit.Context(t, 5*time.Second)
results := worker.Start(ctx)

got := testkit.Receive(t, results, time.Second)
testkit.Equal(t, "done", got)

testkit.Eventually(t, worker.Idle, time.Second, 10*time.Millisecond)
}
```

### Files

```go
func TestExport(t *testing.T) {
dir := testkit.WriteFiles(t, map[string]string{
"leads.json": `[{"name":"Ana"}]`,
})

out := export.CSV(filepath.Join(dir, "leads.json"))

testkit.EqualFile(t, "export/leads.csv", out)
}
```

`EqualFile` compares output with `testdata/<name>`. Create or refresh the expected files with:

```bash
UPDATE_FILES=1 go test ./...
```

### Table tests and logs

```go
func TestValidateEmail(t *testing.T) {
testkit.Cases(t, map[string]struct {
input   string
wantErr error
}{
"empty": {"", ErrInvalidEmail},
"valid": {"ana@example.com", nil},
}, func(t *testing.T, c struct {
input   string
wantErr error
}) {
testkit.ErrorIs(t, ValidateEmail(c.input), c.wantErr)
})
}

func TestWithLogs(t *testing.T) {
logger := log.New(testkit.LogWriter(t), "", 0)
// logs only show when the test fails or with -v
}
```

### Testing your own helpers

Use `Passes`, `FailsSoft`, `FailsHard`, and `Record` only when you write **your own test helpers**: functions that take `testing.TB` and check something. Many tests trust such a helper, so a bug in it, such as never failing, makes all of them pass silently. Test the helper like any other code. Ordinary tests of your application use the regular checks instead.

Each function runs the helper against a `Recorder` in place of the real test, so its failures are captured instead of failing your test.

#### Check that a helper passes and fails when it should

```go
func EqualMoney(t testing.TB, want, got Money) {
	t.Helper()

	if want.Currency != got.Currency {
		t.Fatalf("currency differs: expected %s, got %s", want.Currency, got.Currency)
	}
	if want.Cents != got.Cents {
		t.Errorf("amount differs: expected %d, got %d", want.Cents, got.Cents)
	}
}

func TestEqualMoney(t *testing.T) {
	testkit.Passes(t, func(tb testing.TB) {
		EqualMoney(tb, Money{"USD", 1999}, Money{"USD", 1999})
	})

	// A different currency makes the rest meaningless, so the helper must stop.
	msg := testkit.FailsHard(t, func(tb testing.TB) {
		EqualMoney(tb, Money{"USD", 1999}, Money{"EUR", 1999})
	})
	testkit.Contains(t, msg, "currency differs")

	// A different amount is an ordinary mismatch, so the test keeps going.
	testkit.FailsSoft(t, func(tb testing.TB) {
		EqualMoney(tb, Money{"USD", 1999}, Money{"USD", 2000})
	})
}
```

`FailsSoft` and `FailsHard` also verify *how* a helper fails. A helper that wrongly uses `Fatal` hides every mismatch after the first one; a helper that wrongly uses `Errorf` lets the next line run with bad data and panic.

#### Check that a helper reports every problem

`Record` returns the `Recorder`, so you can inspect all recorded messages:

```go
func AssertValidUser(t testing.TB, u User) {
	t.Helper()

	if u.ID == "" {
		t.Error("user has no ID")
	}
	if !strings.Contains(u.Email, "@") {
		t.Errorf("user has invalid email %q", u.Email)
	}
}

func TestAssertValidUser(t *testing.T) {
	r := testkit.Record(t, func(tb testing.TB) {
		AssertValidUser(tb, User{})
	})

	testkit.False(t, r.Stopped())
	testkit.Contains(t, r.Message(), "no ID")
	testkit.Contains(t, r.Message(), "invalid email")
}
```

#### Check that a helper skips

Calling a skipping helper directly would skip your own test. Through `Record`, the skip is captured instead:

```go
func RequireDatabase(t testing.TB) string {
	t.Helper()

	dsn := os.Getenv("TEST_DATABASE_DSN")
	if dsn == "" {
		t.Skip("TEST_DATABASE_DSN not set")
	}

	return dsn
}

func TestRequireDatabaseSkips(t *testing.T) {
	t.Setenv("TEST_DATABASE_DSN", "")

	r := testkit.Record(t, func(tb testing.TB) { RequireDatabase(tb) })

	testkit.True(t, r.Skipped())
	testkit.False(t, r.Failed())
}
```

The `Recorder` captures failures and skips only. Other calls, such as `Log`, `Cleanup`, `TempDir`, and `Setenv`, go to the real test, so helpers that use them still work.

## API

### Assertions

| Function                           | Fails with         | Description                                                      |
|------------------------------------|--------------------|------------------------------------------------------------------|
| `Equal(t, want, got)`              | `Errorf`           | Values differ (comparable types).                                |
| `NotEqual(t, notWant, got)`        | `Errorf`           | Values are equal.                                                |
| `DeepEqual(t, want, got)`          | `Errorf`           | Slices, maps, or structs differ; multi-line strings show a diff. |
| `True(t, cond)` / `False(t, cond)` | `Errorf`           | Condition has the wrong value.                                   |
| `Nil(t, v)`                        | `Errorf`           | Value is not nil (typed nils count as nil).                      |
| `NotNil(t, v)`                     | `Fatal`            | Value is nil.                                                    |
| `Pointer(t, want, got)`            | `Fatal` / `Errorf` | Pointer is nil, or points to a different value.                  |
| `Contains(t, s, substr)`           | `Errorf`           | String does not contain the substring.                           |
| `NotContains(t, s, substr)`        | `Errorf`           | String contains the substring.                                   |
| `Len(t, n, v)`                     | `Errorf`           | String, slice, array, map, or channel has a different length.    |
| `Empty(t, v)`                      | `Errorf`           | Value is not nil, zero, or empty.                                |
| `Match(t, want, got)`              | `Errorf`           | Slices do not hold the same elements, ignoring order.            |
| `Panics(t, fn)`                    | `Errorf`           | Function does not panic; returns the recovered value.            |

### Errors

| Function                        | Fails with         | Description                                           |
|---------------------------------|--------------------|-------------------------------------------------------|
| `Error(t, err)`                 | `Fatal`            | Error is nil.                                         |
| `NoError(t, err)`               | `Fatal`            | Error is not nil.                                     |
| `ErrorContains(t, err, substr)` | `Fatal` / `Errorf` | Error is nil, or its message lacks the substring.     |
| `ErrorIs(t, err, target)`       | `Errorf`           | Target is not in the error chain.                     |
| `ErrorAs[E](t, err) E`          | `Fatal`            | No error of type `E` in the chain; returns the match. |
| `MustValue(v, err)(t) T`        | `Fatal`            | Returns `v`, or stops the test if `err` is not nil.   |

### JSON, HTTP, async, files, and time

| Function                                        | Description                                                        |
|-------------------------------------------------|--------------------------------------------------------------------|
| `JSONEqual(t, want, got)`                       | JSON documents differ, ignoring key order and whitespace.          |
| `DecodeJSON[T](t, src) T`                       | Decodes a recorder, response, reader, string, or bytes.            |
| `MustJSON(t, v) []byte`                         | Marshals `v` as indented JSON.                                     |
| `Serve(t, h, method, target, body, headers...)` | Calls a handler directly and returns the recorded response.        |
| `Server(t, h)`                                  | Starts a test server that closes when the test ends.               |
| `RoundTripFunc` / `Response(status, body)`      | Stub outgoing HTTP calls without a server.                         |
| `Eventually(t, cond, timeout, tick)`            | Polls until a condition is true.                                   |
| `Receive(t, ch, timeout) T`                     | Waits for a value on a channel.                                    |
| `Context(t, timeout)`                           | Context cancelled when the test ends or the timeout passes.        |
| `WriteFiles(t, files) string`                   | Creates a temporary directory with the given files.                |
| `ReadFile(t, path) string`                      | Reads a file.                                                      |
| `EqualFile(t, name, got)`                       | Output differs from `testdata/<name>`.                             |
| `NewFakeClock(start) *FakeClock`                | Manually controlled clock: `Now`, `Advance`, `Set`.                |
| `Cases(t, cases, fn)`                           | Runs named table cases as subtests in sorted order.                |
| `LogWriter(t) io.Writer`                        | Sends log output to `t.Log`; works with `log`, `slog`, and others. |
| `PointerOf(v) *T`                               | Returns a pointer to `v` for building fixtures.                    |
| `Diff(want, got) string`                        | Line diff used in failure messages.                                |

### Testing helpers

| Function                  | Description                                                                  |
|---------------------------|------------------------------------------------------------------------------|
| `Passes(t, fn)`           | Reports when `fn` records any failure.                                       |
| `FailsSoft(t, fn) string` | Reports unless `fn` fails and keeps running (`Errorf`); returns the message. |
| `FailsHard(t, fn) string` | Reports unless `fn` fails and stops (`Fatal`); returns the message.          |
| `Record(t, fn) *Recorder` | Runs `fn` with a `Recorder`: `Failed`, `Stopped`, `Skipped`, `Message`.      |

## License

This package is part of the **Entiqon Common** library and is distributed under the MIT License.