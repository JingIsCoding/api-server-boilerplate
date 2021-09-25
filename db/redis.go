package db

import (
	"web-server/config"

	"github.com/go-redis/redis/v8"
)

func NewRedis() redis.Cmdable {
	config := config.Get().RedisConfig
	rdb := redis.NewClient(&redis.Options{
		Addr: config.URL,
	})
	return rdb
}
