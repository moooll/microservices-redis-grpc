package redis

import (
	"context"
	"fmt"

	"github.com/moooll/microservices-redis-grpc/price-generator/internal/models"

	redis "github.com/go-redis/redis/v8"
)

// RedisClient describes client for writing to Redis Streams
type RedisClient struct {
	client     *redis.Client
	ctx        context.Context
	streamName string
}

// NewRedisClient returns new redis client
func NewRedisClient(ctx context.Context, client *redis.Client, streamName string) *RedisClient {
	return &RedisClient{
		client:     client,
		ctx:        ctx,
		streamName: streamName,
	}
}

// Connect connects to Redis and returns *redis.Client entity for working with Redis
func Connect(redisURI string) *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr: redisURI,
	})
	return client
}

// Write writes generated price to Redis Streams
func (c *RedisClient) Write(price models.Price) error {
	val := map[string]interface{}{"price": fmt.Sprint(price)}
	err := c.client.XAdd(c.ctx, &redis.XAddArgs{
		Stream: c.streamName,
		ID:     "",
		Values: val,
	}).Err()
	if err != nil {
		return err
	}

	return nil
}
