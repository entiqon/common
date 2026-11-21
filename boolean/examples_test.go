package boolean_test

import (
	"fmt"

	"github.com/entiqon/common/boolean"
)

func ExampleFrom() {
	b, _ := boolean.From("yes")
	fmt.Println(b)
	// Output: true
}

func ExampleToString() {
	fmt.Println(boolean.ToString(true, "yes", "no"))
	fmt.Println(boolean.ToString(false, "enabled", "disabled"))
	// Output:
	// yes
	// disabled
}
