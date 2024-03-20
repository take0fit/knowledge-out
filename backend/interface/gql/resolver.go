package gql

import (
	"book-action/interface/gql/model"
	"book-action/internal/usecase/user"
	"context"
)

type Resolver struct {
	userUsecase usecase.UserUsecase
}

func NewResolver(userUsecase usecase.UserUsecase) *Resolver {
	return &Resolver{
		userUsecase: userUsecase,
	}
}

// GetUserWithDetails はユーザーIDに基づいてユーザーとその詳細を取得します。
func (r *Resolver) GetUserWithDetails(ctx context.Context, userID string) (*model.User, error) {
	// ユースケースを使用してビジネスロジックを実行
	return r.userUsecase.GetUserWithDetails(userID)
}
