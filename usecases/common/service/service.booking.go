package service

import (
	bookings "github.com/SpoonBuoy/waba/bookings/interfaces"
	"github.com/SpoonBuoy/waba/usecases/common/dto"
	"github.com/SpoonBuoy/waba/usecases/factory"
)

type BookingService struct {
	Repo bookings.BusinessRepo
}

func NewBookingService(business bookings.BusinessRepo) *BookingService {
	return &BookingService{
		Repo: business,
	}
}
func (bc *BookingService) CreateActor(req dto.CreateActorReq) {
	//we create an actor
	actor := factory.NewActor(req.Name, req.Details, req.Service, req.BusinessType)
	var bid int
	bc.Repo.AddActor(actor, bid)
}

func (bc *BookingService) GetActor(id int) bookings.Actor {
	doc := bc.Repo.GetActor(id)
	return doc
}

func (bc *BookingService) GetAllActors(bid int) []bookings.Actor {
	return bc.Repo.GetAllActors(bid)
}

func (bc *BookingService) GetAllAppointments(bid int) []bookings.Appointment {
	return bc.Repo.GetAllAppointments(bid)
}
func (bc *BookingService) GetSlots() []bookings.Slot {
	//get the slots of doc
	// var docId int
	// doc := bc.Repo.GetActor(docId)
	// slots := doc.GetSlots("", "")
	// return slots
	return nil
}
func (bc *BookingService) CreateAppointment(req dto.CreateAppointmentReq) {
	//we have to create an appointment for some doc
	//lets get the doc first
	// var docId int = 0
	// doc := bc.Repo.GetActor(docId)
	// //now lets create a slot from dto
	// slot := medical.DocSlot{}
}
func (bc *BookingService) BookSlot(req dto.BookSlotReq) {
	bc.Repo.BookSlot(req.SlotId)

}
func (bc *BookingService) CreateBusinessService(req dto.CreateServiceReq) {
	// have to be auth
	var bid int
	service := factory.NewService(req.Name, req.BusinessType)
	bc.Repo.AddService(service, bid)
}

func (bc *BookingService) GetAllBusinessServices() []bookings.ActorService {
	var bid int
	return bc.Repo.GetAllServices(bid)
}
