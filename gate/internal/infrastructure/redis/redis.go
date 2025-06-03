package redis

import (
	"context"
	"gate/internal/config"
	"gate/internal/logger"
	"github.com/redis/go-redis/v9"
	"time"
)

func NewRedisEngine(setting config.RedisConfig) *redis.Client {
	clint := redis.NewClient(&redis.Options{
		Addr:         "redis:" + setting.HttpPort,
		ReadTimeout:  time.Second * 5,
		WriteTimeout: time.Second * 5,
		PoolSize:     10,
	})
	ctx := context.Background()

	pong, err := clint.Ping(ctx).Result()
	if err != nil {
		logger.Logger.Error().Err(err).Msg("Redis connection was refused, pong: " + pong)
	}
	return clint
}
