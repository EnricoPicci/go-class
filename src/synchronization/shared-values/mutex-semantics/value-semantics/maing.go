package main

import (
	"fmt"
	"sync"
)

// counter is a global shared variable
var counter int

// the counter type holds only the mutex and increments the global shared variable
type Counter struct {
	mu sync.Mutex
}

func (c Counter) Increment() {
	c.mu.Lock()
	counter++
	c.mu.Unlock()
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
	fmt.Printf("The real value of counter is: %d\n", counter)
}
