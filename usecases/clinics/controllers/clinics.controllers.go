package controllers

import (
	"net/http"

	"github.com/SpoonBuoy/waba/usecases/clinics/interfaces"
	"github.com/ahsmha/gashtools/logger"
	"github.com/gin-gonic/gin"
)

type clinicController struct {
	logger   logger.ILogWriter
	cservice interfaces.IClinicService
}

func NewClinicController(logger logger.ILogWriter) *clinicController {
	return &clinicController{
		logger: logger,
	}
}

func (cc *clinicController) GetBookings(ctx *gin.Context) {
	cc.cservice.GetBookings(ctx)
	ctx.JSON(http.StatusOK, gin.H{"valid": true})
}
