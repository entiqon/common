# Date Parser 📅

Parsing and formatting utilities for Go `time.Time` values.

Part of the [`entiqon`](https://github.com/entiqon/entiqon) `common` module.

---

## ✨ Features

* Parse strings, byte slices, time values, pointers, and numeric timestamps
* Support year-first, day-first, and month-first dates
* Support extended and compact date formats
* Support date-only, date-time, RFC 3339, and RFC 1123 inputs
* Preserve embedded timezones
* Normalize date-only inputs to midnight UTC
* Interpret timezone-less date-times in `time.Local` and convert them to UTC
* Format parsed values with `ParseAndFormat`
* Return descriptive errors for invalid or unsupported inputs

---

## 🕒 Supported Formats

### Year First

Year-first is the default order.

| Input                       | Description                                   |
| --------------------------- | --------------------------------------------- |
| `2026-08-31`                | Date                                          |
| `2026/08/31`                | Date with slash separators                    |
| `20260831`                  | Compact date                                  |
| `2026-08-31 15:00:00`       | Date and time                                 |
| `2026/08/31 15:00:00`       | Date and time with slash separators           |
| `2026-08-31T15:00:00`       | ISO date and time without a timezone          |
| `20260831T150000`           | Compact date and time                         |
| `2026-08-31T15:00:00Z`      | RFC 3339 date and time in UTC                 |
| `2026-08-31T15:00:00-04:00` | RFC 3339 date and time with a timezone offset |
| `2026-08-31T15:00:00+04:00` | RFC 3339 date and time with a timezone offset |
| `20260831T150000Z`          | Compact date and time in UTC                  |
| `20260831T150000-0400`      | Compact date and time with a timezone offset  |
| `20260831T150000+0400`      | Compact date and time with a timezone offset  |

### Day First

Pass `date.DayFirst` when the day is the first component.

| Input                       | Description                                  |
| --------------------------- | -------------------------------------------- |
| `31-08-2026`                | Date                                         |
| `31/08/2026`                | Date with slash separators                   |
| `31082026`                  | Compact date                                 |
| `31-08-2026 15:00:00`       | Date and time                                |
| `31/08/2026 15:00:00`       | Date and time with slash separators          |
| `31-08-2026T15:00:00`       | Date and time with `T` separator             |
| `31082026T150000`           | Compact date and time                        |
| `31-08-2026T15:00:00-04:00` | Date and time with a timezone offset         |
| `31/08/2026T15:00:00-04:00` | Zoned date and time with slash separators    |
| `31082026T150000-0400`      | Compact date and time with a timezone offset |

### Month First

Pass `date.MonthFirst` when the month is the first component.

| Input                       | Description                                  |
| --------------------------- | -------------------------------------------- |
| `08-31-2026`                | Date                                         |
| `08/31/2026`                | Date with slash separators                   |
| `08312026`                  | Compact date                                 |
| `08-31-2026 15:00:00`       | Date and time                                |
| `08/31/2026 15:00:00`       | Date and time with slash separators          |
| `08-31-2026T15:00:00`       | Date and time with `T` separator             |
| `08312026T150000`           | Compact date and time                        |
| `08-31-2026T15:00:00-04:00` | Date and time with a timezone offset         |
| `08/31/2026T15:00:00-04:00` | Zoned date and time with slash separators    |
| `08312026T150000-0400`      | Compact date and time with a timezone offset |

### Additional Formats

| Input                             | Description         |
| --------------------------------- | ------------------- |
| `31 Aug 2026`                     | Named-month date    |
| `Mon, 31 Aug 2026 15:00:00 GMT`   | RFC 1123 date-time  |
| `Mon, 31 Aug 2026 15:00:00 +0000` | RFC 1123Z date-time |

---

## 🔢 Numeric Timestamps

| Input            | Interpretation                                  |
| ---------------- | ----------------------------------------------- |
| 10-digit string  | Unix seconds                                    |
| 13-digit string  | Unix milliseconds                               |
| Signed integer   | Unix seconds                                    |
| Unsigned integer | Unix seconds                                    |
| `float32`        | Unix seconds with the fractional part discarded |
| `float64`        | Unix seconds rounded to millisecond precision   |

Numeric timestamps are returned in UTC.

---

## 🌐 Timezone Handling

| Input                        | Behavior                                           |
| ---------------------------- | -------------------------------------------------- |
| Date only                    | Normalized to `00:00:00 UTC`                       |
| Date-time without a timezone | Interpreted in `time.Local`, then converted to UTC |
| Date-time with a timezone    | Embedded timezone is preserved                     |
| Numeric timestamp            | Returned in UTC                                    |
| Existing `time.Time`         | Returned unchanged                                 |

For example, when `time.Local` is UTC-05:00:

```text
2026-08-31 15:00:00 → 2026-08-31T20:00:00Z
```

---

## ⚠️ Ambiguous Dates

When both numeric components are 12 or lower, the value may represent either
day-first or month-first order:

```text
03-04-2026
```

Specify the intended order explicitly:

```go
dayFirst, err := date.From(
    "03-04-2026",
    date.DayFirst,
)
// 2026-04-03T00:00:00Z

monthFirst, err := date.From(
    "03-04-2026",
    date.MonthFirst,
)
// 2026-03-04T00:00:00Z
```

---

## 📑 API Reference

### `From(value any, orders ...Order) (time.Time, error)`

Converts a supported value into `time.Time`.

Supported values:

| Type             | Behavior                                                 |
| ---------------- | -------------------------------------------------------- |
| `string`         | Parsed using the selected date order                     |
| `*string`        | Dereferenced and parsed                                  |
| `[]byte`         | Converted to a string and parsed                         |
| `time.Time`      | Returned unchanged                                       |
| `*time.Time`     | Dereferenced and returned                                |
| Signed integer   | Interpreted as Unix seconds                              |
| Unsigned integer | Interpreted as Unix seconds                              |
| `float32`        | Interpreted as Unix seconds without fractional precision |
| `float64`        | Interpreted as Unix seconds with millisecond precision   |

`YearFirst` is used when no order is provided. Only one order may be supplied.

### `ParseAndFormat(value, layout string, orders ...Order) string`

Parses `value` and formats it using the supplied Go layout.

An empty layout defaults to `2006-01-02`. The function returns an empty string
when parsing fails.

### `Cleaning(raw string) string`

Normalizes a raw date string before parsing.

---

## 🔹 Usage

```go
package main

import (
    "fmt"
    "time"

    "github.com/entiqon/common/date"
)

func main() {
    yearFirst, err := date.From(
        "2026-08-31T15:00:00-04:00",
    )
    if err != nil {
        panic(err)
    }

    dayFirst, err := date.From(
        "31/08/2026 15:00:00",
        date.DayFirst,
    )
    if err != nil {
        panic(err)
    }

    monthFirst, err := date.From(
        "08/31/2026 15:00:00",
        date.MonthFirst,
    )
    if err != nil {
        panic(err)
    }

    compact, err := date.From("20260831T150000")
    if err != nil {
        panic(err)
    }

    formatted := date.ParseAndFormat(
        "31/08/2026",
        "2006-01-02",
        date.DayFirst,
    )

    fmt.Println(yearFirst.Format(time.RFC3339))
    fmt.Println(dayFirst.Format(time.RFC3339))
    fmt.Println(monthFirst.Format(time.RFC3339))
    fmt.Println(compact.Format(time.RFC3339))
    fmt.Println(formatted)
}
```

---

## ❌ Errors

`From` returns an error when:

* The input is `nil`
* A supported pointer is nil
* A string is empty or unrecognized
* A date is invalid or incomplete
* The input type is unsupported
* More than one date order is supplied
* The date order is invalid

---

## 📌 Summary

* **Parsing:** `From(value any, orders ...Order) (time.Time, error)`
* **Formatting:** `ParseAndFormat(value, layout string, orders ...Order) string`
* **Orders:** `YearFirst`, `DayFirst`, and `MonthFirst`
* **Cleaning:** `Cleaning(raw string) string`
