package products

import (
	aggreate "ddd-go/aggregate"

	"github.com/google/uuid"
)

type ProductRepo interface {
	Get(uuid.UUID) (aggreate.Customer, error)
	Add(aggreate.Customer) error
	Update(aggreate.Customer) error
}
