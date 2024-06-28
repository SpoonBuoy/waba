package interfaces

import "github.com/gin-gonic/gin"

type IActorController interface {
	GetActors(ctx *gin.Context)
	AddActor(ctx *gin.Context)
	UpdateActor(ctx *gin.Context)
	DisableActor(ctx *gin.Context)
}

type IActorService interface {
	GetActors(ctx *gin.Context)
	AddActor(ctx *gin.Context)
	UpdateActor(ctx *gin.Context)
	DisableActor(ctx *gin.Context)
}
