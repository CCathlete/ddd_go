package tavern

import "ddd-go/services"

type OrderService = services.OrderService
type BillingService = services.BillingService

type Tavern struct {
	// An order service to take orders.
	OrderService *OrderService

	// Billing service.
	BillingService *BillingService
}

type TaverConfiguration = func(*Tavern) error

func NewTavern(cfgs ...TaverConfiguration) (tv *Tavern, err error) {
	tv = &Tavern{}

	// Looping through all configs and applying them.
	for _, cfg := range cfgs {
		err = cfg(tv)
	}

	return
}

func WithOrderService(os *OrderService) (cfg TaverConfiguration) {
	cfg = func(tv *Tavern) (err error) {
		// -------------------------------------------------------------

		tv.OrderService = os

		// -------------------------------------------------------------
		return
	}
	return
}

func WithBillingService(bs *BillingService) (cfg TaverConfiguration) {
	cfg = func(tv *Tavern) (err error) {
		// -------------------------------------------------------------

		tv.BillingService = bs

		// -------------------------------------------------------------
		return
	}
	return
}
