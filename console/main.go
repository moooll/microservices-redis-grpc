package main

import (
	"context"

	"github.com/caarlos0/env"
	"github.com/moooll/microservices-redis-grpc/console/internal/config"
	server "github.com/moooll/microservices-redis-grpc/console/internal/grpc"
	"github.com/moooll/microservices-redis-grpc/console/internal/models"
	pb "github.com/moooll/microservices-redis-grpc/price-service/protocol"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

func main() {
	cfg := config.Config{}
	if err := env.Parse(&cfg); err != nil {
		log.Error("error parsing config ", err.Error())
	}

	conn, er := grpcConnect(cfg.GRPC)
	if er != nil {
		log.Error("error connecting to grpc ", er.Error())
	}

	defer func() {
		if errr := conn.Close(); errr != nil {
			log.Error("error connecting to grpc ", errr.Error())
		}
	}()
	
	client := pb.NewPriceServiceClient(conn)
	c := make(chan models.Price)
	r := server.NewPriceReciever(c)
	ctx := context.Background()
	if e := r.GetPrices(ctx, client); e != nil {
		log.Error("error getting prices ", e.Error())
	}
}

func grpcConnect(addr string) (*grpc.ClientConn, error) {
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	return conn, nil
}
