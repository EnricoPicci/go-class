package main

import (
	"fmt"
	"os"
)

// In this example we define an interface type Doer and a shared global variable, sharedInterface, of type Doer.
// The shared variable is assigned values by 2 different goroutines concurrenlty.
// The 2 goroutines assign values of different types to sharedInterface variable, one of type Human and one of type Computer
// So, in a concurrent way, we assign values of different concrete types to a variable of type interface.
// Since the assignement is not protected by synchronization mechanisms, we can have data corruption that leads to wierd behavior:
// a concrete value of a certain type runs the method implementation of a different type; for instance a value of type Human executes the method DoStuff
// implemented by the type Computer, or viceversa
type Doer interface {
	DoStuff() bool
}

var sharedInterface Doer

const human = "Human"
const computer = "Computer"

type Human struct {
	Name string
	Age  int
}

func (d Human) DoStuff() bool {
	if d.Name != human {
		fmt.Printf("I am the Human doing stuff as a %v\n", d.Name)
		return false
	}

	return true
}

type Machine struct {
	Name  string
	Model string
}

func (d Machine) DoStuff() bool {
	if d.Name != computer {
		fmt.Printf("I am the Machine doing stuff as a %v\n", d.Name)
		return false
	}

	return true
}

func main() {

	// this is a function that calls DoStuff many times and checks whether DoStuff returns false.
	// If it returns false it means that the interface value is corrupted and the corruption derives that the assignement to the shared variable
	// sharedInterface (the assignement is a write operation) is run concurrently and is not protected by a synchronization mechanism
	doStuffManyTimes := func(d Doer) {
		for i := 0; i < 1000000; i++ {
			// this is the assignement operation (write operation) which is not protected and therefore can cause data corruption
			sharedInterface = d
			if !sharedInterface.DoStuff() {
				fmt.Printf("Something strange happened in iteration %v\n", i)
				os.Exit(1)
			}
		}
		fmt.Println("DONE")
	}

	d_1 := Human{human, 20}
	go doStuffManyTimes(d_1)

	d_2 := Machine{computer, "M24"}
	go doStuffManyTimes(d_2)

	// stop forever used here just for the laziness of not adding wait groups
	select {}
}
