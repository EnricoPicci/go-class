package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"
)

var counter int
var mu = sync.Mutex{}

var sleepPtr *int
var logPtr *bool

// higher level function which returns an handler function - uses currying to set the sleep time
func handlerFact(sleep *int) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if *logPtr {
			log.Printf("Request %v coming in", counter)
		}
		var _counter int
		mu.Lock()
		counter++
		// use a local variable to copy the counter - in this way we can pass the local variable down to the calling functions
		// without the risk of having it changed by other executions of the handler (which can be invoked concurrently by concurrent requests coming in)
		_counter = counter
		mu.Unlock()

		// simulate work by sleeping
		time.Sleep(time.Duration(*sleep) * time.Millisecond)

		fmt.Fprintf(w, "Request %v completed", _counter)
	}
}

func sleepQueryHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%v", *sleepPtr)
}

func main() {
	sleepPtr = flag.Int("sleep", 10, "number of milliseconds each request processing sleeps to simulate work")
	portPtr := flag.Int("port", 8080, "the port the server runs on")
	logPtr = flag.Bool("log", false, "log on the standard output")
	flag.Parse()

	http.HandleFunc("/dostuff", handlerFact(sleepPtr))
	http.HandleFunc("/sleep", sleepQueryHandler)

	fmt.Printf("Starting server on port %v - the processing of each request takes %v milliseconds\n", *portPtr, *sleepPtr)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", *portPtr), nil))
}
