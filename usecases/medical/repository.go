package medical

import (
	"time"

	bookings "github.com/SpoonBuoy/waba/bookings/interfaces"
)

func (mb MedicalBusiness) GetSlots(docId int, bid int) []bookings.Slot {
	var clinic Clinic
	err := mb.Db.Where("id = ?", bid).Preload("Doctor").First(&clinic).Error
	if err != nil {
		HandleDbErr(err)
		return nil
	}
	doc := docById(clinic.Doctors, docId)
	actor := NewDoctor(nil, doc.Name)
	return actor.GetSlots()
}
func (mb MedicalBusiness) GetAllActors(bid int) []bookings.Actor {
	//gets all actors
	var clinic MedicalBusiness
	err := mb.Db.Where("id = ?", bid).Preload("Doctor").First(&clinic).Error
	if err != nil {
		HandleDbErr(err)
		return nil
	}
	return clinic.Doctors
}
func (mb MedicalBusiness) GetAllAppointments(bid int) []bookings.Appointment {
	var clinic MedicalBusiness
	err := mb.Db.Where("id = ?", bid).Preload("DocAppointment").First(&clinic).Error
	if err != nil {
		HandleDbErr(err)
		return nil
	}
	return clinic.Appointments
}
func (mb MedicalBusiness) GetAllServices(bid int) []bookings.ActorService {
	var clinic MedicalBusiness
	err := mb.Db.Where("id = ?", bid).Preload("DermoService").First(&clinic).Error
	if err != nil {
		HandleDbErr(err)
		return nil
	}
	return clinic.Services
}
func (mb MedicalBusiness) AddService(svc bookings.ActorService, bid int) {

}
func NewMedicalBusiness() bookings.BusinessRepo {
	return MedicalBusiness{}
}

// DocSlot--Slot
func (s *DocSlot) IsAvailable() bool {
	return s.Available
}
func (s *DocSlot) FromTo() (time.Time, time.Time) {
	return s.From, s.To
}
