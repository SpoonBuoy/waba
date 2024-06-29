package service

import (
	bookings "github.com/SpoonBuoy/waba/bookings/interfaces"
	"github.com/SpoonBuoy/waba/usecases/medical"
)

type BookingService struct {
	Ctrl bookings.Business
}

func NewBookingService() *BookingService {
	return &BookingService{}
}
func (bc *BookingService) CreateActor(req medical.CreateActorReq) {
	//we create a doctor
	doctor := medical.Doctor{}
	bc.Ctrl.AddActor(&doctor)
}

func (bc *BookingService) GetActor(id int) {
	bc.Ctrl.GetActor(id)
}

func (bc *BookingService) GetAllActors() {
	bc.Ctrl.GetAllActors()
}

func (bc *BookingService) GetAllAppointments() {
	bc.Ctrl.GetAllAppointments()
}
func (bc *BookingService) GetSlots() {
	//get the slots of doc
	var docId int
	doc := bc.Ctrl.GetActor(docId)
	slots := doc.GetSlots()
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
