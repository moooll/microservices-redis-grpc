package main

import (
	"context"
	"net"

	"github.com/caarlos0/env/v6"
	"github.com/moooll/microservices-redis-grpc/price-service/internal/config"
	"github.com/moooll/microservices-redis-grpc/price-service/internal/models"
	"github.com/moooll/microservices-redis-grpc/price-service/internal/redis"
	pb "github.com/moooll/microservices-redis-grpc/price-service/protocol"
	log "github.com/sirupsen/logrus"
	rpc "github.com/moooll/microservices-redis-grpc/price-service/internal/grpc"
	"google.golang.org/grpc"
)

type GRPCServer struct {
	pb.UnimplementedPriceServiceServer
} 

func main() {
	cfg := config.Config{}
	if err := env.Parse(&cfg); err != nil {
		log.Error("error parsing config: ", err.Error())
	}

	rdb := redis.Connect(cfg.RedisURI)
	var streams []string
	streams = append(streams, "prices", "$")
	client := redis.NewClient(context.Background(), rdb, streams)
	c := make(chan models.Price)
	go func() {
		if e := grpcConnect(cfg.PortGRPC); e != nil {
			log.Error("error connecting to gRPC: ", e.Error())
		}
	}()
	go func() {
		for {
			er := client.Read(c)
			if er != nil {
				log.Error("error reading from redis streams:", er.Error())
			}
		}
	}()
	wait := make(chan bool)
	<-wait
}

func grpcConnect(port string) error {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		return err
	}

	s := grpc.NewServer()
	server := rpc.PriceStream{}
	pb.RegisterPriceServiceServer(s, server)
	er := s.Serve(lis)
	if er != nil {
		return er
	}

	return nil
}
