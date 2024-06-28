package clinics

import (
	"time"

	typeinterfaces "github.com/SpoonBuoy/waba/usecases/bookings/type-interfaces"
)

// should implement ActorService
type DermoService struct {
	Name string
}

// should implement appointments
type DocAppointment struct {
	Doctor string
	User   string
	Slot   Slot
}

// it should implement actor
type Doctor struct {
	Name         string
	Service      typeinterfaces.ActorServices
	Details      string
	Slots        []Slot
	Appointments []typeinterfaces.Appointment
}

// it should implement BusinessController
type MedicalBusiness struct {
	Name    string
	Doctors []typeinterfaces.Actor
}
type Slot struct {
	From      time.Time
	To        time.Time
	Available bool
}

func (ds DermoService) Type() string {
	return ds.Name
}

func NewDermoService(name string) typeinterfaces.ActorServices {
	return DermoService{Name: name}
}

func NewDocAppointment(doc string, user string, slot Slot) typeinterfaces.Appointment {
	return DocAppointment{
		Doctor: doc,
		User:   user,
		Slot:   slot,
	}
}

func NewDoctor(svc typeinterfaces.ActorServices, name string) typeinterfaces.Actor {
	return Doctor{
		Service: svc,
		Name:    name,
	}
}

func NewSlot(from time.Time, to time.Time) Slot {
	return Slot{
		From:      from,
		To:        to,
		Available: true,
	}
}
func NewMedicalBusiness() typeinterfaces.BusinessController {
	return MedicalBusiness{}
}
