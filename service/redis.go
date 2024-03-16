package service

import (
	"fmt"
	"github.com/go-redis/redis/v8"
)

func InitRedis() *redis.Client {
	cf := GetConfig()
	return redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%v:%v", cf.RedisHost, cf.RedisPort),
		Password: "",
		DB:       0,
	})
}
