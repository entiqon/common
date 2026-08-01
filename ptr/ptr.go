package ptr

import "strings"

// EmptyString returns a pointer to an empty string.
func EmptyString() *string {
	return new(string)
}

// ToPtr returns a pointer to v.
func ToPtr[T any](v T) *T {
	return &v
}

// FromPtr dereferences p if non-nil, otherwise returns def.
func FromPtr[T any](p *T, def T) T {
	if p != nil {
		return *p
	}
	return def
}

// NilIfEmpty returns nil if v is empty; otherwise it returns a pointer to v.
func NilIfEmpty(v string) *string {
	if v == "" {
		return nil
	}
	return &v
}

// NilIfBlank returns nil if v contains only whitespace;
// otherwise it returns a pointer to v.
func NilIfBlank(v string) *string {
	if strings.TrimSpace(v) == "" {
		return nil
	}
	return &v
}
