package main

import "sync"

// In this example 2 goroutines lock and unlock 2 mutexes in a different order.
// The first one lockss	mu_1 and then locks mu_2. The second one locks mu_2 and then locks mu_1.
// This is a deadlock which is detected by the runtime since, at a certain point, the runtime sees
// that all goroutines are in waiting state.

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

	// wait forever, until the runtime detects a deadlock
	select {}
}

var lock_1_then_2_executions int
var lock_unlock_2_executions int

func lock_1_then_2() {
	mu_1.Lock()
	// call a function that locks another mutex
	lock_unlock_2()
	// do something
	lock_1_then_2_executions++
	mu_1.Unlock()
}
func lock_unlock_2() {
	mu_2.Lock()
	// do something
	lock_unlock_2_executions++
	mu_2.Unlock()
}

var lock_2_then_1_executions int
var lock_unlock_1_executions int

func lock_2_then_1() {
	mu_2.Lock()
	// call a function that locks another mutex
	lock_unlock_1()
	// do something
	lock_2_then_1_executions++
	mu_2.Unlock()
}
func lock_unlock_1() {
	mu_1.Lock()
	// do something
	lock_unlock_1_executions++
	mu_1.Unlock()
}
