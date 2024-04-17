package repository

import (
	"github.com/take0fit/knowledge-out/internal/domain/entity"
)

type InputRepository interface {
	GetInputs(inputId string) ([]*entity.Input, error)
	GetInputDetail(userId string) (*entity.Input, error)
	CreateInput(*entity.Input) error
}
