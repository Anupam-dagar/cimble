package router

import (
	"cimble/controllers"
	"cimble/middlewares"

	"github.com/gin-gonic/gin"
)

func (engine Router) OrganisationRoute(routeGroup *gin.RouterGroup) {
	organisationRoute := routeGroup.Group("/organisation", middlewares.AuthoriseJwt())
	{
		oc := controllers.NewOrganisationController()

		organisationRoute.POST("/", oc.CreateOrganisation)
		organisationRoute.PUT("/:id", oc.UpdateOrganisation)
		organisationRoute.GET("/", oc.GetOrganisations)
		organisationRoute.DELETE("/:id", oc.DeleteOrganisation)
	}
}
