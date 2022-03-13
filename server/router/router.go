package router

import "github.com/gin-gonic/gin"

type Router struct {
	Router *gin.Engine
}

func SetupRoutes() Router {
	router := Router{
		Router: gin.Default(),
	}

	apiRouter := router.Router.Group("/api")

	router.AuthRoute(apiRouter)
	router.OrganisationRoute(apiRouter)
	router.ProjectRoute(apiRouter)
	router.ConfigurationRoute(apiRouter)

	return router
}
