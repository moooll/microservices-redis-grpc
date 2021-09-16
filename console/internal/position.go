package internal

import (
	"context"
	"errors"
	"fmt"

	rpc "github.com/moooll/microservices-redis-grpc/console/internal/grpc"
	"github.com/moooll/microservices-redis-grpc/console/internal/models"
)

type PositionOpener struct {
	Rc    rpc.PriceReciever
	Ctx   context.Context
	Input chan Input
	Er    chan error
}

func (po *PositionOpener) ManagePositions() {
	var (
		companyName string
		open        bool
		err         error
	)
	for {
		for in := range po.Input {
			companyName = in.CompanyName
			open = in.Open
			err = in.Err
			if err != nil {
				po.Er <- err
				return
			}
			if companyName == "" {
				po.Er <- errors.New("companyName is empty")
				return
			}

			generatedPrice := po.Rc.GetLatestPrice(companyName)
			position := models.Position{}
			if open {
				position.Open = true
			}

			position.Price.CompanyName = companyName
			position.Price.BuyPrice = generatedPrice.BuyPrice
			spread, e := rpc.GetProfitAndLoss(po.Ctx, position)
			if e != nil {
				po.Er <- e
			}

			fmt.Println("spread:", spread.CompanyName, spread.Spread)
		}

		
	}
}