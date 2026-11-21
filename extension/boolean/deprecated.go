// Package boolean keeps the old import path working after the package moved
// to github.com/entiqon/common/boolean.
//
// Deprecated: Use github.com/entiqon/common/boolean instead. This package
// will be removed in v2.0.0.
package boolean

import moved "github.com/entiqon/common/boolean"

// ParseFrom parses value into a bool.
//
// Deprecated: Use boolean.From from github.com/entiqon/common/boolean instead.
func ParseFrom(value any) (bool, error) {
	return moved.From(value)
}

// BoolToString returns trueStr if b is true, otherwise falseStr.
//
// Deprecated: Use boolean.ToString from github.com/entiqon/common/boolean instead.
func BoolToString(b bool, trueStr, falseStr string) string {
	return moved.ToString(b, trueStr, falseStr)
}
