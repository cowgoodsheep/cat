package models

type CatInfo struct {
	ID            uint   `json:"cat_id"`
	Name          string `json:"cat_name"`
	Type          string `json:"cat_type"`
	Gender        string `json:"cat_gender" gorm:"varchar(10)"`
	Date          string `json:"cat_date"`
	Weight        int    `json:"cat_weight"`
	Sterilization bool   `json:"cat_sterilization" gorm:"varchar(10)"`
	Adapt         bool   `json:"cat_adapt" gorm:"varchar(10)"`
	Avatar        string `json:"cat_avatar"`
	Owner         uint   `json:"cat_owner"`
}
