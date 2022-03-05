package controllers

import (
	"cimble/models"
	"cimble/services"
	"cimble/utilities"
	"fmt"
	"net/http"

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
	var signUpPayload models.SignUp
	err := ctx.ShouldBindJSON(&signUpPayload)

	if err != nil {
		fmt.Printf("Error binding json: %v\n", err)
		utilities.ResponseWithError(ctx, http.StatusInternalServerError, err)
		return
	}

	err = ac.AuthService.SignUp(signUpPayload)
	if err != nil {
		fmt.Printf("Error signing up user: %v\n", err)
		utilities.ResponseWithError(ctx, http.StatusInternalServerError, err)
	}

	utilities.ResponseWithSuccess(ctx, http.StatusOK, signUpPayload)
}

func (ac *AuthController) Register(ctx *gin.Context) {
	ac.AuthService.Register()
	fmt.Println("Register Controller")
}
