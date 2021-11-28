package main

import (
	"fmt"
	"github.com/go-redis/redis/v8"
)

func initRedis() *redis.Client {
	cf := getConfig()
	return redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%v:%v", cf.RedisHost, cf.RedisPort),
		Password: "",
		DB:       0,
	})
}
