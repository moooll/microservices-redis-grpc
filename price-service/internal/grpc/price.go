// Package grpc contains methods for messaging via grpc
package grpc

import (
	"github.com/moooll/microservices-redis-grpc/price-service/internal/models"
	pb "github.com/moooll/microservices-redis-grpc/price-service/protocol"
)

type PriceStream struct {
	fromRedis chan models.Price
	pb.UnimplementedPriceServiceServer
}

func NewPriceStream(fromRedis chan models.Price) *PriceStream {
	return &PriceStream{
		fromRedis: fromRedis,
	}
}

func (p PriceStream) StreamPrice(req *pb.PriceRequest, stream pb.PriceService_StreamPriceServer) error {
	// var m []models.Price
	// p := <- 
	// for i := 0; i < 9; i++ {
	// 	m = append(m, GeneratePrice("apple"))
	// }
	for price := range p.fromRedis {
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
