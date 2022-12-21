package main

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"time"

	"github.com/EnricoPicci/go-class/src/testhelpers"
)

// this function processes a file line by line
// it is a slow processor since, for every line read, it simulates some sort of heavy processing by sleeping for some time
func processFileLineByLine(ctx context.Context, path string) {
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
			// in this case there is no need to clenup and free resources, since this is done already by the deferred function which closes the file
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
	// we want to read this file with our slowReader
	var path = testhelpers.FilePath("/canti-divina-commedia/01 - Inferno - CANTO PRIMO.txt")

	// we create a context that times out after a bit less than 1 second, which means that we allow our slowReader to read the first buffer only
	// when the Read method of slowReader gets called for the second time, then the context will have already timed out and therefore
	// there will be no second read
	ctx, cancel := context.WithTimeout(context.Background(), 900*time.Millisecond)
	defer cancel()

	processFileLineByLine(ctx, path)

	fmt.Println("Program terminating")
}
