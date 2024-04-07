package usecase

import (
	"book-action/internal/application/dto"
	"book-action/internal/domain/model"
	"book-action/internal/domain/repository"
	"context"
	"fmt"
)

type ResourceUseCaseInteractor struct {
	resourceRepo repository.ResourceRepository
}

func NewResourceInteractor(resourceRepo repository.ResourceRepository) *ResourceUseCaseInteractor {
	return &ResourceUseCaseInteractor{
		resourceRepo: resourceRepo,
	}
}

func (u *ResourceUseCaseInteractor) GetResourceDetail(resourceId string) (*model.Resource, error) {
	resource, err := u.resourceRepo.GetResourceDetail(resourceId)
	if err != nil {
		return nil, err
	}

	return resource, nil
}

func (u *ResourceUseCaseInteractor) CreateResource(ctx context.Context, input dto.ResourceCreateInput) (*model.Resource, error) {

	resourceModel := model.NewResource(
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
