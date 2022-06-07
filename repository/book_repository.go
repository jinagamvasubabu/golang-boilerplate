package repository

import (
	"context"
	"errors"

	Logger "github.com/jinagamvasubabu/golang-boilerplate/adapters/logger"
	"github.com/jinagamvasubabu/golang-boilerplate/model"
	"gorm.io/gorm"
)

type bookRepository struct {
	db *gorm.DB
}

type BookRepository interface {
	AddBook(ctx context.Context, book model.Book) error
	GetAllBooks(ctx context.Context) ([]model.Book, error)
	GetBook(ctx context.Context, id int32) (model.Book, error)
}

func NewBookRepository(ctx context.Context, db *gorm.DB) BookRepository {
	return bookRepository{
		db: db,
	}
}

func (b bookRepository) AddBook(ctx context.Context, book model.Book) error {
	// Append to the Books table
	if result := b.db.Create(&book); result.Error != nil {
		Logger.Errorf("Error while creating book = %s", result.Error.Error())
		return result.Error
	}
	return nil
}

func (b bookRepository) GetBook(ctx context.Context, id int32) (model.Book, error) {
	// Find book by Id
	var book model.Book

	if result := b.db.First(&book, id); result.Error != nil {
		Logger.Errorf("Error while returning the book = %s", result.Error)
		return book, result.Error
	}
	if book.Id == 0 {
		return book, errors.New("No book found")
	}
	return book, nil
}

func (b bookRepository) GetAllBooks(ctx context.Context) ([]model.Book, error) {
	books := []model.Book{}
	if result := b.db.Find(&books); result.Error != nil {
		Logger.Errorf("Error while finding books = %s", result.Error.Error())
		return books, result.Error
	}
	return books, nil
}
