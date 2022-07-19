package main

import (
	"context"
	"flag"
	"fmt"
	"time"

	"github.com/EnricoPicci/go-class.git/src/testhelpers"
)

// This program implements the forkJoin pattern. It launches some subtasks, passing them a context, and wait for their partial results on a result channel.
// If any of the subtasks errors, the error is sent to an error channel.
// If an error is received on the error channel, then the context is cancelled and the other subtasks end their processing gracefully
func main() {
	failAfter := flag.Int("failAfter", 0, "Milliseconds after which the faultyFunction fails. If not set, the faultyFunction does not fail")
	flag.Parse()

	subtasks := 3
	sleepBeforeError := time.Duration(*failAfter) * time.Millisecond

	resChan := make(chan any)
	errChan := make(chan error)

	// Create a context that can be cancelled. It will be cancelled if an error is received on the error channel
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	results := make([]any, subtasks)
	errors := make([]error, subtasks)

	// we want to read this file with our slowReader and we want to process it with our slow processor
	// while these 2 goroutines slowly perform their work, there is the case that faultyFunction fails
	var path = testhelpers.FilePath("/canti-divina-commedia/01 - Inferno - CANTO PRIMO.txt")

	go slowlyReadFileToString(ctx, path, resChan, errChan)
	go slowlyProcess(ctx, path, resChan, errChan)
	go faultyFunction(ctx, sleepBeforeError, resChan, errChan)

	for i := 0; i < subtasks; i++ {
		select {
		case res := <-resChan:
			results[i] = res
		case err := <-errChan:
			fmt.Printf("An error occurred: %v\n", err)
			errors[i] = err
			cancel()
		}
	}

	fmt.Println("These are the results received")
	for _, r := range results {
		fmt.Println(r)
	}
	fmt.Println("\nThese are the errors received")
	for _, e := range errors {
		fmt.Println(e)
	}
	fmt.Println("Program terminating")
}
