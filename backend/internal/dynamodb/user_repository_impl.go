package dynamodb

import (
	"book-action/internal/graph/model"
	"book-action/internal/repository"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

type DynamoDBUserRepository struct {
	client *dynamodb.DynamoDB
}

func NewDynamoDBUserRepository() repository.UserRepository {
	// 環境変数から設定を取得
	endpoint := os.Getenv("DYNAMODB_ENDPOINT")
	region := os.Getenv("AWS_REGION")

	// カスタムセッションの作成
	sess := session.Must(session.NewSession(&aws.Config{
		Region:   aws.String(region),
		Endpoint: aws.String(endpoint),
	}))

	return &DynamoDBUserRepository{
		client: dynamodb.New(sess),
	}
}

func (r *DynamoDBUserRepository) GetUserWithDetails(userID string) (*model.User, error) {
	// DynamoDBからユーザー情報を取得するロジックを実装します。
	// この例ではシンプルな取得方法を示しますが、実際にはユーザー、書籍、入力、出力データを結合して取得する必要があります。

	// 仮の実装例
	result, err := r.client.GetItem(&dynamodb.GetItemInput{
		TableName: aws.String("YourTableName"),
		Key: map[string]*dynamodb.AttributeValue{
			"UserID": {
				S: aws.String(userID),
			},
		},
	})
	if err != nil {
		return nil, fmt.Errorf("failed to get user from DynamoDB: %w", err)
	}

	var user model.User
	err = dynamodbattribute.UnmarshalMap(result.Item, &user)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal DynamoDB result to user: %w", err)
	}

	// ここで関連する書籍、入力、出力データを取得し、userオブジェクトに追加します。

	return &user, nil
}
