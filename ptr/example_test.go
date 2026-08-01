package ptr_test

import (
	"fmt"

	"github.com/entiqon/common/ptr"
)

func ExampleToPtr() {
	name := ptr.ToPtr("John Doe")

	fmt.Println(*name)

	// Output:
	// John Doe
}

func ExampleFromPtr() {
	var name *string

	fmt.Println(ptr.FromPtr(name, "Unknown"))

	name = ptr.ToPtr("John Doe")

	fmt.Println(ptr.FromPtr(name, "Unknown"))

	// Output:
	// Unknown
	// John Doe
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
