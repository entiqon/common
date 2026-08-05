package text

// Ellipsis truncates a string to max characters and appends "..." if
// truncation occurs.
//
// The ellipsis is appended after truncation and is not included in max.
//
// If max is less than or equal to 0, an empty string is returned.
// If the string is shorter than or equal to max characters, it is returned
// unchanged.
func Ellipsis(value string, max int) string {
	if max <= 0 {
		return ""
	}

	runes := []rune(value)
	if len(runes) <= max {
		return value
	}

	return Truncate(value, max) + "..."
}
