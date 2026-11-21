// Package testkit provides dependency-free helpers for writing Go tests:
// assertions, error checks, JSON and HTTP helpers, polling for async code,
// temporary files, expected-output files, a fake clock, table tests, and
// helpers for testing your own test helpers.
//
// Import it only from _test.go files, so it never ends up in production builds.
//
// Conventions:
//   - Arguments go expected first, then got: Equal(t, want, got).
//   - Value checks report with t.Errorf and let the test keep going, so one
//     run shows every mismatch.
//   - Checks whose failure makes the rest of the test meaningless (a missing
//     error, a nil pointer, a helper that must return a value) stop the test
//     with t.Fatal.
package testkit
