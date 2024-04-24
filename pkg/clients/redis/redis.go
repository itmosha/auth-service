package redis

import (
	"github.com/itmosha/auth-service/internal/config"
	"github.com/redis/go-redis/v9"
)

type RedisClient struct {
	*redis.Client
}

func NewRedisClient(cacheCfg *config.Cache) *RedisClient {
	return &RedisClient{
		redis.NewClient(&redis.Options{
			Addr:     cacheCfg.Host + ":" + cacheCfg.Port,
			Password: cacheCfg.Pass,
		}),
	}
}
