package text_test

import (
	"fmt"

	"github.com/entiqon/common/text"
)

func ExamplePadLeft() {
	fmt.Println(text.PadLeft("42", 5, '0'))

	// Output:
	// 00042
}

func ExamplePadRight() {
	fmt.Println(text.PadRight("42", 5, '.'))

	// Output:
	// 42...
}

func ExamplePadCenter() {
	fmt.Println(text.PadCenter("Go", 6, ' '))

	// Output:
	//   Go
}

func ExampleTruncate() {
	fmt.Println(text.Truncate("Hello World", 5))

	// Output:
	// Hello
}

func ExampleEllipsis() {
	fmt.Println(text.Ellipsis("Hello World", 5))

	// Output:
	// Hello...
}

func ExampleCapitalize() {
	fmt.Println(text.Capitalize("hello"))

	// Output:
	// Hello
}

func ExampleUncapitalize() {
	fmt.Println(text.Uncapitalize("Hello"))

	// Output:
	// hello
}

func ExampleReverse() {
	fmt.Println(text.Reverse("Hello"))

	// Output:
	// olleH
}
