package service

import (
	"fmt"
	"github.com/redis/go-redis/v9"
)

func InitRedis() *redis.Client {
	cf := GetConfig()
	return redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%v:%v", cf.RedisHost, cf.RedisPort),
		Password: "",
		DB:       0,
	})
}
