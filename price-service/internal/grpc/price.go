// Package grpc contains methods for messaging via gRPC
package grpc

import (
	"github.com/moooll/microservices-redis-grpc/price-generator/models"
	pb "github.com/moooll/microservices-redis-grpc/price-service/protocol"
)

// PriceStream implements pb.PriceServiceServer and contains a channel to receive from redis
type PriceStream struct {
	pb.UnimplementedPriceServiceServer
	fromRedis chan models.Price
}

// NewPriceStream returns new PriceStream
func NewPriceStream(fromRedis chan models.Price) *PriceStream {
	return &PriceStream{
		fromRedis: fromRedis,
	}
}

// StreamPrice implements pb.StreamPrice
func (p PriceStream) StreamPrice(req *pb.PriceRequest, stream pb.PriceService_StreamPriceServer) error {
	for price := range p.fromRedis {
		if err := stream.Send(&pb.PriceResponse{
			Id:          price.ID.String(),
			CompanyName: price.CompanyName,
			Price:       price.Price,
		}); err != nil {
			return err
		}
	}
	return nil
}
