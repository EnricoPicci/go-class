package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"sync"
	"time"
)

func main() {
	var url_GO = "https://en.wikipedia.org/wiki/Go_(programming_language)"
	var url_Gr = "https://en.wikipedia.org/wiki/Robert_Griesemer"
	var url_Pi = "https://en.wikipedia.org/wiki/Rob_Pike"
	var url_Th = "https://en.wikipedia.org/wiki/Ken_Thompson"

	fmt.Println("Execute a bunch of get http requests which should all be successful")
	executeConcurrentlyAllOrNothing([]string{url_GO, url_Gr, url_Pi, url_Th})

	fmt.Print("\n\n\n")

	fmt.Println("Try to execute a bunch of get http requests while one fails and therefore all are cancelled")
	executeConcurrentlyAllOrNothing([]string{url_GO, url_Gr, url_Pi, url_Th, "blah"})

}

func executeConcurrentlyAllOrNothing(urls []string) {
	// the parent context is created as well as the cancel function for it (use parentCtx, parentCancel names instead of the conventional ctx, cancel
	// with the aim to increase the clarity of the example)
	parentCtx, parentCancel := context.WithCancel(context.Background())
	defer parentCancel()

	var wg sync.WaitGroup
	wg.Add(len(urls))

	// launch the requests in parallel using one goroutine per request
	for _, url := range urls {
		go executeGet(parentCtx, url, parentCancel, &wg)
	}

	wg.Wait()

}

// the context passed to this function is a child of the parent context created by its caller, i.e. the function executeConcurrentlyAllOrNothing
// the cancel function passed, on the contrary, is the function that will cancel the parent context and, therefore, also all of its children
func executeGet(childCtx context.Context, url string, parentCancel context.CancelFunc, wg *sync.WaitGroup) {
	defer wg.Done()

	start := time.Now()

	// we pass childCtx to each http request - when the parent context is cancelled, then all of its children contexts will be cancelled
	req, err := http.NewRequestWithContext(childCtx, "GET", url, nil)
	if err != nil {
		log.Fatalf("%v", err)
	}

	client := http.DefaultClient
	resp, err := client.Do(req)
	if err != nil {
		// if an error occurs, then the parent context is cancelled, causing all of its children to be cancelled
		parentCancel()
		fmt.Printf("- An error %v has cancelled the get %v\n", err, url)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	fmt.Printf("- %v have been read from %v in %v ms\n", len(body), url, time.Since(start).Milliseconds())
}
