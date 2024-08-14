package model

type AdminModel struct {
	AdminID  uint   `gorm:"PrimaryKey"`
	Username string `json:"username"`
	Password string `json:"password"`
}
