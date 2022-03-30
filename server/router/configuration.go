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
		configurationRoute.PUT("/:projectId/:id", oc.UpdateConfiguration)
		configurationRoute.GET("/:projectId", oc.GetConfigurations)
		configurationRoute.GET("/:projectId/json", oc.GetFormattedConfigurations)
		configurationRoute.DELETE("/:projectId/:id", oc.DeleteConfiguration)
	}
}
