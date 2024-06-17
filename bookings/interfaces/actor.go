package bookings

import "time"

// Actor is someone who provides service
type Actor interface {
	BookSlot(Slot)
	FreeSlot()
	NextFreeSlot() Slot
	GetService()
	AddAppointment(Appointment)
	GetAllAppointments() []Appointment
	SlotFactory(from time.Time, to time.Time, duration time.Duration)
}
