package rds

import (
	"slient/misc"
	"sync"

	"github.com/go-redis/redis/v8"
)

const (
	DefaultKey = "redis-key"
)

var m sync.Map

func Register(dsn string, optName ...string) {
	m.LoadOrStore(misc.GetOptName(append(optName, DefaultKey)...), func() *redis.Client {
		opt, err := redis.ParseURL(dsn)
		if err != nil {
			panic(err)
		}
		return redis.NewClient(opt)
	}())
}

func GetInstenance(keys ...string) *redis.Client {
	key := DefaultKey
	for _, k := range keys {
		if k != "" {
			key = k
		}
	}

	val, ok := m.Load(key)
	if !ok {
		return nil
	}

	return val.(*redis.Client)
}
