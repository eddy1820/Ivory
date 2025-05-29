package setup

import (
	"context"
	"fmt"
	"gate/internal/infrastructure/setting"
	"github.com/redis/go-redis/v9"
	"log"
	"time"
)

func NewRedisEngine(setting *setting.RedisSettings) *redis.Client {
	clint := redis.NewClient(&redis.Options{
		Addr:         "redis:" + setting.HttpPort,
		ReadTimeout:  time.Second * 5,
		WriteTimeout: time.Second * 5,
		PoolSize:     10,
	})
	ctx := context.Background()

	pong, err := clint.Ping(ctx).Result()
	if err != nil {
		log.Fatalln("!!!!  Redis connection was refused")
	}
	fmt.Println(pong)
	return clint
}
