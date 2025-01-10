package services

import (
	aggreate "ddd-go/aggregate"
	"ddd-go/domain/customers"
	inmemcust "ddd-go/domain/customers/in-memory"
	"ddd-go/domain/products"
	inmemprod "ddd-go/domain/products/in-memory"
	"ddd-go/valueobject"
	"log"
	"time"

	"github.com/google/uuid"
)

var Business aggreate.Customer

func init() {
	Business, _ = aggreate.NewCustomer("Tavern")
}

// A function that applies a configuration to the order service.
// We define all sorts of configuration appliers below.
type OrderConfiguration = func(*OrderService) error

type OrderService struct {
	customers customers.CustomerRepo
	products  products.ProductRepo
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
func WithCustomerRepo(repo customers.CustomerRepo,
) (cfg OrderConfiguration) {
	cfg = func(os *OrderService) (err error) {

		os.customers = repo

		return
	}
	return
}

func WithProductRepo(repo products.ProductRepo,
) (cfg OrderConfiguration) {
	cfg = func(os *OrderService) (err error) {

		os.products = repo

		return
	}
	return
}

// Configuring a memory customer repo to the order service.
func WithMemoryCustomerRepo() (cfg OrderConfiguration) {

	repo := inmemcust.New()
	cfg = WithCustomerRepo(repo)

	return
}

// Getting a slice of products and creating a product repo from it.
// Then, registering the repo to the service.
func WithMemoryProductRepo(prods []aggreate.Product,
) (cfg OrderConfiguration) {

	// Creating the product repo from our products.
	repo := inmemprod.New()
	for _, pd := range prods {
		if err := repo.Add(pd); err != nil {
			cfg = func(os *OrderService) error {
				return err
			}
		}
	}

	// Registering the repo to our service.
	cfg = WithProductRepo(repo)

	return
}

// Getting a customer's ID and a list of desired products,
// and registers a transaction in the customer's data.
func (s *OrderService) CreateOrder(
	cstID uuid.UUID,
	prodIDs []uuid.UUID,
) (err error) {

	var cost float64
	var prods []aggreate.Product

	// Fetching the products by their ids.
	for _, pid := range prodIDs {
		prod, err := s.products.Get(pid)
		if err != nil {
			return err
		}
		prods = append(prods, prod)
		cost += prod.GetPrice()
	}

	// Fetching the customer from the repo.
	cst, err := s.customers.Get(cstID)
	if err != nil {
		return
	}

	// Creating a transaction.
	transaction := valueobject.Transaction{
		Amount:    cost,
		From:      cst.GetID(),
		To:        Business.GetID(),
		CreatedAt: time.Now(),
	}
	log.Println("List of products: ", prods)
	log.Println("Transaction: ", transaction)

	// Registering the list of products to the customer.

	return
}
