package model

type Resource struct {
	ID         string `json:"id"`
	UserID     string `json:"user_id"`
	Name       string `json:"name"`
	Detail     string `json:"detail"`
	CategoryId string `json:"category_id"`
}
