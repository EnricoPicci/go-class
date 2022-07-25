package main

import (
	"context"
	"flag"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// In this example we implement drop pattern with timeout

type Request struct {
	Param        int
	created      time.Time
	waitDuration time.Duration
}

// set the size of the worker pool
var poolSize *int

// time unit utilized to calculate durations
var timeUnit = time.Millisecond

// average interval between requests
var avgReqInterval *int

// ratio between average time spent to process a request and the interval between subsequent requests - a ratio of 5, for instance,
// means that 5 requests come in, on average, in the time that one request is processed
var procIntervalRatio *float64

var muReqProcessed sync.Mutex
var requestProcessed = 0
var muReqDropped sync.Mutex
var requestDropped = 0

var muReqWaiting sync.Mutex
var requestWaiting = 0
var maxWaitingQueueSize = 0
var maxWaitDuration time.Duration

var muReqIdleTime sync.Mutex
var workersIdleTime time.Duration

var extraordinaryEvent *bool

func main() {
	poolSize = flag.Int("poolSize", 5, "number of workers in the worker pool")
	timeout := flag.Int("timeout", 100, "timeout in miliseconds after which a request waiting to be processed is dropped")
	avgReqInterval = flag.Int("avgReqInterval", 100, "average interval between requests in miliseconds")
	procIntervalRatio = flag.Float64("procIntervalRatio", 10, "ratio between average time spent to process a request and the interval between subsequent requests")
	numReq := flag.Int("numReq", 100, "number of requests coming in to be processed")
	extraordinaryEvent = flag.Bool("extraordinaryEvent", false, "if true, it is possible that an extraordinary event occurs that generates a peak in processing time "+
		"(this simulates for instance a transient problem on a DB that slows down the processing of some requests)")
	flag.Parse()

	fmt.Println("Start processing requests")
	fmt.Print("\n")

	flag.VisitAll(func(f *flag.Flag) {
		fmt.Printf("%s: %s  (%s)\n", f.Name, f.Value, f.Usage)
	})
	fmt.Print("\n")

	// the channel that provides requests to the pool is unbuffered
	inPoolCh := make(chan Request)

	var wgPool sync.WaitGroup
	wgPool.Add(*poolSize)

	// start the worker pool
	startPool(inPoolCh, &wgPool, *poolSize)

	ctx := context.Background()

	var wgReq sync.WaitGroup

	// we simulate a stream of incoming requests
	for i := 0; i < *numReq; i++ {
		// random interval between each incoming request - the random interval is centered around the average interval
		var intervalBetweenRequests = time.Duration(rand.Intn((*avgReqInterval)*2)) * timeUnit
		time.Sleep(time.Duration(intervalBetweenRequests))
		req := Request{i, time.Now(), 0}

		wgReq.Add(1)

		fmt.Printf("Sent %v\n", req.Param)
		increaseRequestWaiting(req)

		// within this goroutine we implement the drop with timeout pattern
		go processOrDropWithTimeout(ctx, req, inPoolCh, &wgReq, *timeout)
	}

	// wait for all the gouroutines launched for each request to either be able to pass the request to the pool or to drop it
	wgReq.Wait()
	// when there are no more requests that can enter the pool we close the in pool channel
	close(inPoolCh)

	// This Wait makes sure that we do not shut down the program before all requests in the channel have been completely processed
	// Without this wait we run the risk of terminating main, and therfore shutting down the whole program, while some requests
	// are still being processed by some of the workers
	wgPool.Wait()

	fmt.Printf("\n\nRequests processed: %v\n", requestProcessed)
	fmt.Printf("Requests dropped: %v\n", requestDropped)

	fmt.Printf("Max waiting queue size: %v\n", maxWaitingQueueSize)
	fmt.Printf("Max wait duration: %v\n", maxWaitDuration)

	fmt.Printf("Average idle time for a worker: %v\n", time.Duration(int(workersIdleTime) / *poolSize))
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
	var startIdleTime time.Time
	for req := range inCh {
		if !startIdleTime.IsZero() {
			muReqIdleTime.Lock()
			workersIdleTime = workersIdleTime + time.Since(startIdleTime)
			muReqIdleTime.Unlock()
		}
		waitDuration := time.Since(req.created)
		req.waitDuration = waitDuration
		if maxWaitDuration < waitDuration {
			maxWaitDuration = waitDuration
		}

		decreaseRequestWaiting("start processing", req)

		// random delay that simulates the work done while processing a request
		avgProcTime := float64((*avgReqInterval)) * (*procIntervalRatio)
		var processingTime = time.Duration(rand.Intn(int(avgProcTime)*2)) * timeUnit
		// if extraordinaryEvent is true, we can see a peak in the time required to process a request
		if *extraordinaryEvent {
			// in the 2% of cases we have an extraordinary long processing time
			r := rand.Intn(100)
			if r < 2 {
				fmt.Println(">>>>>>>>>>>>>>>>> Extraordinary event <<<<<<<<<<<<<<<<")
				processingTime = processingTime * 10
				fmt.Printf(">>>>>>>>>>>>>>>>> %v <<<<<<<<<<<<<<<<\n", processingTime)
			}
		}

		time.Sleep(time.Duration(processingTime))
		fmt.Printf("===>>>> Request executed with parameter %v - wait time %v\n", req.Param, req.waitDuration)
		muReqProcessed.Lock()
		requestProcessed++
		muReqProcessed.Unlock()
		startIdleTime = time.Now()
	}
	wg.Done()
	fmt.Printf("Worker %v shutting down\n", i)
}

func increaseRequestWaiting(req Request) {
	muReqWaiting.Lock()
	if maxWaitingQueueSize < requestWaiting {
		maxWaitingQueueSize = requestWaiting
	}
	requestWaiting++
	rw := requestWaiting // make a copy so that the following print prints exactly this value
	muReqWaiting.Unlock()
	fmt.Printf("Requests waiting (sent - %v) %v - %v\n", req.Param, rw, maxWaitingQueueSize)
}

func decreaseRequestWaiting(reason string, req Request) {
	muReqWaiting.Lock()
	requestWaiting--
	muReqWaiting.Unlock()
	fmt.Printf("Requests waiting (%v - %v)\n", reason, req.Param)
}

// this function implements the drop with timeout pattern
func processOrDropWithTimeout(ctx context.Context, req Request, inPoolCh chan<- Request, wgReq *sync.WaitGroup, timeout int) {
	defer wgReq.Done()

	// the timeout context
	ctx, cancel := context.WithTimeout(ctx, time.Duration(timeout)*timeUnit)
	defer cancel()

	// wait until the request is processed or the context times out
	select {
	case inPoolCh <- req:
		// the request is sent to the pool
	case <-ctx.Done():
		// the context times out and the request is dropped
		decreaseRequestWaiting("drop", req)
		fmt.Printf("Request %v dropped\n", req.Param)
		muReqDropped.Lock()
		requestDropped++
		muReqDropped.Unlock()
	}
}
