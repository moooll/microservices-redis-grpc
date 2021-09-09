// Package grpc contains utilities for connecting as grpc-client
package grpc

import (
	pb "github.com/moooll/microservices-redis-grpc/price-service/protocol"
)

func getPrices(client pb.PriceService_StreamPriceClient)
