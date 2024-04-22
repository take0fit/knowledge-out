package dto

import (
	"github.com/take0fit/knowledge-out/internal/domain/entity"
)

type OutputUser struct {
	Id        string
	Nickname  string
	Birthday  *string
	Age       *int
	CreatedAt string
	UpdatedAt string
}

func NewOutputUser(user *entity.User) *OutputUser {
	return &OutputUser{
		Id:        user.Id,
		Nickname:  string(user.Nickname),
		Birthday:  user.Birthday.String(),
		Age:       user.Birthday.Age(),
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}

type OutputUsers []*OutputUser

func NewOutputUsers(users []*entity.User) OutputUsers {
	outputUsers := make([]*OutputUser, len(users))
	for i, user := range users {
		outputUsers[i] = &OutputUser{
			Id:        user.Id,
			Nickname:  string(user.Nickname),
			Birthday:  user.Birthday.String(),
			Age:       user.Birthday.Age(),
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
		}
	}

	return outputUsers
}

type InputCreateUser struct {
	Nickname string
	Birthday *string
}

func NewInputCreateUser(nickname string, birthday *string) *InputCreateUser {
	return &InputCreateUser{
		Nickname: nickname,
		Birthday: birthday,
	}
}
