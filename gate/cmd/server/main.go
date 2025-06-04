package main

import (
	"gate/internal/config"
	"gate/internal/infrastructure/database"
	"gate/internal/infrastructure/redis"
	"gate/internal/logger"
	"gate/internal/router"
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

	server, err := router.NewServer(db, rdb, cfg.Token)
	if err != nil {
		logger.Logger.Error().Err(err).Msg("cannot create server")
		return
	}
	err = server.Start(cfg.Server.HttpPort)
	if err != nil {
		logger.Logger.Error().Err(err).Msg("cannot start server")
	}
}
