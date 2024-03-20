package server

import (
	"book-action/interface/gql"
	"book-action/interface/gql/generated"
	"book-action/internal/infrastructure/db/dynamodb"
	"book-action/internal/usecase/user"
	"log"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
)

func main() {
	// DynamoDBUserRepositoryのインスタンスを生成
	userRepo := dynamodb.NewDynamoDBUserRepository()

	// 他の依存関係をセットアップ
	userUsecase := usecase.NewUserInteractor(userRepo)
	resolver := gql.NewResolver(userUsecase)

	// GraphQLサーバーの設定
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: resolver}))

	// サーバーの起動
	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", "8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
