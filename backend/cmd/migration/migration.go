package main

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"os"
)

func main() {
	endpoint := os.Getenv("DYNAMODB_ENDPOINT")
	if endpoint == "" {
		panic("DYNAMODB_ENDPOINT is unset")
	}

	region := os.Getenv("AWS_REGION")
	if region == "" {
		panic("AWS_REGION is unset")
	}

	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithRegion(region),
		// カスタムエンドポイントの設定
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
		panic(fmt.Errorf("unable to load SDK config, %w", err))
	}

	svc := dynamodb.NewFromConfig(cfg)

	tableName := "MyDataModel"
	_, err = svc.CreateTable(context.TODO(), &dynamodb.CreateTableInput{
		TableName: &tableName,
		KeySchema: []types.KeySchemaElement{
			{AttributeName: aws.String("Id"), KeyType: types.KeyTypeHash},        // パーティションキー
			{AttributeName: aws.String("DataType"), KeyType: types.KeyTypeRange}, // ソートキー
		},
		AttributeDefinitions: []types.AttributeDefinition{
			{AttributeName: aws.String("Id"), AttributeType: types.ScalarAttributeTypeS},
			{AttributeName: aws.String("DataType"), AttributeType: types.ScalarAttributeTypeS},
			{AttributeName: aws.String("DataValue"), AttributeType: types.ScalarAttributeTypeS}, // 追加の属性
		},
		// GSI1の定義 (DataValueを基にしたクエリをサポートするため)
		GlobalSecondaryIndexes: []types.GlobalSecondaryIndex{
			{
				IndexName: aws.String("DataValueIndex"),
				KeySchema: []types.KeySchemaElement{
					{AttributeName: aws.String("DataValue"), KeyType: types.KeyTypeHash},
					{AttributeName: aws.String("Id"), KeyType: types.KeyTypeRange},
				},
				Projection: &types.Projection{
					ProjectionType: types.ProjectionTypeAll,
				},
			},
		},
		BillingMode: types.BillingModePayPerRequest,
	})
	if err != nil {
		panic(fmt.Errorf("failed to create table %s, %w", tableName, err))
	}

	fmt.Printf("Table %s created successfully\n", tableName)
}
