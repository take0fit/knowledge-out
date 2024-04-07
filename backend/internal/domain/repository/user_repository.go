package repository

import (
	"book-action/internal/domain/model"
)

type UserRepository interface {
	GetUserDetail(userId string) (*model.User, error)
	CreateUser(*model.User) error
}
