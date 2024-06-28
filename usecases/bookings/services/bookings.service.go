package services

import (
	"github.com/SpoonBuoy/waba/usecases/bookings/interfaces"
	"github.com/ahsmha/gashtools/logger"
	"github.com/gin-gonic/gin"
)

type bookingservice struct {
	logger logger.ILogWriter
	repo   interfaces.IBookingsRepository
}

func NewBookingsService(logger logger.ILogWriter, repo interfaces.IBookingsRepository) *bookingservice {
	return &bookingservice{
		logger: logger,
		repo:   repo,
	}
}

// booker interface
func (bs *bookingservice) GetBookings(ctx *gin.Context) {
	bs.repo.GetBookings(ctx)
	return
}

func (bs *bookingservice) Book(ctx *gin.Context) {
	bs.repo.Book(ctx)
	return
}
