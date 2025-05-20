package main

import (
	"gate/api"
	"gate/global"
	"gate/pkg/setting"
	"gate/pkg/setup"
	"github.com/rs/zerolog/log"
)

func main() {
	setupSetting()
	//setupDBEngine()
	//setupRedisEngine()
	runGinServer()
}

func runGinServer() {
	server, err := api.NewServer(global.TokenSetting)
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

func setupDBEngine() {
	var err error
	global.DB, err = setup.NewDBEngine(global.DatabaseSetting)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot start db")
	}
}

func setupRedisEngine() {
	global.Redis = setup.NewRedisEngine(global.RedisSetting)
}
