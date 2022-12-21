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
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var wg sync.WaitGroup
	wg.Add(len(urls))

	for _, url := range urls {
		go executeGet(ctx, url, cancel, &wg)
	}

	wg.Wait()

}

func executeGet(ctx context.Context, url string, cancel context.CancelFunc, wg *sync.WaitGroup) {
	defer wg.Done()

	start := time.Now()

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		log.Fatalf("%v", err)
	}

	client := http.DefaultClient
	resp, err := client.Do(req)
	if err != nil {
		cancel()
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
