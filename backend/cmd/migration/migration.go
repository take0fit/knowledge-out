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
	region := os.Getenv("AWS_REGION")
	if region == "" {
		region = "us-west-2" // デフォルトのリージョン設定
	}

	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithRegion(region),
		// カスタムエンドポイントの設定
		config.WithEndpointResolverWithOptions(aws.EndpointResolverWithOptionsFunc(
			func(service, region string, options ...interface{}) (aws.Endpoint, error) {
				if endpoint != "" {
					return aws.Endpoint{
						PartitionID:   "aws",
						URL:           endpoint,
						SigningRegion: region,
					}, nil
				}
				// カスタムエンドポイントが指定されていない場合、デフォルトの解決方法を使用
				return aws.Endpoint{}, &aws.EndpointNotFoundError{}
			},
		)),
	)
	if err != nil {
		panic(fmt.Errorf("unable to load SDK config, %w", err))
	}

	svc := dynamodb.NewFromConfig(cfg)

	tableName := "UserEntities"
	_, err = svc.CreateTable(context.TODO(), &dynamodb.CreateTableInput{
		TableName: &tableName,
		KeySchema: []types.KeySchemaElement{
			{AttributeName: aws.String("UserID"), KeyType: types.KeyTypeHash},                // パーティションキー
			{AttributeName: aws.String("EntityType#CreatedAt"), KeyType: types.KeyTypeRange}, // ソートキー
		},
		AttributeDefinitions: []types.AttributeDefinition{
			{AttributeName: aws.String("UserID"), AttributeType: types.ScalarAttributeTypeS},
			{AttributeName: aws.String("EntityType#CreatedAt"), AttributeType: types.ScalarAttributeTypeS},
			// LSI用の属性定義
			{AttributeName: aws.String("ResourceID"), AttributeType: types.ScalarAttributeTypeS},
			{AttributeName: aws.String("InputID"), AttributeType: types.ScalarAttributeTypeS},
			{AttributeName: aws.String("OutputID"), AttributeType: types.ScalarAttributeTypeS},
			{AttributeName: aws.String("Age#JobId"), AttributeType: types.ScalarAttributeTypeS},
		},
		LocalSecondaryIndexes: []types.LocalSecondaryIndex{
			{
				IndexName: aws.String("ResourceIDIndex"),
				KeySchema: []types.KeySchemaElement{
					{AttributeName: aws.String("UserID"), KeyType: types.KeyTypeHash},
					{AttributeName: aws.String("ResourceID"), KeyType: types.KeyTypeRange},
				},
				Projection: &types.Projection{
					ProjectionType: types.ProjectionTypeAll,
				},
			},
			{
				IndexName: aws.String("InputIDIndex"),
				KeySchema: []types.KeySchemaElement{
					{AttributeName: aws.String("UserID"), KeyType: types.KeyTypeHash},
					{AttributeName: aws.String("InputID"), KeyType: types.KeyTypeRange},
				},
				Projection: &types.Projection{
					ProjectionType: types.ProjectionTypeAll,
				},
			},
			{
				IndexName: aws.String("OutputIDIndex"),
				KeySchema: []types.KeySchemaElement{
					{AttributeName: aws.String("UserID"), KeyType: types.KeyTypeHash},
					{AttributeName: aws.String("OutputID"), KeyType: types.KeyTypeRange},
				},
				Projection: &types.Projection{
					ProjectionType: types.ProjectionTypeAll,
				},
			},
			{
				IndexName: aws.String("AgeJobIdIndex"),
				KeySchema: []types.KeySchemaElement{
					{AttributeName: aws.String("UserID"), KeyType: types.KeyTypeHash},
					{AttributeName: aws.String("Age#JobId"), KeyType: types.KeyTypeRange},
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
