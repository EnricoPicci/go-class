package main

import (
	"fmt"
	"sync"
)

type t struct{ val string }

func (aT t) tellMe(wg *sync.WaitGroup) {
	fmt.Printf("I am a %T with val %v in a goroutine\n", aT, aT.val)
	wg.Done()
}

func doStuff(wg *sync.WaitGroup) {
	fmt.Println("I am the function 'doStuff' doing stuff in a goroutine")
	wg.Done()
}

func main() {
	var wg sync.WaitGroup
	// we are going to launch 3 goroutines
	// the actual sequence in which they run can not be predicted since the Go scheduler is not
	wg.Add(3)

	// launch a goroutine as literal (anonymous) function
	go func() {
		fmt.Println("I am an anonymous function doing stuff in a goroutine")
		wg.Done()
	}()

	// launch a goroutine as a function
	go doStuff(&wg)

	// launch a goroutine as a method of a struct
	var aT t = t{"aT"}
	go aT.tellMe(&wg)

	wg.Wait()
	fmt.Printf("\nTerminating the program\n")
}
