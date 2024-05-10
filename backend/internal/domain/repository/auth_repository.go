package repository

import (
	"github.com/take0fit/knowledge-out/internal/domain/entity"
	"golang.org/x/oauth2"
)

type AuthRepository interface {
	ExchangeToken(code string) (*oauth2.Token, error)
	SaveAuthenticationData(token *oauth2.Token, user *entity.User) error
}
