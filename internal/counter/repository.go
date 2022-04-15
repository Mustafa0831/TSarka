package counter

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
)

type counterRepository struct {
	client *redis.Client
	key    string
}

//Counter ...
type Counter interface {
	Set(ctx context.Context, num uint64) error
	Get(ctx context.Context) (string, error)
}

//NewCounterRepository ...
func NewCounterRepository(client *redis.Client) *counterRepository {
	return &counterRepository{client: client, key: "counter"}
}

func (c *counterRepository) Set(ctx context.Context, num uint64) error {
	if c.key == "" {
		return fmt.Errorf("redisSet key: %s empty", c.key)
	}
	return c.client.Set(ctx, c.key, num, 0).Err()
}

func (c *counterRepository) Get(ctx context.Context) (string, error) {
	if c.key == "" {
		return "", fmt.Errorf("redisGet key: %s empty", c.key)
	}
	val, err := c.client.Get(ctx, c.key).Result()
	if err == redis.Nil {
		return val, fmt.Errorf("redisGet key: %s not exist", c.key)
	}
	return val, nil
}
