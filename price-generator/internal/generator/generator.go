// Package generator produces generated prices for the shares
package generator

import (
	"math/rand"

	"github.com/moooll/microservices-redis-grpc/price-generator/internal/models"

	"github.com/google/uuid"
)

// GeneratePrice generates new price
func GeneratePrice(company string) (price models.Price) {
	var ack = float32(rand.Intn(160-110) + 110)
	price = models.Price{}
	price.ID = uuid.New()
	price.CompanyName = company
	price.BuyPrice = ack
	return price
}
