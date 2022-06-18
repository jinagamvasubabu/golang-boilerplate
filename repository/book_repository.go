package repository

import (
	"context"

	"github.com/jinagamvasubabu/golang-boilerplate/model"
)

type BookRepository interface {
	AddBook(ctx context.Context, book model.Book) error
	GetAllBooks(ctx context.Context) ([]model.Book, error)
	GetBook(ctx context.Context, id int32) (model.Book, error)
}
