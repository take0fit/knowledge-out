package repository

import (
	"book-action/internal/domain/model"
)

type UserRepository interface {
	GetUserDetails(userID string) (*model.User, error)
	CreateUser(*model.User) error
}
