package order

type Item struct {
	Description string
	Price       float64
}

func (i Item) GetPrice() float64 {
	return i.Price
}

type Order struct {
	Id          string
	Description string
	Items       []Item
}

func (o Order) GetItems() []Item {
	return o.Items
}
