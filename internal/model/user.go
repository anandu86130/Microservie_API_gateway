package model

type UserModel struct {
	UserID uint   `gorm:"PrimaryKey"`
	Email  string `json:"email" validate:"required"`
}

type VerifyOTP struct {
	Email string `json:"email"`
	Otp   string `json:"otp"`
}
