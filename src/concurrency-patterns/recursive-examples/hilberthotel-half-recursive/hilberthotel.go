package main

import "fmt"

type WelcomeEnvelope struct {
	BusNumber       int
	PassengerNumber int
	RoomNumber      int
}

func (e WelcomeEnvelope) String() string {
	return fmt.Sprintf("Bus %v - Passenger %v - Room %v", e.BusNumber, e.PassengerNumber, e.RoomNumber)
}

func TriangleRowGenerator(upTo int) chan []int {
	ch := make(chan []int)

	go func(_upTo int, _ch chan []int) {
		counter := 0
		i := 0
		for {
			counter++
			var row []int
			for j := 0; j < counter; j++ {
				if i+j == _upTo {
					if row != nil {
						_ch <- row
					}
					goto ret
				}
				if row == nil {
					row = make([]int, counter)
				}
				row[j] = i + j + 1
			}
			_ch <- row
			i = i + counter
		}
	ret:
		close(_ch)
	}(upTo, ch)

	return ch
}

func NewWelcomeEnvelope(busNumber int, passengerNumber int, row []int) WelcomeEnvelope {
	return WelcomeEnvelope{busNumber, passengerNumber, row[len(row)-busNumber]}
}

func WelcomeEnvelopesForBus(rowsCh chan []int, hilbertCh chan WelcomeEnvelope, busNumber int) {
	passengerNumber := 1

	row, more := <-rowsCh
	if !more {
		close(hilbertCh)
		return
	}
	passengerNumber = GiveToHilbert(busNumber, passengerNumber, row, hilbertCh)

	nextBusRowsCh := make(chan []int)
	nextBusNumber := busNumber + 1
	go WelcomeEnvelopesForBus(nextBusRowsCh, hilbertCh, nextBusNumber)
	for row := range rowsCh {
		passengerNumber = GiveToHilbert(busNumber, passengerNumber, row, hilbertCh)
		nextBusRowsCh <- row
	}

	close(nextBusRowsCh)
}

func GiveToHilbert(busNumber int, passengerNumber int, row []int, hilbertCh chan WelcomeEnvelope) int {
	welcomeEnvelope := NewWelcomeEnvelope(busNumber, passengerNumber, row)
	if welcomeEnvelope.RoomNumber == 0 {
		return -1
	}
	hilbertCh <- welcomeEnvelope
	passengerNumber++
	return passengerNumber
}

func GoHilbert(upTo int) {
	rowsCh := TriangleRowGenerator(upTo)
	hilbertCh := make(chan WelcomeEnvelope)
	firstBusNumber := 1

	guestCounter := 0

	go WelcomeEnvelopesForBus(rowsCh, hilbertCh, firstBusNumber)

	for envelope := range hilbertCh {
		guestCounter++
		fmt.Println(envelope)
	}

	fmt.Println()
	fmt.Printf("%v guests have been given a room by Hilber at his Hotel", guestCounter)
}
