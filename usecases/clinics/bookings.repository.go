package clinics

import (
	"time"

	typeinterfaces "github.com/SpoonBuoy/waba/usecases/bookings/type-interfaces"
)

// func (cr clinicRepo) GetBookings(ctx *gin.Context) {
// 	return
// }
// func (cr clinicRepo) Book(ctx *gin.Context) {
// 	return
// }
// func (cr clinicRepo) UpdateAvailability(ctx *gin.Context) {
// 	return
// }
// func (cr clinicRepo) GetActor(ctx *gin.Context) {
// 	return
// }
// func (cr clinicRepo) GetActors(ctx *gin.Context) {
// 	return
// }
// func (cr clinicRepo) AddActor(ctx *gin.Context) {
// 	return
// }
// func (cr clinicRepo) UpdateActor(ctx *gin.Context) {
// 	return
// }
// func (cr clinicRepo) DisableActor(ctx *gin.Context) {
// 	return
// }

func (da DocAppointment) Book() {
	//books an appointment
}

func (d Doctor) BookSlot(slot typeinterfaces.Slot) {
	//book a slot

}
func (d Doctor) FreeSlot() {
	//free a slot
}
func (d Doctor) GetService() {
	//returns the service type
}
func (d Doctor) NextFreeSlot() typeinterfaces.Slot {
	//gets next free slot for this doc
	return d.Slots[0]
}
func (d Doctor) AddAppointment(apt typeinterfaces.Appointment) {
	//adds appointment to doc
	d.Appointments = append(d.Appointments, apt)
}
func (d Doctor) GetAllAppointments() []typeinterfaces.Appointment {
	return d.Appointments
}
func (d *Doctor) SlotFactory(from time.Time, to time.Time, duration time.Duration) {
	//creates a slot factory at first time for a doc
	var slots []Slot
	//while from is less than to
	for from.Before(to) {
		thisSlotTo := from.Add(duration)
		slot := NewSlot(from, thisSlotTo)
		from = thisSlotTo
		slots = append(slots, slot)
	}
	d.Slots = slots
}

func (mb MedicalBusiness) AddActor(doc typeinterfaces.Actor) {
	//adds doctor
}

func (mb MedicalBusiness) GetActor(i int) typeinterfaces.Actor {
	//gets ith doctor
	return mb.Doctors[i]
}

func (mb MedicalBusiness) GetAllActors() []typeinterfaces.Actor {
	//gets all actors
	return mb.Doctors
}

func (s *Slot) IsAvailable() bool {
	return s.Available
}
func (s *Slot) FromTo() (time.Time, time.Time) {
	return s.From, s.To
}

// func main() {
// 	myClinic := NewMedicalBusiness()

// 	//lets create a dermoservice
// 	dermo := NewDermoService("Dermo")

// 	//lets create a dermo doc
// 	myDermoDoc := NewDoctor(dermo, "shadab")

// 	//lets create some slots for doc each 1hr duration for 24 hours
// 	myDermoDoc.SlotFactory(time.Now(), time.Now().Add(time.Hour*24), time.Hour)

// 	//lets add some dermo doctors to it
// 	for i := 0; i < 2; i++ {
// 		//create a dermo service
// 		myClinic.AddActor(myDermoDoc)
// 	}

// 	//lets book 2 appointments for our shadab

// 	for i := 0; i < 2; i++ {
// 		//let's get next free slot for doc
// 		nextFreeSlot := myDermoDoc.NextFreeSlot()

// 		//will book a slot
// 		myDermoDoc.BookSlot(nextFreeSlot)

// 		//create an appointment
// 		apt := NewDocAppointment("shadab", "user1", nextFreeSlot)
// 		myDermoDoc.AddAppointment(apt)
// 	}

// 	//check all the appointments of doc
// 	apts := myDermoDoc.GetAllAppointments()

// 	fmt.Println(apts)

// }
