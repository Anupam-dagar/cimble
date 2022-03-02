package router

import (
	"cimble/controllers"

	"github.com/gin-gonic/gin"
)

func (engine Router) AuthRoute(routeGroup *gin.RouterGroup) {
	authRoute := routeGroup.Group("/auth")
	{
		ac := controllers.NewAuthController()

		authRoute.POST("/login", ac.Login)
		authRoute.POST("/signup", ac.SignUp)
		authRoute.POST("/register", ac.Register)
	}
}
