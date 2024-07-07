package inputs

type RegisterUserInput struct {
	UserUsername string `json:"user_username" binding:"required"`
	UserEmail    string `json:"user_email" binding:"required"`
	UserPassword string `json:"user_password" binding:"required"`
}

type LoginUserInput struct {
	UserUsername string `json:"user_username" binding:"required"`
	UserPassword string `json:"user_password" binding:"required"`
}
