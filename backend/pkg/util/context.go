package util

import (
	"context"
	"fmt"
	"strconv"
)

// GetUserIDFromContext は、コンテキストからユーザーIDを取得し、それを整数型に変換します。
func GetUserIDFromContext(ctx context.Context) (int, error) {
	userIDValue := ctx.Value("userID")
	if userIDValue == nil {
		return 0, fmt.Errorf("userID is not found in context")
	}

	userIDStr, ok := userIDValue.(string)
	if !ok {
		return 0, fmt.Errorf("userID in context is not a string")
	}

	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		return 0, fmt.Errorf("failed to convert userID to int: %v", err)
	}

	return userID, nil
}
