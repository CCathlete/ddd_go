package aggreate

import (
	"ddd-go/entity"

	"github.com/google/uuid"
)

// We want the underlying item to be private so we define an unexported alias.
type item = entity.Item

type Product struct {
	*item
	price    float64
	quantity int
}

// Factory.
func NewProduct(
	name, description string,
	price float64,
) (p Product, err error) {
	if name == "" || description == "" {
		err = ErrMissingValue
	}

	// Quantity defaults to 0. We set it in another method after creation.
	p = Product{
		item: &entity.Item{
			ID:          uuid.New(),
			Name:        name,
			Description: description,
		},
		price: price,
	}

	return
}

func (p *Product) GetID() (id uuid.UUID) {
	id = p.ID

	return
}

func (p *Product) GetDescription() (des string) {
	des = p.Description

	return
}

func (p *Product) GetName() (name string) {
	name = p.Description

	return
}

func (p *Product) GetPrice() (price float64) {
	price = p.price

	return
}

func (p *Product) GetQuantity() (q int) {
	q = p.quantity

	return
}

func (p *Product) SetID(id uuid.UUID) {
	p.ID = id

	return
}

func (p *Product) SetDescription(des string) {
	p.Description = des

	return
}

func (p *Product) SetName(name string) {
	p.Description = name

	return
}

func (p *Product) SetPrice(price float64) {
	p.price = price

	return
}

func (p *Product) SetQuantity(q int) {
	p.quantity = q

	return
}
