package rds

import (
	"context"
	"slient/conf"
	"sync"
	"time"

	"github.com/go-redis/redis/v8"
)

var m sync.Map
var rdb *redis.Client

func GetInstenance(key string, db ...int) *redis.Client {
	m.LoadOrStore(key, redis.NewClient(&redis.Options{
		Addr:     conf.GetValue("redis_addr", "localhost:6379"),
		Password: conf.GetValue("redis_pwd", "123456"),
		DB:       0,
	}))
}

func Set(ctx context.Context, key string, value interface{}, expiration time.Duration) *redis.StatusCmd {
	return rdb.Set(ctx, key, value, expiration)
}

func Get(ctx context.Context, key string) *redis.StringCmd {
	return rdb.Get(ctx, key)
}
