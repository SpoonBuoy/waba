package medical

import (
	"time"

	bookings "github.com/SpoonBuoy/waba/bookings/interfaces"
	"github.com/gin-gonic/gin"
)

func HandleDbErr(err error) {}

// DermoService--ActorService implementation
func (ds DermoService) Type() string {
	return ds.Name
}

// DocAppointment--Appointment
func (da DocAppointment) Book() {
	//books an appointment
}

// Doctor--Actor
func (d Doctor) BookSlot(slot bookings.Slot) {
	//book a slot

}
func (d Doctor) FreeSlot() {
	//free a slot
}
func (d Doctor) GetService() {
	//returns the service type
}
func (d Doctor) NextFreeSlot() bookings.Slot {
	//gets next free slot for this doc
	return d.Slots[0]
}
func (d Doctor) AddAppointment(apt bookings.Appointment) {
	//adds appointment to doc
	d.Appointments = append(d.Appointments, apt)
}
func (d Doctor) GetAllAppointments() []bookings.Appointment {
	return d.Appointments
}
func (d Doctor) GetSlots() []bookings.Slot {
	return d.Slots
}
func (d *Doctor) SlotFactory(from time.Time, to time.Time, duration time.Duration) {
	//creates a slot factory at first time for a doc
	var slots []bookings.Slot
	//while from is less than to
	for from.Before(to) {
		thisSlotTo := from.Add(duration)
		slot := NewDocSlot(from, thisSlotTo)
		from = thisSlotTo
		slots = append(slots, &slot)
	}
	d.Slots = slots
}

// Medical Bussiness -- Business
func (mb MedicalBusiness) AddActor(c *gin.Context, doc bookings.Actor, bid int) {
	//adds doctor
}

func (mb MedicalBusiness) GetActor(c *gin.Context, id int) bookings.Actor {
	//gets ith doctor
	//return mb.Doctors[i]
	var doc Doctor
	err := mb.Db.Where("id = ?", id).First(&doc).Error
	if err != nil {
		HandleDbErr(err)
		return nil
	}
	return &doc
}

// helper func
func docById(c *gin.Context, docs []bookings.Actor, id int) bookings.Actor {
	return docs[0]
}
func (mb MedicalBusiness) GetSlots(c *gin.Context, docId int, bid int) []bookings.Slot {
	var clinic MedicalBusiness
	err := mb.Db.Where("id = ?", bid).Preload("Doctor").First(&clinic).Error
	if err != nil {
		HandleDbErr(err)
		return nil
	}
	doc := docById(clinic.Doctors, docId)
	return doc.GetSlots()
}
func (mb MedicalBusiness) GetAllActors(c *gin.Context, bid int) []bookings.Actor {
	//gets all actors
	var clinic MedicalBusiness
	err := mb.Db.Where("id = ?", bid).Preload("Doctor").First(&clinic).Error
	if err != nil {
		HandleDbErr(err)
		return nil
	}
	return clinic.Doctors
}
func (mb MedicalBusiness) GetAllAppointments(c *gin.Context, bid int) []bookings.Appointment {
	var clinic MedicalBusiness
	err := mb.Db.Where("id = ?", bid).Preload("DocAppointment").First(&clinic).Error
	if err != nil {
		HandleDbErr(err)
		return nil
	}
	return clinic.Appointments
}
func (mb MedicalBusiness) GetAllServices(c *gin.Context, bid int) []bookings.ActorService {
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
