package services

import (
	"ddd-go/domain/customer"
	inmemory "ddd-go/domain/in-memory"

	"github.com/google/uuid"
)

// A function that applies a configuration to the order service.
// We define all sorts of configuration appliers below.
type OrderConfiguration = func(*OrderService) error

type OrderService struct {
	customers customer.CustomerRepo
}

func NewOrderService(cfgs ...OrderConfiguration,
) (os *OrderService, err error) {
	os = &OrderService{}

	// Looping through all configs and applying them.
	for _, cfg := range cfgs {
		err = cfg(os)
	}

	return
}

// Configuring a customer repo to the order service.
func WithCustomerRepo(repo customer.CustomerRepo,
) (cfg OrderConfiguration) {
	cfg = func(os *OrderService) (err error) {

		os.customers = repo

		return
	}
	return
}

// Configuring a memory customer repo to the order service.
func WithMemoryCustomerRepo() (cfg OrderConfiguration) {

	repo := inmemory.New()
	cfg = WithCustomerRepo(repo)

	return
}

// Getting a customer's ID and a list of desired products,
// and registers a transaction in the customer's data.
func (s *OrderService) CreateOrder(
	cstID uuid.UUID,
	prodIDs []uuid.UUID,
) (err error) {

	// Fetching the customer from the repo.
	cst, err := s.customers.Get(cstID)
	if err != nil {
		return
	}

	// Registering the list of products to the customer.

	return
}
