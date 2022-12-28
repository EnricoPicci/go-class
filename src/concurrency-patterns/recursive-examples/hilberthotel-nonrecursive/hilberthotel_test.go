package main

import (
	"sort"
	"testing"
)

func TestWelcomeEnvelopesForBus(t *testing.T) {
	envelopes := WelcomeEnvelopes(11)

	expectedEnvelopes := []WelcomeEnvelope{
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

	expectedNumOfEnvelopes := len(expectedEnvelopes)
	gotNumOfEnvelopes := len(envelopes)
	if gotNumOfEnvelopes != expectedNumOfEnvelopes {
		t.Errorf("expected %v, got %v", expectedNumOfEnvelopes, gotNumOfEnvelopes)
	}

	sort.Slice(envelopes, func(i, j int) bool {
		if envelopes[i].BusNumber == envelopes[j].BusNumber {
			return envelopes[i].PassengerNumber < envelopes[j].PassengerNumber
		}
		return envelopes[i].BusNumber < envelopes[j].BusNumber
	})

	for i, gotEnvelop := range envelopes {
		expectedEnvelop := expectedEnvelopes[i]
		if gotEnvelop.RoomNumber != expectedEnvelop.RoomNumber {
			t.Errorf("Room number in envelope %v ==> expected %v - got %v", i, expectedEnvelop, gotEnvelop)
		}
	}
}

func TestGoHilbert(t *testing.T) {
	GoHilbert(16)
}
