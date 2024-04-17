package repository

import (
	"github.com/take0fit/knowledge-out/internal/domain/entity"
)

type ResourceRepository interface {
	ListResourcesByUserId(userId string) ([]*entity.Resource, error)
	GetResourceDetail(resourceId string) (*entity.Resource, error)
	CreateResource(*entity.Resource) error
}
