package redis

import (
	"context"

	"github.com/go-redis/redis"
	log "github.com/sirupsen/logrus")

type RedisClient struct {
	client *redis.Client
	ctx    context.Context
	stream []string
}

func NewRedisClient(ctx context.Context, client *redis.Client, stream []string) *RedisClient {
	return &RedisClient{
		ctx:    ctx,
		client: client,
		stream: stream,
	}
}

// Connect connects to redis streams and returns client
func Connect(redisURI string) *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr: redisURI,
	})

	return client
}

// func CreateConsumerGroup(c *redis.Client, stream string, start string) error {
// 	err := c.XGroupCreate(stream, "gc-1", "0").Err()
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }

func (client *RedisClient) Read() error {
	a, err := client.client.XRead(&redis.XReadArgs{
		Streams: client.stream,
		
	}).Result()

	if err != nil {
		return err
	}

	for _, v := range a {
		log.Info(v.Messages)
	}

	return nil
}
