// Package redis contains tools for connecting and reading from Redis
package redis

import (
	"context"

	"github.com/go-redis/redis"
	"github.com/moooll/microservices-redis-grpc/price-generator/models"
	"github.com/pquerna/ffjson/ffjson"
)

// Client is used to work with unexported Redis client from other packages
type Client struct {
	client *redis.Client
	ctx    context.Context
	stream []string
}

// NewClient connects to Redis and returns client
func NewClient(ctx context.Context, client *redis.Client, stream []string) *Client {
	return &Client{
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

// Read reads messages from Redis Streams
func (client *Client) Read(c chan models.Price) error {
	a, err := client.client.XRead(&redis.XReadArgs{
		Streams: client.stream,
	}).Result()

	if err != nil {
		return err
	}

	for _, v := range a {
		for _, f := range v.Messages {
			var price models.Price
			p := f.Values["price"].(string)
			er := ffjson.Unmarshal([]byte(p), &price)
			if er != nil {
				return er
			}

			c <- price
		}
	}

	return nil
}
