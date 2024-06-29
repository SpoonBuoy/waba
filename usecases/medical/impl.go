package medical

import (
	"time"

	bookings "github.com/SpoonBuoy/waba/bookings/interfaces"
)

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
func (mb MedicalBusiness) AddActor(doc bookings.Actor) {
	//adds doctor
}

func (mb MedicalBusiness) GetActor(i int) bookings.Actor {
	//gets ith doctor
	return mb.Doctors[i]
}

func (mb MedicalBusiness) GetAllActors() []bookings.Actor {
	//gets all actors
	return mb.Doctors
}
func (mb MedicalBusiness) GetAllAppointments() []bookings.Appointment {
	return nil
}
func (mb MedicalBusiness) GetAllServices() []bookings.ActorService {
	return nil
}
func (mb MedicalBusiness) AddService(svc bookings.ActorService) {

}
func NewMedicalBusiness() bookings.Business {
	return MedicalBusiness{}
}

// DocSlot--Slot
func (s *DocSlot) IsAvailable() bool {
	return s.Available
}
func (s *DocSlot) FromTo() (time.Time, time.Time) {
	return s.From, s.To
}
