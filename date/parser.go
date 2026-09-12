package date

import (
	"errors"
	"fmt"
	"math"
	"strconv"
	"strings"
	"time"

	"github.com/entiqon/common/date/internal"
)

// Order identifies the component order of numeric dates.
type Order uint8

const (
	YearFirst Order = iota
	DayFirst
	MonthFirst
)

type parseLayout struct {
	value    string
	dateOnly bool
	zoned    bool
}

// From converts value into a time.Time.
//
// YearFirst is used by default. Use DayFirst or MonthFirst when the year is not
// the first component.
func From(value any, orders ...Order) (time.Time, error) {
	order, err := resolveOrder(orders)
	if err != nil {
		return time.Time{}, err
	}

	if value == nil {
		return time.Time{}, errors.New("date.From: value is nil")
	}

	switch typed := value.(type) {
	case time.Time:
		return typed, nil

	case *time.Time:
		if typed == nil {
			return time.Time{}, errors.New(
				"date.From: *time.Time is nil",
			)
		}

		return *typed, nil

	case string:
		return parseString(typed, order)

	case *string:
		if typed == nil {
			return time.Time{}, errors.New(
				"date.From: *string is nil",
			)
		}

		return parseString(*typed, order)

	case []byte:
		return parseString(string(typed), order)

	case int:
		return fromSigned(int64(typed))

	case int8:
		return fromSigned(int64(typed))

	case int16:
		return fromSigned(int64(typed))

	case int32:
		return fromSigned(int64(typed))

	case int64:
		return fromSigned(typed)

	case uint:
		return fromUnsigned(uint64(typed))

	case uint8:
		return fromUnsigned(uint64(typed))

	case uint16:
		return fromUnsigned(uint64(typed))

	case uint32:
		return fromUnsigned(uint64(typed))

	case uint64:
		return fromUnsigned(typed)

	case uintptr:
		return fromUnsigned(uint64(typed))

	case float32:
		return time.Unix(int64(typed), 0).UTC(), nil

	case float64:
		return fromFloat64(typed), nil

	default:
		return time.Time{}, fmt.Errorf(
			"date.From: unsupported type %T",
			value,
		)
	}
}

// ParseAndFormat parses value and formats it with layout.
//
// An empty layout defaults to "2006-01-02". An empty string is returned when
// parsing fails.
func ParseAndFormat(
	value string,
	layout string,
	orders ...Order,
) string {
	if layout == "" {
		layout = "2006-01-02"
	}

	parsed, err := From(value, orders...)
	if err != nil {
		return ""
	}

	return parsed.Format(layout)
}

func resolveOrder(orders []Order) (Order, error) {
	switch len(orders) {
	case 0:
		return YearFirst, nil
	case 1:
	default:
		return YearFirst, errors.New(
			"date.From: only one date order may be specified",
		)
	}

	switch orders[0] {
	case YearFirst, DayFirst, MonthFirst:
		return orders[0], nil
	default:
		return YearFirst, fmt.Errorf(
			"date.From: unsupported date order %d",
			orders[0],
		)
	}
}

func parseString(
	value string,
	order Order,
) (time.Time, error) {
	value = strings.TrimSpace(value)
	if value == "" {
		return time.Time{}, errors.New(
			"date.From: value is empty",
		)
	}

	if parsed, handled, err := parseDigits(value, order); handled {
		return parsed, err
	}

	for _, candidate := range layoutsFor(order) {
		parsed, err := parseWithLayout(value, candidate)
		if err == nil {
			return parsed, nil
		}
	}

	return time.Time{}, fmt.Errorf(
		"date.From: unsupported date %q",
		value,
	)
}

func parseWithLayout(
	value string,
	candidate parseLayout,
) (time.Time, error) {
	if candidate.zoned {
		return time.Parse(candidate.value, value)
	}

	location := time.Local
	if candidate.dateOnly {
		location = time.UTC
	}

	parsed, err := time.ParseInLocation(
		candidate.value,
		value,
		location,
	)
	if err != nil {
		return time.Time{}, err
	}

	if candidate.dateOnly {
		return parsed, nil
	}

	return parsed.UTC(), nil
}

func parseDigits(
	value string,
	order Order,
) (time.Time, bool, error) {
	if !isDigits(value) {
		return time.Time{}, false, nil
	}

	switch len(value) {
	case 8:
		parsed, err := time.ParseInLocation(
			compactDateLayout(order),
			value,
			time.UTC,
		)
		if err != nil {
			return time.Time{}, true, fmt.Errorf(
				"date.From: invalid compact date %q: %w",
				value,
				err,
			)
		}

		return parsed, true, nil

	case 10:
		seconds, _ := strconv.ParseInt(value, 10, 64)

		return time.Unix(seconds, 0).UTC(), true, nil

	case 13:
		milliseconds, _ := strconv.ParseInt(value, 10, 64)

		return time.UnixMilli(milliseconds).UTC(), true, nil

	default:
		return time.Time{}, false, nil
	}
}

func isDigits(value string) bool {
	for _, char := range value {
		if char < '0' || char > '9' {
			return false
		}
	}

	return true
}

func layoutsFor(order Order) []parseLayout {
	layouts := []parseLayout{
		{
			value: time.RFC3339Nano,
			zoned: true,
		},
		{
			value: time.RFC1123,
			zoned: true,
		},
		{
			value: time.RFC1123Z,
			zoned: true,
		},
	}

	switch order {
	case DayFirst:
		return append(layouts, dayFirstLayouts()...)

	case MonthFirst:
		return append(layouts, monthFirstLayouts()...)

	default:
		return append(layouts, yearFirstLayouts()...)
	}
}

func yearFirstLayouts() []parseLayout {
	return []parseLayout{
		{
			value:    "2006-01-02",
			dateOnly: true,
		},
		{
			value:    "2006/01/02",
			dateOnly: true,
		},
		{
			value: "2006-01-02 15:04:05",
		},
		{
			value: "2006/01/02 15:04:05",
		},
		{
			value: "2006-01-02T15:04:05",
		},
		{
			value: "2006/01/02T15:04:05",
		},
		{
			value: "20060102T150405",
		},
		{
			value: "20060102T150405Z0700",
			zoned: true,
		},
		{
			value:    "02 Jan 2006",
			dateOnly: true,
		},
	}
}

func dayFirstLayouts() []parseLayout {
	return []parseLayout{
		{
			value:    "02-01-2006",
			dateOnly: true,
		},
		{
			value:    "02/01/2006",
			dateOnly: true,
		},
		{
			value: "02-01-2006 15:04:05",
		},
		{
			value: "02/01/2006 15:04:05",
		},
		{
			value: "02-01-2006T15:04:05",
		},
		{
			value: "02/01/2006T15:04:05",
		},
		{
			value: "02012006T150405",
		},
		{
			value: "02-01-2006T15:04:05Z07:00",
			zoned: true,
		},
		{
			value: "02/01/2006T15:04:05Z07:00",
			zoned: true,
		},
		{
			value: "02012006T150405Z0700",
			zoned: true,
		},
		{
			value:    "02 Jan 2006",
			dateOnly: true,
		},
	}
}

func monthFirstLayouts() []parseLayout {
	return []parseLayout{
		{
			value:    "01-02-2006",
			dateOnly: true,
		},
		{
			value:    "01/02/2006",
			dateOnly: true,
		},
		{
			value: "01-02-2006 15:04:05",
		},
		{
			value: "01/02/2006 15:04:05",
		},
		{
			value: "01-02-2006T15:04:05",
		},
		{
			value: "01/02/2006T15:04:05",
		},
		{
			value: "01022006T150405",
		},
		{
			value: "01-02-2006T15:04:05Z07:00",
			zoned: true,
		},
		{
			value: "01/02/2006T15:04:05Z07:00",
			zoned: true,
		},
		{
			value: "01022006T150405Z0700",
			zoned: true,
		},
	}
}

func compactDateLayout(order Order) string {
	switch order {
	case DayFirst:
		return "02012006"

	case MonthFirst:
		return "01022006"

	default:
		return "20060102"
	}
}

func fromSigned(value int64) (time.Time, error) {
	seconds, err := internal.ToSecondsSigned(value)
	if err != nil {
		return time.Time{}, err
	}

	return time.Unix(seconds, 0).UTC(), nil
}

func fromUnsigned(value uint64) (time.Time, error) {
	seconds, err := internal.ToSecondsUnsigned(value)
	if err != nil {
		return time.Time{}, err
	}

	return time.Unix(seconds, 0).UTC(), nil
}

func fromFloat64(value float64) time.Time {
	totalMilliseconds := math.Round(value * 1e3)
	totalNanoseconds := int64(
		totalMilliseconds,
	) * 1_000_000

	seconds := totalNanoseconds / 1_000_000_000
	nanoseconds := totalNanoseconds % 1_000_000_000

	if nanoseconds < 0 {
		seconds--
		nanoseconds += 1_000_000_000
	}

	return time.Unix(seconds, nanoseconds).UTC()
}
