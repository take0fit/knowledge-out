package dto

import (
	"github.com/take0fit/knowledge-out/internal/domain/entity"
)

type OutputInput struct {
	Id         string
	UserId     string
	ResourceId string
	Name       string
	Detail     *string
	CategoryId int
	CreatedAt  string
	UpdatedAt  string
}

func NewOutputInput(input *entity.Input) *OutputInput {
	return &OutputInput{
		Id:         input.Id,
		UserId:     input.UserId,
		ResourceId: input.ResourceId,
		Name:       input.Name,
		Detail:     input.Detail,
		CategoryId: input.CategoryId,
		CreatedAt:  input.CreatedAt,
		UpdatedAt:  input.UpdatedAt,
	}
}

type OutputInputs []*OutputInput

func NewOutputInputs(inputs []*entity.Input) OutputInputs {
	outputInputs := make([]*OutputInput, len(inputs))
	for i, input := range inputs {
		outputInputs[i] = &OutputInput{
			Id:   input.Id,
			Name: input.Name,
		}
	}

	return outputInputs
}

type InputCreateInput struct {
	UserId     string
	ResourceId string
	Name       string
	Detail     *string
	CategoryId int
}

func NewInputCreateInput(
	userId string,
	resourceId string,
	name string,
	detail *string,
	categoryId int,
) *InputCreateInput {
	return &InputCreateInput{
		UserId:     userId,
		ResourceId: resourceId,
		Name:       name,
		Detail:     detail,
		CategoryId: categoryId,
	}
}
