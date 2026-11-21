package boolean_test

import (
	"testing"

	"github.com/entiqon/common/boolean"
)

func TestToString(t *testing.T) {
	tests := []struct {
		input    bool
		trueStr  string
		falseStr string
		want     string
	}{
		{true, "yes", "no", "yes"},
		{false, "yes", "no", "no"},
		{true, "enabled", "disabled", "enabled"},
		{false, "enabled", "disabled", "disabled"},
	}

	for _, tt := range tests {
		got := boolean.ToString(tt.input, tt.trueStr, tt.falseStr)
		if got != tt.want {
			t.Errorf("ToString(%v, %q, %q) = %q; want %q",
				tt.input, tt.trueStr, tt.falseStr, got, tt.want)
		}
	}
}
