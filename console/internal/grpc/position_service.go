package grpc

import (
	"context"

	"github.com/google/uuid"
	"github.com/moooll/microservices-redis-grpc/console/internal/models"
	server "github.com/moooll/microservices-redis-grpc/position-service/communication/server"
	pb "github.com/moooll/microservices-redis-grpc/position-service/protocol"
)

// GetProfitAndLoss requests ProfitAndLoss from PositionsService  
// and returns spread for the position manipulations
func  GetProfitAndLoss(ctx context.Context, position models.Position) (models.Spread, error) {
	c := server.ProfitAndLoss{}
	req := &pb.ProfitAndLossRequest{
		Id:          position.Price.ID.String(),
		CompanyName: position.Price.CompanyName,
		BuyPrice:    position.Price.BuyPrice,
	}
	if position.Open {
		req.Position = pb.ProfitAndLossRequest_Position(pb.ProfitAndLossRequest_Position_value["OPEN"])
	} else {  
		req.Position = pb.ProfitAndLossRequest_Position(pb.ProfitAndLossRequest_Position_value["CLOSE"])
		req.SellPrice = position.Price.SellPrice
	}

	resp, err := c.GetProfitAndLoss(ctx, req)
	if err != nil {
		return models.Spread{}, err
	}

	id, er := uuid.Parse(resp.Id)
	if er != nil {
		return models.Spread{}, er
	}

	pnl := models.Spread{
		Id: id,
		CompanyName: resp.CompanyName,
		Spread: resp.ProfitAndLoss,
	}

	return pnl, nil
}
