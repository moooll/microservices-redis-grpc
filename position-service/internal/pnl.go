package internal

// CalculateProfitAndLoss calculates profit and loss for provided arguments
func CalculateProfitAndLoss(bid, ask float32) float32 {
	pnl := bid - ask
	return pnl
}
