package main

import (
	"context"
	"sync"

	"github.com/caarlos0/env"
	"github.com/moooll/microservices-redis-grpc/console/internal"
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

	connPositionService, er := grpcConnect(cfg.GRPCPosition)
	if er != nil {
		log.Error("error connecting to grpc ", er.Error())
	}

	defer func() {
		if errr := connPositionService.Close(); errr != nil {
			log.Error("error connecting to grpc ", errr.Error())
		}
	}()

	connPriceService, er := grpcConnect(cfg.GRPCPrice)
	if er != nil {
		log.Error("error connecting to grpc ", er.Error())
	}

	defer func() {
		if errr := connPriceService.Close(); errr != nil {
			log.Error("error connecting to grpc ", errr.Error())
		}
	}()

	ctx := context.Background()
	priceServClient := pb.NewPriceServiceClient(connPriceService)
	c := make(map[string]models.Price)
	r := server.NewPriceReciever(c, &sync.Mutex{})
	go func(){
		for {
			if e := r.GetPrices(ctx, priceServClient); e != nil {
				log.Error("error getting prices! ", e.Error())
			}
		}
	}()

	erchan := make(chan error)
	inchan := make(chan internal.Input)

	positionManager := &internal.PositionOpener{
		Rc: *r,
		Er: erchan,
		Ctx: ctx,
		Input: inchan,
	}

	// input := make(chan internal.Input)
	go func(inchan chan internal.Input) {
		for {
			in := internal.ScanInput()
			if in.Err != nil {
				log.Error("error scanning the input ", in.Err.Error())
			}

			inchan <- in
		}
	}(inchan)

	go func() {
		positionManager.ManagePositions()
		for e := range positionManager.Er {
			log.Error("error scanning the input ", e.Error())
		}
	}() 

	wait := make(chan bool)
	<- wait
}

func grpcConnect(addr string) (*grpc.ClientConn, error) {
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	return conn, nil
}
