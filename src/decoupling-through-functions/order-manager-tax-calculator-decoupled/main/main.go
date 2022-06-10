package main

import (
	"fmt"

	"github.com/EnricoPicci/go-class.git/src/decoupling-through-functions/order-manager-tax-calculator-decoupled/order"
	"github.com/EnricoPicci/go-class.git/src/decoupling-through-functions/order-manager-tax-calculator-decoupled/orderManager"
	"github.com/EnricoPicci/go-class.git/src/decoupling-through-functions/order-manager-tax-calculator-decoupled/taxCalculatorComplex"
	"github.com/EnricoPicci/go-class.git/src/decoupling-through-functions/order-manager-tax-calculator-decoupled/taxCalculatorSimple"
)

func main() {
	// Configure the tax calculator
	orderManager.SetTaxCalculateLogic(taxCalculatorSimple.CalculateTax)
	orderManager.SetTaxCalculateLogic(taxCalculatorComplex.CalculateTax)

	// Create an order and calculate the final price
	o := order.Order{Description: "An order", Items: []order.Item{{Description: "An item", Price: 10.0}, {Description: "A second item", Price: 20.0}}}
	finalPrice := orderManager.CalculatePrice(o)
	fmt.Printf(("The final price is %f"), finalPrice)
}
