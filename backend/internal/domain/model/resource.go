package model

type Resource struct {
	ID         string `json:"id"`
	UserID     string `json:"user_id"`
	Name       string `json:"name"`
	CategoryId string `json:"category_id"`
}
