package main

import (
	"context"
	"gate/internal/config"
	"gate/internal/infrastructure/database"
	"gate/internal/infrastructure/redis"
	"gate/internal/logger"
	"gate/internal/router"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	cfg, err := config.LoadConfig("internal/config/config.yaml")
	if err != nil {
		logger.Logger.Error().Err(err).Msg("Failed to load config")
		panic("Failed to load config")
	}

	logger.Init(cfg.Server.RunMode)

	db, err := database.NewDBEngine(cfg.Database, cfg.Server)
	if err != nil {
		logger.Logger.Error().Err(err).Msg("cannot start db")
	}

	rdb := redis.NewRedisEngine(cfg.Redis)

	server, err := router.NewServer(db, rdb, cfg.Token, cfg.Server.HttpPort)
	if err != nil {
		logger.Logger.Error().Err(err).Msg("cannot create server")
		return
	}

	go server.Start()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)
	<-stop
	logger.Logger.Info().Msg("Shutting down gracefully...")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		logger.Logger.Error().Err(err).Msg("HTTP shutdown error")
	}
}
