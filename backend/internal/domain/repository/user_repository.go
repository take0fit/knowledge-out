package repository

import (
	"github.com/take0fit/knowledge-out/internal/domain/entity"
)

type UserRepository interface {
	GetUserDetail(userId string) (*entity.User, error)
	ListUsersSortedByCreatedAt(ascending bool) ([]*entity.User, error)
	CreateUser(*entity.User) error
}
