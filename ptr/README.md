# ptr

<h6 align="left">
  Part of the
  <a href="https://github.com/entiqon/entiqon">entiqon</a>::<a href="https://github.com/entiqon/common">common</a>
  toolkit.
</h6>

The `ptr` package provides lightweight, generic helpers for working with pointers in Go. It simplifies common pointer operations such as creating pointers, safely dereferencing values, working with optional values, and normalizing empty or blank strings.

The package requires **Go 1.18 or later** because it uses generics.

## Installation

```bash
go get github.com/entiqon/common/ptr
```

## Features

- Create pointers from values.
- Safely dereference pointers with a default value.
- Create pointers to empty strings.
- Convert empty strings to `nil`.
- Convert blank (whitespace-only) strings to `nil`.
- Generic implementation with zero external dependencies.

## Usage

### Create a Pointer

Use `Of` to create a pointer from any value.

```go
package main

import (
	"fmt"

	"github.com/entiqon/common/ptr"
)

func main() {
	name := ptr.Of("John")

	fmt.Println(*name)
}
```

Output:

```text
John
```

---

### Dereference with a Default Value

Use `Deref` to safely read a pointer, returning a default value when the pointer is `nil`.

```go
package main

import (
	"fmt"

	"github.com/entiqon/common/ptr"
)

func main() {
	var age *int

	fmt.Println(ptr.Deref(age, 18))

	age = ptr.Of(25)

	fmt.Println(ptr.Deref(age, 18))
}
```

Output:

```text
18
25
```

---

### Create an Empty String Pointer

Use `EmptyString` when a `*string` is required and the value should be the empty string.

```go
package main

import (
	"fmt"

	"github.com/entiqon/common/ptr"
)

func main() {
	value := ptr.EmptyString()

	fmt.Printf("%q\n", *value)
}
```

Output:

```text
""
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

	value := ptr.NilIfEmpty("hello")

	if value != nil {
		fmt.Println(*value)
	}
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

	if value != nil {
		fmt.Println(*value)
	}
}
```

Output:

```text
<nil>
 hello 
```

## API

| Function                       | Description                                                                                    |
|--------------------------------|------------------------------------------------------------------------------------------------|
| `Of[T any](v T) *T`            | Returns a pointer to `v`.                                                                      |
| `Deref[T any](p *T, def T) T`  | Returns the dereferenced value if `p` is not `nil`; otherwise returns `def`.                   |
| `EmptyString() *string`        | Returns a pointer to an empty string (`""`).                                                   |
| `NilIfEmpty(v string) *string` | Returns `nil` if `v` is empty; otherwise returns a pointer to `v`.                             |
| `NilIfBlank(v string) *string` | Returns `nil` if `v` is empty or contains only whitespace; otherwise returns a pointer to `v`. |

### Deprecated

| Function                        | Replacement                  |
|---------------------------------|------------------------------|
| `ToPtr[T any](v T) *T`          | Use `Of(v)` instead.         |
| `FromPtr[T any](p *T, def T) T` | Use `Deref(p, def)` instead. |

## When to Use

The `ptr` package is useful when working with:

- Optional struct fields.
- JSON, XML, and YAML serialization.
- Database models with nullable or optional values.
- Configuration objects.
- DTOs and API request/response models.
- APIs that use pointers to distinguish between zero values and missing values.

## License

This package is part of the **Entiqon Common** library and is distributed under the MIT License.