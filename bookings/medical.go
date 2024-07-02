package main

// import (
// 	"fmt"
// 	"time"

// 	bookings "github.com/SpoonBuoy/waba/bookings/interfaces"
// )

// // should implement ActorService
// type DermoService struct {
// 	Name string
// }

// func (ds DermoService) Type() string {
// 	return ds.Name
// }

// func NewDermoService(name string) bookings.ActorService {
// 	return DermoService{Name: name}
// }

// // should implement appointments
// type DocAppointment struct {
// 	Doctor string
// 	User   string
// 	Slot   bookings.Slot
// }

// func NewDocAppointment(doc string, user string, slot bookings.Slot) bookings.Appointment {
// 	return DocAppointment{
// 		Doctor: doc,
// 		User:   user,
// 		Slot:   slot,
// 	}
// }
// func (da DocAppointment) Book() bookings.Appointment {
// 	//books an appointment
// 	return nil
// }

// // it should implement actor
// type Doctor struct {
// 	Name         string
// 	Service      bookings.ActorService
// 	Details      string
// 	Slots        []bookings.Slot
// 	Appointments []bookings.Appointment
// }

// func NewDoctor(svc bookings.ActorService, name string) bookings.Actor {
// 	return &Doctor{
// 		Service: svc,
// 		Name:    name,
// 	}
// }
// func (d Doctor) BookSlot(slot bookings.Slot) {
// 	//book a slot

// }
// func (d Doctor) GetSlots() []bookings.Slot {
// 	return d.Slots
// }
// func (d Doctor) RemoveExpiredSlots() {

// }
// func (d Doctor) FreeSlot(s bookings.Slot) {
// 	//free a slot
// }
// func (d Doctor) GetService() string {
// 	//returns the service type
// 	return ""
// }
// func (d Doctor) NextFreeSlot() bookings.Slot {
// 	//gets next free slot for this doc
// 	return d.Slots[0]
// }
// func (d Doctor) AddAppointment(apt bookings.Appointment) {
// 	//adds appointment to doc
// 	d.Appointments = append(d.Appointments, apt)
// }
// func (d Doctor) GetAllAppointments() []bookings.Appointment {
// 	return d.Appointments
// }
// func (d *Doctor) SlotFactory(from time.Time, to time.Time, duration time.Duration) {
// 	//creates a slot factory at first time for a doc
// 	var slots []bookings.Slot
// 	//while from is less than to
// 	for from.Before(to) {
// 		thisSlotTo := from.Add(duration)
// 		slot := NewSlot(from, thisSlotTo)
// 		from = thisSlotTo
// 		slots = append(slots, &slot)
// 	}
// 	d.Slots = slots
// }

// // it should implement BusinessController
// type MedicalBusiness struct {
// 	Name    string
// 	Doctors []bookings.Actor
// }

// func (mb MedicalBusiness) AddActor(doc bookings.Actor, bid int) {
// 	//adds doctor
// }

// func (mb MedicalBusiness) GetActor(i int) bookings.Actor {
// 	//gets ith doctor
// 	return mb.Doctors[i]
// }

// func (mb MedicalBusiness) GetAllActors(bid int) []bookings.Actor {
// 	//gets all actors
// 	return mb.Doctors
// }
// func (mb MedicalBusiness) GetAllAppointments(bid int) []bookings.Appointment {
// 	return nil
// }
// func (mb MedicalBusiness) GetAllServices(bid int) []bookings.ActorService {
// 	return nil
// }
// func (mb MedicalBusiness) AddService(svc bookings.ActorService, bid int) {

// }
// func NewMedicalBusiness() bookings.BusinessRepo {
// 	return MedicalBusiness{}
// }

// type Slot struct {
// 	From      time.Time
// 	To        time.Time
// 	Available bool
// }

// func (s *Slot) IsAvailable() bool {
// 	return s.Available
// }
// func (s *Slot) FromTo() (time.Time, time.Time) {
// 	return s.From, s.To
// }
// func NewSlot(from time.Time, to time.Time) Slot {
// 	return Slot{
// 		From:      from,
// 		To:        to,
// 		Available: true,
// 	}
// }

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
// 		myClinic.AddActor(myDermoDoc, 0)
// 	}

// 	//lets book 2 appointments for our shadab

// 	for i := 0; i < 2; i++ {
// 		//let's get next free slot for doc
// 		nextFreeSlot := myDermoDoc.NextFreeSlot()

// 		//will book a slot
// 		// myDermoDoc.BookSlot(nextFreeSlot)

// 		//create an appointment
// 		apt := NewDocAppointment("shadab", "user1", nextFreeSlot)
// 		myDermoDoc.AddAppointment(apt)
// 	}

// 	//check all the appointments of doc
// 	apts := myDermoDoc.GetAllAppointments()

// 	fmt.Println(apts)

// }
