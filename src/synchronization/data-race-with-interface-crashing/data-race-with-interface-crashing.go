package main

import (
	"fmt"
	"os"
)

// In this example we define an interface type Doer and a shared global variable, sharedInterface, of type Doer.
// The shared variable is assigned values by 2 different goroutines concurrenlty.
// The 2 goroutines assign values of different types to sharedInterface variable, one of type Human and one of type Computer
// So, in a concurrent way, we assign values of different concrete types to a variable of type interface.
// Since the assignement is not protected by synchronization mechanisms, we can have data corruption that leads to runtime errors
type Doer interface {
	DoStuff() bool
}

var sharedInterface Doer

const human = "Human"
const computer = "Computer"

type Human struct {
	Name           string
	MyLuckyNumbers []int
}

func (d Human) DoStuff() bool {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in f", r)
			fmt.Println("These are my corrupted lucky numbers")
			for _, v := range d.MyLuckyNumbers {
				fmt.Println(">>>>>>>******", v)
			}
		}
	}()

	// do something with the MyLuckyNumbers
	// if the receiver d happens to be corrupted (which is possible since the share variable sharedInterface is not protected against data races)
	// then there may be a runtime error
	if d.MyLuckyNumbers[5] != 5 {
		fmt.Println("My sixth lucky number should be 5")
		return false
	}

	return true
}

type Machine struct {
	Name  string
	Model string
}

func (d Machine) DoStuff() bool {
	// do nothing and return always true
	return true
}

func main() {

	// this is a function that calls DoStuff many times and checks whether DoStuff returns false.
	// If it returns false it means that the interface value is corrupted and the corruption derives that the assignement to the shared variable
	// sharedInterface (the assignement is a write operation) is run concurrently and is not protected by a synchronization mechanism
	doStuffManyTimes := func(d Doer) {
		i := 0
		for {
			// this is the assignement operation (write operation) which is not protected and therefore can cause data corruption
			sharedInterface = d
			if !sharedInterface.DoStuff() {
				fmt.Printf("Something strange happened in iteration %v\n", i)
				os.Exit(1)
			}
			i++
		}
	}

	luckyNumbers := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	d_1 := Human{human, luckyNumbers}
	go doStuffManyTimes(&d_1)

	d_2 := Machine{computer, "M24"}
	go doStuffManyTimes(d_2)

	// stop forever used here just for the laziness of not adding wait groups
	select {}
}
