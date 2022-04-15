package redisclient

import (
	"context"

	"github.com/go-redis/redis/v8"
)

//NewClient ...
func NewClient(addr, pass string) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: pass,
		DB:       0,
	})
}

//Ping ...
func Ping(ctx context.Context, client *redis.Client) error {
	return client.Ping(ctx).Err()
}
