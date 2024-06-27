package controllers

import (
	"net/http"

	"github.com/SpoonBuoy/waba/usecases/bookings/interfaces"
	"github.com/ahsmha/gashtools/logger"
	"github.com/gin-gonic/gin"
)

type bookingsController struct {
	logger   logger.ILogWriter
	bservice interfaces.IBookingsService
}

func newBookingsController(logger logger.ILogWriter, service interfaces.IBookingsService) *bookingsController {
	return &bookingsController{
		logger:   logger,
		bservice: service,
	}
}

func BookingRoutes(bRouter *gin.RouterGroup) {
	r := BookingsController
	bRouter.GET("/hi", BookingsController.GetBookings)
	bRouter.GET("/hii", r.GetBookings)
}

func (bc *bookingsController) GetBookings(ctx *gin.Context) {
	bc.bservice.GetBookings(ctx)
	ctx.JSON(http.StatusOK, gin.H{"valid": true})
}
