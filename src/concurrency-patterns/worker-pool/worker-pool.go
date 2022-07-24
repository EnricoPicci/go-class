package main

import (
	"fmt"
	"sync"
	"time"
)

// In this example we implement worker pool pattern in its simplest form

type Request struct {
	Param int
}

func main() {
	fmt.Println("Start processing requests")

	// the channel that provides requests to the pool is unbuffered
	inCh := make(chan Request)

	// set the size of the worker pool
	poolSize := 3

	var wg sync.WaitGroup
	wg.Add(poolSize)

	// start the workers
	i := 0
	for i < poolSize {
		go doWork(inCh, &wg, i)
		i++
	}

	// add a delay just to make sure that the processing of the stream starts after all workers have started
	time.Sleep(1 * time.Millisecond)

	// we simulate a stream of incoming requests with an array of Request values
	requests := []Request{{1}, {2}, {3}, {4}, {5}, {6}, {7}, {8}, {9}, {10}}

	for _, req := range requests {
		inCh <- req
		fmt.Printf("Request %v sent\n", req)
	}

	// when the stream completes we close the channel
	close(inCh)

	// This Wait makes sure that we do not shut down the program before all requests in the channel have been completely processed
	// Without this wait we run the risk of terminating main, and therfore shutting down the whole program, while some requests
	// are still being processed by some of the workers
	wg.Wait()

	fmt.Println("All requests processed")
}

func doWork(inCh <-chan Request, wg *sync.WaitGroup, i int) {
	fmt.Printf("Worker %v started\n", i)
	for req := range inCh {
		time.Sleep(100 * time.Millisecond)
		fmt.Printf("Request executed with parameter  %v\n", req.Param)
	}
	wg.Done()
	fmt.Printf("Worker %v shutting down\n", i)
}
