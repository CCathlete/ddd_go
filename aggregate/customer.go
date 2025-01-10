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

func (c *Customer) GetID() (id uuid.UUID) {
	id = c.person.ID

	return
}

func (c *Customer) GetName() (name string) {
	name = c.person.Name

	return
}

func (c *Customer) GetAge() (age int) {
	age = c.person.Age

	return
}

func (c *Customer) SetID(id uuid.UUID) {
	if c.person == nil {
		id, _ := uuid.NewUUID()
		c.person = &entity.Person{
			ID: id,
		}
	}

	c.person.ID = id
}

func (c *Customer) SetName(name string) {
	if c.person == nil {
		id, _ := uuid.NewUUID()
		c.person = &entity.Person{
			ID: id,
		}
	}
	c.person.Name = name
}

func (c *Customer) SetAge(age int) {
	if c.person == nil {
		id, _ := uuid.NewUUID()
		c.person = &entity.Person{
			ID: id,
		}
	}
	c.person.Age = age
}

func (c *Customer) AddProducts(
	toAdd []*entity.Item,
) {
	c.products = append(c.products, toAdd...)
}
