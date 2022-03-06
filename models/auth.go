package models

import "github.com/segmentio/ksuid"

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
		CreatedBy: createdBy,
		UpdatedBy: createdBy,
	}
}
