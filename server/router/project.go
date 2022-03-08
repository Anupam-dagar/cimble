package router

import (
	"cimble/controllers"
	"cimble/middlewares"

	"github.com/gin-gonic/gin"
)

func (engine Router) ProjectRoute(routeGroup *gin.RouterGroup) {
	projectRoute := routeGroup.Group("/project", middlewares.AuthoriseJwt())
	{
		pc := controllers.NewProjectController()

		projectRoute.POST("/", pc.CreateProject)
	}
}
