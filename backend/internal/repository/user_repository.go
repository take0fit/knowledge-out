package repository

import (
	"book-action/internal/graph/model"
)

type UserRepository interface {
	GetUserWithDetails(userID string) (*model.User, error)
}
