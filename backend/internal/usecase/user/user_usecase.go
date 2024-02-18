package usecase

import "book-action/internal/graph/model"

type UserUsecase interface {
	GetUserWithDetails(id string) (*model.User, error)
	//CreateUser(username, email string) (*model.User, error)
	//GetAllUsers() ([]*model.User, error)
}
