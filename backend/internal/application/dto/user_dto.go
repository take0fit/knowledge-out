package dto

import (
	"github.com/take0fit/knowledge-out/internal/domain/entity"
	"time"
)

type OutputUser struct {
	Id       string
	Nickname string
	Birthday *time.Time
	Age      *int
}

func NewOutputUser(user *entity.User) *OutputUser {
	return &OutputUser{
		Id:       user.Id,
		Nickname: string(user.Nickname),
		Birthday: user.GetBirthday(),
		Age:      user.Birthday.Age(),
	}
}

type OutputUsers []*OutputUser

func NewOutputUsers(users []*entity.User) OutputUsers {
	outputUsers := make([]*OutputUser, len(users))
	for i, user := range users {
		outputUsers[i] = &OutputUser{
			Id:       user.Id,
			Nickname: string(user.Nickname),
		}
	}

	return outputUsers
}

type InputCreateUser struct {
	Nickname string
	Birthday time.Time
}

func NewInputCreateUser(nickname string, birthday time.Time) *InputCreateUser {
	return &InputCreateUser{
		Nickname: nickname,
		Birthday: birthday,
	}
}
