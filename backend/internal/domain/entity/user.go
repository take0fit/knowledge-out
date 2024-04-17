package entity

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/take0fit/knowledge-out/internal/domain/valueobject"
	"time"
)

type User struct {
	Id        string               `DynamoDB:"Id"`
	Nickname  valueobject.NickName `DynamoDB:"Nickname"`
	CreatedAt string               `DynamoDB:"CreatedAt"`
	UpdatedAt string               `DynamoDB:"UpdatedAt"`
	Birthday  valueobject.Birthday
}

func NewUser(nickname valueobject.NickName, birthday time.Time) *User {
	userId := fmt.Sprintf("User#%s", uuid.New().String())
	newUser := User{
		Id:        userId,
		Nickname:  nickname,
		Birthday:  valueobject.NewBirthday(birthday),
		CreatedAt: time.Now().String(),
		UpdatedAt: time.Now().String(),
	}

	return &newUser
}

func (b User) GetBirthday() *time.Time {
	if b.Birthday.Valid {
		return &b.Birthday.Time
	}
	return nil
}
