package controllers

import (
	"log"
	"net/http"
	"time"

	"github.com/SpoonBuoy/waba/dto"
	"github.com/SpoonBuoy/waba/service"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type BusinessController struct {
	Secret          []byte
	businessService service.BusinessService
}

func NewBusinessController(businessServ service.BusinessService, secret string) *BusinessController {
	return &BusinessController{
		businessService: businessServ,
		Secret:          []byte(secret),
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
	bid := ctx.GetInt("businessId")
	c := bc.businessService.CreateContext(req, uint(bid))
	ctx.JSON(http.StatusOK, gin.H{"data": c})
}

func (bc *BusinessController) GetContexts(ctx *gin.Context) {
	//get business id from token
	var bid uint = 1
	cs := bc.businessService.GetAllContexts(bid)
	ctx.JSON(http.StatusOK, gin.H{"data": cs})
}

func (bc *BusinessController) SetActiveCtx(ctx *gin.Context) {
	var req dto.SwitchActiveCtxReq
	err := ctx.BindJSON(&req)
	HandleBindErr(err, "SetActiveCtx", ctx)
	bid := ctx.GetInt("businessId")
	res := bc.businessService.SetActiveContext(req, uint(bid))
	ctx.JSON(http.StatusOK, gin.H{"data": res})
}

func (bc *BusinessController) Login(ctx *gin.Context) {
	//first validate credentials
	var req dto.LoginReq
	err := ctx.BindJSON(&req)
	HandleBindErr(err, "Login", ctx)
	business, err := bc.businessService.ValidateBusiness(req)

	if err != nil {
		//invalid credentials
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid login credentials"})
	}

	//valid credentials
	token, err := bc.CreateToken(business.ID, business.Name)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "failed to create token"})
	}

	ctx.JSON(http.StatusOK, gin.H{"token": token, "business": business})
}

func (bc *BusinessController) Logout(ctx *gin.Context) {

}

var exp = time.Now().Add(time.Hour * 24).Unix()
var secret = []byte("some_key")

func (bc *BusinessController) CreateToken(bId uint, name string) (string, error) {
	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		jwt.MapClaims{
			"sub": bId,
			"exp": exp,
		},
	)
	tokenStr, err := token.SignedString(bc.Secret)

	if err != nil {
		return "", err
	}
	return tokenStr, nil
}
