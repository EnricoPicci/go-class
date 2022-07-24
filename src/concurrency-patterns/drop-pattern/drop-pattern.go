package main

import (
	"flag"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// In this example we implement drop pattern

type Request struct {
	Param int
}

// set the size of the worker pool
var poolSize *int

// average interval between requests
var avgReqInterval *int

// ratio between average time spent to process a request and the interval between subsequent requests - a ratio of 5, for instance,
// means that 5 requests come in, on average, in the time that one request is processed
var procIntervalRatio *float64
var timeUnit = time.Millisecond

var requestProcessed = 0
var requestDropped = 0
var bufferChanCumulativeOccupation = 0
var maxBufferChanOccupation = 0

func main() {
	poolSize = flag.Int("poolSize", 5, "number of workers in the worker pool")
	chanCap := flag.Int("chanCap", 5, "capacity of the channel")
	avgReqInterval = flag.Int("avgReqInterval", 100, "average interval between requests in miliseconds")
	procIntervalRatio = flag.Float64("procIntervalRatio", 10, "ratio between average time spent to process a request and the interval between subsequent requests")
	numReq := flag.Int("numReq", 100, "number of requests coming in to be processed")
	flag.Parse()

	fmt.Println("Start processing requests")
	fmt.Print("\n")

	flag.VisitAll(func(f *flag.Flag) {
		fmt.Printf("%s: %s  (%s)\n", f.Name, f.Value, f.Usage)
	})
	fmt.Print("\n")

	// the channel that provides requests to the pool is buffered
	// the capacity of the buffer is the number of requests we accept to remain in the queue
	// if a request comes in when the channel buffer is full, such request is dropped
	inCh := make(chan Request, *chanCap)

	var wg sync.WaitGroup
	wg.Add(*poolSize)

	// start the worker pool
	startPool(inCh, &wg, *poolSize)

	// we simulate a stream of incoming requests
	for i := 0; i < *numReq; i++ {
		// random intervale between each incoming request - the random interval is centered around the average interval
		var intervalBetweenRequests = time.Duration(rand.Intn((*avgReqInterval)*2)) * timeUnit

		time.Sleep(time.Duration(intervalBetweenRequests))
		// with this select we implement the drop pattern
		select {
		// if there is room in the channel buffer, then the request is sent to the pool
		case inCh <- Request{i}:
			bufferChanCumulativeOccupation = bufferChanCumulativeOccupation + len(inCh)
			if maxBufferChanOccupation < len(inCh) {
				maxBufferChanOccupation = len(inCh)
			}
		// if the channel buffer is full the request is dropped
		default:
			fmt.Printf("Request %v dropped\n", i)
			requestDropped++
			bufferChanCumulativeOccupation = bufferChanCumulativeOccupation + (*chanCap)
			maxBufferChanOccupation = *chanCap
		}
	}

	// when the stream completes we close the channel
	close(inCh)

	// This Wait makes sure that we do not shut down the program before all requests in the channel have been completely processed
	// Without this wait we run the risk of terminating main, and therfore shutting down the whole program, while some requests
	// are still being processed by some of the workers
	wg.Wait()

	fmt.Printf("\n\nRequests processed: %v\n", requestProcessed)
	fmt.Printf("Requests dropped: %v\n", requestDropped)
	fmt.Printf("Average buffer occupation: %v\n", bufferChanCumulativeOccupation / *numReq)
	fmt.Printf("Max buffer occupation: %v\n", maxBufferChanOccupation)
}

func startPool(inCh <-chan Request, wg *sync.WaitGroup, poolSize int) {
	// start the workers
	i := 0
	for i < poolSize {
		go doWork(inCh, wg, i)
		i++
	}
}

func doWork(inCh <-chan Request, wg *sync.WaitGroup, i int) {
	fmt.Printf("Worker %v started\n", i)
	for req := range inCh {
		// random delay that simulates the work done while processing a request

		avgProcTime := float64((*avgReqInterval)) * (*procIntervalRatio)
		var processingTime = time.Duration(rand.Intn(int(avgProcTime)*2)) * timeUnit

		time.Sleep(time.Duration(processingTime))
		fmt.Printf("Request executed with parameter  %v\n", req.Param)
		requestProcessed++
	}
	wg.Done()
	fmt.Printf("Worker %v shutting down\n", i)
}
