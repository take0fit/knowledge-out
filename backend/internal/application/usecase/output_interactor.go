package usecase

import (
	"book-action/internal/application/dto"
	"book-action/internal/domain/model"
	"book-action/internal/domain/repository"
	"context"
	"fmt"
)

type OutputUseCaseInteractor struct {
	outputRepo repository.OutputRepository
}

func NewOutputInteractor(outputRepo repository.OutputRepository) *OutputUseCaseInteractor {
	return &OutputUseCaseInteractor{
		outputRepo: outputRepo,
	}
}

func (u *OutputUseCaseInteractor) GetOutputDetail(outputId string) (*model.Output, error) {
	output, err := u.outputRepo.GetOutputDetail(outputId)
	if err != nil {
		return nil, err
	}

	return output, nil
}

func (u *OutputUseCaseInteractor) CreateOutput(ctx context.Context, input dto.OutputCreateInput) (*model.Output, error) {
	outputModel := model.NewOutput(
		input.UserId,
		input.InputIds,
		input.Name,
		input.Detail,
		input.CategoryId,
	)

	err := u.outputRepo.CreateOutput(outputModel)
	if err != nil {
		return nil, fmt.Errorf("failed to create output: %w", err)
	}

	return outputModel, nil
}
