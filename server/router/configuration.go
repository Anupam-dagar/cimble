package router

import (
	"cimble/controllers"
	"cimble/middlewares"

	"github.com/gin-gonic/gin"
)

func (engine Router) ConfigurationRoute(routeGroup *gin.RouterGroup) {
	configurationRoute := routeGroup.Group("/configuration", middlewares.AuthoriseJwt())
	{
		oc := controllers.NewConfigurationController()

		configurationRoute.POST("/:projectId", oc.CreateConfiguration)
	}
}
