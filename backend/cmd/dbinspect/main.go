package main

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
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

	result, err := svc.DescribeTable(context.TODO(), &dynamodb.DescribeTableInput{
		TableName: aws.String("MyDataModel"),
	})
	if err != nil {
		fmt.Printf("error: %v\n", err)
		os.Exit(1)
	}

	table := result.Table
	fmt.Printf("Table Name: %s\n", *table.TableName)
	fmt.Printf("Table Status: %s\n", table.TableStatus)
	fmt.Printf("Provisioned Read Capacity Units: %d\n", table.ProvisionedThroughput.ReadCapacityUnits)
	fmt.Printf("Provisioned Write Capacity Units: %d\n", table.ProvisionedThroughput.WriteCapacityUnits)
	fmt.Println("Key Schema:")
	for _, element := range table.KeySchema {
		fmt.Printf("  Attribute Name: %s, KeyType: %s\n", *element.AttributeName, element.KeyType)
	}

	// LSIの情報を出力
	if table.LocalSecondaryIndexes != nil {
		fmt.Println("Local Secondary Indexes:")
		for _, lsi := range table.LocalSecondaryIndexes {
			fmt.Printf("  Index Name: %s\n", *lsi.IndexName)
			fmt.Println("  Key Schema:")
			for _, element := range lsi.KeySchema {
				fmt.Printf("    Attribute Name: %s, KeyType: %s\n", *element.AttributeName, element.KeyType)
			}
			fmt.Println("  Projection:")
			fmt.Printf("    Projection Type: %s\n", lsi.Projection.ProjectionType)
			if lsi.Projection.NonKeyAttributes != nil {
				fmt.Println("    Non-Key Attributes:")
				for _, attr := range lsi.Projection.NonKeyAttributes {
					fmt.Printf("      %s\n", attr)
				}
			}
		}
	}

}
