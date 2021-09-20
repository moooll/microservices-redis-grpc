// Package server contains tools for gRPC interacton with ConsoleService as a server
package server

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v4"
	"github.com/moooll/microservices-redis-grpc/position-service/communication/client"
	"github.com/moooll/microservices-redis-grpc/position-service/internal"
	"github.com/moooll/microservices-redis-grpc/position-service/internal/db"
	"github.com/moooll/microservices-redis-grpc/position-service/internal/models"
	pb "github.com/moooll/microservices-redis-grpc/position-service/protocol"
)

// ProfitAndLoss implements pb.ProfitAndLossServer
type ProfitAndLoss struct {
	pb.UnimplementedProfitAndLossServer
	cl       *client.GetPriceService
	conn     *pgx.Conn
	serverID int
}

// NewProfitAndLoss returns new ProfitAndLoss instance
func NewProfitAndLoss(cl client.GetPriceService, conn *pgx.Conn) (p ProfitAndLoss) {
	return ProfitAndLoss{
		cl:   &cl,
		conn: conn,
	}
}

// GetProfitAndLoss implements pb.GetProfitAndLoss method
func (p *ProfitAndLoss) GetProfitAndLoss(ctx context.Context, req *pb.ProfitAndLossRequest) (*pb.ProfitAndLossResponse, error) {
	generatedPrice, err := p.cl.GetLatestPrice(req.CompanyName)
	if err != nil {
		return &pb.ProfitAndLossResponse{}, err
	}

	switch req.Position.String() {
	case "OPEN":
		pnl := getPnLForOpen(req, generatedPrice)
		id, er := uuid.Parse(req.Id)
		if er != nil {
			return &pb.ProfitAndLossResponse{}, er
		}

		db.AddPosition(ctx, p.conn, models.Position{
			ID:            id,
			ServerID:      p.serverID,
			CompanyName:   req.CompanyName,
			SellPrice:     req.SellPrice,
			BuyPrice:      req.BuyPrice,
			ProfitAndLoss: pnl.ProfitAndLoss,
			Open:          true,
		})
		return pnl, nil
	case "CLOSE":
		pnl := getPnLForClosed(req)
		db.UpdPosition(ctx, p.conn, p.serverID, req.SellPrice, pnl.ProfitAndLoss)
		return pnl, nil
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
