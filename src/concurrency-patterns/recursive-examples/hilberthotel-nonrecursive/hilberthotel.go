package main

import "fmt"

type WelcomeEnvelope struct {
	BusNumber       int
	PassengerNumber int
	RoomNumber      int
}

func NewWelcomeEnvelope(busNumber int, passengerNumber int, row []int) WelcomeEnvelope {
	return WelcomeEnvelope{busNumber, passengerNumber, row[len(row)-busNumber]}
}
func (e WelcomeEnvelope) String() string {
	return fmt.Sprintf("Bus %v - Passenger %v - Room %v", e.BusNumber, e.PassengerNumber, e.RoomNumber)
}

func WelcomeEnvelopes(upTo int) []WelcomeEnvelope {
	rows := [][]int{}
	counter := 0
	i := 0
	for {
		counter++
		var row []int
		for j := 0; j < counter; j++ {
			if i+j == upTo {
				if row != nil {
					rows = append(rows, row)
				}
				goto rowsReady
			}
			if row == nil {
				row = make([]int, counter)
			}
			row[j] = i + j + 1
		}
		rows = append(rows, row)
		i = i + counter
	}
rowsReady:

	welcomeEnvelopes := []WelcomeEnvelope{}
	rowNumber := 1
	var passengerNumbersForBus []int

	for _, row := range rows {
		passengerNumbersForBus = append(passengerNumbersForBus, 1)
		for busNumber := 0; busNumber < rowNumber; busNumber++ {
			passengerNumber := passengerNumbersForBus[busNumber]
			welcomeEnvelope := NewWelcomeEnvelope(busNumber+1, passengerNumber, row)
			if welcomeEnvelope.RoomNumber > 0 {
				welcomeEnvelopes = append(welcomeEnvelopes, welcomeEnvelope)
			}
			passengerNumbersForBus[busNumber]++
		}
		rowNumber++
	}

	return welcomeEnvelopes
}

func GoHilbert(upTo int) {
	envelopes := WelcomeEnvelopes(upTo)

	for _, envelope := range envelopes {
		fmt.Println(envelope)
	}

	fmt.Println()
	fmt.Printf("%v guests have been given a room by Hilber at his Hotel", len(envelopes))
}
