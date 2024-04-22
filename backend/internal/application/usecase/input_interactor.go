package usecase

import (
	"context"
	"fmt"
	"github.com/take0fit/knowledge-out/internal/application/dto"
	"github.com/take0fit/knowledge-out/internal/domain/entity"
	"github.com/take0fit/knowledge-out/internal/domain/repository"
)

type InputUseCaseInteractor struct {
	inputRepo repository.InputRepository
}

func NewInputInteractor(inputRepo repository.InputRepository) *InputUseCaseInteractor {
	return &InputUseCaseInteractor{
		inputRepo: inputRepo,
	}
}

func (u *InputUseCaseInteractor) GetInputDetail(inputId string) (*entity.Input, error) {
	inputModel, err := u.inputRepo.GetInputDetail(inputId)
	if err != nil {
		return nil, err
	}

	return inputModel, nil
}

func (u *InputUseCaseInteractor) CreateInput(ctx context.Context, input *dto.InputCreateInput) (*entity.Input, error) {

	inputModel := entity.NewInput(
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

func (u *InputUseCaseInteractor) GetInputListByUserId(userId string) (dto.OutputInputs, error) {
	inputs, err := u.inputRepo.ListInputsByUserId(userId)
	if err != nil {
		return nil, err
	}

	return dto.NewOutputInputs(inputs), nil
}

func (u *InputUseCaseInteractor) GetInputDetails(inputId string) (*dto.OutputInput, error) {
	input, err := u.inputRepo.GetInputDetail(inputId)
	if err != nil {
		return nil, err
	}

	return dto.NewOutputInput(input), nil
}
