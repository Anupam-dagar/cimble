package controllers

import (
	"cimble/services"
	"fmt"

	"github.com/gin-gonic/gin"
)

type AuthControllerInterface interface {
	Login(*gin.Context)
	SignUp(*gin.Context)
	Register(*gin.Context)
}

type AuthController struct {
	AuthService services.AuthServiceInterface
}

func NewAuthController() AuthControllerInterface {
	ac := new(AuthController)
	ac.AuthService = services.NewAuthService()
	return ac
}

func (ac *AuthController) Login(ctx *gin.Context) {
	ac.AuthService.Login()
	fmt.Println("Login Controller")
}

func (ac *AuthController) SignUp(ctx *gin.Context) {
	ac.AuthService.SignUp()
	fmt.Println("SignUp Controller")
}

func (ac *AuthController) Register(ctx *gin.Context) {
	ac.AuthService.Register()
	fmt.Println("Register Controller")
}
