package main

import (
	"context"
	"net"
	"sync"

	"github.com/caarlos0/env"
	rpc "github.com/moooll/microservices-redis-grpc/position-service/communication/client"
	rpcserver "github.com/moooll/microservices-redis-grpc/position-service/communication/server"
	"github.com/moooll/microservices-redis-grpc/position-service/internal/config"
	"github.com/moooll/microservices-redis-grpc/position-service/internal/models"
	pbserver "github.com/moooll/microservices-redis-grpc/position-service/protocol"
	pb "github.com/moooll/microservices-redis-grpc/price-service/protocol"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

func main() {
	cfg := config.Config{}
	if err := env.Parse(&cfg); err != nil {
		log.Error("error parsing config: ", err.Error())
	}

	conn, client := grpcClientConnect(cfg.GRPCClientAddr)

	defer func() {
		if e := conn.Close(); e != nil {
			log.Error("error parsing config: ", e.Error())
		}
	}()
	c := make(chan models.Price)
	latestPrice := make(map[string]models.Price)
	cl := rpc.NewGetPriceService(c, latestPrice, &sync.Mutex{})
	go func() {
		if er := launchGRPCServer(cfg.GRPCServerPort, *cl); er != nil {
			log.Error("error launching gRPC server: ", er.Error())
		}
	}()

	er := cl.GetPrice(context.Background(), client)
	if er != nil {
		log.Error("error parsing config: ", er.Error())
	}

	var wait chan struct{}
	<-wait
}

func grpcClientConnect(grpcAddr string) (*grpc.ClientConn, pb.PriceServiceClient) {
	conn, er := grpc.Dial(grpcAddr, grpc.WithInsecure())
	if er != nil {
		log.Error("error parsing config: ", er.Error())
	}

	return conn, pb.NewPriceServiceClient(conn)
}

func launchGRPCServer(port string, g rpc.GetPriceService) error {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		return err
	}

	s := grpc.NewServer()
	server := rpcserver.NewProfitAndLoss(g)
	pbserver.RegisterProfitAndLossServer(s, server)
	if er := s.Serve(lis); er != nil {
		return er
	}

	return nil
}
