package inmemprod

import (
	aggreate "ddd-go/aggregate"
	"ddd-go/domain/products"
	"fmt"
	"sync"

	"github.com/google/uuid"
)

// An in memory implementation of the product repo.

type MemoryProductRepo struct {
	products map[uuid.UUID]aggreate.Product
	sync.Mutex
}

// ---------------------------------------------------------------------
func New() (mr *MemoryProductRepo) {
	mr = &MemoryProductRepo{
		products: make(map[uuid.UUID]aggreate.Product),
	}

	return
}

// ---------------------------------------------------------------------
func (mr *MemoryProductRepo) Get(
	id uuid.UUID,
) (pd aggreate.Product, err error) {

	pd, ok := mr.products[id]
	if !ok {
		// The error is defined in the domain (business logic).
		err = products.ErrProductNotFound
	}

	return
}

// ---------------------------------------------------------------------
func (mr *MemoryProductRepo) Add(pd aggreate.Product) (err error) {

	// Since we're adding a product, we could init the inner storage instead of returning an error.
	if mr.products == nil {
		// We need to lock our repo when doing modifications.
		mr.Lock()
		mr.products = map[uuid.UUID]aggreate.Product{}
		mr.Unlock()
	}

	if _, ok := mr.products[pd.GetID()]; ok {
		// Product already exists in repo.
		err = fmt.Errorf(
			"product already exists: %w",
			products.ErrFailedToAddProduct,
		)

		return
	}

	mr.Lock()
	mr.products[pd.GetID()] = pd
	mr.Unlock()

	return
}

// ---------------------------------------------------------------------
func (mr *MemoryProductRepo) Update(pd aggreate.Product) (err error) {

	if mr.products == nil {
		err = fmt.Errorf(
			"uninitialised inner storage: %w", products.ErrUpdateProduct,
		)
	}

	if _, ok := mr.products[pd.GetID()]; !ok {
		// Product doesn't exist in repo.
		err = fmt.Errorf(
			"product doesn't exist: %w",
			products.ErrUpdateProduct,
		)

		return
	}

	mr.Lock()
	mr.products[pd.GetID()] = pd
	mr.Unlock()

	return
}

// ---------------------------------------------------------------------
func (mr *MemoryProductRepo) GetAll() (
	pds []aggreate.Product, err error,
) {
	// In this case we don't return an error but the interface must include returning an error since other implementation or the repo, like a db might return errors.

	for _, pd := range mr.products {
		pds = append(pds, pd)
	}

	return
}

// ---------------------------------------------------------------------
func (mr *MemoryProductRepo) Delete(id uuid.UUID) (err error) {

	if mr.products == nil {
		// We need to lock our repo when doing modifications.
		err = fmt.Errorf(
			"uninitialised inner storage: %w", products.ErrDeleteProduct,
		)
	}

	if _, ok := mr.products[id]; !ok {
		// Product doesn't exist in repo.
		err = fmt.Errorf(
			"product doesn't exist: %w",
			products.ErrUpdateProduct,
		)

		return
	}

	mr.Lock()
	delete(mr.products, id)
	mr.Unlock()

	return
}
