package handler

import (
	"net/http"

	"github.com/gorilla/mux"
)

func InitRouter(h handler) *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/books", h.GetAllBooks).Methods(http.MethodGet)
	router.HandleFunc("/books/{id}", h.GetBook).Methods(http.MethodGet)
	router.HandleFunc("/books", h.AddBook).Methods(http.MethodPost)
	return router
}
