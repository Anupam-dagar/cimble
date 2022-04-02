package router

import (
	"cimble/middlewares"

	"github.com/gin-gonic/gin"
)

type Router struct {
	Router *gin.Engine
}

func SetupRoutes() Router {
	router := Router{
		Router: gin.Default(),
	}

	router.Router.Use(middlewares.CORS())
	apiRouter := router.Router.Group("/api")

	router.AuthRoute(apiRouter)
	router.OrganisationRoute(apiRouter)
	router.ProjectRoute(apiRouter)
	router.ConfigurationRoute(apiRouter)
	router.ApiSecretsRoute(apiRouter)

	return router
}
