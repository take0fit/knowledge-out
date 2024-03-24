package usecase

import (
	"book-action/interface/gql/model"
)

type UserUsecase interface {
	GetUserWithDetails(id string) (*model.User, error)
	CreateUser() (*model.User, error)
	//GetAllUsers() ([]*model.User, error)
}
