package repository

import "book-action/pkg/model"

type BookRepository interface {
	FindBookByID(id string) (*model.Book, error)
	SaveBook(book *model.Book) error
}
