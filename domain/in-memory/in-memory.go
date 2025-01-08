package inmemory

import (
	aggreate "ddd-go/aggregate"
	"ddd-go/domain/customer"
	"fmt"
	"sync"

	"github.com/google/uuid"
)

// An in memory implementation of the customer repo.

type MemoryRepo struct {
	customers map[uuid.UUID]aggreate.Customer
	sync.Mutex
}

// ---------------------------------------------------------------------
func New() (mr *MemoryRepo) {
	mr = &MemoryRepo{
		customers: make(map[uuid.UUID]aggreate.Customer),
	}

	return
}

// ---------------------------------------------------------------------
func (mr *MemoryRepo) Get(id uuid.UUID) (cst aggreate.Customer, err error) {

	cst, ok := mr.customers[id]
	if !ok {
		// The error is defined in the domain (business logic).
		err = customer.ErrCustomerNotFound
		cst = aggreate.Customer{}
	}

	return
}

// ---------------------------------------------------------------------
func (mr *MemoryRepo) Add(cst aggreate.Customer) (err error) {

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
			customer.ErrFailedToAddCustomer,
		)

		return
	}

	mr.Lock()
	mr.customers[cst.GetID()] = cst
	mr.Unlock()

	return
}

// ---------------------------------------------------------------------
func (mr *MemoryRepo) Update(cst aggreate.Customer) (err error) {

	if mr.customers == nil {
		// We need to lock our repo when doing modifications.
		err = fmt.Errorf(
			"no customers in repo: %w", customer.ErrUpdateCustomer,
		)
	}

	if _, ok := mr.customers[cst.GetID()]; !ok {
		// Customer doesn't exist in repo.
		err = fmt.Errorf(
			"customer doesn't exist: %w",
			customer.ErrUpdateCustomer,
		)

		return
	}

	mr.Lock()
	mr.customers[cst.GetID()] = cst
	mr.Unlock()

	return
}
