package services

import (
	"cimble/models"
	"cimble/repositories"
	"cimble/utilities"
	"fmt"
)

type AuthServiceInterface interface {
	Login(models.Login) error
	SignUp(models.SignUp) error
	Register()
}

type AuthService struct {
	UserRepository         repositories.UserRepositoryInterface
	UserPasswordRepository repositories.UserPasswordRepositoryInterface
}

func NewAuthService() AuthServiceInterface {
	as := new(AuthService)
	as.UserRepository = repositories.NewUserRepository()
	as.UserPasswordRepository = repositories.NewUserPasswordRepository()
	return as
}

func (as *AuthService) Login(loginPayload models.Login) error {
	passwordBytes := utilities.GenerateSha512Hash(loginPayload.Password)
	passwordHash := utilities.ByteToString(passwordBytes)

	savedPasswordBytes, err := as.UserPasswordRepository.GetUserPassword(loginPayload.Email)
	if err != nil {
		fmt.Printf(`Error getting user password: %v`, err)
		return err
	}

	savedPasswordHash := utilities.ByteToString(savedPasswordBytes)
	if passwordHash != savedPasswordHash {
		return fmt.Errorf("passwords mismatch")
	}

	fmt.Println("Login Success")
	return nil
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
