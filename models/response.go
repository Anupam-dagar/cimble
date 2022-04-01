package models

type ErrorResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

type LoginResponse struct {
	Id           string `json:"id"`
	Email        string `json:"email"`
	Token        string `json:"token"`
	RefreshToken string `json:"refreshToken"`
}

func (lr LoginResponse) ConstructLoginResponse(
	userId string,
	email string,
	token string,
	refreshToken string,
) LoginResponse {
	return LoginResponse{
		Id:           userId,
		Email:        email,
		Token:        token,
		RefreshToken: refreshToken,
	}
}

type Pagination struct {
	CurrentPage int `json:"currentPage"`
	TotalPages  int `json:"totalPages"`
}

type OrganisationsResponse struct {
	Organisations []OrganisationModel `json:"organisations"`
	Page          Pagination          `json:"page"`
}
