package models

type SignUp struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	Password  string `json:"password,omitempty"`
}

func (s SignUp) ToCreateUserEntity(createdBy string) User {
	return User{
		FirstName: s.FirstName,
		LastName:  s.LastName,
		Email:     s.Email,
		CreatedBy: createdBy,
	}
}
