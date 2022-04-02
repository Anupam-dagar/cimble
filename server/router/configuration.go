package router

import (
	"cimble/controllers"
	"cimble/middlewares"

	"github.com/gin-gonic/gin"
)

func (engine Router) ConfigurationRoute(routeGroup *gin.RouterGroup) {
	configurationRoute := routeGroup.Group("/configuration")
	{
		oc := controllers.NewConfigurationController()

		configurationRoute.POST("/:projectId", middlewares.AuthoriseRequest(), oc.CreateConfiguration)
		configurationRoute.PUT("/:projectId/:id", middlewares.AuthoriseRequest(), oc.UpdateConfiguration)
		configurationRoute.GET("/:projectId", middlewares.AuthoriseRequest(), oc.GetConfigurations)
		configurationRoute.GET("/:projectId/:organisationId/json", middlewares.AuthoriseRequest(), oc.GetFormattedConfigurations)
		configurationRoute.DELETE("/:projectId/:id", middlewares.AuthoriseRequest(), oc.DeleteConfiguration)
	}
}
