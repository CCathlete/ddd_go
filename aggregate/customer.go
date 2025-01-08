package aggreate

import (
	"ddd-go/entity"
	"ddd-go/valueobject"

	"github.com/google/uuid"
)

// Aggregate = combines multiple entities to a full object.
// Person is the root entity of Customer.

type Customer struct {
	person       *entity.Person
	products     []*entity.Item
	transactions []valueobject.Transaction // Not a pointer since a valueobject doesn't change (immutable).
}

// Factory to create a new customer.
func NewCustomer(name string) (cust Customer, err error) {

	if name == "" {
		err = ErrInvalidNamePerson
	}

	person := &entity.Person{
		Name: name,
		ID:   uuid.New(),
	}

	cust = Customer{
		person:       person,
		products:     make([]*entity.Item, 0),
		transactions: make([]valueobject.Transaction, 0),
	}

	return
}
