package clinics

import (
	"github.com/SpoonBuoy/waba/db"
	bookingController "github.com/SpoonBuoy/waba/usecases/bookings/controllers"
	"github.com/SpoonBuoy/waba/usecases/bookings/interfaces"
	"github.com/SpoonBuoy/waba/usecases/bookings/services"
	clinicRepo "github.com/SpoonBuoy/waba/usecases/clinics/repository"
	"github.com/ahsmha/gashtools/logger"
	"github.com/gin-gonic/gin"
)

var ClinicsController interfaces.IBookingsController

// this does the injection
func Init(logger logger.ILogWriter, db *db.Db) {
	clinicRepo := clinicRepo.NewClinicRepo(db, logger)
	clinicService := services.NewBookingsService(logger, clinicRepo)
	// BookingsController = newBookingsController(logger, services.BookingsService)
	ClinicsController = bookingController.NewBookingsController(logger, clinicService)

}

func MedicalRoutes(r *gin.RouterGroup) {
	med := r.Group("med")
	{
		med.GET("/appointments", ClinicsController.GetBookings)
	}
}
