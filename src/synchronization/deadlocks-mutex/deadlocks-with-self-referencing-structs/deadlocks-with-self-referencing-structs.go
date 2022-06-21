package main

import (
	"fmt"
	"sync"
	"time"
)

// This program simulates a deadlock with self-referencing structs.
// There are 2 user defined types at play here:
// 1) an Order that references a Customer with pointer semantics
// 2) a Customer that references its Orders with pointer semantics
//
// Thera are also 2 functionalities at play
// 1) Issue an Order: When the user issues a order, it changes the status of the order and it adds the receivable to the customer.
// 2) Score a Customer: At the occurence of specific events, not necessarly correlated to the issuing of an order, a the score of
// a customer is calculated and, depending on the result, a customer can be set to VIP status which also means that its last order
// receives a particular discount
//
// The 2 functionalities are run by 2 goroutines.
//
// The 2 functionalities lock the same customer and the same order in a reverse sequence.
// a) the order issue functionality first locks the order (to change the satatus) and then the customer (to add the receivable)
// b) the customer score functionality first locks the customer (to calculate the score) and then the order (to add the discount)
//
// Therefore, this crisscrossing of locks can lead to a deadlock situation.
//
// It has to be noticed that the deadlock is not systematic. The 2 functions called simultaneously (see "launchIssueAnsScore" function)
// normally work without any deadlock. Sometimes though the scheduler schedules the 2 goroutines in a way that the enter the
// crisscrossing of locks and this leads to a deadlock.
// You can appreciate this looking at the "main" function, which calls the "launchIssueAnsScore" function in an infinite loop. At some
// point the loop is interrupted by the deadlock.

type Customer struct {
	mu     sync.Mutex
	Name   string
	Orders []*Order
	// more stuff
	Receivables []Receivable
	Vip         bool
}

func NewCustomer(name string) *Customer {
	return &Customer{Name: name}
}
func (c *Customer) Score(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Score 1 - Try to lock custumer %p\n", c)
	c.mu.Lock()
	fmt.Printf("Score 2  - custumer %p locked\n", c)
	// calculate the score
	// we just assume that the score calculation returns the fact that the customer is a vip
	c.Vip = true
	// find somehow the last order
	o := c.lastOrder()
	o.SetDiscount(0.1)
	fmt.Printf("Score 3 - Unlock costumer %p\n", c)
	c.mu.Unlock()
	fmt.Printf("Score 4  - Custumer %p unlocked\n", c)
}
func (c *Customer) AddReceiveble(o *Order) {
	fmt.Printf("AddReceiveble 1 - Try to lock custumer %p\n", c)
	c.mu.Lock()
	fmt.Printf("AddReceiveble 2  - custumer %p locked\n", c)
	// calculate somehow the receivable
	r := calcReceivable(o)
	c.Receivables = append(c.Receivables, r)
	fmt.Printf("AddReceiveble 3 - Unlock costumer %p\n", c)
	c.mu.Unlock()
	fmt.Printf("AddReceiveble 4  - Custumer %p unlocked\n", c)
}
func (c *Customer) lastOrder() *Order {
	if len(c.Orders) == 0 {
		panic("there should be at least one order for this simulation to work")
	}
	return c.Orders[len(c.Orders)-1]
}

type Order struct {
	mu       sync.Mutex
	Id       string
	Customer *Customer
	// more stuff
	Status   string
	Discount float32
}

func NewOrder(id string, customer *Customer) *Order {
	return &Order{Id: id, Customer: customer}
}
func (o *Order) Issue(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Issue 1 - Try to lock order %p\n", o)
	o.mu.Lock()
	fmt.Printf("Issue 2  - order %p locked\n", o)
	o.Status = "confirmed"
	// do some more  stuff
	o.Customer.AddReceiveble(o)
	fmt.Printf("Issue 3 - Unlock order %p\n", o)
	o.mu.Unlock()
	fmt.Printf("Issue 4  - order %p unlocked\n", o)
}

func (o *Order) SetDiscount(d float32) {
	fmt.Printf("SetDiscount 1 - Try to lock order %p\n", o)
	o.mu.Lock()
	fmt.Printf("SetDiscount 2  - order %p locked\n", o)
	// do some more  stuff
	o.Discount = d
	fmt.Printf("SetDiscount 3 - Unlock order %p\n", o)
	o.mu.Unlock()
	fmt.Printf("SetDiscount 4  - order %p unlocked\n", o)
}

type Receivable struct {
}

func calcReceivable(o *Order) Receivable {
	return Receivable{}
}

func main() {
	c := NewCustomer("Andrea")
	o := NewOrder("2", c)
	o.Status = "in progress"
	c.Orders = append(c.Orders, o)
	counter := 0
	for {
		counter++
		fmt.Println(">>>>>>>>>> Iteration", counter)
		time.Sleep(1 * time.Millisecond)
		launchIssueAnsScore(c, o)
	}
}

func launchIssueAnsScore(c *Customer, o *Order) {
	var wg sync.WaitGroup
	wg.Add(2)
	go o.Issue(&wg)
	go c.Score(&wg)
	wg.Wait()
}
