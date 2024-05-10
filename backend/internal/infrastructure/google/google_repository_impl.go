package google

import (
	"encoding/json"
	"fmt"
	"github.com/take0fit/knowledge-out/internal/domain/entity"
	"github.com/take0fit/knowledge-out/internal/domain/repository"
	"net/http"
)

type GoogleRepositoryImpl struct {
	client *http.Client
}

func NewGoogleRepository(client *http.Client) repository.GoogleRepository {
	return &GoogleRepositoryImpl{
		client: client,
	}
}

func (r *GoogleRepositoryImpl) FetchUserInfo(accessToken string) (*entity.GoogleUser, error) {
	req, err := http.NewRequest("GET", "https://www.googleapis.com/oauth2/v3/userinfo", nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Authorization", "Bearer "+accessToken)

	resp, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to fetch user info: %s", resp.Status)
	}

	var userInfo entity.GoogleUser
	if err := json.NewDecoder(resp.Body).Decode(&userInfo); err != nil {
		return nil, err
	}

	return &userInfo, nil
}
