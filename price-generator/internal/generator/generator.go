package generator

import (
	"price-generator/internal/models"

	"github.com/google/uuid"
)

// GeneratePrice generates new price
func GeneratePrice(company string) (price models.Price) {
	var share = float32(154.3)
	ack := share + float32(1.0)
	// todo: generate
	price = models.Price{}
	price.ID = uuid.New()
	price.CompanyName = company
	price.BuyPrice = ack
	return price
}
