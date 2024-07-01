package medical

import (
	"time"

	bookings "github.com/SpoonBuoy/waba/bookings/interfaces"
	"gorm.io/gorm"
)

// should implement ActorService
type MedicalService struct {
	Name string
}

func NewMedicalService(name string) bookings.ActorService {
	return MedicalService{Name: name}
}

// should implement appointments
type DocAppointment struct {
	Doctor string
	User   string
	Slot   DocSlot
}

func NewDocAppointment(doc string, user string, slot DocSlot) bookings.Appointment {
	return DocAppointment{
		Doctor: doc,
		User:   user,
		Slot:   slot,
	}
}

// it should implement actor
type Doctor struct {
	Name         string
	Service      MedicalService
	Details      string
	Slots        []DocSlot
	Appointments []DocAppointment
}

func NewDoc(name string, svc MedicalService, details string) Doctor {
	return Doctor{
		Name:    name,
		Service: svc,
		Details: details,
	}
}

// it should implement BusinessRepo
type MedicalBusiness struct {
	Name         string
	Appointments []DocAppointment
	Doctors      []Doctor
	Services     []MedicalService
	Db           *gorm.DB
}

func NewMedicalBusiness(name string, db *gorm.DB) MedicalBusiness {
	return MedicalBusiness{
		Name: name,
		Db:   db,
	}
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
