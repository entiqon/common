// Package date provides utilities for parsing and formatting time.Time values.
//
// From accepts strings, byte slices, time values, pointers, and numeric Unix
// timestamps. Date-only values are normalized to midnight UTC, while inputs
// containing a timezone preserve it.
//
// Numeric dates use YearFirst by default. DayFirst or MonthFirst can be passed
// when the year is not the first component.
//
// Main functions:
//
//   - From(value any, orders ...Order) (time.Time, error)
//   - ParseAndFormat(value, layout string) string
//   - Cleaning(raw string) string
//
// Examples:
//
//	t, err := date.From("2026-08-31")
//	// t = 2026-08-31 00:00:00 +0000 UTC
//
//	t, err = date.From("31-08-2026", date.DayFirst)
//	t, err = date.From("08-31-2026", date.MonthFirst)
//
//	formatted := date.ParseAndFormat("20260831", "2006-01-02")
//	// formatted = "2026-08-31"
package date
