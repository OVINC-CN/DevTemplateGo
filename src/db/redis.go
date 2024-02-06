package db

import (
	"context"
	"fmt"
	"github.com/OVINC-CN/DevTemplateGo/src/utils"
	"github.com/redis/go-redis/v9"
)

var Redis *redis.Client

func InitRedisConnection(host, port, password string, db int) {
	Redis = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", host, port),
		Password: password,
		DB:       db,
	})
	status := Redis.Ping(context.Background())
	utils.Logger.Infof("[InitRedisConnectionSuccess] %T %s", Redis, status)
}
