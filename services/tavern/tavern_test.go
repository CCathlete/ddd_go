package tavern

import (
	"ddd-go/services/billing"
	"ddd-go/services/order"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestOrder(t *testing.T) {

	// Setting up a stub OrderService + parameters for an order.
	os, cstID, prodIDs, cost := order.StubOrderService(t)

	// Setting up a stub BillingService.
	bs := billing.StubBillingService(t)

	// Setting up configurations for the tavern.
	cfgs := []TaverConfiguration{
		WithOrderService(os),
		WithBillingService(bs),
	}

	// Creating the tavern service.
	tv, err := New(cfgs...)
	require.NoError(t, err)

	// -----------------------------------------------------------------
	// Creating an order.
	// -----------------------------------------------------------------
	resultCost, err := tv.Order(cstID, prodIDs)
	require.NoError(t, err)
	require.Equal(t, cost, resultCost)

}
