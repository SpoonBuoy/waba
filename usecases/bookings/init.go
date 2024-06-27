package bookings

import (
	"github.com/SpoonBuoy/waba/db"
	"github.com/SpoonBuoy/waba/usecases/bookings/controllers"
	"github.com/ahsmha/gashtools/logger"
)

// maybe some other inits other than di
// configs etc
func Init(logger logger.ILogWriter, db *db.Db) {
	controllers.Init(logger, db)
}
