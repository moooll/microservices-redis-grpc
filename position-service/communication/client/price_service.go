// Package client provides tools for gRPC interaction with PriceService as a client
package client

import (
	"context"
	"errors"
	"fmt"
	"io"
	"sync"

	"github.com/google/uuid"
	prmodels "github.com/moooll/microservices-redis-grpc/price-generator/models"
// 	models 	"github.com/moooll/microservices-redis-grpc/position-service/internal/models"

	pbclient "github.com/moooll/microservices-redis-grpc/price-service/protocol"
)

// GetPriceService implements interface
type GetPriceService struct {
	C           chan prmodels.Price
	LatestPrice map[string]prmodels.Price
	Mu          *sync.Mutex
}

// NewGetPriceService returns new GetPriceService
func NewGetPriceService(c chan prmodels.Price, latestPrice map[string]prmodels.Price, mu *sync.Mutex) *GetPriceService {
	return &GetPriceService{
		C:           c,
		LatestPrice: latestPrice,
		Mu:          mu,
	}
}

// GetPrice receives prices from the PriceService via gRPC
func (p *GetPriceService) GetPrice(ctx context.Context, client pbclient.PriceServiceClient) error {
	var recievedPrice prmodels.Price

	stream, err := client.StreamPrice(ctx, &pbclient.PriceRequest{})
	if err != nil {
		return err
	}

	for {
		resp, er := stream.Recv()
		if er == io.EOF {
			break
		}

		if er != nil {
			return er
		}

		id, e := uuid.Parse(resp.Id)
		if e != nil {
			return e
		}

		recievedPrice = prmodels.Price{
			ID:          id,
			CompanyName: resp.CompanyName,
			Price:       resp.Price,
		}

		p.Mu.Lock()
		p.LatestPrice[resp.CompanyName] = recievedPrice
		p.Mu.Unlock()
		fmt.Println("received:", p.LatestPrice)
	}
	return nil
}

// GetLatestPrice returns latest price received from the PriceService
func (p *GetPriceService) GetLatestPrice(companyName string) (prmodels.Price, error) {
	price, ok := p.LatestPrice[companyName]
	if !ok {
		return prmodels.Price{}, errors.New("no prices received")
	}

	return price, nil
}
