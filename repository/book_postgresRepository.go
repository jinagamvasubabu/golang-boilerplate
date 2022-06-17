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

func PostgresNewBookRepository(ctx context.Context, db *gorm.DB) BookRepository {
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

//MONGODB
/*

//Repository layer
cfg := config.GetConfig()

type bookrepository struct {
	db     *mgo.Database
}
type BookRepository interface {
	AddBook(ctx context.Context, book model.Book) error
	GetAllBooks(ctx context.Context) ([]model.Book, error)
	GetBook(ctx context.Context, id int32) (model.Book, error)
}

func NewBookRepository(ctx context.Context, db *mgo.Database) Repository {
	return repo{
		db:     db,
	}
}

func (b bookRepository) GetBook(ctx context.Context,id int32) (model.Book, error) {
	var book model.Book
	if err := b.db.C(cfg.DB).FindOne(bson.M{"id":id}); err != nil {
		Logger.Errorf("Error while finding books = %s", err.Error())
		return book, err
	}
	if book.Id == 0 {
		return book, errors.New("no book found")
	}
	return book, nil
}

func (b bookRepository) GetAllBooks(ctx context.Context) ([]model.Book, error) {
	booksEntity := []model.BookEntity{}
	if err := b.db.C(cfg.DB).Find(bson.M{}).All(&booksEntity); err != nil {
		Logger.Errorf("Error while finding books = %s", err.Error())
		return booksEntity, err
	}
	books := []models.Book{}
	for _, b := range booksEntity {
		books = append(books, models.Book{
			Id:        b.Id.Hex(),
			Title:     b.Title,
			Author:    b.Author,
			Desc:      b.Desc,
		})
	}
	return books, nil
}

func (b bookRepository) AddBook(ctx context.Context, book model.Book) error {

	if todo.Title == "" {
		return errors.New("title is required")
	}

	bookEntity := models.BookEntity{
		Id:        bson.ObjectIdHex(book.Id),
		Title:     book.Title,
		Author:	   book.Author,
		Desc:      book.Desc
	}

	if err := repo.db.C(cfg.DB).Insert(&bookEntity); err != nil {
		Logger.Errorf("Error while creating book = %s", err.Error())
		return err
	}

	return nil
}

*/
