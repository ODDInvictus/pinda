package image

import (
	"context"

	"github.com/redis/go-redis/v9"
)

func Init(rdb *redis.Client) {
	go compressImage(rdb)
}

func compressImage(rdb *redis.Client) {
	ctx := context.Background()

	rdb.Subscribe(ctx, "compress-image")
}