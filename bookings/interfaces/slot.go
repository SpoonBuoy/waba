package bookings

import "time"

// slot -- self explanatory
type Slot interface {
	IsAvailable() bool
	FromTo() (time.Time, time.Time)
}
