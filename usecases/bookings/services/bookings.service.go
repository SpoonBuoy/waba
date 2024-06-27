package services

import (
	"context"

	"github.com/SpoonBuoy/waba/usecases/bookings/interfaces"
	"github.com/ahsmha/gashtools/logger"
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

func (cs *bookingservice) GetBookings(ctx context.Context) {
	cs.crepo.GetBookings(ctx)
	return
}
