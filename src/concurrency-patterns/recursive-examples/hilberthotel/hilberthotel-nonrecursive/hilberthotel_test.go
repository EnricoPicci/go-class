package main

import (
	"sort"
	"testing"

	"github.com/EnricoPicci/go-class/src/concurrency-patterns/recursive-examples/hilberthotel"
)

func TestWelcomeKitsForBus(t *testing.T) {
	Kits := WelcomeKits(11)

	expectedKits := []hilberthotel.WelcomeKit{
		{BusNumber: 1, PassengerNumber: 1, RoomNumber: 1},
		{BusNumber: 1, PassengerNumber: 2, RoomNumber: 3},
		{BusNumber: 1, PassengerNumber: 3, RoomNumber: 6},
		{BusNumber: 1, PassengerNumber: 4, RoomNumber: 10},
		{BusNumber: 2, PassengerNumber: 1, RoomNumber: 2},
		{BusNumber: 2, PassengerNumber: 2, RoomNumber: 5},
		{BusNumber: 2, PassengerNumber: 3, RoomNumber: 9},
		{BusNumber: 3, PassengerNumber: 1, RoomNumber: 4},
		{BusNumber: 3, PassengerNumber: 2, RoomNumber: 8},
		{BusNumber: 4, PassengerNumber: 1, RoomNumber: 7},
		{BusNumber: 5, PassengerNumber: 1, RoomNumber: 11},
	}

	expectedNumOfKits := len(expectedKits)
	gotNumOfKits := len(Kits)
	if gotNumOfKits != expectedNumOfKits {
		t.Errorf("expected %v, got %v", expectedNumOfKits, gotNumOfKits)
	}

	sort.Slice(Kits, func(i, j int) bool {
		if Kits[i].BusNumber == Kits[j].BusNumber {
			return Kits[i].PassengerNumber < Kits[j].PassengerNumber
		}
		return Kits[i].BusNumber < Kits[j].BusNumber
	})

	for i, gotEnvelop := range Kits {
		expectedEnvelop := expectedKits[i]
		if gotEnvelop.RoomNumber != expectedEnvelop.RoomNumber {
			t.Errorf("Room number in Kit %v ==> expected %v - got %v", i, expectedEnvelop, gotEnvelop)
		}
	}
}

func TestGoHilbertMassive(t *testing.T) {
	numOfPassengers := 1000000

	GoHilbert(numOfPassengers)
}
