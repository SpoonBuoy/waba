package service

import (
	bookings "github.com/SpoonBuoy/waba/bookings/interfaces"
	"github.com/SpoonBuoy/waba/usecases/medical"
)

type BookingService struct {
	Ctrl bookings.Business
}

func NewBookingService(business bookings.BusinessController) *BookingService {
	return &BookingService{
		Ctrl: business,
	}
}
func (bc *BookingService) CreateActor(req medical.CreateActorReq) {
	//we create a doctor
	doctor := medical.Doctor{}
	var bid int
	bc.Ctrl.AddActor(&doctor, bid)
}

func (bc *BookingService) GetActor(id int) bookings.Actor {
	doc := bc.Ctrl.GetActor(id)
	return doc
}

func (bc *BookingService) GetAllActors(bid int) {
	bc.Ctrl.GetAllActors(bid)
}

func (bc *BookingService) GetAllAppointments(bid int) {
	bc.Ctrl.GetAllAppointments(bid)
}
func (bc *BookingService) GetSlots() []bookings.Slot {
	//get the slots of doc
	var docId int
	doc := bc.Ctrl.GetActor(docId)
	slots := doc.GetSlots()
	return slots
}
func (bc *BookingService) CreateAppointment(req medical.CreateAppointmentReq) {
	//we have to create an appointment for some doc
	//lets get the doc first
	var docId int = 0
	doc := bc.Ctrl.GetActor(docId)
	//now lets create a slot from dto
	slot := medical.DocSlot{}
	doc.BookSlot(&slot)
}

func (bc *BookingService) CreateBusinessService(req medical.CreateBusinessReq) {

}

func (bc *BookingService) GetAllBusinessServices() {

}
