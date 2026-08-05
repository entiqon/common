package text

import "strings"

// PadLeft prepends pad to value until it reaches the specified length.
//
// If value is already at least length bytes long, it is returned unchanged.
func PadLeft(value string, length int, pad rune) string {
	if len(value) >= length {
		return value
	}

	return padString(value, length-len(value), 0, pad)
}

// PadRight appends pad to value until it reaches the specified length.
//
// If value is already at least length bytes long, it is returned unchanged.
func PadRight(value string, length int, pad rune) string {
	if len(value) >= length {
		return value
	}

	return padString(value, 0, length-len(value), pad)
}

// PadCenter prepends and appends pad to value until it reaches the specified
// length.
//
// If an odd number of padding characters is required, the extra character is
// appended to the right.
//
// If value is already at least length bytes long, it is returned unchanged.
func PadCenter(value string, length int, pad rune) string {
	if len(value) >= length {
		return value
	}

	padding := length - len(value)
	left := padding / 2
	right := padding - left

	return padString(value, left, right, pad)
}

func padString(
	value string,
	left int,
	right int,
	pad rune,
) string {
	return strings.Repeat(string(pad), left) +
		value +
		strings.Repeat(string(pad), right)
}
