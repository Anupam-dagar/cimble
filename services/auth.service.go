package services

import (
	"fmt"
)

type AuthServiceInterface interface {
	Login()
	SignUp()
	Register()
}

type AuthService struct {
}

func NewAuthService() AuthServiceInterface {
	as := new(AuthService)
	return as
}

func (as *AuthService) Login() {
	fmt.Println("Login Service")
}

func (as *AuthService) SignUp() {
	fmt.Println("SignUp Service")
}

func (as *AuthService) Register() {
	fmt.Println("Register Service")
}
