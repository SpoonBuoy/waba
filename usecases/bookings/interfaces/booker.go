package interfaces

import "github.com/gin-gonic/gin"

type IBookerController interface {
	GetBookings(ctx *gin.Context)
	Book(ctx *gin.Context)
}

type IBookerService interface {
	GetBookings(ctx *gin.Context)
	Book(ctx *gin.Context)
}
