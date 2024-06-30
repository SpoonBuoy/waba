package controller

import (
	"net/http"

	"github.com/SpoonBuoy/waba/usecases/common/service"
	"github.com/SpoonBuoy/waba/usecases/medical"
	"github.com/ahsmha/gashtools/logger"
	"github.com/gin-gonic/gin"
)

type BookingController struct {
	bservice *service.BookingService
	logger   logger.ILogWriter
}

func NewBookingController(svc *service.BookingService, logger logger.ILogWriter) *BookingController {
	return &BookingController{
		bservice: svc,
		logger:   logger,
	}
}

// utils
func sendResponse(ctx *gin.Context, code int, message string, data any) {
	ctx.JSON(code, gin.H{
		"message": message,
		"data":    data,
	})
}

func (bc *BookingController) CreateActor(ctx *gin.Context) {
	var req medical.CreateActorReq
	if err := ctx.BindJSON(req); err != nil {
		// bc.logger.
		sendResponse(ctx, http.StatusUnprocessableEntity, "failed to bind req", nil)
	}
	err := bc.bservice.CreateActor(ctx, req)
	if err != nil {
		// logger
		sendResponse(ctx, http.StatusUnprocessableEntity, "failed creation", nil)
	}
	sendResponse(ctx, http.StatusOK, "created success", nil)
}

func (bc *BookingController) GetActor(ctx *gin.Context) {
	var actorId int
	//return actor
	actor := bc.bservice.GetActor(actorId)
	ctx.JSON(200, gin.H{"actor": actor})
}

func (bc *BookingController) GetAllActors(ctx *gin.Context) {

}

func (bc *BookingController) GetAllAppointments(ctx *gin.Context) {

}

func (bc *BookingController) CreateAppointment(ctx *gin.Context) {

}

func (bc *BookingController) CreateBusinessService(ctx *gin.Context) {

}

func (bc *BookingController) GetAllBusinessServices(ctx *gin.Context) {

}
