package main

import (
	"sort"
	"testing"
)

func TestHilbertHospitality(t *testing.T) {

	kits := Hilbert(11)

	expectedKits := []WelcomeKit{
		{1, 1, 1},
		{1, 2, 3},
		{1, 3, 6},
		{1, 4, 10},
		{2, 1, 2},
		{2, 2, 5},
		{2, 3, 9},
		{3, 1, 4},
		{3, 2, 8},
		{4, 1, 7},
		{5, 1, 11},
	}

	expectedNumOfKits := len(expectedKits)
	gotNumOfKits := len(kits)
	if gotNumOfKits != expectedNumOfKits {
		t.Errorf("expected %v, got %v", expectedNumOfKits, gotNumOfKits)
	}

	sort.Slice(kits, func(i, j int) bool {
		if kits[i].BusNumber == kits[j].BusNumber {
			return kits[i].PassengerNumber < kits[j].PassengerNumber
		}
		return kits[i].BusNumber < kits[j].BusNumber
	})

	for i, gotEnvelop := range kits {
		expectedEnvelop := expectedKits[i]
		if gotEnvelop.RoomNumber != expectedEnvelop.RoomNumber {
			t.Errorf("Room number in envelope %v ==> expected %v - got %v", i, expectedEnvelop, gotEnvelop)
		}
	}
}
