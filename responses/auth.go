package responses

type RegisterResponse struct {
	UserUsername string `json:"user_username"`
	UserEmail    string `json:"user_email"`
	UserPassword string `json:"user_password"`
}

type LoginResponse struct {
	UserUsername string `json:"user_username"`
	UserEmail    string `json:"user_email"`
	UserToken    string `json:"user_token"`
}
