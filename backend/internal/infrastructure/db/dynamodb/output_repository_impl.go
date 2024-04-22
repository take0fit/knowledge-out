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

type DynamoOutputRepository struct {
	client *dynamodb.Client
}

func NewDynamoOutputRepository(client *dynamodb.Client) repository.OutputRepository {
	return &DynamoOutputRepository{
		client: client,
	}
}

func (r *DynamoOutputRepository) ListOutputsByUserId(userId string) ([]*entity.Output, error) {
	gsiName := "DataTypeDataValueIndex"
	dataType := "OutputUserId#CategoryId"

	input := &dynamodb.QueryInput{
		TableName:              aws.String("MyDataModel"),
		IndexName:              aws.String(gsiName),
		KeyConditionExpression: aws.String("DataType = :dataType and begins_with(DataValue, :userIdPrefix)"),
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":dataType":     &types.AttributeValueMemberS{Value: dataType},
			":userIdPrefix": &types.AttributeValueMemberS{Value: userId + "#"},
		},
	}

	result, err := r.client.Query(context.TODO(), input)
	if err != nil {
		return nil, fmt.Errorf("failed to query items from DynamoDB: %w", err)
	}

	outputs := make([]*entity.Output, 0)
	for _, item := range result.Items {
		var output entity.Output
		err := attributevalue.UnmarshalMap(item, &output)
		if err != nil {
			return nil, fmt.Errorf("failed to unmarshal result item: %w", err)
		}
		outputs = append(outputs, &output)
	}

	return outputs, nil
}

func (r *DynamoOutputRepository) ListOutputsByUserIdAndCategoryId(userId string, categoryId int) ([]*entity.Output, error) {
	gsiName := "DataTypeDataValueIndex"
	partitionKeyName := "OutputUserId#CategoryId"

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

	outputs := make([]*entity.Output, 0)
	for _, item := range result.Items {
		var output entity.Output
		err := attributevalue.UnmarshalMap(item, &output)
		if err != nil {
			return nil, fmt.Errorf("failed to unmarshal result item: %w", err)
		}
		outputs = append(outputs, &output)
	}

	return outputs, nil
}

func (r *DynamoOutputRepository) GetOutputDetail(outputId string) (*entity.Output, error) {
	dataType := "OutputUserId#CategoryId"

	result, err := r.client.GetItem(context.TODO(), &dynamodb.GetItemInput{
		TableName: aws.String("MyDataModel"),
		Key: map[string]types.AttributeValue{
			"Id":       &types.AttributeValueMemberS{Value: outputId},
			"DataType": &types.AttributeValueMemberS{Value: dataType},
		},
	})
	if err != nil {
		panic(fmt.Errorf("failed to get item from DynamoDB, %w", err))
	}

	var output entity.Output
	err = attributevalue.UnmarshalMap(result.Item, &output)
	if err != nil {
		panic(fmt.Errorf("failed to unmarshal result item, %w", err))
	}

	return &output, nil
}

func (r *DynamoOutputRepository) CreateOutput(output *entity.Output) error {
	tableName := "MyDataModel"
	dataType := "OutputUserId#CategoryId"
	strCategoryId := strconv.Itoa(output.CategoryId)
	dataValue := output.UserId + "#" + strCategoryId

	nameItem := map[string]types.AttributeValue{
		"Id":               &types.AttributeValueMemberS{Value: output.Id},
		"DataType":         &types.AttributeValueMemberS{Value: dataType},
		"DataValue":        &types.AttributeValueMemberS{Value: dataValue},
		"UserId":           &types.AttributeValueMemberS{Value: output.UserId},
		"OutputName":       &types.AttributeValueMemberS{Value: output.Name},
		"OutputDetail":     &types.AttributeValueMemberS{Value: *output.Detail},
		"OutputCategoryId": &types.AttributeValueMemberN{Value: strCategoryId},
		"OutputCreatedAt":  &types.AttributeValueMemberS{Value: output.CreatedAt},
		"OutputUpdatedAt":  &types.AttributeValueMemberS{Value: output.UpdatedAt},
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
