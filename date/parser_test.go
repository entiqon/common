package date_test

import (
	"math"
	"testing"
	"time"

	"github.com/entiqon/common/date"
	"github.com/entiqon/common/date/internal"
	"github.com/entiqon/common/test"
)

func TestDate(t *testing.T) {
	t.Run("From", func(t *testing.T) {
		t.Run("Error", func(t *testing.T) {
			t.Run("Nil", func(t *testing.T) {
				var value any

				_, err := date.From(value)
				if err == nil {
					t.Fatal("expected error, got nil")
				}
			})

			t.Run("NilTimePointer", func(t *testing.T) {
				var value *time.Time

				_, err := date.From(value)
				if err == nil {
					t.Fatal("expected error, got nil")
				}
			})

			t.Run("NilStringPointer", func(t *testing.T) {
				var value *string

				_, err := date.From(value)
				if err == nil {
					t.Fatal("expected error, got nil")
				}
			})

			t.Run("EmptyString", func(t *testing.T) {
				_, err := date.From(" ")

				if err == nil {
					t.Fatal("expected error, got nil")
				}
			})

			t.Run("UnsupportedType", func(t *testing.T) {
				_, err := date.From(struct{}{})

				if err == nil {
					t.Fatal("expected error, got nil")
				}
			})

			t.Run("InvalidOrder", func(t *testing.T) {
				_, err := date.From(
					"2026-08-31",
					date.Order(99),
				)

				if err == nil {
					t.Fatal("expected error, got nil")
				}
			})

			t.Run("MultipleOrders", func(t *testing.T) {
				_, err := date.From(
					"2026-08-31",
					date.YearFirst,
					date.DayFirst,
				)

				if err == nil {
					t.Fatal("expected error, got nil")
				}
			})

			t.Run("UnrecognizedString", func(t *testing.T) {
				_, err := date.From("not-a-date")

				if err == nil {
					t.Fatal("expected error, got nil")
				}
			})

			t.Run("InvalidDate", func(t *testing.T) {
				_, err := date.From("2026-02-30")

				if err == nil {
					t.Fatal("expected error, got nil")
				}
			})

			t.Run("IncompleteCompactDate", func(t *testing.T) {
				_, err := date.From("2026083")

				if err == nil {
					t.Fatal("expected error, got nil")
				}
			})

			t.Run("InvalidCompactDate", func(t *testing.T) {
				_, err := date.From("20260230")

				if err == nil {
					t.Fatal("expected invalid compact date error, got nil")
				}
			})

			t.Run("SignedMilliseconds", func(t *testing.T) {
				_, err := date.From(int64(1_734_412_800_123))

				if err == nil {
					t.Fatal("expected millisecond guard error, got nil")
				}
			})

			t.Run("UnsignedMilliseconds", func(t *testing.T) {
				_, err := date.From(uint64(1_734_412_800_123))

				if err == nil {
					t.Fatal("expected millisecond guard error, got nil")
				}
			})

			t.Run("UnsignedOverflow", func(t *testing.T) {
				_, err := date.From(uint64(math.MaxInt64) + 1)

				if err == nil {
					t.Fatal("expected overflow error, got nil")
				}
			})
		})

		t.Run("Time", func(t *testing.T) {
			t.Run("Value", func(t *testing.T) {
				want := time.Date(
					2026,
					time.August,
					31,
					15,
					0,
					0,
					0,
					time.UTC,
				)

				got, err := date.From(want)
				if err != nil {
					t.Fatal(err)
				}

				test.AssertEqual(t, want, got)
			})

			t.Run("Pointer", func(t *testing.T) {
				want := time.Date(
					2026,
					time.August,
					31,
					15,
					0,
					0,
					0,
					time.UTC,
				)

				got, err := date.From(&want)
				if err != nil {
					t.Fatal(err)
				}

				test.AssertEqual(t, want, got)
			})
		})

		t.Run("String", func(t *testing.T) {
			t.Run("Pointer", func(t *testing.T) {
				value := "2026-08-31"
				want := time.Date(
					2026,
					time.August,
					31,
					0,
					0,
					0,
					0,
					time.UTC,
				)

				got, err := date.From(&value)
				assertTime(t, got, want, err)
			})

			t.Run("Bytes", func(t *testing.T) {
				value := []byte("2026-08-31")
				want := time.Date(
					2026,
					time.August,
					31,
					0,
					0,
					0,
					0,
					time.UTC,
				)

				got, err := date.From(value)
				assertTime(t, got, want, err)
			})

			t.Run("YearFirst", func(t *testing.T) {
				t.Run("Date", func(t *testing.T) {
					tests := []struct {
						name  string
						value string
					}{
						{
							name:  "Dash",
							value: "2026-08-31",
						},
						{
							name:  "Slash",
							value: "2026/08/31",
						},
						{
							name:  "Compact",
							value: "20260831",
						},
					}

					want := time.Date(
						2026,
						time.August,
						31,
						0,
						0,
						0,
						0,
						time.UTC,
					)

					for _, tc := range tests {
						t.Run(tc.name, func(t *testing.T) {
							got, err := date.From(tc.value)
							assertTime(t, got, want, err)
						})
					}
				})

				t.Run("DateTime", func(t *testing.T) {
					originalLocal := time.Local
					time.Local = time.FixedZone(
						"TEST-0500",
						-5*60*60,
					)
					t.Cleanup(func() {
						time.Local = originalLocal
					})

					tests := []struct {
						name  string
						value string
					}{
						{
							name:  "Space",
							value: "2026-08-31 15:00:00",
						},
						{
							name:  "T",
							value: "2026-08-31T15:00:00",
						},
						{
							name:  "Slash",
							value: "2026/08/31 15:00:00",
						},
						{
							name:  "Compact",
							value: "20260831T150000",
						},
					}

					want := time.Date(
						2026,
						time.August,
						31,
						20,
						0,
						0,
						0,
						time.UTC,
					)

					for _, tc := range tests {
						t.Run(tc.name, func(t *testing.T) {
							got, err := date.From(tc.value)
							assertTime(t, got, want, err)
						})
					}
				})
			})

			t.Run("DayFirst", func(t *testing.T) {
				t.Run("Date", func(t *testing.T) {
					tests := []struct {
						name  string
						value string
					}{
						{
							name:  "Dash",
							value: "31-08-2026",
						},
						{
							name:  "Slash",
							value: "31/08/2026",
						},
						{
							name:  "Compact",
							value: "31082026",
						},
					}

					want := time.Date(
						2026,
						time.August,
						31,
						0,
						0,
						0,
						0,
						time.UTC,
					)

					for _, tc := range tests {
						t.Run(tc.name, func(t *testing.T) {
							got, err := date.From(
								tc.value,
								date.DayFirst,
							)
							assertTime(t, got, want, err)
						})
					}
				})

				t.Run("DateTime", func(t *testing.T) {
					originalLocal := time.Local
					time.Local = time.FixedZone(
						"TEST-0500",
						-5*60*60,
					)
					t.Cleanup(func() {
						time.Local = originalLocal
					})

					tests := []struct {
						name  string
						value string
					}{
						{
							name:  "Dash",
							value: "31-08-2026 15:00:00",
						},
						{
							name:  "Slash",
							value: "31/08/2026 15:00:00",
						},
						{
							name:  "T",
							value: "31-08-2026T15:00:00",
						},
						{
							name:  "Compact",
							value: "31082026T150000",
						},
					}

					want := time.Date(
						2026,
						time.August,
						31,
						20,
						0,
						0,
						0,
						time.UTC,
					)

					for _, tc := range tests {
						t.Run(tc.name, func(t *testing.T) {
							got, err := date.From(
								tc.value,
								date.DayFirst,
							)
							assertTime(t, got, want, err)
						})
					}
				})
			})

			t.Run("MonthFirst", func(t *testing.T) {
				t.Run("Date", func(t *testing.T) {
					tests := []struct {
						name  string
						value string
					}{
						{
							name:  "Dash",
							value: "08-31-2026",
						},
						{
							name:  "Slash",
							value: "08/31/2026",
						},
						{
							name:  "Compact",
							value: "08312026",
						},
					}

					want := time.Date(
						2026,
						time.August,
						31,
						0,
						0,
						0,
						0,
						time.UTC,
					)

					for _, tc := range tests {
						t.Run(tc.name, func(t *testing.T) {
							got, err := date.From(
								tc.value,
								date.MonthFirst,
							)
							assertTime(t, got, want, err)
						})
					}
				})

				t.Run("DateTime", func(t *testing.T) {
					originalLocal := time.Local
					time.Local = time.FixedZone(
						"TEST-0500",
						-5*60*60,
					)
					t.Cleanup(func() {
						time.Local = originalLocal
					})

					tests := []struct {
						name  string
						value string
					}{
						{
							name:  "Dash",
							value: "08-31-2026 15:00:00",
						},
						{
							name:  "Slash",
							value: "08/31/2026 15:00:00",
						},
						{
							name:  "T",
							value: "08-31-2026T15:00:00",
						},
						{
							name:  "Compact",
							value: "08312026T150000",
						},
					}

					want := time.Date(
						2026,
						time.August,
						31,
						20,
						0,
						0,
						0,
						time.UTC,
					)

					for _, tc := range tests {
						t.Run(tc.name, func(t *testing.T) {
							got, err := date.From(
								tc.value,
								date.MonthFirst,
							)
							assertTime(t, got, want, err)
						})
					}
				})
			})

			t.Run("Zoned", func(t *testing.T) {
				t.Run("UTC", func(t *testing.T) {
					got, err := date.From(
						"2026-08-31T15:00:00Z",
					)
					want := time.Date(
						2026,
						time.August,
						31,
						15,
						0,
						0,
						0,
						time.UTC,
					)

					assertTime(t, got, want, err)
				})

				t.Run("NegativeOffset", func(t *testing.T) {
					got, err := date.From(
						"2026-08-31T15:00:00-04:00",
					)
					if err != nil {
						t.Fatalf("unexpected error: %v", err)
					}

					_, offset := got.Zone()
					if offset != -4*60*60 {
						t.Fatalf(
							"expected offset %d, got %d",
							-4*60*60,
							offset,
						)
					}

					if got.Hour() != 15 {
						t.Fatalf(
							"expected hour 15, got %d",
							got.Hour(),
						)
					}
				})

				t.Run("PositiveOffset", func(t *testing.T) {
					got, err := date.From(
						"2026-08-31T15:00:00+04:00",
					)
					if err != nil {
						t.Fatalf("unexpected error: %v", err)
					}

					_, offset := got.Zone()
					if offset != 4*60*60 {
						t.Fatalf(
							"expected offset %d, got %d",
							4*60*60,
							offset,
						)
					}
				})

				t.Run("CompactUTC", func(t *testing.T) {
					got, err := date.From(
						"20260831T150000Z",
					)
					want := time.Date(
						2026,
						time.August,
						31,
						15,
						0,
						0,
						0,
						time.UTC,
					)

					assertTime(t, got, want, err)
				})

				t.Run("CompactOffset", func(t *testing.T) {
					got, err := date.From(
						"20260831T150000-0400",
					)
					if err != nil {
						t.Fatalf("unexpected error: %v", err)
					}

					_, offset := got.Zone()
					if offset != -4*60*60 {
						t.Fatalf(
							"expected offset %d, got %d",
							-4*60*60,
							offset,
						)
					}
				})

				t.Run("RFC1123", func(t *testing.T) {
					got, err := date.From(
						"Mon, 31 Aug 2026 15:00:00 GMT",
					)
					if err != nil {
						t.Fatalf("unexpected error: %v", err)
					}

					if got.Location().String() != "GMT" {
						t.Fatalf(
							"expected GMT, got %s",
							got.Location(),
						)
					}
				})
			})

			t.Run("Ambiguous", func(t *testing.T) {
				t.Run("DayFirst", func(t *testing.T) {
					got, err := date.From(
						"03-04-2026",
						date.DayFirst,
					)
					want := time.Date(
						2026,
						time.April,
						3,
						0,
						0,
						0,
						0,
						time.UTC,
					)

					assertTime(t, got, want, err)
				})

				t.Run("MonthFirst", func(t *testing.T) {
					got, err := date.From(
						"03-04-2026",
						date.MonthFirst,
					)
					want := time.Date(
						2026,
						time.March,
						4,
						0,
						0,
						0,
						0,
						time.UTC,
					)

					assertTime(t, got, want, err)
				})
			})
		})

		t.Run("EpochString", func(t *testing.T) {
			t.Run("Seconds", func(t *testing.T) {
				got, err := date.From("1700000000")
				want := time.Unix(1_700_000_000, 0).UTC()

				assertTime(t, got, want, err)
			})

			t.Run("Milliseconds", func(t *testing.T) {
				got, err := date.From("1700000000123")
				want := time.Unix(
					1_700_000_000,
					123_000_000,
				).UTC()

				assertTime(t, got, want, err)
			})
		})

		t.Run("Integer", func(t *testing.T) {
			tests := []struct {
				name  string
				value any
				want  time.Time
			}{
				{
					name:  "Int",
					value: int(42),
					want:  time.Unix(42, 0).UTC(),
				},
				{
					name:  "Int8",
					value: int8(-7),
					want:  time.Unix(-7, 0).UTC(),
				},
				{
					name:  "Int16",
					value: int16(12345),
					want:  time.Unix(12345, 0).UTC(),
				},
				{
					name:  "Int32",
					value: int32(2_147_000_000),
					want:  time.Unix(2_147_000_000, 0).UTC(),
				},
				{
					name:  "Int64",
					value: int64(1_700_000_000),
					want:  time.Unix(1_700_000_000, 0).UTC(),
				},
				{
					name:  "Uint",
					value: uint(42),
					want:  time.Unix(42, 0).UTC(),
				},
				{
					name:  "Uint8",
					value: uint8(7),
					want:  time.Unix(7, 0).UTC(),
				},
				{
					name:  "Uint16",
					value: uint16(12345),
					want:  time.Unix(12345, 0).UTC(),
				},
				{
					name:  "Uint32",
					value: uint32(2_147_000_000),
					want:  time.Unix(2_147_000_000, 0).UTC(),
				},
				{
					name:  "Uint64",
					value: uint64(1_700_000_000),
					want:  time.Unix(1_700_000_000, 0).UTC(),
				},
				{
					name:  "Uintptr",
					value: uintptr(42),
					want:  time.Unix(42, 0).UTC(),
				},
			}

			for _, tc := range tests {
				t.Run(tc.name, func(t *testing.T) {
					got, err := date.From(tc.value)
					assertTime(t, got, tc.want, err)
				})
			}
		})

		t.Run("Float", func(t *testing.T) {
			t.Run("Float32", func(t *testing.T) {
				got, err := date.From(
					float32(1_700_000_000.123),
				)
				want := time.Unix(
					1_700_000_000,
					0,
				).UTC()

				assertTime(t, got, want, err)
			})

			t.Run("Float32Negative", func(t *testing.T) {
				got, err := date.From(float32(-1.987))
				want := time.Unix(-1, 0).UTC()

				assertTime(t, got, want, err)
			})

			t.Run("Float64", func(t *testing.T) {
				got, err := date.From(
					1_700_000_000.123,
				)
				want := time.Unix(
					1_700_000_000,
					123_000_000,
				).UTC()

				assertTime(t, got, want, err)
			})

			t.Run("Float64Carry", func(t *testing.T) {
				got, err := date.From(
					1_700_000_000.9995,
				)
				want := time.Unix(
					1_700_000_001,
					0,
				).UTC()

				assertTime(t, got, want, err)
			})

			t.Run("Float64Negative", func(t *testing.T) {
				got, err := date.From(-1.0005)
				want := time.Unix(
					-2,
					999_000_000,
				).UTC()

				assertTime(t, got, want, err)
			})
		})
	})

	t.Run("ParseAndFormat", func(t *testing.T) {
		tests := []struct {
			name   string
			value  string
			layout string
			want   string
		}{
			{
				name:   "DefaultLayout",
				value:  "2026-08-31",
				layout: "",
				want:   "2026-08-31",
			},
			{
				name:   "Dash",
				value:  "2026-08-31",
				layout: "2006-01-02",
				want:   "2026-08-31",
			},
			{
				name:   "Slash",
				value:  "2026/08/31",
				layout: "2006-01-02",
				want:   "2026-08-31",
			},
			{
				name:   "Compact",
				value:  "20260831",
				layout: "2006-01-02",
				want:   "2026-08-31",
			},
			{
				name:   "CustomOutput",
				value:  "2026-08-31",
				layout: "02 Jan 2006",
				want:   "31 Aug 2026",
			},
			{
				name:   "Invalid",
				value:  "not-a-date",
				layout: "2006-01-02",
				want:   "",
			},
			{
				name:   "Empty",
				value:  "",
				layout: "2006-01-02",
				want:   "",
			},
		}

		for _, tc := range tests {
			t.Run(tc.name, func(t *testing.T) {
				got := date.ParseAndFormat(
					tc.value,
					tc.layout,
				)
				if got != tc.want {
					t.Fatalf(
						"expected %q, got %q",
						tc.want,
						got,
					)
				}
			})
		}
	})

	t.Run("DecimalDigits", func(t *testing.T) {
		tests := []struct {
			name  string
			value uint64
			want  int
		}{
			{
				name:  "Zero",
				value: 0,
				want:  1,
			},
			{
				name:  "OneDigit",
				value: 7,
				want:  1,
			},
			{
				name:  "TwoDigits",
				value: 42,
				want:  2,
			},
			{
				name:  "TenDigits",
				value: 1_700_000_000,
				want:  10,
			},
		}

		for _, tc := range tests {
			t.Run(tc.name, func(t *testing.T) {
				got := internal.DecimalDigits(int64(tc.value))
				if got != tc.want {
					t.Fatalf(
						"expected %d, got %d",
						tc.want,
						got,
					)
				}
			})
		}
	})
}

func assertTime(
	t *testing.T,
	got time.Time,
	want time.Time,
	err error,
) {
	t.Helper()

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if !got.Equal(want) {
		t.Fatalf("expected %v, got %v", want, got)
	}
}
