package middlewares

import (
	"cimble/services"
	"cimble/utilities"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func AuthoriseJwt() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader("Authorization")
		tokenString := authHeader[len("Bearer "):]

		token, err := services.NewAuthService().ValidateToken(tokenString)
		if token.Valid {
			claims := token.Claims.(jwt.MapClaims)
			ctx.Set("id", claims["id"].(string))
			ctx.Set("email", claims["email"].(string))
			return
		}

		fmt.Printf(`Invalid jwt token: %v\n`, err)
		utilities.ResponseWithError(ctx, http.StatusUnauthorized, err)
	}
}
