package services

import (
	"cimble/models"
	"cimble/repositories"
	"cimble/utilities"
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
	user := signUpPayload.CreateUserEntity("test")
	passwordHash := utilities.GenerateSha512Hash(signUpPayload.Password)
	userPassword := models.UserPassword{
		UserId:       user.ID,
		PasswordHash: passwordHash,
		CreatedBy:    "test",
		UpdatedBy:    "test",
	}

	err = as.UserRepository.AddUser(user, userPassword)

	return err
}

func (as *AuthService) Register() {
	fmt.Println("Register Service")
}
