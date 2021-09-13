package client

import (
	"context"
	"errors"
	"io"
	"log"
	"sync"

	"github.com/moooll/microservices-redis-grpc/position-service/internal/models"
	pb "github.com/moooll/microservices-redis-grpc/price-service/protocol"
)

type PriceServiceClient struct {
	C           chan models.Price
	LatestPrice map[string]models.Price
	Mu          *sync.Mutex
}

func (p *PriceServiceClient) GetPrice(ctx context.Context, client pb.PriceServiceClient) error {
	var recievedPrice models.Price

	stream, err := client.StreamPrice(ctx, &pb.PriceRequest{})
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

		recievedPrice = models.Price{
			ID:          resp.Id,
			CompanyName: resp.CompanyName,
			BuyPrice:    resp.BuyPrice,
			SellPrice:   resp.SellPrice,
		}

		log.Print(recievedPrice)
		p.Mu.Lock()
		p.LatestPrice["latest"] = recievedPrice
		p.Mu.Unlock()
	}
	return nil
}

// GetLatestPrice returns latest price recieved from price-service 
func (p *PriceServiceClient) GetLatestPrice() (models.Price, error) {
	price, ok := p.LatestPrice["latest"]
	if !ok {
		return models.Price{}, errors.New("no prices recieved!")
	}

	return price, nil
}
