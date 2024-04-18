package entity

import (
	"fmt"
	"github.com/google/uuid"
	"time"
)

type Output struct {
	Id         string `dynamodbav:"Id"`
	UserId     string `dynamodbav:"UserId"`
	InputIds   []string
	Name       string  `dynamodbav:"Name"`
	Detail     *string `dynamodbav:"Detail,omitempty"`
	CategoryId int     `dynamodbav:"CategoryId"`
	CreatedAt  string  `dynamodbav:"CreatedAt"`
	UpdatedAt  string  `dynamodbav:"UpdatedAt"`
}

func NewOutput(
	userId string,
	inputIds []string,
	name string,
	detail *string,
	categoryId int,
) *Output {
	outputId := fmt.Sprintf("Output#%s", uuid.New().String())
	newOutput := Output{
		Id:         outputId,
		UserId:     userId,
		InputIds:   inputIds,
		Name:       name,
		Detail:     detail,
		CategoryId: categoryId,
		CreatedAt:  time.Now().Format("2006-01-02 15:04:05"),
		UpdatedAt:  time.Now().Format("2006-01-02 15:04:05"),
	}

	return &newOutput
}
