package services

import (
	"cimble/models"
	"cimble/repositories"
	"cimble/utilities"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
)

type AuthServiceInterface interface {
	Login(models.Login) (models.LoginResponse, error)
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

func (as *AuthService) Login(loginPayload models.Login) (loginResponse models.LoginResponse, err error) {
	passwordBytes := utilities.GenerateSha512Hash(loginPayload.Password)
	passwordHash := utilities.ByteToString(passwordBytes)

	user, err := as.UserPasswordRepository.GetUserPassword(loginPayload.Email)
	if err != nil {
		fmt.Printf(`Error getting user password: %v`, err)
		return loginResponse, err
	}

	savedPasswordHash := utilities.ByteToString(user.PasswordHash)
	if passwordHash != savedPasswordHash {
		return loginResponse, fmt.Errorf("either email or password is wrong")
	}

	token, err := as.GenerateToken(user.Email, user.ID)

	return loginResponse.ConstructLoginResponse(user, token), err
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

func (as *AuthService) GenerateToken(email string, userId string) (string, error) {
	claims := &models.JwtClaims{
		Email: email,
		Id:    userId,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 1).Unix(),
			Issuer:    "cimble",
			IssuedAt:  time.Now().Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	jwtSecret := utilities.GetEnvironmentVariableString("JWT_SECRET")
	jwtToken, err := token.SignedString([]byte(jwtSecret))
	if err != nil {
		fmt.Printf(`Error signing jwt token: %v`, err)
		return "", err
	}

	return jwtToken, nil
}
