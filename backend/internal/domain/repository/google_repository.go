package repository

import (
	"github.com/take0fit/knowledge-out/internal/domain/entity"
)

type GoogleRepository interface {
	FetchUserInfo(accessToken string) (*entity.GoogleUser, error)
}
