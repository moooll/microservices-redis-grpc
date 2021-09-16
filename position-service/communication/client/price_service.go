package client

import (
	"context"
	"errors"
	"fmt"
	"io"
	"sync"

	"github.com/google/uuid"
	"github.com/moooll/microservices-redis-grpc/position-service/internal/models"
	pbclient "github.com/moooll/microservices-redis-grpc/price-service/protocol"
)

// GetPriceService implements interface
type GetPriceService struct {
	C           chan models.Price
	LatestPrice map[string]models.Price
	Mu          *sync.Mutex
}

//
func NewGetPriceService() *GetPriceService {
	c := make(chan models.Price)
	latestPrice := make(map[string]models.Price)
	mu := &sync.Mutex{}
	return &GetPriceService{
		C:           c,
		LatestPrice: latestPrice,
		Mu:          mu,
	}
}

func (p *GetPriceService) GetPrice(ctx context.Context, client pbclient.PriceServiceClient) error {
	var recievedPrice models.Price

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

		recievedPrice = models.Price{
			ID:          id,
			CompanyName: resp.CompanyName,
			BuyPrice:    resp.BuyPrice,
			SellPrice:   resp.SellPrice,
		}

		p.Mu.Lock()
		p.LatestPrice[resp.CompanyName] = recievedPrice
		p.Mu.Unlock()
		fmt.Println("recieved:", p.LatestPrice)
	}
	return nil
}

// GetLatestPrice returns latest price recieved from price-service
func (p *GetPriceService) GetLatestPrice(companyName string) (models.Price, error) {
	price, ok := p.LatestPrice[companyName]
	if !ok {
		return models.Price{}, errors.New("no prices recieved!")
	}

	return price, nil
}
