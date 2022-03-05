package services

import (
	"cimble/models"
	"cimble/repositories"
	"fmt"
)

type AuthServiceInterface interface {
	Login()
	SignUp(models.SignUp) error
	Register()
}

type AuthService struct {
	UserRepository repositories.UserRepositoryInterface
}

func NewAuthService() AuthServiceInterface {
	as := new(AuthService)
	as.UserRepository = repositories.NewUserRepository()
	return as
}

func (as *AuthService) Login() {
	fmt.Println("Login Service")
}

func (as *AuthService) SignUp(signUpPayload models.SignUp) (err error) {
	user := signUpPayload.ToCreateUserEntity("test")
	err = as.UserRepository.AddUser(user)

	return
}

func (as *AuthService) Register() {
	fmt.Println("Register Service")
}
