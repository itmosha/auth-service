package redis

import (
	"github.com/go-redis/redis"
	"github.com/itmosha/auth-service/internal/config"
)

type RedisClient struct {
	Cache *redis.Client
}

func NewRedisClient(cacheCfg *config.Cache) *RedisClient {
	return &RedisClient{
		redis.NewClient(&redis.Options{
			Addr:     cacheCfg.Host,
			Password: cacheCfg.Pass,
		}),
	}
}
