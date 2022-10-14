package taxCalculatorComplex

import "github.com/EnricoPicci/go-class/src/decoupling-through-functions/order-manager-tax-calculator-decoupled/order"

// the logic implemented is pretty stupid, but it's not important for the example
func CalculateTax(o order.Order) float64 {
	tax := 0.0
	for _, item := range o.Items {
		if item.Price > 10.0 {
			tax = tax + item.Price*0.2
			break
		}
		tax = tax + item.Price*0.1
	}
	return tax
}
