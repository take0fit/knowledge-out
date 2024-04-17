package main

import (
	"github.com/take0fit/knowledge-out/interface/gql"
	"github.com/take0fit/knowledge-out/interface/gql/generated"
	"github.com/take0fit/knowledge-out/internal/application/usecase"
	"github.com/take0fit/knowledge-out/internal/infrastructure/db/dynamodb"
	"log"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/rs/cors"
)

func main() {
	// DynamoDBUserRepositoryのインスタンスを生成
	client := dynamodb.NewClient()
	userRepo := dynamodb.NewDynamoUserRepository(client)
	resourceRepo := dynamodb.NewDynamoResourceRepository(client)
	inputRepo := dynamodb.NewDynamoInputRepository(client)
	outputRepo := dynamodb.NewDynamoOutputRepository(client)

	userUsecase := usecase.NewUserInteractor(userRepo)
	resourceUsecase := usecase.NewResourceInteractor(resourceRepo)
	inputUsecase := usecase.NewInputInteractor(inputRepo)
	outputUsecase := usecase.NewOutputInteractor(outputRepo)
	resolver := gql.NewResolver(
		userUsecase,
		resourceUsecase,
		inputUsecase,
		outputUsecase,
	)

	// GraphQLサーバーの設定
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: resolver}))

	// サーバーの起動
	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", cors.AllowAll().Handler(srv))

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", "8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
