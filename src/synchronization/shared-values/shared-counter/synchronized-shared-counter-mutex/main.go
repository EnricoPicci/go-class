// In this example the global variable counter is incremented concurrently by many goroutines which run in parallel
// Synchronization is guaranteed by mutex.
package main

import (
	"fmt"
	"sync"
)

var globalCounter int = 0

func main() {

	numberOfGoroutines := 100
	iterations := 1000

	var wg sync.WaitGroup
	wg.Add(numberOfGoroutines)

	var mu sync.Mutex

	for i := 0; i < numberOfGoroutines; i++ {
		go func() {
			for j := 0; j < iterations; j++ {
				// PROTECTED update of the global variable
				mu.Lock()
				globalCounter++
				mu.Unlock()
			}
			wg.Done()
		}()
	}

	wg.Wait()

	fmt.Printf("The value of counter should be equal to the number of times it has been incremented: %d\n", numberOfGoroutines*iterations)
	fmt.Printf("The real value of counter is: %d\n", globalCounter)
}
