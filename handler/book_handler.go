package handler

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"fmt"

	"github.com/gorilla/mux"
	"github.com/jinagamvasubabu/golang-boilerplate/adapters/kafka"
	Logger "github.com/jinagamvasubabu/golang-boilerplate/adapters/logger"
	"github.com/jinagamvasubabu/golang-boilerplate/model"
	"github.com/jinagamvasubabu/golang-boilerplate/model/dto"
	"github.com/jinagamvasubabu/golang-boilerplate/service"
)

type handler struct {
	bookService service.BookService
}

func NewBookHandler(bookService service.BookService) handler {
	return handler{bookService}
}

func (h handler) AddBook(w http.ResponseWriter, r *http.Request) {
	// Read to request body
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		// Send a 400 bad_request response
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(fmt.Sprintf("Request is invalid = %s", err.Error()))
	}

	var book model.Book
	json.Unmarshal(body, &book)

	// Append to the Books table
	if result := h.bookService.AddBook(context.Background(), book); result != nil {
		// Send a 400 bad_request response
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(fmt.Sprintf("Error = %s", err.Error()))
	} else {
		// Send a 201 created response
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode("Created")
		producer, err := kafka.NewSyncProducer()
		if err != nil {
			Logger.Errorf(fmt.Sprintf("Error while creating syncProducer: %s", err.Error()))
			panic(err)
		}
		if err := producer.PublishMessage(context.Background(), dto.Message{Value: "Created"}); err != nil {
			Logger.Errorf(fmt.Sprintf("Error while publishing message: %s", err.Error()))
			panic(err)
		}
	}
}

func (h handler) GetAllBooks(w http.ResponseWriter, r *http.Request) {
	var books []model.Book
	var err error

	if books, err = h.bookService.GetAllBooks(context.Background()); err != nil {
		// Send a 400 bad_request response
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(fmt.Sprintf("Request is invalid = %s", err.Error()))
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(books)
	Logger.Info("Books fetched")
}

func (h handler) GetBook(w http.ResponseWriter, r *http.Request) {
	// Read dynamic id parameter
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	// Find book by Id
	var book model.Book
	var err error

	if book, err = h.bookService.GetBook(context.Background(), int32(id)); err != nil {
		// Send a 400 bad_request response
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(fmt.Sprintf("Request is invalid = %s", err.Error()))
	} else {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(book)
	}
	Logger.Info("Book fetched")
}
