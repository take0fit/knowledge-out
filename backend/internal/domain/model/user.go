package model

import (
	"fmt"
	"github.com/google/uuid"
	"time"
)

type User struct {
	Id        string `dynamodbav:"Id"`
	Name      string `dynamodbav:"userName"`
	Age       int    `dynamodbav:"Age"`
	CreatedAt string `dynamodbav:"CreatedAt"`
}

func NewUser(name *UserName, age int) *User {
	userId := fmt.Sprintf("User#%s", uuid.New().String())
	newUser := User{
		Id:        userId,
		Name:      string(*name),
		Age:       age,
		CreatedAt: time.Now().String(),
	}

	return &newUser
}
