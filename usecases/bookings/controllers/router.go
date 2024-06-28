package controllers

import (
	"sync"

	"github.com/SpoonBuoy/waba/db"
	"github.com/SpoonBuoy/waba/usecases/bookings/interfaces"
	"github.com/SpoonBuoy/waba/usecases/bookings/services"
	clinicRepo "github.com/SpoonBuoy/waba/usecases/clinics/repository"
	"github.com/ahsmha/gashtools/logger"
)

var BookingsController interfaces.IBookingsController
var _bookingsController sync.Once

// this does the injection
func Init(logger logger.ILogWriter, db *db.Db) {
	_bookingsController.Do(func() {
		clinicRepo.InitClinicRepo()
		services.Init(logger, db, clinicRepo.ClinicRepo)
		// BookingsController = newBookingsController(logger, services.BookingsService)
		ClinicsController = newBookingsController(logger, services.BookingsService, clinicRepo.ClinicRepo)
	})
}
