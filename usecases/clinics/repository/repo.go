package repository

import (
	"github.com/SpoonBuoy/waba/db"
	"github.com/ahsmha/gashtools/logger"
)

type clinicRepo struct {
	db     db.Db
	logger logger.ILogWriter
}

func newClinicRepo(db db.Db, logger logger.ILogWriter) *clinicRepo {
	return &clinicRepo{
		db:     db,
		logger: logger,
	}
}
