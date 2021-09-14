// Package grpc contains utilities for connecting as grpc-client
package grpc

import (
	"context"
	"io"

	"github.com/google/uuid"
	"github.com/moooll/microservices-redis-grpc/console/internal/models"
	pb "github.com/moooll/microservices-redis-grpc/price-service/protocol"
	log "github.com/sirupsen/logrus"
)

type PriceReciever struct {
	C chan models.Price
}

func NewPriceReciever(c chan models.Price) *PriceReciever {
	return &PriceReciever{
		C: c,
	}
}

func (p *PriceReciever) GetPrices(ctx context.Context, client pb.PriceServiceClient) error {
	stream, err := client.StreamPrice(ctx, &pb.PriceRequest{})
	if err != nil {
		return err
	}

	price := models.Price{}
	for {
		recieved, er := stream.Recv()
		if er == io.EOF {
			break
		}

		if er != nil {
			return er
		}

		id, e := uuid.Parse(recieved.GetId())
		if e != nil {
			return e
		}

		price.ID = id
		price.CompanyName = recieved.GetCompanyName()
		price.BuyPrice = recieved.GetBuyPrice()
		price.SellPrice = recieved.GetSellPrice()
		log.Info(price)

	}
	return nil
}

func RecievePrice() models.Price {
	// read latest price from map
}