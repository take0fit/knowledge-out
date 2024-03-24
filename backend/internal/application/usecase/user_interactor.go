package usecase

import (
	"book-action/internal/application/dto"
	"book-action/internal/domain/model"
	"book-action/internal/domain/repository"
	"context"
	"fmt"
	"github.com/google/uuid"
)

type UserUseCaseInteractor struct {
	userRepo repository.UserRepository
}

func NewUserInteractor(userRepo repository.UserRepository) *UserUseCaseInteractor {
	return &UserUseCaseInteractor{
		userRepo: userRepo,
	}
}

func (u *UserUseCaseInteractor) GetUserDetails(userID string) (*model.User, error) {
	user, err := u.userRepo.GetUserDetails(userID)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u *UserUseCaseInteractor) CreateUser(ctx context.Context, input dto.UserCreateInput) (*model.User, error) {
	// UUIDを生成してユーザーIDを作成
	userID := fmt.Sprintf("User#%s", uuid.New().String())
	user := &model.User{
		ID:   userID,
		Name: input.Name,
		Age:  input.Age,
	}

	err := u.userRepo.CreateUser(user)
	if err != nil {
		return nil, fmt.Errorf("failed to create user: %w", err)
	}

	return user, nil
}
