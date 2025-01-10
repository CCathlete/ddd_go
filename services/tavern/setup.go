package tavern

import (
	"ddd-go/services/billing"
	"ddd-go/services/order"
)

type OrderService = order.OrderService
type BillingService = billing.BillingService

type Service struct {
	// An order service to take orders.
	OrderService *OrderService

	// Billing service.
	BillingService *BillingService
}

type TaverConfiguration = func(*Service) error

func New(cfgs ...TaverConfiguration) (tv *Service, err error) {
	tv = &Service{}

	// Looping through all configs and applying them.
	for _, cfg := range cfgs {
		err = cfg(tv)
	}

	return
}

func WithOrderService(os *OrderService) (cfg TaverConfiguration) {
	cfg = func(tv *Service) (err error) {
		// -------------------------------------------------------------

		tv.OrderService = os

		// -------------------------------------------------------------
		return
	}
	return
}

func WithBillingService(bs *BillingService) (cfg TaverConfiguration) {
	cfg = func(tv *Service) (err error) {
		// -------------------------------------------------------------

		tv.BillingService = bs

		// -------------------------------------------------------------
		return
	}
	return
}
