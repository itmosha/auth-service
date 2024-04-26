package storage

import (
	"context"
	"errors"
	"fmt"
	"time"

	common "github.com/itmosha/auth-service/internal/storage"
	"github.com/itmosha/auth-service/pkg/clients/redis"
	_redis "github.com/redis/go-redis/v9"
)

type CacheRedis struct {
	cli *redis.RedisClient
}

func NewCacheRedis(cli *redis.RedisClient) *CacheRedis {
	return &CacheRedis{cli: cli}
}

func (c *CacheRedis) SetRegisterCode(ctx *context.Context, uid, code string) (err error) {
	return c.cli.Set(*ctx, fmt.Sprintf("register:%s", uid), code, time.Minute*30).Err()
}

func (c *CacheRedis) GetRegisterCode(ctx *context.Context, uid string) (code string, err error) {
	code, err = c.cli.Get(*ctx, fmt.Sprintf("register:%s", uid)).Result()
	if errors.Is(err, _redis.Nil) {
		err = common.ErrRegisterCodeNotExist
	}
	return
}

func (c *CacheRedis) DelRegisterCode(ctx *context.Context, uid string) (err error) {
	return c.cli.Del(*ctx, fmt.Sprintf("register:%s", uid)).Err()
}

func (c *CacheRedis) SetLoginCode(ctx *context.Context, uid, code string) (err error) {
	return c.cli.Set(*ctx, fmt.Sprintf("login:%s", uid), code, time.Minute*30).Err()
}

func (c *CacheRedis) GetLoginCode(ctx *context.Context, uid string) (code string, err error) {
	code, err = c.cli.Get(*ctx, fmt.Sprintf("login:%s", uid)).Result()
	if errors.Is(err, _redis.Nil) {
		err = common.ErrRegisterCodeNotExist
	}
	return
}

func (c *CacheRedis) DelLoginCode(ctx *context.Context, uid string) (err error) {
	return c.cli.Del(*ctx, fmt.Sprintf("login:%s", uid)).Err()
}
