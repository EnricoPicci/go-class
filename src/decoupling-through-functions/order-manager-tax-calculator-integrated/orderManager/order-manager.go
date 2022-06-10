package orderManager

import (
	"github.com/EnricoPicci/go-class.git/src/decoupling-through-functions/order-manager-tax-calculator-integrated/order"
	"github.com/EnricoPicci/go-class.git/src/decoupling-through-functions/order-manager-tax-calculator-integrated/taxCalculator"
)

func CalculatePrice(o order.Order) float64 {
	tot := 0.0
	for _, item := range o.Items {
		tot = tot + item.Price
	}
	tax := taxCalculator.CalculateTax(o)
	return tot + tax
}
