// Package internal contains files and directories that should not be imported
package internal

// CalculateProfitAndLoss calculates profit and loss for provided arguments
func CalculateProfitAndLoss(bid, ask float32) float32 {
	pnl := bid - ask
	return pnl
}
