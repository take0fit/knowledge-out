package usecase

import (
	"context"
	"github.com/golang-jwt/jwt/v4"
	"github.com/take0fit/knowledge-out/internal/domain/entity"
	"github.com/take0fit/knowledge-out/internal/domain/repository"
	"io/ioutil"
	"os"
	"time"
)

type AuthUseCaseInteractor struct {
	ar  repository.AuthRepository
	gr  repository.GoogleRepository
	psr repository.ParameterStoreRepository
}

func NewAuthUseCaseInteractor(
	ar repository.AuthRepository,
	gr repository.GoogleRepository,
	psr repository.ParameterStoreRepository,
) *AuthUseCaseInteractor {
	return &AuthUseCaseInteractor{
		ar:  ar,
		gr:  gr,
		psr: psr,
	}
}

func (uc *AuthUseCaseInteractor) Authenticate(ctx context.Context, code string) (string, error) {
	// トークン交換
	token, err := uc.ar.ExchangeToken(code)
	if err != nil {
		return "", err
	}

	// Google APIからユーザーデータを取得
	googleUserInfo, err := uc.gr.FetchUserInfo(token.AccessToken)
	if err != nil {
		return "", err
	}

	user, err := entity.MapGoogleUserToUser(googleUserInfo)
	if err != nil {
		return "", err
	}

	// 認証データを保存
	err = uc.ar.SaveAuthenticationData(token, user)
	if err != nil {
		return "", err
	}

	var privateKeyStr string
	if os.Getenv("ENVIRONMENT") == "production" {
		privateKeyStr, err = getKeyFromSSM(ctx, uc, os.Getenv("AWS_SSM_PRIVATE_KEY_PARAMETER"))
	} else {
		privateKeyStr, err = getKeyFromFile(os.Getenv("PRIVATE_KEY_PATH"))
	}

	if err != nil {
		return "", err
	}

	jwtToken, err := generateJWT(privateKeyStr, user.Id)
	if err != nil {
		return "", err
	}

	return jwtToken, nil
}

func generateJWT(privateKeyStr string, userId string) (string, error) {
	privateKey, err := jwt.ParseRSAPrivateKeyFromPEM([]byte(privateKeyStr))
	if err != nil {
		return "", err
	}

	claims := jwt.MapClaims{
		"sub": userId,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	signedToken, err := token.SignedString(privateKey)
	if err != nil {
		return "", err
	}
	return signedToken, nil
}

func getKeyFromSSM(ctx context.Context, uc *AuthUseCaseInteractor, parameterName string) (string, error) {
	secretKeyParameter, err := uc.psr.GetParameter(ctx, parameterName)
	if err != nil {
		return "", err
	}
	return secretKeyParameter, nil
}

// ファイルから秘密鍵を取得
func getKeyFromFile(filePath string) (string, error) {
	key, err := ioutil.ReadFile(filePath)
	if err != nil {
		return "", err
	}
	return string(key), nil
}
