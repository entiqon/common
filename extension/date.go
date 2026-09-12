package extension

import (
	"time"

	"github.com/entiqon/common/date"
)

// Date parses value into a time.Time.
// It returns zero time if parsing fails.
func Date(value any, orders ...date.Order) time.Time {
	result, err := date.From(value, orders...)
	if err != nil {
		return time.Time{}
	}

	return result
}

// DateOr parses value into a time.Time.
// It returns def if parsing fails.
func DateOr(
	value any,
	def time.Time,
	orders ...date.Order,
) time.Time {
	result, err := date.From(value, orders...)
	if err != nil {
		return def
	}

	return result
}
