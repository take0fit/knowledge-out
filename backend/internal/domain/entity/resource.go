package entity

import (
	"fmt"
	"github.com/google/uuid"
	"time"
)

type Resource struct {
	Id         string
	UserId     string
	Name       string
	Detail     *string
	CategoryId int
	CreatedAt  string
	UpdatedAt  string
}

func NewResource(
	userId string,
	name string,
	detail *string,
	categoryId int,
) *Resource {
	resourceId := fmt.Sprintf("Resource#%s", uuid.New().String())
	newResource := Resource{
		Id:         resourceId,
		UserId:     userId,
		Name:       name,
		Detail:     detail,
		CategoryId: categoryId,
		CreatedAt:  time.Now().Format("2006-01-02 15:04:05"),
		UpdatedAt:  time.Now().Format("2006-01-02 15:04:05"),
	}

	return &newResource
}
