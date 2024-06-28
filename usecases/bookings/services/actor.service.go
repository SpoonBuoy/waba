package services

import "github.com/gin-gonic/gin"

func (bs *bookingservice) GetActors(ctx *gin.Context) {
	bs.crepo.GetDoctors(ctx)
}
func (bs *bookingservice) AddActor(ctx *gin.Context) {
	bs.crepo.AddDoctor(ctx)
}
func (bs *bookingservice) UpdateActor(ctx *gin.Context) {
	bs.crepo.UpdateDoctor(ctx)
}
func (bs *bookingservice) DisableActor(ctx *gin.Context) {
	bs.crepo.DisableDoctor(ctx)
}
