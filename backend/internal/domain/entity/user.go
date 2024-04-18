package entity

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/take0fit/knowledge-out/internal/domain/valueobject"
	"time"
)

type User struct {
	Id        string `DynamoDB:"Id"`
	Nickname  valueobject.NickName
	Birthday  valueobject.Birthday
	CreatedAt string `DynamoDB:"CreatedAt"`
	UpdatedAt string `DynamoDB:"UpdatedAt"`
}

func NewUser(nickname string, birthday *string) (*User, error) {
	userId := fmt.Sprintf("User#%s", uuid.New().String())
	nicknameObj, err := valueobject.NewUserNickname(nickname)
	if err != nil {
		return nil, err
	}
	newUser := User{
		Id:        userId,
		Nickname:  nicknameObj,
		Birthday:  valueobject.NewBirthday(birthday),
		CreatedAt: time.Now().Format("2006-01-02 15:04:05"),
		UpdatedAt: time.Now().Format("2006-01-02 15:04:05"),
	}

	return &newUser, nil
}

func (b User) GetBirthday() *time.Time {
	if b.Birthday.Valid {
		return &b.Birthday.Time
	}

	return nil
}
