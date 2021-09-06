package main

import (
	"context"
	"price-generator/internal/config"
	"price-generator/internal/generator"
	"price-generator/internal/redis"

	"github.com/caarlos0/env/v6"
	log "github.com/sirupsen/logrus"
)

func main() {
	cfg := config.Config{}
	if err := env.Parse(&cfg); err != nil {
		log.Errorln("error parsing config: ", err.Error())
	}

	rdb := redis.Connect(cfg.RedisURI)
	redisClient := redis.NewRedisClient(context.Background(), rdb, "prices")
	// todo: remove hardcode
	generatedPrice := generator.GeneratePrice("apple")
	er := redisClient.Write(generatedPrice)
	if er != nil {
		log.Errorln("error writing to redis: ", er.Error())
	}
}
