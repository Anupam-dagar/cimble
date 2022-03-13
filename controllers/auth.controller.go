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
	RefreshToken(*gin.Context)
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
	var loginPayload models.Login
	err := ctx.ShouldBindJSON(&loginPayload)
	if err != nil {
		fmt.Printf("Error binding json: %v\n", err)
		utilities.ResponseWithErrorCode(ctx, http.StatusBadRequest, err)
		return
	}

	loginResponse, err := ac.AuthService.Login(loginPayload)
	if err != nil {
		fmt.Printf("Error login user: %v\n", err)
		utilities.ResponseWithErrorCode(ctx, http.StatusInternalServerError, err)
		return
	}

	utilities.ResponseWithSuccess(ctx, http.StatusOK, loginResponse)
}

func (ac *AuthController) SignUp(ctx *gin.Context) {
	var signUpPayload models.SignUp
	err := ctx.ShouldBindJSON(&signUpPayload)

	if err != nil {
		fmt.Printf("Error binding json: %v\n", err)
		utilities.ResponseWithErrorCode(ctx, http.StatusBadRequest, err)
		return
	}

	err = ac.AuthService.SignUp(signUpPayload)
	if err != nil {
		fmt.Printf("Error signing up user: %v\n", err)
		utilities.ResponseWithErrorCode(ctx, http.StatusInternalServerError, err)
		return
	}

	utilities.ResponseWithSuccess(ctx, http.StatusOK, signUpPayload)
}

func (ac *AuthController) RefreshToken(ctx *gin.Context) {
	userId, ok := ctx.Get("id")
	if !ok {
		err := fmt.Errorf("invalid user id")
		utilities.ResponseWithErrorCode(ctx, http.StatusBadRequest, err)
		return
	}

	refreshTokenResponse, err := ac.AuthService.RefreshToken(fmt.Sprintf("%v", userId))
	if err != nil {
		fmt.Printf("Error generating token pair: %v\n", err)
		utilities.ResponseWithErrorCode(ctx, http.StatusInternalServerError, err)
		return
	}

	utilities.ResponseWithSuccess(ctx, http.StatusOK, refreshTokenResponse)
}

func (ac *AuthController) Register(ctx *gin.Context) {
	ac.AuthService.Register()
	fmt.Println("Register Controller")
}
