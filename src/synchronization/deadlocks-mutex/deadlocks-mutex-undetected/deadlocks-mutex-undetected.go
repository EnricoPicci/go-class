package main

import (
	"fmt"
	"sync"
	"time"
)

// In this example 2 goroutines lock and unlock 2 mutexes in a different order.
// The first one lockss	mu_1 and then locks mu_2. The second one locks mu_2 and then locks mu_1.
// This is a deadlock but the runtime does not detect it since we have also launched a third goroutine.
// Even if the thirg goroutine does nothing but wait for a timeout, the runtime imagines
// that sooner or later it could unlock one of the mutexed and resume the processing of the other 2 goroutines.
//
// This example is to show that the support for deaclock detection provided by the runtime is very rudimentary and
// does not cover most of the real cases, where few goroutines got stuck in a dealock while the rest of the program works.
// The end result of such situations are unpredictable and difficult to debug.

var mu_1 sync.Mutex
var mu_2 sync.Mutex

func main() {
	go func() {
		for {
			lock_1_then_2()
		}
	}()
	go func() {
		for {
			lock_2_then_1()
		}
	}()

	// launch a third gorourine with a timeout - the fact that there is a third goroutine alive prevents the runtime from detecting a deadlock
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		time.Sleep(3 * time.Second)
		wg.Done()
	}()

	wg.Wait()

	fmt.Println("Program terminated")

}

var lock_1_then_2_executions int

func lock_1_then_2() {
	mu_1.Lock()
	mu_2.Lock()
	// do something
	lock_1_then_2_executions++
	mu_2.Unlock()
	mu_1.Unlock()
}

var lock_2_then_1_executions int

func lock_2_then_1() {
	mu_2.Lock()
	mu_1.Lock()
	// do something
	lock_2_then_1_executions++
	mu_1.Unlock()
	mu_2.Unlock()
}
