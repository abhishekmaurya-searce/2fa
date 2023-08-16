package graph

type User_response struct {
	Name     string `gorm:"type:varchar(255);not null"`
	Email    string `gorm:"primary_key;not null"`
	Password string `gorm:"not null"`

	Otp_enabled  bool `gorm:"default:false;"`
	Otp_verified bool `gorm:"default:false;"`

	Otp_secret   string `gorm:"type:varchar(255);not null"`
	Otp_auth_url string `gorm:"type:varchar(255);not null"`

	PublicKey string `gorm:"type:varchar(255);not null"`
}
