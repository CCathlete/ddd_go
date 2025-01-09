package products

import (
	aggreate "ddd-go/aggregate"

	"github.com/google/uuid"
)

type ProductRepo interface {
	GetAll() ([]aggreate.Product, error)
	Get(uuid.UUID) (aggreate.Customer, error)
	Add(aggreate.Product) error
	Update(aggreate.Product) error
	Delete(uuid.UUID) error
}
