package aggreate

import (
	"ddd-go/entity"
	"ddd-go/valueobject"

	"github.com/google/uuid"
)

// Aggregate = combines multiple entities to a full object.
// Person is the root entity of Customer.

type Customer struct {
	*entity.Person
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
		Person:       person,
		products:     make([]*entity.Item, 0),
		transactions: make([]valueobject.Transaction, 0),
	}

	return
}

func (c *Customer) GetID() (id uuid.UUID) {
	id = c.ID

	return
}

func (c *Customer) GetName() (name string) {
	name = c.Name

	return
}

func (c *Customer) GetAge() (age int) {
	age = c.Age

	return
}

func (c *Customer) SetID(id uuid.UUID) {
	if c.Person == nil {
		id, _ := uuid.NewUUID()
		c.Person = &entity.Person{
			ID: id,
		}
	}

	c.ID = id
}

func (c *Customer) SetName(name string) {
	if c.Person == nil {
		id, _ := uuid.NewUUID()
		c.Person = &entity.Person{
			ID: id,
		}
	}
	c.Name = name
}

func (c *Customer) SetAge(age int) {
	if c.Person == nil {
		id, _ := uuid.NewUUID()
		c.Person = &entity.Person{
			ID: id,
		}
	}
	c.Age = age
}

func (c *Customer) AddProducts(
	toAdd []*entity.Item,
) {
	c.products = append(c.products, toAdd...)
}
