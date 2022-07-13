package main

import (
	"context"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

var waitForCancelPtr *bool
var timeoutPtr *int

func handler(w http.ResponseWriter, r *http.Request) {
	log.Print("Request coming in")

	// create a context with timouout starting from the request context
	ctx, cancel := context.WithTimeout(r.Context(), time.Duration(*timeoutPtr)*time.Millisecond)
	defer cancel()
	// process the request
	resp, err := process(ctx, waitForCancelPtr)
	if err != nil {
		// check if the context has an error
		if e := ctx.Err(); e != nil {
			// the error can be generated because of a timeout or because of a cancellation of the context or of any of its parents
			switch err {
			case context.Canceled:
				log.Println("Request cancelled")
			case context.DeadlineExceeded:
				log.Println("Request time out")
			default:
				log.Printf("Another type of error %v, %T\n", err, err)
			}
			fmt.Fprintf(w, "Request errored with  error %v", e)
			return
		}
		log.Printf("Error in processing the request: %v", err)
		fmt.Fprintf(w, "Request errored with  error %v", err)
		return
	}

	fmt.Fprint(w, resp)
}

// this function simulates some work done to process a request - in particular we execute an http get and therefore we can pass
// the context to the http layer which is context aware
// if waitCancel is true, then it waits for the client to close the request (e.g. closing the browser)
func process(ctx context.Context, waitCancel *bool) (string, error) {
	var url = "https://en.wikipedia.org/wiki/Go_(programming_language)"

	// if we have set the waitCancel flag then we wait until the request is cancelled, i.e. closing the browser which has generated the request
	if *waitCancel {
		log.Println("Waiting for the client to close the request. If the client is a browser just close the browser tab.")
		<-ctx.Done()
		// if the channel returned by Done() has fired a signal (in this case a close signal) then we return the error of the context
		return "", ctx.Err()
	}

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return "", err
	}
	client := http.DefaultClient
	resp, err := client.Do(req)
	// if the context timeouts, then an error is returned and we return the error to the caller
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}

func main() {
	waitForCancelPtr = flag.Bool("waitCancel", false, "set it to true if you want to wait for the client to cancel the request")
	portPtr := flag.Int("port", 8080, "the port the server runs on")
	timeoutPtr = flag.Int("timeout", 1, "timeout in milliseconds set on the context derived from the request")
	flag.Parse()

	endpoint := "/process"

	http.HandleFunc(endpoint, handler)

	fmt.Printf("Starting server on port %v\n", *portPtr)
	fmt.Printf("Launch a request from a browser at the address: http://localhost:%v%v\n", *portPtr, endpoint)
	if *waitForCancelPtr {
		fmt.Println("Each requests will wait until the client closes the request, e.g. by closing the broser")
	} else {
		fmt.Printf("Starting server on port 8080 - the timeout to process the request is %v milliseconds\n", *timeoutPtr)
	}

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", *portPtr), nil))
}
