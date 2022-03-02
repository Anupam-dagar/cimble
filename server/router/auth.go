package router

import "github.com/gin-gonic/gin"

func (engine Router) AuthRoute(routeGroup *gin.RouterGroup) {
	authRoute := routeGroup.Group("/auth")
	{
		authRoute.POST("/login")
		authRoute.POST("/signup")
		authRoute.POST("/register")
	}
}
