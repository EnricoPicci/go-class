// In this example the global variable counter is incremented concurrently by many goroutines which run in parallel
// with no synchronization implemented.
// Considering the relative high number of goroutines and iterations there is the high risk that, at the end of the processing,
// the counter will not hold the number expected just considering the number of times the increment operation has been executed
package main

import (
	"flag"
	"fmt"
	"runtime"
	"sync"
)

var globalCounter = 0

func main() {
	oneProc := flag.Bool("oneProc", false, "run only with one processor")
	flag.Parse()

	numberOfGoroutines := 100
	iterations := 1000

	if *oneProc {
		runtime.GOMAXPROCS(1)
	}

	var wg sync.WaitGroup
	wg.Add(numberOfGoroutines)

	for i := 0; i < numberOfGoroutines; i++ {
		go func() {
			for j := 0; j < iterations; j++ {
				// unprotected update of the global variable
				globalCounter++
			}
			wg.Done()
		}()
	}

	wg.Wait()

	fmt.Printf("The value of counter should be equal to the number of times it has been incremented: %d\n", numberOfGoroutines*iterations)
	fmt.Printf("The real value of counter is: %d\n", globalCounter)
}
