package responses

type UserResponse struct {
	UserUsername    string `json:"user_username"`
	UserEmail       string `json:"user_email"`
	UserFirstName   string `json:"user_first_name"`
	UserLastName    string `json:"user_last_name"`
	UserAddress     string `json:"user_address"`
	UserPhoneNumber string `json:"user_phone_number"`
}
