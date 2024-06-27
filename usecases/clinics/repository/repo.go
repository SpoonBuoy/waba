package repository

import (
	"sync"

	"github.com/SpoonBuoy/waba/db"
	"github.com/ahsmha/gashtools/logger"
)

var (
	ClinicRepo      *clinicRepo
	_clinicRepoOnce sync.Once
)

type clinicRepo struct {
	db     *db.Db
	logger logger.ILogWriter
}

func newClinicRepo(db *db.Db, logger logger.ILogWriter) *clinicRepo {
	return &clinicRepo{
		db:     db,
		logger: logger,
	}
}

func InitClinicRepo(logger logger.ILogWriter, db *db.Db) {
	_clinicRepoOnce.Do(func() {
		ClinicRepo = newClinicRepo(db, logger)
	})
}
