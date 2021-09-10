package server

import (
	"context"

	"github.com/moooll/microservices-redis-grpc/position-service/internal/models"
	"github.com/moooll/microservices-redis-grpc/position-service/internal/pnl"
	pb "github.com/moooll/microservices-redis-grpc/position-service/protocol"
)

type ProfitAndLoss struct {
	C chan models.Price
}

func (p ProfitAndLoss) GetProfitAndLoss(ctx context.Context, req *pb.ProfitAndLossRequest) *pb.ProfitAndLossResponse {
	generatedPrice := <- p.C
	switch pb.ProfitAndLossRequestPosition_name {
	case "OPEN":
		return getPnLForOpen(req, generatedPrice)
	case "CLOSE":
		return getPnLForClosed(req)
	case "UNDEF":
		return &pb.ProfitAndLossResponse{}
	}

	return &pb.ProfitAndLossResponse{}
}

func getPnLForOpen(req *pb.ProfitAndLossRequest, generatedPrice models.Price) *pb.ProfitAndLossResponse {
	pnl := pnl.CalculateProfitAndLoss(req.BuyPrice, generatedPrice.BuyPrice)
	return &pb.ProfitAndLossResponse{
		Id:            req.Id,
		CompanyName:   req.CompanyName,
		ProfitAndLoss: pnl,
	}
}

func getPnLForClosed(req *pb.ProfitAndLossRequest) *pb.ProfitAndLossResponse {
	pnl := pnl.CalculateProfitAndLoss(req.SellPrice, req.BuyPrice)
	return &pb.ProfitAndLossResponse{
		Id:            req.Id,
		CompanyName:   req.CompanyName,
		ProfitAndLoss: pnl,
	}
}
