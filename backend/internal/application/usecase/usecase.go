package usecase

import (
	"github.com/take0fit/knowledge-out/interface/gql/model"
)

type UserUsecase interface {
	GetUserWithDetails(id string) (*model.User, error)
	CreateUser() (*model.User, error)
	//GetAllUsers() ([]*model.User, error)
}

type ResourceUsecase interface {
	GetResourceWithDetails(id string) (*model.Resource, error)
	CreateResource() (*model.Resource, error)
}

type InputUsecase interface {
	GetInputWithDetails(id string) (*model.Input, error)
	CreateInput() (*model.Input, error)
}
type OutputUsecase interface {
	GetOutputWithDetails(id string) (*model.Output, error)
	CreateOutput() (*model.Output, error)
}
