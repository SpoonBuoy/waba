package services

import (
	"github.com/SpoonBuoy/waba/usecases/bookings/interfaces"
	"github.com/ahsmha/gashtools/logger"
	"github.com/gin-gonic/gin"
)

type bookingservice struct {
	logger logger.ILogWriter
	crepo  interfaces.IClinicsRepository
}

func newBookingsService(logger logger.ILogWriter, clinicRepo interfaces.IClinicsRepository) *bookingservice {
	return &bookingservice{
		logger: logger,
		crepo:  clinicRepo,
	}
}

// booker interface
func (bs *bookingservice) GetBookings(ctx *gin.Context) {
	bs.crepo.GetBookings(ctx)
	return
}

func (bs *bookingservice) Book(ctx *gin.Context) {
	bs.crepo.Book(ctx)
	return
}
