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

type DynamoResourceRepository struct {
	client *dynamodb.Client
}

func NewDynamoResourceRepository(client *dynamodb.Client) repository.ResourceRepository {
	return &DynamoResourceRepository{
		client: client,
	}
}

func (r *DynamoResourceRepository) ListResourcesByUserId(userId string) ([]*entity.Resource, error) {
	gsiName := "DataTypeDataValueIndex"
	dataType := "ResourceUserId#CategoryId"

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

	resources := make([]*entity.Resource, 0)
	for _, item := range result.Items {
		var resource entity.Resource
		err := attributevalue.UnmarshalMap(item, &resource)
		if err != nil {
			return nil, fmt.Errorf("failed to unmarshal result item: %w", err)
		}
		resources = append(resources, &resource)
	}

	return resources, nil
}

func (r *DynamoResourceRepository) ListResourcesByUserIdAndCategoryId(userId string, categoryId int) ([]*entity.Resource, error) {
	gsiName := "DataTypeDataValueIndex"
	partitionKeyName := "ResourceUserId#CategoryId"

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

	resources := make([]*entity.Resource, 0)
	for _, item := range result.Items {
		var resource entity.Resource
		err := attributevalue.UnmarshalMap(item, &resource)
		if err != nil {
			return nil, fmt.Errorf("failed to unmarshal result item: %w", err)
		}
		resources = append(resources, &resource)
	}

	return resources, nil
}

func (r *DynamoResourceRepository) GetResourceDetail(resourceId string) (*entity.Resource, error) {
	dataType := "ResourceUserId#CategoryId"

	result, err := r.client.GetItem(context.TODO(), &dynamodb.GetItemInput{
		TableName: aws.String("MyDataModel"),
		Key: map[string]types.AttributeValue{
			"Id":       &types.AttributeValueMemberS{Value: resourceId},
			"DataType": &types.AttributeValueMemberS{Value: dataType},
		},
	})
	if err != nil {
		panic(fmt.Errorf("failed to get item from DynamoDB, %w", err))
	}

	var resource entity.Resource
	err = attributevalue.UnmarshalMap(result.Item, &resource)
	if err != nil {
		panic(fmt.Errorf("failed to unmarshal result item, %w", err))
	}

	return &resource, nil
}

func (r *DynamoResourceRepository) CreateResource(resource *entity.Resource) error {
	tableName := "MyDataModel"
	dataType := "ResourceUserId#CategoryId"
	strCategoryId := strconv.Itoa(resource.CategoryId)
	dataValue := resource.UserId + "#" + strCategoryId

	nameItem := map[string]types.AttributeValue{
		"Id":                 &types.AttributeValueMemberS{Value: resource.Id},
		"DataType":           &types.AttributeValueMemberS{Value: dataType},
		"DataValue":          &types.AttributeValueMemberS{Value: dataValue},
		"UserId":             &types.AttributeValueMemberS{Value: resource.UserId},
		"ResourceName":       &types.AttributeValueMemberS{Value: resource.Name},
		"ResourceDetail":     &types.AttributeValueMemberS{Value: *resource.Detail},
		"ResourceCategoryId": &types.AttributeValueMemberN{Value: strCategoryId},
		"ResourceCreatedAt":  &types.AttributeValueMemberS{Value: resource.CreatedAt},
		"ResourceUpdatedAt":  &types.AttributeValueMemberS{Value: resource.UpdatedAt},
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
