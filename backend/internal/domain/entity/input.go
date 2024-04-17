package entity

import (
	"fmt"
	"github.com/google/uuid"
	"time"
)

type Input struct {
	Id         string
	UserId     string
	ResourceId string
	Name       string
	Detail     string
	CategoryId int
	CreatedAt  time.Time
}

func NewInput(
	userId string,
	resourceId string,
	name string,
	detail string,
	categoryId int,
) *Input {
	inputId := fmt.Sprintf("Input#%s", uuid.New().String())
	newInput := Input{
		Id:         inputId,
		UserId:     userId,
		ResourceId: resourceId,
		Name:       name,
		Detail:     detail,
		CategoryId: categoryId,
		CreatedAt:  time.Now(),
	}

	return &newInput
}
