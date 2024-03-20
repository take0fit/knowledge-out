package usecase

import (
	"book-action/interface/gql/model"
	"book-action/internal/domain/repository"
)

type UserUseCaseInteractor struct {
	repo repository.UserRepository
}

func NewUserInteractor(repo repository.UserRepository) *UserUseCaseInteractor {
	return &UserUseCaseInteractor{
		repo: repo,
	}
}

func (u *UserUseCaseInteractor) GetUserWithDetails(userID string) (*model.User, error) {
	return u.repo.GetUserWithDetails(userID)
}
