package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

func main() {
	var url = "https://en.wikipedia.org/wiki/Go_(programming_language)"

	fmt.Println("Execute a Get http request using all defaults")
	executeGetDefault(url)

	fmt.Print("\n\n\n")

	fmt.Println("Execute the same Get http request passing a context which has a timeout that should be big enough for the request to complete")

	executeWithTimeout(url, 1*time.Second)

	fmt.Print("\n\n\n")

	fmt.Println("Execute the same Get http request again, this time passing a context which has a timeout that should be too short for the request to complete")

	executeWithTimeout(url, 1*time.Millisecond)
}

func executeGetDefault(url string) {
	start := time.Now()

	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%v have been read in %v ms", len(body), time.Since(start).Milliseconds())
}

func executeWithTimeout(url string, d time.Duration) {
	start := time.Now()

	ctx, cancel := context.WithTimeout(context.Background(), d)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		log.Fatalf("%v", err)
	}

	client := http.DefaultClient
	resp, err := client.Do(req)
	if err != nil {
		// if the context has an error, it means that it has timeouted
		if e := ctx.Err(); e != nil {
			log.Printf("%v", e)
			return
		}
		// if there is no error in the context and we end up here, then we have an unexpected problem
		panic(err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%v have been read in %v ms", len(body), time.Since(start).Milliseconds())
}
