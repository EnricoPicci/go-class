package main

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"time"
)

// this function processes a file line by line
// it is a slow processor since, for every line read, it simulates heavy processing by sleeping for some time
func slowlyProcess(ctx context.Context, path string, resChan chan any, errChan chan error) {

	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	sc := bufio.NewScanner(f)
	linesCount := 0

	for sc.Scan() {
		select {
		case <-ctx.Done():
			fmt.Printf("The context has been cancelled - the error is %v\n", ctx.Err())
			fmt.Printf("Processing cancelled after reading %v lines\n", linesCount)
			errChan <- fmt.Errorf("slowlyProcess failed with error '%v'", ctx.Err())
			return
		default:
			// sleep at each line read to simulate some processing for each line
			time.Sleep(10 * time.Millisecond)
			linesCount++
			fmt.Printf("Processed line %v\n", linesCount)
		}
	}
	if err := sc.Err(); err != nil {
		panic(fmt.Sprintf("scan file error: %v", err))
	}

	resChan <- fmt.Sprintf("%v lines processed by the slow processor", linesCount)
}
