package service

import (
	"fmt"
	"github.com/go-redis/redis/v8"
)

func InitRedis() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%v:%v", cfg.RedisHost, cfg.RedisPort),
		Password: "",
		DB:       0,
	})
}
