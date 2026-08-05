package text

import "unicode"

// Capitalize returns value with its first character converted to upper case.
//
// If value is empty, it is returned unchanged.
func Capitalize(value string) string {
	runes := []rune(value)
	if len(runes) == 0 {
		return value
	}

	runes[0] = unicode.ToUpper(runes[0])

	return string(runes)
}

// Uncapitalize returns value with its first character converted to lower case.
//
// If value is empty, it is returned unchanged.
func Uncapitalize(value string) string {
	runes := []rune(value)
	if len(runes) == 0 {
		return value
	}

	runes[0] = unicode.ToLower(runes[0])

	return string(runes)
}
