package main

import (
	"fmt"
	"os"
)

// type Speaker interface {
// 	Speak() bool
// }

// type Ben struct {
// 	name string
// }

// func (b *Ben) Speak() bool {
// 	if b.name != "Ben" {
// 		fmt.Printf("Ben says %v\n", b.name)
// 		return false
// 	}
// 	return true
// }

// type Jerry struct {
// 	name string
// }

// func (b *Jerry) Speak() bool {
// 	if b.name != "Jerry" {
// 		fmt.Printf("Jerry says %v\n", b.name)
// 		return false
// 	}
// 	return true
// }

// func main() {
// 	fmt.Println("Starting with data race and interface")
// 	var person Speaker

// 	ben := Ben{"Ben"}
// 	jerry := Jerry{"Jerry"}

// 	go func() {
// 		for {
// 			person = &ben
// 			if !person.Speak() {
// 				os.Exit(1)
// 			}
// 		}
// 	}()

// 	go func() {
// 		for {
// 			person = &jerry
// 			if !person.Speak() {
// 				os.Exit(1)
// 			}
// 		}
// 	}()

// 	select {}
// }

type Doer interface {
	DoStuff() bool
}

var human = "Human"
var computer = "Computer"

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
	var sharedInterface Doer

	doStuffManyTimes := func(d Doer) {
		for i := 0; i < 1000000; i++ {
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

	// stop forever
	select {}
}
