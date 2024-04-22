package usecase

import (
	"context"
	"fmt"
	"github.com/take0fit/knowledge-out/internal/application/dto"
	"github.com/take0fit/knowledge-out/internal/domain/entity"
	"github.com/take0fit/knowledge-out/internal/domain/repository"
)

type ResourceUseCaseInteractor struct {
	resourceRepo repository.ResourceRepository
}

func NewResourceInteractor(resourceRepo repository.ResourceRepository) *ResourceUseCaseInteractor {
	return &ResourceUseCaseInteractor{
		resourceRepo: resourceRepo,
	}
}

func (u *ResourceUseCaseInteractor) GetResourceDetail(resourceId string) (*entity.Resource, error) {
	resource, err := u.resourceRepo.GetResourceDetail(resourceId)
	if err != nil {
		return nil, err
	}

	return resource, nil
}

func (u *ResourceUseCaseInteractor) CreateResource(ctx context.Context, input *dto.InputCreateResource) (*entity.Resource, error) {

	resourceModel := entity.NewResource(
		input.UserId,
		input.Name,
		input.Detail,
		input.CategoryId,
	)

	err := u.resourceRepo.CreateResource(resourceModel)
	if err != nil {
		return nil, fmt.Errorf("failed to create resource: %w", err)
	}

	return resourceModel, nil
}

func (u *ResourceUseCaseInteractor) GetResourceListByUserId(userId string) (dto.OutputResources, error) {
	resources, err := u.resourceRepo.ListResourcesByUserId(userId)
	if err != nil {
		return nil, err
	}

	return dto.NewOutputResources(resources), nil
}

func (u *ResourceUseCaseInteractor) GetResourceDetails(resourceId string) (*dto.OutputResource, error) {
	resource, err := u.resourceRepo.GetResourceDetail(resourceId)
	if err != nil {
		return nil, err
	}

	return dto.NewOutputResource(resource), nil
}
