package usecase

import (
	"context"
	"fmt"
	"github.com/take0fit/knowledge-out/internal/application/dto"
	"github.com/take0fit/knowledge-out/internal/domain/entity"
	"github.com/take0fit/knowledge-out/internal/domain/repository"
)

type OutputUseCaseInteractor struct {
	outputRepo repository.OutputRepository
}

func NewOutputInteractor(outputRepo repository.OutputRepository) *OutputUseCaseInteractor {
	return &OutputUseCaseInteractor{
		outputRepo: outputRepo,
	}
}

func (u *OutputUseCaseInteractor) GetOutputDetail(outputId string) (*entity.Output, error) {
	output, err := u.outputRepo.GetOutputDetail(outputId)
	if err != nil {
		return nil, err
	}

	return output, nil
}

func (u *OutputUseCaseInteractor) CreateOutput(ctx context.Context, input *dto.InputCreateOutput) (*entity.Output, error) {
	outputModel := entity.NewOutput(
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
