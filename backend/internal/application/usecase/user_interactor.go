package usecase

import (
	"book-action/internal/application/dto"
	"book-action/internal/domain/model"
	"book-action/internal/domain/repository"
	"context"
	"fmt"
)

type UserUseCaseInteractor struct {
	userRepo repository.UserRepository
}

func NewUserInteractor(userRepo repository.UserRepository) *UserUseCaseInteractor {
	return &UserUseCaseInteractor{
		userRepo: userRepo,
	}
}

func (u *UserUseCaseInteractor) GetUserDetails(userId string) (*model.User, error) {
	userModel, err := u.userRepo.GetUserDetail(userId)
	if err != nil {
		return nil, err
	}

	return userModel, nil
}

func (u *UserUseCaseInteractor) CreateUser(ctx context.Context, input dto.UserCreateInput) (*model.User, error) {

	userName, err := model.NewUserName(input.Name)
	if err != nil {
		return nil, fmt.Errorf("failed to create user: %w", err)
	}

	newUser := model.NewUser(userName, input.Age)

	err = u.userRepo.CreateUser(newUser)
	if err != nil {
		return nil, fmt.Errorf("failed to create user: %w", err)
	}

	return newUser, nil
}
