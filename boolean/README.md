# boolean

<h6 align="left">
  Part of the
  <a href="https://github.com/entiqon/entiqon">entiqon</a>::<a href="https://github.com/entiqon/common">common</a>
  toolkit.
</h6>

The `boolean` package parses common boolean representations into `bool` values and formats booleans as custom strings.

The package requires **Go 1.18 or later**.

## Installation

```bash
go get github.com/entiqon/common/boolean
```

## Features

- Parse `bool`, `string`, integer, unsigned integer, and floating-point values into `bool`.
- Accept `true/false`, `1/0`, `yes/no`, `y/n`, `t/f`, and `on/off`, ignoring case and surrounding whitespace.
- Return an error for `nil`, unknown strings, and unsupported types instead of guessing.
- Fall back to a default value instead of handling an error with `FromOr`.
- Format a `bool` as any pair of strings.
- Zero external dependencies.

## Usage

### Parse a Value

Use `From` to read a boolean from configuration, query parameters, form values, or numbers.

```go
package main

import (
	"fmt"

	"github.com/entiqon/common/boolean"
)

func main() {
	for _, input := range []any{"yes", " OFF ", 1, 0.0, "maybe"} {
		value, err := boolean.From(input)
		if err != nil {
			fmt.Println("error:", err)
			continue
		}

		fmt.Println(value)
	}
}
```

Output:

```text
true
false
true
false
error: invalid boolean string: "maybe"
```

---

### Parse with a Fallback

Use `FromOr` when an invalid or missing value should simply use a default.

```go
package main

import (
	"fmt"
	"os"

	"github.com/entiqon/common/boolean"
)

func main() {
	os.Setenv("FEATURE_X", "on")

	fmt.Println(boolean.FromOr(os.Getenv("FEATURE_X"), false))
	fmt.Println(boolean.FromOr(os.Getenv("NOT_SET"), false))
	fmt.Println(boolean.FromOr("maybe", true))
}
```

Output:

```text
true
false
true
```

---

### Format a Boolean

Use `ToString` to turn a `bool` into labels for display.

```go
package main

import (
	"fmt"

	"github.com/entiqon/common/boolean"
)

func main() {
	fmt.Println(boolean.ToString(true, "enabled", "disabled"))
	fmt.Println(boolean.ToString(false, "yes", "no"))
}
```

Output:

```text
enabled
no
```

## API

| Function                                             | Description                                                              |
|------------------------------------------------------|--------------------------------------------------------------------------|
| `From(value any) (bool, error)`                      | Parses `value` into a `bool`; returns an error for unsupported input.    |
| `FromOr(value any, fallback bool) bool`              | Parses `value` into a `bool`; returns `fallback` for unsupported input.  |
| `ToString(b bool, trueStr, falseStr string) string`  | Returns `trueStr` if `b` is true; otherwise returns `falseStr`.          |

### Accepted Values

| Input                                   | `true`                                  | `false`                                   |
|-----------------------------------------|-----------------------------------------|-------------------------------------------|
| `bool`                                  | `true`                                  | `false`                                   |
| `string` (case-insensitive, trimmed)    | `true`, `1`, `yes`, `y`, `t`, `on`      | `false`, `0`, `no`, `n`, `f`, `off`       |
| integers, unsigned integers, floats     | any non-zero value                      | `0`                                       |

### Deprecated

This package moved from `common/extension/boolean`, and its `extension` shortcuts are replaced by `FromOr`. The old functions still work but will be removed in v2.0.0.

| Old function                                    | Replacement (`common/boolean`)   |
|-------------------------------------------------|----------------------------------|
| `extension/boolean.ParseFrom(value)`            | `From(value)`                    |
| `extension/boolean.BoolToString(b, t, f)`       | `ToString(b, t, f)`              |
| `extension.Boolean(value)`                      | `FromOr(value, false)`           |
| `extension.BooleanOr(value, fallback)`          | `FromOr(value, fallback)`        |

## License

This package is part of the **Entiqon Common** library and is distributed under the MIT License.
