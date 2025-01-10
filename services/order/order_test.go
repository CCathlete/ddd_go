package order

import (
	aggreate "ddd-go/aggregate"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func initProducts(t *testing.T,
) (cost float64, prods []aggreate.Product) {

	// Setting up initial data for our products.
	data := []struct {
		name, description string
		price             float64
	}{
		{
			name:        "Beer",
			description: "Liquid bread.",
			price:       5.5,
		},
		{
			name:        "Milk",
			description: "Yummy baby food.",
			price:       4,
		},
		{
			name:        "Red wine",
			description: "Shatou Le Blabla.",
			price:       8,
		},
		{
			name:        "White wine",
			description: "Fish juice.",
			price:       7.5,
		},
	}

	// Creating a product from each data segment and appending to our product slice.
	for _, segment := range data {
		newProd, err := aggreate.NewProduct(
			segment.name, segment.description, segment.price,
		)
		require.NoError(t, err)
		prods = append(prods, newProd)
		cost += segment.price
	}

	return
}

func StubOrderService(t *testing.T,
) (
	os *OrderService,
	cstID uuid.UUID,
	prodIDs []uuid.UUID,
	cost float64,
) {

	// Loading a product list.
	cost, prods := initProducts(t)
	// -----------------------------------------------------------------
	// Creating a configuration setting for our order service.
	// -----------------------------------------------------------------
	configs := []OrderConfiguration{
		WithMemoryCustomerRepo(),
		WithMemoryProductRepo(prods),
	}
	// -----------------------------------------------------------------
	// Creating the order service.
	// -----------------------------------------------------------------
	os, err := NewOrderService(configs...)
	require.NoError(t, err)
	// -----------------------------------------------------------------
	// Creating a new customer.
	// -----------------------------------------------------------------
	cst, err := aggreate.NewCustomer("KenC")
	// -----------------------------------------------------------------
	// Adding the customer to the customer repo through the service.
	// -----------------------------------------------------------------
	err = os.customers.Add(cst)
	require.NoError(t, err)
	// -----------------------------------------------------------------
	// Setting up parameters for a new order.
	// -----------------------------------------------------------------

	cstID = cst.GetID()
	for _, pd := range prods {
		prodIDs = append(prodIDs, pd.GetID())
	}

	return
}

func TestCreateOrder(t *testing.T) {
	os, cstID, prodIDs, cost := StubOrderService(t)
	// -----------------------------------------------------------------
	// Creating a new order.
	// -----------------------------------------------------------------
	resultCost, err := os.CreateOrder(cstID, prodIDs)
	require.NoError(t, err)
	require.Equal(t, cost, resultCost)
}
