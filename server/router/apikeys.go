package router

import (
	"cimble/controllers"
	"cimble/middlewares"

	"github.com/gin-gonic/gin"
)

func (engine Router) ApiSecretsRoute(routeGroup *gin.RouterGroup) {
	apiKeyRoute := routeGroup.Group("/api-keys", middlewares.AuthoriseJwt())
	{
		akc := controllers.NewApiKeysController()

		apiKeyRoute.POST("/", akc.CreateApiKey)
	}
}
