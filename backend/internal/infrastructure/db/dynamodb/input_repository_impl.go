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

type DynamoInputRepository struct {
	client *dynamodb.Client
}

func NewDynamoInputRepository(client *dynamodb.Client) repository.InputRepository {
	return &DynamoInputRepository{
		client: client,
	}
}

func (r *DynamoInputRepository) ListInputsByUserId(userId string) ([]*entity.Input, error) {
	gsiName := "DataTypeDataValueIndex"
	partitionKeyName := "InputUserId#CategoryId"

	input := &dynamodb.QueryInput{
		TableName:              aws.String("MyDataModel"),
		IndexName:              aws.String(gsiName),
		KeyConditionExpression: aws.String(partitionKeyName + " BEGINS_WITH :userIdPrefix"),
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":userIdPrefix": &types.AttributeValueMemberS{Value: userId + "#"},
		},
	}

	result, err := r.client.Query(context.TODO(), input)
	if err != nil {
		return nil, fmt.Errorf("failed to query items from DynamoDB: %w", err)
	}

	inputs := make([]*entity.Input, 0)
	for _, item := range result.Items {
		var input entity.Input
		err := attributevalue.UnmarshalMap(item, &input)
		if err != nil {
			return nil, fmt.Errorf("failed to unmarshal result item: %w", err)
		}
		inputs = append(inputs, &input)
	}

	return inputs, nil
}

func (r *DynamoInputRepository) ListInputsByUserIdAndCategoryId(userId string, categoryId int) ([]*entity.Input, error) {
	gsiName := "DataTypeDataValueIndex"
	partitionKeyName := "InputUserId#CategoryId"

	input := &dynamodb.QueryInput{
		TableName:              aws.String("MyDataModel"),
		IndexName:              aws.String(gsiName),
		KeyConditionExpression: aws.String(partitionKeyName + " = :partitionKey"),
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":partitionKey": &types.AttributeValueMemberS{
				Value: userId + "#" + strconv.Itoa(categoryId),
			},
		},
	}

	result, err := r.client.Query(context.TODO(), input)
	if err != nil {
		return nil, fmt.Errorf("failed to query items from DynamoDB: %w", err)
	}

	inputs := make([]*entity.Input, 0)
	for _, item := range result.Items {
		var input entity.Input
		err := attributevalue.UnmarshalMap(item, &input)
		if err != nil {
			return nil, fmt.Errorf("failed to unmarshal result item: %w", err)
		}
		inputs = append(inputs, &input)
	}

	return inputs, nil
}

func (r *DynamoInputRepository) GetInputDetail(inputId string) (*entity.Input, error) {
	dataType := "InputUserId#CategoryId"

	result, err := r.client.GetItem(context.TODO(), &dynamodb.GetItemInput{
		TableName: aws.String("MyDataModel"),
		Key: map[string]types.AttributeValue{
			"Id":       &types.AttributeValueMemberS{Value: inputId},
			"DataType": &types.AttributeValueMemberS{Value: dataType},
		},
	})
	if err != nil {
		panic(fmt.Errorf("failed to get item from DynamoDB, %w", err))
	}

	var input entity.Input
	err = attributevalue.UnmarshalMap(result.Item, &input)
	if err != nil {
		panic(fmt.Errorf("failed to unmarshal result item, %w", err))
	}

	return &input, nil
}

func (r *DynamoInputRepository) CreateInput(input *entity.Input) error {
	tableName := "MyDataModel"
	dataType := "InputUserId#CategoryId"
	strCategoryId := strconv.Itoa(input.CategoryId)
	dataValue := input.UserId + "#" + strCategoryId

	nameItem := map[string]types.AttributeValue{
		"Id":              &types.AttributeValueMemberS{Value: input.Id},
		"DataType":        &types.AttributeValueMemberS{Value: dataType},
		"DataValue":       &types.AttributeValueMemberS{Value: dataValue},
		"UserId":          &types.AttributeValueMemberS{Value: input.UserId},
		"ResourceId":      &types.AttributeValueMemberS{Value: input.ResourceId},
		"InputName":       &types.AttributeValueMemberS{Value: input.Name},
		"InputDetail":     &types.AttributeValueMemberS{Value: *input.Detail},
		"InputCategoryId": &types.AttributeValueMemberN{Value: strCategoryId},
		"InputCreatedAt":  &types.AttributeValueMemberS{Value: input.CreatedAt},
		"InputUpdatedAt":  &types.AttributeValueMemberS{Value: input.UpdatedAt},
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
