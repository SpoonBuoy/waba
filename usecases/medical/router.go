package medical

import (
	"github.com/SpoonBuoy/waba/usecases/common/controller"
	"github.com/SpoonBuoy/waba/usecases/common/service"
	"github.com/gin-gonic/gin"
)

var bc *controller.BookingController

func InitMed() {
	medBus := NewMedicalBusiness()
	ser := service.NewBookingService()
	bc = controller.NewBookingController(ser)
}

func MedicalRouter(med *gin.RouterGroup) {
	med.POST("addactor", bc.CreateActor)
	med.POST("addservice", bc.CreateBusinessService)
	med.GET("actor", bc.GetActor)
	med.GET("actors", bc.GetAllActors)
	med.GET("appointments", bc.GetAllAppointments)
	med.GET("services", bc.GetAllBusinessServices)
}
