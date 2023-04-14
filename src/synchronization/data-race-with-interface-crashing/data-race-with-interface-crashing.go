package main

import (
	"fmt"
	"os"
	"unsafe"
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
	// recover from a runtime error just to show why the runtime error has occurred
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in DoStuff implementation for type Human", r)
			fmt.Println("Probably I am executing the DoStuff implementation for type Human on a concrete value of Computer")
			intComp := (*Computer)(unsafe.Pointer(&d))
			fmt.Println("Actually we see that, if doing some unsafe magic, we turn a Human into a Computer the value of the Model is: ", intComp.Model)
		}
	}()

	// do something with the MyLuckyNumbers
	// if the receiver d happens to be corrupted (which is possible since the share variable sharedInterface is not protected against data races)
	// and points to a concrete value of type Computer rather than of type Human, then a runtime error occurs.
	//
	// Why the runtime error occurs.
	// In the program there is only 1 concrete value of type Computer with the attribute "Model" set to "M24".
	// The DoStuff implementation for the type Human looks for the second attribute which is MyLuckyNumbers (of type []int) when everything is good,
	// while it is the attribute Model (of type string) if the data held by the interface is corrupted and the interfaces points to a contrete value of type Computer.
	// The attrinute Model has value "M24", which is a string of length 3,
	// The logic below (i.e. the logic of DoStuff of type Human) wants to access the 5th element, hence an "index out of range" runtime error we experience.
	// If we change the Model of the Computer value to a string with length greater than 5, then the runtime error does not occur any more.
	if d.MyLuckyNumbers[5] != 5 {
		fmt.Println("My sixth lucky number should be 5")
		return false
	}

	return true
}

type Computer struct {
	Name  string
	Model string
}

func (d Computer) DoStuff() bool {
	// do nothing and return always true
	return true
}

func main() {

	// this is a function that calls DoStuff many times and checks whether DoStuff returns false.
	// If it returns false it means that the interface value is corrupted and the corruption derives that the assignement to the shared variable
	// sharedInterface (the assignement is a write operation) is run concurrently and is not protected by a synchronization mechanism
	doStuffManyTimes := func(d Doer) {
		for i := 0; i < 10000000; i++ {
			// this is the assignement operation (write operation) which is not protected and therefore can cause data corruption
			sharedInterface = d
			if !sharedInterface.DoStuff() {
				fmt.Printf("Something strange happened in iteration %v\n", i)
				os.Exit(1)
			}
		}
		fmt.Println("DONE")
	}

	luckyNumbers := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	d_1 := Human{human, luckyNumbers}
	go doStuffManyTimes(&d_1)

	d_2 := Computer{computer, "M24"}
	go doStuffManyTimes(d_2)

	// stop forever used here just for the laziness of not adding wait groups
	select {}
}
