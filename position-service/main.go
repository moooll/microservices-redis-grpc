package main

import (
	"context"
	"net"
	"net/rpc"

	"github.com/caarlos0/env"
	"github.com/moooll/microservices-redis-grpc/position-service/internal/config"
	rpc "github.com/moooll/microservices-redis-grpc/position-service/internal/grpc"
	pb "github.com/moooll/microservices-redis-grpc/price-service/protocol"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

func main() {
	cfg := config.Config{}
	if err := env.Parse(&cfg); err != nil {
		log.Error("error parsing config: ", err.Error())
	}

	client := grpcClientConnect(cfg.GRPCClientAddr)

	go func() {
		if er := launchGRPCServer(cfg.GRPCServerPort, client); er != nil {
			log.Error("error launching gRPC server: ", er.Error())
		}
	}()

	er := rpc.GetPrice(context.Background(), client)
	if er != nil {
		log.Error("error parsing config: ", er.Error())
	}
}

func grpcClientConnect(grpcAddr string) pb.PriceServiceClient {
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

func launchGRPCServer(port string, client pb.PriceServiceClient) error {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		return err
	}

	s := grpc.NewServer()
	server := rpc.NewProfitAndLoss(client)
	pb.RegisterPriceServiceServer(s, server)
	if er := s.Serve(lis); er != nil {
		return er
	}

	return nil
}
