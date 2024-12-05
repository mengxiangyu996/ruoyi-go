package dal

import (
	"context"
	"strconv"

	"github.com/go-redis/redis/v8"
)

type RedisConfig struct {
	Host     string
	Port     int
	Database int
	Password string
}

var Redis *redis.Client

func initRedis(config *RedisConfig) {

	Redis = redis.NewClient(&redis.Options{
		Addr:     config.Host + ":" + strconv.Itoa(config.Port),
		Password: config.Password,
		DB:       config.Database,
	})

	_, err := Redis.Ping(context.Background()).Result()
	if err != nil {
		panic(err)
	}
}