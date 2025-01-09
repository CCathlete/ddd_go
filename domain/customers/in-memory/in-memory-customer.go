package inmemory

import (
	aggreate "ddd-go/aggregate"
	"ddd-go/domain/customers"
	"fmt"
	"sync"

	"github.com/google/uuid"
)

// An in memory implementation of the customer repo.

type MemoryCustomerRepo struct {
	customers map[uuid.UUID]aggreate.Customer
	sync.Mutex
}

// ---------------------------------------------------------------------
func New() (mr *MemoryCustomerRepo) {
	mr = &MemoryCustomerRepo{
		customers: make(map[uuid.UUID]aggreate.Customer),
	}

	return
}

// ---------------------------------------------------------------------
func (mr *MemoryCustomerRepo) Get(id uuid.UUID) (cst aggreate.Customer, err error) {

	cst, ok := mr.customers[id]
	if !ok {
		// The error is defined in the domain (business logic).
		err = customers.ErrCustomerNotFound
		cst = aggreate.Customer{}
	}

	return
}

// ---------------------------------------------------------------------
func (mr *MemoryCustomerRepo) Add(cst aggreate.Customer) (err error) {

	if mr.customers == nil {
		// We need to lock our repo when doing modifications.
		mr.Lock()
		mr.customers = map[uuid.UUID]aggreate.Customer{}
		mr.Unlock()
	}

	if _, ok := mr.customers[cst.GetID()]; ok {
		// Customer already exists in repo.
		err = fmt.Errorf(
			"customer already exists: %w",
			customers.ErrFailedToAddCustomer,
		)

		return
	}

	mr.Lock()
	mr.customers[cst.GetID()] = cst
	mr.Unlock()

	return
}

// ---------------------------------------------------------------------
func (mr *MemoryCustomerRepo) Update(cst aggreate.Customer) (err error) {

	if mr.customers == nil {
		// We need to lock our repo when doing modifications.
		err = fmt.Errorf(
			"no customers in repo: %w", customers.ErrUpdateCustomer,
		)
	}

	if _, ok := mr.customers[cst.GetID()]; !ok {
		// Customer doesn't exist in repo.
		err = fmt.Errorf(
			"customer doesn't exist: %w",
			customers.ErrUpdateCustomer,
		)

		return
	}

	mr.Lock()
	mr.customers[cst.GetID()] = cst
	mr.Unlock()

	return
}
