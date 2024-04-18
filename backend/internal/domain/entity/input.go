package entity

import (
	"fmt"
	"github.com/google/uuid"
	"time"
)

type Input struct {
	Id         string  `dynamodbav:"Id"`
	UserId     string  `dynamodbav:"UserId"`
	ResourceId string  `dynamodbav:"ResourceId"`
	Name       string  `dynamodbav:"Name"`
	Detail     *string `dynamodbav:"Detail,omitempty"`
	CategoryId int     `dynamodbav:"CategoryId"`
	CreatedAt  string  `dynamodbav:"CreatedAt"`
	UpdatedAt  string  `dynamodbav:"UpdatedAt"`
}

func NewInput(
	userId string,
	resourceId string,
	name string,
	detail *string,
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
		CreatedAt:  time.Now().Format("2006-01-02 15:04:05"),
		UpdatedAt:  time.Now().Format("2006-01-02 15:04:05"),
	}

	return &newInput
}
