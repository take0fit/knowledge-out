package gql

import (
	"book-action/internal/application/usecase"
)

type Resolver struct {
	userUsecase *usecase.UserUseCaseInteractor
}

func NewResolver(userUsecase *usecase.UserUseCaseInteractor) *Resolver {
	return &Resolver{
		userUsecase: userUsecase,
	}
}
