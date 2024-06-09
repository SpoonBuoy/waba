package controllers

import (
	"log"
	"net/http"

	"github.com/SpoonBuoy/waba/dto"
	"github.com/SpoonBuoy/waba/service"
	"github.com/gin-gonic/gin"
)

type BusinessController struct {
	businessService service.BusinessService
}

func NewBusinessController(businessServ service.BusinessService) *BusinessController {
	return &BusinessController{
		businessService: businessServ,
	}
}

func HandleBindErr(err error, where string, ctx *gin.Context) {
	if err != nil {
		log.Printf("failed to unmarshall at %s with err : %v", where, err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		ctx.Abort()
	}
}

func (bc *BusinessController) AddBusiness(ctx *gin.Context) {
	var req dto.CreateBusinessReq
	err := ctx.BindJSON(&req)
	HandleBindErr(err, "AddBusiness", ctx)

	b := bc.businessService.CreateBusiness(req)
	ctx.JSON(http.StatusOK, gin.H{"data": b})
}

func (bc *BusinessController) AddContext(ctx *gin.Context) {
	var req dto.CreateCtxReq
	err := ctx.BindJSON(&req)
	HandleBindErr(err, "AddContext", ctx)

	c := bc.businessService.CreateContext(req)
	ctx.JSON(http.StatusOK, gin.H{"data": c})
}

func (bc *BusinessController) GetContexts(ctx *gin.Context) {
	//get business id from token
	var bid uint
	cs := bc.businessService.GetAllContexts(bid)
	ctx.JSON(http.StatusOK, gin.H{"data": cs})
}

func (bc *BusinessController) SetActiveCtx(ctx *gin.Context) {
	var req dto.SwitchActiveCtxReq
	err := ctx.BindJSON(&req)
	HandleBindErr(err, "SetActiveCtx", ctx)

	res := bc.businessService.SetActiveContext(req)
	ctx.JSON(http.StatusOK, gin.H{"data": res})
}
