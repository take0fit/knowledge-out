package dynamo_db

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/take0fit/knowledge-out/internal/domain/entity"
	"golang.org/x/oauth2"
	"log"
	"time"
)

type AuthRepository struct {
	OAuthConfig    *oauth2.Config
	DynamoDBClient *dynamodb.Client
}

func NewAuthRepository(oauthConfig *oauth2.Config, dynamoDBClient *dynamodb.Client) *AuthRepository {
	return &AuthRepository{
		OAuthConfig:    oauthConfig,
		DynamoDBClient: dynamoDBClient,
	}
}

func (r *AuthRepository) ExchangeToken(code string) (*oauth2.Token, error) {
	return r.OAuthConfig.Exchange(context.Background(), code)
}

func (r *AuthRepository) SaveAuthenticationData(token *oauth2.Token, user *entity.User) error {
	log.Print("SaveAuthenticationData")
	var birthday, gender string

	if user.Birthday.Valid {
		birthday = *user.Birthday.String()
	} else {
		birthday = "" // または適切なデフォルト値
	}

	if user.Gender.Valid {
		gender = *user.Gender.String()
	} else {
		gender = "" // または適切なデフォルト値
	}

	dataType := "User#CreatedAt"
	_, err := r.DynamoDBClient.PutItem(context.TODO(), &dynamodb.PutItemInput{
		TableName: aws.String("MyDataModel"),
		Item: map[string]types.AttributeValue{
			"Id":           &types.AttributeValueMemberS{Value: user.Id},
			"DataType":     &types.AttributeValueMemberS{Value: dataType},
			"DataValue":    &types.AttributeValueMemberS{Value: user.CreatedAt},
			"Nickname":     &types.AttributeValueMemberS{Value: user.Nickname.String()},
			"lastName":     &types.AttributeValueMemberS{Value: user.LastName},
			"FirstName":    &types.AttributeValueMemberS{Value: user.FirstName},
			"Birthday":     &types.AttributeValueMemberS{Value: birthday},
			"Gender":       &types.AttributeValueMemberS{Value: gender},
			"CreatedAt":    &types.AttributeValueMemberS{Value: user.CreatedAt},
			"UpdatedAt":    &types.AttributeValueMemberS{Value: user.UpdatedAt},
			"GoogleUserID": &types.AttributeValueMemberS{Value: user.GoogleUserId},
			"AccessToken":  &types.AttributeValueMemberS{Value: token.AccessToken},
			"RefreshToken": &types.AttributeValueMemberS{Value: token.RefreshToken},
			"TokenType":    &types.AttributeValueMemberS{Value: token.TokenType},
			"Expiry":       &types.AttributeValueMemberS{Value: token.Expiry.Format(time.RFC3339)},
		},
	})

	if err != nil {
		log.Printf("failed to put item into DynamoDB: %v", err)
		return fmt.Errorf("failed to put item into DynamoDB: %w", err)
	}

	return err
}
