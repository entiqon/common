package ptr_test

import (
	"fmt"

	"github.com/entiqon/common/ptr"
)

func ExampleOf() {
	name := ptr.Of("John Doe")

	fmt.Println(*name)

	// Output:
	// John Doe
}

func ExampleDeref() {
	var name *string

	fmt.Println(ptr.Deref(name, "Unknown"))

	name = ptr.Of("John Doe")

	fmt.Println(ptr.Deref(name, "Unknown"))

	// Output:
	// Unknown
	// John Doe
}

func ExampleEmptyString() {
	value := ptr.EmptyString()

	fmt.Printf("%q\n", *value)

	// Output:
	// ""
}

func ExampleNilIfEmpty() {
	fmt.Println(ptr.NilIfEmpty(""))

	name := ptr.NilIfEmpty("John Doe")
	if name != nil {
		fmt.Println(*name)
	}

	// Output:
	// <nil>
	// John Doe
}

func ExampleNilIfBlank() {
	fmt.Println(ptr.NilIfBlank("   "))

	name := ptr.NilIfBlank("John Doe")
	if name != nil {
		fmt.Println(*name)
	}

	// Output:
	// <nil>
	// John Doe
}
