package rds

import (
	"context"
	"testing"
	"time"

	"github.com/magiconair/properties/assert"
)

var ctx = context.Background()

func TestSet(t *testing.T) {
	_, err := Set(ctx, "pwd", "123456", time.Hour).Result()
	t.Log(err)
}

func TestGet(t *testing.T) {
	r, err := Get(ctx, "pwd").Result()
	assert.Equal(t, err, nil)
	t.Log(r)
}
