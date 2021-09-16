// Package grpc contains utilities for connecting as grpc-client
package grpc

import (
	"context"
	"io"
	"sync"

	"github.com/google/uuid"
	"github.com/moooll/microservices-redis-grpc/console/internal/models"
	pb "github.com/moooll/microservices-redis-grpc/price-service/protocol"
)

// PriceReciever describes a type to store prices so that to get tha latest ones for each company
type PriceReciever struct {
	C  map[string]models.Price
	Mu *sync.Mutex
}

// NewPriceReciever returns new *PriceReciever
func NewPriceReciever(c map[string]models.Price, mu *sync.Mutex) *PriceReciever {
	return &PriceReciever{
		C:  c,
		Mu: mu,
	}
}

// GetPrices recieves stream of prices from PriceService
func (p *PriceReciever) GetPrices(ctx context.Context, client pb.PriceServiceClient) error {
	stream, err := client.StreamPrice(ctx, &pb.PriceRequest{})
	if err != nil {
		return err
	}

	price := models.Price{}
	for {
		recieved, er := stream.Recv()
		if er == io.EOF {
			return er
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
		p.Mu.Lock()
		p.C[price.CompanyName] = price
		p.Mu.Unlock()
	}
	return nil
}

// GetLatestPrice returns latest price from PriceService
func (p *PriceReciever) GetLatestPrice(companyName string) models.Price {
	pr, ok := p.C[companyName]
	if !ok {
		return models.Price{}
	}

	return pr
}
