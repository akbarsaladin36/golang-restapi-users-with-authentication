package inputs

type ProfileUserInput struct {
	UserUsername    string `json:"user_username" binding:"required"`
	UserEmail       string `json:"user_email" binding:"required"`
	UserFirstName   string `json:"user_first_name"`
	UserLastName    string `json:"user_last_name"`
	UserAddress     string `json:"user_address"`
	UserPhoneNumber string `json:"user_phone_number"`
}

type NewUserInput struct {
	UserUsername    string `json:"user_username" binding:"required"`
	UserPassword    string `json:"user_password" binding:"required"`
	UserEmail       string `json:"user_email" binding:"required"`
	UserFirstName   string `json:"user_first_name"`
	UserLastName    string `json:"user_last_name"`
	UserAddress     string `json:"user_address"`
	UserPhoneNumber string `json:"user_phone_number"`
}
