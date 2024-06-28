package services

import (
	"sync"

	"github.com/SpoonBuoy/waba/db"
	"github.com/SpoonBuoy/waba/usecases/bookings/interfaces"
	"github.com/ahsmha/gashtools/logger"
)

var BookingsService interfaces.IBookingsService
var _bookingsService sync.Once

func Init(logger logger.ILogWriter, db *db.Db, repo interfaces.IBookingsRepository) {
	_bookingsService.Do(func() {
		BookingsService = NewBookingsService(logger, repo)
	})
}
