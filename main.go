
package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	log "github.com/jinagamvasubabu/JITScheduler/adapters/logger"
	db "github.com/jinagamvasubabu/JITScheduler/adapters/persistence"
	"github.com/jinagamvasubabu/JITScheduler/config"
	"github.com/jinagamvasubabu/JITScheduler/handlers"
)

func main() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		oscall := <-c
		log.Info("oscall", fmt.Sprintf("%s:%+v", "oscall", oscall.String()))
		cancel()
	}()

	if err := run(ctx); err != nil {
		log.Error("failed to serve:")
	}
}

func run(ctx context.Context) (err error) {
	//Initialize the config
	if err := config.InitConfig(); err != nil {
		log.Error("error while loading the config", fmt.Sprintf("%s:%s", err, err.Error()))
		panic(err)
	}
	//Logger
	log.InitLogger()

	//Initialize the Database
	DB := db.InitDatabase()

	//Initiliaze the handlers
	h := handlers.New(DB)
	router := handlers.InitRouter(h)

	srv := &http.Server{
		Addr:    fmt.Sprintf("%s:%d", config.GetConfig().HOST, config.GetConfig().PORT),
		Handler: router,
	}

	log.Info("Server started")
	//Start the server
	if err := srv.ListenAndServe(); err != nil {
		log.Error("error while starting the server", fmt.Sprintf("%s:%s", err, err.Error()))
		panic(err)
	}
	<-ctx.Done()

	log.Info("server stopped")

	ctxShutDown, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer func() {
		cancel()
	}()

	if err = srv.Shutdown(ctxShutDown); err != nil {
		log.Error("server Shutdown Failed")
	}

	log.Info("server exited properly")

	if err == http.ErrServerClosed {
		err = nil
	}

	return
}
