package order

import (
	aggreate "ddd-go/aggregate"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func initProducts(t *testing.T) (prods []aggreate.Product) {

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
	}

	return
}

func TestCreateOrder(t *testing.T) {
	// Loading a product list.
	prods := initProducts(t)
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
	var args struct {
		ids   []uuid.UUID
		cstID uuid.UUID
	}
	args.cstID = cst.GetID()
	for _, pd := range prods {
		args.ids = append(args.ids, pd.GetID())
	}
	// -----------------------------------------------------------------
	// Creating a new order.
	// -----------------------------------------------------------------
	err = os.CreateOrder(args.cstID, args.ids)
	require.NoError(t, err)
}
