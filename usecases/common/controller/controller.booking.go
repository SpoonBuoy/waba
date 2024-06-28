package controller

import "github.com/gin-gonic/gin"

type BookingController struct{}

func NewBookingController() *BookingController {
	return &BookingController{}
}

func (bc *BookingController) CreateActor(*gin.Context) {

}

func (bc *BookingController) GetActor(*gin.Context) {

}

func (bc *BookingController) GetAllActors(*gin.Context) {

}

func (bc *BookingController) GetAllAppointments(*gin.Context) {

}

func (bc *BookingController) CreateAppointment(*gin.Context) {

}

func (bc *BookingController) CreateBusinessService(*gin.Context) {

}

func (bc *BookingController) GetAllBusinessServices(*gin.Context) {

}
