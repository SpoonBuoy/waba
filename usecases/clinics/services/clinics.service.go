package services

import (
	"context"

	"github.com/SpoonBuoy/waba/usecases/clinics/interfaces"
	"github.com/ahsmha/gashtools/logger"
)

type clinicService struct {
	logger logger.ILogWriter
	crepo  interfaces.IClinicRepository
}

func NewClinicService(logger logger.ILogWriter, clinicRepo interfaces.IClinicRepository) *clinicService {
	return &clinicService{
		logger: logger,
		crepo:  clinicRepo,
	}
}

func (cs *clinicService) GetBookings(ctx context.Context) {
	cs.crepo.GetBookings(ctx)
	return
}
