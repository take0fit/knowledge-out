package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.43

import (
	"book-action/internal/graph/generated"
	"book-action/internal/graph/model"
	"context"
	"fmt"
)

var users = []*model.User{
	{ID: "1", Username: "user1", Email: "user1@example.com"},
	{ID: "2", Username: "user2", Email: "user2@example.com"},
}

var publishAt = "2024"
var books = []*model.Book{
	{ID: "1", Title: "Book One", Author: "Author One", PublishedAt: &publishAt},
	{ID: "2", Title: "Book Two", Author: "Author Two", PublishedAt: &publishAt},
}

// CreateUser - ユーザーの作成
func (r *mutationResolver) CreateUser(ctx context.Context, username string, email string) (*model.User, error) {
	newUser := &model.User{
		ID:       fmt.Sprintf("%d", len(users)+1), // 簡易的なID生成
		Username: username,
		Email:    email,
	}
	users = append(users, newUser)
	return newUser, nil
}

// CreateBook - 書籍の作成
func (r *mutationResolver) CreateBook(ctx context.Context, title string, author string, publishedAt *string) (*model.Book, error) {
	newBook := &model.Book{
		ID:          fmt.Sprintf("%d", len(books)+1), // 簡易的なID生成
		Title:       title,
		Author:      author,
		PublishedAt: publishedAt,
	}
	books = append(books, newBook)
	return newBook, nil
}

// Users - ユーザー一覧の取得
func (r *queryResolver) Users(ctx context.Context) ([]*model.User, error) {
	return users, nil
}

// User - 特定のユーザーの取得
func (r *queryResolver) User(ctx context.Context, id string) (*model.User, error) {
	for _, user := range users {
		if user.ID == id {
			return user, nil
		}
	}
	return nil, fmt.Errorf("user not found")
}

// Books - 書籍一覧の取得
func (r *queryResolver) Books(ctx context.Context) ([]*model.Book, error) {
	return books, nil
}

// Book - 特定の書籍の取得
func (r *queryResolver) Book(ctx context.Context, id string) (*model.Book, error) {
	for _, book := range books {
		if book.ID == id {
			return book, nil
		}
	}
	return nil, fmt.Errorf("book not found")
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }