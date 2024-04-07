package usecase

import (
	"book-action/internal/application/dto"
	"book-action/internal/domain/model"
	"book-action/internal/domain/repository"
	"context"
	"fmt"
)

type InputUseCaseInteractor struct {
	inputRepo repository.InputRepository
}

func NewInputInteractor(inputRepo repository.InputRepository) *InputUseCaseInteractor {
	return &InputUseCaseInteractor{
		inputRepo: inputRepo,
	}
}

func (u *InputUseCaseInteractor) GetInputDetail(inputId string) (*model.Input, error) {
	inputModel, err := u.inputRepo.GetInputDetail(inputId)
	if err != nil {
		return nil, err
	}

	return inputModel, nil
}

func (u *InputUseCaseInteractor) CreateInput(ctx context.Context, input dto.InputCreateInput) (*model.Input, error) {

	inputModel := model.NewInput(
		input.UserId,
		input.ResourceId,
		input.Name,
		input.Detail,
		input.CategoryId,
	)

	err := u.inputRepo.CreateInput(inputModel)
	if err != nil {
		return nil, fmt.Errorf("failed to create input: %w", err)
	}

	return inputModel, nil
}
