package gql

import (
	"github.com/take0fit/knowledge-out/internal/application/usecase"
)

type Resolver struct {
	userUsecase     *usecase.UserUseCaseInteractor
	resourceUsecase *usecase.ResourceUseCaseInteractor
	inputUsecase    *usecase.InputUseCaseInteractor
	outputUsecase   *usecase.OutputUseCaseInteractor
}

func NewResolver(
	userUsecase *usecase.UserUseCaseInteractor,
	resourceUsecase *usecase.ResourceUseCaseInteractor,
	inputUsecase *usecase.InputUseCaseInteractor,
	outputUsecase *usecase.OutputUseCaseInteractor,
) *Resolver {
	return &Resolver{
		userUsecase:     userUsecase,
		resourceUsecase: resourceUsecase,
		inputUsecase:    inputUsecase,
		outputUsecase:   outputUsecase,
	}
}
