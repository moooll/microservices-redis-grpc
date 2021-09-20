package main

import (
	"context"
	"time"

	"github.com/moooll/microservices-redis-grpc/price-generator/internal/config"
	"github.com/moooll/microservices-redis-grpc/price-generator/internal/generator"
	"github.com/moooll/microservices-redis-grpc/price-generator/internal/redis"

	"github.com/caarlos0/env/v6"
	log "github.com/sirupsen/logrus"
)

func main() {
	cfg := config.Config{}
	if err := env.Parse(&cfg); err != nil {
		log.Error("error parsing config: ", err.Error())
	}

	rdb := redis.Connect(cfg.RedisURI)
	redisClient := redis.NewClient(context.Background(), rdb, "prices")
	generate(*redisClient)
	go func() {
		for {
			generate(*redisClient)
		}
	}()
	wait := make(chan bool)
	<-wait
}

func generate(redisClient redis.Client) {
	applePrice := generator.GeneratePrice("apple")
	er := redisClient.Write(applePrice)
	if er != nil {
		log.Error("error writing to redis: ", er.Error())
	}

	time.Sleep(10 * time.Second)
	msPrice := generator.GeneratePrice("microsoft")
	erro := redisClient.Write(msPrice)
	if erro != nil {
		log.Error("error writing to redis: ", erro.Error())
	}

	time.Sleep(10 * time.Second)
	sonyPrice := generator.GeneratePrice("sony")
	e := redisClient.Write(sonyPrice)
	if e != nil {
		log.Error("error writing to redis: ", e.Error())
	}
}
