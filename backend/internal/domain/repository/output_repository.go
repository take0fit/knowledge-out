package repository

import (
	"github.com/take0fit/knowledge-out/internal/domain/entity"
)

type OutputRepository interface {
	ListOutputsByUserId(userId string) ([]*entity.Output, error)
	ListOutputsByUserIdAndCategoryId(userId string, categoryId int) ([]*entity.Output, error)
	GetOutputDetail(outputId string) (*entity.Output, error)
	CreateOutput(output *entity.Output) error
}
