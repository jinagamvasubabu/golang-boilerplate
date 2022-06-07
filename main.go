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
	ctx := context.Background()
	// Handle sigterm and await termChan signal
	termChan := make(chan os.Signal)
	signal.Notify(termChan, syscall.SIGTERM, syscall.SIGINT)
	//Initialize the config
	if err := config.InitConfig(); err != nil {
		log.Error("error while loading the config", fmt.Sprintf("%s:%s", err, err.Error()))
	}
	//Logger
	log.InitLogger()
	//Initialize the Database
	DB, err := db.InitDatabase()
	if err != nil {
		log.Errorf("Error:%s", err.Error())
	}
	//Initiliaze the repository
	bookRepository := repository.NewBookRepository(ctx, DB)
	//Initiliaze the service
	bookService := service.NewBookService(ctx, bookRepository)
	//Initiliaze the handler
	bookHandler := handler.NewBookHandler(bookService)

	router := handler.InitRouter(bookHandler)
	srv := &http.Server{
		Addr:    fmt.Sprintf("%s:%d", config.GetConfig().HOST, config.GetConfig().PORT),
		Handler: router,
	}
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
