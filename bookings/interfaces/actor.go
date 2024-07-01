package bookings

import "time"

// Actor is someone who provides service
type Actor interface {
	BookSlot(Slot)
	FreeSlot(Slot)
	NextFreeSlot() Slot
	GetService() string
	AddAppointment(Appointment)
	GetAllAppointments() []Appointment
	SlotFactory(from time.Time, to time.Time, duration time.Duration)
	GetSlots() []Slot
	RemoveExpiredSlots()
}
