package main

import (
	"context"
	"fmt"
	"io"
	"os"
	"strings"
	"time"
)

// this a type implementing the io.Reader interface
type slowReaderCtx struct {
	ctx context.Context
	r   io.Reader
}

// Reads slowly, sleeping for some time between each read
func (r *slowReaderCtx) Read(p []byte) (n int, err error) {
	if err := r.ctx.Err(); err != nil {
		return 0, err
	}
	time.Sleep(100 * time.Millisecond)
	fmt.Print("Read operation executed\n")
	return r.r.Read(p)
}

// NewReader returns a context-aware io.Reader.
func NewSlowReader(ctx context.Context, r io.Reader) io.Reader {
	return &slowReaderCtx{
		ctx: ctx,
		r:   r,
	}
}

// this function read slowly a file and write its content into a string
// It reads slowly because we are using the slowReaderCtx type.
// Once all the file has been read, the string containing the content of the file is sent to the resChan
// If the context is cancelled while the read is still on fly, then the function writes the error on the errChan and returns
func slowlyReadFileToString(ctx context.Context, path string, resChan chan any, errChan chan error) {

	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}

	// here we create the slowReader
	slowReader := NewSlowReader(ctx, f)

	// here we create the writer
	var w strings.Builder

	buf := make([]byte, 1024)

	// we use io.CopyBuffer since we want to set a buffer small enough so that it will require more reads and therefore more time to complete the copy
	res, err := io.CopyBuffer(&w, slowReader, buf)
	if err != nil {
		fmt.Printf("Error while copying: %v\n", err)
		errChan <- fmt.Errorf("slowlyReadFileToString failed with error '%v'", ctx.Err())
		return
	}
	resChan <- fmt.Sprintf("%v characters read slowly from the file", res)
}
