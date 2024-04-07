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

func (r *DynamoUserRepository) GetUserDetail(userId string) (*model.User, error) {
	dataType := "userCreatedAt"

	result, err := r.client.GetItem(context.TODO(), &dynamodb.GetItemInput{
		TableName: aws.String("MyDataModel"),
		Key: map[string]types.AttributeValue{
			"Id":       &types.AttributeValueMemberS{Value: userId},
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
	dataType := "userCreatedAt"

	nameItem := map[string]types.AttributeValue{
		"Id":        &types.AttributeValueMemberS{Value: user.Id},
		"DataType":  &types.AttributeValueMemberS{Value: dataType},
		"DataValue": &types.AttributeValueMemberS{Value: user.CreatedAt.String()},
		"UserName":  &types.AttributeValueMemberS{Value: user.Name},
		"Age":       &types.AttributeValueMemberN{Value: strconv.Itoa(user.Age)},
		"CreatedAt": &types.AttributeValueMemberS{Value: user.CreatedAt.String()},
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

func (r *DynamoUserRepository) ListUsersSortedByCreatedAt(ascending bool) ([]*model.User, error) {
	gsiName := "DataValueIndex"
	dataType := "userCreatedAt"

	input := &dynamodb.QueryInput{
		TableName:              aws.String("MyDataModel"),
		IndexName:              aws.String(gsiName),
		KeyConditionExpression: aws.String("DataType = :dataType"),
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":dataType": &types.AttributeValueMemberS{Value: dataType},
		},
		ScanIndexForward: aws.Bool(ascending), // trueで昇順、falseで降順
	}

	// Query実行
	result, err := r.client.Query(context.TODO(), input)
	if err != nil {
		return nil, fmt.Errorf("failed to query items from DynamoDB: %w", err)
	}

	// 結果をモデルにアンマーシャル
	users := make([]*model.User, 0)
	for _, item := range result.Items {
		var user model.User
		err = attributevalue.UnmarshalMap(item, &user)
		if err != nil {
			return nil, fmt.Errorf("failed to unmarshal result item: %w", err)
		}
		users = append(users, &user)
	}

	return users, nil
}
