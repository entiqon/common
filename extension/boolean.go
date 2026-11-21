package extension

import "github.com/entiqon/common/boolean"

// Boolean converts value to a bool, returning false when it cannot.
//
// Deprecated: Use boolean.FromOr(value, false) from github.com/entiqon/common/boolean instead.
func Boolean(value any) bool {
	return boolean.FromOr(value, false)
}

// BooleanOr converts value to a bool, returning fallback when it cannot.
//
// Deprecated: Use boolean.FromOr from github.com/entiqon/common/boolean instead.
func BooleanOr(value any, fallback bool) bool {
	return boolean.FromOr(value, fallback)
}
