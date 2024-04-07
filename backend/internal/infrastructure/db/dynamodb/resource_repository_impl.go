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

type DynamoResourceRepository struct {
	client *dynamodb.Client
}

func NewDynamoResourceRepository(client *dynamodb.Client) repository.ResourceRepository {
	return &DynamoResourceRepository{
		client: client,
	}
}

func (r *DynamoResourceRepository) GetResourceDetail(resourceId string) (*model.Resource, error) {
	dataType := "resourceUserId"

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

	var resource model.Resource
	err = attributevalue.UnmarshalMap(result.Item, &resource)
	if err != nil {
		panic(fmt.Errorf("failed to unmarshal result item, %w", err))
	}

	return &resource, nil
}

func (r *DynamoResourceRepository) CreateResource(resource *model.Resource) error {
	tableName := "MyDataModel"
	dataType := "resourceUserId"

	nameItem := map[string]types.AttributeValue{
		"Id":                 &types.AttributeValueMemberS{Value: resource.Id},
		"DataType":           &types.AttributeValueMemberS{Value: dataType},
		"DataValue":          &types.AttributeValueMemberS{Value: resource.UserId},
		"UserId":             &types.AttributeValueMemberS{Value: resource.UserId},
		"ResourceName":       &types.AttributeValueMemberS{Value: resource.Name},
		"ResourceDetail":     &types.AttributeValueMemberS{Value: resource.Detail},
		"ResourceCategoryId": &types.AttributeValueMemberN{Value: strconv.Itoa(resource.CategoryId)},
		"ResourceCreatedAt":  &types.AttributeValueMemberS{Value: resource.CreatedAt.String()},
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

func (r *DynamoResourceRepository) ListResourcesByUserId(userId string) ([]*model.Resource, error) {
	gsiName := "DataValueIndex"
	partitionKeyName := "DataType" // または "DataValue"、実際の設計に依存

	input := &dynamodb.QueryInput{
		TableName:              aws.String("MyDataModel"),
		IndexName:              aws.String(gsiName), // GSI名
		KeyConditionExpression: aws.String(partitionKeyName + " BEGINS_WITH :userIdPrefix"),
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":userIdPrefix": &types.AttributeValueMemberS{Value: userId + "#"},
		},
	}

	// Query実行
	result, err := r.client.Query(context.TODO(), input)
	if err != nil {
		return nil, fmt.Errorf("failed to query items from DynamoDB: %w", err)
	}

	resources := make([]*model.Resource, 0)
	for _, item := range result.Items {
		var resource model.Resource
		err := attributevalue.UnmarshalMap(item, &resource)
		if err != nil {
			return nil, fmt.Errorf("failed to unmarshal result item: %w", err)
		}
		resources = append(resources, &resource)
	}

	return resources, nil
}
