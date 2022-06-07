package service

import (
	"context"

	Logger "github.com/jinagamvasubabu/JITScheduler/adapters/logger"
	"github.com/jinagamvasubabu/JITScheduler/model"
	"github.com/jinagamvasubabu/JITScheduler/repository"
)

type bookService struct {
	bookRepository repository.BookRepository
}

type BookService interface {
	AddBook(ctx context.Context, book model.Book) error
	GetAllBooks(ctx context.Context) ([]model.Book, error)
	GetBook(ctx context.Context, id int32) (model.Book, error)
}

func NewBookService(ctx context.Context, bookRepository repository.BookRepository) BookService {
	return bookService{
		bookRepository: bookRepository,
	}
}

func (b bookService) AddBook(ctx context.Context, book model.Book) error {
	if err := b.bookRepository.AddBook(ctx, book); err != nil {
		Logger.Errorf("Error while creating the book = %s", err.Error())
		return err
	}
	return nil
}

func (b bookService) GetAllBooks(ctx context.Context) ([]model.Book, error) {
	books, err := b.bookRepository.GetAllBooks(ctx)
	if err != nil {
		Logger.Errorf("Error while fetching all the books = %s", err.Error())
		return books, err
	}
	return books, err
}

func (b bookService) GetBook(ctx context.Context, id int32) (model.Book, error) {
	book, err := b.bookRepository.GetBook(ctx, id)
	if err != nil {
		Logger.Errorf("Error while fetching the book = %s", err.Error())
		return book, err
	}
	return book, err
}
