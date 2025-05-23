package main

import (
	"gate/internal/infrastructure/global"
	"gate/internal/infrastructure/setting"
	"gate/pkg/setup"
	"gate/router"
	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
)

func main() {
	setupSetting()
	db := setupDBEngine()
	//setupRedisEngine()
	runGinServer(db)
}

func runGinServer(db *gorm.DB) {
	server, err := router.NewServer(db, global.TokenSetting)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot create server")
		return
	}
	err = server.Start(global.ServerSetting.HttpPort)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot start server")
	}
}

func setupSetting() {
	settings, err := setting.NewSetting()
	err = settings.ReadSection("Server", &global.ServerSetting)
	err = settings.ReadSection("App", &global.AppSetting)
	err = settings.ReadSection("Database", &global.DatabaseSetting)
	err = settings.ReadSection("Redis", &global.RedisSetting)
	err = settings.ReadSection("Token", &global.TokenSetting)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot setup setting")
	}
}

func setupDBEngine() *gorm.DB {
	var err error
	db, err := setup.NewDBEngine(global.DatabaseSetting)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot start db")
	}
	return db
}

func setupRedisEngine() {
	global.Redis = setup.NewRedisEngine(global.RedisSetting)
}
