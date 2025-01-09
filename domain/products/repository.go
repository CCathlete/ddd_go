package products

import (
	aggreate "ddd-go/aggregate"

	"github.com/google/uuid"
)

type ProductRepo interface {
	GetAll() ([]aggreate.Product, error)
	Get(uuid.UUID) (aggreate.Product, error)
	Add(aggreate.Product) error
	Update(aggreate.Product) error
	Delete(uuid.UUID) error
}
