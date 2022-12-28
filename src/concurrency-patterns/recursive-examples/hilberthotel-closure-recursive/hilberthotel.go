package main

import "fmt"

type WelcomeKit struct {
	BusNumber       int
	PassengerNumber int
	RoomNumber      int
}

func (e WelcomeKit) String() string {
	return fmt.Sprintf("Bus %v - Passenger %v - Room %v", e.BusNumber, e.PassengerNumber, e.RoomNumber)
}
func NewWelcomeKit(busNumber int, passengerNumber int, roomNmber int) WelcomeKit {
	return WelcomeKit{busNumber, passengerNumber, roomNmber}
}

func BusClerk(busNumber int) func(i int) WelcomeKit {
	var count = 0
	var passengerNumber = 1
	var nextClerkCh func(i int) WelcomeKit

	return func(i int) WelcomeKit {
		count++
		if count == passengerNumber {
			passengerNumber++
			count = 0
			return NewWelcomeKit(busNumber, passengerNumber-1, i)
		}

		if nextClerkCh == nil {
			nextClerkCh = BusClerk(busNumber + 1)
		}
		return nextClerkCh(i)
	}
}

func Hilbert(upTo int) []WelcomeKit {
	var wellcomeKits = []WelcomeKit{}
	var firstBusClerk = BusClerk(1)
	for i := 1; i <= upTo; i++ {
		kit := firstBusClerk(i)
		wellcomeKits = append(wellcomeKits, kit)
	}

	fmt.Println()
	fmt.Printf("%v guests have been given a room by Hilber at his Hotel\n", upTo)

	return wellcomeKits
}
