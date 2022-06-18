// In this example the global variable counter is incremented concurrently by many goroutines which run in parallel
// Synchronization is guaranteed by atomic instructions.
package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var globalCounter int64 = 0

func main() {

	numberOfGoroutines := 100
	iterations := 1000

	var wg sync.WaitGroup
	wg.Add(numberOfGoroutines)

	for i := 0; i < numberOfGoroutines; i++ {
		go func() {
			for j := 0; j < iterations; j++ {
				// PROTECTED update of the global variable
				atomic.AddInt64(&globalCounter, 1)
			}
			wg.Done()
		}()
	}

	wg.Wait()

	fmt.Printf("The value of counter should be equal to the number of times it has been incremented: %d\n", numberOfGoroutines*iterations)
	fmt.Printf("The real value of counter is: %d\n", globalCounter)
}
