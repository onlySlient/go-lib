package rds

import (
	"context"
	"slient/conf"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var ctx = context.Background()

func init() {
	Register(conf.PG_DSN.Value("redis://127.0.0.1:6379/0"))
}

func TestGetInstenance(t *testing.T) {
	rds := GetInstenance()
	assert.NotEqual(t, nil, rds)

	_, err := rds.Set(ctx, "test", "result", time.Second).Result()
	assert.Equal(t, nil, err)

	result, err := rds.Get(ctx, "test").Result()
	assert.Equal(t, nil, err)

	t.Log(result)
}
