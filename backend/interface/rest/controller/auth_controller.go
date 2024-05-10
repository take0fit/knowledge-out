package controller

import (
	"encoding/json"
	"github.com/take0fit/knowledge-out/internal/application/usecase"
	"golang.org/x/oauth2"
	"log"
	"net/http"
)

type AuthController struct {
	AuthUseCase *usecase.AuthUseCaseInteractor
	OAuthConfig *oauth2.Config
}

func NewAuthController(authUseCase *usecase.AuthUseCaseInteractor, oauthConfig *oauth2.Config) *AuthController {
	return &AuthController{
		AuthUseCase: authUseCase,
		OAuthConfig: oauthConfig,
	}
}

func (c *AuthController) Login(w http.ResponseWriter, r *http.Request) {
	url := c.OAuthConfig.AuthCodeURL("state", oauth2.AccessTypeOffline)
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

func (c *AuthController) Callback(w http.ResponseWriter, r *http.Request) {
	code := r.URL.Query().Get("code")

	jwtToken, err := c.AuthUseCase.Authenticate(r.Context(), code)
	if err != nil {
		log.Printf("Failed : %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(map[string]string{
		"jwt": jwtToken,
	})
}
