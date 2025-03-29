package main

import (
	"FinanceSystem/internal/config"
	"FinanceSystem/internal/repository"
	"FinanceSystem/internal/service"
	"FinanceSystem/internal/transport"
	"FinanceSystem/internal/transport/handlers"
	"context"
	"github.com/gin-gonic/gin"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {

	cfg := config.MustLoad()

	if cfg.Env == "local" || cfg.Env == "dev" {
		gin.SetMode(gin.DebugMode)
	}
	if cfg.Env == "prod" {
		gin.SetMode(gin.ReleaseMode)
	}

	db, err := repository.NewSqliteDB(cfg.StoragePath)

	if err != nil {
		log.Fatalf("Failed to init db: %s", err.Error())
	}

	repo := repository.NewRepository(db)
	services := service.NewService(repo)
	handler := handlers.NewHandler(services)

	stop := make(chan struct{})
	services.StartPeriodicRecalculation(30*24*time.Hour, stop)

	r := handler.InitRoutes()
	srv := new(transport.Server)

	go func() {
		if err := srv.Run(cfg.Address, cfg.Timeout, cfg.IdleTimeout, r); err != nil {
			log.Fatalf("Error while running http server: %v", err)
		}
	}()
	log.Print("Server started")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	log.Print("Server stopped")

	if err := srv.Shutdown(context.Background()); err != nil {
		log.Printf("Error while shutting down server: %v", err)
	}

	if err := db.Close(); err != nil {
		log.Printf("Error while closing db: %v", err)
	}

}
