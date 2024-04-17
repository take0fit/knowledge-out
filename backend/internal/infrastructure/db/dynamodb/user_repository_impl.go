package dynamodb

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/take0fit/knowledge-out/internal/domain/entity"
	"github.com/take0fit/knowledge-out/internal/domain/repository"
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

func (r *DynamoUserRepository) GetUserDetail(userId string) (*entity.User, error) {
	dataType := "UserCreatedAt"

	result, err := r.client.GetItem(context.TODO(), &dynamodb.GetItemInput{
		TableName: aws.String("MyDataModel"),
		Key: map[string]types.AttributeValue{
			"Id":       &types.AttributeValueMemberS{Value: userId},
			"DataType": &types.AttributeValueMemberS{Value: dataType},
		},
	})
	if err != nil {
		return nil, fmt.Errorf("failed to get item from DynamoDB, %w", err)
	}
	if len(result.Item) == 0 {
		return nil, fmt.Errorf("not found")
	}

	var user entity.User
	err = attributevalue.UnmarshalMap(result.Item, &user)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal result item, %w", err)
	}

	return &user, nil
}

func (r *DynamoUserRepository) CreateUser(user *entity.User) error {
	tableName := "MyDataModel"
	dataType := "UserCreatedAt"

	nameItem := map[string]types.AttributeValue{
		"Id":        &types.AttributeValueMemberS{Value: user.Id},
		"DataType":  &types.AttributeValueMemberS{Value: dataType},
		"DataValue": &types.AttributeValueMemberS{Value: user.CreatedAt},
		"UserName":  &types.AttributeValueMemberS{Value: user.Name},
		"Age":       &types.AttributeValueMemberN{Value: strconv.Itoa(user.Birthday)},
		"CreatedAt": &types.AttributeValueMemberS{Value: user.CreatedAt},
	}

	_, err := r.client.PutItem(context.TODO(), &dynamodb.PutItemInput{
		TableName: aws.String(tableName),
		Item:      nameItem,
	})
	if err != nil {
		return fmt.Errorf("failed to put item into DynamoDB: %w", err)
	}

	return nil
}

func (r *DynamoUserRepository) ListUsersSortedByCreatedAt(ascending bool) ([]*entity.User, error) {
	gsiName := "DataTypeDataValueIndex"
	dataType := "UserCreatedAt"

	input := &dynamodb.QueryInput{
		TableName:              aws.String("MyDataModel"),
		IndexName:              aws.String(gsiName),
		KeyConditionExpression: aws.String("DataType = :dataType AND DataValue BETWEEN :startValue AND :endValue"),
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":dataType":   &types.AttributeValueMemberS{Value: dataType},
			":startValue": &types.AttributeValueMemberS{Value: "2024-01-01T00:00:00Z"},
			":endValue":   &types.AttributeValueMemberS{Value: "2024-12-31T23:59:59Z"},
		},
		ScanIndexForward: aws.Bool(ascending), // trueで昇順、falseで降順
	}

	result, err := r.client.Query(context.TODO(), input)
	if err != nil {
		return nil, fmt.Errorf("failed to query items from DynamoDB: %w", err)
	}

	users := make([]*entity.User, 0)
	for _, item := range result.Items {
		var user entity.User
		err = attributevalue.UnmarshalMap(item, &user)
		if err != nil {
			return nil, fmt.Errorf("failed to unmarshal result item: %w", err)
		}
		users = append(users, &user)
	}

	return users, nil
}
