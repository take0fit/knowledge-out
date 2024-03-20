package repository

import (
	"book-action/interface/gql/model"
)

type UserRepository interface {
	GetUserWithDetails(userID string) (*model.User, error)
}
