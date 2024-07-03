package router

import (
	"github.com/SpoonBuoy/waba/usecases/common/controller"
	"github.com/SpoonBuoy/waba/usecases/common/service"
	"github.com/SpoonBuoy/waba/usecases/medical"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var bc *controller.BookingController

func InitMed() {
	medBus := medical.NewMedicalBusiness("gasha clinic", &gorm.DB{})
	ser := service.NewBookingService(medBus)
	bc = controller.NewBookingController(ser)
}

func MedicalRouter(med *gin.RouterGroup) {
	med.POST("/doctor", bc.CreateActor)
	med.POST("/service", bc.CreateBusinessService)
	med.GET("/actor", bc.GetActor)
	med.GET("/actors", bc.GetAllActors)
	med.GET("/appointments", bc.GetAllAppointments)
	med.GET("/services", bc.GetAllBusinessServices)
}
