package tavern

import (
	"log"

	"github.com/google/uuid"
)

func (tv *Service) Order(
	cstID uuid.UUID,
	prodIDs []uuid.UUID,
) (cost float64, err error) {

	cost, err = tv.OrderService.CreateOrder(cstID, prodIDs)
	log.Println("Order cost: ", cost)

	return
}
