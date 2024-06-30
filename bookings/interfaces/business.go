package bookings

import "github.com/gin-gonic/gin"

// manages all the operations for business
type Business interface {
	AddActor(*gin.Context, Actor, int) error
	GetActor(*gin.Context, int) Actor
	GetAllActors(*gin.Context, int) []Actor
	GetAllAppointments(*gin.Context, int) []Appointment
	GetAllServices(*gin.Context, int) []ActorService
	AddService(*gin.Context, ActorService, int)
}
