package services

import (
	"sync"

	"github.com/SpoonBuoy/waba/db"
	"github.com/SpoonBuoy/waba/usecases/clinics/repository"
	"github.com/ahsmha/gashtools/logger"
)

var BookingsService *bookingservice
var _bookingsService sync.Once

func Init(logger logger.ILogWriter, db *db.Db) {
	_bookingsService.Do(func() {
		repository.InitClinicRepo(logger, db)
		BookingsService = newBookingsService(logger, repository.ClinicRepo)
	})
}
