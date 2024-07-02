package bookings

import "time"

// slot -- self explanatory
type Slot interface {
	Book()
	Free()
	IsAvailable() bool
	FromTo() (time.Time, time.Time)
}
