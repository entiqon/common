# text

<h6 align="left">
  Package: <a href="https://github.com/entiqon/entiqon">entiqon</a>/<a href="https://github.com/entiqon/common">common</a>/text
</h6>

The `text` package provides common string utilities that complement Go's
standard library.

It includes helpers for padding, truncation, and presentation while maintaining
a small, focused, and idiomatic API.

## Installation

```bash
go get github.com/entiqon/common/text
```

## Features

- Left, right, and center padding
- Unicode-aware truncation
- Ellipsis formatting

## Functions

| Function    | Description                                           |
|-------------|-------------------------------------------------------|
| `PadLeft`   | Left-pads a string to a fixed length.                 |
| `PadRight`  | Right-pads a string to a fixed length.                |
| `PadCenter` | Centers a string within a fixed length.               |
| `Truncate`  | Truncates a string to a maximum number of characters. |
| `Ellipsis`  | Truncates a string and appends `"..."` when needed.   |

## Usage

### PadLeft

```go
text.PadLeft("42", 5, '0')
// 00042
```

### PadRight

```go
text.PadRight("42", 5, '.')
// 42...
```

### PadCenter

```go
text.PadCenter("Go", 6, ' ')
// "  Go  "
```

### Truncate

```go
text.Truncate("Hello World", 5)
// Hello
```

### Ellipsis

```go
text.Ellipsis("Hello World", 5)
// Hello...
```

## Notes

Padding functions operate on byte length, making them suitable for ASCII
identifiers, codes, and fixed-width text.

Truncation functions operate on Unicode characters (runes) to avoid splitting
multi-byte characters.

The `Ellipsis` function truncates the string to the specified number of
characters before appending `"..."` when truncation occurs. The ellipsis is
not included in the specified maximum length.

## License

This package is part of the **Entiqon Common** library and is distributed under
the MIT License.