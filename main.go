package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	log "github.com/jinagamvasubabu/golang-boilerplate/adapters/logger"
	db "github.com/jinagamvasubabu/golang-boilerplate/adapters/persistence"
	"github.com/jinagamvasubabu/golang-boilerplate/config"
	handler "github.com/jinagamvasubabu/golang-boilerplate/handler"
	"github.com/jinagamvasubabu/golang-boilerplate/repository"
	"github.com/jinagamvasubabu/golang-boilerplate/service"
)

func main() {
	defer recoverPanic()
	ctx := context.Background()
	//Initialize the config

	if err := config.InitConfig(); err != nil {
		log.Error("error while loading the config", fmt.Sprintf("%s:%s", err, err.Error()))
	}
	//Logger
	log.InitLogger()

	//get config to check DB Type
	cfg := config.GetConfig()
	var bookRepository repository.BookRepository

	if cfg.DBType == "postgres" {
		DB, err := db.InitPostgresDatabase()
		if err != nil {
			log.Errorf("Error:%s", err.Error())
		}
		bookRepository = repository.NewPostgresBookRepository(ctx, DB)
	} else {
		DB, err := db.InitMongoDatabase()
		if err != nil {
			log.Errorf("Error:%s", err.Error())
		}
		bookRepository = repository.NewMongoBookRepository(ctx, DB, cfg)
	}

	//Initiliaze the  service.
	bookService := service.NewBookService(ctx, bookRepository)

	//Initiliaze the handler
	bookHandler := handler.NewBookHandler(bookService)
	//router
	router := handler.InitRouter(bookHandler)
	srv := &http.Server{
		Addr:    fmt.Sprintf("%s:%d", config.GetConfig().HOST, config.GetConfig().PORT),
		Handler: router,
	}
	// Handle sigterm and await termChan signal
	termChan := make(chan os.Signal)
	signal.Notify(termChan, syscall.SIGTERM, syscall.SIGINT)
	//Graceful shutdown on OS signals (CTRL+C, etc)
	go func() {
		<-termChan // Blocks here until interrupted
		log.Info("SIGTERM received. Shutdown process initiated\n")
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		srv.Shutdown(ctx)
	}()
	log.Info("Server started")
	//Start the server
	if err := srv.ListenAndServe(); err != nil {
		log.Infof("HTTP server interupted, Error - %s:%s", err, err.Error())
	}
	log.Info("Server Stopped")
}

//function to recover a panic
func recoverPanic() {
	if r := recover(); r != nil {
		log.Info("Recovered from panic!!!")
	}
}
