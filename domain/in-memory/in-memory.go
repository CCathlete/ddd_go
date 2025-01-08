package inmemory

import (
	aggreate "ddd-go/aggregate"
	"sync"

	"github.com/google/uuid"
)

// An in memory implementation of the customer repo.

type MemoryRepo struct {
	customers map[uuid.UUID]aggreate.Customer
	sync.Mutex
}

func New() (mr *MemoryRepo) {
	mr = &MemoryRepo{
		customers: make(map[uuid.UUID]aggreate.Customer),
	}

	return
}

func (mr *MemoryRepo) Get(uuid.UUID) (cst aggreate.Customer, err error) {

	cst = aggreate.Customer{}

	return
}

func (mr *MemoryRepo) Add(aggreate.Customer) (err error) {

	return
}

func (mr *MemoryRepo) Update(aggreate.Customer) (err error) {

	return
}
