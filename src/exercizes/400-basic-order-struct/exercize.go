// go run ./src/exercizes/400-basic-order-struct
package main

import "fmt"

type Customer struct {
	Name  string
	TaxID string
}
type Item struct {
	Description string
	ID          string
}
type Order struct {
	ID       string
	Customer Customer
	Item     Item
	Quantity int
}

func main() {
	cust := Customer{"Jane Smith", "1234-abc"}
	item := Item{"Gucci black shooes", "Gucci-123"}
	order := Order{
		"xyz", cust, item, 2,
	}
	fmt.Println("Customer name: ", order.Customer.Name)
	fmt.Println("Item description: ", order.Item.Description)
	fmt.Println("Order: ", order)
}
