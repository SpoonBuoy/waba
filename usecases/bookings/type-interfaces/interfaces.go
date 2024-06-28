package typeinterfaces

import (
	"time"
)

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

type ActorServices interface {
	Type() string
}

// appointment is a contract between actor and user
type Appointment interface {
	Book()
}

// manages all the operations for business
type BusinessController interface {
	AddActor(Actor)
	GetActor(int) Actor
	GetAllActors() []Actor
}

// slot -- self explanatory
type Slot interface {
	IsAvailable() bool
	FromTo() (time.Time, time.Time)
}
