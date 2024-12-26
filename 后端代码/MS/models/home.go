package models

type HomeInfo struct {
	ID                uint   `json:"home_id"`
	Cat_id            uint   `json:"cat_id"`
	Cat_name          string `json:"cat_name"`
	Cat_avatar        string `json:"cat_avatar"`
	Tot_amount        int    `json:"tot_amount"`
	Cat_male          int    `json:"cat_male"`
	Cat_female        int    `json:"cat_female"`
	Cat_adapt         int    `json:"cat_adapt"`
	Cat_sterilization int    `json:"cat_sterilization"`
	Date              string `json:"home_date"`
}
