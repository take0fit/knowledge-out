package ssm

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ssm"
	"github.com/take0fit/knowledge-out/internal/domain/repository"
)

type ParameterStoreRepositoryImpl struct {
	Client *ssm.Client
}

func NewParameterStoreRepository(client *ssm.Client) repository.ParameterStoreRepository {
	return &ParameterStoreRepositoryImpl{Client: client}
}

func (repo *ParameterStoreRepositoryImpl) GetParameter(ctx context.Context, name string) (string, error) {
	output, err := repo.Client.GetParameter(ctx, &ssm.GetParameterInput{
		Name:           &name,
		WithDecryption: aws.Bool(true),
	})
	if err != nil {
		return "", err
	}
	return *output.Parameter.Value, nil
}
