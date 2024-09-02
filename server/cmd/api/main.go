package main

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/charmingruby/densor/config"
	"github.com/charmingruby/densor/internal/common/api/api_rest"
	"github.com/charmingruby/densor/internal/sensor"
	"github.com/charmingruby/densor/internal/sensor/database/postgres_repository"
	"github.com/charmingruby/densor/pkg/postgres"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	if err := godotenv.Load(); err != nil {
		slog.Warn("CONFIGURATION: .env file not found")
	}

	cfg, err := config.NewConfig()
	if err != nil {
		slog.Error(fmt.Sprintf("CONFIGURATION: %s", err.Error()))
		os.Exit(1)
	}

	db, err := postgres.NewPostgresConnection(cfg)
	if err != nil {
		slog.Error(fmt.Sprintf("DATABASE: %s", err.Error()))
		os.Exit(1)
	}

	router := gin.Default()
	api_rest.SetupCORS(router)

	initDependencies(db, router)

	server := api_rest.NewServer(router, cfg.ServerConfig.Port)

	go func() {
		if err := server.Start(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			slog.Error(fmt.Sprintf("REST SERVER: %s", err.Error()))
			os.Exit(1)
		}
	}()

	// Graceful shutdown
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM)
	<-sc

	slog.Info("HTTP Server interruption received!")

	ctx, shutdownRelease := context.WithTimeout(context.Background(), 10*time.Second)
	defer shutdownRelease()

	if err := server.Server.Shutdown(ctx); err != nil {
		slog.Error(fmt.Sprintf("GRACEFUL SHUTDOWN REST SERVER: %s", err.Error()))
		os.Exit(1)
	}

	slog.Info("Gracefully shutdown!")
}

func initDependencies(db *sqlx.DB, router *gin.Engine) {
	sensorRepo, err := postgres_repository.NewPostgresSensorRepository(db)
	if err != nil {
		slog.Error(fmt.Sprintf("DATABASE REPOSITORY: %s", err.Error()))
		os.Exit(1)
	}

	sectorRepo, err := postgres_repository.NewPostgresSectorRepository(db)
	if err != nil {
		slog.Error(fmt.Sprintf("DATABASE REPOSITORY: %s", err.Error()))
		os.Exit(1)
	}

	sensorCategoryRepo, err := postgres_repository.NewPostgresSensorCategoryRepository(db)
	if err != nil {
		slog.Error(fmt.Sprintf("DATABASE REPOSITORY: %s", err.Error()))
		os.Exit(1)
	}

	equipmentRepo, err := postgres_repository.NewPostgresEquipmentRepository(db)
	if err != nil {
		slog.Error(fmt.Sprintf("DATABASE REPOSITORY: %s", err.Error()))
		os.Exit(1)
	}

	sensorSvc := sensor.NewService(sensorRepo, sensorCategoryRepo, sectorRepo, equipmentRepo)
	sensor.NewHTTPService(router, sensorSvc)
}
