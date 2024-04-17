package repository

import (
	"github.com/take0fit/knowledge-out/internal/domain/entity"
)

type OutputRepository interface {
	GetOutputs(outputId string) ([]*entity.Output, error)
	GetOutputDetail(userId string) (*entity.Output, error)
	CreateOutput(*entity.Output) error
}
