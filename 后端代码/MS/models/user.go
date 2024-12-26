package models

type UserInfo struct {
	ID       uint   `json:"user_id"`
	Name     string `json:"user_name"`
	Password string `json:"user_password"`
	Gender   string `json:"user_gender" gorm:"varchar(10)"`
	Phone    string `json:"user_phone" gorm:"varchar(11)"`
	Age      uint   `json:"user_age"`
	Avatar   string `json:"user_avatar"`
}
