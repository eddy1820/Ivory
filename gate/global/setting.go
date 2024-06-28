package global

import (
	"gate/pkg/setting"
)

var (
	ServerSetting   *setting.ServerSettings
	AppSetting      *setting.AppSettings
	DatabaseSetting *setting.DatabaseSettings
	RedisSetting    *setting.RedisSettings
	TokenSetting    *setting.TokenSettings
)
