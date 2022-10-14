package main

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/EnricoPicci/go-class/src/testhelpers"
)

// this function processes a file line by line
// it is a slow processor since, for every line read, it simulates heacy processing by sleeping for some time
func processFileLineByLine(ctx context.Context, path string, wg *sync.WaitGroup) {
	defer wg.Done()

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
			return
		default:
			// sleep at each line read to simulate some processing for each line
			time.Sleep(100 * time.Millisecond)
			linesCount++
		}
	}
	if err := sc.Err(); err != nil {
		panic(fmt.Sprintf("scan file error: %v", err))
	}
}

func main() {
	// we wanty to read this file with our slowReader
	var path = testhelpers.FilePath("/canti-divina-commedia/01 - Inferno - CANTO PRIMO.txt")

	// we create a context that times out after a bit less than 1 second, which means that we allow our slowReader to read the first buffer only
	// when the Read method of slowReader gets called for the second time, then the context will have already timed out and therefore
	// there will be no second read
	ctx, cancel := context.WithTimeout(context.Background(), 900*time.Millisecond)
	defer cancel()

	var wg sync.WaitGroup
	wg.Add(1)
	go processFileLineByLine(ctx, path, &wg)

	wg.Wait()

	fmt.Println("Program terminating")
}
