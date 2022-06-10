package order

type Item struct {
	Description string
	Price       float64
}

type Order struct {
	Id          string
	Description string
	Items       []Item
}
