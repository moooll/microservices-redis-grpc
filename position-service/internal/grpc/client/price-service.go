package client

import (
	"context"
	"io"
	"log"

	"github.com/moooll/microservices-redis-grpc/position-service/internal/models"
	pb "github.com/moooll/microservices-redis-grpc/price-service/protocol"
)

type PriceServiceClient struct {
	C chan models.Price
}

func (p *PriceServiceClient)GetPrice(ctx context.Context, client pb.PriceServiceClient) error {
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
			ID: resp.Id, 
			CompanyName: resp.CompanyName, 
			BuyPrice: resp.BuyPrice, 
			SellPrice: resp.SellPrice,
		}

		log.Print(recievedPrice)
		p.C <- recievedPrice
	}
	return nil
}
