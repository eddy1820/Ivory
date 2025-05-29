package setting

import (
	"time"
)

type ServerSettings struct {
	RunMode      string
	HttpPort     string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

func (s *ServerSettings) IsDebug() bool {
	return s.RunMode == "debug"
}

type AppSettings struct {
	DefaultPageSize     int
	MaxPageSize         int
	LogSavePath         string
	LogFileName         string
	LogFileExt          string
	UploadSavePath      string
	UploadServerUrl     string
	UploadImageMaxSize  int
	UploadImageAllowExt []string
}

type DatabaseSettings struct {
	DBType      string
	UserName    string
	Password    string
	Host        string
	DBName      string
	TablePrefix string
	Charset     string
	ParseTime   bool
	MaxIdleConn int
	MaxOpenConn int
}

type RedisSettings struct {
	HttpPort string
}

type TokenSettings struct {
	Secret string
	Expire time.Duration
}
