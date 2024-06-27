package controllers

import (
	"sync"

	"github.com/SpoonBuoy/waba/db"
	"github.com/SpoonBuoy/waba/usecases/bookings/services"
	"github.com/ahsmha/gashtools/logger"
)

var BookingsController bookingsController
var _bookingsController sync.Once

// this does the injection
func Init(logger logger.ILogWriter, db *db.Db) {
	_bookingsController.Do(func() {
		services.Init(logger, db)
		BookingsController = *newBookingsController(logger, services.BookingsService)
	})
}
