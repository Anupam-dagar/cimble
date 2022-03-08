package models

import (
	"github.com/golang-jwt/jwt"
	"github.com/segmentio/ksuid"
)

type SignUp struct {
	FirstName string `json:"firstName" binding:"required"`
	LastName  string `json:"lastName" binding:"required"`
	Email     string `json:"email" binding:"required"`
	Password  string `json:"password" binding:"required"`
}

func (s SignUp) CreateUserEntity(createdBy string) User {
	return User{
		ID:        ksuid.New().String(),
		FirstName: s.FirstName,
		LastName:  s.LastName,
		Email:     s.Email,
		BaseEntity: BaseEntity{
			CreatedBy: createdBy,
			UpdatedBy: createdBy,
		},
	}
}

type Login struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type JwtClaims struct {
	Email string `json:"email"`
	Id    string `json:"id"`
	jwt.StandardClaims
}

type RefreshTokenClaims struct {
	Id string `json:"id"`
	jwt.StandardClaims
}

type RefreshToken struct {
	RefreshToken string `json:"refreshToken" binding:"required"`
}
