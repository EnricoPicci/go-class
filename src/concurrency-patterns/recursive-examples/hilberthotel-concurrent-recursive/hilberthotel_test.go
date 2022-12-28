package main

import (
	"sort"
	"testing"
)

func TestRoomKeysClerk_10(t *testing.T) {
	keysCh := make(chan int)
	go RoomKeysClerk(10, keysCh)

	roomNumbers := []int{}

	for n := range keysCh {
		roomNumbers = append(roomNumbers, n)
	}

	var expectedNums = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for i, r := range roomNumbers {
		if r != expectedNums[i] {
			t.Errorf("Room numbers ==> expcted %v - got %v", expectedNums[i], r)
		}
	}

}

func TestHilbertHospitality(t *testing.T) {
	keysCh := make(chan int)
	go RoomKeysClerk(11, keysCh)

	hilbertCh := make(chan WelcomeKit)
	go BusClerk(1, keysCh, hilbertCh)

	kits := []WelcomeKit{}
	for kit := range hilbertCh {
		kits = append(kits, kit)
	}

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

func TestGoHilbert(t *testing.T) {
	GoHilbert(16)
}
