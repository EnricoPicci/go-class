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

func RoomKeysClerk(upTo int, keysCh chan<- int) {
	for i := 0; i < upTo; i++ {
		keysCh <- i + 1
	}
	close(keysCh)
}

func BusClerk(busNumber int, roomKeysCh <-chan int, welcomeKitCh chan<- WelcomeKit) {
	var count = 0
	var passengerNumber = 1
	var nextClerkCh chan int

	for roomKey := range roomKeysCh {
		count++
		if nextClerkCh == nil {
			nextClerkCh = make(chan int)
			go BusClerk(busNumber+1, nextClerkCh, welcomeKitCh)
		}
		if count == passengerNumber {
			welcomeKitCh <- NewWelcomeKit(busNumber, passengerNumber, roomKey)
			passengerNumber++
			count = 0
			continue
		}
		nextClerkCh <- roomKey
	}

	if nextClerkCh != nil {
		close(nextClerkCh)
	} else {
		close(welcomeKitCh)
	}
}

func GoHilbert(upTo int) {
	keysCh := make(chan int)
	go RoomKeysClerk(upTo, keysCh)

	welcomeKitCh := make(chan WelcomeKit)
	go BusClerk(1, keysCh, welcomeKitCh)

	guestCounter := 0
	for envelope := range welcomeKitCh {
		guestCounter++
		fmt.Println(envelope)
	}

	fmt.Println()
	fmt.Printf("%v guests have been given a room by Hilber at his Hotel\n", guestCounter)
}
