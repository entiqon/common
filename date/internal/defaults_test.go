package internal_test

import (
	"testing"

	"github.com/entiqon/common/date/internal"
	"github.com/entiqon/common/test"
)

func TestDefaults(t *testing.T) {
	t.Run("DefaultLayouts", func(t *testing.T) {
		got := internal.DefaultLayouts()

		test.AssertEqual(t, 6, len(got))
		test.AssertEqual(t, "2006-01-02", got[0])
		test.AssertEqual(t, "2006/01/02", got[1])
		test.AssertEqual(t, "2006-01-02 15:04:05", got[2])
		test.AssertEqual(t, "2006/01/02 15:04:05", got[3])
		test.AssertEqual(t, "02 Jan 2006", got[4])
		test.AssertEqual(t, "Mon, 02 Jan 2006 15:04:05 MST", got[5])
	})
}
