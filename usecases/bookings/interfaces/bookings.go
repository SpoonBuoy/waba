package interfaces

import "github.com/gin-gonic/gin"

type IBookingsController interface {
	IBookerController
	IActorController
}

type IBookingsService interface {
	IBookerService
	IActorService
}

type IBookingsRepository interface {
	GetBookings(ctx *gin.Context)
	Book(ctx *gin.Context)
	UpdateAvailability(ctx *gin.Context)
	GetActors(ctx *gin.Context)
	AddActor(ctx *gin.Context)
	UpdateActor(ctx *gin.Context)
	DisableActor(ctx *gin.Context)
}
