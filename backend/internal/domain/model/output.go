package model

import "time"

type Output struct {
	ID        string    `json:"id"`
	InputID   string    `json:"inputId"`
	Name      string    `json:"name"`
	Detail    string    `json:"detail"`
	CreatedAt time.Time `json:"createdAt"`
}
