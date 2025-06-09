package redis

import (
	"context"

	"user-service/config"

	"github.com/redis/go-redis/v9"
)

var Rdb *redis.Client
var Ctx = context.Background()

func InitRedis() {
	Rdb = redis.NewClient(&redis.Options{
		Addr: config.AppConfig.RedisAddress,
	})
}
