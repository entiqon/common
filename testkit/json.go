package testkit

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"
	"testing"
)

// JSONEqual reports when expected and got are not the same JSON document.
// Key order and whitespace are ignored.
func JSONEqual(
	t testing.TB,
	expected string,
	got string,
) {
	t.Helper()

	var want, have any
	if err := json.Unmarshal([]byte(expected), &want); err != nil {
		t.Fatalf("expected value is not valid JSON: %v\n%s", err, expected)
	}
	if err := json.Unmarshal([]byte(got), &have); err != nil {
		t.Fatalf("got value is not valid JSON: %v\n%s", err, got)
	}

	if !reflect.DeepEqual(want, have) {
		t.Errorf("JSON differs (-expected +got):\n%s",
			Diff(string(MustJSON(t, want)), string(MustJSON(t, have))))
	}
}

// MustJSON marshals v as indented JSON, failing the test immediately on error.
func MustJSON(
	t testing.TB,
	v any,
) []byte {
	t.Helper()

	b, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		t.Fatalf("MustJSON: %v", err)
	}

	return b
}

// DecodeJSON decodes JSON into a T, failing the test immediately on error.
// src may be a *httptest.ResponseRecorder, *http.Response, io.Reader, string
// or []byte.
//
//	user := testkit.DecodeJSON[UserResponse](t, rec)
func DecodeJSON[T any](
	t testing.TB,
	src any,
) T {
	t.Helper()

	var r io.Reader
	switch x := src.(type) {
	case *httptest.ResponseRecorder:
		r = x.Body

	case *http.Response:
		defer func() {
			_ = x.Body.Close()
		}()
		r = x.Body

	case string:
		r = strings.NewReader(x)

	case []byte:
		r = bytes.NewReader(x)

	case io.Reader:
		r = x

	default:
		t.Fatalf("DecodeJSON: unsupported source %T", src)

	}

	var v T
	if err := json.NewDecoder(r).Decode(&v); err != nil {
		t.Fatalf("DecodeJSON: %v", err)
	}

	return v
}
