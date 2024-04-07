package model

import (
	"fmt"
	"github.com/google/uuid"
	"time"
)

type User struct {
	Id        string    `dynamodbav:"Id"`
	Name      string    `dynamodbav:"UserName"`
	Age       int       `dynamodbav:"Age"`
	CreatedAt time.Time `dynamodbav:"CreatedAt"`
}

func NewUser(name *UserName, age int) *User {
	userId := fmt.Sprintf("User#%s", uuid.New().String())
	newUser := User{
		Id:        userId,
		Name:      string(*name),
		Age:       age,
		CreatedAt: time.Now(),
	}

	return &newUser
}
