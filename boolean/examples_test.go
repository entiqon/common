package boolean_test

import (
	"fmt"

	"github.com/entiqon/common/boolean"
)

func ExampleFrom() {
	b, err := boolean.From("yes")
	fmt.Println(b, err)

	_, err = boolean.From("maybe")
	fmt.Println(err)
	// Output:
	// true <nil>
	// invalid boolean string: "maybe"
}

func ExampleFromOr() {
	fmt.Println(boolean.FromOr("on", false))
	fmt.Println(boolean.FromOr("maybe", true))
	fmt.Println(boolean.FromOr(nil, false))
	// Output:
	// true
	// true
	// false
}

func ExampleToString() {
	fmt.Println(boolean.ToString(true, "yes", "no"))
	fmt.Println(boolean.ToString(false, "enabled", "disabled"))
	// Output:
	// yes
	// disabled
}
