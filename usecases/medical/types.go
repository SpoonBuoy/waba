package medical

import (
	"time"

	bookings "github.com/SpoonBuoy/waba/bookings/interfaces"
	"gorm.io/gorm"
)

// should implement ActorService
type DermoService struct {
	Name string
}

func NewDermoService(name string) bookings.ActorService {
	return DermoService{Name: name}
}

// should implement appointments
type DocAppointment struct {
	Doctor string
	User   string
	Slot   bookings.Slot
}

func NewDocAppointment(doc string, user string, slot bookings.Slot) bookings.Appointment {
	return DocAppointment{
		Doctor: doc,
		User:   user,
		Slot:   slot,
	}
}

// it should implement actor
type Doctor struct {
	Name         string
	Service      bookings.ActorService
	Details      string
	Slots        []bookings.Slot
	Appointments []bookings.Appointment
}

func NewDoctor(svc bookings.ActorService, name string, details string, slots []bookings.Slot, appts []bookings.Appointment) bookings.Actor {
	return &Doctor{
		Service:      svc,
		Name:         name,
		Details:      details,
		Slots:        slots,
		Appointments: appts,
	}
}

// it should implement BusinessRepo
type MedicalBusiness struct {
	Name         string
	Appointments []bookings.Appointment
	Doctors      []bookings.Actor
	Services     []bookings.ActorService
	Db           *gorm.DB
}

type DocSlot struct {
	From      time.Time
	To        time.Time
	Available bool
}

func NewDocSlot(from time.Time, to time.Time) DocSlot {
	return DocSlot{
		From:      from,
		To:        to,
		Available: true,
	}
}
