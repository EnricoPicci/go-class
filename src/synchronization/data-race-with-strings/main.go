// This example has been inspired by [this blog](https://dave.cheney.net/2014/06/27/ice-cream-makers-and-data-races).
// In the example there is a shared variable, sharedString, of whose type is a string.
// There are 2 goroutines that concurrently set the value of this variable: one assignes a long string, one assignes a short string.
// There is also a goroutine that reads concurrently the same shared variable and checks whether it can happen that
// the lenght of the string is not coherent with its content. When this occurs, a message is printed and the program exits.

package main

import (
	"log"
)

var sharedString string

func main() {
	longString := "1234567890"
	shortString := "1"

	readers := 1

	var setLongString, setShortString, readString func()

	setLongString = func() {
		sharedString = longString
		go setShortString()
	}
	setShortString = func() {
		sharedString = shortString
		go setLongString()
	}

	readString = func() {
		i := 0
		for {
			// checks whether the content stored in the shared variable and its lenght are coherent
			if sharedString == shortString && len(sharedString) != len(shortString) {
				log.Fatalf("'sharedString == shortString' is true but 'len(sharedString) != len(shortString)' is also true ---- iteration %v", i)
			}
			if sharedString == longString && len(sharedString) != len(longString) {
				log.Fatalf("'sharedString == longString' is true but 'len(sharedString) != len(longString)' is also true ---- iteration %v", i)
			}
			i++
		}
	}

	for i := 0; i < readers; i++ {
		go readString()
	}

	setLongString()

	select {}
}
