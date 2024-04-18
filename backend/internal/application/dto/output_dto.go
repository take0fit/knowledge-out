package dto

import "github.com/take0fit/knowledge-out/internal/domain/entity"

type OutputOutput struct {
	Id         string
	UserId     string
	InputIds   []string
	Name       string
	Detail     *string
	CategoryId int
	CreatedAt  string
	UpdatedAt  string
}

func NewOutputOutput(output *entity.Output) *OutputOutput {
	return &OutputOutput{
		Id:         output.Id,
		UserId:     output.UserId,
		InputIds:   output.InputIds,
		Name:       output.Name,
		Detail:     output.Detail,
		CategoryId: output.CategoryId,
		CreatedAt:  output.CreatedAt,
		UpdatedAt:  output.UpdatedAt,
	}
}

type OutputOutputs []*OutputOutput

func NewOutputOutputs(outputs []*entity.Output) OutputOutputs {
	outputOutputs := make([]*OutputOutput, len(outputs))
	for i, output := range outputs {
		outputOutputs[i] = &OutputOutput{
			Id:   output.Id,
			Name: output.Name,
		}
	}

	return outputOutputs
}

type InputCreateOutput struct {
	UserId     string
	InputIds   []string
	Name       string
	Detail     *string
	CategoryId int
}

func NewInputCreateOutput(
	userId string,
	inputIds []string,
	name string,
	detail *string,
	categoryId int,
) *InputCreateOutput {
	return &InputCreateOutput{
		UserId:     userId,
		InputIds:   inputIds,
		Name:       name,
		Detail:     detail,
		CategoryId: categoryId,
	}
}
