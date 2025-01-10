package tavern

import (
	"ddd-go/services/billing"
	"ddd-go/services/order"
	"testing"
)

func TestOrder(t *testing.T) {

	// Setting up a stub OrderService + parameters for an order.
	os, cstID, prodIDs, cost := order.StubOrderService(t)

	// Setting up a stub BillingService.
	bs := billing.StubBillingService(t)

}
