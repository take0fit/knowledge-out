package dto

import (
	"github.com/take0fit/knowledge-out/internal/domain/entity"
)

type OutputResource struct {
	Id         string
	UserId     string
	Name       string
	Detail     *string
	CategoryId int
}

func NewOutputResource(resource *entity.Resource) *OutputResource {
	return &OutputResource{
		Id:         resource.Id,
		UserId:     resource.UserId,
		Name:       resource.Name,
		Detail:     resource.Detail,
		CategoryId: resource.CategoryId,
	}
}

type OutputResources []*OutputResource

func NewOutputResources(resources []*entity.Resource) OutputResources {
	outputResources := make([]*OutputResource, len(resources))
	for i, resource := range resources {
		outputResources[i] = &OutputResource{
			Id:   resource.Id,
			Name: resource.Name,
		}
	}

	return outputResources
}

type InputCreateResource struct {
	UserId     string
	Name       string
	Detail     *string
	CategoryId int
}

func NewInputCreateResource(
	userId string,
	name string,
	detail *string,
	categoryId int,
) *InputCreateResource {
	return &InputCreateResource{
		UserId:     userId,
		Name:       name,
		Detail:     detail,
		CategoryId: categoryId,
	}
}
