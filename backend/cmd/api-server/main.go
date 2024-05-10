package main

import (
	"context"
	"fmt"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/ssm"
	"github.com/joho/godotenv"
	"github.com/rs/cors"
	"github.com/take0fit/knowledge-out/interface/gql"
	"github.com/take0fit/knowledge-out/interface/gql/generated"
	"github.com/take0fit/knowledge-out/interface/rest"
	"github.com/take0fit/knowledge-out/interface/rest/controller"
	"github.com/take0fit/knowledge-out/internal/application/usecase"
	"github.com/take0fit/knowledge-out/internal/domain/repository"
	"github.com/take0fit/knowledge-out/internal/infrastructure/aws"
	dynamoDbInfra "github.com/take0fit/knowledge-out/internal/infrastructure/aws/dynamo_db"
	s3Infra "github.com/take0fit/knowledge-out/internal/infrastructure/aws/s3"
	ssmInfra "github.com/take0fit/knowledge-out/internal/infrastructure/aws/ssm"
	googleInfra "github.com/take0fit/knowledge-out/internal/infrastructure/google"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"log"
	"net/http"
	"os"
	"strconv"
)

const defaultAppPort = 8080

func main() {
	loadEnvironment()
	setupLog()
	ctx := context.Background()
	cfg := aws.NewConfig(ctx)
	dynamoClient := dynamodb.NewFromConfig(*cfg)
	ssmClient := ssm.NewFromConfig(*cfg)
	s3Client := s3.NewFromConfig(*cfg)
	appPort := getAppPort()
	oauthConfig := setupOAuth()

	authRepo, googleRepo, userRepo, resourceRepo, inputRepo, outputRepo, fileStorageRepo, ParameterStoreRepo := setupRepositories(dynamoClient, s3Client, ssmClient, oauthConfig)
	authUsecase, userUsecase, resourceUsecase, inputUsecase, outputUsecase := setupUseCases(authRepo, googleRepo, userRepo, resourceRepo, inputRepo, outputRepo, fileStorageRepo, ParameterStoreRepo)

	authController := controller.NewAuthController(authUsecase, oauthConfig)
	handler := rest.NewHandler(authController)
	http.ListenAndServe(":8080", handler)

	setupGraphqlServer(userUsecase, resourceUsecase, inputUsecase, outputUsecase, appPort)
}

func loadEnvironment() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
}

func getAppPort() int {
	appPortStr := os.Getenv("APP_PORT")
	appPort, err := strconv.Atoi(appPortStr)
	if err != nil {
		return defaultAppPort
	}
	return appPort
}

func setupOAuth() *oauth2.Config {
	return &oauth2.Config{
		ClientID:     os.Getenv("GOOGLE_OAUTH2_CLIENT_ID"),
		ClientSecret: os.Getenv("GOOGLE_OAUTH2_CLIENT_SECRET"),
		RedirectURL:  "http://localhost:8080/v1/api/callback",
		Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email"},
		Endpoint:     google.Endpoint,
	}
}

func setupRepositories(
	dynamoDbClient *dynamodb.Client,
	s3Client *s3.Client,
	ssmClient *ssm.Client,
	oauthConfig *oauth2.Config,
) (
	repository.AuthRepository,
	repository.GoogleRepository,
	repository.UserRepository,
	repository.ResourceRepository,
	repository.InputRepository,
	repository.OutputRepository,
	repository.FileStorageRepository,
	repository.ParameterStoreRepository,
) {
	httpClient := &http.Client{}
	return dynamoDbInfra.NewAuthRepository(oauthConfig, dynamoDbClient),
		googleInfra.NewGoogleRepository(httpClient),
		dynamoDbInfra.NewDynamoUserRepository(dynamoDbClient),
		dynamoDbInfra.NewDynamoResourceRepository(dynamoDbClient),
		dynamoDbInfra.NewDynamoInputRepository(dynamoDbClient),
		dynamoDbInfra.NewDynamoOutputRepository(dynamoDbClient),
		s3Infra.NewS3StorageService(s3Client),
		ssmInfra.NewParameterStoreRepository(ssmClient)
}

func setupUseCases(
	authRepo repository.AuthRepository,
	googleRepo repository.GoogleRepository,
	userRepo repository.UserRepository,
	resourceRepo repository.ResourceRepository,
	inputRepo repository.InputRepository,
	outputRepo repository.OutputRepository,
	fileStorageRepo repository.FileStorageRepository,
	parameterStoreRepo repository.ParameterStoreRepository,
) (
	*usecase.AuthUseCaseInteractor,
	*usecase.UserUseCaseInteractor,
	*usecase.ResourceUseCaseInteractor,
	*usecase.InputUseCaseInteractor,
	*usecase.OutputUseCaseInteractor,
) {
	return usecase.NewAuthUseCaseInteractor(authRepo, googleRepo, parameterStoreRepo),
		usecase.NewUserInteractor(userRepo),
		usecase.NewResourceInteractor(resourceRepo),
		usecase.NewInputInteractor(inputRepo),
		usecase.NewOutputInteractor(outputRepo)
}

func setupGraphqlServer(
	userUsecase *usecase.UserUseCaseInteractor,
	resourceUsecase *usecase.ResourceUseCaseInteractor,
	inputUsecase *usecase.InputUseCaseInteractor,
	outputUsecase *usecase.OutputUseCaseInteractor,
	port int,
) {
	resolver := gql.NewResolver(userUsecase, resourceUsecase, inputUsecase, outputUsecase)
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: resolver}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", cors.AllowAll().Handler(srv))
	log.Printf("Connect to http://localhost:%d/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}

func setupLog() {
	log.SetOutput(os.Stderr)
	// ログにファイル名と行数を含める
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}
