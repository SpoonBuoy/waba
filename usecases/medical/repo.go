package medical

import bookings "github.com/SpoonBuoy/waba/bookings/interfaces"

// Medical Bussiness -- Business
func (mb MedicalBusiness) AddActor(doc bookings.Actor, bid int) {
	//adds doctor
	var clinic MedicalBusiness
	var doctor Doctor = doc.(Doctor)
	err := mb.Db.Model(&clinic).Association("Doctor").Append(doctor)
	if err != nil {
		HandleDbErr(err)
	}
}
func (mb MedicalBusiness) AddService(svc bookings.ActorService, bid int) {
	var medSvc MedicalService = svc.(MedicalService)
	var clinic MedicalBusiness
	err := mb.Db.Model(&clinic).Association("Service").Append(medSvc)
	if err != nil {
		HandleDbErr(err)
	}
}
func (mb MedicalBusiness) GetAllActors(bid int) []bookings.Actor {
	var docs []bookings.Actor
	var clinic MedicalBusiness
	err := mb.Db.Where("id = ?", bid).Preload("Doctor").First(&clinic).Error
	if err != nil {
		HandleDbErr(err)
		return nil
	}
	for _, doc := range clinic.Doctors {
		docs = append(docs, doc)
	}
	return docs
}
func (mb MedicalBusiness) GetAllAppointments(bid int) []bookings.Appointment {
	var docAppts []bookings.Appointment
	var clinic MedicalBusiness
	err := mb.Db.Where("id = ?", bid).Preload("DocAppointment").First(&clinic).Error
	if err != nil {
		HandleDbErr(err)
		return nil
	}
	for _, appt := range clinic.Appointments {
		docAppts = append(docAppts, appt)
	}
	return docAppts
}
func (mb MedicalBusiness) GetAllServices(bid int) []bookings.ActorService {
	var svcs []bookings.ActorService
	var clinic MedicalBusiness
	err := mb.Db.Where("id = ?", bid).Preload("Service").First(&clinic).Error
	if err != nil {
		HandleDbErr(err)
		return nil
	}
	for _, svc := range clinic.Services {
		svcs = append(svcs, svc)
	}
	return svcs
}

func (mb MedicalBusiness) GetActor(id int) bookings.Actor {
	var doc Doctor
	err := mb.Db.Where("id = ?", id).Preload("DocAppointment").Preload("DocSlot").First(&doc).Error
	if err != nil {
		HandleDbErr(err)
		return nil
	}
	return &doc
}
