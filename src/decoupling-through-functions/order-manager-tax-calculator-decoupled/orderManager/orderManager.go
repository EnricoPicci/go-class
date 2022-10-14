package orderManager

import (
	"github.com/EnricoPicci/go-class/src/decoupling-through-functions/order-manager-tax-calculator-decoupled/order"
)

// taxCalculateLogic is a variable which stores a function that calculates the tax for a given order.
var taxCalculateLogic func(order.Order) float64

// SetTaxCalculateLogic sets the function that calculates the tax for a given order.
func SetTaxCalculateLogic(logic func(order.Order) float64) {
	taxCalculateLogic = logic
}

func CalculatePrice(o order.Order) float64 {
	tot := 0.0
	for _, item := range o.Items {
		tot = tot + item.Price
	}
	if taxCalculateLogic == nil {
		panic("taxCalculateLogic is not set")
	}
	tax := taxCalculateLogic(o)
	return tot + tax
}
