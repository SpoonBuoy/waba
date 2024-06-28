package controllers

import "github.com/gin-gonic/gin"

func (bc *bookingsController) GetActors(ctx *gin.Context) {
	bc.bservice.GetActors(ctx)
}
func (bc *bookingsController) AddActor(ctx *gin.Context) {
	bc.bservice.AddActor(ctx)
}
func (bc *bookingsController) UpdateActor(ctx *gin.Context) {
	bc.bservice.UpdateActor(ctx)
}
func (bc *bookingsController) DisableActor(ctx *gin.Context) {
	bc.bservice.DisableActor(ctx)
}
