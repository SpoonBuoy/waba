package controller

import (
	"github.com/SpoonBuoy/waba/usecases/common/service"
	"github.com/gin-gonic/gin"
)

type BookingController struct {
	Svc *service.BookingService
}

func NewBookingController(svc *service.BookingService) *BookingController {
	return &BookingController{
		Svc: svc,
	}
}

func (bc *BookingController) CreateActor(*gin.Context) {

}

func (bc *BookingController) GetActor(ctx *gin.Context) {

	var actorId int
	//return actor
	actor := bc.Svc.Ctrl.GetActor(actorId)
	ctx.JSON(200, gin.H{"actor": actor})
}

func (bc *BookingController) GetAllActors(*gin.Context) {

}

func (bc *BookingController) GetAllAppointments(*gin.Context) {

}

func (bc *BookingController) CreateAppointment(*gin.Context) {

}

func (bc *BookingController) CreateBusinessService(*gin.Context) {

}

func (bc *BookingController) GetAllBusinessServices(*gin.Context) {

}
