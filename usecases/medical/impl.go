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
	d.Appointments = append(d.Appointments, apt.(DocAppointment))
}

func (d Doctor) GetAllAppointments() []bookings.Appointment {
	var res []bookings.Appointment
	for _, appt := range d.Appointments {
		res = append(res, appt)
	}
	return res
}

func (d Doctor) GetSlots() []bookings.Slot {
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

func (s DocSlot) IsAvailable() bool {
	return s.Available
}
func (s DocSlot) FromTo() (time.Time, time.Time) {
	return s.From, s.To
}
