package middleware

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

type RequestContext struct {
	BusinessId uint
}

func GetRequestKeys(ctx *gin.Context) RequestContext {
	bid, _ := strconv.Atoi(fmt.Sprint(ctx.Keys["businessID"]))
	return RequestContext{
		BusinessId: uint(bid),
	}
}
