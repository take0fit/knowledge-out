package aws

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"os"
)

var endpoint string
var region string

func init() {
	// 環境変数から設定を取得
	endpoint = os.Getenv("DYNAMODB_ENDPOINT")
	if endpoint == "" {
		panic("DYNAMODB_ENDPOINT is unset")
	}

	region = os.Getenv("AWS_REGION")
	if region == "" {
		panic("AWS_REGION is unset")
	}
}

func NewConfig(ctx context.Context) *aws.Config {
	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithRegion(region),
		config.WithEndpointResolverWithOptions(aws.EndpointResolverWithOptionsFunc(
			func(service, region string, options ...interface{}) (aws.Endpoint, error) {
				return aws.Endpoint{
					PartitionID:   "aws",
					URL:           endpoint,
					SigningRegion: region,
				}, nil
			},
		)),
	)
	if err != nil {
		panic("unable to load SDK config, " + err.Error())
	}

	return &cfg
}
