package controller

import (
	"github.com/SpoonBuoy/waba/usecases/common/dto"
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

func HandelBindErr(err error) {

}

func (bc *BookingController) CreateActor(ctx *gin.Context) {
	var req dto.CreateActorReq
	err := ctx.BindJSON(&req)
	HandelBindErr(err)
	bc.Svc.CreateActor(req)
	ctx.JSON(Ok())
}

func (bc *BookingController) GetActor(ctx *gin.Context) {

	var actorId int
	//return actor
	actor := bc.Svc.GetActor(actorId)
	ctx.JSON(Ok(KV("actor", actor)))
}

func (bc *BookingController) GetAllActors(ctx *gin.Context) {
	//business id need to be get from auth
	var bid int
	actors := bc.Svc.GetAllActors(bid)
	ctx.JSON(Ok(KV("actors", actors)))
}

func (bc *BookingController) GetAllAppointments(*gin.Context) {

}

func (bc *BookingController) CreateAppointment(*gin.Context) {

}

func (bc *BookingController) BookSlot(ctx *gin.Context) {
	var req dto.BookSlotReq
	err := ctx.BindJSON(&req)
	HandelBindErr(err)
	bc.Svc.BookSlot(req)
	ctx.JSON(Ok())
}

func (bc *BookingController) CreateBusinessService(ctx *gin.Context) {
	var req dto.CreateServiceReq
	err := ctx.BindJSON(&req)
	HandelBindErr(err)
	bc.Svc.CreateBusinessService(req)
	ctx.JSON(Ok())
}

func (bc *BookingController) GetAllBusinessServices(*gin.Context) {

}
