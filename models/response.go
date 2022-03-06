package models

type ErrorResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

type LoginResponse struct {
	Id    string `json:"id"`
	Email string `json:"email"`
	Token string `json:"token"`
}

func (lr LoginResponse) ConstructLoginResponse(user UserJoinUserPassword, token string) LoginResponse {
	return LoginResponse{
		Id:    user.ID,
		Email: user.Email,
		Token: token,
	}
}
