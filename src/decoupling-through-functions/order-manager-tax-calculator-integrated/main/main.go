package main

import (
	"fmt"

	"github.com/EnricoPicci/go-class/src/decoupling-through-functions/order-manager-tax-calculator-integrated/order"
	"github.com/EnricoPicci/go-class/src/decoupling-through-functions/order-manager-tax-calculator-integrated/orderManager"
)

func main() {
	o := order.Order{Description: "An order", Items: []order.Item{{Description: "An item", Price: 10.0}, {Description: "A second item", Price: 20.0}}}
	finalPrice := orderManager.CalculatePrice(o)
	fmt.Printf(("The final price is %f"), finalPrice)
}
