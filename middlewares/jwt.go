package middlewares

import (
	"cimble/constants"
	"cimble/services"
	"cimble/utilities"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func AuthoriseRequest() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader("Authorization")
		err := fmt.Errorf("unauthorised access")
		isApiKeyPath := constants.ApiKeyEligiblePaths[ctx.FullPath()]

		// no auth data sent
		if authHeader == "" && !isApiKeyPath {
			fmt.Println("No jwt token")
			utilities.ResponseWithErrorCode(ctx, http.StatusForbidden, err)
			return
		}

		// api key authorisation
		if authHeader == "" && isApiKeyPath {
			apiKey := ctx.Query("apiKey")
			organisationId := ctx.Param("organisationId")
			if apiKey == "" {
				err := fmt.Errorf(string(constants.ApiKeyNotPresent))
				utilities.ResponseWithError(ctx, err)
				return
			}

			apiKeyService := services.NewApiKeyService()
			isValid, err := apiKeyService.IsValidApiKey(organisationId, apiKey)
			if err != nil {
				err = fmt.Errorf("error validating api key")
				utilities.ResponseWithErrorCode(ctx, http.StatusInternalServerError, err)
				return
			}

			if !isValid {
				err := fmt.Errorf("invalid api key")
				utilities.ResponseWithErrorCode(ctx, http.StatusBadRequest, err)
				return
			}

			return
		}

		// jwt authorisation
		tokenString := authHeader[len("Bearer "):]

		token, err := services.NewAuthService().ValidateToken(tokenString)
		if token.Valid {
			claims := token.Claims.(jwt.MapClaims)
			ctx.Set("id", claims["id"].(string))
			ctx.Set("email", claims["email"].(string))
			return
		}

		fmt.Printf("Invalid jwt token: %v\n", err)
		utilities.ResponseWithErrorCode(ctx, http.StatusForbidden, err)
	}
}
