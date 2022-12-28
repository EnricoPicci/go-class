package main

import (
	"fmt"

	"github.com/EnricoPicci/go-class/src/concurrency-patterns/recursive-examples/hilberthotel"
)

func BusClerk(busNumber int) func(i int) hilberthotel.WelcomeKit {
	var count = 0
	var passengerNumber = 1
	var nextClerkCh func(i int) hilberthotel.WelcomeKit

	return func(i int) hilberthotel.WelcomeKit {
		count++
		if count == passengerNumber {
			passengerNumber++
			count = 0
			return hilberthotel.NewWelcomeKit(busNumber, passengerNumber-1, i)
		}

		if nextClerkCh == nil {
			nextClerkCh = BusClerk(busNumber + 1)
		}
		return nextClerkCh(i)
	}
}

func Hilbert(upTo int) []hilberthotel.WelcomeKit {
	var wellcomeKits = []hilberthotel.WelcomeKit{}
	var firstBusClerk = BusClerk(1)
	for i := 1; i <= upTo; i++ {
		kit := firstBusClerk(i)
		wellcomeKits = append(wellcomeKits, kit)
	}

	fmt.Println()
	fmt.Printf("%v guests have been given a room by Hilber at his Hotel\n", upTo)

	return wellcomeKits
}
