package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type Auth struct {
	Secret []byte
}

func NewAuthMiddleware(secret string) *Auth {
	return &Auth{Secret: []byte(secret)}
}

func (au *Auth) Verify(ctx *gin.Context) {
	tokenStr := ctx.Request.Header.Get("Authorization")
	tokenStr = tokenStr[len("Bearer "):] // remove "Bearer " prefix
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return au.Secret, nil
	})
	if err != nil {
		//not able to verify
		ctx.JSON(401, gin.H{"error": "failed to parse token: " + err.Error()})
		return
	}
	if !token.Valid {
		//token not valid
		ctx.JSON(401, gin.H{"error": "invalid token"})
	}

	busId, err := token.Claims.GetSubject()

	if err != nil {
		//failed to get business id from token
		ctx.JSON(401, gin.H{"error": "data not found in token"})
	}
	ctx.Set("businessID", busId)
	ctx.Next()
}
