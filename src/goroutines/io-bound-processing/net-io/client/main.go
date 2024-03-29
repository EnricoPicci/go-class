package main

import (
	"flag"
	"fmt"
	"runtime"
	"time"

	"github.com/EnricoPicci/go-class/src/goroutines/io-bound-processing/net-io/client/send"
)

func main() {
	requestsPtr := flag.Int("requests", 100, "number of requests to send to the server")
	concurrentPtr := flag.Int("concurrent", 1, "concurrency level used, i.e. number of goroutines sending requests to the server concurrently")
	maxProcs := flag.Int("maxprocs", runtime.NumCPU(), "maximum number of cores to use")
	flag.Parse()

	runtime.GOMAXPROCS(*maxProcs)

	fmt.Printf("Send %v requests to the server with %v concurrent goroutines and %v active cores\n", *requestsPtr, *concurrentPtr, *maxProcs)

	sleep := send.GetSleepTime()
	fmt.Printf("Each request takes %v milliseconds to be processed\n", sleep)

	start := time.Now()
	send.CallServer(*requestsPtr, *concurrentPtr)

	fmt.Println(">>>>>>>>>>>> Time to process all requests:", time.Since(start))

}
