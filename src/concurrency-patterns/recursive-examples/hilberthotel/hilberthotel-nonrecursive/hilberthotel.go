package hilberthotelnonrecursive

import (
	"fmt"

	"github.com/EnricoPicci/go-class/src/concurrency-patterns/recursive-examples/hilberthotel"
)

func NewWelcomeKit(busNumber int, passengerNumber int, row []int) hilberthotel.WelcomeKit {
	return hilberthotel.NewWelcomeKit(busNumber, passengerNumber, row[len(row)-busNumber])
}

func WelcomeKits(upTo int) []hilberthotel.WelcomeKit {
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

	welcomeKits := []hilberthotel.WelcomeKit{}
	rowNumber := 1
	var passengerNumbersForBus []int

	for _, row := range rows {
		passengerNumbersForBus = append(passengerNumbersForBus, 1)
		for busNumber := 0; busNumber < rowNumber; busNumber++ {
			passengerNumber := passengerNumbersForBus[busNumber]
			welcomeKit := NewWelcomeKit(busNumber+1, passengerNumber, row)
			if welcomeKit.RoomNumber > 0 {
				welcomeKits = append(welcomeKits, welcomeKit)
			}
			passengerNumbersForBus[busNumber]++
		}
		rowNumber++
	}

	return welcomeKits
}

func GoHilbert(upTo int, verbose bool) []hilberthotel.WelcomeKit {
	kits := WelcomeKits(upTo)

	if verbose {
		fmt.Println()
		fmt.Printf("%v guests have been given a room by Hilber at his Hotel", len(kits))
	}

	return kits
}
