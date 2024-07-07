package models

import (
	"time"
)

type User struct {
	UserID          int    `gorm:"primaryKey" json:"user_id"`
	UserUsername    string `json:"user_username" gorm:"type:varchar(100)"`
	UserEmail       string `json:"user_email" gorm:"type:varchar(100)"`
	UserPassword    string `json:"user_password" gorm:"type:varchar(255)"`
	UserFirstName   string `json:"user_first_name" gorm:"type:varchar(100)"`
	UserLastName    string `json:"user_last_name" gorm:"type:varchar(100)"`
	UserAddress     string `json:"user_address" gorm:"type:text"`
	UserPhoneNumber string `json:"user_phone_number" gorm:"type:varchar(30)"`
	UserCreatedDate time.Time
	UserUpdatedDate time.Time
}
