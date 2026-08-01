package ptr

import "strings"

// EmptyString returns a pointer to "".
func EmptyString() *string {
	return new(string)
}

// ToPtr returns a pointer to value.
//
// Deprecated: use Of instead.
func ToPtr[T any](value T) *T {
	return Of(value)
}

// Of returns a pointer to value.
func Of[T any](value T) *T {
	return &value
}

// FromPtr dereferences p if non-nil, otherwise returns def.
//
// Deprecated: use Deref instead.
func FromPtr[T any](p *T, def T) T {
	return Deref(p, def)
}

// Deref dereferences p if non-nil, otherwise returns def.
func Deref[T any](p *T, def T) T {
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

// NilIfBlank returns nil if v is blank; otherwise it returns a pointer to v.
func NilIfBlank(v string) *string {
	if strings.TrimSpace(v) == "" {
		return nil
	}
	return &v
}
