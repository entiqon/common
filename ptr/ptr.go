package ptr

import "strings"

// EmptyString returns a pointer to "".
func EmptyString() *string {
	return Of("")
}

// Of returns a pointer to value.
func Of[T any](value T) *T {
	v := &value

	return v
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

	return Of(v)
}

// NilIfBlank returns nil if v is blank; otherwise it returns a pointer to v.
func NilIfBlank(v string) *string {
	if strings.TrimSpace(v) == "" {
		return nil
	}

	return Of(v)
}

// ToPtr returns a pointer to value.
//
// Deprecated: Use Of instead. ToPtr will be removed in v2.0.0.
func ToPtr[T any](value T) *T {
	return Of(value)
}

// FromPtr dereferences p if non-nil, otherwise returns def.
//
// Deprecated: Use Deref instead. FromPtr will be removed in v2.0.0.
func FromPtr[T any](p *T, def T) T {
	return Deref(p, def)
}
