package main

import (
	"context"

	"github.com/caarlos0/env/v6"
	"github.com/moooll/microservices-redis-grpc/price-service/internal/config"
	"github.com/moooll/microservices-redis-grpc/price-service/internal/redis"
	log "github.com/sirupsen/logrus"
)

func main() {
	cfg := config.Config{}
	if err := env.Parse(&cfg); err != nil {
		log.Errorln("error parsing config: ", err.Error())
	}

	rdb := redis.Connect(cfg.RedisURI)
	var streams []string
	streams = append(streams, "prices", "$")
	client := redis.NewClient(context.Background(), rdb, streams)
	go func() {
		for {
			er := client.Read()
			if er != nil {
				log.Errorln("error reading from redis streams:", er.Error())
			}
		}
	}()
	wait := make(chan bool)
	<-wait
}
