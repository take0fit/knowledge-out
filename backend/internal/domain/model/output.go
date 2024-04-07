package model

import (
	"fmt"
	"github.com/google/uuid"
	"time"
)

type Output struct {
	Id         string
	UserId     string
	InputIds   []string
	Name       string
	Detail     string
	CategoryId int
	CreatedAt  time.Time
}

func NewOutput(
	userId string,
	inputIds []string,
	name string,
	detail string,
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
		CreatedAt:  time.Now(),
	}

	return &newOutput
}
