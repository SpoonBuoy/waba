package medical

import (
	"time"

	bookings "github.com/SpoonBuoy/waba/bookings/interfaces"
)

func HandleDbErr(err error) {}

// MedicalService--ActorService implementation
func (ds MedicalService) Type() string {
	return ds.Name
}

func (ds MedicalService) GetDetails() *bookings.ServiceDetails {
	return nil
}

// DocAppointment--Appointment
func (da DocAppointment) Book() bookings.Appointment {
	//create an appointment
	appt := NewDocAppointment(da.Doctor, da.User, da.Slot)
	return appt
}

// Doctor--Actor
func (d Doctor) GetDetails() string {
	return ""
}
func (d Doctor) GetService() string {
	//returns the service type
	return d.Service.Type()
}
func (d Doctor) NextFreeSlot() bookings.Slot {
	//gets next free slot for this doc
	return d.Slots[0]
}
func (d Doctor) AddAppointment(apt bookings.Appointment) {
	//adds appointment to doc
	var docApt DocAppointment = apt.(DocAppointment)
	//apt.Book(d.Name, docApt.User, docApt.Slot)
	docApt.Book()
	d.Appointments = append(d.Appointments, docApt)
}

func (d Doctor) GetAllAppointments() []bookings.Appointment {
	var res []bookings.Appointment
	for _, appt := range d.Appointments {
		res = append(res, appt)
	}
	return res
}

func (d Doctor) GetSlots(from time.Time, to time.Time) []bookings.Slot {
	var res []bookings.Slot
	for _, slot := range d.Slots {
		res = append(res, slot)
	}
	return res
}
func (d Doctor) SlotFactory(from time.Time, to time.Time, duration time.Duration) {
	//creates a slot factory at first time for a doc
	var slots []DocSlot
	//while from is less than to
	for from.Before(to) {
		thisSlotTo := from.Add(duration)
		slot := NewDocSlot(from, thisSlotTo)
		from = thisSlotTo
		slots = append(slots, slot)
	}
	d.Slots = slots
}

func (s DocSlot) Book() {
	s.Available = false
	return
}
func (s DocSlot) Free() {
	s.Available = true
	return
}
func (s DocSlot) IsAvailable() bool {
	return s.Available
}
func (s DocSlot) FromTo() (time.Time, time.Time) {
	return s.From, s.To
}
