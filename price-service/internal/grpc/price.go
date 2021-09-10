// Package grpc contains methods for messaging via grpc
package grpc

import (
	"math/rand"

	"github.com/google/uuid"
	"github.com/moooll/microservices-redis-grpc/price-service/internal/models"
	pb "github.com/moooll/microservices-redis-grpc/price-service/protocol"
)

type PriceStream struct {
	// fromRedis chan models.Price
	pb.UnimplementedPriceServiceServer
}

func GeneratePrice(company string) (price models.Price) {
	var ack = float32(rand.Intn(160-110) + 110)
	price = models.Price{}
	price.ID = uuid.New()
	price.CompanyName = company
	price.BuyPrice = ack
	return price
}

func (p PriceStream) StreamPrice(req *pb.PriceRequest, stream pb.PriceService_StreamPriceServer) error {
	var m []models.Price
	for i := 0; i < 9; i ++ {
		m = append(m, GeneratePrice("apple"))
	}
	for _, price := range m {
		if err := stream.Send(&pb.PriceResponse{
			Id:          price.ID.String(),
			CompanyName: price.CompanyName,
			BuyPrice:    price.BuyPrice,
			SellPrice:   price.SellPrice,
		}); err != nil {
			return err
		}
	}
	return nil
}
