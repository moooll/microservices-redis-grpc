package server

import (
	"context"

	"github.com/moooll/microservices-redis-grpc/position-service/communication/client"
	"github.com/moooll/microservices-redis-grpc/position-service/internal"
	"github.com/moooll/microservices-redis-grpc/position-service/internal/models"
	pb "github.com/moooll/microservices-redis-grpc/position-service/protocol"
)

type ProfitAndLoss struct {
	C chan models.Price
	pb.UnimplementedProfitAndLossServer
	cl client.GetPriceService
}

// NewProfitAndLoss returns new ProfitAndLoss instance
func NewProfitAndLoss(cl client.GetPriceService) (p ProfitAndLoss) {
	return ProfitAndLoss{
		cl: cl,
	}
}

func (p ProfitAndLoss) GetProfitAndLoss(ctx context.Context, req *pb.ProfitAndLossRequest) (*pb.ProfitAndLossResponse, error) {
	generatedPrice, err := p.cl.GetLatestPrice()
	if err != nil {
		return &pb.ProfitAndLossResponse{}, err
	}

	switch req.Position.String() {
	case "OPEN":
		return getPnLForOpen(req, generatedPrice), nil
	case "CLOSE":
		return getPnLForClosed(req), nil
	default:
		return &pb.ProfitAndLossResponse{}, nil
	}
}

func getPnLForOpen(req *pb.ProfitAndLossRequest, generatedPrice models.Price) *pb.ProfitAndLossResponse {
	pnl := internal.CalculateProfitAndLoss(req.BuyPrice, generatedPrice.BuyPrice)
	return &pb.ProfitAndLossResponse{
		Id:            req.Id,
		CompanyName:   req.CompanyName,
		ProfitAndLoss: pnl,
	}
}

func getPnLForClosed(req *pb.ProfitAndLossRequest) *pb.ProfitAndLossResponse {
	pnl := internal.CalculateProfitAndLoss(req.SellPrice, req.BuyPrice)
	return &pb.ProfitAndLossResponse{
		Id:            req.Id,
		CompanyName:   req.CompanyName,
		ProfitAndLoss: pnl,
	}
}
