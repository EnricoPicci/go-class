package taxCalculatorSimple

import "github.com/EnricoPicci/go-class.git/src/decoupling-through-functions/order-manager-tax-calculator-decoupled/order"

func CalculateTax(o order.Order) float64 {
	tot := 0.0
	for _, item := range o.Items {
		tot = tot + item.Price
	}
	return tot * 0.2
}
