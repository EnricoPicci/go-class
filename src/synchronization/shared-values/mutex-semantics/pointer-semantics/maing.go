package main

import (
	"fmt"
	"sync"
)

type Counter struct {
	counter int
	mu      sync.Mutex
}

func (c *Counter) Increment() {
	c.mu.Lock()
	c.counter++
	c.mu.Unlock()
}
func (c *Counter) Value() int {
	return c.counter
}

var globalCounter = &Counter{}

func main() {
	numberOfGoroutines := 100
	iterations := 1000

	var wg sync.WaitGroup
	wg.Add(numberOfGoroutines)

	for i := 0; i < numberOfGoroutines; i++ {
		go func() {
			for j := 0; j < iterations; j++ {
				// unprotected update of the global variable
				globalCounter.Increment()
			}
			wg.Done()
		}()
	}

	wg.Wait()

	fmt.Printf("The value of counter should be equal to the number of times it has been incremented: %d\n", numberOfGoroutines*iterations)
	fmt.Printf("The real value of counter is: %d\n", globalCounter.Value())
}
