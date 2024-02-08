package usecase

import (
	"book-action/internal/repository"
	"book-action/pkg/model"
)

type BookUseCase struct {
	bookRepo repository.BookRepository
}

func NewBookUseCase(bookRepo repository.BookRepository) *BookUseCase {
	return &BookUseCase{
		bookRepo: bookRepo,
	}
}

func (uc *BookUseCase) GetBook(id string) (*model.Book, error) {
	return uc.bookRepo.FindBookByID(id)
}

func (uc *BookUseCase) AddBook(book *model.Book) error {
	return uc.bookRepo.SaveBook(book)
}
