package config

import (
	"fmt"
	"github.com/spf13/viper"
	"time"
)

type Config struct {
	Server   ServerConfig   `mapstructure:"Server"`
	App      AppConfig      `mapstructure:"App"`
	Database DatabaseConfig `mapstructure:"Database"`
	Redis    RedisConfig    `mapstructure:"Redis"`
	Token    TokenConfig    `mapstructure:"Token"`
}

type ServerConfig struct {
	RunMode         string        `mapstructure:"RunMode"`
	HttpPort        string        `mapstructure:"HttpPort"`
	ReadTimeout     string        `mapstructure:"ReadTimeout"`
	WriteTimeout    string        `mapstructure:"WriteTimeout"`
	ReadTimeoutDur  time.Duration `mapstructure:"-"`
	WriteTimeoutDur time.Duration `mapstructure:"-"`
}

func (s *ServerConfig) IsDebug() bool {
	return s.RunMode == "debug"
}

type AppConfig struct {
	DefaultPageSize       int      `mapstructure:"DefaultPageSize"`
	MaxPageSize           int      `mapstructure:"MaxPageSize"`
	DefaultContextTimeout int      `mapstructure:"DefaultContextTimeout"`
	LogSavePath           string   `mapstructure:"LogSavePath"`
	LogFileName           string   `mapstructure:"LogFileName"`
	LogFileExt            string   `mapstructure:"LogFileExt"`
	UploadSavePath        string   `mapstructure:"UploadSavePath"`
	UploadServerUrl       string   `mapstructure:"UploadServerUrl"`
	UploadImageMaxSize    int      `mapstructure:"UploadImageMaxSize"`
	UploadImageAllowExt   []string `mapstructure:"UploadImageAllowExt"`
}

type DatabaseConfig struct {
	DBType      string `mapstructure:"DBType"`
	UserName    string `mapstructure:"UserName"`
	Password    string `mapstructure:"Password"`
	Host        string `mapstructure:"Host"`
	DBName      string `mapstructure:"DBName"`
	Charset     string `mapstructure:"Charset"`
	ParseTime   string `mapstructure:"ParseTime"`
	MaxIdleConn int    `mapstructure:"MaxIdleConn"`
	MaxOpenConn int    `mapstructure:"MaxOpenConn"`
}

type RedisConfig struct {
	HttpPort string `mapstructure:"HttpPort"`
}

type TokenConfig struct {
	Secret    string        `mapstructure:"Secret"`
	Expire    string        `mapstructure:"Expire"`
	ExpireDur time.Duration `mapstructure:"-"`
}

func LoadConfig(path string) (*Config, error) {
	v := viper.New()
	v.SetConfigFile(path)

	if err := v.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("failed to read config: %w", err)
	}

	var cfg Config
	if err := v.Unmarshal(&cfg); err != nil {
		return nil, fmt.Errorf("failed to unmarshal config: %w", err)
	}

	// parse duration strings
	if dur, err := time.ParseDuration(cfg.Server.ReadTimeout); err == nil {
		cfg.Server.ReadTimeoutDur = dur
	}
	if dur, err := time.ParseDuration(cfg.Server.WriteTimeout); err == nil {
		cfg.Server.WriteTimeoutDur = dur
	}
	if dur, err := time.ParseDuration(cfg.Token.Expire); err == nil {
		cfg.Token.ExpireDur = dur
	}

	return &cfg, nil
}
