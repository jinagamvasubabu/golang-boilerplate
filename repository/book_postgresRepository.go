package repository

import (
	"context"
	"errors"

	Logger "github.com/jinagamvasubabu/golang-boilerplate/adapters/logger"
	"github.com/jinagamvasubabu/golang-boilerplate/model"

	"gorm.io/gorm"
)

//POSTGRES
type postgresBookRepository struct {
	db *gorm.DB
}

func NewPostgresBookRepository(ctx context.Context, db *gorm.DB) BookRepository {
	return postgresBookRepository{
		db: db,
	}
}

func (b postgresBookRepository) AddBook(ctx context.Context, book model.Book) error {
	// Append to the Books table
	if result := b.db.Create(&book); result.Error != nil {
		Logger.Errorf("Error while creating book = %s", result.Error.Error())
		return result.Error
	}
	return nil
}

func (b postgresBookRepository) GetBook(ctx context.Context, id int32) (model.Book, error) {
	// Find book by Id
	var book model.Book

	if result := b.db.First(&book, id); result.Error != nil {
		Logger.Errorf("Error while returning the book = %s", result.Error)
		return book, result.Error
	}
	if book.Id == 0 {
		return book, errors.New("no book found")
	}
	return book, nil
}

func (b postgresBookRepository) GetAllBooks(ctx context.Context) ([]model.Book, error) {
	books := []model.Book{}
	if result := b.db.Find(&books); result.Error != nil {
		Logger.Errorf("Error while finding books = %s", result.Error.Error())
		return books, result.Error
	}
	return books, nil
}
