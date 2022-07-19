package main

import (
	"context"
	"fmt"
	"io"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/EnricoPicci/go-class.git/src/testhelpers"
)

// this a type implementing the io.Reader interface
type slowReaderCtx struct {
	ctx context.Context
	r   io.Reader
}

// Reads slowly, sleeping 1 sec between each read
func (r *slowReaderCtx) Read(p []byte) (n int, err error) {
	if err := r.ctx.Err(); err != nil {
		return 0, err
	}
	time.Sleep(1 * time.Second)
	fmt.Println(">>>>> Execute read")
	return r.r.Read(p)
}

// NewReader returns a context-aware io.Reader.
func NewSlowReader(ctx context.Context, r io.Reader) io.Reader {
	return &slowReaderCtx{
		ctx: ctx,
		r:   r,
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

	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}

	// here we create the slowReader
	slowReader := NewSlowReader(ctx, f)
	// here we create the writer (we do not use os.Stdout since this type implements the ReaderFrom interface which prevents from being able to
	// use a small buffer - see https://pkg.go.dev/io#CopyBuffer)
	var w strings.Builder

	var wg sync.WaitGroup
	wg.Add(1)
	go copy(&w, slowReader, &wg)

	wg.Wait()

	fmt.Printf("Heve been able to copy\n %v chars\n", len(w.String()))

	fmt.Println("Program terminating")
}

func copy(w io.Writer, r io.Reader, wg *sync.WaitGroup) {
	defer wg.Done()

	buf := make([]byte, 1024)

	// we use io.CopyBuffer since we want to set a buffer small enough to read only a part of the file
	_, err := io.CopyBuffer(w, r, buf)
	if err != nil {
		fmt.Printf("Error while copying: %v\n", err)
	}
}
