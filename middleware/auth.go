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

func (au *Auth) GetClaimsMap(ctx *gin.Context) jwt.MapClaims {
	tokenStr := ctx.Request.Header.Get("Authorization")
	tokenStr = tokenStr[len("Bearer "):] // remove "Bearer " prefix
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return au.Secret, nil
	})
	if err != nil {
		//not able to verify
		ctx.JSON(401, gin.H{"error": "failed to parse token: " + err.Error()})
		return nil
	}
	if !token.Valid {
		//token not valid
		ctx.JSON(401, gin.H{"error": "invalid token"})
		return nil
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		//failed to get claims
		ctx.JSON(401, gin.H{"error": "no claims found in token"})
		return nil
	}
	return claims
}
func (au *Auth) Verify(ctx *gin.Context) {
	claims := au.GetClaimsMap(ctx)
	if claims == nil {
		ctx.JSON(401, gin.H{"error": "claims not found in token"})
	}
	busId, exists := claims["businessId"].(int)
	if !exists {
		ctx.JSON(401, gin.H{"error": "busId not found in token"})
	}

	ctx.Set("businessID", busId)
	ctx.Next()
}

func (au *Auth) VerifyRole(ctx *gin.Context) {
	claims := au.GetClaimsMap(ctx)
	role, exists := claims["role"].(string)
	if !exists {
		ctx.JSON(401, gin.H{"error": "role not found in token"})
	}
	//bind role with context
	ctx.Set("role", role)
	ctx.Next()

}
