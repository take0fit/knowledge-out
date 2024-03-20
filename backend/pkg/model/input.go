package model

import "time"

type Input struct {
	ID         string    `json:"id"`
	ResourceID string    `json:"resourceId"`
	Name       string    `json:"name"`
	Detail     string    `json:"detail"`
	CreatedAt  time.Time `json:"createdAt"`
}
