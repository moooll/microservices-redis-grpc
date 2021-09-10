package pnl

import "github.com/moooll/microservices-redis-grpc/position-service/internal/models"


func CalculateProfitAndLoss(bid, ask float32) float32 {
	pnl := bid - ask
	return pnl
}
