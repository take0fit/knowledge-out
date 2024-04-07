package repository

import (
	"book-action/internal/domain/model"
)

type ResourceRepository interface {
	ListResourcesByUserId(userId string) ([]*model.Resource, error)
	GetResourceDetail(resourceId string) (*model.Resource, error)
	CreateResource(*model.Resource) error
}
