package global

import (
	"gate/internal/infrastructure/setting"
)

var (
	ServerSetting   *setting.ServerSettings
	AppSetting      *setting.AppSettings
	DatabaseSetting *setting.DatabaseSettings
	RedisSetting    *setting.RedisSettings
	TokenSetting    *setting.TokenSettings
)
