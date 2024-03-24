package dynamodb

import (
	"book-action/internal/domain/model"
	"book-action/internal/domain/repository"
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"strconv"
)

type DynamoUserRepository struct {
	client *dynamodb.Client
}

func NewDynamoUserRepository(client *dynamodb.Client) repository.UserRepository {
	return &DynamoUserRepository{
		client: client,
	}
}

func (r *DynamoUserRepository) GetUserDetails(userID string) (*model.User, error) {
	dataType := "userInfo"

	// DynamoDBからデータを取得
	result, err := r.client.GetItem(context.TODO(), &dynamodb.GetItemInput{
		TableName: aws.String("MyDataModel"),
		Key: map[string]types.AttributeValue{
			"Id":       &types.AttributeValueMemberS{Value: userID},
			"DataType": &types.AttributeValueMemberS{Value: dataType},
		},
	})
	if err != nil {
		panic(fmt.Errorf("failed to get item from DynamoDB, %w", err))
	}

	var user model.User
	err = attributevalue.UnmarshalMap(result.Item, &user)
	if err != nil {
		panic(fmt.Errorf("failed to unmarshal result item, %w", err))
	}

	return &user, nil
}

func (r *DynamoUserRepository) CreateUser(user *model.User) error {
	tableName := "MyDataModel"

	nameItem := map[string]types.AttributeValue{
		"Id":        &types.AttributeValueMemberS{Value: user.ID},
		"DataType":  &types.AttributeValueMemberS{Value: "UserName"},
		"DataValue": &types.AttributeValueMemberS{Value: user.Name},
		"userName":  &types.AttributeValueMemberS{Value: user.Name},
		"Age":       &types.AttributeValueMemberN{Value: strconv.Itoa(user.Age)},
	}

	_, err := r.client.PutItem(context.TODO(), &dynamodb.PutItemInput{
		TableName: aws.String(tableName),
		Item:      nameItem,
	})
	if err != nil {
		return fmt.Errorf("failed to put item into DynamoDB: %w", err)
	}

	if err != nil {
		return fmt.Errorf("failed to put item into DynamoDB: %w", err)
	}

	return nil
}
