package models

type User struct {
	Name     string `gorm:"primary_key;type:varchar(255);not null"`
	Email    string `gorm:"primary_key;not null"`
	Password string `gorm:"not null"`

	Otp_enabled  bool `gorm:"default:false;"`
	Otp_verified bool `gorm:"default:false;"`

	Otp_secret   string
	Otp_auth_url string
}

type RegisterUserInput struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" bindinig:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginUserInput struct {
	Email    string `json:"email" bindinig:"required"`
	Password string `json:"password" binding:"required"`
	Otp      string `json:"otp"`
}
