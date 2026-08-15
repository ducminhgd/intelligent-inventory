package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"go.uber.org/zap"
	"moul.io/chizap"

	"github.com/ducminhgd/intelligent-inventory/internal/adapter/postgresql"
	"github.com/ducminhgd/intelligent-inventory/internal/adapter/rest"
	"github.com/ducminhgd/intelligent-inventory/internal/application/manufacturer"
	"github.com/ducminhgd/intelligent-inventory/internal/infrastructure/config"
	"github.com/ducminhgd/intelligent-inventory/internal/infrastructure/db"
	"github.com/ducminhgd/intelligent-inventory/internal/infrastructure/logger"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		panic(err)
	}

	apiLogger, err := logger.New(cfg.LogLevel)
	if err != nil {
		panic(err)
	}

	database, err := db.InitDB(cfg.Database)
	if err != nil {
		apiLogger.Fatal("Failed to initialize database", zap.Error(err))
	}
	defer func() {
		if err := db.CloseDB(database); err != nil {
			apiLogger.Error("Failed to close database", zap.Error(err))
		}
	}()

	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(chizap.New(apiLogger, &chizap.Opts{
		WithReferer:   true,
		WithUserAgent: true,
	}))
	r.Use(middleware.Recoverer)
	r.Use(middleware.URLFormat)

	r.Mount("/health", rest.HealthRouter())

	manufacturerRepo := postgresql.NewManufacturerRepository(database)
	manufacturerSvc := manufacturer.NewManufacturerService(manufacturerRepo)
	manufacturerAPI := rest.NewManufacturerAPI(manufacturerSvc, apiLogger)
	r.Mount("/", manufacturerAPI.Router())

	server := &http.Server{
		Addr:    cfg.REST.Host + ":" + fmt.Sprintf("%d", cfg.REST.Port),
		Handler: r,
	}

	// Server run context
	serverCtx, serverStopFunc := context.WithCancel(context.Background())
	// Listen for syscall signals for process to interrupt/quit
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	go func() {
		<-sig

		// Shutdown signal with grace period of 30 seconds
		shutdownCtx, _ := context.WithTimeout(serverCtx, 30*time.Second)

		go func() {
			<-shutdownCtx.Done()
			if shutdownCtx.Err() == context.DeadlineExceeded {
				apiLogger.Fatal("graceful shutdown timed out.. forcing exit.")
			}
		}()

		// Trigger graceful shutdown
		err := server.Shutdown(shutdownCtx)
		if err != nil {
			apiLogger.Fatal("Failed to shutdown server", zap.Error(err))
		}
		serverStopFunc()
	}()

	// Run the server
	err = server.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		apiLogger.Fatal("Failed to listen and serve", zap.Error(err))
	}

	// Wait for server context to be stopped
	<-serverCtx.Done()
}
