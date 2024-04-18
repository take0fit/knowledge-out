package repository

import (
	"github.com/take0fit/knowledge-out/internal/domain/entity"
)

type InputRepository interface {
	ListInputsByUserId(userId string) ([]*entity.Input, error)
	ListInputsByUserIdAndCategoryId(userId string, categoryId int) ([]*entity.Input, error)
	GetInputDetail(inputId string) (*entity.Input, error)
	CreateInput(input *entity.Input) error
}
