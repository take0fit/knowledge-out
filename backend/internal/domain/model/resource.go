package model

import (
	"fmt"
	"github.com/google/uuid"
	"time"
)

type Resource struct {
	Id         string
	UserId     string
	Name       string
	Detail     string
	CategoryId int
	CreatedAt  time.Time
}

func NewResource(
	userId string,
	name string,
	detail string,
	categoryId int,
) *Resource {
	resourceId := fmt.Sprintf("Resource#%s", uuid.New().String())
	newResource := Resource{
		Id:         resourceId,
		UserId:     userId,
		Name:       name,
		Detail:     detail,
		CategoryId: categoryId,
		CreatedAt:  time.Now(),
	}

	return &newResource
}
