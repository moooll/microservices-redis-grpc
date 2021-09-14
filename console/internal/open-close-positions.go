package internal

import (
	"errors"

	"github.com/moooll/microservices-redis-grpc/console/internal/models"
)

type Position struct {
	Price models.Price
	Open  bool
}

func ManipulatePositions() error {
	er := make(chan error, 1)
	go func(er chan error) error {
		for {
			companyName, open, err := ScanInput()
			if err != nil {
				er <- err
				return
			}
			if companyName == "" {
				er <- errors.New("companyName is empty!")
				return
			}

			recievedPrice := RecievePrice()
			position := Position{}
			if open {
				position.Open = true
			}

			position.Price.CompanyName = companyName
			position.Price.BuyPrice = recievedPrice
		}
	}(er)
	if err := <- er; 
}
