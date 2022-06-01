package handlers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	log "github.com/sirupsen/logrus"

	"github.com/gorilla/mux"
	"github.com/jinagamvasubabu/JITScheduler/models"
	"gorm.io/gorm"
)

type handler struct {
	DB *gorm.DB
}

func New(db *gorm.DB) handler {
	return handler{db}
}

func (h *handler) AddBook(w http.ResponseWriter, r *http.Request) {
	// Read to request body
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Fatalln(err)
	}

	var book models.Book
	json.Unmarshal(body, &book)

	// Append to the Books table
	if result := h.DB.Create(&book); result.Error != nil {
		log.Info(result.Error)
	}

	// Send a 201 created response
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode("Created")
}

func (h *handler) GetAllBooks(w http.ResponseWriter, r *http.Request) {
	var books []models.Book

	if result := h.DB.Find(&books); result.Error != nil {
		log.Error(result.Error)
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(books)
}

func (h *handler) GetBook(w http.ResponseWriter, r *http.Request) {
	// Read dynamic id parameter
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	// Find book by Id
	var book models.Book

	if result := h.DB.First(&book, id); result.Error != nil {
		log.Error(result.Error)
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(book)
}
