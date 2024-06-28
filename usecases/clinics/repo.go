package clinics

import (
	"sync"

	"github.com/SpoonBuoy/waba/db"
	"github.com/SpoonBuoy/waba/usecases/bookings/interfaces"
	"github.com/ahsmha/gashtools/logger"
)

var (
	ClinicRepo      interfaces.IBookingsRepository
	_clinicRepoOnce sync.Once
)

type clinicRepo struct {
	db     *db.Db
	logger logger.ILogWriter
}

func NewClinicRepo(db *db.Db, logger logger.ILogWriter) interfaces.IBookingsRepository {
	return clinicRepo{
		db:     db,
		logger: logger,
	}
}

func InitClinicRepo(logger logger.ILogWriter, db *db.Db) {
	_clinicRepoOnce.Do(func() {
		ClinicRepo = NewClinicRepo(db, logger)
	})
}
