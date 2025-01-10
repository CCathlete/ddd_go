package order

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCreateOrder(t *testing.T) {
	os, cstID, prodIDs, cost := StubOrderService(t)
	// -----------------------------------------------------------------
	// Creating a new order.
	// -----------------------------------------------------------------
	resultCost, err := os.CreateOrder(cstID, prodIDs)
	require.NoError(t, err)
	require.Equal(t, cost, resultCost)
}
