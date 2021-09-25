package response

type LoginResponse struct {
	Token string `json:"token"`
}

type RefreshTokenResponse struct {
	Token string `json:"token"`
}

func NewLoginResponse(token string) LoginResponse {
	return LoginResponse{
		Token: token,
	}
}

func NewRefreshTokenResponse(token string) RefreshTokenResponse {
	return RefreshTokenResponse{
		Token: token,
	}
}
