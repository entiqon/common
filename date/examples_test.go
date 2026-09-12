package date_test

import (
	"fmt"
	"math"
	"time"

	"github.com/entiqon/common/date"
)

func ExampleFrom_strings() {
	originalLocal := time.Local
	time.Local = time.FixedZone("TEST-0500", -5*60*60)
	defer func() { time.Local = originalLocal }()

	tests := []struct {
		input string
		order []date.Order
	}{
		{"2026-08-31", nil},
		{"2026/08/31", nil},
		{"20260831", nil},
		{"2026-08-31 15:00:00", nil},
		{"2026-08-31T15:00:00", nil},
		{"2026/08/31 15:00:00", nil},
		{"20260831T150000", nil},
		{"31-08-2026", []date.Order{date.DayFirst}},
		{"31/08/2026", []date.Order{date.DayFirst}},
		{"31082026", []date.Order{date.DayFirst}},
		{"31-08-2026 15:00:00", []date.Order{date.DayFirst}},
		{"31/08/2026 15:00:00", []date.Order{date.DayFirst}},
		{"31-08-2026T15:00:00", []date.Order{date.DayFirst}},
		{"31082026T150000", []date.Order{date.DayFirst}},
		{"08-31-2026", []date.Order{date.MonthFirst}},
		{"08/31/2026", []date.Order{date.MonthFirst}},
		{"08312026", []date.Order{date.MonthFirst}},
		{"08-31-2026 15:00:00", []date.Order{date.MonthFirst}},
		{"08/31/2026 15:00:00", []date.Order{date.MonthFirst}},
		{"08-31-2026T15:00:00", []date.Order{date.MonthFirst}},
		{"08312026T150000", []date.Order{date.MonthFirst}},
		{"2026-08-31T15:00:00Z", nil},
		{"2026-08-31T15:00:00-04:00", nil},
		{"2026-08-31T15:00:00+04:00", nil},
		{"20260831T150000Z", nil},
		{"20260831T150000-0400", nil},
		{"20260831T150000+0400", nil},
		{"Mon, 31 Aug 2026 15:00:00 GMT", nil},
	}

	for _, test := range tests {
		value, err := date.From(test.input, test.order...)
		if err != nil {
			fmt.Printf("%s: %v\n", test.input, err)
			continue
		}
		fmt.Printf("%-31s -> %s\n", test.input, value.Format(time.RFC3339))
	}

	// Output:
	// 2026-08-31                      -> 2026-08-31T00:00:00Z
	// 2026/08/31                      -> 2026-08-31T00:00:00Z
	// 20260831                        -> 2026-08-31T00:00:00Z
	// 2026-08-31 15:00:00             -> 2026-08-31T20:00:00Z
	// 2026-08-31T15:00:00             -> 2026-08-31T20:00:00Z
	// 2026/08/31 15:00:00             -> 2026-08-31T20:00:00Z
	// 20260831T150000                 -> 2026-08-31T20:00:00Z
	// 31-08-2026                      -> 2026-08-31T00:00:00Z
	// 31/08/2026                      -> 2026-08-31T00:00:00Z
	// 31082026                        -> 2026-08-31T00:00:00Z
	// 31-08-2026 15:00:00             -> 2026-08-31T20:00:00Z
	// 31/08/2026 15:00:00             -> 2026-08-31T20:00:00Z
	// 31-08-2026T15:00:00             -> 2026-08-31T20:00:00Z
	// 31082026T150000                 -> 2026-08-31T20:00:00Z
	// 08-31-2026                      -> 2026-08-31T00:00:00Z
	// 08/31/2026                      -> 2026-08-31T00:00:00Z
	// 08312026                        -> 2026-08-31T00:00:00Z
	// 08-31-2026 15:00:00             -> 2026-08-31T20:00:00Z
	// 08/31/2026 15:00:00             -> 2026-08-31T20:00:00Z
	// 08-31-2026T15:00:00             -> 2026-08-31T20:00:00Z
	// 08312026T150000                 -> 2026-08-31T20:00:00Z
	// 2026-08-31T15:00:00Z            -> 2026-08-31T15:00:00Z
	// 2026-08-31T15:00:00-04:00       -> 2026-08-31T15:00:00-04:00
	// 2026-08-31T15:00:00+04:00       -> 2026-08-31T15:00:00+04:00
	// 20260831T150000Z                -> 2026-08-31T15:00:00Z
	// 20260831T150000-0400            -> 2026-08-31T15:00:00-04:00
	// 20260831T150000+0400            -> 2026-08-31T15:00:00+04:00
	// Mon, 31 Aug 2026 15:00:00 GMT   -> 2026-08-31T15:00:00Z
}

func ExampleFrom_values() {
	timeValue := time.Date(2026, time.August, 31, 15, 0, 0, 0, time.UTC)
	stringValue := "2026-08-31"
	values := []any{
		timeValue, &timeValue, stringValue, &stringValue, []byte(stringValue),
		int(42), int8(-7), int16(42), int32(42), int64(1_700_000_000),
		uint(42), uint8(7), uint16(42), uint32(42),
		uint64(1_700_000_000), uintptr(42),
	}

	for _, input := range values {
		value, err := date.From(input)
		fmt.Printf("%-12T unix=%d error=%t\n", input, value.Unix(), err != nil)
	}

	seconds, _ := date.From("1700000000")
	milliseconds, _ := date.From("1700000000123")
	float32Value, _ := date.From(float32(-1.987))
	float64Value, _ := date.From(1_700_000_000.123)
	fmt.Println("epoch seconds:", seconds.Format(time.RFC3339Nano))
	fmt.Println("epoch milliseconds:", milliseconds.Format(time.RFC3339Nano))
	fmt.Println("float32:", float32Value.Format(time.RFC3339Nano))
	fmt.Println("float64:", float64Value.Format(time.RFC3339Nano))

	// Output:
	// time.Time    unix=1788188400 error=false
	// *time.Time   unix=1788188400 error=false
	// string       unix=1788134400 error=false
	// *string      unix=1788134400 error=false
	// []uint8      unix=1788134400 error=false
	// int          unix=42 error=false
	// int8         unix=-7 error=false
	// int16        unix=42 error=false
	// int32        unix=42 error=false
	// int64        unix=1700000000 error=false
	// uint         unix=42 error=false
	// uint8        unix=7 error=false
	// uint16       unix=42 error=false
	// uint32       unix=42 error=false
	// uint64       unix=1700000000 error=false
	// uintptr      unix=42 error=false
	// epoch seconds: 2023-11-14T22:13:20Z
	// epoch milliseconds: 2023-11-14T22:13:20.123Z
	// float32: 1969-12-31T23:59:59Z
	// float64: 2023-11-14T22:13:20.123Z
}

func ExampleFrom_ambiguous() {
	dayFirst, _ := date.From("03-04-2026", date.DayFirst)
	monthFirst, _ := date.From("03-04-2026", date.MonthFirst)
	fmt.Println("day first:", dayFirst.Format("2006-01-02"))
	fmt.Println("month first:", monthFirst.Format("2006-01-02"))

	// Output:
	// day first: 2026-04-03
	// month first: 2026-03-04
}

func ExampleFrom_errors() {
	var nilTime *time.Time
	var nilString *string
	inputs := []any{
		nil, nilTime, nilString, "not-a-date", "2026-02-30",
		int64(1_734_412_800_123), uint64(math.MaxInt64) + 1, struct{}{},
	}

	for _, input := range inputs {
		_, err := date.From(input)
		fmt.Printf("%-12T error=%t\n", input, err != nil)
	}

	_, invalidOrder := date.From("2026-08-31", date.Order(99))
	_, multipleOrders := date.From("2026-08-31", date.YearFirst, date.DayFirst)
	fmt.Println("invalid order:", invalidOrder != nil)
	fmt.Println("multiple orders:", multipleOrders != nil)

	// Output:
	// <nil>        error=true
	// *time.Time   error=true
	// *string      error=true
	// string       error=true
	// string       error=true
	// int64        error=true
	// uint64       error=true
	// struct {}    error=true
	// invalid order: true
	// multiple orders: true
}

func ExampleParseAndFormat() {
	fmt.Println(date.ParseAndFormat("2026-08-31", ""))
	fmt.Println(date.ParseAndFormat("2026/08/31", "2006-01-02"))
	fmt.Println(date.ParseAndFormat("20260831", "02 Jan 2006"))
	fmt.Printf("%q\n", date.ParseAndFormat("not-a-date", "2006-01-02"))

	// Output:
	// 2026-08-31
	// 2026-08-31
	// 31 Aug 2026
	// ""
}

func ExampleCleanAndParse() {
	options := date.DefaultCleanParseOptions()
	options.Location = time.FixedZone("TEST-0500", -5*60*60)

	dateOnly, _ := date.CleanAndParse("2026/08/31", nil)
	dateTime, _ := date.CleanAndParse("2026-08-31 15:00:00", options)
	epochSeconds, _ := date.CleanAndParse("1700000000", nil)
	epochMilliseconds, _ := date.CleanAndParse("1700000000123", nil)

	fmt.Println(dateOnly.Format(time.RFC3339Nano))
	fmt.Println(dateTime.Format(time.RFC3339Nano))
	fmt.Println(epochSeconds.Format(time.RFC3339Nano))
	fmt.Println(epochMilliseconds.Format(time.RFC3339Nano))

	// Output:
	// 2026-08-31T00:00:00Z
	// 2026-08-31T20:00:00Z
	// 2023-11-14T22:13:20Z
	// 2023-11-14T22:13:20.123Z
}

func ExampleCleanAndParseAsString() {
	formatted := date.CleanAndParseAsString(
		"2023-11-14T22:13:20Z",
		"20060102",
	)
	fmt.Println(formatted)

	// Output: 20231114
}

func ExampleStrictYYYYMMDDOptions() {
	options := date.StrictYYYYMMDDOptions()
	value, _ := date.CleanAndParse("20240229T150000", options)
	fmt.Println(value.Format(time.RFC3339))

	// Output: 2024-02-29T00:00:00Z
}
