package main

import (
	"context"

	"github.com/caarlos0/env"
	rpc "github.com/moooll/microservices-redis-grpc/position-service/internal/grpc"
	"github.com/moooll/microservices-redis-grpc/position-service/internal/config"
	pb "github.com/moooll/microservices-redis-grpc/price-service/protocol"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

func main() {
	cfg := config.Config{}
	if err := env.Parse(&cfg); err != nil {
		log.Error("error parsing config: ", err.Error())
	}

	client := grpcConnect(cfg.GRPCAddr)
	er := rpc.GetPrice(context.Background(), client)
	if er != nil {
		log.Error("error parsing config: ", er.Error())
	}

	
}

func grpcConnect(grpcAddr string) pb.PriceServiceClient {
	conn, er := grpc.Dial(grpcAddr, grpc.WithInsecure())
	if er != nil {
		log.Error("error parsing config: ", er.Error())
	}

	defer func() {
		if e := conn.Close(); e != nil {
			log.Error("error parsing config: ", e.Error())
		}
	}()

	return pb.NewPriceServiceClient(conn)
}
