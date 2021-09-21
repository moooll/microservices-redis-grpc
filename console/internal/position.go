// Package internal contains tools and types that cannot be imported by other packages
package internal

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	rpc "github.com/moooll/microservices-redis-grpc/console/internal/grpc"
	pb "github.com/moooll/microservices-redis-grpc/position-service/protocol"
)

// PositionManager
type PositionManager struct {
	Rc    *rpc.PriceReciever
	Ctx   context.Context
	Input chan Input
	Er    chan error
	Cl    pb.ProfitAndLossClient
}

// ManagePositions can either open or close a position
func (po *PositionManager) ManagePositions() {
	for {
		for in := range po.Input {
			if in.CompanyName == "" {
				po.Er <- errors.New("companyName is empty")
				return
			}

			// todo: REMOVE THIS SHAME
			time.Sleep(time.Duration(100 * time.Second))
			generatedPrice := po.Rc.GetLatestPrice(in.CompanyName)
			position := &pb.ProfitAndLossRequest{
				CompanyName: in.CompanyName,
			}
			if in.Open {
				id := uuid.New()
				position.Id = id.String()
				position.Position = pb.ProfitAndLossRequest_Position(pb.ProfitAndLossRequest_Position_value["OPEN"])
				position.BuyPrice = generatedPrice.Price
				} else {
				position.Id = in.Id
				position.Position = pb.ProfitAndLossRequest_Position(pb.ProfitAndLossRequest_Position_value["CLOSE"])
				position.SellPrice = generatedPrice.Price
			}

			spread, e := po.Cl.GetProfitAndLoss(po.Ctx, position)
			if e != nil {
				po.Er <- e
			}

			if in.Open {
				fmt.Println("opened:", spread.Id, spread.CompanyName, spread.ProfitAndLoss)
			} else {
				fmt.Println("closed:", spread.Id, spread.CompanyName, spread.ProfitAndLoss)
			}
		}
	}
}
