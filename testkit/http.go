package testkit

import (
	"bytes"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

// Serve sends a request straight to h, with no network, and returns the
// recorded response. body may be nil, a string, []byte or io.Reader; any other
// value is sent as JSON with Content-Type: application/json. headers are
// key/value pairs.
//
//	rec := testkit.Serve(t, h, http.MethodPost, "/leads", CreateLead{Name: "Ana"},
//		"Authorization", "Bearer token")
//	testkit.Equal(t, http.StatusCreated, rec.Code)
func Serve(
	t testing.TB,
	h http.Handler,
	method string,
	target string,
	body any,
	headers ...string,
) *httptest.ResponseRecorder {
	t.Helper()

	if len(headers)%2 != 0 {
		t.Fatal("Serve: headers must be key/value pairs")
	}

	var r io.Reader
	isJSON := false
	switch b := body.(type) {
	case nil:
	case string:
		r = strings.NewReader(b)
	case []byte:
		r = bytes.NewReader(b)
	case io.Reader:
		r = b
	default:
		r = bytes.NewReader(MustJSON(t, b))
		isJSON = true
	}

	req := httptest.NewRequest(method, target, r)
	if isJSON {
		req.Header.Set("Content-Type", "application/json")
	}
	for i := 0; i < len(headers); i += 2 {
		req.Header.Set(headers[i], headers[i+1])
	}

	rec := httptest.NewRecorder()
	h.ServeHTTP(rec, req)

	return rec
}

// Server starts an httptest.Server for h and closes it when the test ends.
func Server(
	t testing.TB,
	h http.Handler,
) *httptest.Server {
	t.Helper()

	srv := httptest.NewServer(h)
	t.Cleanup(srv.Close)

	return srv
}

// RoundTripFunc lets a function act as an http.RoundTripper, to stub calls to
// outside services without a server:
//
//	client := &http.Client{Transport: testkit.RoundTripFunc(
//		func(r *http.Request) (*http.Response, error) {
//			return testkit.Response(http.StatusOK, `{"ok":true}`), nil
//		},
//	)}
type RoundTripFunc func(*http.Request) (*http.Response, error)

// RoundTrip calls f.
func (f RoundTripFunc) RoundTrip(r *http.Request) (*http.Response, error) {
	return f(r)
}

// Response builds a minimal *http.Response with the given status and body.
// headers are key/value pairs.
func Response(
	status int,
	body string,
	headers ...string,
) *http.Response {
	h := make(http.Header)
	for i := 0; i+1 < len(headers); i += 2 {
		h.Set(headers[i], headers[i+1])
	}

	return &http.Response{
		StatusCode:    status,
		Status:        http.StatusText(status),
		Header:        h,
		Body:          io.NopCloser(strings.NewReader(body)),
		ContentLength: int64(len(body)),
		Proto:         "HTTP/1.1",
		ProtoMajor:    1,
		ProtoMinor:    1,
	}
}
