package request

type RefreshTokenRequest struct {
	Token string `json:"token"`
}

type EmailLoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
