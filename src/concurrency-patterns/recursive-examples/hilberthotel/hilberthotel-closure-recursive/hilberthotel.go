package hilberthotelclosurerecursive

import (
	"fmt"
	"time"

	"github.com/EnricoPicci/go-class/src/concurrency-patterns/recursive-examples/hilberthotel"
)

func BusClerk(busNumber int, delay time.Duration) func(i int) hilberthotel.WelcomeKit {
	var count = 0
	var passengerNumber = 1
	var nextClerk func(i int) hilberthotel.WelcomeKit

	return func(i int) hilberthotel.WelcomeKit {
		count++
		if count == passengerNumber {
			passengerNumber++
			count = 0
			return hilberthotel.NewWelcomeKit(busNumber, passengerNumber-1, i, delay)
		}

		if nextClerk == nil {
			nextClerk = BusClerk(busNumber+1, delay)
		}
		return nextClerk(i)
	}
}

func Hilbert(upTo int, delay time.Duration, verbose bool) []hilberthotel.WelcomeKit {
	var wellcomeKits = []hilberthotel.WelcomeKit{}
	var firstBusClerk = BusClerk(1, delay)
	for i := 1; i <= upTo; i++ {
		kit := firstBusClerk(i)
		wellcomeKits = append(wellcomeKits, kit)
	}

	if verbose {
		fmt.Println()
		fmt.Printf("%v guests have been given a room by Hilber at his Hotel\n", upTo)
	}

	return wellcomeKits
}
