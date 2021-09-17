package models

// Position describes position on the market, which has price and can be requested for opening or closing 
type Position struct {
	Price Price
	Open  bool
}
