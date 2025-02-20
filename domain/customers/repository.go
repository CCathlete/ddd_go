package customers

import (
	aggreate "ddd-go/aggregate"

	"github.com/google/uuid"
)

type CustomerRepo interface {
	Get(uuid.UUID) (aggreate.Customer, error)
	Add(aggreate.Customer) error
	Update(aggreate.Customer) error
}
