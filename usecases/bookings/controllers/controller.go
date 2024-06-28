package controllers

import (
	"sync"

	"github.com/SpoonBuoy/waba/db"
	"github.com/SpoonBuoy/waba/usecases/bookings/interfaces"
	"github.com/ahsmha/gashtools/logger"
)

var BookingsController interfaces.IBookingsController
var _bookingsController sync.Once

func Init(logger logger.ILogWriter, db *db.Db, svc interfaces.IBookingsService) {
	_bookingsController.Do(func() {
		BookingsController = NewBookingsController(logger, svc)
	})
}
