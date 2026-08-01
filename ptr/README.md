# ptr

<h6 align="left">Part of the <a href="https://github.com/entiqon/entiqon">entiqon</a>::<a href="https://github.com/entiqon/entiqon">common</a> toolkit.</h6>

The `ptr` package provides small, generic helpers for working with pointers in Go. It simplifies common pointer operations such as creating pointers, safely dereferencing values, and converting empty or blank strings to `nil`.

## Installation

```bash
go get github.com/entiqon/common/ptr
```

## Features

- Create pointers from values.
- Safely dereference pointers with a default value.
- Convert empty strings to `nil`.
- Convert blank (whitespace-only) strings to `nil`.
- Generic implementation with zero dependencies.

## Usage

### Create a Pointer

Use `ToPtr` to create a pointer from any value.

```go
package main

import (
	"fmt"

	"github.com/entiqon/common/ptr"
)

func main() {
	name := ptr.ToPtr("John")

	fmt.Println(*name)
}
```

Output:

```text
John
```

---

### Dereference with a Default Value

Use `FromPtr` to safely read a pointer, returning a default value when the pointer is `nil`.

```go
package main

import (
	"fmt"

	"github.com/entiqon/common/ptr"
)

func main() {
	var age *int

	fmt.Println(ptr.FromPtr(age, 18))

	age = ptr.ToPtr(25)

	fmt.Println(ptr.FromPtr(age, 18))
}
```

Output:

```text
18
25
```

---

### Return `nil` for Empty Strings

Use `NilIfEmpty` when an empty string should be treated as an absent value.

```go
package main

import (
	"fmt"

	"github.com/entiqon/common/ptr"
)

func main() {
	fmt.Println(ptr.NilIfEmpty(""))
	fmt.Println(*ptr.NilIfEmpty("hello"))
}
```

Output:

```text
<nil>
hello
```

---

### Return `nil` for Blank Strings

Use `NilIfBlank` when whitespace-only strings should also be considered absent.

```go
package main

import (
	"fmt"

	"github.com/entiqon/common/ptr"
)

func main() {
	fmt.Println(ptr.NilIfBlank("   "))

	value := ptr.NilIfBlank(" hello ")

	fmt.Println(*value)
}
```

Output:

```text
<nil>
 hello 
```

## API

| Function                        | Description                                                                                    |
|---------------------------------|------------------------------------------------------------------------------------------------|
| `ToPtr[T any](v T) *T`          | Returns a pointer to `v`.                                                                      |
| `FromPtr[T any](p *T, def T) T` | Returns the dereferenced value if `p` is not `nil`; otherwise returns `def`.                   |
| `NilIfEmpty(v string) *string`  | Returns `nil` if `v` is empty; otherwise returns a pointer to `v`.                             |
| `NilIfBlank(v string) *string`  | Returns `nil` if `v` is empty or contains only whitespace; otherwise returns a pointer to `v`. |

## When to Use

The `ptr` package is useful when working with:

- Optional fields in structs.
- JSON and XML serialization.
- Database models with nullable values.
- Configuration values.
- APIs that use pointers to distinguish between zero values and missing values.

## License

This package is part of the Entiqon Common library.