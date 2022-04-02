package router

import (
	"cimble/controllers"
	"cimble/middlewares"

	"github.com/gin-gonic/gin"
)

func (engine Router) AuthRoute(routeGroup *gin.RouterGroup) {
	authRoute := routeGroup.Group("/auth")
	{
		ac := controllers.NewAuthController()

		authRoute.POST("/login", ac.Login)
		authRoute.POST("/signup", ac.SignUp)
		authRoute.POST("/refreshToken", middlewares.AuthoriseRequest(), ac.RefreshToken)
		authRoute.POST("/register", ac.Register)
	}
}
