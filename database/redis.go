package database

import (
	"github.com/go-redis/redis/v8"
	"umarutv/config"
)

var Redis *redis.Client

func init() {
	Redis = redis.NewClient(&redis.Options{
		Addr: config.Redis.Host + ":" + config.Redis.Port,
		DB:   8,
	})
}
