// Package boolean provides helpers for parsing and formatting boolean values.
//
// From parses a bool, string, integer, unsigned integer, or floating-point
// value into a bool:
//   - bool: returned as is.
//   - string: "true", "1", "yes", "y", "t", and "on" are true; "false", "0",
//     "no", "n", "f", and "off" are false. Matching ignores case and
//     surrounding whitespace.
//   - integers, unsigned integers, and floating-point numbers: zero is false,
//     anything else is true.
//
// Any other input, including nil and unknown strings, returns an error.
// FromOr parses the same inputs but returns a fallback instead of an error.
//
// ToString formats a bool as one of two custom strings.
//
// For example:
//
//	ok, err := boolean.From("yes")                         // true, nil
//	debug := boolean.FromOr(os.Getenv("DEBUG"), false)     // false if unset or invalid
//	label := boolean.ToString(false, "enabled", "disabled") // "disabled"
package boolean
