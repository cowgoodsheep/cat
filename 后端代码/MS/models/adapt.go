package models

type AdaptInfo struct {
	ID      uint   `json:"adapt_id"`
	OwnerId uint   `json:"adapt_owner"`
	CatId   uint   `json:"adapt_cat"`
	Date    string `json:"adapt_date"`
	Status  string `json:"adapt_status"`
}
