package repository

import (
	"book-action/internal/domain/model"
)

type OutputRepository interface {
	GetOutputs(outputId string) ([]*model.Output, error)
	GetOutputDetail(userId string) (*model.Output, error)
	CreateOutput(*model.Output) error
}
