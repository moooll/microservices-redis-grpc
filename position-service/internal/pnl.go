package internal

func CalculateProfitAndLoss(bid, ask float32) float32 {
	pnl := bid - ask
	return pnl
}
