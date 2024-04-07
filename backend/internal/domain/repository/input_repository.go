package repository

import (
	"book-action/internal/domain/model"
)

type InputRepository interface {
	GetInputs(inputId string) ([]*model.Input, error)
	GetInputDetail(userId string) (*model.Input, error)
	CreateInput(*model.Input) error
}
