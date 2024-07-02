package bookings

import "time"

// Actor is someone who provides service
type Actor interface {
	NextFreeSlot() Slot
	GetService() string
	GetDetails() string
	GetSlots(from time.Time, to time.Time) []Slot
	SlotFactory(from time.Time, to time.Time, duration time.Duration)

	AddAppointment(Appointment)
	GetAllAppointments() []Appointment
}
