// Package extension provides convenient conversions for common Go values.
//
// Its shortcut functions suppress parsing errors and return either a zero value
// or a caller-provided fallback:
//
//   - BooleanOr(value any, def bool) bool
//   - NumberOr(value any, def float64) float64
//   - FloatOr(value any, def float64) float64
//   - DecimalOr(value any, def string) decimal.Decimal
//   - Date(value any, orders ...date.Order) time.Time
//   - DateOr(value any, def time.Time, orders ...date.Order) time.Time
//
// Date and DateOr accept an optional date order for ambiguous day-first or
// month-first inputs.
//
// Example:
//
//	send := extension.BooleanOr(c.QueryParam("send"), false)
package extension
