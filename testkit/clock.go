package testkit

import (
	"sync"
	"time"
)

// FakeClock is a manually controlled, goroutine-safe clock.
//
// Production code should not import testkit. Have it depend on its own small
// interface or a func, and pass clock.Now in tests:
//
//	type Service struct{ now func() time.Time } // time.Now in production
//	svc := Service{now: clock.Now}              // in tests
type FakeClock struct {
	mu  sync.Mutex
	now time.Time
}

// NewFakeClock returns a FakeClock set to start. A zero start uses
// 2024-01-01 00:00:00 UTC so tests are reproducible.
func NewFakeClock(start time.Time) *FakeClock {
	if start.IsZero() {
		start = time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)
	}

	return &FakeClock{now: start}
}

// Now returns the clock's current time.
func (c *FakeClock) Now() time.Time {
	c.mu.Lock()
	defer c.mu.Unlock()

	return c.now
}

// Advance moves the clock forward by d and returns the new time.
func (c *FakeClock) Advance(d time.Duration) time.Time {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.now = c.now.Add(d)

	return c.now
}

// Set moves the clock to t.
func (c *FakeClock) Set(t time.Time) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.now = t
}
