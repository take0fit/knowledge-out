package usecase

import (
	"context"
	"fmt"
	"github.com/take0fit/knowledge-out/internal/application/dto"
	"github.com/take0fit/knowledge-out/internal/domain/entity"
	"github.com/take0fit/knowledge-out/internal/domain/repository"
)

type UserUseCaseInteractor struct {
	userRepo repository.UserRepository
}

func NewUserInteractor(userRepo repository.UserRepository) *UserUseCaseInteractor {
	return &UserUseCaseInteractor{
		userRepo: userRepo,
	}
}

func (u *UserUseCaseInteractor) GetUserList() (dto.OutputUsers, error) {
	users, err := u.userRepo.ListUsersSortedByCreatedAt(true)
	if err != nil {
		return nil, err
	}

	return dto.NewOutputUsers(users), nil
}

func (u *UserUseCaseInteractor) GetUserDetails(userId string) (*dto.OutputUser, error) {
	user, err := u.userRepo.GetUserDetail(userId)
	if err != nil {
		return nil, err
	}

	return dto.NewOutputUser(user), nil
}

func (u *UserUseCaseInteractor) CreateUser(ctx context.Context, input *dto.InputCreateUser) (*dto.OutputUser, error) {
	newUser, err := entity.NewUser(input.Nickname, input.Birthday, nil)
	if err != nil {
		return nil, err
	}

	err = u.userRepo.CreateUser(newUser)
	if err != nil {
		return nil, fmt.Errorf("failed to create user: %w", err)
	}

	return dto.NewOutputUser(newUser), nil
}
