package main

import (
	"context"
	"errors"
	"fmt"
	"time"
)

func faultyFunction(ctx context.Context, sleep time.Duration, resChan chan any, errChan chan error) {
	if sleep == time.Duration(0) {
		resChan <- "The faultyFunction is not going to fail. To make faulty pass a duration > 0"
		return
	}
	time.Sleep(sleep)
	fmt.Print("the faulty function is going  to fail\n\n")
	errChan <- errors.New("faultyFunction errored")
}
