package main

import (
	"context"
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

func main() {
	// AWS設定をロード
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion("us-west-2"))
	if err != nil {
		panic(fmt.Errorf("unable to load SDK config, %w", err))
	}

	// DynamoDBクライアントの初期化
	svc := dynamodb.NewFromConfig(cfg)

	userIDs := []string{"user-00001", "user-00002"}
	for _, userID := range userIDs {
		entityType := "User"
		now := time.Now().Format(time.RFC3339)
		putItem(svc, userID, entityType+"#"+now, map[string]string{
			"UserName": "John Doe",
			"Birthday": "1990-01-01",
		})

		// ResourceIDごとのループ（リソースを示す）
		resourceIDs := []string{"resource-00001", "resource-00002"}
		for _, resourceID := range resourceIDs {
			entityType = "Resource"
			now = time.Now().Add(time.Minute).Format(time.RFC3339) // 時間をずらす
			putItem(svc, userID, entityType+"#"+now, map[string]string{
				"ResourceID":   resourceID,
				"ResourceName": "Sample Resource",
				"ResourceUrl":  "https://example.com/resource",
			})

			// InputIDごとのループ（インプットを示す）
			inputIDs := []string{"input-00001", "input-00002"}
			for _, inputID := range inputIDs {
				entityType = "Input"
				now = time.Now().Add(2 * time.Minute).Format(time.RFC3339) // 時間をさらにずらす
				putItem(svc, userID, entityType+"#"+now, map[string]string{
					"InputID":     inputID,
					"InputTitle":  "Sample Input",
					"InputDetail": "This is a sample input detail.",
				})

				// OutputIDごとのループ（アウトプットを示す）
				outputIDs := []string{"output-00001", "output-00002"}
				for _, outputID := range outputIDs {
					entityType = "Output"
					now = time.Now().Add(3 * time.Minute).Format(time.RFC3339) // 時間をさらにずらす
					putItem(svc, userID, entityType+"#"+now, map[string]string{
						"OutputID":     outputID,
						"OutputTitle":  "Sample Output",
						"OutputDetail": "This is a sample output detail.",
					})
				}
			}
		}
	}
	fmt.Println("Dummy data inserted successfully.")
}

func putItem(svc *dynamodb.Client, userID, sortKey string, additionalAttributes map[string]string) {
	item := map[string]types.AttributeValue{
		"UserID":               &types.AttributeValueMemberS{Value: userID},
		"EntityType#CreatedAt": &types.AttributeValueMemberS{Value: sortKey},
	}

	// 追加属性をアイテムに含める
	for key, value := range additionalAttributes {
		item[key] = &types.AttributeValueMemberS{Value: value}
	}

	_, err := svc.PutItem(context.TODO(), &dynamodb.PutItemInput{
		TableName: aws.String("UserEntities"),
		Item:      item,
	})
	if err != nil {
		panic(fmt.Errorf("failed to put item: %w", err))
	}
}
