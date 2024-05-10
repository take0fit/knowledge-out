package repository

import "context"

type ParameterStoreRepository interface {
	GetParameter(ctx context.Context, name string) (string, error)
}
