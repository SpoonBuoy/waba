package controller

import (
	"github.com/SpoonBuoy/waba/usecases/common/service"
	"github.com/SpoonBuoy/waba/usecases/medical"
	"github.com/gin-gonic/gin"
)

var bc *BookingController

func InitMed() {
	medBus := medical.NewMedicalBusiness()
	ser := service.NewBookingService(medBus)
	bc = NewBookingController(ser)
}

func MedicalRouter(med *gin.RouterGroup) {
	med.POST("addactor", bc.CreateActor)
	med.POST("addservice", bc.CreateBusinessService)
	med.GET("actor", bc.GetActor)
	med.GET("actors", bc.GetAllActors)
	med.GET("appointments", bc.GetAllAppointments)
	med.GET("services", bc.GetAllBusinessServices)
}
