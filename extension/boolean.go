package extension

import "github.com/entiqon/common/boolean"

// Boolean parses a value into a bool.
// Returns false if parsing fails.
func Boolean(value any) bool {
	v, err := boolean.From(value)
	if err != nil {
		return false
	}
	return v
}

// BooleanOr parses a value into a bool.
// Returns the provided default if parsing fails.
func BooleanOr(value any, def bool) bool {
	v, err := boolean.From(value)
	if err != nil {
		return def
	}
	return v
}
