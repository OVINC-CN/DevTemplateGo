package db

import (
	"context"
	"fmt"
	"github.com/OVINC-CN/DevTemplateGo/src/configs"
	"github.com/OVINC-CN/DevTemplateGo/src/utils"
	"github.com/redis/go-redis/v9"
	"time"
)

var Redis redisClient

type redisClient struct {
	backend *redis.Client
}

func (client *redisClient) buildCacheKey(keyType, key string) string {
	return fmt.Sprintf("%s:%s:%s", configs.Config.RedisPrefix, keyType, key)
}

func (client *redisClient) Ping(ctx context.Context) *redis.StatusCmd {
	return client.backend.Ping(ctx)
}

func (client *redisClient) Set(ctx context.Context, keyType, key string, value interface{}, expiration time.Duration) *redis.StatusCmd {
	cacheKey := client.buildCacheKey(keyType, key)
	return client.backend.Set(ctx, cacheKey, value, expiration)
}

func (client *redisClient) Get(ctx context.Context, keyType, key string) *redis.StringCmd {
	cacheKey := client.buildCacheKey(keyType, key)
	return client.backend.Get(ctx, cacheKey)
}

func (client *redisClient) Del(ctx context.Context, keyType string, keys ...string) *redis.IntCmd {
	var cacheKeys []string
	for _, key := range keys {
		cacheKeys = append(cacheKeys, client.buildCacheKey(keyType, key))
	}
	return client.backend.Del(ctx, cacheKeys...)
}

func InitRedisConnection(host, port, password string, db, maxConnections int) {
	Redis = redisClient{
		backend: redis.NewClient(
			&redis.Options{
				Addr:           fmt.Sprintf("%s:%s", host, port),
				Password:       password,
				DB:             db,
				MaxActiveConns: maxConnections,
			},
		),
	}
	status := Redis.Ping(context.Background())
	utils.Logger.Infof("[InitRedisConnectionSuccess] %T %s", Redis, status)
}
