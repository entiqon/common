package boolean

import (
	"testing"
)

func TestParseFrom(t *testing.T) {
	got, err := ParseFrom("yes")
	if err != nil || !got {
		t.Errorf(`ParseFrom("yes") = %v, %v; want true, nil`, got, err)
	}

	if _, err := ParseFrom("maybe"); err == nil {
		t.Error(`ParseFrom("maybe") expected an error`)
	}
}

func TestBoolToString(t *testing.T) {
	if got := BoolToString(false, "on", "off"); got != "off" {
		t.Errorf(`BoolToString(false, "on", "off") = %q; want "off"`, got)
	}
}
