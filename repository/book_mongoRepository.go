//Repository layer
package repository

import (
	"context"
	"errors"

	Logger "github.com/jinagamvasubabu/golang-boilerplate/adapters/logger"
	"github.com/jinagamvasubabu/golang-boilerplate/config"
	"github.com/jinagamvasubabu/golang-boilerplate/model"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type mongoBookRepository struct {
	db  *mgo.Database
	cfg config.Config
}


func NewMongoBookRepository(ctx context.Context, db *mgo.Database, cfg config.Config) BookRepository {
	return mongoBookRepository{
		db:  db,
		cfg: cfg,
	}
}

func (b mongoBookRepository) GetBook(ctx context.Context, id int32) (model.Book, error) {
	var bookEntity model.MongoBookEntity{}
	if err := b.db.C(b.cfg.DB).Find(bson.M{"id": ObjectId(id)}).One(&bookEntity); err != nil {
		Logger.Errorf("Error while finding books = %s", err)
		return nil, err
	}
	book := model.Book{
		Id:     bookEntity.Id.int32(),
		Title:  bookEntity.Title,
		Author: bookEntity.Author,
		Desc:   bookEntity.Desc,
	}

	return book, nil
}

func (b mongoBookRepository) GetAllBooks(ctx context.Context) ([]model.Book, error) {
	booksEntity := []model.MongoBookEntity{}
	if err := b.db.C(b.cfg.DB).Find(bson.M{}).All(&booksEntity); err != nil {
		Logger.Errorf("Error while finding books = %s", err)
		return booksEntity, err
	}
	books := []model.Book{}
	for _, b := range booksEntity {
		books = append(books, model.Book{
			Id:     b.Id.int32(),
			Title:  b.Title,
			Author: b.Author,
			Desc:   b.Desc,
		})
	}
	return books, nil
}

func (b mongoBookRepository) AddBook(ctx context.Context, book model.Book) error {

	if book.Title == "" {
		return errors.New("title is required")
	}

	bookEntity := model.MongoBookEntity{
		Id:     bson.ObjectId(book.Id),
		Title:  book.Title,
		Author: book.Author,
		Desc:   book.Desc,
	}

	if err := b.db.C(b.cfg.DB).Insert(&bookEntity); err != nil {
		Logger.Errorf("Error while creating book = %s", err.Error())
		return err
	}

	return nil
}
